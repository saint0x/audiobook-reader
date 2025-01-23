package main

import (
	"database/sql"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/google/uuid"

	"github.com/ledongthuc/pdf"
)

// processPDF processes a PDF file and returns a Book object
func processPDF(file multipart.File, filename string) (*Book, error) {
	// Create temporary file
	tmpFile, err := os.CreateTemp("", "book-*.pdf")
	if err != nil {
		return nil, fmt.Errorf("error creating temp file: %v", err)
	}
	defer os.Remove(tmpFile.Name())

	// Copy uploaded file to temp file
	if _, err := io.Copy(tmpFile, file); err != nil {
		return nil, fmt.Errorf("error copying file: %v", err)
	}

	// Open PDF file for reading
	pdfFile, reader, err := pdf.Open(tmpFile.Name())
	if err != nil {
		return nil, fmt.Errorf("error opening PDF: %v", err)
	}
	defer pdfFile.Close()

	// Get total number of pages
	totalPages := reader.NumPage()

	// Create book object
	book := &Book{
		ID:          uuid.New().String(),
		Title:       strings.TrimSuffix(filename, ".pdf"),
		Author:      "Unknown", // Can be updated later
		FileURL:     "",        // Will be set after uploading to UploadThing
		PageCount:   totalPages,
		CurrentPage: 0,
		Language:    "en", // Default to English, can be updated later
		Status:      "pending",
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	return book, nil
}

func processBook(db *sql.DB, book *Book) error {
	// Download PDF from UploadThing
	req, err := http.NewRequest("GET", book.FileURL, nil)
	if err != nil {
		return fmt.Errorf("failed to create request: %v", err)
	}

	// Add UploadThing authentication
	req.Header.Set("Authorization", "Bearer "+AppConfig.UploadThingToken)

	// Make the request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return fmt.Errorf("failed to download PDF: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("failed to download PDF: status code %d", resp.StatusCode)
	}

	// Create temporary file
	tmpFile, err := os.CreateTemp("", "book-*.pdf")
	if err != nil {
		return fmt.Errorf("failed to create temp file: %v", err)
	}
	defer os.Remove(tmpFile.Name())

	// Copy PDF to temp file
	if _, err := io.Copy(tmpFile, resp.Body); err != nil {
		return fmt.Errorf("failed to copy PDF to temp file: %v", err)
	}

	// Open PDF file
	pdfFile, reader, err := pdf.Open(tmpFile.Name())
	if err != nil {
		return fmt.Errorf("failed to open PDF: %v", err)
	}
	defer pdfFile.Close()

	// Update book page count
	totalPages := reader.NumPage()
	book.PageCount = totalPages
	book.UpdatedAt = time.Now()
	if err := UpdateBook(db, book); err != nil {
		return fmt.Errorf("failed to update book: %v", err)
	}

	// Process each page
	for pageNum := 1; pageNum <= totalPages; pageNum++ {
		page := reader.Page(pageNum)
		if page.V.IsNull() {
			continue
		}

		text, err := page.GetPlainText(nil)
		if err != nil {
			return fmt.Errorf("failed to get text from page %d: %v", pageNum, err)
		}

		// Create audio segment
		segment := &AudioSegment{
			ID:        uuid.New().String(),
			BookID:    book.ID,
			Content:   text,
			Status:    "pending",
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		}

		if err := SaveAudioSegment(db, segment); err != nil {
			return fmt.Errorf("failed to save audio segment: %v", err)
		}

		// Start TTS processing in background
		go processAudioSegment(db, segment)
	}

	return nil
}
