package http

import (
	"bytes"
	"context"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/digitaistudios/crims-backend/internal/adapters/repo_pb"
	"github.com/digitaistudios/crims-backend/internal/ports"
)

type fakeGameRepository struct {
	createResult ports.GameRecord
	createErr    error
	getByID      ports.GameRecord
	getByIDErr   error
	getByCode    ports.GameRecord
	getByCodeErr error
}

func (f fakeGameRepository) CreateGame(ctx context.Context, input ports.GameRecordInput) (ports.GameRecord, error) {
	if f.createErr != nil {
		return ports.GameRecord{}, f.createErr
	}
	return f.createResult, nil
}

func (f fakeGameRepository) GetGameByID(ctx context.Context, id string) (ports.GameRecord, error) {
	if f.getByIDErr != nil {
		return ports.GameRecord{}, f.getByIDErr
	}
	return f.getByID, nil
}

func (f fakeGameRepository) GetGameByCode(ctx context.Context, code string) (ports.GameRecord, error) {
	if f.getByCodeErr != nil {
		return ports.GameRecord{}, f.getByCodeErr
	}
	return f.getByCode, nil
}

func TestCreateGameHandler_OK(t *testing.T) {
	repo := fakeGameRepository{
		createResult: ports.GameRecord{ID: "game-1", Code: "ABCD", State: "INVESTIGATION", Seed: "seed-1"},
	}
	handler := NewCreateGameHandler(repo)

	payload := []byte(`{"code":"ABCD","state":"INVESTIGATION","seed":"seed-1"}`)
	request := httptest.NewRequest(http.MethodPost, "/api/games", bytes.NewReader(payload))
	response := httptest.NewRecorder()

	handler.ServeHTTP(response, request)

	if response.Code != http.StatusCreated {
		t.Fatalf("expected status 201, got %d", response.Code)
	}
}

func TestCreateGameHandler_InvalidInput(t *testing.T) {
	repo := fakeGameRepository{}
	handler := NewCreateGameHandler(repo)

	request := httptest.NewRequest(http.MethodPost, "/api/games", bytes.NewReader([]byte(`{invalid`)))
	response := httptest.NewRecorder()

	handler.ServeHTTP(response, request)

	if response.Code != http.StatusBadRequest {
		t.Fatalf("expected status 400, got %d", response.Code)
	}
}

func TestGetGameByIDHandler_OK(t *testing.T) {
	repo := fakeGameRepository{
		getByID: ports.GameRecord{ID: "game-1", Code: "ABCD", State: "INVESTIGATION", Seed: "seed-1"},
	}
	handler := NewGetGameByIDHandler(repo)

	request := httptest.NewRequest(http.MethodGet, "/api/games/game-1", nil)
	request = request.WithContext(context.WithValue(request.Context(), idParamKey, "game-1"))
	response := httptest.NewRecorder()

	handler.ServeHTTP(response, request)

	if response.Code != http.StatusOK {
		t.Fatalf("expected status 200, got %d", response.Code)
	}
}

func TestGetGameByCodeHandler_NotFound(t *testing.T) {
	repo := fakeGameRepository{getByCodeErr: repo_pb.ErrRecordNotFound}
	handler := NewGetGameByCodeHandler(repo)

	request := httptest.NewRequest(http.MethodGet, "/api/games/by-code/ABCD", nil)
	request = request.WithContext(context.WithValue(request.Context(), codeParamKey, "ABCD"))
	response := httptest.NewRecorder()

	handler.ServeHTTP(response, request)

	if response.Code != http.StatusNotFound {
		t.Fatalf("expected status 404, got %d", response.Code)
	}
}
