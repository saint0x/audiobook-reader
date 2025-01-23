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

	// Create audio_segments table
	_, err = db.Exec(`
		CREATE TABLE IF NOT EXISTS audio_segments (
			id TEXT PRIMARY KEY,
			book_id TEXT NOT NULL,
			segment_number INTEGER NOT NULL,
			content TEXT NOT NULL,
			audio_url TEXT,
			duration REAL,
			status TEXT NOT NULL,
			created_at DATETIME NOT NULL,
			updated_at DATETIME NOT NULL,
			FOREIGN KEY (book_id) REFERENCES books(id)
		)
	`)
	if err != nil {
		return fmt.Errorf("error creating audio_segments table: %v", err)
	}

	return nil
}

// GetBookByID retrieves a book by its ID
func GetBookByID(db *sql.DB, id string) (*Book, error) {
	var book Book
	query := `
		SELECT id, title, author, cover_url, file_url, 
			page_count, current_page, language, status,
			created_at, updated_at
		FROM books WHERE id = ?
	`

	err := db.QueryRow(query, id).Scan(
		&book.ID,
		&book.Title,
		&book.Author,
		&book.CoverURL,
		&book.FileURL,
		&book.PageCount,
		&book.CurrentPage,
		&book.Language,
		&book.Status,
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
			id, title, author, cover_url, file_url,
			page_count, current_page, language, status,
			created_at, updated_at
		) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)
	`

	_, err := db.Exec(query,
		book.ID,
		book.Title,
		book.Author,
		book.CoverURL,
		book.FileURL,
		book.PageCount,
		book.CurrentPage,
		book.Language,
		book.Status,
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

// GetAudioSegments retrieves all audio segments for a book
func GetAudioSegments(bookID string) ([]AudioSegment, error) {
	query := `
		SELECT id, book_id, content, audio_url, status, created_at, updated_at
		FROM audio_segments
		WHERE book_id = ?
		ORDER BY created_at ASC
	`

	rows, err := db.Query(query, bookID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var segments []AudioSegment
	for rows.Next() {
		var segment AudioSegment
		err := rows.Scan(
			&segment.ID,
			&segment.BookID,
			&segment.Content,
			&segment.AudioURL,
			&segment.Status,
			&segment.CreatedAt,
			&segment.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}
		segments = append(segments, segment)
	}

	return segments, nil
}

// CreateAudioSegment creates a new audio segment
func CreateAudioSegment(db *sql.DB, segment *AudioSegment) error {
	query := `
		INSERT INTO audio_segments (
			id, book_id, content, audio_url, status,
			created_at, updated_at
		) VALUES (?, ?, ?, ?, ?, ?, ?)
	`

	_, err := db.Exec(query,
		segment.ID,
		segment.BookID,
		segment.Content,
		segment.AudioURL,
		segment.Status,
		formatDateTime(segment.CreatedAt),
		formatDateTime(segment.UpdatedAt),
	)

	return err
}

// UpdateAudioSegment updates an existing audio segment
func UpdateAudioSegment(db *sql.DB, segment *AudioSegment) error {
	query := `
		UPDATE audio_segments 
		SET content = ?, audio_url = ?, status = ?, updated_at = ?
		WHERE id = ? AND book_id = ?
	`

	_, err := db.Exec(query,
		segment.Content,
		segment.AudioURL,
		segment.Status,
		formatDateTime(segment.UpdatedAt),
		segment.ID,
		segment.BookID,
	)

	return err
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

// CleanupDuplicateBooks removes duplicate books keeping only the latest version
func CleanupDuplicateBooks(db *sql.DB) error {
	// First, get all books grouped by title, keeping only the latest version
	query := `
		WITH RankedBooks AS (
			SELECT *,
				ROW_NUMBER() OVER (PARTITION BY title ORDER BY created_at DESC) as rn
			FROM books
		)
		DELETE FROM books 
		WHERE id IN (
			SELECT id 
			FROM RankedBooks 
			WHERE rn > 1
		)
	`

	_, err := db.Exec(query)
	if err != nil {
		return fmt.Errorf("error cleaning up duplicate books: %v", err)
	}

	// Also clean up orphaned audio segments
	_, err = db.Exec(`
		DELETE FROM audio_segments 
		WHERE book_id NOT IN (SELECT id FROM books)
	`)
	if err != nil {
		return fmt.Errorf("error cleaning up orphaned audio segments: %v", err)
	}

	return nil
}
