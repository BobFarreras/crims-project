package http

import (
	"bytes"
	"context"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/digitaistudios/crims-backend/internal/ports"
)

type fakePersonService struct {
	createResult ports.PersonRecord
	createErr    error
	listResult   []ports.PersonRecord
	listErr      error
}

func (f fakePersonService) CreatePerson(ctx context.Context, input ports.PersonRecordInput) (ports.PersonRecord, error) {
	if f.createErr != nil {
		return ports.PersonRecord{}, f.createErr
	}
	return f.createResult, nil
}

func (f fakePersonService) GetPersonByID(ctx context.Context, id string) (ports.PersonRecord, error) {
	return ports.PersonRecord{}, nil
}

func (f fakePersonService) ListPersonsByGame(ctx context.Context, gameID string) ([]ports.PersonRecord, error) {
	if f.listErr != nil {
		return nil, f.listErr
	}
	return f.listResult, nil
}

func TestCreatePersonHandler_OK(t *testing.T) {
	service := fakePersonService{createResult: ports.PersonRecord{ID: "person-1"}}
	handler := NewCreatePersonHandler(service)

	payload := []byte(`{"gameId":"game-1","name":"Witness","officialStory":"story","truthStory":"truth","stress":10,"credibility":80}`)
	request := httptest.NewRequest(http.MethodPost, "/api/persons", bytes.NewReader(payload))
	response := httptest.NewRecorder()

	handler.ServeHTTP(response, request)

	if response.Code != http.StatusCreated {
		t.Fatalf("expected status 201, got %d", response.Code)
	}
}

func TestListPersonsByGameHandler_OK(t *testing.T) {
	service := fakePersonService{listResult: []ports.PersonRecord{{ID: "person-1"}}}
	handler := NewListPersonsByGameHandler(service)

	request := httptest.NewRequest(http.MethodGet, "/api/games/game-1/persons", nil)
	request = request.WithContext(context.WithValue(request.Context(), idParamKey, "game-1"))
	response := httptest.NewRecorder()

	handler.ServeHTTP(response, request)

	if response.Code != http.StatusOK {
		t.Fatalf("expected status 200, got %d", response.Code)
	}
}
