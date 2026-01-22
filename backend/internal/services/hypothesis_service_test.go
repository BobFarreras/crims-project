package services

import (
	"context"
	"errors"
	"testing"

	"github.com/digitaistudios/crims-backend/internal/ports"
)

type fakeHypothesisRepository struct {
	createResult ports.HypothesisRecord
	createErr    error
	listResult   []ports.HypothesisRecord
	listErr      error
}

func (f *fakeHypothesisRepository) CreateHypothesis(ctx context.Context, input ports.HypothesisRecordInput) (ports.HypothesisRecord, error) {
	if f.createErr != nil {
		return ports.HypothesisRecord{}, f.createErr
	}
	return f.createResult, nil
}

func (f *fakeHypothesisRepository) GetHypothesisByID(ctx context.Context, id string) (ports.HypothesisRecord, error) {
	return ports.HypothesisRecord{}, nil
}

func (f *fakeHypothesisRepository) ListHypothesesByGame(ctx context.Context, gameID string) ([]ports.HypothesisRecord, error) {
	if f.listErr != nil {
		return nil, f.listErr
	}
	return f.listResult, nil
}

func TestHypothesisService_CreateHypothesis_Invalid(t *testing.T) {
	service := NewHypothesisService(&fakeHypothesisRepository{})

	_, err := service.CreateHypothesis(context.Background(), ports.HypothesisRecordInput{})
	if !errors.Is(err, ErrInvalidHypothesisInput) {
		t.Fatalf("expected ErrInvalidHypothesisInput, got %v", err)
	}
}

func TestHypothesisService_CreateHypothesis_OK(t *testing.T) {
	repo := &fakeHypothesisRepository{createResult: ports.HypothesisRecord{ID: "hyp-1"}}
	service := NewHypothesisService(repo)

	result, err := service.CreateHypothesis(context.Background(), ports.HypothesisRecordInput{
		GameID:        "game-1",
		Title:         "Main theory",
		StrengthScore: 42,
		Status:        "PLAUSIBLE",
	})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if result.ID != "hyp-1" {
		t.Fatalf("expected id hyp-1, got %s", result.ID)
	}
}

func TestHypothesisService_ListHypothesesByGame_Invalid(t *testing.T) {
	service := NewHypothesisService(&fakeHypothesisRepository{})

	_, err := service.ListHypothesesByGame(context.Background(), "")
	if !errors.Is(err, ErrMissingGameID) {
		t.Fatalf("expected ErrMissingGameID, got %v", err)
	}
}
