package config

import (
	"log"
	"mailmind-api/pkg/utils"
	"time"

	"github.com/joho/godotenv"
)

type AppConfig struct {
	BackEndDomain       string
	ServerPort          string
	DatabaseURI         string
	DatabaseName        string
	AccessTokenSecret   string
	AccessTokenMaxAge   time.Duration
	GoogleClientID      string
	GoogleClientSecret  string
	GoogleRedirectURL   string
	GeminiAPIKey        string
}

func LoadConfig() (*AppConfig, error) {
	err := godotenv.Load("./.env")
	if err != nil {
		log.Println("Warning: No .env file found, using default environment variables.")
	}
	config := &AppConfig{
		BackEndDomain:      getEnvOrFatal("BACK_END_DOMAIN", "string").(string),
		ServerPort:         getEnvOrFatal("SERVER_PORT", "string").(string),
		DatabaseURI:        getEnvOrFatal("DATABASE_URI", "string").(string),
		DatabaseName:       getEnvOrFatal("DATABASE_NAME", "string").(string),
		AccessTokenSecret:  getEnvOrFatal("ACCESS_TOKEN_SECRET", "string").(string),
		AccessTokenMaxAge:  getEnvOrFatal("ACCESS_TOKEN_MAX_AGE", "duration").(time.Duration),
		GoogleClientID:     getEnvOrFatal("GOOGLE_CLIENT_ID", "string").(string),
		GoogleClientSecret: getEnvOrFatal("GOOGLE_CLIENT_SECRET", "string").(string),
		GoogleRedirectURL:  getEnvOrFatal("GOOGLE_REDIRECT_URL", "string").(string),
		GeminiAPIKey:       getEnvOrFatal("GEMINI_API_KEY", "string").(string),
	}
	return config, nil
}

func getEnvOrFatal(key, expectedType string) interface{} {
	val, err := utils.GetEnv(key, expectedType)
	if err != nil {
		log.Fatalf("Failed to get environment variable %s: %v", key, err)
	}
	return val
}
