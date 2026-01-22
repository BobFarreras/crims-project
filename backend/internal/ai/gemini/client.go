package gemini

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"time"
)

var ErrMissingAPIKey = errors.New("missing Gemini API key")

type Config struct {
	APIKey  string
	BaseURL string
	Timeout time.Duration
}

type Client struct {
	apiKey     string
	baseURL    string
	httpClient *http.Client
}

func NewClient(cfg Config) (*Client, error) {
	if cfg.APIKey == "" {
		return nil, ErrMissingAPIKey
	}
	if cfg.BaseURL == "" {
		cfg.BaseURL = "https://generativelanguage.googleapis.com/v1beta"
	}
	if cfg.Timeout <= 0 {
		cfg.Timeout = 10 * time.Second
	}

	return &Client{
		apiKey:  cfg.APIKey,
		baseURL: cfg.BaseURL,
		httpClient: &http.Client{
			Timeout: cfg.Timeout,
		},
	}, nil
}

func (c *Client) GenerateNarrative(ctx context.Context, prompt string) (string, error) {
	requestBody := map[string]interface{}{
		"contents": []map[string]interface{}{
			{
				"parts": []map[string]string{{"text": prompt}},
			},
		},
	}
	data, err := json.Marshal(requestBody)
	if err != nil {
		return "", err
	}

	url := fmt.Sprintf("%s/models/gemini-1.5-flash:generateContent?key=%s", c.baseURL, c.apiKey)
	request, err := http.NewRequestWithContext(ctx, http.MethodPost, url, bytes.NewReader(data))
	if err != nil {
		return "", err
	}
	request.Header.Set("Content-Type", "application/json")

	response, err := c.httpClient.Do(request)
	if err != nil {
		return "", err
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		return "", fmt.Errorf("unexpected status %d", response.StatusCode)
	}

	var payload struct {
		Candidates []struct {
			Content struct {
				Parts []struct {
					Text string `json:"text"`
				} `json:"parts"`
			} `json:"content"`
		} `json:"candidates"`
	}
	if err := json.NewDecoder(response.Body).Decode(&payload); err != nil {
		return "", err
	}
	if len(payload.Candidates) == 0 || len(payload.Candidates[0].Content.Parts) == 0 {
		return "", errors.New("empty response")
	}

	return payload.Candidates[0].Content.Parts[0].Text, nil
}
