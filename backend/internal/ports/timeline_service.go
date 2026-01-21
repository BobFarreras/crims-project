package ports

import "context"

type TimelineService interface {
	CreateEntry(ctx context.Context, input TimelineRecordInput) (TimelineRecord, error)
	GetEntryByID(ctx context.Context, id string) (TimelineRecord, error)
	ListEntriesByGame(ctx context.Context, gameID string) ([]TimelineRecord, error)
}
