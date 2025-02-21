package v1

import (
	"mailmind-api/internal/controllers"

	"github.com/gin-gonic/gin"
)

func UserRoutes(rg *gin.RouterGroup, userController *controllers.UserController) {
	usersRoute := rg.Group("/users")
	usersRoute.GET("/:userId", userController.GetUserByID)
	usersRoute.GET("/email/:email", userController.GetUserByEmail)
}
