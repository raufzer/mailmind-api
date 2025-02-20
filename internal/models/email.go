package models

import "time"

type Email struct {
    ID          string    `bson:"_id,omitempty" json:"id"`
    UserID      string    `bson:"user_id" json:"user_id"`
    EmailID     string    `bson:"email_id" json:"email_id"` 
    Sender      string    `bson:"sender" json:"sender"`
    Recipient   string    `bson:"recipient" json:"recipient"`
    Subject     string    `bson:"subject" json:"subject"`
    Body        string    `bson:"body" json:"body"`
    IsRead      bool      `bson:"is_read" json:"is_read"`
    ReceivedAt  time.Time `bson:"received_at" json:"received_at"`
}
