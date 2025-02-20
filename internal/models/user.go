package models

import "time"

type User struct {
	ID           string       `bson:"_id,omitempty" json:"id"`
	GoogleID     string       `bson:"google_id" json:"google_id"`
	Email        string       `bson:"email" json:"email"`
	Name         string       `bson:"name" json:"name"`
	AccessToken  string       `bson:"access_token" json:"-"`
	RefreshToken string       `bson:"refresh_token" json:"-"`
	TokenExpiry  time.Time    `bson:"token_expiry" json:"-"`
	Settings     UserSettings `bson:"settings" json:"settings"`
}

type UserSettings struct {
	PreferredTone string `bson:"preferred_tone" json:"preferred_tone"`
	AutoSend      bool   `bson:"auto_send" json:"auto_send"`
}
