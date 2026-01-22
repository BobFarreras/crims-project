package middleware

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/digitaistudios/crims-backend/internal/ports"
)

type fakePocketBaseClient struct{}

func (f fakePocketBaseClient) Ping(_ context.Context) error {
	return nil
}

func (f fakePocketBaseClient) CreateUser(_, _, _, _, _ string) error {
	return nil
}

func (f fakePocketBaseClient) AuthWithPassword(_, _ string) (*ports.AuthResponse, error) {
	return nil, nil
}

func (f fakePocketBaseClient) RefreshAuth(_ string) (*ports.AuthResponse, error) {
	return nil, nil
}

func TestAuthMiddleware_MissingToken(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "/secure", nil)
	response := httptest.NewRecorder()

	handler := AuthMiddleware(fakePocketBaseClient{}, func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	})

	handler.ServeHTTP(response, request)

	if response.Code != http.StatusUnauthorized {
		t.Fatalf("expected 401, got %d", response.Code)
	}
}
