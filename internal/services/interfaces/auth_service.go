package interfaces

import (
	"mailmind-api/internal/models"
)

type AuthService interface {
	Logout(userID string) error
	RefreshAccessToken(userID, refreshToken string) (string, error)
	ValidateToken(token string) (string, error)
	GoogleConnect(code string) (*models.User, string, string, string, error)
}
