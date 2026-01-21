package ports

import "context"

type HypothesisService interface {
	CreateHypothesis(ctx context.Context, input HypothesisRecordInput) (HypothesisRecord, error)
	GetHypothesisByID(ctx context.Context, id string) (HypothesisRecord, error)
	ListHypothesesByGame(ctx context.Context, gameID string) ([]HypothesisRecord, error)
}
