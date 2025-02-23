package services

import (
	"context"
	"mailmind-api/internal/models"
	"mailmind-api/internal/repositories/interfaces"
	"mailmind-api/pkg/utils"
	"net/http"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type EmailService struct {
	emailRepo interfaces.EmailRepository
	draftRepo interfaces.DraftRepository
}

func NewEmailService(emailRepo interfaces.EmailRepository, draftRepo interfaces.DraftRepository) *EmailService {
	return &EmailService{
		emailRepo: emailRepo,
		draftRepo: draftRepo,
	}
}

// 📧 Send an Email
func (s *EmailService) SendEmail(ctx context.Context, email *models.Email) error {
	// Simulate sending the email via an external email service
	err := utils.SendEmail(email)
	if err != nil {
		return utils.NewCustomError(1001, err.Error()) // Assuming 1001 is the error code for "failed to send email"
	}

	// Save the sent email
	email.ID = primitive.NewObjectID()
	email.SentAt = time.Now()
	err = s.emailRepo.SaveEmail(ctx, email)
	if err != nil {
		return utils.NewCustomError(http.StatusInternalServerError, "failed to save sent email")
	}

	return nil
}

// ✍️ Save a New Draft
func (s *EmailService) SaveDraft(ctx context.Context, draft *models.Draft) error {
	draft.ID = primitive.NewObjectID()
	draft.CreatedAt = time.Now()
	draft.UpdatedAt = time.Now()

	err := s.draftRepo.SaveDraft(ctx, draft)
	if err != nil {
		return utils.NewCustomError(http.StatusInternalServerError, "failed to save draft")
	}
	return nil
}

// 🔍 Get a Draft by ID
func (s *EmailService) GetDraft(ctx context.Context, draftID primitive.ObjectID) (*models.Draft, error) {
	draft, err := s.draftRepo.GetDraftByID(ctx, draftID)
	if err != nil {
		return nil, utils.NewCustomError(http.StatusNotFound, "draft not found")
	}
	return draft, nil
}
