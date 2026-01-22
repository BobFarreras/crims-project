package services

import (
	"context"
	"errors"

	"github.com/digitaistudios/crims-backend/internal/ports"
)

var ErrInvalidPlayerInput = errors.New("invalid player input")

type PlayerService struct {
	repo ports.PlayerRepository
}

var _ ports.PlayerService = (*PlayerService)(nil)

func NewPlayerService(repo ports.PlayerRepository) *PlayerService {
	return &PlayerService{repo: repo}
}

func (s *PlayerService) CreatePlayer(ctx context.Context, input ports.PlayerRecordInput) (ports.PlayerRecord, error) {
	if input.GameID == "" || input.UserID == "" || input.Status == "" {
		return ports.PlayerRecord{}, ErrInvalidPlayerInput
	}

	return s.repo.CreatePlayer(ctx, input)
}

func (s *PlayerService) GetPlayerByID(ctx context.Context, id string) (ports.PlayerRecord, error) {
	if id == "" {
		return ports.PlayerRecord{}, ErrMissingGameID
	}

	return s.repo.GetPlayerByID(ctx, id)
}

func (s *PlayerService) ListPlayersByGame(ctx context.Context, gameID string) ([]ports.PlayerRecord, error) {
	if gameID == "" {
		return nil, ErrMissingGameID
	}

	return s.repo.ListPlayersByGame(ctx, gameID)
}
