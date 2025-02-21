package interfaces

import (
	"context"
	"mailmind-api/internal/dto/request"
	"mailmind-api/internal/models"

)

type AIService interface {
	GenerateReply(ctx context.Context, req request.GenerateReplyRequest, userID string) (*models.AIResponse, error)
}
