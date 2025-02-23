package utils

import (
	"fmt"
	"mailmind-api/internal/models"
)

// Simulate sending an email via an external service (Gmail API, SMTP, etc.)
func SendEmail(email *models.Email) error {
	fmt.Printf("📧 Sending email to: %v\n", email.To)
	fmt.Printf("📬 Subject: %s\n", email.Subject)
	fmt.Printf("📄 Body: %s\n", email.Body)
	
	// Simulate successful email sending
	return nil
}
