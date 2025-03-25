package models

import (
	"time"
)

// Bookmark represents a user's bookmark in a book
type Bookmark struct {
	ID         string    `json:"id"`
	BookID     string    `json:"bookId"`
	UserID     string    `json:"userId"`
	PageNumber int       `json:"pageNumber"`
	Note       string    `json:"note"`
	CreatedAt  time.Time `json:"createdAt"`
	UpdatedAt  time.Time `json:"updatedAt"`
}

// CreateBookmarkRequest represents the request to create a bookmark
type CreateBookmarkRequest struct {
	PageNumber int    `json:"pageNumber"`
	Note       string `json:"note"`
}
