package services

import (
	"context"
	"mailmind-api/config"
	"mailmind-api/internal/integrations"
	"mailmind-api/internal/models"
	"mailmind-api/internal/repositories/interfaces"
	"mailmind-api/pkg/utils"
	"net/http"

	"github.com/go-redis/redis/v8"
	"go.mongodb.org/mongo-driver/mongo"
)

type AuthService struct {
	userRepository  interfaces.UserRepository
	redisRepository interfaces.RedisRepository
	config          *config.AppConfig
}

func NewAuthService(userRepo interfaces.UserRepository, redisRepo interfaces.RedisRepository, config *config.AppConfig) *AuthService {
	return &AuthService{
		userRepository:  userRepo,
		redisRepository: redisRepo,
		config:          config,
	}
}
func (s *AuthService) RefreshAccessToken(userID, refreshToken string) (string, error) {
	storedToken, err := s.redisRepository.GetRefreshToken(userID)
	if err != nil {
		if err == redis.Nil {
			return "", utils.NewCustomError(http.StatusUnauthorized, "Refresh token expired or not found")
		}
		return "", utils.NewCustomError(http.StatusInternalServerError, "Failed to retrieve refresh token")
	}
	if storedToken != refreshToken {
		return "", utils.NewCustomError(http.StatusUnauthorized, "Invalid Token")
	}

	accessToken, err := utils.GenerateToken(userID, s.config.AccessTokenMaxAge, "access", s.config.AccessTokenSecret)
	if err != nil {
		return "", utils.NewCustomError(http.StatusInternalServerError, "Failed to generate access token")
	}

	return accessToken, nil
}

func (s *AuthService) Logout(userID string) error {
	err := s.redisRepository.InvalidateRefreshToken(userID)
	if err != nil {
		return utils.NewCustomError(http.StatusInternalServerError, "Failed to delete refresh token")
	}
	return nil
}

func (s *AuthService) ValidateToken(token string) (string, error) {
	claims, err := utils.ValidateToken(token, s.config.RefreshTokenSecret, "refresh")
	if err != nil {
		return "", utils.NewCustomError(http.StatusUnauthorized, "Invalid or expired token")
	}
	return claims.ID, nil
}

func (s *AuthService) GoogleConnect(code string) (*models.User, string, string, string, error) {
	oauthConfig := integrations.InitializeGoogleOAuthConfig(s.config.GoogleClientID, s.config.GoogleClientSecret, s.config.GoogleRedirectURL)
    ctx := context.Background()
	token, err := oauthConfig.Exchange(ctx, code)
	if err != nil {
		return nil, "", "", "", utils.NewCustomError(http.StatusBadRequest, "Failed to exchange authorization code for token")
	}

	userInfo, err := integrations.FetchGoogleUserInfo(oauthConfig, token)
	if err != nil {
		return nil, "", "", "", utils.NewCustomError(http.StatusInternalServerError, "Failed to fetch user information from Google")
	}

	existingUser, err := s.userRepository.GetUserByEmail(ctx, userInfo.Email)
	if err != nil && err != mongo.ErrNoDocuments {
		return nil, "", "", "", utils.NewCustomError(http.StatusInternalServerError, "Database error")
	}

	if existingUser == nil {
		newUser := &models.User{
			Name:     userInfo.Name,
			Email:    userInfo.Email,
			GoogleID: userInfo.ID,
		}
		_ = s.userRepository.CreateUser(ctx, newUser)
		return newUser, "", "", "register", nil
	}

	return existingUser, "", "", "login", nil
}
