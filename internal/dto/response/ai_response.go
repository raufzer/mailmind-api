package response

import (
	"mailmind-api/internal/models"
	"time"
)

type AIResponse struct {
	ID             string    `json:"id"`
	EmailID        string    `json:"email_id"`
	GeneratedReply string    `json:"generated_reply"`
	CreatedAt      time.Time `json:"created_at"`
}

func ToAIResponse(ai *models.AIResponse) AIResponse {
	return AIResponse{
		ID:             ai.ID,
		EmailID:        ai.EmailID,
		GeneratedReply: ai.GeneratedReply,
		CreatedAt:      ai.CreatedAt,
	}
}
