package interfaces

import (
	"context"
	"mailmind-api/internal/models"


)

type AIService interface {
	GenerateReply(ctx context.Context, content string) (*models.AIResponse, error)
}
