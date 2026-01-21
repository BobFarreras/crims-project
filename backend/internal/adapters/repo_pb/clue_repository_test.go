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

func TestClueRepository_CreateClue_OK(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			w.WriteHeader(http.StatusMethodNotAllowed)
			return
		}
		if r.URL.Path != "/api/collections/clues/records" {
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
			"id":          "clue-1",
			"gameId":      "game-1",
			"type":        "OBJECT",
			"state":       "DISCOVERED",
			"reliability": 80,
			"facts":       map[string]interface{}{"note": "test"},
		})
	}))
	defer server.Close()

	client, err := NewClient(Config{BaseURL: server.URL, Timeout: 2 * time.Second})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	repo := NewClueRepository(client)
	result, err := repo.CreateClue(context.Background(), ClueRecordInput{
		GameID:      "game-1",
		Type:        "OBJECT",
		State:       "DISCOVERED",
		Reliability: 80,
		Facts:       map[string]interface{}{"note": "test"},
	})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if result.ID != "clue-1" {
		t.Fatalf("expected id clue-1, got %s", result.ID)
	}
}

func TestClueRepository_ListCluesByGame_OK(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			w.WriteHeader(http.StatusMethodNotAllowed)
			return
		}
		if !strings.HasPrefix(r.URL.Path, "/api/collections/clues/records") {
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
					"id":          "clue-1",
					"gameId":      "game-1",
					"type":        "OBJECT",
					"state":       "DISCOVERED",
					"reliability": 80,
				},
			},
		})
	}))
	defer server.Close()

	client, err := NewClient(Config{BaseURL: server.URL, Timeout: 2 * time.Second})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	repo := NewClueRepository(client)
	result, err := repo.ListCluesByGame(context.Background(), "game-1")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if len(result) != 1 {
		t.Fatalf("expected 1 clue, got %d", len(result))
	}
}
