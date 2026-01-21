package services

import (
	"context"
	"errors"
	"testing"

	"github.com/digitaistudios/crims-backend/internal/ports"
)

type fakeGameRepository struct {
	lastCreate ports.GameRecordInput
	createResult ports.GameRecord
	createErr    error
	getByID      ports.GameRecord
	getByIDErr   error
	getByCode    ports.GameRecord
	getByCodeErr error
}

func (f *fakeGameRepository) CreateGame(ctx context.Context, input ports.GameRecordInput) (ports.GameRecord, error) {
	f.lastCreate = input
	if f.createErr != nil {
		return ports.GameRecord{}, f.createErr
	}
	return f.createResult, nil
}

func (f *fakeGameRepository) GetGameByID(ctx context.Context, id string) (ports.GameRecord, error) {
	if f.getByIDErr != nil {
		return ports.GameRecord{}, f.getByIDErr
	}
	return f.getByID, nil
}

func (f *fakeGameRepository) GetGameByCode(ctx context.Context, code string) (ports.GameRecord, error) {
	if f.getByCodeErr != nil {
		return ports.GameRecord{}, f.getByCodeErr
	}
	return f.getByCode, nil
}

func TestGameService_CreateGame_Invalid(t *testing.T) {
	service := NewGameService(&fakeGameRepository{})

	_, err := service.CreateGame(context.Background(), ports.GameRecordInput{})
	if !errors.Is(err, ErrInvalidGameInput) {
		t.Fatalf("expected ErrInvalidGameInput, got %v", err)
	}
}

func TestGameService_CreateGame_OK(t *testing.T) {
	repo := &fakeGameRepository{createResult: ports.GameRecord{ID: "game-1"}}
	service := NewGameService(repo)

	result, err := service.CreateGame(context.Background(), ports.GameRecordInput{
		Code:  "ABCD",
		State: "INVESTIGATION",
		Seed:  "seed-1",
	})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if result.ID != "game-1" {
		t.Fatalf("expected id game-1, got %s", result.ID)
	}
	if repo.lastCreate.Code != "ABCD" {
		t.Fatalf("expected repo call, got %s", repo.lastCreate.Code)
	}
}

func TestGameService_GetGameByID_Invalid(t *testing.T) {
	service := NewGameService(&fakeGameRepository{})

	_, err := service.GetGameByID(context.Background(), "")
	if !errors.Is(err, ErrMissingGameID) {
		t.Fatalf("expected ErrMissingGameID, got %v", err)
	}
}

func TestGameService_GetGameByCode_Invalid(t *testing.T) {
	service := NewGameService(&fakeGameRepository{})

	_, err := service.GetGameByCode(context.Background(), "")
	if !errors.Is(err, ErrMissingGameCode) {
		t.Fatalf("expected ErrMissingGameCode, got %v", err)
	}
}
