package ports

import "context"

type AccusationRepository interface {
	CreateAccusation(ctx context.Context, input AccusationRecordInput) (AccusationRecord, error)
	GetAccusationByID(ctx context.Context, id string) (AccusationRecord, error)
	ListAccusationsByGame(ctx context.Context, gameID string) ([]AccusationRecord, error)
}

type AccusationRecordInput struct {
	GameID     string
	PlayerID   string
	SuspectID  string
	MotiveID   string
	EvidenceID string
	Verdict    string
}

type AccusationRecord struct {
	ID         string
	GameID     string
	PlayerID   string
	SuspectID  string
	MotiveID   string
	EvidenceID string
	Verdict    string
}
