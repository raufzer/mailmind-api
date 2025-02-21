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

	RegisterPublicRoutes(basePath, authController, aiController)

	protected := basePath.Group("/")
	protected.Use(middlewares.AuthMiddleware(appConfig))
	// RegisterProtectedRoutes(
	// 	protected,
	// 	userController,
	// )
}

func RegisterPublicRoutes(
	router *gin.RouterGroup,
	authController *controllers.AuthController,
	aiController *controllers.AIController,
) {
	AuthRoutes(router, authController)
	AIRoutes(router, aiController)
}

// func RegisterProtectedRoutes(
// 	router *gin.RouterGroup,
// 	userController *controllers.UserController,
// ) {

// 	userGroup := router.Group("/user")
// 	userGroup.Use(middlewares.RoleMiddleware("user"))
// 	RegisterUserRoutes(userGroup, userController)
// }

// func RegisterUserRoutes(
// 	router *gin.RouterGroup,
// 	userController *controllers.UserController,
// ) {
// 	UserRoutes(router, userController)
// }
