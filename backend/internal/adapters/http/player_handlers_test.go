package http

import (
	"bytes"
	"context"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/digitaistudios/crims-backend/internal/ports"
)

type fakePlayerService struct {
	createResult ports.PlayerRecord
	createErr    error
	listResult   []ports.PlayerRecord
	listErr      error
}

func (f fakePlayerService) CreatePlayer(ctx context.Context, input ports.PlayerRecordInput) (ports.PlayerRecord, error) {
	if f.createErr != nil {
		return ports.PlayerRecord{}, f.createErr
	}
	return f.createResult, nil
}

func (f fakePlayerService) GetPlayerByID(ctx context.Context, id string) (ports.PlayerRecord, error) {
	return ports.PlayerRecord{}, nil
}

func (f fakePlayerService) ListPlayersByGame(ctx context.Context, gameID string) ([]ports.PlayerRecord, error) {
	if f.listErr != nil {
		return nil, f.listErr
	}
	return f.listResult, nil
}

func TestCreatePlayerHandler_OK(t *testing.T) {
	service := fakePlayerService{createResult: ports.PlayerRecord{ID: "player-1"}}
	handler := NewCreatePlayerHandler(service)

	payload := []byte(`{"gameId":"game-1","userId":"user-1","role":"DETECTIVE","status":"ONLINE","isHost":true}`)
	request := httptest.NewRequest(http.MethodPost, "/api/players", bytes.NewReader(payload))
	response := httptest.NewRecorder()

	handler.ServeHTTP(response, request)

	if response.Code != http.StatusCreated {
		t.Fatalf("expected status 201, got %d", response.Code)
	}
}

func TestListPlayersByGameHandler_OK(t *testing.T) {
	service := fakePlayerService{listResult: []ports.PlayerRecord{{ID: "player-1"}}}
	handler := NewListPlayersByGameHandler(service)

	request := httptest.NewRequest(http.MethodGet, "/api/games/game-1/players", nil)
	request = request.WithContext(context.WithValue(request.Context(), idParamKey, "game-1"))
	response := httptest.NewRecorder()

	handler.ServeHTTP(response, request)

	if response.Code != http.StatusOK {
		t.Fatalf("expected status 200, got %d", response.Code)
	}
}
