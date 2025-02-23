package response

import (
	"mailmind-api/internal/models"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type UserResponse struct {
	ID       primitive.ObjectID   `json:"id"`
	Email    string               `json:"email"`
	Name     string               `json:"name"`
	Settings UserSettingsResponse `json:"settings"`
}

type UserSettingsResponse struct {
	PreferredTone string `json:"preferred_tone"`
	AutoSend      bool   `json:"auto_send"`
}

func ToUserResponse(user *models.User) UserResponse {
	return UserResponse{
		ID:       user.ID,
		Email:    user.Email,
		Name:     user.Name,
		Settings: UserSettingsResponse{
			PreferredTone: user.Settings.PreferredTone,
			AutoSend:      user.Settings.AutoSend,
		},
	}
}
