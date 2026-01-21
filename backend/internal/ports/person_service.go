package ports

import "context"

type PersonService interface {
	CreatePerson(ctx context.Context, input PersonRecordInput) (PersonRecord, error)
	GetPersonByID(ctx context.Context, id string) (PersonRecord, error)
	ListPersonsByGame(ctx context.Context, gameID string) ([]PersonRecord, error)
}
