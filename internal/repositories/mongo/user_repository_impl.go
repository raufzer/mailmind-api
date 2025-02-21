package mongo

import (
	"context"
	"mailmind-api/internal/models"
	"mailmind-api/internal/repositories/interfaces"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserRepositoryMongo struct {
	collection *mongo.Collection
	ctx        context.Context
}

func NewUserRepository(db *mongo.Database) interfaces.UserRepository {
	return &UserRepositoryMongo{collection: db.Collection("users")}
}

func (r *UserRepositoryMongo) GetUserByID(ctx context.Context, userID string) (*models.User, error) {
	objID, err := primitive.ObjectIDFromHex(userID)
	if err != nil {
		return nil, err
	}
	var user models.User
	err = r.collection.FindOne(ctx, bson.M{"_id": objID}).Decode(&user)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *UserRepositoryMongo) GetUserByEmail(ctx context.Context, email string) (*models.User, error) {
	var user models.User
	err := r.collection.FindOne(ctx, bson.M{"email": email}).Decode(&user)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *UserRepositoryMongo) CreateUser(ctx context.Context, user *models.User) error {
	_, err := r.collection.InsertOne(ctx, user)
	return err
}

func (r *UserRepositoryMongo) UpdateUserSettings(ctx context.Context, userID string, settings *models.UserSettings) error {
    objID, err := primitive.ObjectIDFromHex(userID)
    if err != nil {
        return err
    }
    _, err = r.collection.UpdateOne(ctx, bson.M{"_id": objID}, bson.M{"$set": bson.M{"settings": settings}})
    return err
}