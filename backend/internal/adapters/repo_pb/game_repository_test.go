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

func TestGameRepository_CreateGame_OK(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			w.WriteHeader(http.StatusMethodNotAllowed)
			return
		}
		if r.URL.Path != "/api/collections/games/records" {
			w.WriteHeader(http.StatusNotFound)
			return
		}

		var payload map[string]string
		if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		if payload["code"] != "ABCD" {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		if payload["state"] != "INVESTIGATION" {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		if payload["seed"] != "seed-1" {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		w.WriteHeader(http.StatusOK)
		_ = json.NewEncoder(w).Encode(map[string]string{
			"id":    "game-1",
			"code":  "ABCD",
			"state": "INVESTIGATION",
			"seed":  "seed-1",
		})
	}))
	defer server.Close()

	client, err := NewClient(Config{BaseURL: server.URL, Timeout: 2 * time.Second})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	repo := NewGameRepository(client)
	result, err := repo.CreateGame(context.Background(), GameRecordInput{
		Code:  "ABCD",
		State: "INVESTIGATION",
		Seed:  "seed-1",
	})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if result.ID != "game-1" {
		t.Fatalf("expected id game-1, got %s", result.ID)
	}
}

func TestGameRepository_CreateGame_HTTPError(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusInternalServerError)
	}))
	defer server.Close()

	client, err := NewClient(Config{BaseURL: server.URL, Timeout: 2 * time.Second})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	repo := NewGameRepository(client)
	_, err = repo.CreateGame(context.Background(), GameRecordInput{Code: "ABCD"})
	if err == nil {
		t.Fatal("expected error, got nil")
	}
}

func TestGameRepository_GetGameByID_OK(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			w.WriteHeader(http.StatusMethodNotAllowed)
			return
		}
		if r.URL.Path != "/api/collections/games/records/game-1" {
			w.WriteHeader(http.StatusNotFound)
			return
		}

		w.WriteHeader(http.StatusOK)
		_ = json.NewEncoder(w).Encode(map[string]string{
			"id":    "game-1",
			"code":  "ABCD",
			"state": "INVESTIGATION",
			"seed":  "seed-1",
		})
	}))
	defer server.Close()

	client, err := NewClient(Config{BaseURL: server.URL, Timeout: 2 * time.Second})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	repo := NewGameRepository(client)
	result, err := repo.GetGameByID(context.Background(), "game-1")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if result.Code != "ABCD" {
		t.Fatalf("expected code ABCD, got %s", result.Code)
	}
}

func TestGameRepository_GetGameByCode_OK(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			w.WriteHeader(http.StatusMethodNotAllowed)
			return
		}
		if !strings.HasPrefix(r.URL.Path, "/api/collections/games/records") {
			w.WriteHeader(http.StatusNotFound)
			return
		}

		query := r.URL.Query().Get("filter")
		if query != "code='ABCD'" {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		w.WriteHeader(http.StatusOK)
		_ = json.NewEncoder(w).Encode(map[string]interface{}{
			"items": []map[string]string{
				{
					"id":    "game-1",
					"code":  "ABCD",
					"state": "INVESTIGATION",
					"seed":  "seed-1",
				},
			},
		})
	}))
	defer server.Close()

	client, err := NewClient(Config{BaseURL: server.URL, Timeout: 2 * time.Second})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	repo := NewGameRepository(client)
	result, err := repo.GetGameByCode(context.Background(), "ABCD")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if result.ID != "game-1" {
		t.Fatalf("expected id game-1, got %s", result.ID)
	}
}
