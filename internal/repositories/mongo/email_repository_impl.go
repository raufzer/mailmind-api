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

func (r *EmailRepositoryMongo) SaveEmail(ctx context.Context, email *models.Email) error {
	_, err := r.collection.InsertOne(ctx, email)
	return err
}

func (r *EmailRepositoryMongo) GetEmailByID(ctx context.Context, emailID string) (*models.Email, error) {
	var email models.Email
	err := r.collection.FindOne(ctx, bson.M{"id": emailID}).Decode(&email)
	return &email, err
}
