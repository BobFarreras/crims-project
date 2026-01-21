package http

import (
	"bytes"
	"context"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/digitaistudios/crims-backend/internal/ports"
)

type fakeEventService struct {
	createResult ports.EventRecord
	createErr    error
	listResult   []ports.EventRecord
	listErr      error
}

func (f fakeEventService) CreateEvent(ctx context.Context, input ports.EventRecordInput) (ports.EventRecord, error) {
	if f.createErr != nil {
		return ports.EventRecord{}, f.createErr
	}
	return f.createResult, nil
}

func (f fakeEventService) GetEventByID(ctx context.Context, id string) (ports.EventRecord, error) {
	return ports.EventRecord{}, nil
}

func (f fakeEventService) ListEventsByGame(ctx context.Context, gameID string) ([]ports.EventRecord, error) {
	if f.listErr != nil {
		return nil, f.listErr
	}
	return f.listResult, nil
}

func TestCreateEventHandler_OK(t *testing.T) {
	service := fakeEventService{createResult: ports.EventRecord{ID: "event-1"}}
	handler := NewCreateEventHandler(service)

	payload := []byte(`{"gameId":"game-1","timestamp":"2026-01-21T10:00:00Z","locationId":"loc-1"}`)
	request := httptest.NewRequest(http.MethodPost, "/api/events", bytes.NewReader(payload))
	response := httptest.NewRecorder()

	handler.ServeHTTP(response, request)

	if response.Code != http.StatusCreated {
		t.Fatalf("expected status 201, got %d", response.Code)
	}
}

func TestListEventsByGameHandler_OK(t *testing.T) {
	service := fakeEventService{listResult: []ports.EventRecord{{ID: "event-1"}}}
	handler := NewListEventsByGameHandler(service)

	request := httptest.NewRequest(http.MethodGet, "/api/games/game-1/events", nil)
	request = request.WithContext(context.WithValue(request.Context(), idParamKey, "game-1"))
	response := httptest.NewRecorder()

	handler.ServeHTTP(response, request)

	if response.Code != http.StatusOK {
		t.Fatalf("expected status 200, got %d", response.Code)
	}
}
