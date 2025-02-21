package interfaces

import (
	"context"
	"mailmind-api/internal/models"
)

type UserRepository interface {
	CreateUser(ctx context.Context, user *models.User) error
	GetUserByID(ctx context.Context, userID string) (*models.User, error)
	GetUserByEmail(ctx context.Context, email string) (*models.User, error)
	UpdateUserSettings(ctx context.Context, userID string, settings *models.UserSettings) error
}
