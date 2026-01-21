package ports

import "context"

// GameService defineix la logica d'aplicacio per Game.
type GameService interface {
	CreateGame(ctx context.Context, input GameRecordInput) (GameRecord, error)
	GetGameByID(ctx context.Context, id string) (GameRecord, error)
	GetGameByCode(ctx context.Context, code string) (GameRecord, error)
}
