package services

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"os"
)

type SearchAIResponse struct {
	Choices []struct {
		Message struct {
			Content string `json:"content"`
		} `json:"message"`
	} `json:"choices"`
}

func SearchCode(query string, code string) (string, error) {

	apiKey := os.Getenv("OPENROUTER_API_KEY")

	prompt := `
You are a code search assistant.

User question:
` + query + `

Repository code:
` + code + `

Find the file or lines that answer the question.
Explain shortly and mention the file and line numbers.
`

	requestBody := map[string]interface{}{
		"model": "meta-llama/llama-3-8b-instruct",
		"messages": []map[string]string{
			{
				"role":    "user",
				"content": prompt,
			},
		},
	}

	jsonBody, _ := json.Marshal(requestBody)

	req, _ := http.NewRequest(
		"POST",
		"https://openrouter.ai/api/v1/chat/completions",
		bytes.NewBuffer(jsonBody),
	)

	req.Header.Set("Authorization", "Bearer "+apiKey)
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}

	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}

	defer resp.Body.Close()

	body, _ := io.ReadAll(resp.Body)

	var result SearchAIResponse
	json.Unmarshal(body, &result)

	return result.Choices[0].Message.Content, nil
}
