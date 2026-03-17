package main

import (
	"os"

	"codebase-analaizer/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	// database.ConnectDB() // Removed DB temporarily

	r := gin.Default()

	// CORS
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

	// ROUTES
	routes.SetupRoutes(r)

	// Get PORT from environment
	port := os.Getenv("PORT")
	if port == "" {
		port = "8081"
	}

	r.Run(":" + port)
}
