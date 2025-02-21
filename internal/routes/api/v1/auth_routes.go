package v1

import (
	"mailmind-api/internal/controllers"

	"github.com/gin-gonic/gin"
)

func AuthRoutes(rg *gin.RouterGroup, authController *controllers.AuthController) {
	authRoute := rg.Group("/auth")
	authRoute.GET("/connect", authController.GoogleConnect)
	authRoute.GET("/callback", authController.GoogleCallbackConnect)
	authRoute.POST("/logout", authController.Logout)
	authRoute.POST("/refresh-token", authController.RefreshToken)

}
