package services

import (
	"context"
	"errors"
	"testing"

	"github.com/digitaistudios/crims-backend/internal/ports"
)

type fakeLobbyGameRepo struct {
	getByCodeResult ports.GameRecord
	getByCodeErr    error
}

func (f *fakeLobbyGameRepo) CreateGame(ctx context.Context, input ports.GameRecordInput) (ports.GameRecord, error) {
	return ports.GameRecord{}, nil
}

func (f *fakeLobbyGameRepo) GetGameByID(ctx context.Context, id string) (ports.GameRecord, error) {
	return ports.GameRecord{}, nil
}

func (f *fakeLobbyGameRepo) GetGameByCode(ctx context.Context, code string) (ports.GameRecord, error) {
	if f.getByCodeErr != nil {
		return ports.GameRecord{}, f.getByCodeErr
	}
	return f.getByCodeResult, nil
}

type fakeLobbyPlayerRepo struct {
	createResult ports.PlayerRecord
	createErr    error
}

func (f *fakeLobbyPlayerRepo) CreatePlayer(ctx context.Context, input ports.PlayerRecordInput) (ports.PlayerRecord, error) {
	if f.createErr != nil {
		return ports.PlayerRecord{}, f.createErr
	}
	return f.createResult, nil
}

func (f *fakeLobbyPlayerRepo) GetPlayerByID(ctx context.Context, id string) (ports.PlayerRecord, error) {
	return ports.PlayerRecord{}, nil
}

func (f *fakeLobbyPlayerRepo) ListPlayersByGame(ctx context.Context, gameID string) ([]ports.PlayerRecord, error) {
	return nil, nil
}

func TestLobbyService_JoinGame_Invalid(t *testing.T) {
	service := NewLobbyService(&fakeLobbyGameRepo{}, &fakeLobbyPlayerRepo{})

	_, err := service.JoinGame(context.Background(), "", "user-1", "DETECTIVE")
	if !errors.Is(err, ErrInvalidLobbyInput) {
		t.Fatalf("expected ErrInvalidLobbyInput, got %v", err)
	}
}

func TestLobbyService_JoinGame_OK(t *testing.T) {
	gameRepo := &fakeLobbyGameRepo{getByCodeResult: ports.GameRecord{ID: "game-1", Code: "ABCD"}}
	playerRepo := &fakeLobbyPlayerRepo{createResult: ports.PlayerRecord{ID: "player-1"}}
	service := NewLobbyService(gameRepo, playerRepo)

	result, err := service.JoinGame(context.Background(), "ABCD", "user-1", "DETECTIVE")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if result.ID != "player-1" {
		t.Fatalf("expected id player-1, got %s", result.ID)
	}
}
