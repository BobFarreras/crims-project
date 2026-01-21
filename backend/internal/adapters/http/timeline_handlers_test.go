package http

import (
	"bytes"
	"context"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/digitaistudios/crims-backend/internal/ports"
)

type fakeTimelineService struct {
	createResult ports.TimelineRecord
	createErr    error
	listResult   []ports.TimelineRecord
	listErr      error
}

func (f fakeTimelineService) CreateEntry(ctx context.Context, input ports.TimelineRecordInput) (ports.TimelineRecord, error) {
	if f.createErr != nil {
		return ports.TimelineRecord{}, f.createErr
	}
	return f.createResult, nil
}

func (f fakeTimelineService) GetEntryByID(ctx context.Context, id string) (ports.TimelineRecord, error) {
	return ports.TimelineRecord{}, nil
}

func (f fakeTimelineService) ListEntriesByGame(ctx context.Context, gameID string) ([]ports.TimelineRecord, error) {
	if f.listErr != nil {
		return nil, f.listErr
	}
	return f.listResult, nil
}

func TestCreateTimelineHandler_OK(t *testing.T) {
	service := fakeTimelineService{createResult: ports.TimelineRecord{ID: "entry-1"}}
	handler := NewCreateTimelineHandler(service)

	payload := []byte(`{"gameId":"game-1","timestamp":"2026-01-21T10:00:00Z","title":"Arrival","description":"Victim arrives"}`)
	request := httptest.NewRequest(http.MethodPost, "/api/timeline", bytes.NewReader(payload))
	response := httptest.NewRecorder()

	handler.ServeHTTP(response, request)

	if response.Code != http.StatusCreated {
		t.Fatalf("expected status 201, got %d", response.Code)
	}
}

func TestListTimelineByGameHandler_OK(t *testing.T) {
	service := fakeTimelineService{listResult: []ports.TimelineRecord{{ID: "entry-1"}}}
	handler := NewListTimelineByGameHandler(service)

	request := httptest.NewRequest(http.MethodGet, "/api/games/game-1/timeline", nil)
	request = request.WithContext(context.WithValue(request.Context(), idParamKey, "game-1"))
	response := httptest.NewRecorder()

	handler.ServeHTTP(response, request)

	if response.Code != http.StatusOK {
		t.Fatalf("expected status 200, got %d", response.Code)
	}
}
