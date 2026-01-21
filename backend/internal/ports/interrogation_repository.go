package ports

import "context"

type InterrogationRepository interface {
	CreateInterrogation(ctx context.Context, input InterrogationRecordInput) (InterrogationRecord, error)
	GetInterrogationByID(ctx context.Context, id string) (InterrogationRecord, error)
	ListInterrogationsByGame(ctx context.Context, gameID string) ([]InterrogationRecord, error)
}

type InterrogationRecordInput struct {
	GameID   string
	PersonID string
	Question string
	Answer   string
	Tone     string
}

type InterrogationRecord struct {
	ID       string
	GameID   string
	PersonID string
	Question string
	Answer   string
	Tone     string
}
