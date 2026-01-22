package http

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/digitaistudios/crims-backend/internal/platform/web"
	"github.com/digitaistudios/crims-backend/internal/ports"
	"github.com/digitaistudios/crims-backend/internal/services"
)

type createForensicRequest struct {
	GameID     string `json:"gameId"`
	ClueID     string `json:"clueId"`
	Result     string `json:"result"`
	Confidence int    `json:"confidence"`
	Status     string `json:"status"`
}

func NewCreateForensicHandler(service ports.ForensicService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var payload createForensicRequest
		if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
			web.RespondError(w, http.StatusBadRequest, "invalid payload", "invalid_payload")
			return
		}

		result, err := service.CreateAnalysis(r.Context(), ports.ForensicRecordInput{
			GameID:     payload.GameID,
			ClueID:     payload.ClueID,
			Result:     payload.Result,
			Confidence: payload.Confidence,
			Status:     payload.Status,
		})
		if err != nil {
			if errors.Is(err, services.ErrInvalidForensicInput) {
				web.RespondError(w, http.StatusBadRequest, "missing fields", "missing_fields")
				return
			}
			web.RespondError(w, http.StatusInternalServerError, "create failed", "create_failed")
			return
		}

		web.RespondJSON(w, http.StatusCreated, result)
	}
}

func NewListForensicsByGameHandler(service ports.ForensicService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		gameID, _ := r.Context().Value(idParamKey).(string)
		result, err := service.ListAnalysesByGame(r.Context(), gameID)
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
