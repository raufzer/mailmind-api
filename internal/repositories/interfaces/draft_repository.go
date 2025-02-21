package interfaces

import (
	"context"
	"mailmind-api/internal/models"

)
type DraftRepository interface {
	CreateDraft(ctx context.Context, draft *models.Draft) error
	GetDraftByID(ctx context.Context, draftID string) (*models.Draft, error)
}
