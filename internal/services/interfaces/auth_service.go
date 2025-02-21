package interfaces

import (
	"context"
	"mailmind-api/internal/models"
)

type AuthService interface {
	Logout(ctx context.Context, userID string) error
	RefreshAccessToken(ctx context.Context, userID, refreshToken string) (string, error)
	ValidateToken(ctx context.Context, token string) (string, error)
	GoogleConnect(ctx context.Context, code string) (*models.User, string, string, string, error)
}
