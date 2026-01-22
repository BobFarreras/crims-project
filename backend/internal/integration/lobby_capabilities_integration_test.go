package integration

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"
	"time"

	apihttp "github.com/digitaistudios/crims-backend/internal/adapters/http"
	"github.com/digitaistudios/crims-backend/internal/adapters/repo_pb"
	"github.com/digitaistudios/crims-backend/internal/ports"
	"github.com/digitaistudios/crims-backend/internal/services"
)

func TestLobbyJoin_WithCapabilities(t *testing.T) {
	baseURL := os.Getenv("PB_URL")
	if baseURL == "" {
		t.Skip("PB_URL not set")
	}

	client, err := repo_pb.NewClient(repo_pb.Config{BaseURL: baseURL, Timeout: 5 * time.Second})
	if err != nil {
		t.Fatalf("create pb client: %v", err)
	}

	gameRepo := repo_pb.NewGameRepository(client)
	playerRepo := repo_pb.NewPlayerRepository(client)
	service := services.NewLobbyService(gameRepo, playerRepo)
	server := httptest.NewServer(apihttp.NewLobbyJoinHandler(service))
	defer server.Close()

	game := createTestGame(t, gameRepo)

	payload := map[string]interface{}{
		"gameCode":     game.Code,
		"userId":       "user-1",
		"capabilities": []string{"DETECTIVE", "ANALYST"},
	}
	body, err := json.Marshal(payload)
	if err != nil {
		t.Fatalf("marshal payload: %v", err)
	}

	response, err := http.Post(server.URL, "application/json", bytes.NewReader(body))
	if err != nil {
		t.Fatalf("join request: %v", err)
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusCreated {
		t.Fatalf("expected status 201, got %d", response.StatusCode)
	}

	var result map[string]interface{}
	if err := json.NewDecoder(response.Body).Decode(&result); err != nil {
		t.Fatalf("decode response: %v", err)
	}
	if _, ok := result["capabilities"]; !ok {
		t.Fatalf("expected capabilities in response")
	}
}

func TestLobbyJoin_AssistantNoCapabilities(t *testing.T) {
	baseURL := os.Getenv("PB_URL")
	if baseURL == "" {
		t.Skip("PB_URL not set")
	}

	client, err := repo_pb.NewClient(repo_pb.Config{BaseURL: baseURL, Timeout: 5 * time.Second})
	if err != nil {
		t.Fatalf("create pb client: %v", err)
	}

	gameRepo := repo_pb.NewGameRepository(client)
	playerRepo := repo_pb.NewPlayerRepository(client)
	service := services.NewLobbyService(gameRepo, playerRepo)
	server := httptest.NewServer(apihttp.NewLobbyJoinHandler(service))
	defer server.Close()

	game := createTestGame(t, gameRepo)

	payload := map[string]interface{}{
		"gameCode":     game.Code,
		"userId":       "user-2",
		"capabilities": []string{},
	}
	body, err := json.Marshal(payload)
	if err != nil {
		t.Fatalf("marshal payload: %v", err)
	}

	response, err := http.Post(server.URL, "application/json", bytes.NewReader(body))
	if err != nil {
		t.Fatalf("join request: %v", err)
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusCreated {
		t.Fatalf("expected status 201, got %d", response.StatusCode)
	}
}

func createTestGame(t *testing.T, repo ports.GameRepository) ports.GameRecord {
	t.Helper()

	code := fmt.Sprintf("G%04d", time.Now().UnixNano()%10000)
	game, err := repo.CreateGame(context.Background(), ports.GameRecordInput{
		Code:  code,
		State: "LOBBY",
		Seed:  "seed",
	})
	if err != nil {
		if errors.Is(err, repo_pb.ErrUnexpectedStatus) && strings.Contains(err.Error(), "status 403") {
			t.Skip("PocketBase denies game creation (403)")
		}
		t.Fatalf("create game: %v", err)
	}

	return game
}
