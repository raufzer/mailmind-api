package controllers

import (
	"mailmind-api/internal/dto/request"
	"mailmind-api/internal/dto/response"
	"mailmind-api/internal/models"
	"mailmind-api/internal/services/interfaces"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type EmailController struct {
	emailService interfaces.EmailService
}

func NewEmailController(emailService interfaces.EmailService) *EmailController {
	return &EmailController{
		emailService: emailService,
	}
}

// 📧 Send Email Endpoint
func (c *EmailController) SendEmail(ctx *gin.Context) {
	var req request.SendEmailRequest

	// 🔍 Validate Request
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, response.Response{Code: 400, Status: "Bad Request", Message: "Invalid request payload"})
		return
	}

	// Convert request to model
	email := &models.Email{
		UserID:  req.UserID,
		To:      req.To,
		Subject: req.Subject,
		Body:    req.Body,
	}

	// Call the service to send email
	if err := c.emailService.SendEmail(ctx, email); err != nil {
		ctx.JSON(http.StatusInternalServerError, response.Response{Code: 500, Status: "Internal Server Error", Message: "Failed to send email"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Email sent successfully"})
}

// ✍️ Save Draft Endpoint
func (c *EmailController) SaveDraft(ctx *gin.Context) {
	var req request.SaveDraftRequest

	// 🔍 Validate Request
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, response.Response{Code: 400, Status: "Bad Request", Message: "Invalid request payload"})
		return
	}

	// Convert request to model
	draft := &models.Draft{
		UserID:  req.UserID,
		To:      req.To,
		Subject: req.Subject,
		Body:    req.Body,
	}

	// Call the service to save draft
	if err := c.emailService.SaveDraft(ctx, draft); err != nil {
		ctx.JSON(http.StatusInternalServerError, response.Response{Code: 500, Status: "Internal Server Error", Message: "Failed to save draft"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Draft saved successfully"})
}

// 🔍 Get Draft Endpoint
func (c *EmailController) GetDraft(ctx *gin.Context) {
	draftIDParam := ctx.Param("id")
	draftID, err := primitive.ObjectIDFromHex(draftIDParam)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, response.Response{Code: 400, Status: "Bad Request", Message: "Invalid draft ID"})
		return
	}

	draft, err := c.emailService.GetDraft(ctx, draftID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, response.Response{Code: 500, Status: "Internal Server Error", Message: "Failed to get draft"})
		return
	}

	ctx.JSON(http.StatusOK, response.ToDraftResponse(draft))
}
