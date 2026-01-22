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

type InterrogationRecord = ports.InterrogationRecord
type InterrogationRecordInput = ports.InterrogationRecordInput

type InterrogationRepository struct {
	client *Client
}

var _ ports.InterrogationRepository = (*InterrogationRepository)(nil)

func NewInterrogationRepository(client *Client) *InterrogationRepository {
	return &InterrogationRepository{client: client}
}

func (i *InterrogationRepository) CreateInterrogation(ctx context.Context, input InterrogationRecordInput) (InterrogationRecord, error) {
	payload, err := json.Marshal(map[string]interface{}{
		"gameId":   input.GameID,
		"personId": input.PersonID,
		"question": input.Question,
		"answer":   input.Answer,
		"tone":     input.Tone,
	})
	if err != nil {
		return InterrogationRecord{}, err
	}

	request, err := http.NewRequestWithContext(ctx, http.MethodPost, i.client.baseURL+"/api/collections/interrogations/records", bytes.NewReader(payload))
	if err != nil {
		return InterrogationRecord{}, err
	}
	request.Header.Set("Content-Type", "application/json")

	response, err := i.client.httpClient.Do(request)
	if err != nil {
		return InterrogationRecord{}, err
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK && response.StatusCode != http.StatusCreated {
		return InterrogationRecord{}, fmt.Errorf("%w: status %d", ErrUnexpectedStatus, response.StatusCode)
	}

	var record InterrogationRecord
	if err := json.NewDecoder(response.Body).Decode(&record); err != nil {
		return InterrogationRecord{}, err
	}

	return record, nil
}

func (i *InterrogationRepository) GetInterrogationByID(ctx context.Context, id string) (InterrogationRecord, error) {
	request, err := http.NewRequestWithContext(ctx, http.MethodGet, i.client.baseURL+"/api/collections/interrogations/records/"+id, nil)
	if err != nil {
		return InterrogationRecord{}, err
	}

	response, err := i.client.httpClient.Do(request)
	if err != nil {
		return InterrogationRecord{}, err
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		return InterrogationRecord{}, fmt.Errorf("%w: status %d", ErrUnexpectedStatus, response.StatusCode)
	}

	var record InterrogationRecord
	if err := json.NewDecoder(response.Body).Decode(&record); err != nil {
		return InterrogationRecord{}, err
	}

	return record, nil
}

func (i *InterrogationRepository) ListInterrogationsByGame(ctx context.Context, gameID string) ([]InterrogationRecord, error) {
	query := url.Values{}
	query.Set("filter", fmt.Sprintf("gameId='%s'", gameID))
	url := i.client.baseURL + "/api/collections/interrogations/records?" + query.Encode()

	request, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}

	response, err := i.client.httpClient.Do(request)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("%w: status %d", ErrUnexpectedStatus, response.StatusCode)
	}

	var payload struct {
		Items []InterrogationRecord `json:"items"`
	}
	if err := json.NewDecoder(response.Body).Decode(&payload); err != nil {
		return nil, err
	}

	return payload.Items, nil
}
