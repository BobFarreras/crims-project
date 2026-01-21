package repo_pb

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/digitaistudios/crims-backend/internal/ports"
)

var (
	ErrMissingPBURL    = errors.New("missing PocketBase base URL")
	ErrUnexpectedStatus = errors.New("unexpected PocketBase status")
)

type Config struct {
	BaseURL string
	Timeout time.Duration
}

type Client struct {
	baseURL    string
	httpClient *http.Client
}

var _ ports.PocketBaseClient = (*Client)(nil)

func NewClient(cfg Config) (*Client, error) {
	if cfg.BaseURL == "" {
		return nil, ErrMissingPBURL
	}
	if cfg.Timeout <= 0 {
		cfg.Timeout = 5 * time.Second
	}

	return &Client{
		baseURL: cfg.BaseURL,
		httpClient: &http.Client{
			Timeout: cfg.Timeout,
		},
	}, nil
}

func (c *Client) Ping(ctx context.Context) error {
	request, err := http.NewRequestWithContext(ctx, http.MethodGet, c.baseURL+"/api/health", nil)
	if err != nil {
		return err
	}

	response, err := c.httpClient.Do(request)
	if err != nil {
		return err
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		return fmt.Errorf("%w: status %d", ErrUnexpectedStatus, response.StatusCode)
	}

	return nil
}
