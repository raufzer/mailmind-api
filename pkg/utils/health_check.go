package utils

import (
	"context"
	"database/sql"
	"fmt"
	"net/http"

	"github.com/cloudinary/cloudinary-go/v2"
	"github.com/go-redis/redis/v8"
	"golang.org/x/oauth2"
)

func CheckDatabaseHealth(db *sql.DB) error {
	if err := db.Ping(); err != nil {
		return fmt.Errorf("database health check failed: %v", err)
	}
	return nil
}

func CheckCacheHealth(redisClient *redis.Client) error {
	ctx := context.Background()
	if err := redisClient.Ping(ctx).Err(); err != nil {
		return fmt.Errorf("cache health check failed: %w", err)
	}
	return nil
}

func CheckCloudinaryHealth(cloudName, apiKey, apiSecret string) error {
	cld, err := cloudinary.NewFromParams(cloudName, apiKey, apiSecret)
	if err != nil {

		return fmt.Errorf("failed to create Cloudinary client: %v", err)
	}

	_, err = cld.Admin.Ping(context.Background())
	if err != nil {

		return fmt.Errorf("failed to ping Cloudinary: %v", err)
	}

	return nil
}

func CheckGoogleOAuthHealth(oauthConfig *oauth2.Config) error {

	url := oauthConfig.AuthCodeURL("", oauth2.AccessTypeOffline)

	resp, err := http.Get(url)
	if err != nil {

		return fmt.Errorf("failed to reach Google OAuth authorization endpoint: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("google OAuth returned error with status code %d", resp.StatusCode)
	}

	return nil
}
