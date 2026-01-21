package http

import (
	"bytes"
	"context"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/digitaistudios/crims-backend/internal/ports"
)

type fakeHypothesisService struct {
	createResult ports.HypothesisRecord
	createErr    error
	listResult   []ports.HypothesisRecord
	listErr      error
}

func (f fakeHypothesisService) CreateHypothesis(ctx context.Context, input ports.HypothesisRecordInput) (ports.HypothesisRecord, error) {
	if f.createErr != nil {
		return ports.HypothesisRecord{}, f.createErr
	}
	return f.createResult, nil
}

func (f fakeHypothesisService) GetHypothesisByID(ctx context.Context, id string) (ports.HypothesisRecord, error) {
	return ports.HypothesisRecord{}, nil
}

func (f fakeHypothesisService) ListHypothesesByGame(ctx context.Context, gameID string) ([]ports.HypothesisRecord, error) {
	if f.listErr != nil {
		return nil, f.listErr
	}
	return f.listResult, nil
}

func TestCreateHypothesisHandler_OK(t *testing.T) {
	service := fakeHypothesisService{createResult: ports.HypothesisRecord{ID: "hyp-1"}}
	handler := NewCreateHypothesisHandler(service)

	payload := []byte(`{"gameId":"game-1","title":"Main theory","strengthScore":42,"status":"PLAUSIBLE"}`)
	request := httptest.NewRequest(http.MethodPost, "/api/hypotheses", bytes.NewReader(payload))
	response := httptest.NewRecorder()

	handler.ServeHTTP(response, request)

	if response.Code != http.StatusCreated {
		t.Fatalf("expected status 201, got %d", response.Code)
	}
}

func TestListHypothesesByGameHandler_OK(t *testing.T) {
	service := fakeHypothesisService{listResult: []ports.HypothesisRecord{{ID: "hyp-1"}}}
	handler := NewListHypothesesByGameHandler(service)

	request := httptest.NewRequest(http.MethodGet, "/api/games/game-1/hypotheses", nil)
	request = request.WithContext(context.WithValue(request.Context(), idParamKey, "game-1"))
	response := httptest.NewRecorder()

	handler.ServeHTTP(response, request)

	if response.Code != http.StatusOK {
		t.Fatalf("expected status 200, got %d", response.Code)
	}
}
