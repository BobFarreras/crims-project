package integration

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"
	"os"
	"testing"
	"time"

	"github.com/digitaistudios/crims-backend/internal/adapters/repo_pb"
)

func TestPocketBase_GameLifecycle(t *testing.T) {
	baseURL := os.Getenv("PB_URL")
	adminToken := os.Getenv("PB_ADMIN_TOKEN")
	if baseURL == "" || adminToken == "" {
		t.Skip("PB_URL or PB_ADMIN_TOKEN not set")
	}

	client := &http.Client{Timeout: 5 * time.Second}
	payload := map[string]string{
		"code":  "ABCD",
		"state": "INVESTIGATION",
		"seed":  "seed-1",
	}
	data, err := json.Marshal(payload)
	if err != nil {
		t.Fatalf("marshal payload: %v", err)
	}

	request, err := http.NewRequest(http.MethodPost, baseURL+"/api/collections/games/records", bytes.NewReader(data))
	if err != nil {
		t.Fatalf("create request: %v", err)
	}
	request.Header.Set("Authorization", "Bearer "+adminToken)
	request.Header.Set("Content-Type", "application/json")

	response, err := client.Do(request)
	if err != nil {
		t.Fatalf("create record: %v", err)
	}
	defer response.Body.Close()
	if response.StatusCode != http.StatusOK && response.StatusCode != http.StatusCreated {
		t.Fatalf("unexpected status: %d", response.StatusCode)
	}

	var created struct {
		ID string `json:"id"`
	}
	if err := json.NewDecoder(response.Body).Decode(&created); err != nil {
		t.Fatalf("decode create response: %v", err)
	}

	repoClient, err := repo_pb.NewClient(repo_pb.Config{BaseURL: baseURL, Timeout: 2 * time.Second})
	if err != nil {
		t.Fatalf("repo client: %v", err)
	}

	repo := repo_pb.NewGameRepository(repoClient)
	game, err := repo.GetGameByID(context.Background(), created.ID)
	if err != nil {
		t.Fatalf("get game: %v", err)
	}
	if game.Code != "ABCD" {
		t.Fatalf("expected code ABCD, got %s", game.Code)
	}

	deleteReq, err := http.NewRequest(http.MethodDelete, baseURL+"/api/collections/games/records/"+created.ID, nil)
	if err != nil {
		t.Fatalf("delete request: %v", err)
	}
	deleteReq.Header.Set("Authorization", "Bearer "+adminToken)

	deleteResp, err := client.Do(deleteReq)
	if err != nil {
		t.Fatalf("delete record: %v", err)
	}
	defer deleteResp.Body.Close()
}
