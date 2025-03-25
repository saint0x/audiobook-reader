package models

import (
	"time"
)

// Book represents a PDF book in the system
type Book struct {
	ID          string    `json:"id"`
	Title       string    `json:"title"`
	Author      string    `json:"author"`
	CoverURL    string    `json:"coverUrl"`
	FileURL     string    `json:"fileUrl"`
	PageCount   int       `json:"pageCount"`
	CurrentPage int       `json:"currentPage"`
	Language    string    `json:"language"`
	Status      string    `json:"status"`
	CreatedAt   time.Time `json:"createdAt"`
	UpdatedAt   time.Time `json:"updatedAt"`
	Categories  []string  `json:"categories,omitempty"`
	Tags        []string  `json:"tags,omitempty"`
}

// ReadingProgress tracks a user's reading progress for a book
type ReadingProgress struct {
	ID                string  `json:"id"`
	BookID            string  `json:"bookId"`
	UserID            string  `json:"userId"`
	CurrentPage       int     `json:"currentPage"`
	TotalPages        int     `json:"totalPages"`
	CompletionPercent float64 `json:"completionPercent"`
	LastReadAt        string  `json:"lastReadAt"`
}
