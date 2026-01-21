package services

import (
	"context"
	"errors"
	"testing"

	"github.com/digitaistudios/crims-backend/internal/ports"
)

type fakeTimelineRepository struct {
	createResult ports.TimelineRecord
	createErr    error
	listResult   []ports.TimelineRecord
	listErr      error
}

func (f *fakeTimelineRepository) CreateEntry(ctx context.Context, input ports.TimelineRecordInput) (ports.TimelineRecord, error) {
	if f.createErr != nil {
		return ports.TimelineRecord{}, f.createErr
	}
	return f.createResult, nil
}

func (f *fakeTimelineRepository) GetEntryByID(ctx context.Context, id string) (ports.TimelineRecord, error) {
	return ports.TimelineRecord{}, nil
}

func (f *fakeTimelineRepository) ListEntriesByGame(ctx context.Context, gameID string) ([]ports.TimelineRecord, error) {
	if f.listErr != nil {
		return nil, f.listErr
	}
	return f.listResult, nil
}

func TestTimelineService_CreateEntry_Invalid(t *testing.T) {
	service := NewTimelineService(&fakeTimelineRepository{})

	_, err := service.CreateEntry(context.Background(), ports.TimelineRecordInput{})
	if !errors.Is(err, ErrInvalidTimelineInput) {
		t.Fatalf("expected ErrInvalidTimelineInput, got %v", err)
	}
}

func TestTimelineService_CreateEntry_OK(t *testing.T) {
	repo := &fakeTimelineRepository{createResult: ports.TimelineRecord{ID: "entry-1"}}
	service := NewTimelineService(repo)

	result, err := service.CreateEntry(context.Background(), ports.TimelineRecordInput{
		GameID:      "game-1",
		Timestamp:   "2026-01-21T10:00:00Z",
		Title:       "Arrival",
		Description: "Victim arrives",
	})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if result.ID != "entry-1" {
		t.Fatalf("expected id entry-1, got %s", result.ID)
	}
}

func TestTimelineService_ListEntriesByGame_Invalid(t *testing.T) {
	service := NewTimelineService(&fakeTimelineRepository{})

	_, err := service.ListEntriesByGame(context.Background(), "")
	if !errors.Is(err, ErrMissingGameID) {
		t.Fatalf("expected ErrMissingGameID, got %v", err)
	}
}
