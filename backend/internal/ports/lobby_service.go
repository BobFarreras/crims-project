package ports

import "context"

type LobbyService interface {
	CreateLobby(ctx context.Context, userID string, capabilities []string) (LobbyState, error)
	JoinGame(ctx context.Context, gameCode, userID string, capabilities []string) (PlayerRecord, error)
	ListPlayers(ctx context.Context, gameID string) ([]PlayerRecord, error)
}

type LobbyState struct {
	Game   GameRecord
	Player PlayerRecord
}
