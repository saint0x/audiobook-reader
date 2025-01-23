package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/joho/godotenv"
)

type Config struct {
	// Server
	Port       string
	BackendURL string
	Env        string

	// Database
	DBPath string

	// File Storage
	UploadDir     string
	PDFDir        string
	CoverDir      string
	MaxUploadSize int64

	// Replicate API
	ReplicateAPIToken  string
	ReplicateAPIURL    string
	KokoroModelVersion string

	// UploadThing
	UploadThingURL    string
	UploadThingToken  string
	UploadThingSecret string
	UploadThingAppID  string

	// TTS
	TTSMaxChunkSize   int
	TTSDefaultLang    string
	TTSDefaultSpeaker int
	TTSDefaultSpeed   float64
	TTSTopK           int
	TTSTopP           float64
	TTSTemperature    float64

	// CORS
	AllowedOrigins []string
}

var AppConfig Config

// LoadConfig loads the configuration from environment variables
func LoadConfig() error {
	// Load .env file if it exists
	godotenv.Load()

	AppConfig = Config{
		// Server
		Port:       getEnv("PORT", "8080"),
		BackendURL: getEnv("BACKEND_URL", "http://localhost:8080"),
		Env:        getEnv("ENV", "development"),

		// Database
		DBPath: getEnv("DB_PATH", "./ereader.db"),

		// File Storage
		UploadDir:     getEnv("UPLOAD_DIR", "./uploads"),
		PDFDir:        getEnv("PDF_DIR", "./uploads/pdfs"),
		CoverDir:      getEnv("COVER_DIR", "./uploads/covers"),
		MaxUploadSize: getEnvAsInt64("MAX_UPLOAD_SIZE", 32) * 1024 * 1024, // Convert MB to bytes

		// Replicate API
		ReplicateAPIToken:  getEnvRequired("REPLICATE_API_TOKEN"),
		ReplicateAPIURL:    getEnv("REPLICATE_API_URL", "https://api.replicate.com/v1"),
		KokoroModelVersion: getEnv("KOKORO_MODEL_VERSION", "82c9a78a6fff5fa0e4acbd0f8eb453e0c4b6d5d90e316b544b8c7e04c42fba42"),

		// UploadThing
		UploadThingURL:    getEnv("UPLOADTHING_URL", "https://uploadthing.com/api"),
		UploadThingToken:  getEnvRequired("UPLOADTHING_TOKEN"),
		UploadThingSecret: getEnvRequired("UPLOADTHING_SECRET"),
		UploadThingAppID:  getEnvRequired("UPLOADTHING_APP_ID"),

		// TTS
		TTSMaxChunkSize:   getEnvAsInt("TTS_MAX_CHUNK_SIZE", 1000),
		TTSDefaultLang:    getEnv("TTS_DEFAULT_LANGUAGE", "en"),
		TTSDefaultSpeaker: getEnvAsInt("TTS_DEFAULT_SPEAKER_ID", 0),
		TTSDefaultSpeed:   getEnvAsFloat64("TTS_DEFAULT_SPEED", 1.0),
		TTSTopK:           getEnvAsInt("TTS_TOP_K", 50),
		TTSTopP:           getEnvAsFloat64("TTS_TOP_P", 0.8),
		TTSTemperature:    getEnvAsFloat64("TTS_TEMPERATURE", 0.8),

		// CORS
		AllowedOrigins: strings.Split(getEnv("ALLOWED_ORIGINS", "http://localhost:3000"), ","),
	}

	// Create required directories
	for _, dir := range []string{AppConfig.UploadDir, AppConfig.PDFDir, AppConfig.CoverDir} {
		if err := os.MkdirAll(dir, 0755); err != nil {
			return fmt.Errorf("error creating directory %s: %v", dir, err)
		}
	}

	return nil
}

// Helper functions to get environment variables
func getEnv(key, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultValue
}

func getEnvRequired(key string) string {
	value, exists := os.LookupEnv(key)
	if !exists {
		panic(fmt.Sprintf("Required environment variable %s is not set", key))
	}
	return value
}

func getEnvAsInt(key string, defaultValue int) int {
	valueStr := getEnv(key, "")
	if value, err := strconv.Atoi(valueStr); err == nil {
		return value
	}
	return defaultValue
}

func getEnvAsInt64(key string, defaultValue int64) int64 {
	valueStr := getEnv(key, "")
	if value, err := strconv.ParseInt(valueStr, 10, 64); err == nil {
		return value
	}
	return defaultValue
}

func getEnvAsFloat64(key string, defaultValue float64) float64 {
	valueStr := getEnv(key, "")
	if value, err := strconv.ParseFloat(valueStr, 64); err == nil {
		return value
	}
	return defaultValue
}
