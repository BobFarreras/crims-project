package http

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/digitaistudios/crims-backend/internal/platform/web"
	"github.com/digitaistudios/crims-backend/internal/ports"
	"github.com/digitaistudios/crims-backend/internal/services"
)

type createHypothesisRequest struct {
	GameID        string   `json:"gameId"`
	Title         string   `json:"title"`
	StrengthScore int      `json:"strengthScore"`
	Status        string   `json:"status"`
	NodeIDs       []string `json:"nodeIds"`
}

func NewCreateHypothesisHandler(service ports.HypothesisService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var payload createHypothesisRequest
		if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
			web.RespondError(w, http.StatusBadRequest, "invalid payload", "invalid_payload")
			return
		}

		result, err := service.CreateHypothesis(r.Context(), ports.HypothesisRecordInput{
			GameID:        payload.GameID,
			Title:         payload.Title,
			StrengthScore: payload.StrengthScore,
			Status:        payload.Status,
			NodeIDs:       payload.NodeIDs,
		})
		if err != nil {
			if errors.Is(err, services.ErrInvalidHypothesisInput) {
				web.RespondError(w, http.StatusBadRequest, "missing fields", "missing_fields")
				return
			}
			web.RespondError(w, http.StatusInternalServerError, "create failed", "create_failed")
			return
		}

		web.RespondJSON(w, http.StatusCreated, result)
	}
}

func NewListHypothesesByGameHandler(service ports.HypothesisService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		gameID, _ := r.Context().Value(idParamKey).(string)
		result, err := service.ListHypothesesByGame(r.Context(), gameID)
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
