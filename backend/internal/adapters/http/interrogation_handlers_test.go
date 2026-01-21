package http

import (
	"bytes"
	"context"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/digitaistudios/crims-backend/internal/ports"
)

type fakeInterrogationService struct {
	createResult ports.InterrogationRecord
	createErr    error
	listResult   []ports.InterrogationRecord
	listErr      error
}

func (f fakeInterrogationService) CreateInterrogation(ctx context.Context, input ports.InterrogationRecordInput) (ports.InterrogationRecord, error) {
	if f.createErr != nil {
		return ports.InterrogationRecord{}, f.createErr
	}
	return f.createResult, nil
}

func (f fakeInterrogationService) GetInterrogationByID(ctx context.Context, id string) (ports.InterrogationRecord, error) {
	return ports.InterrogationRecord{}, nil
}

func (f fakeInterrogationService) ListInterrogationsByGame(ctx context.Context, gameID string) ([]ports.InterrogationRecord, error) {
	if f.listErr != nil {
		return nil, f.listErr
	}
	return f.listResult, nil
}

func TestCreateInterrogationHandler_OK(t *testing.T) {
	service := fakeInterrogationService{createResult: ports.InterrogationRecord{ID: "int-1"}}
	handler := NewCreateInterrogationHandler(service)

	payload := []byte(`{"gameId":"game-1","personId":"person-1","question":"Where were you?","answer":"At home.","tone":"neutral"}`)
	request := httptest.NewRequest(http.MethodPost, "/api/interrogations", bytes.NewReader(payload))
	response := httptest.NewRecorder()

	handler.ServeHTTP(response, request)

	if response.Code != http.StatusCreated {
		t.Fatalf("expected status 201, got %d", response.Code)
	}
}

func TestListInterrogationsByGameHandler_OK(t *testing.T) {
	service := fakeInterrogationService{listResult: []ports.InterrogationRecord{{ID: "int-1"}}}
	handler := NewListInterrogationsByGameHandler(service)

	request := httptest.NewRequest(http.MethodGet, "/api/games/game-1/interrogations", nil)
	request = request.WithContext(context.WithValue(request.Context(), idParamKey, "game-1"))
	response := httptest.NewRecorder()

	handler.ServeHTTP(response, request)

	if response.Code != http.StatusOK {
		t.Fatalf("expected status 200, got %d", response.Code)
	}
}
