package request

import "go.mongodb.org/mongo-driver/bson/primitive"

type SendEmailRequest struct {
	UserID  primitive.ObjectID `json:"user_id" binding:"required"`
	To      string             `json:"to" binding:"required,email"`
	Subject string             `json:"subject" binding:"required"`
	Body    string             `json:"body" binding:"required"`
}

type SaveDraftRequest struct {
	UserID  primitive.ObjectID `json:"user_id" binding:"required"`
	To      string             `json:"to" binding:"required,email"`
	Subject string             `json:"subject" binding:"required"`
	Body    string             `json:"body" binding:"required"`
}
