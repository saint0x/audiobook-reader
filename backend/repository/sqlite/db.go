package sqlite

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/mattn/go-sqlite3"
)

// DB represents a database connection
type DB struct {
	*sql.DB
}

// NewDB creates a new database connection
func NewDB(dbPath string) (*DB, error) {
	db, err := sql.Open("sqlite3", dbPath)
	if err != nil {
		return nil, fmt.Errorf("error opening database: %v", err)
	}

	return &DB{db}, nil
}

// InitDB initializes the database connection and creates tables
func (db *DB) InitDB() error {
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

// CleanupDuplicateBooks removes duplicate book entries based on title and author
func (db *DB) CleanupDuplicateBooks() error {
	query := `
		WITH duplicates AS (
			SELECT title, author, COUNT(*) as count
			FROM books
			GROUP BY title, author
			HAVING count > 1
		)
		DELETE FROM books
		WHERE (title, author) IN (
			SELECT title, author FROM duplicates
		)
		AND id NOT IN (
			SELECT MIN(id)
			FROM books
			GROUP BY title, author
		)
	`

	_, err := db.Exec(query)
	if err != nil {
		return fmt.Errorf("error cleaning up duplicate books: %v", err)
	}

	return nil
}
