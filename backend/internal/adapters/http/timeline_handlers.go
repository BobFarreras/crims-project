package http

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/digitaistudios/crims-backend/internal/platform/web"
	"github.com/digitaistudios/crims-backend/internal/ports"
	"github.com/digitaistudios/crims-backend/internal/services"
)

type createTimelineRequest struct {
	GameID      string `json:"gameId"`
	Timestamp   string `json:"timestamp"`
	Title       string `json:"title"`
	Description string `json:"description"`
	EventID     string `json:"eventId"`
}

func NewCreateTimelineHandler(service ports.TimelineService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var payload createTimelineRequest
		if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
			web.RespondError(w, http.StatusBadRequest, "invalid payload", "invalid_payload")
			return
		}

		result, err := service.CreateEntry(r.Context(), ports.TimelineRecordInput{
			GameID:      payload.GameID,
			Timestamp:   payload.Timestamp,
			Title:       payload.Title,
			Description: payload.Description,
			EventID:     payload.EventID,
		})
		if err != nil {
			if errors.Is(err, services.ErrInvalidTimelineInput) {
				web.RespondError(w, http.StatusBadRequest, "missing fields", "missing_fields")
				return
			}
			web.RespondError(w, http.StatusInternalServerError, "create failed", "create_failed")
			return
		}

		web.RespondJSON(w, http.StatusCreated, result)
	}
}

func NewListTimelineByGameHandler(service ports.TimelineService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		gameID, _ := r.Context().Value(idParamKey).(string)
		result, err := service.ListEntriesByGame(r.Context(), gameID)
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
