package services

import (
	"context"
	"errors"
	"testing"

	"github.com/digitaistudios/crims-backend/internal/ports"
)

type fakeForensicRepository struct {
	createResult ports.ForensicRecord
	createErr    error
	listResult   []ports.ForensicRecord
	listErr      error
}

func (f *fakeForensicRepository) CreateAnalysis(ctx context.Context, input ports.ForensicRecordInput) (ports.ForensicRecord, error) {
	if f.createErr != nil {
		return ports.ForensicRecord{}, f.createErr
	}
	return f.createResult, nil
}

func (f *fakeForensicRepository) GetAnalysisByID(ctx context.Context, id string) (ports.ForensicRecord, error) {
	return ports.ForensicRecord{}, nil
}

func (f *fakeForensicRepository) ListAnalysesByGame(ctx context.Context, gameID string) ([]ports.ForensicRecord, error) {
	if f.listErr != nil {
		return nil, f.listErr
	}
	return f.listResult, nil
}

func TestForensicService_CreateAnalysis_Invalid(t *testing.T) {
	service := NewForensicService(&fakeForensicRepository{})

	_, err := service.CreateAnalysis(context.Background(), ports.ForensicRecordInput{})
	if !errors.Is(err, ErrInvalidForensicInput) {
		t.Fatalf("expected ErrInvalidForensicInput, got %v", err)
	}
}

func TestForensicService_CreateAnalysis_OK(t *testing.T) {
	repo := &fakeForensicRepository{createResult: ports.ForensicRecord{ID: "forensic-1"}}
	service := NewForensicService(repo)

	result, err := service.CreateAnalysis(context.Background(), ports.ForensicRecordInput{
		GameID:     "game-1",
		ClueID:     "clue-1",
		Result:     "match",
		Confidence: 90,
		Status:     "DONE",
	})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if result.ID != "forensic-1" {
		t.Fatalf("expected id forensic-1, got %s", result.ID)
	}
}

func TestForensicService_ListAnalysesByGame_Invalid(t *testing.T) {
	service := NewForensicService(&fakeForensicRepository{})

	_, err := service.ListAnalysesByGame(context.Background(), "")
	if !errors.Is(err, ErrMissingGameID) {
		t.Fatalf("expected ErrMissingGameID, got %v", err)
	}
}
