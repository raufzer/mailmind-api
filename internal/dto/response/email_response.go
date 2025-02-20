package response

import (
	"mailmind-api/internal/models"
	"time"
)

type EmailResponse struct {
	ID         string    `json:"id"`
	EmailID    string    `json:"email_id"`
	Sender     string    `json:"sender"`
	Recipient  string    `json:"recipient"`
	Subject    string    `json:"subject"`
	Body       string    `json:"body"`
	IsRead     bool      `json:"is_read"`
	ReceivedAt time.Time `json:"received_at"`
}

func ToEmailResponse(email *models.Email) EmailResponse {
	return EmailResponse{
		ID:         email.ID,
		EmailID:    email.EmailID,
		Sender:     email.Sender,
		Recipient:  email.Recipient,
		Subject:    email.Subject,
		Body:       email.Body,
		IsRead:     email.IsRead,
		ReceivedAt: email.ReceivedAt,
	}
}
