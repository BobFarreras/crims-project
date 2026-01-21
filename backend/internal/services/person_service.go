package services

import (
	"context"
	"errors"

	"github.com/digitaistudios/crims-backend/internal/ports"
)

var ErrInvalidPersonInput = errors.New("invalid person input")

type PersonService struct {
	repo ports.PersonRepository
}

var _ ports.PersonService = (*PersonService)(nil)

func NewPersonService(repo ports.PersonRepository) *PersonService {
	return &PersonService{repo: repo}
}

func (s *PersonService) CreatePerson(ctx context.Context, input ports.PersonRecordInput) (ports.PersonRecord, error) {
	if input.GameID == "" || input.Name == "" || input.OfficialStory == "" || input.TruthStory == "" {
		return ports.PersonRecord{}, ErrInvalidPersonInput
	}

	return s.repo.CreatePerson(ctx, input)
}

func (s *PersonService) GetPersonByID(ctx context.Context, id string) (ports.PersonRecord, error) {
	if id == "" {
		return ports.PersonRecord{}, ErrMissingGameID
	}

	return s.repo.GetPersonByID(ctx, id)
}

func (s *PersonService) ListPersonsByGame(ctx context.Context, gameID string) ([]ports.PersonRecord, error) {
	if gameID == "" {
		return nil, ErrMissingGameID
	}

	return s.repo.ListPersonsByGame(ctx, gameID)
}
