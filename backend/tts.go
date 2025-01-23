package main

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"os"
	"time"
)

// TTSRequest represents a request to the Replicate API
type TTSRequest struct {
	Version string   `json:"version"`
	Input   TTSInput `json:"input"`
}

// TTSInput represents the input parameters for the TTS model
type TTSInput struct {
	Text     string `json:"text"`
	Language string `json:"language"`
}

// TTSResponse represents a response from the Replicate API
type TTSResponse struct {
	ID     string `json:"id"`
	Status string `json:"status"`
	Output string `json:"output"`
	Error  string `json:"error"`
}

// processAudioSegment processes a text segment and generates audio
func processAudioSegment(db *sql.DB, segment *AudioSegment) error {
	// Get Kokoro model version from environment
	modelVersion := os.Getenv("KOKORO_MODEL_VERSION")
	if modelVersion == "" {
		return fmt.Errorf("KOKORO_MODEL_VERSION not set")
	}

	// Prepare request
	reqBody := TTSRequest{
		Version: modelVersion,
		Input: TTSInput{
			Text:     segment.Content,
			Language: "en", // Default to English for now
		},
	}

	// Convert request to JSON
	jsonData, err := json.Marshal(reqBody)
	if err != nil {
		return fmt.Errorf("error marshaling request: %v", err)
	}

	// Create HTTP request
	req, err := http.NewRequest("POST", os.Getenv("REPLICATE_API_URL")+"/predictions", bytes.NewBuffer(jsonData))
	if err != nil {
		return fmt.Errorf("error creating request: %v", err)
	}

	// Set headers
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Token "+os.Getenv("REPLICATE_API_TOKEN"))

	// Send request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return fmt.Errorf("error sending request: %v", err)
	}
	defer resp.Body.Close()

	// Read response
	var ttsResp TTSResponse
	if err := json.NewDecoder(resp.Body).Decode(&ttsResp); err != nil {
		return fmt.Errorf("error decoding response: %v", err)
	}

	// Check for error
	if ttsResp.Error != "" {
		return fmt.Errorf("TTS error: %s", ttsResp.Error)
	}

	// Poll for completion
	for ttsResp.Status != "succeeded" {
		time.Sleep(2 * time.Second)

		// Get prediction status
		req, err = http.NewRequest("GET", os.Getenv("REPLICATE_API_URL")+"/predictions/"+ttsResp.ID, nil)
		if err != nil {
			return fmt.Errorf("error creating status request: %v", err)
		}

		req.Header.Set("Authorization", "Token "+os.Getenv("REPLICATE_API_TOKEN"))

		resp, err = client.Do(req)
		if err != nil {
			return fmt.Errorf("error checking status: %v", err)
		}

		if err := json.NewDecoder(resp.Body).Decode(&ttsResp); err != nil {
			resp.Body.Close()
			return fmt.Errorf("error decoding status: %v", err)
		}
		resp.Body.Close()

		if ttsResp.Status == "failed" {
			return fmt.Errorf("TTS processing failed: %s", ttsResp.Error)
		}
	}

	// Download audio file from Replicate
	resp, err = http.Get(ttsResp.Output)
	if err != nil {
		return fmt.Errorf("error downloading audio: %v", err)
	}
	defer resp.Body.Close()

	// Create multipart form data for UploadThing
	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)

	// Create form file
	part, err := writer.CreateFormFile("file", fmt.Sprintf("%s.mp3", segment.ID))
	if err != nil {
		return fmt.Errorf("error creating form file: %v", err)
	}

	// Copy audio data to form file
	if _, err := io.Copy(part, resp.Body); err != nil {
		return fmt.Errorf("error copying audio data: %v", err)
	}

	// Close multipart writer
	writer.Close()

	// Create request to UploadThing
	uploadReq, err := http.NewRequest("POST", os.Getenv("UPLOADTHING_URL")+"/upload", body)
	if err != nil {
		return fmt.Errorf("error creating upload request: %v", err)
	}

	// Set headers for UploadThing
	uploadReq.Header.Set("Content-Type", writer.FormDataContentType())
	uploadReq.Header.Set("Authorization", "Bearer "+os.Getenv("UPLOADTHING_TOKEN"))

	// Send upload request
	uploadResp, err := client.Do(uploadReq)
	if err != nil {
		return fmt.Errorf("error uploading to UploadThing: %v", err)
	}
	defer uploadResp.Body.Close()

	if uploadResp.StatusCode != http.StatusOK {
		return fmt.Errorf("UploadThing error: %s", uploadResp.Status)
	}

	// Parse UploadThing response
	var uploadResult struct {
		URL string `json:"url"`
	}
	if err := json.NewDecoder(uploadResp.Body).Decode(&uploadResult); err != nil {
		return fmt.Errorf("error decoding UploadThing response: %v", err)
	}

	// Update segment with UploadThing URL and status
	segment.AudioURL = uploadResult.URL
	segment.Status = "completed"

	// Update database
	if err := UpdateAudioSegment(db, segment); err != nil {
		return fmt.Errorf("error updating segment: %v", err)
	}

	return nil
}

// generateTTS generates audio for the given text using Kokoro TTS
func generateTTS(text string) (string, error) {
	// Prepare request to Replicate API
	reqBody := map[string]interface{}{
		"version": os.Getenv("KOKORO_MODEL_VERSION"),
		"input": map[string]interface{}{
			"text":       text,
			"language":   "en",
			"speaker_id": 0,
		},
	}

	jsonData, err := json.Marshal(reqBody)
	if err != nil {
		return "", fmt.Errorf("error marshaling request: %v", err)
	}

	// Create request
	req, err := http.NewRequest("POST", os.Getenv("REPLICATE_API_URL")+"/predictions", bytes.NewBuffer(jsonData))
	if err != nil {
		return "", fmt.Errorf("error creating request: %v", err)
	}

	// Set headers
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Token "+os.Getenv("REPLICATE_API_TOKEN"))

	// Send request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", fmt.Errorf("error sending request: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusCreated {
		body, _ := io.ReadAll(resp.Body)
		return "", fmt.Errorf("unexpected status code %d: %s", resp.StatusCode, body)
	}

	// Parse response
	var prediction struct {
		ID     string `json:"id"`
		Status string `json:"status"`
		Output string `json:"output"`
	}
	if err := json.NewDecoder(resp.Body).Decode(&prediction); err != nil {
		return "", fmt.Errorf("error decoding response: %v", err)
	}

	// Poll for completion
	for prediction.Status != "succeeded" {
		time.Sleep(1 * time.Second)

		req, err = http.NewRequest("GET", os.Getenv("REPLICATE_API_URL")+"/predictions/"+prediction.ID, nil)
		if err != nil {
			return "", fmt.Errorf("error creating poll request: %v", err)
		}
		req.Header.Set("Authorization", "Token "+os.Getenv("REPLICATE_API_TOKEN"))

		resp, err = client.Do(req)
		if err != nil {
			return "", fmt.Errorf("error polling: %v", err)
		}
		defer resp.Body.Close()

		if err := json.NewDecoder(resp.Body).Decode(&prediction); err != nil {
			return "", fmt.Errorf("error decoding poll response: %v", err)
		}

		if prediction.Status == "failed" {
			return "", fmt.Errorf("TTS generation failed")
		}
	}

	// Download the audio file from Replicate
	audioResp, err := http.Get(prediction.Output)
	if err != nil {
		return "", fmt.Errorf("error downloading audio: %v", err)
	}
	defer audioResp.Body.Close()

	// Create multipart form data for UploadThing
	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)

	// Create form file
	part, err := writer.CreateFormFile("file", fmt.Sprintf("%s.mp3", generateID()))
	if err != nil {
		return "", fmt.Errorf("error creating form file: %v", err)
	}

	// Copy audio data to form file
	if _, err := io.Copy(part, audioResp.Body); err != nil {
		return "", fmt.Errorf("error copying audio data: %v", err)
	}

	// Close multipart writer
	writer.Close()

	// Create request to UploadThing
	uploadReq, err := http.NewRequest("POST", os.Getenv("UPLOADTHING_URL")+"/upload", body)
	if err != nil {
		return "", fmt.Errorf("error creating upload request: %v", err)
	}

	// Set headers for UploadThing
	uploadReq.Header.Set("Content-Type", writer.FormDataContentType())
	uploadReq.Header.Set("Authorization", "Bearer "+os.Getenv("UPLOADTHING_TOKEN"))

	// Send upload request
	uploadResp, err := client.Do(uploadReq)
	if err != nil {
		return "", fmt.Errorf("error uploading to UploadThing: %v", err)
	}
	defer uploadResp.Body.Close()

	if uploadResp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("UploadThing error: %s", uploadResp.Status)
	}

	// Parse UploadThing response
	var uploadResult struct {
		URL string `json:"url"`
	}
	if err := json.NewDecoder(uploadResp.Body).Decode(&uploadResult); err != nil {
		return "", fmt.Errorf("error decoding UploadThing response: %v", err)
	}

	return uploadResult.URL, nil
}
