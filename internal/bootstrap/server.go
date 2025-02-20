package bootstrap

import (
	"mailmind-api/config"
	"mailmind-api/internal/middlewares"
	v1 "mailmind-api/internal/routes/api/v1"

	"github.com/gin-gonic/gin"
)

func CreateServer(appConfig *config.AppConfig) *gin.Engine {
	// Set Gin mode
	gin.SetMode(gin.ReleaseMode)
	server := gin.Default()

	// CORS setup
	server.Use(config.SetupCORS(appConfig.FrontEndDomain, appConfig.BackEndDomain))

	// Global middleware
	server.Use(gin.Recovery())
	server.Use(middlewares.ErrorHandlingMiddleware())
	server.Use(middlewares.LoggingMiddleware())
	server.Use(middlewares.RateLimiter(20, 10))
	server.MaxMultipartMemory = 8 << 20 // 8 MiB

	return server
}
