package controllers

import (
	"context"

	"mailmind-api/internal/dto/request"
	"mailmind-api/internal/dto/response"
	"mailmind-api/internal/services/interfaces"
	"net/http"

	"github.com/gin-gonic/gin"
)

type AIController struct {
	aiService interfaces.AIService
}

func NewAIController(aiService interfaces.AIService) *AIController {
	return &AIController{
		aiService: aiService,
	}
}

func (c *AIController) GenerateReply(ctx *gin.Context) {
	var req request.GenerateReplyRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, response.Response{Code: 400, Status: "Bad Request", Message: "Invalid request payload"})
		return
	}

	reply, err := c.aiService.GenerateReply(context.Background(), req.Content)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, response.Response{Code: 500, Status: "Internal Server Error", Message: "Failed to generate reply"})
		return
	}

	ctx.JSON(http.StatusOK, response.Response{Code: 200, Status: "OK", Message: "Reply generated successfully", Data: response.ToAIResponse(reply)})
}
