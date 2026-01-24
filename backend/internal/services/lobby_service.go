package services

import (
	"context"
	"crypto/rand"
	"errors"
	"math/big"
	"strings"

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

func (s *LobbyService) JoinGame(ctx context.Context, gameCode, userID string, capabilities []string) (ports.PlayerRecord, error) {
	if gameCode == "" || userID == "" {
		return ports.PlayerRecord{}, ErrInvalidLobbyInput
	}

	game, err := s.gameRepo.GetGameByCode(ctx, gameCode)
	if err != nil {
		return ports.PlayerRecord{}, err
	}

	return s.playerRepo.CreatePlayer(ctx, ports.PlayerRecordInput{
		GameID:       game.ID,
		UserID:       userID,
		Capabilities: capabilities,
		Status:       "ONLINE",
		IsHost:       false,
	})
}

func (s *LobbyService) CreateLobby(ctx context.Context, userID string, capabilities []string) (ports.LobbyState, error) {
	if userID == "" {
		return ports.LobbyState{}, ErrInvalidLobbyInput
	}

	code := generateLobbyCode()
	game, err := s.gameRepo.CreateGame(ctx, ports.GameRecordInput{
		Code:  code,
		State: "LOBBY",
		Seed:  "seed",
	})
	if err != nil {
		return ports.LobbyState{}, err
	}

	player, err := s.playerRepo.CreatePlayer(ctx, ports.PlayerRecordInput{
		GameID:       game.ID,
		UserID:       userID,
		Capabilities: capabilities,
		Status:       "ONLINE",
		IsHost:       true,
	})
	if err != nil {
		return ports.LobbyState{}, err
	}

	return ports.LobbyState{Game: game, Player: player}, nil
}

func generateLobbyCode() string {
	const letters = "ABCDEFGHJKLMNPQRSTUVWXYZ"
	var builder strings.Builder
	for i := 0; i < 4; i++ {
		n, _ := rand.Int(rand.Reader, big.NewInt(int64(len(letters))))
		builder.WriteByte(letters[n.Int64()])
	}
	return builder.String()
}

func (s *LobbyService) ListPlayers(ctx context.Context, gameID string) ([]ports.PlayerRecord, error) {
	if gameID == "" {
		return nil, ErrMissingGameID
	}

	return s.playerRepo.ListPlayersByGame(ctx, gameID)
}
