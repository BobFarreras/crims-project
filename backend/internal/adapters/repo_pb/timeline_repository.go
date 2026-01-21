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

type TimelineRecord = ports.TimelineRecord
type TimelineRecordInput = ports.TimelineRecordInput

type TimelineRepository struct {
	client *Client
}

var _ ports.TimelineRepository = (*TimelineRepository)(nil)

func NewTimelineRepository(client *Client) *TimelineRepository {
	return &TimelineRepository{client: client}
}

func (t *TimelineRepository) CreateEntry(ctx context.Context, input TimelineRecordInput) (TimelineRecord, error) {
	payload, err := json.Marshal(map[string]interface{}{
		"gameId":      input.GameID,
		"timestamp":   input.Timestamp,
		"title":       input.Title,
		"description": input.Description,
		"eventId":     input.EventID,
	})
	if err != nil {
		return TimelineRecord{}, err
	}

	request, err := http.NewRequestWithContext(ctx, http.MethodPost, t.client.baseURL+"/api/collections/timeline/records", bytes.NewReader(payload))
	if err != nil {
		return TimelineRecord{}, err
	}
	request.Header.Set("Content-Type", "application/json")

	response, err := t.client.httpClient.Do(request)
	if err != nil {
		return TimelineRecord{}, err
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK && response.StatusCode != http.StatusCreated {
		return TimelineRecord{}, fmt.Errorf("%w: status %d", ErrUnexpectedStatus, response.StatusCode)
	}

	var record TimelineRecord
	if err := json.NewDecoder(response.Body).Decode(&record); err != nil {
		return TimelineRecord{}, err
	}

	return record, nil
}

func (t *TimelineRepository) GetEntryByID(ctx context.Context, id string) (TimelineRecord, error) {
	request, err := http.NewRequestWithContext(ctx, http.MethodGet, t.client.baseURL+"/api/collections/timeline/records/"+id, nil)
	if err != nil {
		return TimelineRecord{}, err
	}

	response, err := t.client.httpClient.Do(request)
	if err != nil {
		return TimelineRecord{}, err
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		return TimelineRecord{}, fmt.Errorf("%w: status %d", ErrUnexpectedStatus, response.StatusCode)
	}

	var record TimelineRecord
	if err := json.NewDecoder(response.Body).Decode(&record); err != nil {
		return TimelineRecord{}, err
	}

	return record, nil
}

func (t *TimelineRepository) ListEntriesByGame(ctx context.Context, gameID string) ([]TimelineRecord, error) {
	query := url.Values{}
	query.Set("filter", fmt.Sprintf("gameId='%s'", gameID))
	url := t.client.baseURL + "/api/collections/timeline/records?" + query.Encode()

	request, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}

	response, err := t.client.httpClient.Do(request)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("%w: status %d", ErrUnexpectedStatus, response.StatusCode)
	}

	var payload struct {
		Items []TimelineRecord `json:"items"`
	}
	if err := json.NewDecoder(response.Body).Decode(&payload); err != nil {
		return nil, err
	}

	return payload.Items, nil
}
