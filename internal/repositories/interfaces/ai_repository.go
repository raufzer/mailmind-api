package interfaces

import (
	"context"
	"mailmind-api/internal/models"

)
type AIRepository interface {
	SaveAIResponse(ctx context.Context, response *models.AIResponse) error
	GetAIResponseByEmailID(ctx context.Context, emailID string) (*models.AIResponse, error)
}