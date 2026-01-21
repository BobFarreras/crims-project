package http

import (
	"bytes"
	"context"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/digitaistudios/crims-backend/internal/ports"
)

type fakeLobbyService struct {
	joinResult ports.PlayerRecord
	joinErr    error
}

func (f fakeLobbyService) JoinGame(ctx context.Context, gameCode, userID, role string) (ports.PlayerRecord, error) {
	if f.joinErr != nil {
		return ports.PlayerRecord{}, f.joinErr
	}
	return f.joinResult, nil
}

func (f fakeLobbyService) ListPlayers(ctx context.Context, gameID string) ([]ports.PlayerRecord, error) {
	return nil, nil
}

func TestLobbyJoinHandler_OK(t *testing.T) {
	service := fakeLobbyService{joinResult: ports.PlayerRecord{ID: "player-1"}}
	handler := NewLobbyJoinHandler(service)

	payload := []byte(`{"gameCode":"ABCD","userId":"user-1","role":"DETECTIVE"}`)
	request := httptest.NewRequest(http.MethodPost, "/api/lobby/join", bytes.NewReader(payload))
	response := httptest.NewRecorder()

	handler.ServeHTTP(response, request)

	if response.Code != http.StatusCreated {
		t.Fatalf("expected status 201, got %d", response.Code)
	}
}
