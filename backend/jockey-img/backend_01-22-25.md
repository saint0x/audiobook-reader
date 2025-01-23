# Jockey Image

Generated: 01-22-2025 at 11:12:42

## Repository Structure

```
backend
│   ├── models.go
│   ├── main.go
│   ├── config.go
│   ├── go.mod
│   ├── db.go
│   ├── pdf.go
│   ├── tts.go
│   ├── storage.go
│   ├── utils.go
│   ├── go.sum
│   ├── schema.sql
│   ├── README.md
    └── pdf-player
```

## File: /Users/saint/Desktop/pdf-player/backend/models.go

```go
package main

import "time"

// Book represents a PDF book in the system
type Book struct {
	ID          string    `json:"id"`
	Title       string    `json:"title"`
	Author      string    `json:"author"`
	CoverURL    string    `json:"coverUrl"`
	Content     string    `json:"content"`
	FilePath    string    `json:"filePath"`
	PageCount   int       `json:"pageCount"`
	CurrentPage int       `json:"currentPage"`
	Language    string    `json:"language"`
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

// AudioSegment represents a segment of text converted to audio
type AudioSegment struct {
	ID            string    `json:"id"`
	BookID        string    `json:"bookId"`
	SegmentNumber int       `json:"segmentNumber"`
	Content       string    `json:"content"`
	AudioURL      string    `json:"audioUrl"`
	Duration      float64   `json:"duration"`
	Status        string    `json:"status"`
	CreatedAt     time.Time `json:"createdAt"`
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
```

## File: /Users/saint/Desktop/pdf-player/backend/config.go

```go
package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/joho/godotenv"
)

type Config struct {
	// Server
	Port       string
	BackendURL string
	Env        string

	// Database
	DBPath string

	// File Storage
	UploadDir     string
	PDFDir        string
	CoverDir      string
	MaxUploadSize int64

	// Replicate API
	ReplicateAPIToken  string
	ReplicateAPIURL    string
	KokoroModelVersion string

	// TTS
	TTSMaxChunkSize   int
	TTSDefaultLang    string
	TTSDefaultSpeaker int
	TTSDefaultSpeed   float64
	TTSTopK           int
	TTSTopP           float64
	TTSTemperature    float64

	// CORS
	AllowedOrigins []string
}

var AppConfig Config

// LoadConfig loads the configuration from environment variables
func LoadConfig() error {
	// Load .env file if it exists
	godotenv.Load()

	AppConfig = Config{
		// Server
		Port:       getEnv("PORT", "8080"),
		BackendURL: getEnv("BACKEND_URL", "http://localhost:8080"),
		Env:        getEnv("ENV", "development"),

		// Database
		DBPath: getEnv("DB_PATH", "./ereader.db"),

		// File Storage
		UploadDir:     getEnv("UPLOAD_DIR", "./uploads"),
		PDFDir:        getEnv("PDF_DIR", "./uploads/pdfs"),
		CoverDir:      getEnv("COVER_DIR", "./uploads/covers"),
		MaxUploadSize: getEnvAsInt64("MAX_UPLOAD_SIZE", 32) * 1024 * 1024, // Convert MB to bytes

		// Replicate API
		ReplicateAPIToken:  getEnvRequired("REPLICATE_API_TOKEN"),
		ReplicateAPIURL:    getEnv("REPLICATE_API_URL", "https://api.replicate.com/v1"),
		KokoroModelVersion: getEnv("KOKORO_MODEL_VERSION", "82c9a78a6fff5fa0e4acbd0f8eb453e0c4b6d5d90e316b544b8c7e04c42fba42"),

		// TTS
		TTSMaxChunkSize:   getEnvAsInt("TTS_MAX_CHUNK_SIZE", 1000),
		TTSDefaultLang:    getEnv("TTS_DEFAULT_LANGUAGE", "en"),
		TTSDefaultSpeaker: getEnvAsInt("TTS_DEFAULT_SPEAKER_ID", 0),
		TTSDefaultSpeed:   getEnvAsFloat64("TTS_DEFAULT_SPEED", 1.0),
		TTSTopK:           getEnvAsInt("TTS_TOP_K", 50),
		TTSTopP:           getEnvAsFloat64("TTS_TOP_P", 0.8),
		TTSTemperature:    getEnvAsFloat64("TTS_TEMPERATURE", 0.8),

		// CORS
		AllowedOrigins: strings.Split(getEnv("ALLOWED_ORIGINS", "http://localhost:3000"), ","),
	}

	// Create required directories
	for _, dir := range []string{AppConfig.UploadDir, AppConfig.PDFDir, AppConfig.CoverDir} {
		if err := os.MkdirAll(dir, 0755); err != nil {
			return fmt.Errorf("error creating directory %s: %v", dir, err)
		}
	}

	return nil
}

// Helper functions to get environment variables
func getEnv(key, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultValue
}

func getEnvRequired(key string) string {
	value, exists := os.LookupEnv(key)
	if !exists {
		panic(fmt.Sprintf("Required environment variable %s is not set", key))
	}
	return value
}

func getEnvAsInt(key string, defaultValue int) int {
	valueStr := getEnv(key, "")
	if value, err := strconv.Atoi(valueStr); err == nil {
		return value
	}
	return defaultValue
}

func getEnvAsInt64(key string, defaultValue int64) int64 {
	valueStr := getEnv(key, "")
	if value, err := strconv.ParseInt(valueStr, 10, 64); err == nil {
		return value
	}
	return defaultValue
}

func getEnvAsFloat64(key string, defaultValue float64) float64 {
	valueStr := getEnv(key, "")
	if value, err := strconv.ParseFloat(valueStr, 64); err == nil {
		return value
	}
	return defaultValue
}
```

## File: /Users/saint/Desktop/pdf-player/backend/go.mod

```mod
module github.com/saint/pdf-player

go 1.21

require (
	github.com/gorilla/mux v1.8.1
	github.com/ledongthuc/pdf v0.0.0-20240201131950-da5b75280b06
	github.com/mattn/go-sqlite3 v1.14.24
)

require github.com/joho/godotenv v1.5.1

require github.com/google/uuid v1.6.0
```

## File: /Users/saint/Desktop/pdf-player/backend/db.go

```go
package main

import (
	"database/sql"
	"fmt"
	"os"
	"time"
)

// InitDB initializes the database connection and creates tables
func InitDB() error {
	var err error
	db, err = sql.Open("sqlite3", "./ereader.db")
	if err != nil {
		return fmt.Errorf("error opening database: %v", err)
	}

	// Read and execute schema.sql
	schema, err := os.ReadFile("schema.sql")
	if err != nil {
		return fmt.Errorf("error reading schema file: %v", err)
	}

	_, err = db.Exec(string(schema))
	if err != nil {
		return fmt.Errorf("error creating schema: %v", err)
	}

	return nil
}

// GetBookByID retrieves a book by its ID
func GetBookByID(db *sql.DB, id string) (*Book, error) {
	var book Book
	query := `
		SELECT id, title, author, cover_url, content, file_path, 
			page_count, current_page, language, 
			created_at, updated_at
		FROM books WHERE id = ?
	`

	err := db.QueryRow(query, id).Scan(
		&book.ID,
		&book.Title,
		&book.Author,
		&book.CoverURL,
		&book.Content,
		&book.FilePath,
		&book.PageCount,
		&book.CurrentPage,
		&book.Language,
		&book.CreatedAt,
		&book.UpdatedAt,
	)

	if err == sql.ErrNoRows {
		return nil, fmt.Errorf("book not found")
	}
	if err != nil {
		return nil, err
	}

	// Get categories
	rows, err := db.Query(`
		SELECT c.name FROM categories c 
		JOIN book_categories bc ON c.id = bc.category_id 
		WHERE bc.book_id = ?`, id)
	if err == nil {
		defer rows.Close()
		for rows.Next() {
			var category string
			rows.Scan(&category)
			book.Categories = append(book.Categories, category)
		}
	}

	// Get tags
	rows, err = db.Query(`
		SELECT t.name FROM tags t 
		JOIN book_tags bt ON t.id = bt.tag_id 
		WHERE bt.book_id = ?`, id)
	if err == nil {
		defer rows.Close()
		for rows.Next() {
			var tag string
			rows.Scan(&tag)
			book.Tags = append(book.Tags, tag)
		}
	}

	return &book, nil
}

// SaveBook saves a new book to the database
func SaveBook(db *sql.DB, book *Book) error {
	query := `
		INSERT INTO books (
			id, title, author, cover_url, content, file_path,
			page_count, current_page, language, created_at, updated_at
		) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)
	`

	_, err := db.Exec(query,
		book.ID,
		book.Title,
		book.Author,
		book.CoverURL,
		book.Content,
		book.FilePath,
		book.PageCount,
		book.CurrentPage,
		book.Language,
		formatDateTime(book.CreatedAt),
		formatDateTime(book.UpdatedAt),
	)

	return err
}

// UpdateReadingProgress updates or creates a reading progress record
func UpdateReadingProgress(db *sql.DB, progress *ReadingProgress) error {
	query := `
		INSERT INTO reading_progress (id, book_id, user_id, current_page, total_pages, completion_percent, last_read_at)
		VALUES (?, ?, ?, ?, ?, ?, CURRENT_TIMESTAMP)
		ON CONFLICT(book_id, user_id) DO UPDATE SET
			current_page = excluded.current_page,
			total_pages = excluded.total_pages,
			completion_percent = excluded.completion_percent,
			last_read_at = CURRENT_TIMESTAMP
	`

	_, err := db.Exec(query,
		progress.ID,
		progress.BookID,
		progress.UserID,
		progress.CurrentPage,
		progress.TotalPages,
		progress.CompletionPercent,
	)

	return err
}

// GetReadingProgress retrieves reading progress for a book and user
func GetReadingProgress(db *sql.DB, bookID, userID string) (*ReadingProgress, error) {
	var progress ReadingProgress
	query := `
		SELECT id, book_id, user_id, current_page, total_pages, completion_percent, last_read_at
		FROM reading_progress
		WHERE book_id = ? AND user_id = ?
	`

	err := db.QueryRow(query, bookID, userID).Scan(
		&progress.ID,
		&progress.BookID,
		&progress.UserID,
		&progress.CurrentPage,
		&progress.TotalPages,
		&progress.CompletionPercent,
		&progress.LastReadAt,
	)

	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}

	return &progress, nil
}

// CreateBookmark creates a new bookmark
func CreateBookmark(db *sql.DB, bookmark *Bookmark) error {
	query := `
		INSERT INTO bookmarks (id, book_id, user_id, page_number, note, created_at, updated_at)
		VALUES (?, ?, ?, ?, ?, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP)
	`

	_, err := db.Exec(query,
		bookmark.ID,
		bookmark.BookID,
		bookmark.UserID,
		bookmark.PageNumber,
		bookmark.Note,
	)

	return err
}

// GetBookmarks retrieves all bookmarks for a book and user
func GetBookmarks(db *sql.DB, bookID, userID string) ([]Bookmark, error) {
	query := `
		SELECT id, book_id, user_id, page_number, note, created_at, updated_at
		FROM bookmarks
		WHERE book_id = ? AND user_id = ?
		ORDER BY page_number ASC
	`

	rows, err := db.Query(query, bookID, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var bookmarks []Bookmark
	for rows.Next() {
		var bookmark Bookmark
		err := rows.Scan(
			&bookmark.ID,
			&bookmark.BookID,
			&bookmark.UserID,
			&bookmark.PageNumber,
			&bookmark.Note,
			&bookmark.CreatedAt,
			&bookmark.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}
		bookmarks = append(bookmarks, bookmark)
	}

	return bookmarks, nil
}

// DeleteBookmark deletes a bookmark
func DeleteBookmark(db *sql.DB, id string) error {
	query := "DELETE FROM bookmarks WHERE id = ?"
	_, err := db.Exec(query, id)
	return err
}

// UpdateBookmark updates a bookmark's note
func UpdateBookmark(db *sql.DB, bookmark *Bookmark) error {
	query := `
		UPDATE bookmarks 
		SET note = ?, updated_at = CURRENT_TIMESTAMP
		WHERE id = ? AND book_id = ? AND user_id = ?
	`

	result, err := db.Exec(query, bookmark.Note, bookmark.ID, bookmark.BookID, bookmark.UserID)
	if err != nil {
		return err
	}

	rows, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rows == 0 {
		return fmt.Errorf("bookmark not found")
	}

	return nil
}

// CreateAudioSegment creates a new audio segment
func CreateAudioSegment(db *sql.DB, segment *AudioSegment) error {
	query := `
		INSERT INTO audio_segments (
			id, book_id, segment_number, content, audio_url, duration, status, created_at
		) VALUES (?, ?, ?, ?, ?, ?, ?, CURRENT_TIMESTAMP)
	`

	_, err := db.Exec(query,
		segment.ID,
		segment.BookID,
		segment.SegmentNumber,
		segment.Content,
		segment.AudioURL,
		segment.Duration,
		segment.Status,
	)

	return err
}

// UpdateAudioSegment updates an existing audio segment
func UpdateAudioSegment(db *sql.DB, segment *AudioSegment) error {
	query := `
		UPDATE audio_segments 
		SET audio_url = ?, duration = ?, status = ?
		WHERE id = ? AND book_id = ?
	`

	result, err := db.Exec(query,
		segment.AudioURL,
		segment.Duration,
		segment.Status,
		segment.ID,
		segment.BookID,
	)
	if err != nil {
		return err
	}

	rows, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rows == 0 {
		return fmt.Errorf("audio segment not found")
	}

	return nil
}

// Category operations
func CreateCategory(db *sql.DB, category *Category) error {
	category.ID = generateID()
	category.CreatedAt = time.Now()

	_, err := db.Exec(`
		INSERT INTO categories (id, name, description, created_at)
		VALUES (?, ?, ?, ?)`,
		category.ID, category.Name, category.Description, category.CreatedAt)

	return err
}

func AddBookToCategory(db *sql.DB, bookID, categoryID string) error {
	_, err := db.Exec(`
		INSERT INTO book_categories (book_id, category_id)
		VALUES (?, ?)`, bookID, categoryID)

	return err
}

// Tag operations
func CreateTag(db *sql.DB, tag *Tag) error {
	tag.ID = generateID()
	tag.CreatedAt = time.Now()

	_, err := db.Exec(`
		INSERT INTO tags (id, name, created_at)
		VALUES (?, ?, ?)`,
		tag.ID, tag.Name, tag.CreatedAt)

	return err
}

func AddBookTag(db *sql.DB, bookID, tagID string) error {
	_, err := db.Exec(`
		INSERT INTO book_tags (book_id, tag_id)
		VALUES (?, ?)`, bookID, tagID)

	return err
}
```

## File: /Users/saint/Desktop/pdf-player/backend/schema.sql

```sql
-- Books table with enhanced metadata
CREATE TABLE IF NOT EXISTS books (
    id TEXT PRIMARY KEY,
    title TEXT NOT NULL,
    author TEXT NOT NULL,
    cover_url TEXT NOT NULL,
    content TEXT NOT NULL,
    file_path TEXT,
    page_count INTEGER DEFAULT 0,
    current_page INTEGER DEFAULT 0,
    language TEXT DEFAULT 'en',
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- Categories for organizing books
CREATE TABLE IF NOT EXISTS categories (
    id TEXT PRIMARY KEY,
    name TEXT NOT NULL UNIQUE,
    description TEXT,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- Book-Category relationship (many-to-many)
CREATE TABLE IF NOT EXISTS book_categories (
    book_id TEXT,
    category_id TEXT,
    PRIMARY KEY (book_id, category_id),
    FOREIGN KEY (book_id) REFERENCES books(id) ON DELETE CASCADE,
    FOREIGN KEY (category_id) REFERENCES categories(id) ON DELETE CASCADE
);

-- Tags for flexible book organization
CREATE TABLE IF NOT EXISTS tags (
    id TEXT PRIMARY KEY,
    name TEXT NOT NULL UNIQUE,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- Book-Tag relationship (many-to-many)
CREATE TABLE IF NOT EXISTS book_tags (
    book_id TEXT,
    tag_id TEXT,
    PRIMARY KEY (book_id, tag_id),
    FOREIGN KEY (book_id) REFERENCES books(id) ON DELETE CASCADE,
    FOREIGN KEY (tag_id) REFERENCES tags(id) ON DELETE CASCADE
);

-- Audio synthesis tracking
CREATE TABLE IF NOT EXISTS audio_segments (
    id TEXT PRIMARY KEY,
    book_id TEXT NOT NULL,
    segment_number INTEGER NOT NULL,
    content TEXT NOT NULL,
    audio_url TEXT,
    duration FLOAT,
    status TEXT CHECK (status IN ('pending', 'processing', 'completed', 'failed')),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (book_id) REFERENCES books(id) ON DELETE CASCADE
);

-- Reading progress tracking
CREATE TABLE IF NOT EXISTS reading_progress (
    id TEXT PRIMARY KEY,
    book_id TEXT NOT NULL,
    current_page INTEGER DEFAULT 0,
    completion_percentage FLOAT DEFAULT 0,
    last_read_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (book_id) REFERENCES books(id) ON DELETE CASCADE,
    UNIQUE(book_id)
);

-- Bookmarks
CREATE TABLE IF NOT EXISTS bookmarks (
    id TEXT PRIMARY KEY,
    book_id TEXT NOT NULL,
    page_number INTEGER NOT NULL,
    note TEXT,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (book_id) REFERENCES books(id) ON DELETE CASCADE
);

-- Create indexes for better query performance
CREATE INDEX IF NOT EXISTS idx_books_title ON books(title);
CREATE INDEX IF NOT EXISTS idx_reading_progress_book ON reading_progress(book_id);
CREATE INDEX IF NOT EXISTS idx_audio_segments_book ON audio_segments(book_id);
CREATE INDEX IF NOT EXISTS idx_bookmarks_book ON bookmarks(book_id); 
```

## File: /Users/saint/Desktop/pdf-player/backend/pdf.go

```go
package main

import (
	"fmt"
	"mime/multipart"
	"strings"
	"time"

	"github.com/google/uuid"
	"github.com/ledongthuc/pdf"
)

// processPDF processes a PDF file and returns a Book object
func processPDF(file multipart.File, filename string) (*Book, error) {
	// Save PDF file temporarily
	filePath, err := fileStorage.SavePDF(file, &multipart.FileHeader{
		Filename: filename,
	})
	if err != nil {
		return nil, fmt.Errorf("error saving PDF: %v", err)
	}

	// Open PDF file for reading
	f, r, err := pdf.Open(filePath)
	if err != nil {
		return nil, fmt.Errorf("error opening PDF: %v", err)
	}
	defer f.Close()

	// Get total number of pages
	totalPages := r.NumPage()

	// Extract text from all pages
	var content strings.Builder
	for pageNum := 1; pageNum <= totalPages; pageNum++ {
		page := r.Page(pageNum)
		if page.V.IsNull() {
			continue
		}

		text, err := page.GetPlainText(nil)
		if err != nil {
			return nil, fmt.Errorf("error extracting text from page %d: %v", pageNum, err)
		}
		content.WriteString(text)
	}

	// Create book object
	book := &Book{
		ID:          uuid.New().String(),
		Title:       strings.TrimSuffix(filename, ".pdf"),
		Author:      "Unknown", // Can be updated later
		FilePath:    filePath,
		Content:     content.String(),
		PageCount:   totalPages,
		CurrentPage: 0,
		Language:    "en", // Default to English, can be updated later
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	return book, nil
}
```

## File: /Users/saint/Desktop/pdf-player/backend/tts.go

```go
package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
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
func processAudioSegment(segment *AudioSegment) error {
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

	// Download audio file
	resp, err = http.Get(ttsResp.Output)
	if err != nil {
		return fmt.Errorf("error downloading audio: %v", err)
	}
	defer resp.Body.Close()

	// Create temporary file
	tmpFile, err := os.CreateTemp("", "tts-*.mp3")
	if err != nil {
		return fmt.Errorf("error creating temp file: %v", err)
	}
	defer os.Remove(tmpFile.Name())

	// Copy audio to temp file
	if _, err := io.Copy(tmpFile, resp.Body); err != nil {
		return fmt.Errorf("error saving audio: %v", err)
	}

	// Save audio file using storage
	audioFile, err := os.Open(tmpFile.Name())
	if err != nil {
		return fmt.Errorf("error opening temp file: %v", err)
	}
	defer audioFile.Close()

	// Save audio and get URL
	audioURL, err := fileStorage.SaveAudio(audioFile, "output.mp3")
	if err != nil {
		return fmt.Errorf("error saving audio file: %v", err)
	}

	// Update segment with audio URL and status
	segment.AudioURL = audioURL
	segment.Status = "completed"

	// Update database
	if err := UpdateAudioSegment(db, segment); err != nil {
		return fmt.Errorf("error updating segment: %v", err)
	}

	return nil
}
```

## File: /Users/saint/Desktop/pdf-player/backend/go.sum

```sum
github.com/google/uuid v1.6.0 h1:NIvaJDMOsjHA8n1jAhLSgzrAzy1Hgr+hNrb57e+94F0=
github.com/google/uuid v1.6.0/go.mod h1:TIyPZe4MgqvfeYDBFedMoGGpEw/LqOeaOT+nhxU+yHo=
github.com/gorilla/mux v1.8.1 h1:TuBL49tXwgrFYWhqrNgrUNEY92u81SPhu7sTdzQEiWY=
github.com/gorilla/mux v1.8.1/go.mod h1:AKf9I4AEqPTmMytcMc0KkNouC66V3BtZ4qD5fmWSiMQ=
github.com/joho/godotenv v1.5.1 h1:7eLL/+HRGLY0ldzfGMeQkb7vMd0as4CfYvUVzLqw0N0=
github.com/joho/godotenv v1.5.1/go.mod h1:f4LDr5Voq0i2e/R5DDNOoa2zzDfwtkZa6DnEwAbqwq4=
github.com/ledongthuc/pdf v0.0.0-20240201131950-da5b75280b06 h1:kacRlPN7EN++tVpGUorNGPn/4DnB7/DfTY82AOn6ccU=
github.com/ledongthuc/pdf v0.0.0-20240201131950-da5b75280b06/go.mod h1:imJHygn/1yfhB7XSJJKlFZKl/J+dCPAknuiaGOshXAs=
github.com/mattn/go-sqlite3 v1.14.24 h1:tpSp2G2KyMnnQu99ngJ47EIkWVmliIizyZBfPrBWDRM=
github.com/mattn/go-sqlite3 v1.14.24/go.mod h1:Uh1q+B4BYcTPb+yiD3kU8Ct7aC0hY9fxUwlHK0RXw+Y=
```

## File: /Users/saint/Desktop/pdf-player/backend/README.md

````md
# PDF Player Backend

This is the backend service for the PDF Player application. It provides APIs for uploading PDFs, extracting text, managing books, and tracking reading progress.

## Prerequisites

- Go 1.16 or later
- SQLite3

## Setup

1. Install dependencies:
```bash
go mod download
```

2. Run the server:
```bash
go run *.go
```

The server will start on port 8080 by default. You can change the port by setting the `PORT` environment variable.

## API Endpoints

### Books
- **POST** `/api/upload` - Upload a PDF file
  - Content-Type: `multipart/form-data`
  - Form field: `file` (PDF file)
  - Returns: Book object with extracted text

- **GET** `/api/books` - Get all books
  - Returns: Array of book objects

- **GET** `/api/book/{id}` - Get a specific book
  - Returns: Single book object

### Reading Progress
- **PUT** `/api/book/{id}/progress` - Update reading progress
  - Body: `{ "currentPage": number, "completion": number }`

- **GET** `/api/book/{id}/progress` - Get reading progress
  - Returns: Reading progress object

### Bookmarks
- **POST** `/api/book/{id}/bookmarks` - Create a bookmark
  - Body: `{ "pageNumber": number, "note": string }`
  - Returns: Created bookmark object

- **GET** `/api/book/{id}/bookmarks` - Get all bookmarks for a book
  - Returns: Array of bookmark objects

### Audio Segments
- **POST** `/api/book/{id}/audio` - Create an audio segment
  - Body: `{ "startPage": number, "endPage": number }`
  - Returns: Created audio segment object

- **PUT** `/api/book/{id}/audio/{segmentId}` - Update audio segment status
  - Body: `{ "status": string, "audioUrl": string, "duration": number }`

### Categories
- **POST** `/api/categories` - Create a category
  - Body: `{ "name": string, "description": string }`
  - Returns: Created category object

- **PUT** `/api/book/{id}/categories/{categoryId}` - Add book to category

### Tags
- **POST** `/api/tags` - Create a tag
  - Body: `{ "name": string }`
  - Returns: Created tag object

- **PUT** `/api/book/{id}/tags/{tagId}` - Add tag to book

## Data Models

### Book
```json
{
  "id": "string",
  "title": "string",
  "author": "string",
  "coverUrl": "string",
  "content": "string",
  "filePath": "string",
  "pageCount": number,
  "currentPage": number,
  "language": "string",
  "createdAt": "datetime",
  "updatedAt": "datetime",
  "categories": ["string"],
  "tags": ["string"]
}
```

### Reading Progress
```json
{
  "id": "string",
  "bookId": "string",
  "currentPage": number,
  "completionPercent": number,
  "lastReadAt": "datetime"
}
```

### Bookmark
```json
{
  "id": "string",
  "bookId": "string",
  "pageNumber": number,
  "note": "string",
  "createdAt": "datetime"
}
```

### Audio Segment
```json
{
  "id": "string",
  "bookId": "string",
  "segmentNumber": number,
  "content": "string",
  "audioUrl": "string",
  "duration": number,
  "status": "string",
  "createdAt": "datetime"
}
```

## Future Enhancements

1. PDF metadata extraction for better book information
2. Custom book cover image upload
3. Text-to-speech synthesis integration using Kokoro TTS
4. Book categories and tags for better organization
5. Full-text search functionality 
````

## File: /Users/saint/Desktop/pdf-player/backend/storage.go

```go
package main

import (
	"fmt"
	"io"
	"mime/multipart"
	"os"
	"path/filepath"
	"strings"

	"github.com/google/uuid"
)

// FileStorage handles file operations
type FileStorage struct {
	baseDir  string
	pdfDir   string
	coverDir string
	audioDir string
}

// NewFileStorage creates a new FileStorage instance
func NewFileStorage(baseDir string) (*FileStorage, error) {
	fs := &FileStorage{
		baseDir:  baseDir,
		pdfDir:   filepath.Join(baseDir, "pdfs"),
		coverDir: filepath.Join(baseDir, "covers"),
		audioDir: filepath.Join(baseDir, "audio"),
	}

	// Create directories if they don't exist
	for _, dir := range []string{fs.pdfDir, fs.coverDir, fs.audioDir} {
		if err := os.MkdirAll(dir, 0755); err != nil {
			return nil, fmt.Errorf("error creating directory %s: %v", dir, err)
		}
	}

	return fs, nil
}

// SavePDF saves a PDF file and returns its path
func (fs *FileStorage) SavePDF(file multipart.File, header *multipart.FileHeader) (string, error) {
	// Generate unique filename
	ext := filepath.Ext(header.Filename)
	if !strings.EqualFold(ext, ".pdf") {
		return "", fmt.Errorf("invalid file type: %s", ext)
	}

	filename := uuid.New().String() + ext
	filePath := filepath.Join(fs.pdfDir, filename)

	return fs.saveFile(file, filePath)
}

// SaveCover saves a cover image and returns its URL
func (fs *FileStorage) SaveCover(file multipart.File, filename string) (string, error) {
	// Generate unique filename
	ext := filepath.Ext(filename)
	if !strings.EqualFold(ext, ".jpg") && !strings.EqualFold(ext, ".jpeg") && !strings.EqualFold(ext, ".png") {
		return "", fmt.Errorf("invalid file type: %s", ext)
	}

	newFilename := uuid.New().String() + ext
	filePath := filepath.Join(fs.coverDir, newFilename)

	_, err := fs.saveFile(file, filePath)
	if err != nil {
		return "", err
	}

	// Return URL-friendly path
	return "/uploads/covers/" + newFilename, nil
}

// SaveAudio saves an audio file and returns its URL
func (fs *FileStorage) SaveAudio(file multipart.File, filename string) (string, error) {
	// Generate unique filename
	ext := filepath.Ext(filename)
	if !strings.EqualFold(ext, ".mp3") && !strings.EqualFold(ext, ".wav") {
		return "", fmt.Errorf("invalid file type: %s", ext)
	}

	newFilename := uuid.New().String() + ext
	filePath := filepath.Join(fs.audioDir, newFilename)

	_, err := fs.saveFile(file, filePath)
	if err != nil {
		return "", err
	}

	// Return URL-friendly path
	return "/uploads/audio/" + newFilename, nil
}

// saveFile is a helper function to save a file
func (fs *FileStorage) saveFile(file multipart.File, filePath string) (string, error) {
	// Create destination file
	dst, err := os.Create(filePath)
	if err != nil {
		return "", fmt.Errorf("error creating file: %v", err)
	}
	defer dst.Close()

	// Copy file contents
	if _, err := io.Copy(dst, file); err != nil {
		os.Remove(filePath) // Clean up on error
		return "", fmt.Errorf("error copying file: %v", err)
	}

	return filePath, nil
}

// DeleteFile deletes a file at the given path
func (fs *FileStorage) DeleteFile(filePath string) error {
	// Ensure the file is within our base directory
	if !strings.HasPrefix(filePath, fs.baseDir) {
		return fmt.Errorf("invalid file path: %s", filePath)
	}

	return os.Remove(filePath)
}

func (fs *FileStorage) GetFilePath(relativePath string) string {
	// Remove leading slash if present
	if len(relativePath) > 0 && relativePath[0] == '/' {
		relativePath = relativePath[1:]
	}

	return filepath.Join(fs.baseDir, relativePath)
}
```

## File: /Users/saint/Desktop/pdf-player/backend/utils.go

```go
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
```

## File: /Users/saint/Desktop/pdf-player/backend/main.go

```go
package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	_ "github.com/mattn/go-sqlite3"
)

// BookBasic represents the basic book information used in the initial implementation
type BookBasic struct {
	ID       string `json:"id"`
	Title    string `json:"title"`
	Author   string `json:"author"`
	CoverURL string `json:"coverUrl"`
	Content  string `json:"content"`
}

var db *sql.DB
var fileStorage *FileStorage

func main() {
	// Load environment variables
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}

	// Initialize database
	var err error
	db, err = sql.Open("sqlite3", "books.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Create tables
	if err := InitDB(); err != nil {
		log.Fatal(err)
	}

	// Initialize file storage
	fileStorage, err = NewFileStorage("uploads")
	if err != nil {
		log.Fatal(err)
	}

	// Create router
	router := mux.NewRouter()

	// Configure CORS
	router.Use(func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			allowedOrigins := strings.Split(os.Getenv("ALLOWED_ORIGINS"), ",")
			origin := r.Header.Get("Origin")
			for _, allowed := range allowedOrigins {
				if origin == allowed {
					w.Header().Set("Access-Control-Allow-Origin", origin)
					break
				}
			}
			w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
			w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
			if r.Method == "OPTIONS" {
				w.WriteHeader(http.StatusOK)
				return
			}
			next.ServeHTTP(w, r)
		})
	})

	// File upload routes
	router.HandleFunc("/api/upload/pdf", uploadPDFHandler).Methods("POST")
	router.HandleFunc("/api/upload/cover", uploadCoverHandler).Methods("POST")

	// Book routes
	router.HandleFunc("/api/books/{id}", getBookHandler).Methods("GET")
	router.HandleFunc("/api/books", getBooksHandler).Methods("GET")

	// Reading progress routes
	router.HandleFunc("/api/progress", updateProgressHandler).Methods("POST")
	router.HandleFunc("/api/progress/{bookId}", getProgressHandler).Methods("GET")

	// Bookmark routes
	router.HandleFunc("/api/bookmarks", createBookmarkHandler).Methods("POST")
	router.HandleFunc("/api/bookmarks/{bookId}", getBookmarksHandler).Methods("GET")
	router.HandleFunc("/api/bookmarks/{id}", updateBookmarkHandler).Methods("PUT")
	router.HandleFunc("/api/bookmarks/{id}", deleteBookmarkHandler).Methods("DELETE")

	// Audio segment routes
	router.HandleFunc("/api/audio/generate", generateAudioHandler).Methods("POST")
	router.HandleFunc("/api/audio/{bookId}", getAudioSegmentsHandler).Methods("GET")

	// Category and tag routes
	router.HandleFunc("/api/categories", getCategoriesHandler).Methods("GET")
	router.HandleFunc("/api/tags", getTagsHandler).Methods("GET")

	// Start server
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	log.Printf("Server starting on port %s", port)
	log.Fatal(http.ListenAndServe(":"+port, router))
}

func uploadPDFHandler(w http.ResponseWriter, r *http.Request) {
	// Parse multipart form with 32MB limit
	if err := r.ParseMultipartForm(32 << 20); err != nil {
		http.Error(w, "File too large", http.StatusBadRequest)
		return
	}

	file, header, err := r.FormFile("file")
	if err != nil {
		http.Error(w, "Error retrieving file", http.StatusBadRequest)
		return
	}
	defer file.Close()

	// Process PDF and save book
	book, err := processPDF(file, header.Filename)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error processing PDF: %v", err), http.StatusInternalServerError)
		return
	}

	// Save book to database
	if err := SaveBook(db, book); err != nil {
		http.Error(w, "Error saving book", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(book)
}

func uploadCoverHandler(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseMultipartForm(10 << 20); err != nil {
		http.Error(w, "File too large", http.StatusBadRequest)
		return
	}

	file, header, err := r.FormFile("cover")
	if err != nil {
		http.Error(w, "Error retrieving file", http.StatusBadRequest)
		return
	}
	defer file.Close()

	coverURL, err := fileStorage.SaveCover(file, header.Filename)
	if err != nil {
		http.Error(w, "Error saving cover", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(map[string]string{"coverUrl": coverURL})
}

func getBookHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	book, err := GetBookByID(db, vars["id"])
	if err != nil {
		http.Error(w, "Book not found", http.StatusNotFound)
		return
	}

	json.NewEncoder(w).Encode(book)
}

func getBooksHandler(w http.ResponseWriter, r *http.Request) {
	query := `
		SELECT id, title, author, cover_url, page_count, language, created_at
		FROM books
		ORDER BY created_at DESC
	`
	rows, err := db.Query(query)
	if err != nil {
		http.Error(w, "Error retrieving books", http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var books []Book
	for rows.Next() {
		var book Book
		err := rows.Scan(
			&book.ID,
			&book.Title,
			&book.Author,
			&book.CoverURL,
			&book.PageCount,
			&book.Language,
			&book.CreatedAt,
		)
		if err != nil {
			http.Error(w, "Error scanning books", http.StatusInternalServerError)
			return
		}
		books = append(books, book)
	}

	json.NewEncoder(w).Encode(books)
}

func updateProgressHandler(w http.ResponseWriter, r *http.Request) {
	var progress ReadingProgress
	if err := json.NewDecoder(r.Body).Decode(&progress); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	// Generate ID if not provided
	if progress.ID == "" {
		progress.ID = uuid.New().String()
	}

	if err := UpdateReadingProgress(db, &progress); err != nil {
		http.Error(w, "Error updating progress", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(progress)
}

func getProgressHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userID := r.Header.Get("X-User-ID") // Get user ID from header

	progress, err := GetReadingProgress(db, vars["bookId"], userID)
	if err != nil {
		http.Error(w, "Error retrieving progress", http.StatusInternalServerError)
		return
	}

	if progress == nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	json.NewEncoder(w).Encode(progress)
}

func createBookmarkHandler(w http.ResponseWriter, r *http.Request) {
	var bookmark Bookmark
	if err := json.NewDecoder(r.Body).Decode(&bookmark); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	// Generate ID if not provided
	if bookmark.ID == "" {
		bookmark.ID = uuid.New().String()
	}

	if err := CreateBookmark(db, &bookmark); err != nil {
		http.Error(w, "Error creating bookmark", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(bookmark)
}

func getBookmarksHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userID := r.Header.Get("X-User-ID") // Get user ID from header

	bookmarks, err := GetBookmarks(db, vars["bookId"], userID)
	if err != nil {
		http.Error(w, "Error retrieving bookmarks", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(bookmarks)
}

func updateBookmarkHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	var bookmark Bookmark
	if err := json.NewDecoder(r.Body).Decode(&bookmark); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	bookmark.ID = vars["id"]
	if err := UpdateBookmark(db, &bookmark); err != nil {
		http.Error(w, "Error updating bookmark", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(bookmark)
}

func deleteBookmarkHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	if err := DeleteBookmark(db, vars["id"]); err != nil {
		http.Error(w, "Error deleting bookmark", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func generateAudioHandler(w http.ResponseWriter, r *http.Request) {
	var segment AudioSegment
	if err := json.NewDecoder(r.Body).Decode(&segment); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	// Generate ID if not provided
	if segment.ID == "" {
		segment.ID = uuid.New().String()
	}

	// Set initial status
	segment.Status = "pending"
	segment.CreatedAt = time.Now()

	if err := CreateAudioSegment(db, &segment); err != nil {
		http.Error(w, "Error creating audio segment", http.StatusInternalServerError)
		return
	}

	// Start TTS processing in background
	go processAudioSegment(&segment)

	json.NewEncoder(w).Encode(segment)
}

func getAudioSegmentsHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	query := `
		SELECT id, book_id, segment_number, content, audio_url, duration, status, created_at
		FROM audio_segments
		WHERE book_id = ?
		ORDER BY segment_number ASC
	`

	rows, err := db.Query(query, vars["bookId"])
	if err != nil {
		http.Error(w, "Error retrieving audio segments", http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var segments []AudioSegment
	for rows.Next() {
		var segment AudioSegment
		err := rows.Scan(
			&segment.ID,
			&segment.BookID,
			&segment.SegmentNumber,
			&segment.Content,
			&segment.AudioURL,
			&segment.Duration,
			&segment.Status,
			&segment.CreatedAt,
		)
		if err != nil {
			http.Error(w, "Error scanning audio segments", http.StatusInternalServerError)
			return
		}
		segments = append(segments, segment)
	}

	json.NewEncoder(w).Encode(segments)
}

func getCategoriesHandler(w http.ResponseWriter, r *http.Request) {
	query := "SELECT id, name, description, created_at FROM categories ORDER BY name"
	rows, err := db.Query(query)
	if err != nil {
		http.Error(w, "Error retrieving categories", http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var categories []Category
	for rows.Next() {
		var category Category
		err := rows.Scan(
			&category.ID,
			&category.Name,
			&category.Description,
			&category.CreatedAt,
		)
		if err != nil {
			http.Error(w, "Error scanning categories", http.StatusInternalServerError)
			return
		}
		categories = append(categories, category)
	}

	json.NewEncoder(w).Encode(categories)
}

func getTagsHandler(w http.ResponseWriter, r *http.Request) {
	query := "SELECT id, name, created_at FROM tags ORDER BY name"
	rows, err := db.Query(query)
	if err != nil {
		http.Error(w, "Error retrieving tags", http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var tags []Tag
	for rows.Next() {
		var tag Tag
		err := rows.Scan(
			&tag.ID,
			&tag.Name,
			&tag.CreatedAt,
		)
		if err != nil {
			http.Error(w, "Error scanning tags", http.StatusInternalServerError)
			return
		}
		tags = append(tags, tag)
	}

	json.NewEncoder(w).Encode(tags)
}
```



---

> 📸 Generated with [Jockey CLI](https://github.com/saint0x/jockey-cli)
