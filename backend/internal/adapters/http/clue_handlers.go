package http

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/digitaistudios/crims-backend/internal/platform/web"
	"github.com/digitaistudios/crims-backend/internal/ports"
	"github.com/digitaistudios/crims-backend/internal/services"
)

type createClueRequest struct {
	GameID      string                 `json:"gameId"`
	Type        string                 `json:"type"`
	State       string                 `json:"state"`
	Reliability int                    `json:"reliability"`
	Facts       map[string]interface{} `json:"facts"`
}

func NewCreateClueHandler(service ports.ClueService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var payload createClueRequest
		if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
			web.RespondError(w, http.StatusBadRequest, "invalid payload", "invalid_payload")
			return
		}

		result, err := service.CreateClue(r.Context(), ports.ClueRecordInput{
			GameID:      payload.GameID,
			Type:        payload.Type,
			State:       payload.State,
			Reliability: payload.Reliability,
			Facts:       payload.Facts,
		})
		if err != nil {
			if errors.Is(err, services.ErrInvalidClueInput) {
				web.RespondError(w, http.StatusBadRequest, "missing fields", "missing_fields")
				return
			}
			web.RespondError(w, http.StatusInternalServerError, "create failed", "create_failed")
			return
		}

		web.RespondJSON(w, http.StatusCreated, result)
	}
}

func NewListCluesByGameHandler(service ports.ClueService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		gameID, _ := r.Context().Value(idParamKey).(string)
		result, err := service.ListCluesByGame(r.Context(), gameID)
		if err != nil {
			if errors.Is(err, services.ErrMissingGameID) {
				web.RespondError(w, http.StatusBadRequest, "missing game id", "missing_game_id")
				return
			}
			web.RespondError(w, http.StatusInternalServerError, "list failed", "list_failed")
			return
		}

		web.RespondJSON(w, http.StatusOK, result)
	}
}
