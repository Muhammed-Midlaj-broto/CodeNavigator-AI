package services

import (
	"codebase-analaizer/models"
	"path/filepath"
	"strings"
)

func AnalyzeFiles(files []models.RepoFile) models.RepoStats {

	var stats models.RepoStats
	stats.Languages = make(map[string]int)

	for _, file := range files {

		// Trees API uses "blob" for files
		if file.Type == "blob" {

			stats.Files = append(stats.Files, file.Name)
			stats.TotalFiles++

			ext := strings.ToLower(filepath.Ext(file.Name))

			switch ext {

			case ".go":
				stats.Languages["Go"]++

			case ".js":
				stats.Languages["JavaScript"]++

			case ".html":
				stats.Languages["HTML"]++

			case ".css":
				stats.Languages["CSS"]++

			case ".md":
				stats.Languages["Markdown"]++

			case ".yml", ".yaml":
				stats.Languages["YAML"]++

			case ".json":
				stats.Languages["JSON"]++

			default:
				stats.Languages["Other"]++
			}
		}

		// Trees API uses "tree" for folders
		if file.Type == "tree" {

			stats.Folders = append(stats.Folders, file.Name)
			stats.TotalFolders++

		}
	}

	return stats
}
