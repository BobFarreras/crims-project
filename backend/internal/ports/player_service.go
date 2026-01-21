package ports

import "context"

type PlayerService interface {
	CreatePlayer(ctx context.Context, input PlayerRecordInput) (PlayerRecord, error)
	GetPlayerByID(ctx context.Context, id string) (PlayerRecord, error)
	ListPlayersByGame(ctx context.Context, gameID string) ([]PlayerRecord, error)
}
