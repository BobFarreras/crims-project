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

type HypothesisRecord = ports.HypothesisRecord
type HypothesisRecordInput = ports.HypothesisRecordInput

type HypothesisRepository struct {
	client *Client
}

var _ ports.HypothesisRepository = (*HypothesisRepository)(nil)

func NewHypothesisRepository(client *Client) *HypothesisRepository {
	return &HypothesisRepository{client: client}
}

func (h *HypothesisRepository) CreateHypothesis(ctx context.Context, input HypothesisRecordInput) (HypothesisRecord, error) {
	payload, err := json.Marshal(map[string]interface{}{
		"gameId":        input.GameID,
		"title":         input.Title,
		"strengthScore": input.StrengthScore,
		"status":        input.Status,
		"nodeIds":       input.NodeIDs,
	})
	if err != nil {
		return HypothesisRecord{}, err
	}

	request, err := http.NewRequestWithContext(ctx, http.MethodPost, h.client.baseURL+"/api/collections/hypotheses/records", bytes.NewReader(payload))
	if err != nil {
		return HypothesisRecord{}, err
	}
	request.Header.Set("Content-Type", "application/json")

	response, err := h.client.httpClient.Do(request)
	if err != nil {
		return HypothesisRecord{}, err
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK && response.StatusCode != http.StatusCreated {
		return HypothesisRecord{}, fmt.Errorf("%w: status %d", ErrUnexpectedStatus, response.StatusCode)
	}

	var record HypothesisRecord
	if err := json.NewDecoder(response.Body).Decode(&record); err != nil {
		return HypothesisRecord{}, err
	}

	return record, nil
}

func (h *HypothesisRepository) GetHypothesisByID(ctx context.Context, id string) (HypothesisRecord, error) {
	request, err := http.NewRequestWithContext(ctx, http.MethodGet, h.client.baseURL+"/api/collections/hypotheses/records/"+id, nil)
	if err != nil {
		return HypothesisRecord{}, err
	}

	response, err := h.client.httpClient.Do(request)
	if err != nil {
		return HypothesisRecord{}, err
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		return HypothesisRecord{}, fmt.Errorf("%w: status %d", ErrUnexpectedStatus, response.StatusCode)
	}

	var record HypothesisRecord
	if err := json.NewDecoder(response.Body).Decode(&record); err != nil {
		return HypothesisRecord{}, err
	}

	return record, nil
}

func (h *HypothesisRepository) ListHypothesesByGame(ctx context.Context, gameID string) ([]HypothesisRecord, error) {
	query := url.Values{}
	query.Set("filter", fmt.Sprintf("gameId='%s'", gameID))
	url := h.client.baseURL + "/api/collections/hypotheses/records?" + query.Encode()

	request, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}

	response, err := h.client.httpClient.Do(request)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("%w: status %d", ErrUnexpectedStatus, response.StatusCode)
	}

	var payload struct {
		Items []HypothesisRecord `json:"items"`
	}
	if err := json.NewDecoder(response.Body).Decode(&payload); err != nil {
		return nil, err
	}

	return payload.Items, nil
}
