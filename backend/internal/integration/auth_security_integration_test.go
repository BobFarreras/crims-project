package integration

import (
	"bytes"
	"encoding/json"
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

func TestAuth_LoginCookieFlags_Production(t *testing.T) {
	baseURL := os.Getenv("PB_URL")
	if baseURL == "" {
		t.Skip("PB_URL not set")
	}

	t.Setenv("ENVIRONMENT", "production")
	t.Setenv("AUTH_COOKIE_SAMESITE", "none")

	cfg, err := config.Load()
	if err != nil {
		t.Fatalf("load config: %v", err)
	}

	pbClient, err := repo_pb.NewClient(repo_pb.Config{BaseURL: baseURL, Timeout: 5 * time.Second})
	if err != nil {
		t.Fatalf("create pb client: %v", err)
	}

	handler := apihttp.NewAuthHandler(pbClient, apihttp.AuthCookieConfig{
		Secure:   cfg.AuthCookieSecure,
		SameSite: cfg.AuthCookieSameSite,
	})

	router := chi.NewRouter()
	apihttp.RegisterAuthRoutes(router, handler)

	server := httptest.NewServer(router)
	defer server.Close()

	username, email, password := registerUser(t, server.URL)
	loginData := loginUserData(t, server.URL, email, password)
	defer cleanupTestUser(t, baseURL, loginData.Token, loginData.UserID)

	cookieHeader := strings.Join(loginData.CookieHeader, "; ")
	if !strings.Contains(cookieHeader, "auth_token=") {
		t.Fatalf("expected auth_token cookie, got %q", cookieHeader)
	}
	if !strings.Contains(strings.ToLower(cookieHeader), "secure") {
		t.Fatalf("expected Secure cookie flag, got %q", cookieHeader)
	}
	if !strings.Contains(cookieHeader, "SameSite=None") {
		t.Fatalf("expected SameSite=None, got %q", cookieHeader)
	}
	if username == "" {
		t.Fatalf("expected username to be set")
	}
}

func TestAuth_SessionEndpoint(t *testing.T) {
	baseURL := os.Getenv("PB_URL")
	if baseURL == "" {
		t.Skip("PB_URL not set")
	}

	cfg, err := config.Load()
	if err != nil {
		t.Fatalf("load config: %v", err)
	}

	pbClient, err := repo_pb.NewClient(repo_pb.Config{BaseURL: baseURL, Timeout: 5 * time.Second})
	if err != nil {
		t.Fatalf("create pb client: %v", err)
	}

	handler := apihttp.NewAuthHandler(pbClient, apihttp.AuthCookieConfig{
		Secure:   cfg.AuthCookieSecure,
		SameSite: cfg.AuthCookieSameSite,
	})

	router := chi.NewRouter()
	apihttp.RegisterAuthRoutes(router, handler)

	server := httptest.NewServer(router)
	defer server.Close()

	_, email, password := registerUser(t, server.URL)
	loginData := loginUserData(t, server.URL, email, password)
	defer cleanupTestUser(t, baseURL, loginData.Token, loginData.UserID)

	request, err := http.NewRequest(http.MethodGet, server.URL+"/api/auth/session", nil)
	if err != nil {
		t.Fatalf("create session request: %v", err)
	}
	for _, cookie := range loginData.Cookies {
		request.AddCookie(cookie)
	}

	response, err := http.DefaultClient.Do(request)
	if err != nil {
		t.Fatalf("session request: %v", err)
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		t.Fatalf("expected 200, got %d", response.StatusCode)
	}
}

func TestAuth_SessionEndpoint_MissingCookie(t *testing.T) {
	baseURL := os.Getenv("PB_URL")
	if baseURL == "" {
		t.Skip("PB_URL not set")
	}

	cfg, err := config.Load()
	if err != nil {
		t.Fatalf("load config: %v", err)
	}

	pbClient, err := repo_pb.NewClient(repo_pb.Config{BaseURL: baseURL, Timeout: 5 * time.Second})
	if err != nil {
		t.Fatalf("create pb client: %v", err)
	}

	handler := apihttp.NewAuthHandler(pbClient, apihttp.AuthCookieConfig{
		Secure:   cfg.AuthCookieSecure,
		SameSite: cfg.AuthCookieSameSite,
	})

	router := chi.NewRouter()
	apihttp.RegisterAuthRoutes(router, handler)

	server := httptest.NewServer(router)
	defer server.Close()

	response, err := http.Get(server.URL + "/api/auth/session")
	if err != nil {
		t.Fatalf("session request: %v", err)
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusUnauthorized {
		t.Fatalf("expected 401, got %d", response.StatusCode)
	}
}

func registerUser(t *testing.T, serverURL string) (string, string, string) {
	t.Helper()

	if email := os.Getenv("PB_TEST_EMAIL"); email != "" {
		password := os.Getenv("PB_TEST_PASSWORD")
		if password == "" {
			t.Skip("PB_TEST_PASSWORD not set")
		}
		username := os.Getenv("PB_TEST_USERNAME")
		if username == "" {
			username = "test-user"
		}
		return username, email, password
	}

	username := "test-user-" + time.Now().UTC().Format("20060102150405.000000000")
	email := username + "@example.com"
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

	registerResp, err := http.Post(serverURL+"/api/auth/register", "application/json", bytes.NewReader(registerBody))
	if err != nil {
		t.Fatalf("register request: %v", err)
	}
	defer registerResp.Body.Close()
	if registerResp.StatusCode != http.StatusOK {
		t.Fatalf("register status: %d", registerResp.StatusCode)
	}

	return username, email, password
}

func loginUser(t *testing.T, serverURL, email, password string) *http.Response {
	t.Helper()

	loginPayload := map[string]string{
		"email":    email,
		"password": password,
	}
	loginBody, err := json.Marshal(loginPayload)
	if err != nil {
		t.Fatalf("marshal login payload: %v", err)
	}

	loginResp, err := http.Post(serverURL+"/api/auth/login", "application/json", bytes.NewReader(loginBody))
	if err != nil {
		t.Fatalf("login request: %v", err)
	}
	if loginResp.StatusCode != http.StatusOK {
		loginResp.Body.Close()
		t.Fatalf("login status: %d", loginResp.StatusCode)
	}

	return loginResp
}
