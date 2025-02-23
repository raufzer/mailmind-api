package interfaces

import (
	"context"
	"mailmind-api/internal/models"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type EmailRepository interface {
	SaveEmail(ctx context.Context, email *models.Email) error
	GetEmailByID(ctx context.Context, emailID primitive.ObjectID) (*models.Email, error)
}

