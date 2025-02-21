package interfaces

import (
	"context"
	"mailmind-api/internal/models"
)

type AuthService interface {
	Logout(ctx context.Context, userID string) error
	ValidateToken(ctx context.Context, token string) (string, error)
	GoogleConnect(ctx context.Context, code string) (*models.User, string, string, error)
}
