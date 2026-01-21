package services

import (
	"context"
	"errors"

	"github.com/digitaistudios/crims-backend/internal/ports"
)

var (
	ErrInvalidGameInput = errors.New("invalid game input")
	ErrMissingGameID    = errors.New("missing game id")
	ErrMissingGameCode  = errors.New("missing game code")
)

type GameService struct {
	repo ports.GameRepository
}

func NewGameService(repo ports.GameRepository) *GameService {
	return &GameService{repo: repo}
}

func (s *GameService) CreateGame(ctx context.Context, input ports.GameRecordInput) (ports.GameRecord, error) {
	if input.Code == "" || input.State == "" || input.Seed == "" {
		return ports.GameRecord{}, ErrInvalidGameInput
	}

	return s.repo.CreateGame(ctx, input)
}

func (s *GameService) GetGameByID(ctx context.Context, id string) (ports.GameRecord, error) {
	if id == "" {
		return ports.GameRecord{}, ErrMissingGameID
	}

	return s.repo.GetGameByID(ctx, id)
}

func (s *GameService) GetGameByCode(ctx context.Context, code string) (ports.GameRecord, error) {
	if code == "" {
		return ports.GameRecord{}, ErrMissingGameCode
	}

	return s.repo.GetGameByCode(ctx, code)
}
