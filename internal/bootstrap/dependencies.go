package bootstrap

import (
	"mailmind-api/config"
	"mailmind-api/internal/controllers"
	"mailmind-api/internal/integrations"
	"mailmind-api/internal/repositories/mongo"
	"mailmind-api/internal/repositories/redis"
	"mailmind-api/internal/services"
	"mailmind-api/pkg/utils"
)

type AppDependencies struct {
	UserController           *controllers.UserController
	AuthController           *controllers.AuthController
	RedisClient              *config.RedisConfig
	SystemController         *controllers.SystemController
}

func InitializeDependencies(cfg *config.AppConfig) (*AppDependencies, error) {
	// Initialize Mongo Database
	dbConfig := config.ConnectDB()

	// Initialize Redis
	redisConfig := config.ConnectRedis(cfg)

	// Initialize logger
	utils.InitLogger()

	// Initialize Cloudinary
	integrations.InitCloudinary(cfg)

	// Initialize Repositories
	userRepo := postgresql.NewUserRepository(dbConfig.DB)
	redisRepo := redis.NewRedisRepository(redisConfig.Client)

	// Initialize Services
	authService := services.NewAuthService(
		userRepo,
		redisRepo,
		cfg,
	)
	userService := services.NewUserService(userRepo)

	// Initialize Controllers
	userController := controllers.NewUserController(userService)
	authController := controllers.NewAuthController(authService, cfg)
	systemController := controllers.NewSystemController(cfg, dbConfig, redisConfig)

	// Return dependencies
	return &AppDependencies{
		UserController:           userController,
		AuthController:           authController,
		RedisClient:              redisConfig,
		SystemController:         systemController,
	}, nil
}
