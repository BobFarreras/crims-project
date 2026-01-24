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
	createResult ports.LobbyState
	createErr    error
	joinResult   ports.PlayerRecord
	joinErr      error
}

type fakeAuthClient struct{}

func (f fakeAuthClient) Ping(_ context.Context) error                              { return nil }
func (f fakeAuthClient) CreateUser(_, _, _, _, _ string) error                     { return nil }
func (f fakeAuthClient) AuthWithPassword(_, _ string) (*ports.AuthResponse, error) { return nil, nil }
func (f fakeAuthClient) RefreshAuth(_ string) (*ports.AuthResponse, error) {
	resp := &ports.AuthResponse{}
	resp.Record.ID = "user-1"
	return resp, nil
}
func (f fakeAuthClient) UpdateUserName(_, _, _ string) error { return nil }

func (f fakeLobbyService) CreateLobby(ctx context.Context, userID string, capabilities []string) (ports.LobbyState, error) {
	if f.createErr != nil {
		return ports.LobbyState{}, f.createErr
	}
	return f.createResult, nil
}

func (f fakeLobbyService) JoinGame(ctx context.Context, gameCode, userID string, capabilities []string) (ports.PlayerRecord, error) {
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
	handler := NewLobbyJoinHandler(service, fakeAuthClient{})

	payload := []byte(`{"gameCode":"ABCD","userId":"user-1","capabilities":["DETECTIVE"]}`)
	request := httptest.NewRequest(http.MethodPost, "/api/lobby/join", bytes.NewReader(payload))
	request.AddCookie(&http.Cookie{Name: "auth_token", Value: "token"})
	response := httptest.NewRecorder()

	handler.ServeHTTP(response, request)

	if response.Code != http.StatusCreated {
		t.Fatalf("expected status 201, got %d", response.Code)
	}
}

func TestLobbyCreateHandler_OK(t *testing.T) {
	service := fakeLobbyService{createResult: ports.LobbyState{Game: ports.GameRecord{Code: "ABCD"}}}
	handler := NewLobbyCreateHandler(service, fakeAuthClient{})

	payload := []byte(`{"userId":"user-1","capabilities":["DETECTIVE"]}`)
	request := httptest.NewRequest(http.MethodPost, "/api/lobby/create", bytes.NewReader(payload))
	request.AddCookie(&http.Cookie{Name: "auth_token", Value: "token"})
	response := httptest.NewRecorder()

	handler.ServeHTTP(response, request)

	if response.Code != http.StatusCreated {
		t.Fatalf("expected status 201, got %d", response.Code)
	}
}
