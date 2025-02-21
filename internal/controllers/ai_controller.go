package controllers

import (
	"mailmind-api/internal/dto/request"
	"mailmind-api/internal/dto/response"
	serviceInterfaces "mailmind-api/internal/services/interfaces"
	"net/http"

	"github.com/gin-gonic/gin"

)

type AIController struct {
	aiService serviceInterfaces.AIService
}

func NewAIController(aiService serviceInterfaces.AIService) *AIController {
	return &AIController{
		aiService: aiService,
	}
}

func (c *AIController) GenerateReply(ctx *gin.Context) {
	var req request.GenerateReplyRequest

	// Bind JSON request
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, response.Response{
			Code:    http.StatusBadRequest,
			Status:  "Bad Request",
			Message: "Invalid request body",
		})
		return
	}

	userID := ctx.MustGet("user_id").(string)

	// Call service to generate AI reply
	aiResponse, err := c.aiService.GenerateReply(ctx, req, userID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, response.Response{
			Code:    http.StatusInternalServerError,
			Status:  "Internal Server Error",
			Message: "Failed to generate reply"})
		return
	}
     
	ctx.JSON(http.StatusOK, response.Response{
		Code:    http.StatusOK,
		Status:  "OK",
		Message: "Reply generated successfully",
		Data:    response.ToAIResponse(aiResponse),
	})
}
