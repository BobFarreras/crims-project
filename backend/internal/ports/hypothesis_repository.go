package ports

import "context"

type HypothesisRepository interface {
	CreateHypothesis(ctx context.Context, input HypothesisRecordInput) (HypothesisRecord, error)
	GetHypothesisByID(ctx context.Context, id string) (HypothesisRecord, error)
	ListHypothesesByGame(ctx context.Context, gameID string) ([]HypothesisRecord, error)
}

type HypothesisRecordInput struct {
	GameID        string
	Title         string
	StrengthScore int
	Status        string
	NodeIDs       []string
}

type HypothesisRecord struct {
	ID            string
	GameID        string
	Title         string
	StrengthScore int
	Status        string
	NodeIDs       []string
}
