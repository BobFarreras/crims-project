package services

import (
	"context"
	"errors"

	"github.com/digitaistudios/crims-backend/internal/ports"
)

var ErrInvalidLobbyInput = errors.New("invalid lobby input")

type LobbyService struct {
	gameRepo   ports.GameRepository
	playerRepo ports.PlayerRepository
}

var _ ports.LobbyService = (*LobbyService)(nil)

func NewLobbyService(gameRepo ports.GameRepository, playerRepo ports.PlayerRepository) *LobbyService {
	return &LobbyService{gameRepo: gameRepo, playerRepo: playerRepo}
}

func (s *LobbyService) JoinGame(ctx context.Context, gameCode, userID, role string) (ports.PlayerRecord, error) {
	if gameCode == "" || userID == "" || role == "" {
		return ports.PlayerRecord{}, ErrInvalidLobbyInput
	}

	game, err := s.gameRepo.GetGameByCode(ctx, gameCode)
	if err != nil {
		return ports.PlayerRecord{}, err
	}

	return s.playerRepo.CreatePlayer(ctx, ports.PlayerRecordInput{
		GameID: game.ID,
		UserID: userID,
		Role:   role,
		Status: "ONLINE",
		IsHost: false,
	})
}

func (s *LobbyService) ListPlayers(ctx context.Context, gameID string) ([]ports.PlayerRecord, error) {
	if gameID == "" {
		return nil, ErrMissingGameID
	}

	return s.playerRepo.ListPlayersByGame(ctx, gameID)
}
