package response

import (
	"mailmind-api/internal/models"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type DraftResponse struct {
	ID        primitive.ObjectID `json:"id"`
	Recipient string             `json:"recipient"`
	Subject   string             `json:"subject"`
	Body      string             `json:"body"`
	CreatedAt time.Time          `json:"created_at"`
}

func ToDraftResponse(draft *models.Draft) DraftResponse {
	return DraftResponse{
		ID:        draft.ID,
		Recipient: draft.Recipient,
		Subject:   draft.Subject,
		Body:      draft.Body,
		CreatedAt: draft.CreatedAt,
	}
}
