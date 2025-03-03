package services

import (
	"context"
	"mailmind-api/config"
	"mailmind-api/internal/integrations"
	"mailmind-api/internal/models"
	"mailmind-api/internal/repositories/interfaces"
	"mailmind-api/pkg/utils"
	"net/http"

	"go.mongodb.org/mongo-driver/mongo"
)

type AuthService struct {
	userRepository interfaces.UserRepository
	config         *config.AppConfig
}

func NewAuthService(userRepo interfaces.UserRepository, config *config.AppConfig) *AuthService {
	return &AuthService{
		userRepository: userRepo,
		config:         config,
	}
}

func (s *AuthService) ValidateToken(ctx context.Context, token string) (string, error) {
	claims, err := utils.ValidateToken(token, s.config.AccessTokenSecret)
	if err != nil {
		return "",utils.NewCustomError(http.StatusUnauthorized, "Invalid or expired token")
	}
	return claims.ID, nil
}
func (s *AuthService) GoogleConnect(ctx context.Context, code string) (*models.User, string, string, error) {
	oauthConfig := integrations.InitializeGoogleOAuthConfig(s.config.GoogleClientID, s.config.GoogleClientSecret, s.config.GoogleRedirectURL)
	token, err := oauthConfig.Exchange(ctx, code)
	if err != nil {
		return nil, "", "", utils.NewCustomError(http.StatusBadRequest, "Failed to exchange authorization code for token")
	}

	userInfo, err := integrations.FetchGoogleUserInfo(oauthConfig, token)
	if err != nil {
		return nil, "", "", utils.NewCustomError(http.StatusInternalServerError, "Failed to fetch user information from Google")
	}

	existingUser, err := s.userRepository.GetUserByEmail(ctx, userInfo.Email)
	if err != nil && err != mongo.ErrNoDocuments {
		return nil, "", "", utils.NewCustomError(http.StatusInternalServerError, "Database error")
	}

	if existingUser == nil {
		newUser := &models.User{
			Name:         userInfo.Name,
			Email:        userInfo.Email,
			GoogleID:     userInfo.ID,
			ProfileImage: userInfo.ImageURL,
		}
		err = s.userRepository.CreateUser(ctx, newUser)
		if err != nil {
			return nil, "", "", utils.NewCustomError(http.StatusInternalServerError, "Failed to create user")
		}

		accessToken, err := utils.GenerateToken(newUser.ID.Hex(), s.config.AccessTokenMaxAge, s.config.AccessTokenSecret)
		if err != nil {
			return nil, "", "", utils.NewCustomError(http.StatusInternalServerError, "Failed to generate access token")
		}

		return newUser, accessToken, "register", nil
	}

	accessToken, err := utils.GenerateToken(existingUser.ID.Hex(), s.config.AccessTokenMaxAge, s.config.AccessTokenSecret)
	if err != nil {
		return nil, "", "", utils.NewCustomError(http.StatusInternalServerError, "Failed to generate access token")
	}

	return existingUser, accessToken, "login", nil
}
