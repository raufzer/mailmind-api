package mongo

import (
	"context"
	"mailmind-api/internal/models"
	"mailmind-api/internal/repositories/interfaces"

	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserRepositoryMongo struct {
	collection *mongo.Collection
}

func NewUserRepository(db *mongo.Database) interfaces.UserRepository {
	return &UserRepositoryMongo{collection: db.Collection("users")}
}

func (r *UserRepositoryMongo) CreateUser(ctx context.Context, user *models.User) error {
	_, err := r.collection.InsertOne(ctx, user)
	return err
}

func (r *UserRepositoryMongo) GetUserByEmail(ctx context.Context, email string) (*models.User, error) {
	var user models.User
	err := r.collection.FindOne(ctx, bson.M{"email": email}).Decode(&user)
	return &user, err
}

func (r *UserRepositoryMongo) GetUserByID(ctx context.Context, userID uuid.UUID) (*models.User, error) {
	var user models.User
	err := r.collection.FindOne(ctx, bson.M{"id": userID}).Decode(&user)
	return &user, err
}

func (r *UserRepositoryMongo) UpdateUserSettings(ctx context.Context, userID uuid.UUID, settings models.UserSettings) error {
	_, err := r.collection.UpdateOne(ctx, bson.M{"id": userID}, bson.M{"$set": bson.M{"settings": settings}})
	return err
}

func (r *UserRepositoryMongo) DeleteUser(ctx context.Context, userID uuid.UUID) error {
	_, err := r.collection.DeleteOne(ctx, bson.M{"id": userID})
	return err
}
