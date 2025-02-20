package interfaces

import (
	"context"
	"mailmind-api/internal/models"

	"github.com/google/uuid"
)

type UserRepository interface {
	CreateUser(ctx context.Context, user *models.User) error
	GetUserByEmail(ctx context.Context, email string) (*models.User, error)
	GetUserByID(ctx context.Context, userID uuid.UUID) (*models.User, error)
	UpdateUserSettings(ctx context.Context, userID uuid.UUID, settings models.UserSettings) error
	DeleteUser(ctx context.Context, userID uuid.UUID) error
}
