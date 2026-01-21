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

type EventRecord = ports.EventRecord
type EventRecordInput = ports.EventRecordInput

type EventRepository struct {
	client *Client
}

var _ ports.EventRepository = (*EventRepository)(nil)

func NewEventRepository(client *Client) *EventRepository {
	return &EventRepository{client: client}
}

func (e *EventRepository) CreateEvent(ctx context.Context, input EventRecordInput) (EventRecord, error) {
	payload, err := json.Marshal(map[string]interface{}{
		"gameId":       input.GameID,
		"timestamp":    input.Timestamp,
		"locationId":   input.LocationID,
		"participants": input.Participants,
	})
	if err != nil {
		return EventRecord{}, err
	}

	request, err := http.NewRequestWithContext(ctx, http.MethodPost, e.client.baseURL+"/api/collections/events/records", bytes.NewReader(payload))
	if err != nil {
		return EventRecord{}, err
	}
	request.Header.Set("Content-Type", "application/json")

	response, err := e.client.httpClient.Do(request)
	if err != nil {
		return EventRecord{}, err
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK && response.StatusCode != http.StatusCreated {
		return EventRecord{}, fmt.Errorf("%w: status %d", ErrUnexpectedStatus, response.StatusCode)
	}

	var record EventRecord
	if err := json.NewDecoder(response.Body).Decode(&record); err != nil {
		return EventRecord{}, err
	}

	return record, nil
}

func (e *EventRepository) GetEventByID(ctx context.Context, id string) (EventRecord, error) {
	request, err := http.NewRequestWithContext(ctx, http.MethodGet, e.client.baseURL+"/api/collections/events/records/"+id, nil)
	if err != nil {
		return EventRecord{}, err
	}

	response, err := e.client.httpClient.Do(request)
	if err != nil {
		return EventRecord{}, err
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		return EventRecord{}, fmt.Errorf("%w: status %d", ErrUnexpectedStatus, response.StatusCode)
	}

	var record EventRecord
	if err := json.NewDecoder(response.Body).Decode(&record); err != nil {
		return EventRecord{}, err
	}

	return record, nil
}

func (e *EventRepository) ListEventsByGame(ctx context.Context, gameID string) ([]EventRecord, error) {
	query := url.Values{}
	query.Set("filter", fmt.Sprintf("gameId='%s'", gameID))
	url := e.client.baseURL + "/api/collections/events/records?" + query.Encode()

	request, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}

	response, err := e.client.httpClient.Do(request)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("%w: status %d", ErrUnexpectedStatus, response.StatusCode)
	}

	var payload struct {
		Items []EventRecord `json:"items"`
	}
	if err := json.NewDecoder(response.Body).Decode(&payload); err != nil {
		return nil, err
	}

	return payload.Items, nil
}
