package services

import (
	"context"
	"errors"

	"github.com/digitaistudios/crims-backend/internal/ports"
)

var ErrInvalidEventInput = errors.New("invalid event input")
var ErrMissingEventID = errors.New("missing event id")

type EventService struct {
	repo ports.EventRepository
}

var _ ports.EventService = (*EventService)(nil)

func NewEventService(repo ports.EventRepository) *EventService {
	return &EventService{repo: repo}
}

func (s *EventService) CreateEvent(ctx context.Context, input ports.EventRecordInput) (ports.EventRecord, error) {
	if input.GameID == "" || input.Timestamp == "" || input.LocationID == "" {
		return ports.EventRecord{}, ErrInvalidEventInput
	}

	return s.repo.CreateEvent(ctx, input)
}

func (s *EventService) GetEventByID(ctx context.Context, id string) (ports.EventRecord, error) {
	if id == "" {
		return ports.EventRecord{}, ErrMissingEventID
	}

	return s.repo.GetEventByID(ctx, id)
}

func (s *EventService) ListEventsByGame(ctx context.Context, gameID string) ([]ports.EventRecord, error) {
	if gameID == "" {
		return nil, ErrMissingGameID
	}

	return s.repo.ListEventsByGame(ctx, gameID)
}
