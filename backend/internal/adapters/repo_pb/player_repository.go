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

type PlayerRecord = ports.PlayerRecord
type PlayerRecordInput = ports.PlayerRecordInput

type PlayerRepository struct {
	client *Client
}

var _ ports.PlayerRepository = (*PlayerRepository)(nil)

func NewPlayerRepository(client *Client) *PlayerRepository {
	return &PlayerRepository{client: client}
}

func (p *PlayerRepository) CreatePlayer(ctx context.Context, input PlayerRecordInput) (PlayerRecord, error) {
	payload, err := json.Marshal(map[string]interface{}{
		"gameId":       input.GameID,
		"userId":       input.UserID,
		"capabilities": input.Capabilities,
		"status":       input.Status,
		"isHost":       input.IsHost,
	})
	if err != nil {
		return PlayerRecord{}, err
	}

	request, err := http.NewRequestWithContext(ctx, http.MethodPost, p.client.baseURL+"/api/collections/players/records", bytes.NewReader(payload))
	if err != nil {
		return PlayerRecord{}, err
	}
	request.Header.Set("Content-Type", "application/json")

	response, err := p.client.httpClient.Do(request)
	if err != nil {
		return PlayerRecord{}, err
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK && response.StatusCode != http.StatusCreated {
		return PlayerRecord{}, fmt.Errorf("%w: status %d", ErrUnexpectedStatus, response.StatusCode)
	}

	var record PlayerRecord
	if err := json.NewDecoder(response.Body).Decode(&record); err != nil {
		return PlayerRecord{}, err
	}

	return record, nil
}

func (p *PlayerRepository) GetPlayerByID(ctx context.Context, id string) (PlayerRecord, error) {
	request, err := http.NewRequestWithContext(ctx, http.MethodGet, p.client.baseURL+"/api/collections/players/records/"+id, nil)
	if err != nil {
		return PlayerRecord{}, err
	}

	response, err := p.client.httpClient.Do(request)
	if err != nil {
		return PlayerRecord{}, err
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		return PlayerRecord{}, fmt.Errorf("%w: status %d", ErrUnexpectedStatus, response.StatusCode)
	}

	var record PlayerRecord
	if err := json.NewDecoder(response.Body).Decode(&record); err != nil {
		return PlayerRecord{}, err
	}

	return record, nil
}

func (p *PlayerRepository) ListPlayersByGame(ctx context.Context, gameID string) ([]PlayerRecord, error) {
	query := url.Values{}
	query.Set("filter", fmt.Sprintf("gameId='%s'", gameID))
	url := p.client.baseURL + "/api/collections/players/records?" + query.Encode()

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
		Items []PlayerRecord `json:"items"`
	}
	if err := json.NewDecoder(response.Body).Decode(&payload); err != nil {
		return nil, err
	}

	return payload.Items, nil
}
