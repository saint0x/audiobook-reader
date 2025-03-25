package pdf

import (
	"fmt"
	"io"
	"os"
	"strings"
	"time"

	"github.com/google/uuid"
	"github.com/ledongthuc/pdf"

	"backend/domain/models"
)

// ProcessPDF processes a PDF file and returns a Book object
func ProcessPDF(file io.Reader, filename string) (*models.Book, error) {
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
	book := &models.Book{
		ID:          uuid.New().String(),
		Title:       strings.TrimSuffix(filename, ".pdf"),
		Author:      "Unknown", // Can be updated later
		FileURL:     "",        // Will be set after uploading to UploadThing
		PageCount:   totalPages,
		CurrentPage: 0,
		Language:    "en", // Default to English, can be updated later
		Status:      "processing",
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	return book, nil
}

// ExtractText extracts text from a PDF file page by page
func ExtractText(filePath string) ([]string, error) {
	pdfFile, reader, err := pdf.Open(filePath)
	if err != nil {
		return nil, fmt.Errorf("error opening PDF: %v", err)
	}
	defer pdfFile.Close()

	var segments []string
	for pageNum := 1; pageNum <= reader.NumPage(); pageNum++ {
		page := reader.Page(pageNum)
		if page.V.IsNull() {
			continue
		}

		text, err := page.GetPlainText(nil)
		if err != nil {
			return nil, fmt.Errorf("error extracting text from page %d: %v", pageNum, err)
		}

		segments = append(segments, text)
	}

	return segments, nil
}
