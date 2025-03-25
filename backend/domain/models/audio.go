package models

import (
	"time"
)

// AudioSegment represents a segment of text and its corresponding audio
type AudioSegment struct {
	ID        string    `json:"id"`
	BookID    string    `json:"bookId"`
	Content   string    `json:"content"`
	AudioURL  string    `json:"audioUrl"`
	Status    string    `json:"status"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

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
