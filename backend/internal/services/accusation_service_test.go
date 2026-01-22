package services

import (
	"context"
	"errors"
	"testing"

	"github.com/digitaistudios/crims-backend/internal/ports"
)

type fakeAccusationRepository struct {
	createResult ports.AccusationRecord
	createErr    error
	listResult   []ports.AccusationRecord
	listErr      error
}

func (f *fakeAccusationRepository) CreateAccusation(ctx context.Context, input ports.AccusationRecordInput) (ports.AccusationRecord, error) {
	if f.createErr != nil {
		return ports.AccusationRecord{}, f.createErr
	}
	return f.createResult, nil
}

func (f *fakeAccusationRepository) GetAccusationByID(ctx context.Context, id string) (ports.AccusationRecord, error) {
	return ports.AccusationRecord{}, nil
}

func (f *fakeAccusationRepository) ListAccusationsByGame(ctx context.Context, gameID string) ([]ports.AccusationRecord, error) {
	if f.listErr != nil {
		return nil, f.listErr
	}
	return f.listResult, nil
}

func TestAccusationService_CreateAccusation_Invalid(t *testing.T) {
	service := NewAccusationService(&fakeAccusationRepository{})

	_, err := service.CreateAccusation(context.Background(), ports.AccusationRecordInput{})
	if !errors.Is(err, ErrInvalidAccusationInput) {
		t.Fatalf("expected ErrInvalidAccusationInput, got %v", err)
	}
}

func TestAccusationService_CreateAccusation_OK(t *testing.T) {
	repo := &fakeAccusationRepository{createResult: ports.AccusationRecord{ID: "acc-1"}}
	service := NewAccusationService(repo)

	result, err := service.CreateAccusation(context.Background(), ports.AccusationRecordInput{
		GameID:     "game-1",
		PlayerID:   "player-1",
		SuspectID:  "person-1",
		MotiveID:   "motive-1",
		EvidenceID: "clue-1",
	})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if result.ID != "acc-1" {
		t.Fatalf("expected id acc-1, got %s", result.ID)
	}
}

func TestAccusationService_ListAccusationsByGame_Invalid(t *testing.T) {
	service := NewAccusationService(&fakeAccusationRepository{})

	_, err := service.ListAccusationsByGame(context.Background(), "")
	if !errors.Is(err, ErrMissingGameID) {
		t.Fatalf("expected ErrMissingGameID, got %v", err)
	}
}
