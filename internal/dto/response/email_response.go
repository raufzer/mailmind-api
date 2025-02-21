package response

import (
	"mailmind-api/internal/models"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type EmailResponse struct {
	ID         primitive.ObjectID `json:"id"`
	EmailID    primitive.ObjectID `json:"email_id"`
	Sender     string             `json:"sender"`
	Recipient  string             `json:"recipient"`
	Subject    string             `json:"subject"`
	Body       string             `json:"body"`
	ReceivedAt time.Time          `json:"received_at"`
}

func ToEmailResponse(email *models.Email) EmailResponse {
	return EmailResponse{
		ID:         email.ID,
		EmailID:    email.EmailID,
		Sender:     email.Sender,
		Recipient:  email.Recipient,
		Subject:    email.Subject,
		Body:       email.Body,
		ReceivedAt: email.ReceivedAt,
	}
}
