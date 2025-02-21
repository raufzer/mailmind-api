package mongo

import (
	"context"
	"mailmind-api/internal/models"
	"mailmind-api/internal/repositories/interfaces"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type AIResponseRepositoryMongo struct {
	collection *mongo.Collection
}

func NewAIResponseRepository(db *mongo.Database) interfaces.AIResponseRepository {
	return &AIResponseRepositoryMongo{collection: db.Collection("ai_responses")}
}

func (r *AIResponseRepositoryMongo) CreateAIResponse(ctx context.Context, response *models.AIResponse) error {
	_, err := r.collection.InsertOne(ctx, response)
	return err
}

func (r *AIResponseRepositoryMongo) GetAIResponseByEmailID(ctx context.Context, emailID string) (*models.AIResponse, error) {
	var response models.AIResponse
	err := r.collection.FindOne(ctx, bson.M{"email_id": emailID}).Decode(&response)
	if err != nil {
		return nil, err
	}
	return &response, nil
}
