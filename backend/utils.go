package main

import (
	"time"

	"github.com/google/uuid"
)

// Helper function to format datetime for SQLite
func formatDateTime(t time.Time) string {
	return t.UTC().Format("2006-01-02 15:04:05")
}

// generateID generates a unique ID using UUID
func generateID() string {
	return uuid.New().String()
}
