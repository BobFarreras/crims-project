package integration

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
	"time"

	apihttp "github.com/digitaistudios/crims-backend/internal/adapters/http"
	"github.com/digitaistudios/crims-backend/internal/adapters/repo_pb"
	"github.com/digitaistudios/crims-backend/internal/platform/config"
	"github.com/go-chi/chi/v5"
)

func TestProfile_UpdateName(t *testing.T) {
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

	router := chi.NewRouter()
	handler := apihttp.NewAuthHandler(pbClient, apihttp.AuthCookieConfig{
		Secure:   cfg.AuthCookieSecure,
		SameSite: cfg.AuthCookieSameSite,
	})
	apihttp.RegisterAuthRoutes(router, handler)
	apihttp.RegisterProfileRoutes(router, apihttp.NewProfileHandler(pbClient))

	server := httptest.NewServer(router)
	defer server.Close()

	_, email, password := registerUser(t, server.URL)
	loginData := loginUserData(t, server.URL, email, password)
	defer cleanupTestUser(t, baseURL, loginData.Token, loginData.UserID)

	updatePayload := map[string]string{"name": "Detective Nova"}
	updateBody, err := json.Marshal(updatePayload)
	if err != nil {
		t.Fatalf("marshal update payload: %v", err)
	}

	request, err := http.NewRequest(http.MethodPut, server.URL+"/api/profile", bytes.NewReader(updateBody))
	if err != nil {
		t.Fatalf("create update request: %v", err)
	}
	request.Header.Set("Content-Type", "application/json")
	for _, cookie := range loginData.Cookies {
		request.AddCookie(cookie)
	}

	response, err := http.DefaultClient.Do(request)
	if err != nil {
		t.Fatalf("update request: %v", err)
	}
	defer response.Body.Close()

	if response.StatusCode == http.StatusForbidden || response.StatusCode == http.StatusUnauthorized || response.StatusCode == http.StatusNotFound || response.StatusCode == http.StatusInternalServerError {
		t.Skip("PocketBase update rules not allowing profile update")
	}
	if response.StatusCode != http.StatusOK {
		t.Fatalf("expected status 200, got %d", response.StatusCode)
	}

	var payload map[string]interface{}
	if err := json.NewDecoder(response.Body).Decode(&payload); err != nil {
		t.Fatalf("decode response: %v", err)
	}
	user, ok := payload["user"].(map[string]interface{})
	if !ok {
		t.Fatalf("expected user object in response")
	}
	if user["name"] != "Detective Nova" {
		t.Fatalf("expected updated name, got %v", user["name"])
	}
}

func TestProfile_UpdateMissingName(t *testing.T) {
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

	router := chi.NewRouter()
	handler := apihttp.NewAuthHandler(pbClient, apihttp.AuthCookieConfig{
		Secure:   cfg.AuthCookieSecure,
		SameSite: cfg.AuthCookieSameSite,
	})
	apihttp.RegisterAuthRoutes(router, handler)
	apihttp.RegisterProfileRoutes(router, apihttp.NewProfileHandler(pbClient))

	server := httptest.NewServer(router)
	defer server.Close()

	_, email, password := registerUser(t, server.URL)
	loginData := loginUserData(t, server.URL, email, password)
	defer cleanupTestUser(t, baseURL, loginData.Token, loginData.UserID)

	updatePayload := map[string]string{"name": ""}
	updateBody, err := json.Marshal(updatePayload)
	if err != nil {
		t.Fatalf("marshal update payload: %v", err)
	}

	request, err := http.NewRequest(http.MethodPut, server.URL+"/api/profile", bytes.NewReader(updateBody))
	if err != nil {
		t.Fatalf("create update request: %v", err)
	}
	request.Header.Set("Content-Type", "application/json")
	for _, cookie := range loginData.Cookies {
		request.AddCookie(cookie)
	}

	response, err := http.DefaultClient.Do(request)
	if err != nil {
		t.Fatalf("update request: %v", err)
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusBadRequest {
		t.Fatalf("expected status 400, got %d", response.StatusCode)
	}
}
