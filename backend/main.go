package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"
	"time"

	"backend/config"
	"backend/domain/models"
	"backend/repository/sqlite"
	"backend/service/pdf"
	"backend/service/storage"
	"backend/service/tts"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"github.com/gorilla/websocket"
)

var (
	db          *sqlite.DB
	fileStorage *storage.FileStorage
	ttsGen      *tts.Generator
)

// Add WebSocket upgrader
var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true // Allow all origins for now
	},
}

// Add WebSocket connections map
var wsConnections = make(map[string][]*websocket.Conn)

func main() {
	// Load configuration
	if err := config.LoadConfig(); err != nil {
		log.Fatal("Error loading configuration:", err)
	}

	// Initialize database
	var err error
	db, err = sqlite.NewDB(config.AppConfig.DBPath)
	if err != nil {
		log.Fatal("Error opening database:", err)
	}
	defer db.Close()

	// Create tables
	if err := db.InitDB(); err != nil {
		log.Fatal("Error initializing database:", err)
	}

	// Clean up any duplicate books
	if err := db.CleanupDuplicateBooks(); err != nil {
		log.Printf("Warning: Error cleaning up duplicates: %v", err)
	}

	// Initialize file storage
	fileStorage, err = storage.NewFileStorage(config.AppConfig.UploadDir)
	if err != nil {
		log.Fatal("Error initializing file storage:", err)
	}

	// Initialize TTS generator
	ttsGen = tts.NewGenerator(&config.AppConfig, db)

	// Create audio directory if it doesn't exist
	audioDir := filepath.Join(config.AppConfig.UploadDir, "audio")
	if err := os.MkdirAll(audioDir, 0755); err != nil {
		log.Fatal("Error creating audio directory:", err)
	}

	// Initialize router
	router := mux.NewRouter()

	// Add CORS middleware
	router.Use(func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			// Set CORS headers for all responses
			w.Header().Set("Access-Control-Allow-Origin", "*")
			w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
			w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization, X-User-ID")
			w.Header().Set("Access-Control-Max-Age", "3600")

			// Handle preflight requests
			if r.Method == "OPTIONS" {
				w.WriteHeader(http.StatusOK)
				return
			}

			next.ServeHTTP(w, r)
		})
	})

	// Serve static audio files
	audioDir = filepath.Join(config.AppConfig.UploadDir, "audio")
	router.PathPrefix("/audio/").Handler(http.StripPrefix("/audio/", http.FileServer(http.Dir(audioDir))))

	// File upload routes
	router.HandleFunc("/api/upload", uploadPDFHandler).Methods("POST", "OPTIONS")
	router.HandleFunc("/api/upload/pdf", uploadPDFHandler).Methods("POST", "OPTIONS")
	router.HandleFunc("/api/upload/cover", uploadCoverHandler).Methods("POST", "OPTIONS")

	// Book routes
	router.HandleFunc("/api/books/{id}", getBookHandler).Methods("GET")
	router.HandleFunc("/api/books", getBooksHandler).Methods("GET")
	router.HandleFunc("/api/books/{id}/status", getBookStatusHandler).Methods("GET")
	router.HandleFunc("/api/books/{id}/update-url", updateBookURLHandler).Methods("POST")
	router.HandleFunc("/api/books/{id}/process", processBookHandler).Methods("POST")

	// Reading progress routes
	router.HandleFunc("/api/progress", updateProgressHandler).Methods("POST")
	router.HandleFunc("/api/progress/{bookId}", getProgressHandler).Methods("GET")

	// Bookmark routes
	router.HandleFunc("/api/bookmarks", createBookmarkHandler).Methods("POST")
	router.HandleFunc("/api/bookmarks/{bookId}", getBookmarksHandler).Methods("GET")
	router.HandleFunc("/api/bookmarks/{id}", updateBookmarkHandler).Methods("PUT")
	router.HandleFunc("/api/bookmarks/{id}", deleteBookmarkHandler).Methods("DELETE")

	// Audio segment routes
	router.HandleFunc("/api/books/{id}/audio-segments", getAudioSegmentsHandler).Methods("GET", "OPTIONS")
	router.HandleFunc("/api/books/{id}/generate-audio", generateBookAudioHandler).Methods("POST")
	router.HandleFunc("/api/audio/generate", generateAudioHandler).Methods("POST")

	// Category and tag routes
	router.HandleFunc("/api/categories", getCategoriesHandler).Methods("GET")
	router.HandleFunc("/api/tags", getTagsHandler).Methods("GET")

	// WebSocket routes
	router.HandleFunc("/ws/books/{id}", wsHandler)

	// Start server
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	log.Printf("Server starting on port %s", port)
	log.Fatal(http.ListenAndServe(":"+port, router))
}

func uploadPDFHandler(w http.ResponseWriter, r *http.Request) {
	log.Printf("[Upload] Starting new upload request. Method: %s, Content-Type: %s", r.Method, r.Header.Get("Content-Type"))

	var req struct {
		FileURL string `json:"fileUrl"`
		Title   string `json:"title"`
	}

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		log.Printf("[Upload] Error parsing request: %v", err)
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}
	log.Printf("[Upload] Received request for title: %s, URL: %s", req.Title, req.FileURL)

	// Create initial book record
	book := &models.Book{
		ID:        uuid.New().String(),
		Title:     req.Title,
		FileURL:   req.FileURL,
		Status:    "processing",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	log.Printf("[Upload] Created book record with ID: %s", book.ID)

	// Save initial book record
	if err := db.SaveBook(book); err != nil {
		log.Printf("[Upload] Error saving book: %v", err)
		http.Error(w, fmt.Sprintf("Error saving book: %v", err), http.StatusInternalServerError)
		return
	}
	log.Printf("[Upload] Successfully saved book to database")

	// Start processing in background immediately
	go func() {
		log.Printf("[Processing] Starting background processing for book: %s", book.ID)

		// Download the PDF file
		resp, err := http.Get(book.FileURL)
		if err != nil {
			log.Printf("[PDF] Error downloading PDF: %v", err)
			book.Status = "error"
			db.UpdateBook(book)
			return
		}
		defer resp.Body.Close()

		// Process PDF
		processedBook, err := pdf.ProcessPDF(resp.Body, filepath.Base(book.FileURL))
		if err != nil {
			log.Printf("[PDF] Error processing PDF: %v", err)
			book.Status = "error"
			db.UpdateBook(book)
			return
		}

		// Update book with processed information
		book.PageCount = processedBook.PageCount
		book.Author = processedBook.Author
		book.Language = processedBook.Language
		if err := db.UpdateBook(book); err != nil {
			log.Printf("[PDF] Error updating book: %v", err)
			return
		}

		// Create temporary file for PDF processing
		tmpFile, err := os.CreateTemp("", "book-*.pdf")
		if err != nil {
			log.Printf("[PDF] Error creating temp file: %v", err)
			book.Status = "error"
			db.UpdateBook(book)
			return
		}
		defer os.Remove(tmpFile.Name())

		// Copy PDF to temporary file
		if _, err := io.Copy(tmpFile, resp.Body); err != nil {
			log.Printf("[PDF] Error copying to temp file: %v", err)
			book.Status = "error"
			db.UpdateBook(book)
			return
		}

		// Extract text segments
		textSegments, err := pdf.ExtractText(tmpFile.Name())
		if err != nil {
			log.Printf("[PDF] Error extracting text: %v", err)
			book.Status = "error"
			db.UpdateBook(book)
			return
		}

		// Create audio segments
		for i, text := range textSegments {
			segment := &models.AudioSegment{
				ID:        uuid.New().String(),
				BookID:    book.ID,
				Content:   text,
				Status:    "pending",
				CreatedAt: time.Now(),
				UpdatedAt: time.Now(),
			}
			if err := db.SaveAudioSegment(segment); err != nil {
				log.Printf("[PDF] Error saving segment %d: %v", i+1, err)
				continue
			}
		}

		// Update book status
		book.Status = "ready"
		book.UpdatedAt = time.Now()
		if err := db.UpdateBook(book); err != nil {
			log.Printf("[PDF] Error updating book status: %v", err)
			return
		}

		// Get audio segments from the book
		audioSegments, err := db.GetAudioSegments(book.ID)
		if err != nil {
			log.Printf("[Processing] Error getting segments: %v", err)
			return
		}

		// Process each segment
		for _, segment := range audioSegments {
			if segment.Status != "pending" {
				continue
			}

			// Generate audio for the segment
			audioData, err := ttsGen.ProcessAudioSegment(&segment)
			if err != nil {
				log.Printf("[Processing] Error generating audio for segment %s: %v", segment.ID, err)
				segment.Status = "error"
				db.UpdateAudioSegment(&segment)
				continue
			}

			// Save audio to file
			audioFileName := fmt.Sprintf("%s.mp3", segment.ID)
			audioPath := filepath.Join("audio", audioFileName)
			if err := os.WriteFile(audioPath, audioData, 0644); err != nil {
				log.Printf("[Processing] Error saving audio file for segment %s: %v", segment.ID, err)
				segment.Status = "error"
				db.UpdateAudioSegment(&segment)
				continue
			}

			// Update segment with audio URL and status
			segment.AudioURL = audioPath
			segment.Status = "ready"
			segment.UpdatedAt = time.Now()
			if err := db.UpdateAudioSegment(&segment); err != nil {
				log.Printf("[Processing] Error updating segment %s: %v", segment.ID, err)
				continue
			}
		}

		// Update book status to ready
		book.Status = "ready"
		book.UpdatedAt = time.Now()
		if err := db.UpdateBook(book); err != nil {
			log.Printf("[Processing] Error updating book status: %v", err)
		}
		log.Printf("[Processing] Book status updated to ready: %s", book.ID)
	}()

	log.Printf("[Upload] Returning response for book: %s", book.ID)
	// Return immediate response with book ID
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]interface{}{
		"id":     book.ID,
		"status": book.Status,
	})
}

// Add WebSocket handler
func wsHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	bookID := vars["id"]

	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Printf("[WS] Error upgrading connection: %v", err)
		return
	}

	// Add connection to the map
	wsConnections[bookID] = append(wsConnections[bookID], conn)
	log.Printf("[WS] New connection established for book: %s", bookID)

	// Clean up on disconnect
	go func() {
		<-r.Context().Done()
		wsConnections[bookID] = removeConn(wsConnections[bookID], conn)
		log.Printf("[WS] Connection closed for book: %s", bookID)
	}()
}

func removeConn(conns []*websocket.Conn, conn *websocket.Conn) []*websocket.Conn {
	for i, c := range conns {
		if c == conn {
			return append(conns[:i], conns[i+1:]...)
		}
	}
	return conns
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
	book, err := db.GetBookByID(vars["id"])
	if err != nil {
		http.Error(w, "Book not found", http.StatusNotFound)
		return
	}

	json.NewEncoder(w).Encode(book)
}

func getBooksHandler(w http.ResponseWriter, r *http.Request) {
	books, err := db.GetBooks()
	if err != nil {
		http.Error(w, "Error retrieving books", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(books)
}

func updateProgressHandler(w http.ResponseWriter, r *http.Request) {
	var progress models.ReadingProgress
	if err := json.NewDecoder(r.Body).Decode(&progress); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	// Generate ID if not provided
	if progress.ID == "" {
		progress.ID = uuid.New().String()
	}

	if err := db.UpdateReadingProgress(&progress); err != nil {
		http.Error(w, "Error updating progress", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(progress)
}

func getProgressHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userID := r.Header.Get("X-User-ID") // Get user ID from header

	progress, err := db.GetReadingProgress(vars["bookId"], userID)
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
	var bookmark models.Bookmark
	if err := json.NewDecoder(r.Body).Decode(&bookmark); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	// Generate ID if not provided
	if bookmark.ID == "" {
		bookmark.ID = uuid.New().String()
	}

	if err := db.CreateBookmark(&bookmark); err != nil {
		http.Error(w, "Error creating bookmark", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(bookmark)
}

func getBookmarksHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userID := r.Header.Get("X-User-ID") // Get user ID from header

	bookmarks, err := db.GetBookmarks(vars["bookId"], userID)
	if err != nil {
		http.Error(w, "Error retrieving bookmarks", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(bookmarks)
}

func updateBookmarkHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	var bookmark models.Bookmark
	if err := json.NewDecoder(r.Body).Decode(&bookmark); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	bookmark.ID = vars["id"]
	if err := db.UpdateBookmark(&bookmark); err != nil {
		http.Error(w, "Error updating bookmark", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(bookmark)
}

func deleteBookmarkHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	if err := db.DeleteBookmark(vars["id"]); err != nil {
		http.Error(w, "Error deleting bookmark", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func generateAudioHandler(w http.ResponseWriter, r *http.Request) {
	var segment models.AudioSegment
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
	segment.UpdatedAt = time.Now()

	if err := db.SaveAudioSegment(&segment); err != nil {
		http.Error(w, "Error creating audio segment", http.StatusInternalServerError)
		return
	}

	// Start TTS processing in background
	go func() {
		audioData, err := ttsGen.GenerateAudio(segment.Content)
		if err != nil {
			log.Printf("[TTS] Error generating audio: %v", err)
			segment.Status = "error"
			db.UpdateAudioSegment(&segment)
			return
		}

		// Save audio file
		audioFileName := fmt.Sprintf("tts-%s.mp3", segment.ID)
		audioPath := filepath.Join(config.AppConfig.UploadDir, "audio", audioFileName)
		if err := os.WriteFile(audioPath, audioData, 0644); err != nil {
			log.Printf("[TTS] Error saving audio file: %v", err)
			segment.Status = "error"
			db.UpdateAudioSegment(&segment)
			return
		}

		// Update segment
		segment.AudioURL = "/audio/" + audioFileName
		segment.Status = "completed"
		segment.UpdatedAt = time.Now()
		if err := db.UpdateAudioSegment(&segment); err != nil {
			log.Printf("[TTS] Error updating segment: %v", err)
			return
		}
	}()

	json.NewEncoder(w).Encode(segment)
}

func getAudioSegmentsHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	bookID := vars["id"]

	segments, err := db.GetAudioSegments(bookID)
	if err != nil {
		http.Error(w, fmt.Sprintf("Failed to get audio segments: %v", err), http.StatusInternalServerError)
		return
	}

	// Initialize empty array if segments is nil
	if segments == nil {
		segments = []models.AudioSegment{}
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(segments)
}

func getCategoriesHandler(w http.ResponseWriter, r *http.Request) {
	categories, err := db.GetCategories()
	if err != nil {
		http.Error(w, "Error retrieving categories", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(categories)
}

func getTagsHandler(w http.ResponseWriter, r *http.Request) {
	tags, err := db.GetTags()
	if err != nil {
		http.Error(w, "Error retrieving tags", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(tags)
}

func getBookStatusHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	book, err := db.GetBookByID(vars["id"])
	if err != nil {
		http.Error(w, "Book not found", http.StatusNotFound)
		return
	}

	json.NewEncoder(w).Encode(map[string]interface{}{
		"id":     book.ID,
		"status": book.Status,
	})
}

func updateBookURLHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	var req struct {
		CloudURL string `json:"cloudUrl"`
	}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	book, err := db.GetBookByID(vars["id"])
	if err != nil {
		http.Error(w, "Book not found", http.StatusNotFound)
		return
	}

	book.FileURL = req.CloudURL
	if err := db.UpdateBook(book); err != nil {
		http.Error(w, "Error updating book", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func generateBookAudioHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	bookID := vars["id"]

	// Get the book
	book, err := db.GetBookByID(bookID)
	if err != nil {
		http.Error(w, "Book not found", http.StatusNotFound)
		return
	}

	// Get segments from the book
	segments, err := db.GetAudioSegments(book.ID)
	if err != nil {
		log.Printf("[Processing] Error getting segments: %v", err)
		http.Error(w, "Error getting audio segments", http.StatusInternalServerError)
		return
	}

	// Start TTS processing in background
	go func() {
		// Process each segment
		for _, segment := range segments {
			if segment.Status != "pending" {
				continue
			}

			// Generate audio
			audioData, err := ttsGen.GenerateAudio(segment.Content)
			if err != nil {
				log.Printf("[TTS] Error generating audio: %v", err)
				segment.Status = "error"
				db.UpdateAudioSegment(&segment)
				continue
			}

			// Save audio to temporary file for immediate playback
			audioFileName := fmt.Sprintf("tts-%s.mp3", uuid.New().String())
			audioPath := filepath.Join(config.AppConfig.UploadDir, "audio", audioFileName)
			if err := os.WriteFile(audioPath, audioData, 0644); err != nil {
				log.Printf("[TTS] Error writing audio file: %v", err)
				segment.Status = "error"
				db.UpdateAudioSegment(&segment)
				continue
			}

			// Update segment with local URL
			segment.AudioURL = "/audio/" + audioFileName
			segment.Status = "completed"
			segment.UpdatedAt = time.Now()
			if err := db.UpdateAudioSegment(&segment); err != nil {
				log.Printf("[TTS] Error updating audio segment: %v", err)
				continue
			}

			// Notify WebSocket clients about the new audio
			if conns, ok := wsConnections[book.ID]; ok {
				notification := map[string]interface{}{
					"type":    "audio_ready",
					"segment": segment,
				}
				for _, conn := range conns {
					if err := conn.WriteJSON(notification); err != nil {
						log.Printf("[WS] Error sending notification: %v", err)
					}
				}
			}

			// Upload to UploadThing in background
			go func(segmentID string, audioPath string, audioData []byte, segment models.AudioSegment) {
				log.Printf("[Upload] Starting UploadThing upload for segment: %s", segmentID)
				uploadURL, err := uploadToUploadThing(audioPath)
				if err != nil {
					log.Printf("[Upload] Error uploading to UploadThing: %v", err)
					return
				}

				// Update segment with UploadThing URL
				segment.AudioURL = uploadURL
				segment.UpdatedAt = time.Now()
				if err := db.UpdateAudioSegment(&segment); err != nil {
					log.Printf("[Upload] Error updating segment with UploadThing URL: %v", err)
					return
				}

				// Clean up local file
				os.Remove(audioPath)
			}(segment.ID, audioPath, audioData, segment)
		}
	}()

	w.WriteHeader(http.StatusAccepted)
	json.NewEncoder(w).Encode(map[string]string{
		"status":  "processing",
		"message": "Audio generation started",
	})
}

func processBookHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	bookID := vars["id"]

	// Get the book
	book, err := db.GetBookByID(bookID)
	if err != nil {
		http.Error(w, "Book not found", http.StatusNotFound)
		return
	}

	// Start processing in background
	go func() {
		// Download the PDF file
		resp, err := http.Get(book.FileURL)
		if err != nil {
			log.Printf("[PDF] Error downloading PDF: %v", err)
			book.Status = "error"
			db.UpdateBook(book)
			return
		}
		defer resp.Body.Close()

		// Process PDF
		processedBook, err := pdf.ProcessPDF(resp.Body, filepath.Base(book.FileURL))
		if err != nil {
			log.Printf("[PDF] Error processing PDF: %v", err)
			book.Status = "error"
			db.UpdateBook(book)
			return
		}

		// Update book with processed information
		book.PageCount = processedBook.PageCount
		book.Author = processedBook.Author
		book.Language = processedBook.Language
		if err := db.UpdateBook(book); err != nil {
			log.Printf("[PDF] Error updating book: %v", err)
			return
		}

		// Create temporary file for PDF processing
		tmpFile, err := os.CreateTemp("", "book-*.pdf")
		if err != nil {
			log.Printf("[PDF] Error creating temp file: %v", err)
			book.Status = "error"
			db.UpdateBook(book)
			return
		}
		defer os.Remove(tmpFile.Name())

		// Copy PDF to temporary file
		if _, err := io.Copy(tmpFile, resp.Body); err != nil {
			log.Printf("[PDF] Error copying to temp file: %v", err)
			book.Status = "error"
			db.UpdateBook(book)
			return
		}

		// Extract text segments
		textSegments, err := pdf.ExtractText(tmpFile.Name())
		if err != nil {
			log.Printf("[PDF] Error extracting text: %v", err)
			book.Status = "error"
			db.UpdateBook(book)
			return
		}

		// Create audio segments
		for i, text := range textSegments {
			segment := &models.AudioSegment{
				ID:        uuid.New().String(),
				BookID:    book.ID,
				Content:   text,
				Status:    "pending",
				CreatedAt: time.Now(),
				UpdatedAt: time.Now(),
			}
			if err := db.SaveAudioSegment(segment); err != nil {
				log.Printf("[PDF] Error saving segment %d: %v", i+1, err)
				continue
			}
		}

		// Update book status
		book.Status = "ready"
		book.UpdatedAt = time.Now()
		if err := db.UpdateBook(book); err != nil {
			log.Printf("[PDF] Error updating book status: %v", err)
			return
		}

		// Get audio segments from the book
		audioSegments, err := db.GetAudioSegments(book.ID)
		if err != nil {
			log.Printf("[Processing] Error getting segments: %v", err)
			return
		}

		// Process each segment
		for _, segment := range audioSegments {
			if segment.Status != "pending" {
				continue
			}

			// Generate audio for the segment
			audioData, err := ttsGen.ProcessAudioSegment(&segment)
			if err != nil {
				log.Printf("[Processing] Error generating audio for segment %s: %v", segment.ID, err)
				segment.Status = "error"
				db.UpdateAudioSegment(&segment)
				continue
			}

			// Save audio to file
			audioFileName := fmt.Sprintf("%s.mp3", segment.ID)
			audioPath := filepath.Join("audio", audioFileName)
			if err := os.WriteFile(audioPath, audioData, 0644); err != nil {
				log.Printf("[Processing] Error saving audio file for segment %s: %v", segment.ID, err)
				segment.Status = "error"
				db.UpdateAudioSegment(&segment)
				continue
			}

			// Update segment with audio URL and status
			segment.AudioURL = audioPath
			segment.Status = "ready"
			segment.UpdatedAt = time.Now()
			if err := db.UpdateAudioSegment(&segment); err != nil {
				log.Printf("[Processing] Error updating segment %s: %v", segment.ID, err)
				continue
			}
		}

		// Update book status to ready
		book.Status = "ready"
		book.UpdatedAt = time.Now()
		if err := db.UpdateBook(book); err != nil {
			log.Printf("[Processing] Error updating book status: %v", err)
		}
	}()

	w.WriteHeader(http.StatusAccepted)
	json.NewEncoder(w).Encode(map[string]string{
		"status":  "processing",
		"message": "Book processing started",
	})
}

func uploadToUploadThing(filePath string) (string, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return "", fmt.Errorf("error opening file: %v", err)
	}
	defer file.Close()

	// Create multipart form
	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)

	// Add file
	part, err := writer.CreateFormFile("file", filepath.Base(filePath))
	if err != nil {
		return "", fmt.Errorf("error creating form file: %v", err)
	}
	if _, err := io.Copy(part, file); err != nil {
		return "", fmt.Errorf("error copying file: %v", err)
	}
	writer.Close()

	// Create request
	req, err := http.NewRequest("POST", config.AppConfig.UploadThingURL, body)
	if err != nil {
		return "", fmt.Errorf("error creating request: %v", err)
	}

	// Set headers
	req.Header.Set("Content-Type", writer.FormDataContentType())
	req.Header.Set("Authorization", "Bearer "+config.AppConfig.UploadThingToken)

	// Send request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", fmt.Errorf("error uploading to UploadThing: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("UploadThing error: %s", resp.Status)
	}

	// Parse response
	var result struct {
		URL string `json:"url"`
	}
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return "", fmt.Errorf("error decoding response: %v", err)
	}

	return result.URL, nil
}
