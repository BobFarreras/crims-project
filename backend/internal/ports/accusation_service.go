package ports

import "context"

type AccusationService interface {
	CreateAccusation(ctx context.Context, input AccusationRecordInput) (AccusationRecord, error)
	GetAccusationByID(ctx context.Context, id string) (AccusationRecord, error)
	ListAccusationsByGame(ctx context.Context, gameID string) ([]AccusationRecord, error)
}
