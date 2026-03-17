package routes

import (
	"codebase-analaizer/controllers"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine) {

	r.GET("/", controllers.HomePage)

	r.POST("/analyze", controllers.AnalyzeRepo)

	r.GET("/explain-file", controllers.ExplainFile)

	// NEW FEATURES
	r.POST("/explain-line", controllers.ExplainLine)
	r.POST("/ai-search", controllers.AISearch)
}
