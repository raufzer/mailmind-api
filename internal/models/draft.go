package models

import "time"

type Draft struct {
    ID        string    `bson:"_id,omitempty" json:"id"`
    UserID    string    `bson:"user_id" json:"user_id"`
    Recipient string    `bson:"recipient" json:"recipient"`
    Subject   string    `bson:"subject" json:"subject"`
    Body      string    `bson:"body" json:"body"`
    CreatedAt time.Time `bson:"created_at" json:"created_at"`
}
