package repo_pb

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/digitaistudios/crims-backend/internal/ports"
)

var (
	ErrMissingPBURL     = errors.New("missing PocketBase base URL")
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
		cfg.Timeout = 10 * time.Second // Augmentem timeout per seguretat
	}

	// OPTIMITZACIÓ DE RENDIMENT
	// Creem un Transport personalitzat basat en el Default
	t := http.DefaultTransport.(*http.Transport).Clone()

	// Default és 100, està bé
	t.MaxIdleConns = 100

	// IMPORTANT: Default és 2! Això és un coll d'ampolla per APIs
	// Ho pugem a 100 per permetre moltes peticions paral·leles a PocketBase
	t.MaxIdleConnsPerHost = 100
	t.MaxConnsPerHost = 100

	return &Client{
		baseURL: cfg.BaseURL,
		httpClient: &http.Client{
			Timeout:   cfg.Timeout,
			Transport: t,
		},
	}, nil
}

func (c *Client) Ping(ctx context.Context) error {
	// Usem el context per permetre cancel·lació (timeout)
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

// Afegeix això al final de client.go (o dins del mètode)

func (c *Client) CreateUser(username, email, password, passwordConfirm, name string) error {
	payload := map[string]string{
		"username":        username,
		"email":           email,
		"password":        password,
		"passwordConfirm": passwordConfirm,
		"name":            name,
	}

	// Nota: PocketBase requereix Content-Type: application/json
	// Implementació simplificada: usa el teu client HTTP per fer POST a /api/collections/users/records
	// (Aquest codi és un exemple, adapta'l si tens helpers com c.postJSON)
	// 🔥 DEBUG: Anem a veure on estem disparant!
	targetURL := c.baseURL + "/api/collections/users/records"
	fmt.Println("📢 INTENTANT CONNECTAR AMB POCKETBASE A:", targetURL)
	body, _ := json.Marshal(payload)
	req, err := http.NewRequest(http.MethodPost, c.baseURL+"/api/collections/users/records", bytes.NewBuffer(body))
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", "application/json")

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode >= 400 {
		// Podries llegir el body per retornar l'error exacte de PB
		return fmt.Errorf("failed to create user, status: %d", resp.StatusCode)
	}

	return nil
}
