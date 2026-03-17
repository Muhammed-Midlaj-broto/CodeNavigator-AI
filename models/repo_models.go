package models

type RepoFile struct {
	Name string `json:"name"`
	Type string `json:"type"`
}

type RepoStats struct {
	Files        []string       `json:"files"`   // Add this
	Folders      []string       `json:"folders"` // Add this
	TotalFiles   int            `json:"total_files"`
	TotalFolders int            `json:"total_folders"`
	Languages    map[string]int `json:"languages"`
	RepoName     string         `json:"repo_name"`
	Owner        string         `json:"owner"`
	Tree         []*FileNode    `json:"tree"`
}

type FileNode struct {
	Name     string      `json:"name"`
	Type     string      `json:"type"` // "blob" (file) or "tree" (folder)
	Children []*FileNode `json:"children"`
}

type TreeResponse struct {
	Tree []struct {
		Path string `json:"path"`
		Type string `json:"type"`
	} `json:"tree"`
}
