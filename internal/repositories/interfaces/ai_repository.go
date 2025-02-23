package interfaces

import (
	"context"
	"mailmind-api/internal/models"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type AIResponseRepository interface {
	SaveResponse(ctx context.Context, response *models.AIResponse) error
	GetResponseByEmailID(ctx context.Context, emailID primitive.ObjectID) (*models.AIResponse, error)
}
