package interfaces

import (
	"context"
	"mailmind-api/internal/dto/request"
	"mailmind-api/internal/models"
)

type UserService interface {
	GetUserByID(ctx context.Context, userID string) (*models.User, error)
	GetUserByEmail(ctx context.Context, email string) (*models.User, error)
	UpdateUserSettings(ctx context.Context, userID string, settings request.UpdateUserSettingsRequest) error
}
