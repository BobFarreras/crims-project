package ports

import "context"

type ClueService interface {
	CreateClue(ctx context.Context, input ClueRecordInput) (ClueRecord, error)
	GetClueByID(ctx context.Context, id string) (ClueRecord, error)
	ListCluesByGame(ctx context.Context, gameID string) ([]ClueRecord, error)
}
