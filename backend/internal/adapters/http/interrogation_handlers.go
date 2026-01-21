package http

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/digitaistudios/crims-backend/internal/platform/web"
	"github.com/digitaistudios/crims-backend/internal/ports"
	"github.com/digitaistudios/crims-backend/internal/services"
)

type createInterrogationRequest struct {
	GameID   string `json:"gameId"`
	PersonID string `json:"personId"`
	Question string `json:"question"`
	Answer   string `json:"answer"`
	Tone     string `json:"tone"`
}

func NewCreateInterrogationHandler(service ports.InterrogationService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var payload createInterrogationRequest
		if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
			web.RespondError(w, http.StatusBadRequest, "invalid payload", "invalid_payload")
			return
		}

		result, err := service.CreateInterrogation(r.Context(), ports.InterrogationRecordInput{
			GameID:   payload.GameID,
			PersonID: payload.PersonID,
			Question: payload.Question,
			Answer:   payload.Answer,
			Tone:     payload.Tone,
		})
		if err != nil {
			if errors.Is(err, services.ErrInvalidInterrogationInput) {
				web.RespondError(w, http.StatusBadRequest, "missing fields", "missing_fields")
				return
			}
			web.RespondError(w, http.StatusInternalServerError, "create failed", "create_failed")
			return
		}

		web.RespondJSON(w, http.StatusCreated, result)
	}
}

func NewListInterrogationsByGameHandler(service ports.InterrogationService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		gameID, _ := r.Context().Value(idParamKey).(string)
		result, err := service.ListInterrogationsByGame(r.Context(), gameID)
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
