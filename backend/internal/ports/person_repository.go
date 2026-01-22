package ports

import "context"

type PersonRepository interface {
	CreatePerson(ctx context.Context, input PersonRecordInput) (PersonRecord, error)
	GetPersonByID(ctx context.Context, id string) (PersonRecord, error)
	ListPersonsByGame(ctx context.Context, gameID string) ([]PersonRecord, error)
}

type PersonRecordInput struct {
	GameID        string
	Name          string
	OfficialStory string
	TruthStory    string
	Stress        int
	Credibility   int
}

type PersonRecord struct {
	ID            string
	GameID        string
	Name          string
	OfficialStory string
	TruthStory    string
	Stress        int
	Credibility   int
}
