package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID           primitive.ObjectID `bson:"_id,omitempty"`
	GoogleID     string             `bson:"google_id"`
	ProfileImage string             `bson:"profile_image"`
	Email        string             `bson:"email"`
	Name         string             `bson:"name"`
	Settings     UserSettings       `bson:"settings"`
}

type UserSettings struct {
	PreferredTone string `bson:"preferred_tone"`
	AutoSend      bool   `bson:"auto_send"`
}
