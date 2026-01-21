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

type AccusationRecord = ports.AccusationRecord
type AccusationRecordInput = ports.AccusationRecordInput

type AccusationRepository struct {
	client *Client
}

var _ ports.AccusationRepository = (*AccusationRepository)(nil)

func NewAccusationRepository(client *Client) *AccusationRepository {
	return &AccusationRepository{client: client}
}

func (a *AccusationRepository) CreateAccusation(ctx context.Context, input AccusationRecordInput) (AccusationRecord, error) {
	payload, err := json.Marshal(map[string]interface{}{
		"gameId":     input.GameID,
		"playerId":   input.PlayerID,
		"suspectId":  input.SuspectID,
		"motiveId":   input.MotiveID,
		"evidenceId": input.EvidenceID,
		"verdict":    input.Verdict,
	})
	if err != nil {
		return AccusationRecord{}, err
	}

	request, err := http.NewRequestWithContext(ctx, http.MethodPost, a.client.baseURL+"/api/collections/accusations/records", bytes.NewReader(payload))
	if err != nil {
		return AccusationRecord{}, err
	}
	request.Header.Set("Content-Type", "application/json")

	response, err := a.client.httpClient.Do(request)
	if err != nil {
		return AccusationRecord{}, err
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK && response.StatusCode != http.StatusCreated {
		return AccusationRecord{}, fmt.Errorf("%w: status %d", ErrUnexpectedStatus, response.StatusCode)
	}

	var record AccusationRecord
	if err := json.NewDecoder(response.Body).Decode(&record); err != nil {
		return AccusationRecord{}, err
	}

	return record, nil
}

func (a *AccusationRepository) GetAccusationByID(ctx context.Context, id string) (AccusationRecord, error) {
	request, err := http.NewRequestWithContext(ctx, http.MethodGet, a.client.baseURL+"/api/collections/accusations/records/"+id, nil)
	if err != nil {
		return AccusationRecord{}, err
	}

	response, err := a.client.httpClient.Do(request)
	if err != nil {
		return AccusationRecord{}, err
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		return AccusationRecord{}, fmt.Errorf("%w: status %d", ErrUnexpectedStatus, response.StatusCode)
	}

	var record AccusationRecord
	if err := json.NewDecoder(response.Body).Decode(&record); err != nil {
		return AccusationRecord{}, err
	}

	return record, nil
}

func (a *AccusationRepository) ListAccusationsByGame(ctx context.Context, gameID string) ([]AccusationRecord, error) {
	query := url.Values{}
	query.Set("filter", fmt.Sprintf("gameId='%s'", gameID))
	url := a.client.baseURL + "/api/collections/accusations/records?" + query.Encode()

	request, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}

	response, err := a.client.httpClient.Do(request)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("%w: status %d", ErrUnexpectedStatus, response.StatusCode)
	}

	var payload struct {
		Items []AccusationRecord `json:"items"`
	}
	if err := json.NewDecoder(response.Body).Decode(&payload); err != nil {
		return nil, err
	}

	return payload.Items, nil
}
