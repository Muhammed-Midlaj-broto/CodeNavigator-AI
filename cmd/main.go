package main

import (
	"os"

	"codebase-analaizer/routes"

	"github.com/gin-gonic/gin"
)

func main() {

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

	// ✅ ADD THIS
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "CodeNavigator API is running 🚀",
		})
	})

	// ROUTES
	routes.SetupRoutes(r)

	// PORT
	port := os.Getenv("PORT")
	if port == "" {
		port = "8081"
	}

	r.Run(":" + port)
}
