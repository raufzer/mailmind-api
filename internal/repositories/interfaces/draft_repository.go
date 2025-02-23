package interfaces

import (
	"context"
	"mailmind-api/internal/models"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type DraftRepository interface {
	SaveDraft(ctx context.Context, draft *models.Draft) error
	GetDraftByID(ctx context.Context, draftID primitive.ObjectID) (*models.Draft, error)
}
