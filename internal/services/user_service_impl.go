package services

import (
	"database/sql"
	"mailmind-api/internal/dto/request"
	"mailmind-api/internal/models"
	"mailmind-api/internal/repositories/interfaces"
	"mailmind-api/pkg/utils"
	"net/http"
	"time"

	"github.com/google/uuid"
)

type UserService struct {
	userRepository interfaces.UserRepository
}

func NewUserService(userRepo interfaces.UserRepository) *UserService {
	return &UserService{userRepository: userRepo}
}

func (s *UserService) CreateUser(req request.CreateUsersRequest) (*models.User, error) {
	existingUser, err := s.userRepository.GetUserByEmail(req.Email)
	if err != nil {
		if err == sql.ErrNoRows {
			existingUser = nil
		}
	}
	if existingUser != nil {
		return nil, utils.NewCustomError(http.StatusBadRequest, "User already exists")
	} else {

		hashedPassword, err := utils.HashPassword(req.Password)
		if err != nil {
			return nil, utils.NewCustomError(http.StatusInternalServerError, "Password hashing failed")
		}

		user := &models.User{
			Name:      req.Name,
			Email:     req.Email,
			Password:  hashedPassword,
			Role:      req.Role,
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		}

		if err := s.userRepository.CreateUser(user); err != nil {
			return nil, utils.NewCustomError(http.StatusInternalServerError, "User creation failed")
		}

		return user, nil
	}
}

func (s *UserService) GetUser(userID uuid.UUID) (*models.User, error) {
	if userID == uuid.Nil {
		return nil, utils.NewCustomError(http.StatusBadRequest, "Invalid user ID")
	}
	user, err := s.userRepository.GetUserByID(userID)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, utils.NewCustomError(http.StatusNotFound, "User not found")
		}
		return nil, utils.NewCustomError(http.StatusInternalServerError, "Error fetching user")
	}
	return user, nil
}

func (s *UserService) UpdateUser(userID uuid.UUID, req request.UpdateUserRequest) (*models.User, error) {
	updatedUser := &models.User{
		Name:      req.Name,
		Email:     req.Email,
		Password:  req.Password,
		Role:      req.Role,
		UpdatedAt: time.Now(),
	}

	if err := s.userRepository.UpdateUser(userID, updatedUser); err != nil {
		if err == sql.ErrNoRows {
			return nil, utils.NewCustomError(http.StatusNotFound, "User not found")
		}
		return nil, utils.NewCustomError(http.StatusInternalServerError, "Failed to update user")
	}

	return s.userRepository.GetUserByID(userID)
}

func (s *UserService) GetAllUsers() ([]*models.User, error) {
	users, err := s.userRepository.GetAllUsers()
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, utils.NewCustomError(http.StatusNotFound, "Users not found")
		}
		return nil, utils.NewCustomError(http.StatusInternalServerError, "Failed to fetch users")
	}
	return users, nil
}

func (s *UserService) DeleteUser(userID uuid.UUID) error {
	err := s.userRepository.DeleteUser(userID)
	if err != nil {
		if err == sql.ErrNoRows {
			return utils.NewCustomError(http.StatusNotFound, "User not found")
		}
		return utils.NewCustomError(http.StatusInternalServerError, "Failed to delete user")
	}
	return nil
}
