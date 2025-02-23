package v1

import (
	"mailmind-api/internal/controllers"

	"github.com/gin-gonic/gin"
)

func EmailRoutes(router *gin.RouterGroup, emailController *controllers.EmailController) {
	router.POST("/emails/send", emailController.SendEmail)
	router.POST("/emails/draft", emailController.SaveDraft)
	router.GET("/emails/draft/:id", emailController.GetDraft)
}
