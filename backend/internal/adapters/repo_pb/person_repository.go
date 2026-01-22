package repo_pb

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"

	"github.com/digitaistudios/crims-backend/internal/ports"
)

type PersonRecord = ports.PersonRecord
type PersonRecordInput = ports.PersonRecordInput

type PersonRepository struct {
	client *Client
}

var _ ports.PersonRepository = (*PersonRepository)(nil)

func NewPersonRepository(client *Client) *PersonRepository {
	return &PersonRepository{client: client}
}

func (p *PersonRepository) CreatePerson(ctx context.Context, input PersonRecordInput) (PersonRecord, error) {
	payload, err := json.Marshal(map[string]interface{}{
		"gameId":        input.GameID,
		"name":          input.Name,
		"officialStory": input.OfficialStory,
		"truthStory":    input.TruthStory,
		"stress":        input.Stress,
		"credibility":   input.Credibility,
	})
	if err != nil {
		return PersonRecord{}, err
	}

	request, err := http.NewRequestWithContext(ctx, http.MethodPost, p.client.baseURL+"/api/collections/persons/records", bytes.NewReader(payload))
	if err != nil {
		return PersonRecord{}, err
	}
	request.Header.Set("Content-Type", "application/json")

	response, err := p.client.httpClient.Do(request)
	if err != nil {
		return PersonRecord{}, err
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK && response.StatusCode != http.StatusCreated {
		return PersonRecord{}, fmt.Errorf("%w: status %d", ErrUnexpectedStatus, response.StatusCode)
	}

	var record PersonRecord
	if err := json.NewDecoder(response.Body).Decode(&record); err != nil {
		return PersonRecord{}, err
	}

	return record, nil
}

func (p *PersonRepository) GetPersonByID(ctx context.Context, id string) (PersonRecord, error) {
	request, err := http.NewRequestWithContext(ctx, http.MethodGet, p.client.baseURL+"/api/collections/persons/records/"+id, nil)
	if err != nil {
		return PersonRecord{}, err
	}

	response, err := p.client.httpClient.Do(request)
	if err != nil {
		return PersonRecord{}, err
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		return PersonRecord{}, fmt.Errorf("%w: status %d", ErrUnexpectedStatus, response.StatusCode)
	}

	var record PersonRecord
	if err := json.NewDecoder(response.Body).Decode(&record); err != nil {
		return PersonRecord{}, err
	}

	return record, nil
}

func (p *PersonRepository) ListPersonsByGame(ctx context.Context, gameID string) ([]PersonRecord, error) {
	query := url.Values{}
	query.Set("filter", fmt.Sprintf("gameId='%s'", gameID))
	url := p.client.baseURL + "/api/collections/persons/records?" + query.Encode()

	request, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}

	response, err := p.client.httpClient.Do(request)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("%w: status %d", ErrUnexpectedStatus, response.StatusCode)
	}

	var payload struct {
		Items []PersonRecord `json:"items"`
	}
	if err := json.NewDecoder(response.Body).Decode(&payload); err != nil {
		return nil, err
	}

	return payload.Items, nil
}
