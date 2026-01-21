package services

import (
	"context"
	"errors"

	"github.com/digitaistudios/crims-backend/internal/ports"
)

var ErrInvalidForensicInput = errors.New("invalid forensic input")

// ForensicService gestiona analisis forense.
type ForensicService struct {
	repo ports.ForensicRepository
}

var _ ports.ForensicService = (*ForensicService)(nil)

func NewForensicService(repo ports.ForensicRepository) *ForensicService {
	return &ForensicService{repo: repo}
}

func (s *ForensicService) CreateAnalysis(ctx context.Context, input ports.ForensicRecordInput) (ports.ForensicRecord, error) {
	if input.GameID == "" || input.ClueID == "" || input.Result == "" || input.Status == "" {
		return ports.ForensicRecord{}, ErrInvalidForensicInput
	}

	return s.repo.CreateAnalysis(ctx, input)
}

func (s *ForensicService) GetAnalysisByID(ctx context.Context, id string) (ports.ForensicRecord, error) {
	if id == "" {
		return ports.ForensicRecord{}, ErrMissingGameID
	}

	return s.repo.GetAnalysisByID(ctx, id)
}

func (s *ForensicService) ListAnalysesByGame(ctx context.Context, gameID string) ([]ports.ForensicRecord, error) {
	if gameID == "" {
		return nil, ErrMissingGameID
	}

	return s.repo.ListAnalysesByGame(ctx, gameID)
}
