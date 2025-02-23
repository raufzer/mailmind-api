package response

import (
	"mailmind-api/internal/models"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type GeminiAPIResponse struct {
	Text string `json:"text"`
}
type AIResponse struct {
	ID             primitive.ObjectID `json:"id"`
	GeneratedReply string             `json:"generated_reply"`
	CreatedAt      time.Time          `json:"created_at"`
}

func ToAIResponse(ai *models.AIResponse) AIResponse {
	return AIResponse{
		ID:             ai.ID,
		GeneratedReply: ai.GeneratedReply,
		CreatedAt:      ai.CreatedAt,
	}
}
