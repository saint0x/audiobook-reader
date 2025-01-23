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
func (fs *FileStorage) SaveAudio(file io.Reader, filename string) (string, error) {
	// Generate unique filename
	ext := filepath.Ext(filename)
	if ext == "" {
		ext = ".mp3"
	}
	filename = uuid.New().String() + ext

	// Create audio directory if it doesn't exist
	audioDir := filepath.Join(fs.baseDir, "audio")
	if err := os.MkdirAll(audioDir, 0755); err != nil {
		return "", fmt.Errorf("error creating audio directory: %v", err)
	}

	// Create file
	filePath := filepath.Join(audioDir, filename)
	f, err := os.Create(filePath)
	if err != nil {
		return "", fmt.Errorf("error creating file: %v", err)
	}
	defer f.Close()

	// Copy file contents
	if _, err := io.Copy(f, file); err != nil {
		return "", fmt.Errorf("error copying file: %v", err)
	}

	// Return URL path relative to server root
	return fmt.Sprintf("/audio/%s", filename), nil
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
