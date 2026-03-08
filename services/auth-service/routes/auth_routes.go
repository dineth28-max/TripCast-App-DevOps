package routes

import (
	"auth-service/conteollers"

	"github.com/gin-gonic/gin"
)

// SetupAuthRoutes initializes the routes for the auth service
func SetupAuthRoutes(router *gin.Engine, authController *conteollers.AuthController) {
	// Root route to prevent 404 on base URL
	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Welcome to the TripCast Auth Service API",
			"status":  "UP",
		})
	})

	// Health check
	router.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"status": "UP",
		})
	})

	// Auth routes (Prefix removed because API Gateway will handle /api)
	authRoutes := router.Group("/auth")
	{
		authRoutes.POST("/register", authController.Register)
		authRoutes.POST("/login", authController.Login)
		authRoutes.GET("/validate", authController.Validate)
	}
}
