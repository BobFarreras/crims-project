package services

import (
	"context"
	"errors"
	"testing"

	"github.com/digitaistudios/crims-backend/internal/ports"
)

type fakePersonRepository struct {
	createResult ports.PersonRecord
	createErr    error
	listResult   []ports.PersonRecord
	listErr      error
}

func (f *fakePersonRepository) CreatePerson(ctx context.Context, input ports.PersonRecordInput) (ports.PersonRecord, error) {
	if f.createErr != nil {
		return ports.PersonRecord{}, f.createErr
	}
	return f.createResult, nil
}

func (f *fakePersonRepository) GetPersonByID(ctx context.Context, id string) (ports.PersonRecord, error) {
	return ports.PersonRecord{}, nil
}

func (f *fakePersonRepository) ListPersonsByGame(ctx context.Context, gameID string) ([]ports.PersonRecord, error) {
	if f.listErr != nil {
		return nil, f.listErr
	}
	return f.listResult, nil
}

func TestPersonService_CreatePerson_Invalid(t *testing.T) {
	service := NewPersonService(&fakePersonRepository{})

	_, err := service.CreatePerson(context.Background(), ports.PersonRecordInput{})
	if !errors.Is(err, ErrInvalidPersonInput) {
		t.Fatalf("expected ErrInvalidPersonInput, got %v", err)
	}
}

func TestPersonService_CreatePerson_OK(t *testing.T) {
	repo := &fakePersonRepository{createResult: ports.PersonRecord{ID: "person-1"}}
	service := NewPersonService(repo)

	result, err := service.CreatePerson(context.Background(), ports.PersonRecordInput{
		GameID:        "game-1",
		Name:          "Witness",
		OfficialStory: "story",
		TruthStory:    "truth",
		Stress:        10,
		Credibility:   80,
	})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if result.ID != "person-1" {
		t.Fatalf("expected id person-1, got %s", result.ID)
	}
}

func TestPersonService_ListPersonsByGame_Invalid(t *testing.T) {
	service := NewPersonService(&fakePersonRepository{})

	_, err := service.ListPersonsByGame(context.Background(), "")
	if !errors.Is(err, ErrMissingGameID) {
		t.Fatalf("expected ErrMissingGameID, got %v", err)
	}
}
