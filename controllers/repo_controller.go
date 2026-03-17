package controllers

import (
	"codebase-analaizer/services"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func HomePage(c *gin.Context) {

	// optional if you still serve HTML
	c.HTML(http.StatusOK, "index.html", nil)

}

func AnalyzeRepo(c *gin.Context) {

	repoURL := c.PostForm("repo_url")

	if repoURL == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Repository URL required",
		})
		return
	}

	// remove github prefix
	trimmed := strings.TrimPrefix(repoURL, "https://github.com/")
	trimmed = strings.TrimSuffix(trimmed, "/")

	parts := strings.Split(trimmed, "/")

	if len(parts) < 2 {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid GitHub URL",
		})
		return
	}

	owner := parts[0]
	repo := parts[1]

	// get repo file tree
	files, err := services.GetRepoTree(owner, repo)

	if err != nil {

		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})

		return
	}

	// analyze stats
	stats := services.AnalyzeFiles(files)

	// build folder tree
	tree := services.BuildTree(files)

	// return JSON for React
	c.JSON(http.StatusOK, gin.H{

		"RepoName": repo,
		"Owner":    owner,

		"TotalFiles":   stats.TotalFiles,
		"TotalFolders": stats.TotalFolders,

		"Languages": stats.Languages,

		"Tree": tree,
	})

}

func ExplainFile(c *gin.Context) {

	owner := c.Query("owner")
	repo := c.Query("repo")
	path := c.Query("path")

	if owner == "" || repo == "" || path == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "missing parameters",
		})
		return
	}

	code, err := services.GetFileContent(owner, repo, path)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	// Only return code (no AI call here)
	c.JSON(http.StatusOK, gin.H{
		"code": code,
	})
}
