package request

import "go.mongodb.org/mongo-driver/bson/primitive"

type GenerateReplyRequest struct {
	EmailID primitive.ObjectID `json:"email_id" binding:"required"`
	Content string             `json:"content" binding:"required"`
}

type GeminiAPIRequest struct {
	Prompt string `json:"prompt"`
}

