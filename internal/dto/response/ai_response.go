package response

import (
	"mailmind-api/internal/models"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type AIResponse struct {
	ID             primitive.ObjectID `json:"id"`
	EmailID        primitive.ObjectID `json:"email_id"`
	GeneratedReply string             `json:"generated_reply"`
	CreatedAt      time.Time          `json:"created_at"`
}

func ToAIResponse(ai *models.AIResponse) AIResponse {
	return AIResponse{
		ID:             ai.ID,
		EmailID:        ai.EmailID,
		GeneratedReply: ai.GeneratedReply,
		CreatedAt:      ai.CreatedAt,
	}
}
