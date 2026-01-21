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

func TestTimelineRepository_CreateEntry_OK(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			w.WriteHeader(http.StatusMethodNotAllowed)
			return
		}
		if r.URL.Path != "/api/collections/timeline/records" {
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
			"id":          "entry-1",
			"gameId":      "game-1",
			"timestamp":   "2026-01-21T10:00:00Z",
			"title":       "Arrival",
			"description": "Victim arrives",
			"eventId":     "event-1",
		})
	}))
	defer server.Close()

	client, err := NewClient(Config{BaseURL: server.URL, Timeout: 2 * time.Second})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	repo := NewTimelineRepository(client)
	result, err := repo.CreateEntry(context.Background(), TimelineRecordInput{
		GameID:      "game-1",
		Timestamp:   "2026-01-21T10:00:00Z",
		Title:       "Arrival",
		Description: "Victim arrives",
		EventID:     "event-1",
	})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if result.ID != "entry-1" {
		t.Fatalf("expected id entry-1, got %s", result.ID)
	}
}

func TestTimelineRepository_ListEntriesByGame_OK(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			w.WriteHeader(http.StatusMethodNotAllowed)
			return
		}
		if !strings.HasPrefix(r.URL.Path, "/api/collections/timeline/records") {
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
					"id":        "entry-1",
					"gameId":    "game-1",
					"timestamp": "2026-01-21T10:00:00Z",
					"title":     "Arrival",
				},
			},
		})
	}))
	defer server.Close()

	client, err := NewClient(Config{BaseURL: server.URL, Timeout: 2 * time.Second})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	repo := NewTimelineRepository(client)
	result, err := repo.ListEntriesByGame(context.Background(), "game-1")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if len(result) != 1 {
		t.Fatalf("expected 1 entry, got %d", len(result))
	}
}
