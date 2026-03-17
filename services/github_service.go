package services

import (
	"encoding/base64"
	"encoding/json"
	"net/http"

	"codebase-analaizer/models"
)

func GetFileContent(owner string, repo string, path string) (string, error) {

	url := "https://api.github.com/repos/" + owner + "/" + repo + "/contents/" + path

	resp, err := http.Get(url)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	var data struct {
		Content string `json:"content"`
	}

	err = json.NewDecoder(resp.Body).Decode(&data)
	if err != nil {
		return "", err
	}

	decoded, err := base64.StdEncoding.DecodeString(data.Content)
	if err != nil {
		return "", err
	}

	return string(decoded), nil
}

type RepoInfo struct {
	DefaultBranch string `json:"default_branch"`
}

func GetRepoTree(owner string, repo string) ([]models.RepoFile, error) {

	repoURL := "https://api.github.com/repos/" + owner + "/" + repo

	resp, err := http.Get(repoURL)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var repoInfo RepoInfo
	json.NewDecoder(resp.Body).Decode(&repoInfo)

	treeURL := "https://api.github.com/repos/" + owner + "/" + repo +
		"/git/trees/" + repoInfo.DefaultBranch + "?recursive=1"

	treeResp, err := http.Get(treeURL)
	if err != nil {
		return nil, err
	}
	defer treeResp.Body.Close()

	var tree models.TreeResponse
	json.NewDecoder(treeResp.Body).Decode(&tree)

	var files []models.RepoFile

	for _, item := range tree.Tree {

		files = append(files, models.RepoFile{
			Name: item.Path,
			Type: item.Type,
		})

	}

	return files, nil
}
