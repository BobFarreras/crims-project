package services

import (
	"context"
	"errors"
	"testing"

	"github.com/digitaistudios/crims-backend/internal/ports"
)

type fakePlayerRepository struct {
	lastCreate   ports.PlayerRecordInput
	createResult ports.PlayerRecord
	createErr    error
	listResult   []ports.PlayerRecord
	listErr      error
}

func (f *fakePlayerRepository) CreatePlayer(ctx context.Context, input ports.PlayerRecordInput) (ports.PlayerRecord, error) {
	f.lastCreate = input
	if f.createErr != nil {
		return ports.PlayerRecord{}, f.createErr
	}
	return f.createResult, nil
}

func (f *fakePlayerRepository) GetPlayerByID(ctx context.Context, id string) (ports.PlayerRecord, error) {
	return ports.PlayerRecord{}, nil
}

func (f *fakePlayerRepository) ListPlayersByGame(ctx context.Context, gameID string) ([]ports.PlayerRecord, error) {
	if f.listErr != nil {
		return nil, f.listErr
	}
	return f.listResult, nil
}

func TestPlayerService_CreatePlayer_Invalid(t *testing.T) {
	service := NewPlayerService(&fakePlayerRepository{})

	_, err := service.CreatePlayer(context.Background(), ports.PlayerRecordInput{})
	if !errors.Is(err, ErrInvalidPlayerInput) {
		t.Fatalf("expected ErrInvalidPlayerInput, got %v", err)
	}
}

func TestPlayerService_CreatePlayer_OK(t *testing.T) {
	repo := &fakePlayerRepository{createResult: ports.PlayerRecord{ID: "player-1"}}
	service := NewPlayerService(repo)

	result, err := service.CreatePlayer(context.Background(), ports.PlayerRecordInput{
		GameID:       "game-1",
		UserID:       "user-1",
		Capabilities: []string{"DETECTIVE"},
		Status:       "ONLINE",
	})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if result.ID != "player-1" {
		t.Fatalf("expected id player-1, got %s", result.ID)
	}
}

func TestPlayerService_ListPlayersByGame_Invalid(t *testing.T) {
	service := NewPlayerService(&fakePlayerRepository{})

	_, err := service.ListPlayersByGame(context.Background(), "")
	if !errors.Is(err, ErrMissingGameID) {
		t.Fatalf("expected ErrMissingGameID, got %v", err)
	}
}
