package http

import (
	"bytes"
	"context"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/digitaistudios/crims-backend/internal/ports"
)

type fakeClueService struct {
	createResult ports.ClueRecord
	createErr    error
	listResult   []ports.ClueRecord
	listErr      error
}

func (f fakeClueService) CreateClue(ctx context.Context, input ports.ClueRecordInput) (ports.ClueRecord, error) {
	if f.createErr != nil {
		return ports.ClueRecord{}, f.createErr
	}
	return f.createResult, nil
}

func (f fakeClueService) GetClueByID(ctx context.Context, id string) (ports.ClueRecord, error) {
	return ports.ClueRecord{}, nil
}

func (f fakeClueService) ListCluesByGame(ctx context.Context, gameID string) ([]ports.ClueRecord, error) {
	if f.listErr != nil {
		return nil, f.listErr
	}
	return f.listResult, nil
}

func TestCreateClueHandler_OK(t *testing.T) {
	service := fakeClueService{createResult: ports.ClueRecord{ID: "clue-1"}}
	handler := NewCreateClueHandler(service)

	payload := []byte(`{"gameId":"game-1","type":"OBJECT","state":"DISCOVERED","reliability":80}`)
	request := httptest.NewRequest(http.MethodPost, "/api/clues", bytes.NewReader(payload))
	response := httptest.NewRecorder()

	handler.ServeHTTP(response, request)

	if response.Code != http.StatusCreated {
		t.Fatalf("expected status 201, got %d", response.Code)
	}
}

func TestListCluesByGameHandler_OK(t *testing.T) {
	service := fakeClueService{listResult: []ports.ClueRecord{{ID: "clue-1"}}}
	handler := NewListCluesByGameHandler(service)

	request := httptest.NewRequest(http.MethodGet, "/api/games/game-1/clues", nil)
	request = request.WithContext(context.WithValue(request.Context(), idParamKey, "game-1"))
	response := httptest.NewRecorder()

	handler.ServeHTTP(response, request)

	if response.Code != http.StatusOK {
		t.Fatalf("expected status 200, got %d", response.Code)
	}
}
