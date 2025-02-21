package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID           primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	GoogleID     string             `bson:"google_id" json:"google_id"`
	ProfileImage string             `bson:"profile_image" json:"profile_image"`
	Email        string             `bson:"email" json:"email"`
	Name         string             `bson:"name" json:"name"`
	TokenExpiry  time.Time          `bson:"token_expiry" json:"-"`
	Settings     UserSettings       `bson:"settings" json:"settings"`
}

type UserSettings struct {
	PreferredTone string `bson:"preferred_tone" json:"preferred_tone"`
	AutoSend      bool   `bson:"auto_send" json:"auto_send"`
}
