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

	loginData := loginUserData(t, server.URL, email, password)
	defer cleanupTestUser(t, baseURL, loginData.Token, loginData.UserID)

	cookieHeader := strings.Join(loginData.CookieHeader, "; ")
	if !strings.Contains(cookieHeader, "auth_token=") {
		t.Fatalf("expected auth_token cookie, got %q", cookieHeader)
	}
	if !strings.Contains(strings.ToLower(cookieHeader), "httponly") {
		t.Fatalf("expected HttpOnly cookie, got %q", cookieHeader)
	}
	if loginData.Username != username {
		t.Fatalf("expected username %s, got %v", username, loginData.Username)
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

type loginData struct {
	Token        string
	UserID       string
	Username     string
	CookieHeader []string
	Cookies      []*http.Cookie
}

func loginUserData(t *testing.T, serverURL, email, password string) loginData {
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
	defer loginResp.Body.Close()
	if loginResp.StatusCode != http.StatusOK {
		t.Fatalf("login status: %d", loginResp.StatusCode)
	}

	var response map[string]interface{}
	if err := json.NewDecoder(loginResp.Body).Decode(&response); err != nil {
		t.Fatalf("decode login response: %v", err)
	}
	user, ok := response["user"].(map[string]interface{})
	if !ok {
		t.Fatalf("expected user object in response")
	}
	userID, _ := user["id"].(string)
	username, _ := user["username"].(string)
	if userID == "" {
		t.Fatalf("expected user id in response")
	}

	var token string
	for _, cookie := range loginResp.Cookies() {
		if cookie.Name == "auth_token" {
			token = cookie.Value
			break
		}
	}
	if token == "" {
		t.Fatalf("missing auth_token cookie")
	}

	return loginData{
		Token:        token,
		UserID:       userID,
		Username:     username,
		CookieHeader: loginResp.Header.Values("Set-Cookie"),
		Cookies:      loginResp.Cookies(),
	}
}

func cleanupTestUser(t *testing.T, baseURL, token, userID string) {
	t.Helper()
	adminToken := os.Getenv("PB_ADMIN_TOKEN")
	if os.Getenv("PB_TEST_EMAIL") != "" && adminToken == "" {
		return
	}
	if baseURL == "" || token == "" || userID == "" {
		return
	}

	request, err := http.NewRequest(http.MethodDelete, baseURL+"/api/collections/users/records/"+userID, nil)
	if err != nil {
		return
	}
	if adminToken != "" {
		request.Header.Set("Authorization", "Bearer "+adminToken)
	} else {
		request.Header.Set("Authorization", "Bearer "+token)
	}

	response, err := http.DefaultClient.Do(request)
	if err != nil {
		return
	}
	defer response.Body.Close()

	if response.StatusCode == http.StatusForbidden || response.StatusCode == http.StatusNotFound {
		return
	}
}
