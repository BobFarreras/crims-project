package ports

import "context"

type EventRepository interface {
	CreateEvent(ctx context.Context, input EventRecordInput) (EventRecord, error)
	GetEventByID(ctx context.Context, id string) (EventRecord, error)
	ListEventsByGame(ctx context.Context, gameID string) ([]EventRecord, error)
}

type EventRecordInput struct {
	GameID       string
	Timestamp    string
	LocationID   string
	Participants []string
}

type EventRecord struct {
	ID           string
	GameID       string
	Timestamp    string
	LocationID   string
	Participants []string
}
