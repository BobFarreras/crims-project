package ports

import "context"

// GameRepository defineix el contracte per guardar/carregar partides.
// Això permetrà canviar PocketBase per Postgres sense trencar res.
type GameRepository interface {
	// Exemple de mètode (encara no implementat)
	SaveGame(ctx context.Context, gameID string, data any) error
	GetGame(ctx context.Context, gameID string) (any, error)
}
