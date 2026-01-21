package repo_pb

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/url"

	"github.com/digitaistudios/crims-backend/internal/ports"
)

var ErrRecordNotFound = errors.New("record not found")

type GameRecord = ports.GameRecord
type GameRecordInput = ports.GameRecordInput

type GameRepository struct {
	client *Client
}

var _ ports.GameRepository = (*GameRepository)(nil)

func NewGameRepository(client *Client) *GameRepository {
	return &GameRepository{client: client}
}

func (g *GameRepository) CreateGame(ctx context.Context, input GameRecordInput) (GameRecord, error) {
	payload, err := json.Marshal(map[string]string{
		"code":  input.Code,
		"state": input.State,
		"seed":  input.Seed,
	})
	if err != nil {
		return GameRecord{}, err
	}

	request, err := http.NewRequestWithContext(ctx, http.MethodPost, g.client.baseURL+"/api/collections/games/records", bytes.NewReader(payload))
	if err != nil {
		return GameRecord{}, err
	}
	request.Header.Set("Content-Type", "application/json")

	response, err := g.client.httpClient.Do(request)
	if err != nil {
		return GameRecord{}, err
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK && response.StatusCode != http.StatusCreated {
		return GameRecord{}, fmt.Errorf("%w: status %d", ErrUnexpectedStatus, response.StatusCode)
	}

	var record GameRecord
	if err := json.NewDecoder(response.Body).Decode(&record); err != nil {
		return GameRecord{}, err
	}

	return record, nil
}

func (g *GameRepository) GetGameByID(ctx context.Context, id string) (GameRecord, error) {
	request, err := http.NewRequestWithContext(ctx, http.MethodGet, g.client.baseURL+"/api/collections/games/records/"+id, nil)
	if err != nil {
		return GameRecord{}, err
	}

	response, err := g.client.httpClient.Do(request)
	if err != nil {
		return GameRecord{}, err
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		return GameRecord{}, fmt.Errorf("%w: status %d", ErrUnexpectedStatus, response.StatusCode)
	}

	var record GameRecord
	if err := json.NewDecoder(response.Body).Decode(&record); err != nil {
		return GameRecord{}, err
	}

	return record, nil
}

func (g *GameRepository) GetGameByCode(ctx context.Context, code string) (GameRecord, error) {
	query := url.Values{}
	query.Set("filter", fmt.Sprintf("code='%s'", code))
	url := g.client.baseURL + "/api/collections/games/records?" + query.Encode()

	request, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return GameRecord{}, err
	}

	response, err := g.client.httpClient.Do(request)
	if err != nil {
		return GameRecord{}, err
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		return GameRecord{}, fmt.Errorf("%w: status %d", ErrUnexpectedStatus, response.StatusCode)
	}

	var payload struct {
		Items []GameRecord `json:"items"`
	}
	if err := json.NewDecoder(response.Body).Decode(&payload); err != nil {
		return GameRecord{}, err
	}
	if len(payload.Items) == 0 {
		return GameRecord{}, ErrRecordNotFound
	}

	return payload.Items[0], nil
}
