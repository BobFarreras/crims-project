package gemini

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestNewClient_MissingAPIKey(t *testing.T) {
	_, err := NewClient(Config{})
	if err == nil {
		t.Fatal("expected error, got nil")
	}
}

func TestClient_GenerateNarrative_OK(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte(`{"candidates":[{"content":{"parts":[{"text":"Hello"}]}}]}`))
	}))
	defer server.Close()

	client, err := NewClient(Config{
		APIKey:  "test",
		BaseURL: server.URL,
		Timeout: 2 * time.Second,
	})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	text, err := client.GenerateNarrative(context.Background(), "prompt")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if text == "" {
		t.Fatal("expected non-empty text")
	}
}
