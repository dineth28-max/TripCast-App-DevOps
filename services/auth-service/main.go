package main

import (
	"log"
	"os"

	"auth-service/conteollers"
	"auth-service/repository"
	"auth-service/routes"
	"auth-service/services"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	// Load .env if it exists
	_ = godotenv.Load()

	// Initialize repository
	userRepo := repository.NewUserRepo()

	// Initialize services
	authService := services.NewAuthService(userRepo)

	// Initialize controllers
	authController := conteollers.NewAuthController(authService)

	// Setup Gin router
	router := gin.Default()

	// Setup routes
	routes.SetupAuthRoutes(router, authController)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8888"
	}

	log.Printf("Auth Service starting on port %s", port)
	if err := router.Run(":" + port); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
