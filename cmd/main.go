package main

import (
	"os"

	"codebase-analaizer/routes"

	"github.com/gin-gonic/gin"
)

func main() {

	// Set Gin to release mode (better for production)
	gin.SetMode(gin.ReleaseMode)

	// Create router
	r := gin.Default()

	// CORS middleware
	r.Use(func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(200)
			return
		}

		c.Next()
	})

	// Setup all routes (includes "/" from routes.go)
	routes.SetupRoutes(r)

	// Get PORT (important for deployment platforms)
	port := os.Getenv("PORT")
	if port == "" {
		port = "8081"
	}

	// Start server
	err := r.Run(":" + port)
	if err != nil {
		panic(err)
	}
}
