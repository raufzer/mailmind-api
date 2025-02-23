package v1

import (
	"mailmind-api/config"
	"mailmind-api/internal/controllers"
	"mailmind-api/internal/middlewares"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(
	router *gin.Engine,
	authController *controllers.AuthController,
	aiController *controllers.AIController,
	appConfig *config.AppConfig,
) {

	basePath := router.Group("/v1")

	RegisterPublicRoutes(basePath, authController)

	protected := basePath.Group("/")
	protected.Use(middlewares.AuthMiddleware(appConfig))
	RegisterProtectedRoutes(
		protected,
		aiController,
	)
}

func RegisterPublicRoutes(
	router *gin.RouterGroup,
	authController *controllers.AuthController,
) {
	AuthRoutes(router, authController)
}

func RegisterProtectedRoutes(
	router *gin.RouterGroup,
	aiController *controllers.AIController,
) {
	AIRoutes(router, aiController)
}
