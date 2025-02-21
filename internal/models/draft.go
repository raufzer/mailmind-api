package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Draft struct {
	ID        primitive.ObjectID `bson:"_id,omitempty"`
	UserID    string             `bson:"user_id"`
	Recipient string             `bson:"recipient"`
	Subject   string             `bson:"subject"`
	Body      string             `bson:"body"`
	CreatedAt time.Time          `bson:"created_at"`
}
