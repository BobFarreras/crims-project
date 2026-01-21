package http

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestMetricsHandler_OK(t *testing.T) {
	handler := NewMetricsHandler()

	request := httptest.NewRequest(http.MethodGet, "/api/v1/metrics", nil)
	response := httptest.NewRecorder()

	handler.ServeHTTP(response, request)

	if response.Code != http.StatusOK {
		t.Fatalf("expected 200, got %d", response.Code)
	}
}
