package controllers

import (
	"codebase-analaizer/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

type SearchRequest struct {
	Query string `json:"query"`
	Code  string `json:"code"`
}

func AISearch(c *gin.Context) {

	var req SearchRequest

	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid request",
		})
		return
	}

	result, err := services.SearchCode(req.Query, req.Code)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"result": result,
	})
}
