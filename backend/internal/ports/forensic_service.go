package ports

import "context"

type ForensicService interface {
	CreateAnalysis(ctx context.Context, input ForensicRecordInput) (ForensicRecord, error)
	GetAnalysisByID(ctx context.Context, id string) (ForensicRecord, error)
	ListAnalysesByGame(ctx context.Context, gameID string) ([]ForensicRecord, error)
}
