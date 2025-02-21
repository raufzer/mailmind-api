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
	RedisURI            string
	RedisPassword       string
	AccessTokenSecret   string
	RefreshTokenSecret  string
	AccessTokenMaxAge   time.Duration
	RefreshTokenMaxAge  time.Duration
	GoogleClientID      string
	GoogleClientSecret  string
	GoogleRedirectURL   string
	CloudinaryCloudName string
	CloudinaryAPIKey    string
	CloudinaryAPISecret string
}

func LoadConfig() (*AppConfig, error) {
	err := godotenv.Load("./.env")
	if err != nil {
		log.Println("Warning: No .env file found, using default environment variables.")
	}
	config := &AppConfig{
		BackEndDomain:       getEnvOrFatal("BACK_END_DOMAIN", "string").(string),
		ServerPort:          getEnvOrFatal("SERVER_PORT", "string").(string),
		DatabaseURI:         getEnvOrFatal("DATABASE_URI", "string").(string),
		DatabaseName:        getEnvOrFatal("DATABASE_NAME", "string").(string),
		RedisURI:            getEnvOrFatal("REDIS_URI", "string").(string),
		RedisPassword:       getEnvOrFatal("REDIS_PASSWORD", "string").(string),
		AccessTokenSecret:   getEnvOrFatal("ACCESS_TOKEN_SECRET", "string").(string),
		RefreshTokenSecret:  getEnvOrFatal("REFRESH_TOKEN_SECRET", "string").(string),
		AccessTokenMaxAge:   getEnvOrFatal("ACCESS_TOKEN_MAX_AGE", "duration").(time.Duration),
		RefreshTokenMaxAge:  getEnvOrFatal("REFRESH_TOKEN_MAX_AGE", "duration").(time.Duration),
		GoogleClientID:      getEnvOrFatal("GOOGLE_CLIENT_ID", "string").(string),
		GoogleClientSecret:  getEnvOrFatal("GOOGLE_CLIENT_SECRET", "string").(string),
		GoogleRedirectURL:   getEnvOrFatal("GOOGLE_REDIRECT_URL", "string").(string),
		CloudinaryCloudName: getEnvOrFatal("CLOUDINARY_CLOUD_NAME", "string").(string),
		CloudinaryAPIKey:    getEnvOrFatal("CLOUDINARY_API_KEY", "string").(string),
		CloudinaryAPISecret: getEnvOrFatal("CLOUDINARY_API_SECRET", "string").(string),
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
