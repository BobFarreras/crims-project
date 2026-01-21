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

type ForensicRecord = ports.ForensicRecord
type ForensicRecordInput = ports.ForensicRecordInput

type ForensicRepository struct {
	client *Client
}

var _ ports.ForensicRepository = (*ForensicRepository)(nil)

func NewForensicRepository(client *Client) *ForensicRepository {
	return &ForensicRepository{client: client}
}

func (f *ForensicRepository) CreateAnalysis(ctx context.Context, input ForensicRecordInput) (ForensicRecord, error) {
	payload, err := json.Marshal(map[string]interface{}{
		"gameId":     input.GameID,
		"clueId":     input.ClueID,
		"result":     input.Result,
		"confidence": input.Confidence,
		"status":     input.Status,
	})
	if err != nil {
		return ForensicRecord{}, err
	}

	request, err := http.NewRequestWithContext(ctx, http.MethodPost, f.client.baseURL+"/api/collections/forensics/records", bytes.NewReader(payload))
	if err != nil {
		return ForensicRecord{}, err
	}
	request.Header.Set("Content-Type", "application/json")

	response, err := f.client.httpClient.Do(request)
	if err != nil {
		return ForensicRecord{}, err
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK && response.StatusCode != http.StatusCreated {
		return ForensicRecord{}, fmt.Errorf("%w: status %d", ErrUnexpectedStatus, response.StatusCode)
	}

	var record ForensicRecord
	if err := json.NewDecoder(response.Body).Decode(&record); err != nil {
		return ForensicRecord{}, err
	}

	return record, nil
}

func (f *ForensicRepository) GetAnalysisByID(ctx context.Context, id string) (ForensicRecord, error) {
	request, err := http.NewRequestWithContext(ctx, http.MethodGet, f.client.baseURL+"/api/collections/forensics/records/"+id, nil)
	if err != nil {
		return ForensicRecord{}, err
	}

	response, err := f.client.httpClient.Do(request)
	if err != nil {
		return ForensicRecord{}, err
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		return ForensicRecord{}, fmt.Errorf("%w: status %d", ErrUnexpectedStatus, response.StatusCode)
	}

	var record ForensicRecord
	if err := json.NewDecoder(response.Body).Decode(&record); err != nil {
		return ForensicRecord{}, err
	}

	return record, nil
}

func (f *ForensicRepository) ListAnalysesByGame(ctx context.Context, gameID string) ([]ForensicRecord, error) {
	query := url.Values{}
	query.Set("filter", fmt.Sprintf("gameId='%s'", gameID))
	url := f.client.baseURL + "/api/collections/forensics/records?" + query.Encode()

	request, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}

	response, err := f.client.httpClient.Do(request)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("%w: status %d", ErrUnexpectedStatus, response.StatusCode)
	}

	var payload struct {
		Items []ForensicRecord `json:"items"`
	}
	if err := json.NewDecoder(response.Body).Decode(&payload); err != nil {
		return nil, err
	}

	return payload.Items, nil
}
