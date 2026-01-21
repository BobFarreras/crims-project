package ports

import "context"

type PlayerRepository interface {
	CreatePlayer(ctx context.Context, input PlayerRecordInput) (PlayerRecord, error)
	GetPlayerByID(ctx context.Context, id string) (PlayerRecord, error)
	ListPlayersByGame(ctx context.Context, gameID string) ([]PlayerRecord, error)
}

type PlayerRecordInput struct {
	GameID string
	UserID string
	Role   string
	Status string
	IsHost bool
}

type PlayerRecord struct {
	ID     string
	GameID string
	UserID string
	Role   string
	Status string
	IsHost bool
}
