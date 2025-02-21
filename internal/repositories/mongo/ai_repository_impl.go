package mongo

import (
	"context"
	"mailmind-api/internal/models"
	"mailmind-api/internal/repositories/interfaces"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type AIResponseRepository struct {
	collection *mongo.Collection
	ctx        context.Context
}

func NewAIResponseRepository(db *mongo.Database) interfaces.AIResponseRepository {
	return &AIResponseRepository{
		collection: db.Collection("ai_responses"),
	}
}

func (r *AIResponseRepository) SaveResponse(ctx context.Context, aiResponse *models.AIResponse) error {
	_, err := r.collection.InsertOne(ctx, aiResponse)
	return err
}

func (r *AIResponseRepository) GetResponseByEmailID(ctx context.Context, emailID primitive.ObjectID) (*models.AIResponse, error) {
	var aiResponse models.AIResponse
	err := r.collection.FindOne(ctx, bson.M{"email_id": emailID}).Decode(&aiResponse)
	if err != nil {
		return nil, err
	}
	return &aiResponse, nil
}
