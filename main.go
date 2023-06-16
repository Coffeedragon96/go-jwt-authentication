package main

import (
	"log"
	"os"

	"github.com/coffeedragon96/go-jwt-authentication/routes"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {

	err := godotenv.Load(".env")

	if err != nil {
		log.Fatal("Error loading .env file")
	}
	// Get port from declared variables in environment file
	port := os.Getenv("PORT")

	// Check if port is empty
	if port == "" {
		port = "8000"
	}

	// Create new router using Gin
	router := gin.New()
	router.Use(gin.Logger())

	// Pass router to the routes
	routes.AuthRoutes(router)
	routes.UserRoutes(router)

	// Define APIs
	router.GET("/api-1", func(c *gin.Context) {
		c.JSON(200, gin.H{"success": "Access granted for api-1"})
	})

	router.GET("/api-2", func(c *gin.Context) {
		c.JSON(200, gin.H{"success": "Access granted for api-2"})
	})

	// Run the router
	router.Run(":" + port)
}
