package services

import (
	"context"
	"errors"

	"github.com/digitaistudios/crims-backend/internal/ports"
)

var ErrInvalidHypothesisInput = errors.New("invalid hypothesis input")
var ErrMissingHypothesisID = errors.New("missing hypothesis id")

type HypothesisService struct {
	repo ports.HypothesisRepository
}

var _ ports.HypothesisService = (*HypothesisService)(nil)

func NewHypothesisService(repo ports.HypothesisRepository) *HypothesisService {
	return &HypothesisService{repo: repo}
}

func (s *HypothesisService) CreateHypothesis(ctx context.Context, input ports.HypothesisRecordInput) (ports.HypothesisRecord, error) {
	if input.GameID == "" || input.Title == "" || input.Status == "" {
		return ports.HypothesisRecord{}, ErrInvalidHypothesisInput
	}

	return s.repo.CreateHypothesis(ctx, input)
}

func (s *HypothesisService) GetHypothesisByID(ctx context.Context, id string) (ports.HypothesisRecord, error) {
	if id == "" {
		return ports.HypothesisRecord{}, ErrMissingHypothesisID
	}

	return s.repo.GetHypothesisByID(ctx, id)
}

func (s *HypothesisService) ListHypothesesByGame(ctx context.Context, gameID string) ([]ports.HypothesisRecord, error) {
	if gameID == "" {
		return nil, ErrMissingGameID
	}

	return s.repo.ListHypothesesByGame(ctx, gameID)
}
