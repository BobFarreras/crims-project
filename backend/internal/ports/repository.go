package ports

import "context"

// GameRepository defineix el contracte per guardar/carregar partides.
// Això permetrà canviar PocketBase per Postgres sense trencar res.
type GameRepository interface {
	CreateGame(ctx context.Context, input GameRecordInput) (GameRecord, error)
	GetGameByID(ctx context.Context, id string) (GameRecord, error)
	GetGameByCode(ctx context.Context, code string) (GameRecord, error)
}

type GameRecordInput struct {
	Code  string
	State string
	Seed  string
}

type GameRecord struct {
	ID    string
	Code  string
	State string
	Seed  string
}
