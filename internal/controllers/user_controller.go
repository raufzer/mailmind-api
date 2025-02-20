package controllers

import (
	"mailmind-api/internal/dto/request"
	"mailmind-api/internal/dto/response"
	serviceInterfaces "mailmind-api/internal/services/interfaces"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type UserController struct {
	userService serviceInterfaces.UserService
}

func NewUserController(service serviceInterfaces.UserService) *UserController {
	return &UserController{
		userService: service,
	}
}

func (c *UserController) CreateUser(ctx *gin.Context) {
	var req request.CreateUsersRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		_ = ctx.Error(err)
		ctx.Abort()
		return
	}
	user, err := c.userService.CreateUser(req)
	if err != nil {
		_ = ctx.Error(err)
		return
	}
	ctx.JSON(http.StatusCreated, response.Response{
		Code:    http.StatusCreated,
		Status:  "Created",
		Message: "User created successfully",
		Data:    response.ToUserResponse(user),
	})
}

func (c *UserController) GetUser(ctx *gin.Context) {
	userID, err := uuid.Parse(ctx.Param("userId"))
	if err != nil {
		_ = ctx.Error(err)
		ctx.Abort()
		return
	}
	user, err := c.userService.GetUser(userID)
	if err != nil {
		_ = ctx.Error(err)
		return
	}
	ctx.JSON(http.StatusOK, response.Response{
		Code:    http.StatusOK,
		Status:  "OK",
		Message: "User found",
		Data:    response.ToUserResponse(user),
	})
}

func (c *UserController) UpdateUser(ctx *gin.Context) {
	var req request.UpdateUserRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		_ = ctx.Error(err)
		ctx.Abort()
		return
	}
	id, err := uuid.Parse(ctx.Param("userId"))
	if err != nil {
		_ = ctx.Error(err)
		return
	}
	updatedUser, err := c.userService.UpdateUser(id, req)
	if err != nil {
		_ = ctx.Error(err)
		ctx.Abort()
		return
	}
	ctx.JSON(http.StatusOK, response.Response{
		Code:    http.StatusOK,
		Status:  "OK",
		Message: "User updated successfully",
		Data:    response.ToUserResponse(updatedUser),
	})
}

func (c *UserController) GetAllUsers(ctx *gin.Context) {
	users, err := c.userService.GetAllUsers()
	if err != nil {
		_ = ctx.Error(err)
		return
	}
	ctx.JSON(http.StatusOK, response.Response{
		Code:    http.StatusOK,
		Status:  "OK",
		Message: "Users retrieved successfully",
		Data:    response.ToUsersResponse(users),
	})
}

func (c *UserController) DeleteUser(ctx *gin.Context) {
	userID, err := uuid.Parse(ctx.Param("userId"))
	if err != nil {
		_ = ctx.Error(err)
		ctx.Abort()
		return
	}
	err = c.userService.DeleteUser(userID)
	if err != nil {
		_ = ctx.Error(err)
		return
	}
	ctx.JSON(http.StatusOK, response.Response{
		Code:    http.StatusOK,
		Status:  "OK",
		Message: "User deleted successfully",
	})
}
