package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"

)

type Email struct {
	ID         primitive.ObjectID `bson:"_id,omitempty"`
	UserID     primitive.ObjectID `bson:"user_id"`
	EmailID    primitive.ObjectID `bson:"email_id"`
	Sender     string             `bson:"sender"`
	Recipient  string             `bson:"recipient"`
	Subject    string             `bson:"subject"`
	Body       string             `bson:"body"`
	ReceivedAt time.Time          `bson:"received_at"`
}
