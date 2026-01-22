package ports

import "context"

type TimelineRepository interface {
	CreateEntry(ctx context.Context, input TimelineRecordInput) (TimelineRecord, error)
	GetEntryByID(ctx context.Context, id string) (TimelineRecord, error)
	ListEntriesByGame(ctx context.Context, gameID string) ([]TimelineRecord, error)
}

type TimelineRecordInput struct {
	GameID      string
	Timestamp   string
	Title       string
	Description string
	EventID     string
}

type TimelineRecord struct {
	ID          string
	GameID      string
	Timestamp   string
	Title       string
	Description string
	EventID     string
}
