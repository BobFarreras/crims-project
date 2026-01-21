package http

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/digitaistudios/crims-backend/internal/platform/web"
	"github.com/digitaistudios/crims-backend/internal/ports"
	"github.com/digitaistudios/crims-backend/internal/services"
)

type createEventRequest struct {
	GameID       string   `json:"gameId"`
	Timestamp    string   `json:"timestamp"`
	LocationID   string   `json:"locationId"`
	Participants []string `json:"participants"`
}

func NewCreateEventHandler(service ports.EventService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var payload createEventRequest
		if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
			web.RespondError(w, http.StatusBadRequest, "invalid payload", "invalid_payload")
			return
		}

		result, err := service.CreateEvent(r.Context(), ports.EventRecordInput{
			GameID:       payload.GameID,
			Timestamp:    payload.Timestamp,
			LocationID:   payload.LocationID,
			Participants: payload.Participants,
		})
		if err != nil {
			if errors.Is(err, services.ErrInvalidEventInput) {
				web.RespondError(w, http.StatusBadRequest, "missing fields", "missing_fields")
				return
			}
			web.RespondError(w, http.StatusInternalServerError, "create failed", "create_failed")
			return
		}

		web.RespondJSON(w, http.StatusCreated, result)
	}
}

func NewListEventsByGameHandler(service ports.EventService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		gameID, _ := r.Context().Value(idParamKey).(string)
		result, err := service.ListEventsByGame(r.Context(), gameID)
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
