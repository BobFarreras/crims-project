package http

import (
	"bytes"
	"context"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/digitaistudios/crims-backend/internal/ports"
)

type fakeAccusationService struct {
	createResult ports.AccusationRecord
	createErr    error
	listResult   []ports.AccusationRecord
	listErr      error
}

func (f fakeAccusationService) CreateAccusation(ctx context.Context, input ports.AccusationRecordInput) (ports.AccusationRecord, error) {
	if f.createErr != nil {
		return ports.AccusationRecord{}, f.createErr
	}
	return f.createResult, nil
}

func (f fakeAccusationService) GetAccusationByID(ctx context.Context, id string) (ports.AccusationRecord, error) {
	return ports.AccusationRecord{}, nil
}

func (f fakeAccusationService) ListAccusationsByGame(ctx context.Context, gameID string) ([]ports.AccusationRecord, error) {
	if f.listErr != nil {
		return nil, f.listErr
	}
	return f.listResult, nil
}

func TestCreateAccusationHandler_OK(t *testing.T) {
	service := fakeAccusationService{createResult: ports.AccusationRecord{ID: "acc-1"}}
	handler := NewCreateAccusationHandler(service)

	payload := []byte(`{"gameId":"game-1","playerId":"player-1","suspectId":"person-1","motiveId":"motive-1","evidenceId":"clue-1"}`)
	request := httptest.NewRequest(http.MethodPost, "/api/accusations", bytes.NewReader(payload))
	response := httptest.NewRecorder()

	handler.ServeHTTP(response, request)

	if response.Code != http.StatusCreated {
		t.Fatalf("expected status 201, got %d", response.Code)
	}
}

func TestListAccusationsByGameHandler_OK(t *testing.T) {
	service := fakeAccusationService{listResult: []ports.AccusationRecord{{ID: "acc-1"}}}
	handler := NewListAccusationsByGameHandler(service)

	request := httptest.NewRequest(http.MethodGet, "/api/games/game-1/accusations", nil)
	request = request.WithContext(context.WithValue(request.Context(), idParamKey, "game-1"))
	response := httptest.NewRecorder()

	handler.ServeHTTP(response, request)

	if response.Code != http.StatusOK {
		t.Fatalf("expected status 200, got %d", response.Code)
	}
}
