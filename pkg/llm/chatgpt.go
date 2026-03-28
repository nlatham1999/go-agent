package llm

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
	"time"
)

const defaultChatGPTBaseURL = "https://api.openai.com/v1"
const defaultChatGPTModel = "gpt-4.1-mini"

type ChatGPTClient struct {
	APIKey        string
	Model         string
	BaseURL       string
	PromptID      string
	PromptVersion string
	HTTPClient    *http.Client
}

type responsesRequest struct {
	Model              string          `json:"model,omitempty"`
	Input              string          `json:"input,omitempty"`
	PreviousResponseID string          `json:"previous_response_id,omitempty"`
	Prompt             *responsePrompt `json:"prompt,omitempty"`
}

type responsePrompt struct {
	ID      string `json:"id"`
	Version string `json:"version,omitempty"`
}

type responsesResponse struct {
	ID     string `json:"id"`
	Output []struct {
		Type    string `json:"type"`
		Content []struct {
			Type string `json:"type"`
			Text string `json:"text"`
		} `json:"content"`
	} `json:"output"`
	Error *struct {
		Message string `json:"message"`
	} `json:"error,omitempty"`
}

func NewChatGPTClient(apiKey, model string) *ChatGPTClient {
	if model == "" {
		model = defaultChatGPTModel
	}

	return &ChatGPTClient{
		APIKey:        apiKey,
		Model:         model,
		BaseURL:       defaultChatGPTBaseURL,
		PromptVersion: "1",
		HTTPClient: &http.Client{
			Timeout: 30 * time.Second,
		},
	}
}

func (c *ChatGPTClient) Complete(agent *Agent, prompt string) (string, error) {
	if c == nil {
		return "", fmt.Errorf("chatgpt client is nil")
	}
	if c.APIKey == "" {
		return "", fmt.Errorf("openai api key is empty")
	}
	if agent == nil {
		return "", fmt.Errorf("agent is nil")
	}

	baseURL := c.BaseURL
	if baseURL == "" {
		baseURL = defaultChatGPTBaseURL
	}

	httpClient := c.HTTPClient
	if httpClient == nil {
		httpClient = &http.Client{Timeout: 30 * time.Second}
	}

	requestBody := responsesRequest{
		Model: c.Model,
		Input: prompt,
	}
	if agent.LastResponseID != "" {
		requestBody.PreviousResponseID = agent.LastResponseID
	}
	if c.PromptID != "" {
		requestBody.Prompt = &responsePrompt{
			ID:      c.PromptID,
			Version: c.PromptVersion,
		}
	}

	body, err := json.Marshal(requestBody)
	if err != nil {
		return "", err
	}

	req, err := http.NewRequest(http.MethodPost, baseURL+"/responses", bytes.NewReader(body))
	if err != nil {
		return "", err
	}
	req.Header.Set("Authorization", "Bearer "+c.APIKey)
	req.Header.Set("Content-Type", "application/json")

	resp, err := httpClient.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	var completion responsesResponse
	if err := json.Unmarshal(respBody, &completion); err != nil {
		return "", err
	}

	if resp.StatusCode >= 300 {
		if completion.Error != nil && completion.Error.Message != "" {
			return "", fmt.Errorf("openai api error: %s", completion.Error.Message)
		}
		return "", fmt.Errorf("openai api error: status %d", resp.StatusCode)
	}

	if len(completion.Output) == 0 {
		return "", fmt.Errorf("openai api returned no output")
	}

	agent.LastResponseID = completion.ID

	var outputText strings.Builder
	for _, output := range completion.Output {
		for _, content := range output.Content {
			if content.Text == "" {
				continue
			}
			if outputText.Len() > 0 {
				outputText.WriteString("\n")
			}
			outputText.WriteString(content.Text)
		}
	}

	if outputText.Len() == 0 {
		return "", fmt.Errorf("openai api returned empty text output")
	}

	return strings.TrimSpace(outputText.String()), nil
}
