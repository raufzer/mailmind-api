package mongo

import (
	"context"
	"mailmind-api/internal/models"
	"mailmind-api/internal/repositories/interfaces"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type EmailRepositoryMongo struct {
	collection *mongo.Collection
}

func NewEmailRepository(db *mongo.Database) interfaces.EmailRepository {
	return &EmailRepositoryMongo{collection: db.Collection("emails")}
}

func (r *EmailRepositoryMongo) CreateEmail(ctx context.Context, email *models.Email) error {
	_, err := r.collection.InsertOne(ctx, email)
	return err
}

func (r *EmailRepositoryMongo) GetEmailByID(ctx context.Context, emailID string) (*models.Email, error) {
	var email models.Email
	err := r.collection.FindOne(ctx, bson.M{"_id": emailID}).Decode(&email)
	if err != nil {
		return nil, err
	}
	return &email, nil
}
