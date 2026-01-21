package services

import (
	"context"
	"errors"

	"github.com/digitaistudios/crims-backend/internal/ports"
)

var ErrInvalidTimelineInput = errors.New("invalid timeline input")

type TimelineService struct {
	repo ports.TimelineRepository
}

var _ ports.TimelineService = (*TimelineService)(nil)

func NewTimelineService(repo ports.TimelineRepository) *TimelineService {
	return &TimelineService{repo: repo}
}

func (s *TimelineService) CreateEntry(ctx context.Context, input ports.TimelineRecordInput) (ports.TimelineRecord, error) {
	if input.GameID == "" || input.Timestamp == "" || input.Title == "" {
		return ports.TimelineRecord{}, ErrInvalidTimelineInput
	}

	return s.repo.CreateEntry(ctx, input)
}

func (s *TimelineService) GetEntryByID(ctx context.Context, id string) (ports.TimelineRecord, error) {
	if id == "" {
		return ports.TimelineRecord{}, ErrMissingGameID
	}

	return s.repo.GetEntryByID(ctx, id)
}

func (s *TimelineService) ListEntriesByGame(ctx context.Context, gameID string) ([]ports.TimelineRecord, error) {
	if gameID == "" {
		return nil, ErrMissingGameID
	}

	return s.repo.ListEntriesByGame(ctx, gameID)
}
