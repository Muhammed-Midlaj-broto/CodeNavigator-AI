package controllers

import (
	"codebase-analaizer/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

type ExplainRequest struct {
	Code string `json:"code"`
}

func ExplainLine(c *gin.Context) {

	var req ExplainRequest

	// read request body
	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid request body",
		})
		return
	}

	// check empty line
	if req.Code == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "No code provided",
		})
		return
	}

	// send to AI service
	explanation, err := services.ExplainCode(req.Code)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	// return explanation
	c.JSON(http.StatusOK, gin.H{
		"explanation": explanation,
	})
}
