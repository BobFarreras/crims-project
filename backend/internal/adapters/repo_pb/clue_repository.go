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

type ClueRecord = ports.ClueRecord
type ClueRecordInput = ports.ClueRecordInput

type ClueRepository struct {
	client *Client
}

var _ ports.ClueRepository = (*ClueRepository)(nil)

func NewClueRepository(client *Client) *ClueRepository {
	return &ClueRepository{client: client}
}

func (c *ClueRepository) CreateClue(ctx context.Context, input ClueRecordInput) (ClueRecord, error) {
	payload, err := json.Marshal(map[string]interface{}{
		"gameId":      input.GameID,
		"type":        input.Type,
		"state":       input.State,
		"reliability": input.Reliability,
		"facts":       input.Facts,
	})
	if err != nil {
		return ClueRecord{}, err
	}

	request, err := http.NewRequestWithContext(ctx, http.MethodPost, c.client.baseURL+"/api/collections/clues/records", bytes.NewReader(payload))
	if err != nil {
		return ClueRecord{}, err
	}
	request.Header.Set("Content-Type", "application/json")

	response, err := c.client.httpClient.Do(request)
	if err != nil {
		return ClueRecord{}, err
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK && response.StatusCode != http.StatusCreated {
		return ClueRecord{}, fmt.Errorf("%w: status %d", ErrUnexpectedStatus, response.StatusCode)
	}

	var record ClueRecord
	if err := json.NewDecoder(response.Body).Decode(&record); err != nil {
		return ClueRecord{}, err
	}

	return record, nil
}

func (c *ClueRepository) GetClueByID(ctx context.Context, id string) (ClueRecord, error) {
	request, err := http.NewRequestWithContext(ctx, http.MethodGet, c.client.baseURL+"/api/collections/clues/records/"+id, nil)
	if err != nil {
		return ClueRecord{}, err
	}

	response, err := c.client.httpClient.Do(request)
	if err != nil {
		return ClueRecord{}, err
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		return ClueRecord{}, fmt.Errorf("%w: status %d", ErrUnexpectedStatus, response.StatusCode)
	}

	var record ClueRecord
	if err := json.NewDecoder(response.Body).Decode(&record); err != nil {
		return ClueRecord{}, err
	}

	return record, nil
}

func (c *ClueRepository) ListCluesByGame(ctx context.Context, gameID string) ([]ClueRecord, error) {
	query := url.Values{}
	query.Set("filter", fmt.Sprintf("gameId='%s'", gameID))
	url := c.client.baseURL + "/api/collections/clues/records?" + query.Encode()

	request, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}

	response, err := c.client.httpClient.Do(request)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("%w: status %d", ErrUnexpectedStatus, response.StatusCode)
	}

	var payload struct {
		Items []ClueRecord `json:"items"`
	}
	if err := json.NewDecoder(response.Body).Decode(&payload); err != nil {
		return nil, err
	}

	return payload.Items, nil
}
