package config

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/joho/godotenv"
)

// Config holds all configuration for the application
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
	if err := godotenv.Load(); err != nil {
		return fmt.Errorf("error loading .env file: %v", err)
	}

	AppConfig = Config{
		Port:       getEnv("PORT", "8080"),
		BackendURL: getEnv("BACKEND_URL", "http://localhost:8080"),
		Env:        getEnv("ENV", "development"),

		DBPath: getEnv("DB_PATH", "./ereader.db"),

		UploadDir:     getEnv("UPLOAD_DIR", "./uploads"),
		PDFDir:        getEnv("PDF_DIR", "./uploads/pdfs"),
		CoverDir:      getEnv("COVER_DIR", "./uploads/covers"),
		MaxUploadSize: getEnvInt64("MAX_UPLOAD_SIZE", 10<<20), // 10MB default

		ReplicateAPIToken:  getEnv("REPLICATE_API_TOKEN", ""),
		ReplicateAPIURL:    getEnv("REPLICATE_API_URL", "https://api.replicate.com/v1"),
		KokoroModelVersion: getEnv("KOKORO_MODEL_VERSION", ""),

		UploadThingURL:    getEnv("UPLOADTHING_URL", ""),
		UploadThingToken:  getEnv("UPLOADTHING_TOKEN", ""),
		UploadThingSecret: getEnv("UPLOADTHING_SECRET", ""),
		UploadThingAppID:  getEnv("UPLOADTHING_APP_ID", ""),

		TTSMaxChunkSize:   getEnvInt("TTS_MAX_CHUNK_SIZE", 1000),
		TTSDefaultLang:    getEnv("TTS_DEFAULT_LANG", "en"),
		TTSDefaultSpeaker: getEnvInt("TTS_DEFAULT_SPEAKER", 0),
		TTSDefaultSpeed:   getEnvFloat64("TTS_DEFAULT_SPEED", 1.0),
		TTSTopK:           getEnvInt("TTS_TOP_K", 50),
		TTSTopP:           getEnvFloat64("TTS_TOP_P", 0.8),
		TTSTemperature:    getEnvFloat64("TTS_TEMPERATURE", 0.8),

		AllowedOrigins: strings.Split(getEnv("ALLOWED_ORIGINS", "*"), ","),
	}

	return nil
}

func getEnv(key, fallback string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return fallback
}

func getEnvInt(key string, fallback int) int {
	if value, exists := os.LookupEnv(key); exists {
		if intVal, err := strconv.Atoi(value); err == nil {
			return intVal
		}
	}
	return fallback
}

func getEnvInt64(key string, fallback int64) int64 {
	if value, exists := os.LookupEnv(key); exists {
		if intVal, err := strconv.ParseInt(value, 10, 64); err == nil {
			return intVal
		}
	}
	return fallback
}

func getEnvFloat64(key string, fallback float64) float64 {
	if value, exists := os.LookupEnv(key); exists {
		if floatVal, err := strconv.ParseFloat(value, 64); err == nil {
			return floatVal
		}
	}
	return fallback
}
