package services

import (
	"context"
	"errors"
	"testing"

	"github.com/digitaistudios/crims-backend/internal/ports"
)

type fakeEventRepository struct {
	createResult ports.EventRecord
	createErr    error
	listResult   []ports.EventRecord
	listErr      error
}

func (f *fakeEventRepository) CreateEvent(ctx context.Context, input ports.EventRecordInput) (ports.EventRecord, error) {
	if f.createErr != nil {
		return ports.EventRecord{}, f.createErr
	}
	return f.createResult, nil
}

func (f *fakeEventRepository) GetEventByID(ctx context.Context, id string) (ports.EventRecord, error) {
	return ports.EventRecord{}, nil
}

func (f *fakeEventRepository) ListEventsByGame(ctx context.Context, gameID string) ([]ports.EventRecord, error) {
	if f.listErr != nil {
		return nil, f.listErr
	}
	return f.listResult, nil
}

func TestEventService_CreateEvent_Invalid(t *testing.T) {
	service := NewEventService(&fakeEventRepository{})

	_, err := service.CreateEvent(context.Background(), ports.EventRecordInput{})
	if !errors.Is(err, ErrInvalidEventInput) {
		t.Fatalf("expected ErrInvalidEventInput, got %v", err)
	}
}

func TestEventService_CreateEvent_OK(t *testing.T) {
	repo := &fakeEventRepository{createResult: ports.EventRecord{ID: "event-1"}}
	service := NewEventService(repo)

	result, err := service.CreateEvent(context.Background(), ports.EventRecordInput{
		GameID:     "game-1",
		Timestamp:  "2026-01-21T10:00:00Z",
		LocationID: "loc-1",
	})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if result.ID != "event-1" {
		t.Fatalf("expected id event-1, got %s", result.ID)
	}
}

func TestEventService_ListEventsByGame_Invalid(t *testing.T) {
	service := NewEventService(&fakeEventRepository{})

	_, err := service.ListEventsByGame(context.Background(), "")
	if !errors.Is(err, ErrMissingGameID) {
		t.Fatalf("expected ErrMissingGameID, got %v", err)
	}
}
