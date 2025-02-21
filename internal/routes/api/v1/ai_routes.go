package v1

import (
	"mailmind-api/internal/controllers"

	"github.com/gin-gonic/gin"
)

func AIRoutes(rg *gin.RouterGroup, aiController *controllers.AIController) {
	aiRoute := rg.Group("/ai")
	aiRoute.POST("/generate-reply", aiController.GenerateReply)
}
