package mongo

import (
	"context"
	"mailmind-api/internal/models"
	"mailmind-api/internal/repositories/interfaces"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type DraftRepository struct {
	collection *mongo.Collection
}

func NewDraftRepository(db *mongo.Database) interfaces.DraftRepository {
	return &DraftRepository{
		collection: db.Collection("drafts"),
	}
}

func (r *DraftRepository) SaveDraft(ctx context.Context, draft *models.Draft) error {
	_, err := r.collection.InsertOne(ctx, draft)
	return err
}

func (r *DraftRepository) GetDraftByID(ctx context.Context, draftID primitive.ObjectID) (*models.Draft, error) {
	var draft models.Draft
	err := r.collection.FindOne(ctx, bson.M{"_id": draftID}).Decode(&draft)
	if err != nil {
		return nil, err
	}
	return &draft, nil
}
