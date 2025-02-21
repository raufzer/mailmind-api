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
	userController *controllers.UserController,
	appConfig *config.AppConfig,
) {

	basePath := router.Group("/v1")

	RegisterPublicRoutes(basePath, authController)

	protected := basePath.Group("/")
	protected.Use(middlewares.AuthMiddleware(appConfig))
	RegisterProtectedRoutes(
		protected,
		userController,
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
	userController *controllers.UserController,
) {

	adminGroup := router.Group("/admin")
	adminGroup.Use(middlewares.RoleMiddleware("admin"))
	RegisterAdminRoutes(adminGroup, userController)
}

func RegisterAdminRoutes(
	router *gin.RouterGroup,
	userController *controllers.UserController,
) {
	UserRoutes(router, userController)
}

func RegisterUserRoutes(
	router *gin.RouterGroup,
) {

}