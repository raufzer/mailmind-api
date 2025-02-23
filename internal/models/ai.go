package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type AIResponse struct {
	ID             primitive.ObjectID `bson:"_id,omitempty"`
	GeneratedReply string             `bson:"generated_reply"`
	CreatedAt      time.Time          `bson:"created_at"`
}
