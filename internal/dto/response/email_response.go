package response

import (
	"mailmind-api/internal/models"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type DraftResponse struct {
	ID        primitive.ObjectID `json:"id"`
	UserID    primitive.ObjectID `json:"user_id"`
	To        string           `json:"to"`
	Subject   string             `json:"subject"`
	Body      string             `json:"body"`
	CreatedAt time.Time          `json:"created_at"`
	UpdatedAt time.Time          `json:"updated_at"`
}

func ToDraftResponse(draft *models.Draft) DraftResponse {
	return DraftResponse{
		ID:        draft.ID,
		UserID:    draft.UserID,
		To:        draft.To,
		Subject:   draft.Subject,
		Body:      draft.Body,
		CreatedAt: draft.CreatedAt,
		UpdatedAt: draft.UpdatedAt,
	}
}
