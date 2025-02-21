package interfaces

import (
	"context"
	"mailmind-api/internal/models"

)

type EmailRepository interface {
	CreateEmail(ctx context.Context, email *models.Email) error
	GetEmailByID(ctx context.Context, emailID string) (*models.Email, error)
}