package controllers

import (
	"mailmind-api/config"
	"mailmind-api/internal/dto/request"
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

func (c *AuthController) Login(ctx *gin.Context) {
	var req request.LoginRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		_ = ctx.Error(err)
		ctx.Abort()
		return
	}
	user, accessToken, refreshToken, err := c.authService.Login(req)
	if err != nil {
		_ = ctx.Error(err)
		ctx.Abort()
		return
	}
	isProduction := c.config.ServerPort != "9090"
	utils.SetAuthCookie(ctx, "access_token", accessToken, c.config.AccessTokenMaxAge, c.config.BackEndDomain, isProduction)
	utils.SetAuthCookie(ctx, "refresh_token", refreshToken, c.config.RefreshTokenMaxAge, c.config.BackEndDomain, isProduction)
	ctx.JSON(http.StatusOK, response.Response{
		Code:    http.StatusOK,
		Status:  "OK",
		Message: "Successfully logged in!",
		Data:    response.ToUserResponse(user),
	})
}

func (c *AuthController) RefreshToken(ctx *gin.Context) {
	refreshToken, err := ctx.Cookie("refresh_token")
	if err != nil {
		_ = ctx.Error(err)
		ctx.Abort()
		return
	}
	userID, userRole, err := c.authService.ValidateToken(refreshToken)
	if err != nil {
		_ = ctx.Error(err)
		ctx.Abort()
		return
	}

	accessToken, err := c.authService.RefreshAccessToken(userID, userRole, refreshToken)
	if err != nil {
		_ = ctx.Error(err)
		ctx.Abort()
		return
	}

	isProduction := c.config.ServerPort != "9090"
	utils.SetAuthCookie(ctx, "access_token", accessToken, c.config.AccessTokenMaxAge, c.config.BackEndDomain, isProduction)
	ctx.JSON(http.StatusOK, response.Response{
		Code:    http.StatusOK,
		Status:  "OK",
		Message: "Access token refreshed successfully!",
	})
}

func (c *AuthController) Logout(ctx *gin.Context) {
	refreshToken, err := ctx.Cookie("refresh_token")
	if err != nil {
		_ = ctx.Error(err)
		ctx.Abort()
		return
	}
	userID, _, _ := c.authService.ValidateToken(refreshToken)
	if err := c.authService.Logout(userID, refreshToken); err != nil {
		_ = ctx.Error(err)
		ctx.Abort()
		return
	}
	isProduction := c.config.ServerPort != "9090"
	utils.SetAuthCookie(ctx, "access_token", "", -1, c.config.BackEndDomain, isProduction)
	utils.SetAuthCookie(ctx, "refresh_token", "", -1, c.config.BackEndDomain, isProduction)

	ctx.JSON(http.StatusOK, response.Response{
		Code:    http.StatusOK,
		Status:  "OK",
		Message: "Successfully logged out!",
	})
}

func (c *AuthController) Register(ctx *gin.Context) {
	var req request.CreateUsersRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		_ = ctx.Error(err)
		ctx.Abort()
		return
	}
	user, err := c.authService.Register(req)
	if err != nil {
		_ = ctx.Error(err)
		ctx.Abort()
		return
	}
	ctx.JSON(http.StatusCreated, response.Response{
		Code:    http.StatusCreated,
		Status:  "Created",
		Message: "User created successfully",
		Data:    response.ToUserResponse(user),
	})
}

func (c *AuthController) SendResetOTP(ctx *gin.Context) {
	var req request.SendOTPRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		_ = ctx.Error(err)
		ctx.Abort()
		return
	}
	err := c.authService.SendOTP(req.Email)
	if err != nil {
		_ = ctx.Error(err)
		ctx.Abort()
		return
	}
	ctx.JSON(http.StatusOK, response.Response{
		Code:    http.StatusOK,
		Status:  "OK",
		Message: "OTP sent successfully!",
	})
}

func (c *AuthController) VerifyOTP(ctx *gin.Context) {
	var req request.VerifyOTPRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		_ = ctx.Error(err)
		ctx.Abort()
		return
	}
	resetToken, err := c.authService.VerifyOTP(req.Email, req.OTP)
	if err != nil {
		_ = ctx.Error(err)
		return
	}
	isProduction := c.config.ServerPort != "9090"
	utils.SetAuthCookie(ctx, "reset_token", resetToken, c.config.ResetPasswordTokenMaxAge, c.config.BackEndDomain, isProduction)
	ctx.JSON(http.StatusOK, response.Response{
		Code:    http.StatusOK,
		Status:  "OK",
		Message: "OTP verify successfully!",
	})
}

func (c *AuthController) ResetPassword(ctx *gin.Context) {
	token, err := ctx.Cookie("reset_token")
	if err != nil {
		_ = ctx.Error(err)
		ctx.Abort()
		return
	}
	var req request.ResetPasswordRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		_ = ctx.Error(err)
		ctx.Abort()
		return
	}
	err = c.authService.ResetPassword(req.Email, token, req.NewPassword)
	if err != nil {
		_ = ctx.Error(err)
		ctx.Abort()
		return
	}
	ctx.JSON(http.StatusOK, response.Response{
		Code:    http.StatusOK,
		Status:  "OK",
		Message: "Password reset successfully!",
	})
}

func (c *AuthController) GoogleConnect(ctx *gin.Context) {
	role := ctx.Query("role")
	ctx.SetCookie("role", role, 3600, "/", c.config.BackEndDomain, false, true)
	oauthConfig := integrations.InitializeGoogleOAuthConfig(c.config.GoogleClientID, c.config.GoogleClientSecret, c.config.GoogleRedirectURL)

	authURL := oauthConfig.AuthCodeURL("", oauth2.AccessTypeOffline)

	ctx.Redirect(http.StatusFound, authURL)
}

func (c *AuthController) GoogleCallbackConnect(ctx *gin.Context) {
	role, err := ctx.Cookie("role")
	if err != nil || role == "" {
		ctx.JSON(http.StatusBadRequest, response.Response{
			Code:    http.StatusBadRequest,
			Status:  "Bad Request",
			Message: "Role is missing or expired",
		})
		return
	}

	code := ctx.DefaultQuery("code", "")

	if code == "" {
		ctx.JSON(http.StatusBadRequest, response.Response{
			Code:    http.StatusBadRequest,
			Status:  "Bad Request",
			Message: "Code is required",
		})
		return
	}

	user, accessToken, refreshToken, connect, err := c.authService.GoogleConnect(code, role)
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
		utils.SetAuthCookie(ctx, "access_token", accessToken, c.config.AccessTokenMaxAge, c.config.BackEndDomain, isProduction)
		utils.SetAuthCookie(ctx, "refresh_token", refreshToken, c.config.RefreshTokenMaxAge, c.config.BackEndDomain, isProduction)
		ctx.JSON(http.StatusOK, response.Response{
			Code:    http.StatusOK,
			Status:  "OK",
			Message: "Successfully logged in!",
			Data:    response.ToUserResponse(user),
		})

	}
}
