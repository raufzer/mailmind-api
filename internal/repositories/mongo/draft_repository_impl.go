package mongo

import (
	"context"
	"mailmind-api/internal/models"
	"mailmind-api/internal/repositories/interfaces"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type DraftRepositoryMongo struct {
	collection *mongo.Collection
}

func NewDraftRepository(db *mongo.Database) interfaces.DraftRepository {
	return &DraftRepositoryMongo{collection: db.Collection("drafts")}
}

func (r *DraftRepositoryMongo) CreateDraft(ctx context.Context, draft *models.Draft) error {
	_, err := r.collection.InsertOne(ctx, draft)
	return err
}

func (r *DraftRepositoryMongo) GetDraftByID(ctx context.Context, draftID string) (*models.Draft, error) {
	var draft models.Draft
	err := r.collection.FindOne(ctx, bson.M{"_id": draftID}).Decode(&draft)
	if err != nil {
		return nil, err
	}
	return &draft, nil
}
