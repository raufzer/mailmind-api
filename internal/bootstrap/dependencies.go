package bootstrap

import (
	"mailmind-api/config"
	"mailmind-api/internal/controllers"
	"mailmind-api/internal/integrations"
	mongodb "mailmind-api/internal/repositories/mongo"
	"mailmind-api/internal/repositories/redis"
	"mailmind-api/internal/services"
	"mailmind-api/pkg/utils"

	"go.uber.org/zap"
)

type AppDependencies struct {
	AuthController *controllers.AuthController
	AIController   *controllers.AIController
	RedisClient    *config.RedisConfig
}

func InitializeDependencies(cfg *config.AppConfig) (*AppDependencies, error) {
	// Initialize MongoDB
	dbConfig := config.ConnectDatabase(cfg)
	db := dbConfig.Client.Database(cfg.DatabaseName)
	// Initialize Redis
	redisConfig := config.ConnectRedis(cfg)

	// Initialize logger
	utils.InitLogger()

	// Initialize Cloudinary
	integrations.InitCloudinary(cfg)

	// Initialize Repositories
	userRepo := mongodb.NewUserRepository(db)
	aiRepo := mongodb.NewAIResponseRepository(db)
	redisRepo := redis.NewRedisRepository(redisConfig.Client)

	// Initialize Services
	authService := services.NewAuthService(userRepo, redisRepo, cfg)
	aiService := services.NewAIService(aiRepo, cfg, &zap.Logger{})

	// Initialize Controllers
	authController := controllers.NewAuthController(authService, cfg)
	aiCotnroller := controllers.NewAIController(aiService)

	// Return dependencies
	return &AppDependencies{
		AuthController: authController,
		AIController:   aiCotnroller,
		RedisClient:    redisConfig,
	}, nil
}
