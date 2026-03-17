package services

import (
	"codebase-analaizer/models"
	"strings"
)

func BuildTree(files []models.RepoFile) *models.FileNode {
	root := &models.FileNode{Name: "Project Root", Type: "tree", Children: []*models.FileNode{}}

	for _, file := range files {
		parts := strings.Split(file.Name, "/")
		current := root

		for i, part := range parts {
			isLast := i == len(parts)-1
			child := findChild(current, part)

			if child == nil {
				newType := "tree"
				if isLast && file.Type == "blob" {
					newType = "blob"
				}
				child = &models.FileNode{
					Name:     part,
					Type:     newType,
					Children: []*models.FileNode{},
				}
				current.Children = append(current.Children, child)
			}
			current = child
		}
	}
	return root
}

func findChild(parent *models.FileNode, name string) *models.FileNode {
	for _, child := range parent.Children {
		if child.Name == name {
			return child
		}
	}
	return nil
}
