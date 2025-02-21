package controllers

import (
	"mailmind-api/config"
	"mailmind-api/internal/dto/response"
	"mailmind-api/internal/integrations"
	serviceInterfaces "mailmind-api/internal/services/interfaces"
	"mailmind-api/pkg/utils"
	"net/http"

	"github.com/gin-gonic/gin"
	"golang.org/x/oauth2"
)

type AuthController struct {
	authService serviceInterfaces.AuthService
	config      *config.AppConfig
}

func NewAuthController(service serviceInterfaces.AuthService, config *config.AppConfig) *AuthController {
	return &AuthController{
		authService: service,
		config:      config,
	}
}

func (c *AuthController) Logout(ctx *gin.Context) {
	acesss_token, err := ctx.Cookie("acesss_token")
	if err != nil {
		_ = ctx.Error(err)
		ctx.Abort()
		return
	}
	userID, err := c.authService.ValidateToken(ctx, acesss_token)
	if err != nil {
		_ = ctx.Error(err)
		ctx.Abort()
		return
	}
	if err := c.authService.Logout(ctx, userID); err != nil {
		_ = ctx.Error(err)
		ctx.Abort()
		return
	}
	isProduction := c.config.ServerPort != "9090"
	utils.SetAuthCookie(ctx, "access_token", "", -1, c.config.BackEndDomain, isProduction)

	ctx.JSON(http.StatusOK, response.Response{
		Code:    http.StatusOK,
		Status:  "OK",
		Message: "Successfully logged out!",
	})
}

func (c *AuthController) GoogleConnect(ctx *gin.Context) {
	oauthConfig := integrations.InitializeGoogleOAuthConfig(c.config.GoogleClientID, c.config.GoogleClientSecret, c.config.GoogleRedirectURL)

	authURL := oauthConfig.AuthCodeURL("", oauth2.AccessTypeOffline)

	ctx.Redirect(http.StatusFound, authURL)
}

func (c *AuthController) GoogleCallbackConnect(ctx *gin.Context) {
	code := ctx.DefaultQuery("code", "")

	if code == "" {
		ctx.JSON(http.StatusBadRequest, response.Response{
			Code:    http.StatusBadRequest,
			Status:  "Bad Request",
			Message: "Code is required",
		})
		return
	}

	user, accessToken, connect, err := c.authService.GoogleConnect(ctx, code)
	if err != nil {
		_ = ctx.Error(err)
		ctx.Abort()
		return
	}
	if connect == "register" {
		ctx.JSON(http.StatusOK, response.Response{
			Code:    http.StatusOK,
			Status:  "OK",
			Message: "User successfully created!",
			Data:    response.ToUserResponse(user),
		})
	} else if connect == "login" {
		isProduction := c.config.ServerPort != "9090"
		ctx.SetCookie("acess_token", accessToken, 3600, "/", c.config.BackEndDomain, false, true)
		utils.SetAuthCookie(ctx, "access_token", accessToken, c.config.AccessTokenMaxAge, c.config.BackEndDomain, isProduction)
		ctx.JSON(http.StatusOK, response.Response{
			Code:    http.StatusOK,
			Status:  "OK",
			Message: "Successfully logged in!",
			Data:    response.ToUserResponse(user),
		})

	}
}
