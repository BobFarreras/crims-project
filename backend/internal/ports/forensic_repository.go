package ports

import "context"

type ForensicRepository interface {
	CreateAnalysis(ctx context.Context, input ForensicRecordInput) (ForensicRecord, error)
	GetAnalysisByID(ctx context.Context, id string) (ForensicRecord, error)
	ListAnalysesByGame(ctx context.Context, gameID string) ([]ForensicRecord, error)
}

type ForensicRecordInput struct {
	GameID     string
	ClueID     string
	Result     string
	Confidence int
	Status     string
}

type ForensicRecord struct {
	ID         string
	GameID     string
	ClueID     string
	Result     string
	Confidence int
	Status     string
}
