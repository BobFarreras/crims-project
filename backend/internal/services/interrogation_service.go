package services

import (
	"context"
	"errors"

	"github.com/digitaistudios/crims-backend/internal/ports"
)

var ErrInvalidInterrogationInput = errors.New("invalid interrogation input")

type InterrogationService struct {
	repo ports.InterrogationRepository
}

var _ ports.InterrogationService = (*InterrogationService)(nil)

func NewInterrogationService(repo ports.InterrogationRepository) *InterrogationService {
	return &InterrogationService{repo: repo}
}

func (s *InterrogationService) CreateInterrogation(ctx context.Context, input ports.InterrogationRecordInput) (ports.InterrogationRecord, error) {
	if input.GameID == "" || input.PersonID == "" || input.Question == "" || input.Answer == "" {
		return ports.InterrogationRecord{}, ErrInvalidInterrogationInput
	}

	return s.repo.CreateInterrogation(ctx, input)
}

func (s *InterrogationService) GetInterrogationByID(ctx context.Context, id string) (ports.InterrogationRecord, error) {
	if id == "" {
		return ports.InterrogationRecord{}, ErrMissingGameID
	}

	return s.repo.GetInterrogationByID(ctx, id)
}

func (s *InterrogationService) ListInterrogationsByGame(ctx context.Context, gameID string) ([]ports.InterrogationRecord, error) {
	if gameID == "" {
		return nil, ErrMissingGameID
	}

	return s.repo.ListInterrogationsByGame(ctx, gameID)
}
