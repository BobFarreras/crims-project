package services

import (
	"context"
	"errors"
	"testing"

	"github.com/digitaistudios/crims-backend/internal/ports"
)

type fakeInterrogationRepository struct {
	createResult ports.InterrogationRecord
	createErr    error
	listResult   []ports.InterrogationRecord
	listErr      error
}

func (f *fakeInterrogationRepository) CreateInterrogation(ctx context.Context, input ports.InterrogationRecordInput) (ports.InterrogationRecord, error) {
	if f.createErr != nil {
		return ports.InterrogationRecord{}, f.createErr
	}
	return f.createResult, nil
}

func (f *fakeInterrogationRepository) GetInterrogationByID(ctx context.Context, id string) (ports.InterrogationRecord, error) {
	return ports.InterrogationRecord{}, nil
}

func (f *fakeInterrogationRepository) ListInterrogationsByGame(ctx context.Context, gameID string) ([]ports.InterrogationRecord, error) {
	if f.listErr != nil {
		return nil, f.listErr
	}
	return f.listResult, nil
}

func TestInterrogationService_CreateInterrogation_Invalid(t *testing.T) {
	service := NewInterrogationService(&fakeInterrogationRepository{})

	_, err := service.CreateInterrogation(context.Background(), ports.InterrogationRecordInput{})
	if !errors.Is(err, ErrInvalidInterrogationInput) {
		t.Fatalf("expected ErrInvalidInterrogationInput, got %v", err)
	}
}

func TestInterrogationService_CreateInterrogation_OK(t *testing.T) {
	repo := &fakeInterrogationRepository{createResult: ports.InterrogationRecord{ID: "int-1"}}
	service := NewInterrogationService(repo)

	result, err := service.CreateInterrogation(context.Background(), ports.InterrogationRecordInput{
		GameID:   "game-1",
		PersonID: "person-1",
		Question: "Where were you?",
		Answer:   "At home.",
		Tone:     "neutral",
	})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if result.ID != "int-1" {
		t.Fatalf("expected id int-1, got %s", result.ID)
	}
}

func TestInterrogationService_ListInterrogationsByGame_Invalid(t *testing.T) {
	service := NewInterrogationService(&fakeInterrogationRepository{})

	_, err := service.ListInterrogationsByGame(context.Background(), "")
	if !errors.Is(err, ErrMissingGameID) {
		t.Fatalf("expected ErrMissingGameID, got %v", err)
	}
}
