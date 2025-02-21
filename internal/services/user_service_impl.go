package services

import (
	"context"
	"mailmind-api/internal/dto/request"
	"mailmind-api/internal/models"
	"mailmind-api/internal/repositories/interfaces"
	"mailmind-api/pkg/utils"
	"net/http"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserService struct {
	userRepository interfaces.UserRepository
}

func NewUserService(userRepo interfaces.UserRepository) *UserService {
	return &UserService{userRepository: userRepo}
}

func (s *UserService) GetUserByID(ctx context.Context, userID string) (*models.User, error) {
	if userID == "" {
		return nil, utils.NewCustomError(http.StatusBadRequest, "Invalid user ID")
	}
	user, err := s.userRepository.GetUserByID(ctx, userID)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, utils.NewCustomError(http.StatusNotFound, "User not found")
		}
		return nil, utils.NewCustomError(http.StatusInternalServerError, "Error fetching user")
	}
	return user, nil
}

func (s *UserService) GetUserByEmail(ctx context.Context, email string) (*models.User, error) {
	if email == "" {
		return nil, utils.NewCustomError(http.StatusBadRequest, "Invalid email")
	}
	user, err := s.userRepository.GetUserByEmail(ctx, email)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, utils.NewCustomError(http.StatusNotFound, "User not found")
		}
		return nil, utils.NewCustomError(http.StatusInternalServerError, "Error fetching user")
	}
	return user, nil
}

func (s *UserService) UpdateUserSettings(ctx context.Context, userID string, settings request.UpdateUserSettingsRequest) error {
    if userID == "" {
        return utils.NewCustomError(http.StatusBadRequest, "Invalid user ID")
    }
	_, err := primitive.ObjectIDFromHex(userID)
    if err != nil {
        return utils.NewCustomError(http.StatusBadRequest, "Invalid user ID")
    }
    if err := settings.Validate(); err != nil {
        return utils.NewCustomError(http.StatusBadRequest, err.Error())
    }
    userSettings := models.UserSettings{
        PreferredTone: settings.PreferredTone,
        AutoSend:      settings.AutoSend,
    }
	err = s.userRepository.UpdateUserSettings(ctx, userID, &userSettings)
    if err != nil {
        if err == mongo.ErrNoDocuments {
            return utils.NewCustomError(http.StatusNotFound, "User not found")
        }
        return utils.NewCustomError(http.StatusInternalServerError, "Error updating user settings")
    }
    return nil
}