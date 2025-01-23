package main

import (
	"database/sql"
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

// Category represents a book category
type Category struct {
	ID          string    `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	CreatedAt   time.Time `json:"createdAt"`
}

// Tag represents a book tag
type Tag struct {
	ID        string    `json:"id"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"createdAt"`
}

// Request/Response structures
type UpdateProgressRequest struct {
	CurrentPage int     `json:"currentPage"`
	TotalPages  int     `json:"totalPages"`
	Completion  float64 `json:"completion"`
}

type CreateBookmarkRequest struct {
	PageNumber int    `json:"pageNumber"`
	Note       string `json:"note"`
}

type GenerateAudioRequest struct {
	BookID    string `json:"bookId"`
	StartPage int    `json:"startPage"`
	EndPage   int    `json:"endPage"`
}

// UpdateBook updates an existing book in the database
func UpdateBook(db *sql.DB, book *Book) error {
	query := `
		UPDATE books 
		SET title = ?, author = ?, file_url = ?, 
			page_count = ?, language = ?, updated_at = ?
		WHERE id = ?`

	_, err := db.Exec(query,
		book.Title, book.Author, book.FileURL,
		book.PageCount, book.Language, book.UpdatedAt,
		book.ID)

	return err
}

// SaveAudioSegment saves a new audio segment to the database
func SaveAudioSegment(db *sql.DB, segment *AudioSegment) error {
	query := `
		INSERT INTO audio_segments (
			id, book_id, content, audio_url,
			status, created_at, updated_at
		) VALUES (?, ?, ?, ?, ?, ?, ?)`

	_, err := db.Exec(query,
		segment.ID,
		segment.BookID,
		segment.Content,
		segment.AudioURL,
		segment.Status,
		segment.CreatedAt,
		segment.UpdatedAt,
	)

	return err
}
