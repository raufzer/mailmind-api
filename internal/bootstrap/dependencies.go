package bootstrap

import (
	"mailmind-api/config"
	"mailmind-api/internal/controllers"
	"mailmind-api/internal/integrations"
	mongodb "mailmind-api/internal/repositories/mongo"
	"mailmind-api/internal/repositories/redis"
	"mailmind-api/internal/services"
	"mailmind-api/pkg/utils"
)

type AppDependencies struct {
	AuthController *controllers.AuthController

	RedisClient *config.RedisConfig
}

func InitializeDependencies(cfg *config.AppConfig) (*AppDependencies, error) {
	// Initialize MongoDB
	dbConfig := config.ConnectDatabase(cfg)

	// Initialize Redis
	redisConfig := config.ConnectRedis(cfg)

	// Initialize logger
	utils.InitLogger()

	// Initialize Cloudinary
	integrations.InitCloudinary(cfg)

	// Initialize Repositories
	userRepo := mongodb.NewUserRepository(dbConfig.Client.Database(cfg.DatabaseName))
	redisRepo := redis.NewRedisRepository(redisConfig.Client)

	// Initialize Services
	authService := services.NewAuthService(userRepo, redisRepo, cfg)
	// userService := services.NewUserService(userRepo)

	// Initialize Controllers
	// userController := controllers.NewUserController(userService)
	authController := controllers.NewAuthController(authService, cfg)

	// Return dependencies
	return &AppDependencies{
		// UserController: userController,
		AuthController: authController,
		RedisClient:    redisConfig,
	}, nil
}
