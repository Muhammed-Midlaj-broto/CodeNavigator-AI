package services

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"os"
)

type AIResponse struct {
	Choices []struct {
		Message struct {
			Content string `json:"content"`
		} `json:"message"`
	} `json:"choices"`
}

func ExplainCode(code string) (string, error) {

	apiKey := os.Getenv("OPENROUTER_API_KEY")

	prompt := `
Explain the following code line by line.

Rules:
- Explain EVERY line.
- Do not skip lines.
- Format strictly like this:

Line 1:
Explanation

Line 2:
Explanation

Continue until the end.

Code:
` + code

	requestBody := map[string]interface{}{
		"model": "meta-llama/llama-3-8b-instruct",
		"messages": []map[string]string{
			{
				"role":    "user",
				"content": prompt,
			},
		},
		"max_tokens": 200,
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

	var result AIResponse
	json.Unmarshal(body, &result)

	return result.Choices[0].Message.Content, nil
}
