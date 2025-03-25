package tts

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"time"

	"backend/config"
	"backend/domain/models"
	"backend/repository/sqlite"

	"github.com/google/uuid"
)

// Generator handles text-to-speech generation
type Generator struct {
	config *config.Config
	db     *sqlite.DB
}

// NewGenerator creates a new TTS generator
func NewGenerator(cfg *config.Config, db *sqlite.DB) *Generator {
	return &Generator{
		config: cfg,
		db:     db,
	}
}

// GenerateAudio generates audio for the given text
func (g *Generator) GenerateAudio(text string) ([]byte, error) {
	text = strings.TrimSpace(text)
	if text == "" {
		return nil, fmt.Errorf("cannot generate audio for empty text")
	}

	log.Printf("[TTS] Starting TTS generation for text length: %d", len(text))

	// Generate TTS using the helper function
	audioURL, err := g.generateTTS(text)
	if err != nil {
		return nil, fmt.Errorf("error generating TTS: %v", err)
	}

	// Download the audio file
	audioResp, err := http.Get(audioURL)
	if err != nil {
		return nil, fmt.Errorf("error downloading audio: %v", err)
	}
	defer audioResp.Body.Close()

	return io.ReadAll(audioResp.Body)
}

// generateTTS generates audio for the given text using Kokoro TTS
func (g *Generator) generateTTS(text string) (string, error) {
	// Call Replicate API to generate audio
	replicateURL := g.config.ReplicateAPIURL + "/predictions"

	requestBody := map[string]interface{}{
		"version": g.config.KokoroModelVersion,
		"input": map[string]interface{}{
			"text":     text,
			"language": "en",
		},
	}

	jsonData, err := json.Marshal(requestBody)
	if err != nil {
		return "", fmt.Errorf("error marshaling request: %v", err)
	}

	// Create initial prediction
	req, err := http.NewRequest("POST", replicateURL, bytes.NewBuffer(jsonData))
	if err != nil {
		return "", fmt.Errorf("error creating request: %v", err)
	}

	req.Header.Set("Authorization", "Token "+os.Getenv("REPLICATE_API_TOKEN"))
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", fmt.Errorf("error making request: %v", err)
	}
	defer resp.Body.Close()

	var prediction struct {
		ID     string   `json:"id"`
		Status string   `json:"status"`
		Output []string `json:"output"`
	}
	if err := json.NewDecoder(resp.Body).Decode(&prediction); err != nil {
		return "", fmt.Errorf("error decoding response: %v", err)
	}

	// Poll for completion
	maxAttempts := 60 // 2 minutes total
	for i := 0; i < maxAttempts; i++ {
		time.Sleep(2 * time.Second)

		req, err = http.NewRequest("GET", replicateURL+"/"+prediction.ID, nil)
		if err != nil {
			return "", fmt.Errorf("error creating poll request: %v", err)
		}
		req.Header.Set("Authorization", "Token "+os.Getenv("REPLICATE_API_TOKEN"))

		resp, err = client.Do(req)
		if err != nil {
			return "", fmt.Errorf("error polling prediction: %v", err)
		}

		var result struct {
			Status string   `json:"status"`
			Output []string `json:"output"`
			Error  string   `json:"error"`
		}
		if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
			resp.Body.Close()
			return "", fmt.Errorf("error decoding poll response: %v", err)
		}
		resp.Body.Close()

		log.Printf("[TTS] Poll status: %s", result.Status)

		switch result.Status {
		case "succeeded", "completed":
			if len(result.Output) == 0 {
				return "", fmt.Errorf("no output from model")
			}
			return result.Output[0], nil

		case "failed":
			return "", fmt.Errorf("prediction failed: %s", result.Error)

		case "canceled":
			return "", fmt.Errorf("prediction was canceled")
		}
	}

	return "", fmt.Errorf("prediction timed out after %d attempts", maxAttempts)
}

// ProcessAudioSegment processes a text segment and generates audio
func (g *Generator) ProcessAudioSegment(segment *models.AudioSegment) ([]byte, error) {
	// Get Kokoro model version from environment
	modelVersion := os.Getenv("KOKORO_MODEL_VERSION")
	if modelVersion == "" {
		return nil, fmt.Errorf("KOKORO_MODEL_VERSION not set")
	}

	// Generate audio
	audioData, err := g.GenerateAudio(segment.Content)
	if err != nil {
		return nil, fmt.Errorf("error generating audio: %v", err)
	}

	// Save audio file
	audioFileName := fmt.Sprintf("tts-%s.mp3", segment.ID)
	audioPath := fmt.Sprintf("/audio/%s", audioFileName)

	// Update segment
	segment.AudioURL = audioPath
	segment.Status = "completed"
	segment.UpdatedAt = time.Now()

	return audioData, nil
}

// ProcessTextToSpeech processes all audio segments for a book
func (g *Generator) ProcessTextToSpeech(book *models.Book) error {
	segments, err := g.db.GetAudioSegments(book.ID)
	if err != nil {
		return fmt.Errorf("error getting audio segments: %v", err)
	}

	for _, segment := range segments {
		if segment.Status != "pending" {
			continue
		}

		audioData, err := g.GenerateAudio(segment.Content)
		if err != nil {
			segment.Status = "error"
			g.db.UpdateAudioSegment(&segment)
			continue
		}

		// Save audio to temporary file for immediate playback
		audioFileName := fmt.Sprintf("tts-%s.mp3", uuid.New().String())
		audioPath := filepath.Join(g.config.UploadDir, "audio", audioFileName)
		if err := os.WriteFile(audioPath, audioData, 0644); err != nil {
			segment.Status = "error"
			g.db.UpdateAudioSegment(&segment)
			continue
		}

		// Update segment with local URL
		segment.AudioURL = "/audio/" + audioFileName
		segment.Status = "completed"
		segment.UpdatedAt = time.Now()
		if err := g.db.UpdateAudioSegment(&segment); err != nil {
			continue
		}
	}

	return nil
}
