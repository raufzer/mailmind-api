package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Email struct {
	ID        primitive.ObjectID `bson:"_id,omitempty"`
	UserID    primitive.ObjectID `bson:"user_id"`
	To        string           `bson:"to"`
	CC        []string           `bson:"cc,omitempty"`
	BCC       []string           `bson:"bcc,omitempty"`
	Subject   string             `bson:"subject"`
	Body      string             `bson:"body"`
	SentAt    time.Time          `bson:"sent_at"`
	CreatedAt time.Time          `bson:"created_at"`
}
