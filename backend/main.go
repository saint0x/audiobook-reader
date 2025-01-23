package main

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"github.com/gorilla/websocket"
	_ "github.com/mattn/go-sqlite3"
)

var db *sql.DB
var fileStorage *FileStorage

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
	if err := LoadConfig(); err != nil {
		log.Fatal("Error loading configuration:", err)
	}

	// Initialize database
	var err error
	db, err = sql.Open("sqlite3", AppConfig.DBPath)
	if err != nil {
		log.Fatal("Error opening database:", err)
	}
	defer db.Close()

	// Create tables
	if err := InitDB(); err != nil {
		log.Fatal("Error initializing database:", err)
	}

	// Clean up any duplicate books
	if err := CleanupDuplicateBooks(db); err != nil {
		log.Printf("Warning: Error cleaning up duplicates: %v", err)
	}

	// Initialize file storage
	fileStorage, err = NewFileStorage(AppConfig.UploadDir)
	if err != nil {
		log.Fatal("Error initializing file storage:", err)
	}

	// Create audio directory if it doesn't exist
	audioDir := filepath.Join(AppConfig.UploadDir, "audio")
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
	audioDir = filepath.Join(AppConfig.UploadDir, "audio")
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
	book := &Book{
		ID:        uuid.New().String(),
		Title:     req.Title,
		FileURL:   req.FileURL,
		Status:    "processing",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	log.Printf("[Upload] Created book record with ID: %s", book.ID)

	// Save initial book record
	if err := SaveBook(db, book); err != nil {
		log.Printf("[Upload] Error saving book: %v", err)
		http.Error(w, fmt.Sprintf("Error saving book: %v", err), http.StatusInternalServerError)
		return
	}
	log.Printf("[Upload] Successfully saved book to database")

	// Start processing in background immediately
	go func() {
		log.Printf("[Processing] Starting background processing for book: %s", book.ID)

		// Process PDF first to create segments
		log.Printf("[PDF] Starting PDF processing for book: %s", book.ID)
		if err := processBook(db, book); err != nil {
			log.Printf("[PDF] Error processing PDF: %v", err)
			book.Status = "error"
			UpdateBook(db, book)
			return
		}
		log.Printf("[PDF] Successfully processed PDF for book: %s", book.ID)

		// Now start TTS processing
		log.Printf("[TTS] Starting TTS processing for book: %s", book.ID)
		if err := processTextToSpeech(book); err != nil {
			log.Printf("[TTS] Error processing TTS: %v", err)
			book.Status = "error"
			UpdateBook(db, book)
			return
		}
		log.Printf("[TTS] Successfully processed TTS for book: %s", book.ID)

		log.Printf("[Processing] All processing completed successfully for book: %s", book.ID)
		// Update book status to ready
		book.Status = "ready"
		book.UpdatedAt = time.Now()
		if err := UpdateBook(db, book); err != nil {
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

// Update processTextToSpeech to notify clients
func processTextToSpeech(book *Book) error {
	log.Printf("[TTS] Starting TTS processing for book: %s", book.ID)

	segments, err := GetAudioSegments(book.ID)
	if err != nil {
		log.Printf("[TTS] Error getting audio segments: %v", err)
		return fmt.Errorf("error getting audio segments: %v", err)
	}
	log.Printf("[TTS] Found %d segments to process", len(segments))

	for i, segment := range segments {
		if segment.Status != "pending" {
			log.Printf("[TTS] Skipping segment %d/%d (ID: %s) - status: %s", i+1, len(segments), segment.ID, segment.Status)
			continue
		}

		if len(strings.TrimSpace(segment.Content)) == 0 {
			log.Printf("[TTS] Skipping empty segment %d/%d (ID: %s)", i+1, len(segments), segment.ID)
			segment.Status = "skipped"
			segment.UpdatedAt = time.Now()
			if err := UpdateAudioSegment(db, &segment); err != nil {
				log.Printf("[TTS] Error updating empty segment status: %v", err)
			}
			continue
		}

		log.Printf("[TTS] Processing segment %d/%d (ID: %s) - Content length: %d", i+1, len(segments), segment.ID, len(segment.Content))
		audioData, err := generateTTSAudio(segment.Content)
		if err != nil {
			log.Printf("[TTS] Error generating audio: %v", err)
			segment.Status = "error"
			UpdateAudioSegment(db, &segment)
			return fmt.Errorf("error generating TTS: %v", err)
		}
		log.Printf("[TTS] Successfully generated audio for segment: %s", segment.ID)

		// Save audio to temporary file for immediate playback
		audioFileName := fmt.Sprintf("tts-%s.mp3", uuid.New().String())
		audioPath := filepath.Join(AppConfig.UploadDir, "audio", audioFileName)
		if err := os.WriteFile(audioPath, audioData, 0644); err != nil {
			log.Printf("[TTS] Error writing audio file: %v", err)
			return fmt.Errorf("error writing audio file: %v", err)
		}
		log.Printf("[TTS] Saved audio file: %s", audioPath)

		// Update segment with local URL and notify frontend immediately
		segment.AudioURL = "/audio/" + audioFileName
		segment.Status = "completed"
		segment.UpdatedAt = time.Now()
		if err := UpdateAudioSegment(db, &segment); err != nil {
			log.Printf("[TTS] Error updating audio segment: %v", err)
			return fmt.Errorf("error updating audio segment: %v", err)
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
			log.Printf("[WS] Notified clients about new audio segment: %s", segment.ID)
		}

		// Upload to UploadThing in background
		go func(segmentID string, audioPath string, audioData []byte) {
			log.Printf("[Upload] Starting UploadThing upload for segment: %s", segmentID)
			uploadURL, err := uploadToUploadThing(audioPath)
			if err != nil {
				log.Printf("[Upload] Error uploading to UploadThing: %v", err)
				return
			}
			log.Printf("[Upload] Successfully uploaded to UploadThing: %s", uploadURL)

			// Update segment with UploadThing URL
			segment.AudioURL = uploadURL
			segment.UpdatedAt = time.Now()
			if err := UpdateAudioSegment(db, &segment); err != nil {
				log.Printf("[Upload] Error updating segment with UploadThing URL: %v", err)
				return
			}

			// Clean up local file
			os.Remove(audioPath)
			log.Printf("[Upload] Cleaned up local file: %s", audioPath)
		}(segment.ID, audioPath, audioData)
	}

	log.Printf("[TTS] Completed TTS processing for all segments of book: %s", book.ID)
	return nil
}

func generateTTSAudio(text string) ([]byte, error) {
	text = strings.TrimSpace(text)
	if text == "" {
		return nil, fmt.Errorf("cannot generate audio for empty text")
	}

	log.Printf("[TTS] Starting TTS generation for text length: %d", len(text))

	// Call Replicate API to generate audio
	replicateURL := AppConfig.ReplicateAPIURL + "/predictions"

	requestBody := map[string]interface{}{
		"version": AppConfig.KokoroModelVersion,
		"input": map[string]interface{}{
			"text":     text,
			"language": "en",
		},
	}

	jsonData, err := json.Marshal(requestBody)
	if err != nil {
		return nil, fmt.Errorf("error marshaling request: %v", err)
	}

	// Create initial prediction
	req, err := http.NewRequest("POST", replicateURL, bytes.NewBuffer(jsonData))
	if err != nil {
		return nil, fmt.Errorf("error creating request: %v", err)
	}

	req.Header.Set("Authorization", "Token "+os.Getenv("REPLICATE_API_TOKEN"))
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error making request: %v", err)
	}
	defer resp.Body.Close()

	var prediction struct {
		ID     string   `json:"id"`
		Status string   `json:"status"`
		Output []string `json:"output"`
	}
	if err := json.NewDecoder(resp.Body).Decode(&prediction); err != nil {
		return nil, fmt.Errorf("error decoding response: %v", err)
	}

	// Poll for completion
	maxAttempts := 60 // 2 minutes total
	for i := 0; i < maxAttempts; i++ {
		time.Sleep(2 * time.Second)

		req, err = http.NewRequest("GET", replicateURL+"/"+prediction.ID, nil)
		if err != nil {
			return nil, fmt.Errorf("error creating poll request: %v", err)
		}
		req.Header.Set("Authorization", "Token "+os.Getenv("REPLICATE_API_TOKEN"))

		resp, err = client.Do(req)
		if err != nil {
			return nil, fmt.Errorf("error polling prediction: %v", err)
		}

		var result struct {
			Status string   `json:"status"`
			Output []string `json:"output"`
			Error  string   `json:"error"`
		}
		if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
			resp.Body.Close()
			return nil, fmt.Errorf("error decoding poll response: %v", err)
		}
		resp.Body.Close()

		log.Printf("[TTS] Poll status: %s", result.Status)

		switch result.Status {
		case "succeeded":
			if len(result.Output) == 0 {
				return nil, fmt.Errorf("no output from model")
			}

			// Download the audio file
			audioResp, err := http.Get(result.Output[0])
			if err != nil {
				return nil, fmt.Errorf("error downloading audio: %v", err)
			}
			defer audioResp.Body.Close()

			return io.ReadAll(audioResp.Body)

		case "failed":
			return nil, fmt.Errorf("prediction failed: %s", result.Error)

		case "canceled":
			return nil, fmt.Errorf("prediction was canceled")

		case "completed":
			if len(result.Output) == 0 {
				return nil, fmt.Errorf("no output from model")
			}

			// Download the audio file
			audioResp, err := http.Get(result.Output[0])
			if err != nil {
				return nil, fmt.Errorf("error downloading audio: %v", err)
			}
			defer audioResp.Body.Close()

			return io.ReadAll(audioResp.Body)
		}
	}

	return nil, fmt.Errorf("prediction timed out after %d attempts", maxAttempts)
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
	book, err := GetBookByID(db, vars["id"])
	if err != nil {
		http.Error(w, "Book not found", http.StatusNotFound)
		return
	}

	json.NewEncoder(w).Encode(book)
}

func getBooksHandler(w http.ResponseWriter, r *http.Request) {
	query := `
		SELECT id, title, author, cover_url, page_count, language, created_at
		FROM books
		ORDER BY created_at DESC
	`
	rows, err := db.Query(query)
	if err != nil {
		http.Error(w, "Error retrieving books", http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var books []Book
	for rows.Next() {
		var book Book
		err := rows.Scan(
			&book.ID,
			&book.Title,
			&book.Author,
			&book.CoverURL,
			&book.PageCount,
			&book.Language,
			&book.CreatedAt,
		)
		if err != nil {
			http.Error(w, "Error scanning books", http.StatusInternalServerError)
			return
		}
		books = append(books, book)
	}

	json.NewEncoder(w).Encode(books)
}

func updateProgressHandler(w http.ResponseWriter, r *http.Request) {
	var progress ReadingProgress
	if err := json.NewDecoder(r.Body).Decode(&progress); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	// Generate ID if not provided
	if progress.ID == "" {
		progress.ID = uuid.New().String()
	}

	if err := UpdateReadingProgress(db, &progress); err != nil {
		http.Error(w, "Error updating progress", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(progress)
}

func getProgressHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userID := r.Header.Get("X-User-ID") // Get user ID from header

	progress, err := GetReadingProgress(db, vars["bookId"], userID)
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
	var bookmark Bookmark
	if err := json.NewDecoder(r.Body).Decode(&bookmark); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	// Generate ID if not provided
	if bookmark.ID == "" {
		bookmark.ID = uuid.New().String()
	}

	if err := CreateBookmark(db, &bookmark); err != nil {
		http.Error(w, "Error creating bookmark", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(bookmark)
}

func getBookmarksHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userID := r.Header.Get("X-User-ID") // Get user ID from header

	bookmarks, err := GetBookmarks(db, vars["bookId"], userID)
	if err != nil {
		http.Error(w, "Error retrieving bookmarks", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(bookmarks)
}

func updateBookmarkHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	var bookmark Bookmark
	if err := json.NewDecoder(r.Body).Decode(&bookmark); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	bookmark.ID = vars["id"]
	if err := UpdateBookmark(db, &bookmark); err != nil {
		http.Error(w, "Error updating bookmark", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(bookmark)
}

func deleteBookmarkHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	if err := DeleteBookmark(db, vars["id"]); err != nil {
		http.Error(w, "Error deleting bookmark", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func generateAudioHandler(w http.ResponseWriter, r *http.Request) {
	var segment AudioSegment
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

	if err := CreateAudioSegment(db, &segment); err != nil {
		http.Error(w, "Error creating audio segment", http.StatusInternalServerError)
		return
	}

	// Start TTS processing in background
	go processAudioSegment(db, &segment)

	json.NewEncoder(w).Encode(segment)
}

func getAudioSegmentsHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	bookID := vars["id"]

	segments, err := GetAudioSegments(bookID)
	if err != nil {
		http.Error(w, fmt.Sprintf("Failed to get audio segments: %v", err), http.StatusInternalServerError)
		return
	}

	// Initialize empty array if segments is nil
	if segments == nil {
		segments = []AudioSegment{}
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(segments)
}

func getCategoriesHandler(w http.ResponseWriter, r *http.Request) {
	query := "SELECT id, name, description, created_at FROM categories ORDER BY name"
	rows, err := db.Query(query)
	if err != nil {
		http.Error(w, "Error retrieving categories", http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var categories []Category
	for rows.Next() {
		var category Category
		err := rows.Scan(
			&category.ID,
			&category.Name,
			&category.Description,
			&category.CreatedAt,
		)
		if err != nil {
			http.Error(w, "Error scanning categories", http.StatusInternalServerError)
			return
		}
		categories = append(categories, category)
	}

	json.NewEncoder(w).Encode(categories)
}

func getTagsHandler(w http.ResponseWriter, r *http.Request) {
	query := "SELECT id, name, created_at FROM tags ORDER BY name"
	rows, err := db.Query(query)
	if err != nil {
		http.Error(w, "Error retrieving tags", http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var tags []Tag
	for rows.Next() {
		var tag Tag
		err := rows.Scan(
			&tag.ID,
			&tag.Name,
			&tag.CreatedAt,
		)
		if err != nil {
			http.Error(w, "Error scanning tags", http.StatusInternalServerError)
			return
		}
		tags = append(tags, tag)
	}

	json.NewEncoder(w).Encode(tags)
}

func getBookStatusHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	book, err := GetBookByID(db, vars["id"])
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

	book, err := GetBookByID(db, vars["id"])
	if err != nil {
		http.Error(w, "Book not found", http.StatusNotFound)
		return
	}

	book.FileURL = req.CloudURL
	if err := UpdateBook(db, book); err != nil {
		http.Error(w, "Error updating book", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func generateBookAudioHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	bookID := vars["id"]

	// Get the book
	book, err := GetBookByID(db, bookID)
	if err != nil {
		http.Error(w, "Book not found", http.StatusNotFound)
		return
	}

	// Start TTS processing in background
	go func() {
		if err := processTextToSpeech(book); err != nil {
			log.Printf("Error processing TTS for book %s: %v", book.ID, err)
			return
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
	book, err := GetBookByID(db, bookID)
	if err != nil {
		http.Error(w, "Book not found", http.StatusNotFound)
		return
	}

	// Start processing in background
	go func() {
		if err := processBook(db, book); err != nil {
			log.Printf("Error processing book: %v", err)
			return
		}
	}()

	w.WriteHeader(http.StatusAccepted)
	json.NewEncoder(w).Encode(map[string]string{
		"status":  "processing",
		"message": "Audio generation started",
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
	req, err := http.NewRequest("POST", AppConfig.UploadThingURL, body)
	if err != nil {
		return "", fmt.Errorf("error creating request: %v", err)
	}

	// Set headers
	req.Header.Set("Content-Type", writer.FormDataContentType())
	req.Header.Set("Authorization", "Bearer "+AppConfig.UploadThingToken)

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
