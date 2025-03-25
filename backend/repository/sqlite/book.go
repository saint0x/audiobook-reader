package sqlite

import (
	"fmt"
	"time"

	"backend/domain/models"
)

// SaveBook saves a book to the database
func (db *DB) SaveBook(book *models.Book) error {
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
		book.CreatedAt,
		book.UpdatedAt,
	)

	if err != nil {
		return fmt.Errorf("error saving book: %v", err)
	}

	return nil
}

// GetBookByID retrieves a book by its ID
func (db *DB) GetBookByID(id string) (*models.Book, error) {
	query := `
		SELECT id, title, author, cover_url, file_url,
			   page_count, current_page, language, status,
			   created_at, updated_at
		FROM books
		WHERE id = ?
	`

	book := &models.Book{}
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

	if err != nil {
		return nil, fmt.Errorf("error getting book: %v", err)
	}

	return book, nil
}

// UpdateBook updates an existing book in the database
func (db *DB) UpdateBook(book *models.Book) error {
	query := `
		UPDATE books 
		SET title = ?, author = ?, cover_url = ?, file_url = ?, 
			page_count = ?, current_page = ?, language = ?, status = ?,
			updated_at = ?
		WHERE id = ?
	`

	book.UpdatedAt = time.Now()
	_, err := db.Exec(query,
		book.Title,
		book.Author,
		book.CoverURL,
		book.FileURL,
		book.PageCount,
		book.CurrentPage,
		book.Language,
		book.Status,
		book.UpdatedAt,
		book.ID,
	)

	if err != nil {
		return fmt.Errorf("error updating book: %v", err)
	}

	return nil
}

// GetBooks retrieves all books from the database
func (db *DB) GetBooks() ([]models.Book, error) {
	query := `
		SELECT id, title, author, cover_url, file_url,
			   page_count, current_page, language, status,
			   created_at, updated_at
		FROM books
		ORDER BY created_at DESC
	`

	rows, err := db.Query(query)
	if err != nil {
		return nil, fmt.Errorf("error querying books: %v", err)
	}
	defer rows.Close()

	var books []models.Book
	for rows.Next() {
		var book models.Book
		err := rows.Scan(
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
		if err != nil {
			return nil, fmt.Errorf("error scanning book: %v", err)
		}
		books = append(books, book)
	}

	return books, nil
}

// DeleteBook deletes a book from the database
func (db *DB) DeleteBook(id string) error {
	query := "DELETE FROM books WHERE id = ?"
	_, err := db.Exec(query, id)
	if err != nil {
		return fmt.Errorf("error deleting book: %v", err)
	}
	return nil
}
