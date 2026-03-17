package main

import (
	"os"

	"codebase-analaizer/routes"

	"github.com/gin-gonic/gin"
)

func main() {

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

	// ✅ ADD THIS (serve React build)
	r.Static("/app", "../frontend/build")

	// ✅ ADD THIS (for React routing)
	r.NoRoute(func(c *gin.Context) {
		c.File("../frontend/build/index.html")
	})

	// Setup API routes
	routes.SetupRoutes(r)

	// Get PORT
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
