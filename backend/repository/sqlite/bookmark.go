package sqlite

import (
	"fmt"
	"time"

	"backend/domain/models"
)

// CreateBookmark creates a new bookmark in the database
func (db *DB) CreateBookmark(bookmark *models.Bookmark) error {
	query := `
		INSERT INTO bookmarks (
			id, book_id, user_id, page_number, note,
			created_at, updated_at
		) VALUES (?, ?, ?, ?, ?, ?, ?)
	`

	_, err := db.Exec(query,
		bookmark.ID,
		bookmark.BookID,
		bookmark.UserID,
		bookmark.PageNumber,
		bookmark.Note,
		bookmark.CreatedAt,
		bookmark.UpdatedAt,
	)

	if err != nil {
		return fmt.Errorf("error creating bookmark: %v", err)
	}

	return nil
}

// UpdateBookmark updates an existing bookmark in the database
func (db *DB) UpdateBookmark(bookmark *models.Bookmark) error {
	query := `
		UPDATE bookmarks 
		SET page_number = ?, note = ?, updated_at = ?
		WHERE id = ? AND user_id = ?
	`

	bookmark.UpdatedAt = time.Now()
	_, err := db.Exec(query,
		bookmark.PageNumber,
		bookmark.Note,
		bookmark.UpdatedAt,
		bookmark.ID,
		bookmark.UserID,
	)

	if err != nil {
		return fmt.Errorf("error updating bookmark: %v", err)
	}

	return nil
}

// GetBookmarks retrieves all bookmarks for a book and user
func (db *DB) GetBookmarks(bookID, userID string) ([]models.Bookmark, error) {
	query := `
		SELECT id, book_id, user_id, page_number, note, created_at, updated_at
		FROM bookmarks
		WHERE book_id = ? AND user_id = ?
		ORDER BY page_number ASC
	`

	rows, err := db.Query(query, bookID, userID)
	if err != nil {
		return nil, fmt.Errorf("error querying bookmarks: %v", err)
	}
	defer rows.Close()

	var bookmarks []models.Bookmark
	for rows.Next() {
		var bookmark models.Bookmark
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
			return nil, fmt.Errorf("error scanning bookmark: %v", err)
		}
		bookmarks = append(bookmarks, bookmark)
	}

	return bookmarks, nil
}

// GetBookmarkByID retrieves a bookmark by its ID
func (db *DB) GetBookmarkByID(id string) (*models.Bookmark, error) {
	query := `
		SELECT id, book_id, user_id, page_number, note, created_at, updated_at
		FROM bookmarks
		WHERE id = ?
	`

	bookmark := &models.Bookmark{}
	err := db.QueryRow(query, id).Scan(
		&bookmark.ID,
		&bookmark.BookID,
		&bookmark.UserID,
		&bookmark.PageNumber,
		&bookmark.Note,
		&bookmark.CreatedAt,
		&bookmark.UpdatedAt,
	)

	if err != nil {
		return nil, fmt.Errorf("error getting bookmark: %v", err)
	}

	return bookmark, nil
}

// DeleteBookmark deletes a bookmark from the database
func (db *DB) DeleteBookmark(id string) error {
	query := "DELETE FROM bookmarks WHERE id = ?"
	_, err := db.Exec(query, id)
	if err != nil {
		return fmt.Errorf("error deleting bookmark: %v", err)
	}
	return nil
}
