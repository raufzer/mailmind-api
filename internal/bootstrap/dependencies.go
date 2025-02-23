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
	EmailController *controllers.EmailController
}

func InitializeDependencies(cfg *config.AppConfig) (*AppDependencies, error) {
	// Initialize MongoDB
	dbConfig := config.ConnectDatabase(cfg)
	db := dbConfig.Client.Database(cfg.DatabaseName)


	// Initialize Repositories
	userRepo := mongodb.NewUserRepository(db)
	aiRepo := mongodb.NewAIResponseRepository(db)
	emailRepo := mongodb.NewEmailRepository(db)
	draftRepo := mongodb.NewDraftRepository(db)

	// Initialize Services
	authService := services.NewAuthService(userRepo, cfg)
	aiService := services.NewAIService(aiRepo, cfg)
	emailService := services.NewEmailService(emailRepo, draftRepo)

	// Initialize Controllers
	authController := controllers.NewAuthController(authService, cfg)
	aiCotnroller := controllers.NewAIController(aiService)
	emailcontroller := controllers.NewEmailController(emailService)

	// Return dependencies
	return &AppDependencies{
		AuthController: authController,
		AIController:   aiCotnroller,
		EmailController: emailcontroller,
	}, nil
}
