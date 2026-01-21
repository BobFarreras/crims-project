package ports

import "context"

type ClueRepository interface {
	CreateClue(ctx context.Context, input ClueRecordInput) (ClueRecord, error)
	GetClueByID(ctx context.Context, id string) (ClueRecord, error)
	ListCluesByGame(ctx context.Context, gameID string) ([]ClueRecord, error)
}

type ClueRecordInput struct {
	GameID      string
	Type        string
	State       string
	Reliability int
	Facts       map[string]interface{}
}

type ClueRecord struct {
	ID          string
	GameID      string
	Type        string
	State       string
	Reliability int
	Facts       map[string]interface{}
}
