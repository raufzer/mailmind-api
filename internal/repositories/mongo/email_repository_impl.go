package mongo

import (
	"context"
	"mailmind-api/internal/models"
	"mailmind-api/internal/repositories/interfaces"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type EmailRepository struct {
	collection *mongo.Collection
}

func NewEmailRepository(db *mongo.Database) interfaces.EmailRepository {
	return &EmailRepository{
		collection: db.Collection("emails"),
	}
}

func (r *EmailRepository) SaveEmail(ctx context.Context, email *models.Email) error {
	_, err := r.collection.InsertOne(ctx, email)
	return err
}

func (r *EmailRepository) GetEmailByID(ctx context.Context, emailID primitive.ObjectID) (*models.Email, error) {
	var email models.Email
	err := r.collection.FindOne(ctx, bson.M{"_id": emailID}).Decode(&email)
	if err != nil {
		return nil, err
	}
	return &email, nil
}
