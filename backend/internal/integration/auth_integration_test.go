package integration

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"
	"time"

	apihttp "github.com/digitaistudios/crims-backend/internal/adapters/http"
	"github.com/digitaistudios/crims-backend/internal/adapters/repo_pb"
	"github.com/digitaistudios/crims-backend/internal/platform/config"
	"github.com/go-chi/chi/v5"
)

func TestAuth_LoginFlow(t *testing.T) {
	baseURL := os.Getenv("PB_URL")
	if baseURL == "" {
		t.Skip("PB_URL not set")
	}

	pbClient, err := repo_pb.NewClient(repo_pb.Config{BaseURL: baseURL, Timeout: 5 * time.Second})
	if err != nil {
		t.Fatalf("create pb client: %v", err)
	}

	cfg, err := config.Load()
	if err != nil {
		t.Fatalf("load config: %v", err)
	}

	handler := apihttp.NewAuthHandler(pbClient, apihttp.AuthCookieConfig{
		Secure:   cfg.AuthCookieSecure,
		SameSite: cfg.AuthCookieSameSite,
	})
	router := chi.NewRouter()
	apihttp.RegisterAuthRoutes(router, handler)

	server := httptest.NewServer(router)
	defer server.Close()

	username := fmt.Sprintf("test-user-%d", time.Now().UnixNano())
	email := fmt.Sprintf("%s@example.com", username)
	password := "TestPass!234"

	registerPayload := map[string]string{
		"username":        username,
		"email":           email,
		"password":        password,
		"passwordConfirm": password,
		"name":            "Test User",
	}
	registerBody, err := json.Marshal(registerPayload)
	if err != nil {
		t.Fatalf("marshal register payload: %v", err)
	}

	registerResp, err := http.Post(server.URL+"/api/auth/register", "application/json", bytes.NewReader(registerBody))
	if err != nil {
		t.Fatalf("register request: %v", err)
	}
	defer registerResp.Body.Close()
	if registerResp.StatusCode != http.StatusOK {
		t.Fatalf("register status: %d", registerResp.StatusCode)
	}

	loginPayload := map[string]string{
		"email":    email,
		"password": password,
	}
	loginBody, err := json.Marshal(loginPayload)
	if err != nil {
		t.Fatalf("marshal login payload: %v", err)
	}

	loginResp, err := http.Post(server.URL+"/api/auth/login", "application/json", bytes.NewReader(loginBody))
	if err != nil {
		t.Fatalf("login request: %v", err)
	}
	defer loginResp.Body.Close()
	if loginResp.StatusCode != http.StatusOK {
		t.Fatalf("login status: %d", loginResp.StatusCode)
	}

	cookieHeader := strings.Join(loginResp.Header.Values("Set-Cookie"), "; ")
	if !strings.Contains(cookieHeader, "auth_token=") {
		t.Fatalf("expected auth_token cookie, got %q", cookieHeader)
	}
	if !strings.Contains(strings.ToLower(cookieHeader), "httponly") {
		t.Fatalf("expected HttpOnly cookie, got %q", cookieHeader)
	}

	var response map[string]interface{}
	if err := json.NewDecoder(loginResp.Body).Decode(&response); err != nil {
		t.Fatalf("decode login response: %v", err)
	}
	if _, ok := response["token"]; ok {
		t.Fatalf("login response should not expose token")
	}
	user, ok := response["user"].(map[string]interface{})
	if !ok {
		t.Fatalf("expected user object in response")
	}
	if user["username"] != username {
		t.Fatalf("expected username %s, got %v", username, user["username"])
	}
}

func TestAuth_LoginMissingCredentials(t *testing.T) {
	baseURL := os.Getenv("PB_URL")
	if baseURL == "" {
		t.Skip("PB_URL not set")
	}

	pbClient, err := repo_pb.NewClient(repo_pb.Config{BaseURL: baseURL, Timeout: 5 * time.Second})
	if err != nil {
		t.Fatalf("create pb client: %v", err)
	}

	cfg, err := config.Load()
	if err != nil {
		t.Fatalf("load config: %v", err)
	}

	handler := apihttp.NewAuthHandler(pbClient, apihttp.AuthCookieConfig{
		Secure:   cfg.AuthCookieSecure,
		SameSite: cfg.AuthCookieSameSite,
	})
	router := chi.NewRouter()
	apihttp.RegisterAuthRoutes(router, handler)

	server := httptest.NewServer(router)
	defer server.Close()

	loginPayload := map[string]string{
		"email":    "",
		"password": "secret",
	}
	loginBody, err := json.Marshal(loginPayload)
	if err != nil {
		t.Fatalf("marshal login payload: %v", err)
	}

	loginResp, err := http.Post(server.URL+"/api/auth/login", "application/json", bytes.NewReader(loginBody))
	if err != nil {
		t.Fatalf("login request: %v", err)
	}
	defer loginResp.Body.Close()

	if loginResp.StatusCode != http.StatusBadRequest {
		t.Fatalf("expected status 400, got %d", loginResp.StatusCode)
	}

	var response map[string]interface{}
	if err := json.NewDecoder(loginResp.Body).Decode(&response); err != nil {
		t.Fatalf("decode login response: %v", err)
	}
	if response["code"] != "auth/missing_credentials" {
		t.Fatalf("expected code auth/missing_credentials, got %v", response["code"])
	}
}
