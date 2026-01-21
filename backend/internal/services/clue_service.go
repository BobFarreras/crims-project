package services

import (
	"context"
	"errors"

	"github.com/digitaistudios/crims-backend/internal/ports"
)

var ErrInvalidClueInput = errors.New("invalid clue input")

// ClueService orquestra la logica de Clue.
type ClueService struct {
	repo ports.ClueRepository
}

var _ ports.ClueService = (*ClueService)(nil)

func NewClueService(repo ports.ClueRepository) *ClueService {
	return &ClueService{repo: repo}
}

func (s *ClueService) CreateClue(ctx context.Context, input ports.ClueRecordInput) (ports.ClueRecord, error) {
	if input.GameID == "" || input.Type == "" || input.State == "" {
		return ports.ClueRecord{}, ErrInvalidClueInput
	}

	return s.repo.CreateClue(ctx, input)
}

func (s *ClueService) GetClueByID(ctx context.Context, id string) (ports.ClueRecord, error) {
	if id == "" {
		return ports.ClueRecord{}, ErrMissingGameID
	}

	return s.repo.GetClueByID(ctx, id)
}

func (s *ClueService) ListCluesByGame(ctx context.Context, gameID string) ([]ports.ClueRecord, error) {
	if gameID == "" {
		return nil, ErrMissingGameID
	}

	return s.repo.ListCluesByGame(ctx, gameID)
}
