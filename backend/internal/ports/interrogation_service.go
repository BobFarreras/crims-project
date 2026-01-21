package ports

import "context"

type InterrogationService interface {
	CreateInterrogation(ctx context.Context, input InterrogationRecordInput) (InterrogationRecord, error)
	GetInterrogationByID(ctx context.Context, id string) (InterrogationRecord, error)
	ListInterrogationsByGame(ctx context.Context, gameID string) ([]InterrogationRecord, error)
}
