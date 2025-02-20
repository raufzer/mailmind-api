package models

import "time"

type AIResponse struct {
    ID         string    `bson:"_id,omitempty" json:"id"`
    UserID     string    `bson:"user_id" json:"user_id"`
    EmailID    string    `bson:"email_id" json:"email_id"`
    GeneratedReply string `bson:"generated_reply" json:"generated_reply"`
    CreatedAt  time.Time `bson:"created_at" json:"created_at"`
}
