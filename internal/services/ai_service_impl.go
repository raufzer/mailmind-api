package services

import (
    "context"
     
    "mailmind-api/config"
    "mailmind-api/internal/integrations"
    "mailmind-api/internal/models"
    "mailmind-api/internal/repositories/interfaces"
    "mailmind-api/pkg/utils"
    "net/http"
    "time"
 
    "go.mongodb.org/mongo-driver/bson/primitive"
)

type AIService struct {
    aiResponseRepo interfaces.AIResponseRepository
    config         *config.AppConfig
}

func NewAIService(aiResponseRepo interfaces.AIResponseRepository, config *config.AppConfig) *AIService {
    return &AIService{
        aiResponseRepo: aiResponseRepo,
        config:         config,
    }
}

func (s *AIService) GenerateReply(ctx context.Context, content string) (*models.AIResponse, error) {
    apiKey := s.config.GeminiAPIKey
    if apiKey == "" {
        return nil, utils.NewCustomError(http.StatusInternalServerError, "Gemini API key is missing")
    }

    
    objectID := primitive.NewObjectID()

    reply, err := integrations.CallGeminiAPI(content, apiKey)
    if err != nil {
        return nil, utils.NewCustomError(http.StatusInternalServerError, "Failed to generate reply from AI service")
    }

    if reply == "" {
        return nil, utils.NewCustomError(http.StatusInternalServerError, "Empty response received from AI service")
    }

    aiResponse := &models.AIResponse{
        ID:             objectID,  
        GeneratedReply: reply,
        CreatedAt:      time.Now(),
    }

    if err := s.aiResponseRepo.SaveResponse(ctx, aiResponse); err != nil {
        return nil, utils.NewCustomError(http.StatusInternalServerError, "Failed to save AI response")
    }

    return aiResponse, nil
}