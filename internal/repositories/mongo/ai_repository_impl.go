package mongo

import (
	"context"
	"mailmind-api/internal/models"
	"mailmind-api/internal/repositories/interfaces"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type AIRepositoryMongo struct {
	collection *mongo.Collection
}

func NewAIRepository(db *mongo.Database) interfaces.AIRepository {
	return &AIRepositoryMongo{collection: db.Collection("ai_responses")}
}

func (r *AIRepositoryMongo) SaveAIResponse(ctx context.Context, response *models.AIResponse) error {
	_, err := r.collection.InsertOne(ctx, response)
	return err
}

func (r *AIRepositoryMongo) GetAIResponseByEmailID(ctx context.Context, emailID string) (*models.AIResponse, error) {
	var response models.AIResponse
	err := r.collection.FindOne(ctx, bson.M{"email_id": emailID}).Decode(&response)
	return &response, err
}
