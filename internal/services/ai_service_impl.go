package services

import (
	"context"
	"errors"
	"fmt"
	"mailmind-api/config"
	"mailmind-api/internal/dto/request"
	"mailmind-api/internal/integrations"
	"mailmind-api/internal/models"
	"mailmind-api/internal/repositories/interfaces"
	"strconv"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.uber.org/zap"
)

type AIService struct {
	aiRepo    interfaces.AIResponseRepository
	config    *config.AppConfig
	logger    *zap.Logger
	semaphore chan struct{} // For rate limiting
}

func NewAIService(aiRepo interfaces.AIResponseRepository, config *config.AppConfig, logger *zap.Logger) *AIService {
	return &AIService{
		aiRepo:    aiRepo,
		config:    config,
		logger:    logger,
		semaphore: make(chan struct{}, 10), // Buffered channel for concurrency control
	}
}

func (s *AIService) GenerateReply(ctx context.Context, req request.GenerateReplyRequest, userID string) (*models.AIResponse, error) {
	// Input validation
	if err := validateGenerateRequest(req, userID); err != nil {
		return nil, fmt.Errorf("invalid request: %w", err)
	}

	// Acquire semaphore for rate limiting
	select {
	case s.semaphore <- struct{}{}:
		defer func() { <-s.semaphore }()
	case <-ctx.Done():
		return nil, ctx.Err()
	}

	// Create context with timeout
	timeout := time.Duration(s.config.AITimeoutSeconds) * time.Second
	ctx, cancel := context.WithTimeout(ctx, timeout)
	defer cancel()

	// Generate reply with retries
	var generatedReply string
	var genErr error
	AIMaxRetriesStr := s.config.AIMaxRetries
	AIMaxRetries, err := strconv.Atoi(AIMaxRetriesStr)
	if err != nil {
		return nil, fmt.Errorf("invalid AI max retries value: %w", err)
	}
	for attempt := 1; attempt <= AIMaxRetries; attempt++ {
		generatedReply, genErr = integrations.CallGeminiAPI(req.Content, s.config.GeminiAPIKey)
		if genErr == nil {
			break
		}

		if errors.Is(genErr, context.DeadlineExceeded) || errors.Is(genErr, context.Canceled) {
			break
		}

		s.logger.Warn("AI API attempt failed",
			zap.Int("attempt", attempt),
			zap.Error(genErr),
		)

		select {
		case <-time.After(exponentialBackoff(attempt)):
		case <-ctx.Done():
			return nil, ctx.Err()
		}
	}

	if genErr != nil {
		s.logger.Error("Failed to generate reply after retries",
			zap.String("userID", userID),
			zap.String("emailID", req.EmailID.Hex()),
			zap.Error(genErr),
		)
		return nil, fmt.Errorf("failed to generate reply: %w", genErr)
	}

	// Validate generated content
	if generatedReply == "" {
		return nil, errors.New("empty response from AI model")
	}

	// Create response object
	objID := primitive.NewObjectID()
	userObjID, _ := primitive.ObjectIDFromHex(userID)
	aiResponse := &models.AIResponse{
		ID:             objID,
		UserID:         userObjID,
		EmailID:        req.EmailID,
		GeneratedReply: generatedReply,
		CreatedAt:      time.Now(),
	}

	// Save with retry
	if err := s.saveWithRetry(ctx, aiResponse); err != nil {
		s.logger.Error("Failed to save AI response",
			zap.String("responseID", objID.Hex()),
			zap.Error(err),
		)
		return nil, fmt.Errorf("failed to persist response: %w", err)
	}

	return aiResponse, nil
}

func validateGenerateRequest(req request.GenerateReplyRequest, userID string) error {
	if _, err := primitive.ObjectIDFromHex(userID); err != nil {
		return fmt.Errorf("invalid user ID: %w", err)
	}
	if req.EmailID.IsZero() {
		return errors.New("empty email ID")
	}
	if req.Content == "" {
		return errors.New("empty content")
	}
	return nil
}

func exponentialBackoff(attempt int) time.Duration {
	return time.Duration(1<<uint(attempt)) * time.Second
}

func (s *AIService) saveWithRetry(ctx context.Context, response *models.AIResponse) error {
	const maxRetries = 3
	var lastErr error

	for i := 0; i < maxRetries; i++ {
		err := s.aiRepo.SaveResponse(ctx, response)
		if err == nil {
			return nil
		}

		if errors.Is(err, context.Canceled) || errors.Is(err, context.DeadlineExceeded) {
			return err
		}

		lastErr = err
		time.Sleep(time.Duration(i+1) * 100 * time.Millisecond)
	}

	return fmt.Errorf("failed after %d retries: %w", maxRetries, lastErr)
}
