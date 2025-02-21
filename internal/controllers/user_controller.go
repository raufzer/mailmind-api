package controllers

import (
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

func (c *UserController) GetUserByID(ctx *gin.Context) {
	userID, err := uuid.Parse(ctx.Param("userId"))
	if err != nil {
		_ = ctx.Error(err)
		ctx.Abort()
		return
	}
	user, err := c.userService.GetUserByID(ctx, userID.String())
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

func (c *UserController) GetUserByEmail(ctx *gin.Context) {
	email := ctx.Param("email")
	user, err := c.userService.GetUserByID(ctx, email)
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
