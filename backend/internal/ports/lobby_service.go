package ports

import "context"

type LobbyService interface {
	JoinGame(ctx context.Context, gameCode, userID string, capabilities []string) (PlayerRecord, error)
	ListPlayers(ctx context.Context, gameID string) ([]PlayerRecord, error)
}
