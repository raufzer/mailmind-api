package interfaces

import (
	"context"
	"mailmind-api/internal/models"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type EmailService interface {
	SendEmail(ctx context.Context, email *models.Email) error
	SaveDraft(ctx context.Context, draft *models.Draft) error
	GetDraft(ctx context.Context, draftID primitive.ObjectID) (*models.Draft, error)
}
