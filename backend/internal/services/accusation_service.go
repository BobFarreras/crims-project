package services

import (
	"context"
	"errors"

	"github.com/digitaistudios/crims-backend/internal/ports"
)

var ErrInvalidAccusationInput = errors.New("invalid accusation input")
var ErrMissingAccusationID = errors.New("missing accusation id")

type AccusationService struct {
	repo ports.AccusationRepository
}

var _ ports.AccusationService = (*AccusationService)(nil)

func NewAccusationService(repo ports.AccusationRepository) *AccusationService {
	return &AccusationService{repo: repo}
}

func (s *AccusationService) CreateAccusation(ctx context.Context, input ports.AccusationRecordInput) (ports.AccusationRecord, error) {
	if input.GameID == "" || input.PlayerID == "" || input.SuspectID == "" || input.MotiveID == "" || input.EvidenceID == "" {
		return ports.AccusationRecord{}, ErrInvalidAccusationInput
	}

	return s.repo.CreateAccusation(ctx, input)
}

func (s *AccusationService) GetAccusationByID(ctx context.Context, id string) (ports.AccusationRecord, error) {
	if id == "" {
		return ports.AccusationRecord{}, ErrMissingAccusationID
	}

	return s.repo.GetAccusationByID(ctx, id)
}

func (s *AccusationService) ListAccusationsByGame(ctx context.Context, gameID string) ([]ports.AccusationRecord, error) {
	if gameID == "" {
		return nil, ErrMissingGameID
	}

	return s.repo.ListAccusationsByGame(ctx, gameID)
}
