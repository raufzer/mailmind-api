package bootstrap

import (
	"mailmind-api/config"
	"mailmind-api/internal/controllers"
	mongodb "mailmind-api/internal/repositories/mongo"
	"mailmind-api/internal/services"
)

type AppDependencies struct {
	AuthController *controllers.AuthController
	AIController   *controllers.AIController
}

func InitializeDependencies(cfg *config.AppConfig) (*AppDependencies, error) {
	// Initialize MongoDB
	dbConfig := config.ConnectDatabase(cfg)
	db := dbConfig.Client.Database(cfg.DatabaseName)


	// Initialize Repositories
	userRepo := mongodb.NewUserRepository(db)
	aiRepo := mongodb.NewAIResponseRepository(db)

	// Initialize Services
	authService := services.NewAuthService(userRepo, cfg)
	aiService := services.NewAIService(aiRepo, cfg)

	// Initialize Controllers
	authController := controllers.NewAuthController(authService, cfg)
	aiCotnroller := controllers.NewAIController(aiService)

	// Return dependencies
	return &AppDependencies{
		AuthController: authController,
		AIController:   aiCotnroller,
	}, nil
}
