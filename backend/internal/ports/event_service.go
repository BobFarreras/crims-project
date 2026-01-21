package ports

import "context"

type EventService interface {
	CreateEvent(ctx context.Context, input EventRecordInput) (EventRecord, error)
	GetEventByID(ctx context.Context, id string) (EventRecord, error)
	ListEventsByGame(ctx context.Context, gameID string) ([]EventRecord, error)
}
