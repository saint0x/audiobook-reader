package storage

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

	// Create destination file
	dst, err := os.Create(filePath)
	if err != nil {
		return "", fmt.Errorf("error creating file: %v", err)
	}
	defer dst.Close()

	// Copy file contents
	if _, err := io.Copy(dst, file); err != nil {
		return "", fmt.Errorf("error copying file: %v", err)
	}

	return filePath, nil
}

// SaveCover saves a cover image and returns its URL
func (fs *FileStorage) SaveCover(file multipart.File, filename string) (string, error) {
	// Generate unique filename
	ext := filepath.Ext(filename)
	if !isValidImageExt(ext) {
		return "", fmt.Errorf("invalid file type: %s", ext)
	}

	filename = uuid.New().String() + ext
	filePath := filepath.Join(fs.coverDir, filename)

	// Create destination file
	dst, err := os.Create(filePath)
	if err != nil {
		return "", fmt.Errorf("error creating file: %v", err)
	}
	defer dst.Close()

	// Copy file contents
	if _, err := io.Copy(dst, file); err != nil {
		return "", fmt.Errorf("error copying file: %v", err)
	}

	return "/covers/" + filename, nil
}

// SaveAudio saves an audio file and returns its URL
func (fs *FileStorage) SaveAudio(data []byte, filename string) (string, error) {
	filePath := filepath.Join(fs.audioDir, filename)

	if err := os.WriteFile(filePath, data, 0644); err != nil {
		return "", fmt.Errorf("error writing audio file: %v", err)
	}

	return "/audio/" + filename, nil
}

func isValidImageExt(ext string) bool {
	ext = strings.ToLower(ext)
	return ext == ".jpg" || ext == ".jpeg" || ext == ".png" || ext == ".gif"
}
