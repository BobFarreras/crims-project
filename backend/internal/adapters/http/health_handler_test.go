package http

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	// 🔥 ASSEGURA'T QUE TENS AQUEST IMPORT!
	"github.com/digitaistudios/crims-backend/internal/platform/web"
	"github.com/digitaistudios/crims-backend/internal/ports"
)

type fakePocketBaseClient struct {
	err error
}

func (f fakePocketBaseClient) Ping(ctx context.Context) error {
	return f.err
}

func (f fakePocketBaseClient) CreateUser(username, email, password, passwordConfirm, name string) error {
	return f.err
}

// 🔥 AFEGEIX AQUEST MÈTODE QUE FALTAVA (I usa ports.AuthResponse)
func (f fakePocketBaseClient) AuthWithPassword(identity, password string) (*ports.AuthResponse, error) {
	return nil, f.err
}

func (f fakePocketBaseClient) RefreshAuth(token string) (*ports.AuthResponse, error) {
	return nil, f.err
}

func (f fakePocketBaseClient) UpdateUserName(_, _, _ string) error {
	return f.err
}

func TestHealthHandler_OK(t *testing.T) {
	handler := NewHealthHandler(fakePocketBaseClient{})

	request := httptest.NewRequest(http.MethodGet, "/api/health", nil)
	response := httptest.NewRecorder()

	handler.ServeHTTP(response, request)

	if response.Code != http.StatusOK {
		t.Fatalf("expected status 200, got %d", response.Code)
	}

	var payload map[string]string
	if err := json.NewDecoder(response.Body).Decode(&payload); err != nil {
		t.Fatalf("failed to decode response: %v", err)
	}

	if payload["status"] != "healthy" {
		t.Fatalf("expected status healthy, got %s", payload["status"])
	}
	if payload["pocketbase"] != "ok" {
		t.Fatalf("expected pocketbase ok, got %s", payload["pocketbase"])
	}
}

func TestHealthHandler_PocketBaseDown(t *testing.T) {
	pingErr := errors.New("connection failed")
	handler := NewHealthHandler(fakePocketBaseClient{err: pingErr})

	request := httptest.NewRequest(http.MethodGet, "/api/health", nil)
	response := httptest.NewRecorder()

	handler.ServeHTTP(response, request)

	if response.Code != http.StatusServiceUnavailable {
		t.Fatalf("expected status 503, got %d", response.Code)
	}

	var payload web.ErrorResponse
	if err := json.NewDecoder(response.Body).Decode(&payload); err != nil {
		t.Fatalf("failed to decode response: %v", err)
	}

	if payload.Code != "pocketbase_unavailable" {
		t.Fatalf("expected error code pocketbase_unavailable, got %s", payload.Code)
	}
}
