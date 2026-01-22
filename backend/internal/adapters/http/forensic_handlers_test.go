package http

import (
	"bytes"
	"context"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/digitaistudios/crims-backend/internal/ports"
)

type fakeForensicService struct {
	createResult ports.ForensicRecord
	createErr    error
	listResult   []ports.ForensicRecord
	listErr      error
}

func (f fakeForensicService) CreateAnalysis(ctx context.Context, input ports.ForensicRecordInput) (ports.ForensicRecord, error) {
	if f.createErr != nil {
		return ports.ForensicRecord{}, f.createErr
	}
	return f.createResult, nil
}

func (f fakeForensicService) GetAnalysisByID(ctx context.Context, id string) (ports.ForensicRecord, error) {
	return ports.ForensicRecord{}, nil
}

func (f fakeForensicService) ListAnalysesByGame(ctx context.Context, gameID string) ([]ports.ForensicRecord, error) {
	if f.listErr != nil {
		return nil, f.listErr
	}
	return f.listResult, nil
}

func TestCreateAnalysisHandler_OK(t *testing.T) {
	service := fakeForensicService{createResult: ports.ForensicRecord{ID: "forensic-1"}}
	handler := NewCreateForensicHandler(service)

	payload := []byte(`{"gameId":"game-1","clueId":"clue-1","result":"match","confidence":90,"status":"DONE"}`)
	request := httptest.NewRequest(http.MethodPost, "/api/forensics", bytes.NewReader(payload))
	response := httptest.NewRecorder()

	handler.ServeHTTP(response, request)

	if response.Code != http.StatusCreated {
		t.Fatalf("expected status 201, got %d", response.Code)
	}
}

func TestListAnalysesByGameHandler_OK(t *testing.T) {
	service := fakeForensicService{listResult: []ports.ForensicRecord{{ID: "forensic-1"}}}
	handler := NewListForensicsByGameHandler(service)

	request := httptest.NewRequest(http.MethodGet, "/api/games/game-1/forensics", nil)
	request = request.WithContext(context.WithValue(request.Context(), idParamKey, "game-1"))
	response := httptest.NewRecorder()

	handler.ServeHTTP(response, request)

	if response.Code != http.StatusOK {
		t.Fatalf("expected status 200, got %d", response.Code)
	}
}
