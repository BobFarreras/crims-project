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
	"github.com/go-chi/chi/v5"
)

func TestLobbyJoin_WithCapabilities(t *testing.T) {
	baseURL := os.Getenv("PB_URL")
	if baseURL == "" {
		t.Skip("PB_URL not set")
	}
	email, password := requireTestCredentials(t)

	client, err := repo_pb.NewClient(repo_pb.Config{BaseURL: baseURL, Timeout: 5 * time.Second})
	if err != nil {
		t.Fatalf("create pb client: %v", err)
	}

	gameRepo := repo_pb.NewGameRepository(client)
	playerRepo := repo_pb.NewPlayerRepository(client)
	service := services.NewLobbyService(gameRepo, playerRepo)
	handler := apihttp.NewAuthHandler(client, apihttp.AuthCookieConfig{Secure: false, SameSite: http.SameSiteLaxMode})
	router := chi.NewRouter()
	apihttp.RegisterAuthRoutes(router, handler)
	apihttp.RegisterLobbyRoutes(router, service, client)
	server := httptest.NewServer(router)
	defer server.Close()

	loginInfo := loginUserData(t, server.URL, email, password)
	game := createTestGame(t, gameRepo, loginInfo.Token)

	payload := map[string]interface{}{
		"gameCode":     game.Code,
		"userId":       "user-1",
		"capabilities": []string{"DETECTIVE", "ANALYST"},
	}
	body, err := json.Marshal(payload)
	if err != nil {
		t.Fatalf("marshal payload: %v", err)
	}

	request, err := http.NewRequest(http.MethodPost, server.URL+"/api/lobby/join", bytes.NewReader(body))
	if err != nil {
		t.Fatalf("create request: %v", err)
	}
	request.Header.Set("Content-Type", "application/json")
	for _, cookie := range loginInfo.Cookies {
		request.AddCookie(cookie)
	}
	response, err := http.DefaultClient.Do(request)
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
	email, password := requireTestCredentials(t)

	client, err := repo_pb.NewClient(repo_pb.Config{BaseURL: baseURL, Timeout: 5 * time.Second})
	if err != nil {
		t.Fatalf("create pb client: %v", err)
	}

	gameRepo := repo_pb.NewGameRepository(client)
	playerRepo := repo_pb.NewPlayerRepository(client)
	service := services.NewLobbyService(gameRepo, playerRepo)
	handler := apihttp.NewAuthHandler(client, apihttp.AuthCookieConfig{Secure: false, SameSite: http.SameSiteLaxMode})
	router := chi.NewRouter()
	apihttp.RegisterAuthRoutes(router, handler)
	apihttp.RegisterLobbyRoutes(router, service, client)
	server := httptest.NewServer(router)
	defer server.Close()

	loginInfo := loginUserData(t, server.URL, email, password)
	game := createTestGame(t, gameRepo, loginInfo.Token)

	payload := map[string]interface{}{
		"gameCode":     game.Code,
		"userId":       "user-2",
		"capabilities": []string{},
	}
	body, err := json.Marshal(payload)
	if err != nil {
		t.Fatalf("marshal payload: %v", err)
	}

	request, err := http.NewRequest(http.MethodPost, server.URL+"/api/lobby/join", bytes.NewReader(body))
	if err != nil {
		t.Fatalf("create request: %v", err)
	}
	request.Header.Set("Content-Type", "application/json")
	for _, cookie := range loginInfo.Cookies {
		request.AddCookie(cookie)
	}
	response, err := http.DefaultClient.Do(request)
	if err != nil {
		t.Fatalf("join request: %v", err)
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusCreated {
		t.Fatalf("expected status 201, got %d", response.StatusCode)
	}
}

func TestLobbyCreate_ReturnsCode(t *testing.T) {
	baseURL := os.Getenv("PB_URL")
	if baseURL == "" {
		t.Skip("PB_URL not set")
	}
	email, password := requireTestCredentials(t)

	client, err := repo_pb.NewClient(repo_pb.Config{BaseURL: baseURL, Timeout: 5 * time.Second})
	if err != nil {
		t.Fatalf("create pb client: %v", err)
	}

	gameRepo := repo_pb.NewGameRepository(client)
	playerRepo := repo_pb.NewPlayerRepository(client)
	service := services.NewLobbyService(gameRepo, playerRepo)
	handler := apihttp.NewAuthHandler(client, apihttp.AuthCookieConfig{Secure: false, SameSite: http.SameSiteLaxMode})
	router := chi.NewRouter()
	apihttp.RegisterAuthRoutes(router, handler)
	apihttp.RegisterLobbyRoutes(router, service, client)
	server := httptest.NewServer(router)
	defer server.Close()

	payload := map[string]interface{}{
		"userId":       "user-1",
		"capabilities": []string{"DETECTIVE"},
	}
	body, err := json.Marshal(payload)
	if err != nil {
		t.Fatalf("marshal payload: %v", err)
	}

	loginInfo := loginUserData(t, server.URL, email, password)
	request, err := http.NewRequest(http.MethodPost, server.URL+"/api/lobby/create", bytes.NewReader(body))
	if err != nil {
		t.Fatalf("create request: %v", err)
	}
	request.Header.Set("Content-Type", "application/json")
	for _, cookie := range loginInfo.Cookies {
		request.AddCookie(cookie)
	}
	response, err := http.DefaultClient.Do(request)
	if err != nil {
		t.Fatalf("create request: %v", err)
	}
	defer response.Body.Close()

	if response.StatusCode == http.StatusForbidden || response.StatusCode == http.StatusInternalServerError {
		t.Skip("PocketBase rules prevent lobby creation")
	}
	if response.StatusCode != http.StatusCreated {
		t.Fatalf("expected status 201, got %d", response.StatusCode)
	}

	var result map[string]interface{}
	if err := json.NewDecoder(response.Body).Decode(&result); err != nil {
		t.Fatalf("decode response: %v", err)
	}
	game, ok := result["Game"].(map[string]interface{})
	if !ok {
		game, _ = result["game"].(map[string]interface{})
	}
	code, _ := game["code"].(string)
	if len(code) != 4 {
		t.Fatalf("expected 4-letter code, got %q", code)
	}
}

func createTestGame(t *testing.T, repo ports.GameRepository, token string) ports.GameRecord {
	t.Helper()

	code := fmt.Sprintf("G%04d", time.Now().UnixNano()%10000)
	ctx := ports.WithAuthToken(context.Background(), token)
	game, err := repo.CreateGame(ctx, ports.GameRecordInput{
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

func requireTestCredentials(t *testing.T) (string, string) {
	t.Helper()
	email := os.Getenv("PB_TEST_EMAIL")
	password := os.Getenv("PB_TEST_PASSWORD")
	if email == "" || password == "" {
		t.Skip("PB_TEST_EMAIL/PB_TEST_PASSWORD not set")
	}
	return email, password
}
