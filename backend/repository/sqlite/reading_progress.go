package sqlite

import (
	"fmt"

	"backend/domain/models"
)

// UpdateReadingProgress updates a user's reading progress for a book
func (db *DB) UpdateReadingProgress(progress *models.ReadingProgress) error {
	query := `
		INSERT OR REPLACE INTO reading_progress (
			id, book_id, user_id, current_page, total_pages,
			completion_percent, last_read_at
		) VALUES (?, ?, ?, ?, ?, ?, ?)
	`

	_, err := db.Exec(query,
		progress.ID,
		progress.BookID,
		progress.UserID,
		progress.CurrentPage,
		progress.TotalPages,
		progress.CompletionPercent,
		progress.LastReadAt,
	)

	if err != nil {
		return fmt.Errorf("error updating reading progress: %v", err)
	}

	return nil
}

// GetReadingProgress retrieves a user's reading progress for a book
func (db *DB) GetReadingProgress(bookID, userID string) (*models.ReadingProgress, error) {
	query := `
		SELECT id, book_id, user_id, current_page, total_pages,
			   completion_percent, last_read_at
		FROM reading_progress
		WHERE book_id = ? AND user_id = ?
	`

	progress := &models.ReadingProgress{}
	err := db.QueryRow(query, bookID, userID).Scan(
		&progress.ID,
		&progress.BookID,
		&progress.UserID,
		&progress.CurrentPage,
		&progress.TotalPages,
		&progress.CompletionPercent,
		&progress.LastReadAt,
	)

	if err != nil {
		return nil, fmt.Errorf("error getting reading progress: %v", err)
	}

	return progress, nil
}

// DeleteReadingProgress deletes a user's reading progress for a book
func (db *DB) DeleteReadingProgress(bookID, userID string) error {
	query := "DELETE FROM reading_progress WHERE book_id = ? AND user_id = ?"
	_, err := db.Exec(query, bookID, userID)
	if err != nil {
		return fmt.Errorf("error deleting reading progress: %v", err)
	}
	return nil
}
