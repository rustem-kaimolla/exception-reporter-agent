package llm

import (
	"context"
	"encoding/json"
	"exception-reporter-agent/model"
	"fmt"
	"github.com/go-resty/resty/v2"
)

type Client struct {
	apiKey string
	model  string
}

func NewClient(apiKey, model string) *Client {
	return &Client{apiKey: apiKey, model: model}
}

func (c *Client) Process(ctx context.Context, exc model.ExceptionPayload) (*model.LLMResponse, error) {
	prompt := fmt.Sprintf(`Ты анализируешь исключения из Laravel-приложения. Дай краткое описание (summary), описание бага (description) и приоритет (Low/Medium/High).

Message: %s
File: %s:%d
App: %s
Env: %s
Trace: %v

Ответ в JSON:
{
  "summary": "...",
  "description": "...",
  "priority": "..."
}
`, exc.Message, exc.File, exc.Line, exc.App, exc.Env, exc.Trace)

	req := map[string]interface{}{
		"model": c.model,
		"messages": []map[string]string{
			{"role": "system", "content": "Ты эксперт по анализу исключений."},
			{"role": "user", "content": prompt},
		},
		"temperature": 0.2,
	}

	var result struct {
		Choices []struct {
			Message struct {
				Content string `json:"content"`
			} `json:"message"`
		} `json:"choices"`
	}

	client := resty.New()

	_, err := client.R().
		SetContext(ctx).
		SetHeader("Authorization", "Bearer "+c.apiKey).
		SetHeader("Content-Type", "application/json").
		SetBody(req).
		SetResult(&result).
		Post("https://api.openai.com/v1/chat/completions")

	if err != nil {
		return nil, err
	}

	if len(result.Choices) == 0 || result.Choices[0].Message.Content == "" {
		return nil, fmt.Errorf("LLM returned empty response")
	}

	content := result.Choices[0].Message.Content

	var parsed model.LLMResponse
	err = json.Unmarshal([]byte(content), &parsed)
	if err != nil {
		return nil, fmt.Errorf("llm returned invalid JSON: %w\nraw: %s", err, content)
	}

	return &parsed, nil
}
