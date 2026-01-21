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

func TestPersonRepository_CreatePerson_OK(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			w.WriteHeader(http.StatusMethodNotAllowed)
			return
		}
		if r.URL.Path != "/api/collections/persons/records" {
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
			"id":            "person-1",
			"gameId":        "game-1",
			"name":          "Witness",
			"officialStory": "story",
			"truthStory":    "truth",
			"stress":        10,
			"credibility":   80,
		})
	}))
	defer server.Close()

	client, err := NewClient(Config{BaseURL: server.URL, Timeout: 2 * time.Second})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	repo := NewPersonRepository(client)
	result, err := repo.CreatePerson(context.Background(), PersonRecordInput{
		GameID:        "game-1",
		Name:          "Witness",
		OfficialStory: "story",
		TruthStory:    "truth",
		Stress:        10,
		Credibility:   80,
	})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if result.ID != "person-1" {
		t.Fatalf("expected id person-1, got %s", result.ID)
	}
}

func TestPersonRepository_ListPersonsByGame_OK(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			w.WriteHeader(http.StatusMethodNotAllowed)
			return
		}
		if !strings.HasPrefix(r.URL.Path, "/api/collections/persons/records") {
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
					"id":            "person-1",
					"gameId":        "game-1",
					"name":          "Witness",
					"officialStory": "story",
					"truthStory":    "truth",
					"stress":        10,
					"credibility":   80,
				},
			},
		})
	}))
	defer server.Close()

	client, err := NewClient(Config{BaseURL: server.URL, Timeout: 2 * time.Second})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	repo := NewPersonRepository(client)
	result, err := repo.ListPersonsByGame(context.Background(), "game-1")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if len(result) != 1 {
		t.Fatalf("expected 1 person, got %d", len(result))
	}
}
