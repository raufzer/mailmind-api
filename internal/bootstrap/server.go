package bootstrap

import (
	"mailmind-api/config"
	"mailmind-api/internal/middlewares"

	"github.com/gin-gonic/gin"
)

func CreateServer(appConfig *config.AppConfig) *gin.Engine {
	// Set Gin mode
	gin.SetMode(gin.ReleaseMode)
	server := gin.Default()

	// CORS setup
	server.Use(config.SetupCORS(appConfig.BackEndDomain))

	// Global middleware
	server.Use(gin.Recovery())
	server.Use(middlewares.ErrorHandlingMiddleware())
	server.Use(middlewares.RateLimiter(20, 10))
	server.MaxMultipartMemory = 8 << 20 // 8 MiB

	return server
}
