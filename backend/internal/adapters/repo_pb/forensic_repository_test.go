package repo_pb

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"
)

func TestForensicRepository_CreateAnalysis_OK(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			w.WriteHeader(http.StatusMethodNotAllowed)
			return
		}
		if r.URL.Path != "/api/collections/forensics/records" {
			w.WriteHeader(http.StatusNotFound)
			return
		}

		var payload map[string]interface{}
		if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		if payload["gameId"] != "game-1" {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		w.WriteHeader(http.StatusOK)
		_ = json.NewEncoder(w).Encode(map[string]interface{}{
			"id":         "forensic-1",
			"gameId":     "game-1",
			"clueId":     "clue-1",
			"result":     "match",
			"confidence": 90,
			"status":     "DONE",
		})
	}))
	defer server.Close()

	client, err := NewClient(Config{BaseURL: server.URL, Timeout: 2 * time.Second})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	repo := NewForensicRepository(client)
	result, err := repo.CreateAnalysis(context.Background(), ForensicRecordInput{
		GameID:     "game-1",
		ClueID:     "clue-1",
		Result:     "match",
		Confidence: 90,
		Status:     "DONE",
	})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if result.ID != "forensic-1" {
		t.Fatalf("expected id forensic-1, got %s", result.ID)
	}
}

func TestForensicRepository_ListAnalysesByGame_OK(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			w.WriteHeader(http.StatusMethodNotAllowed)
			return
		}
		if !strings.HasPrefix(r.URL.Path, "/api/collections/forensics/records") {
			w.WriteHeader(http.StatusNotFound)
			return
		}

		query := r.URL.Query().Get("filter")
		if query != "gameId='game-1'" {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		w.WriteHeader(http.StatusOK)
		_ = json.NewEncoder(w).Encode(map[string]interface{}{
			"items": []map[string]interface{}{
				{
					"id":         "forensic-1",
					"gameId":     "game-1",
					"clueId":     "clue-1",
					"result":     "match",
					"confidence": 90,
				},
			},
		})
	}))
	defer server.Close()

	client, err := NewClient(Config{BaseURL: server.URL, Timeout: 2 * time.Second})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	repo := NewForensicRepository(client)
	result, err := repo.ListAnalysesByGame(context.Background(), "game-1")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if len(result) != 1 {
		t.Fatalf("expected 1 analysis, got %d", len(result))
	}
}
