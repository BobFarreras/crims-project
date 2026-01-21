package services

import (
	"context"
	"errors"
	"testing"

	"github.com/digitaistudios/crims-backend/internal/ports"
)

type fakeClueRepository struct {
	createResult ports.ClueRecord
	createErr    error
	listResult   []ports.ClueRecord
	listErr      error
}

func (f *fakeClueRepository) CreateClue(ctx context.Context, input ports.ClueRecordInput) (ports.ClueRecord, error) {
	if f.createErr != nil {
		return ports.ClueRecord{}, f.createErr
	}
	return f.createResult, nil
}

func (f *fakeClueRepository) GetClueByID(ctx context.Context, id string) (ports.ClueRecord, error) {
	return ports.ClueRecord{}, nil
}

func (f *fakeClueRepository) ListCluesByGame(ctx context.Context, gameID string) ([]ports.ClueRecord, error) {
	if f.listErr != nil {
		return nil, f.listErr
	}
	return f.listResult, nil
}

func TestClueService_CreateClue_Invalid(t *testing.T) {
	service := NewClueService(&fakeClueRepository{})

	_, err := service.CreateClue(context.Background(), ports.ClueRecordInput{})
	if !errors.Is(err, ErrInvalidClueInput) {
		t.Fatalf("expected ErrInvalidClueInput, got %v", err)
	}
}

func TestClueService_CreateClue_OK(t *testing.T) {
	repo := &fakeClueRepository{createResult: ports.ClueRecord{ID: "clue-1"}}
	service := NewClueService(repo)

	result, err := service.CreateClue(context.Background(), ports.ClueRecordInput{
		GameID:      "game-1",
		Type:        "OBJECT",
		State:       "DISCOVERED",
		Reliability: 80,
	})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if result.ID != "clue-1" {
		t.Fatalf("expected id clue-1, got %s", result.ID)
	}
}

func TestClueService_ListCluesByGame_Invalid(t *testing.T) {
	service := NewClueService(&fakeClueRepository{})

	_, err := service.ListCluesByGame(context.Background(), "")
	if !errors.Is(err, ErrMissingGameID) {
		t.Fatalf("expected ErrMissingGameID, got %v", err)
	}
}
