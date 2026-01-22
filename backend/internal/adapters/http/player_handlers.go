package http

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/digitaistudios/crims-backend/internal/platform/web"
	"github.com/digitaistudios/crims-backend/internal/ports"
	"github.com/digitaistudios/crims-backend/internal/services"
)

type createPlayerRequest struct {
	GameID       string   `json:"gameId"`
	UserID       string   `json:"userId"`
	Capabilities []string `json:"capabilities"`
	Status       string   `json:"status"`
	IsHost       bool     `json:"isHost"`
}

func NewCreatePlayerHandler(service ports.PlayerService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var payload createPlayerRequest
		if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
			web.RespondError(w, http.StatusBadRequest, "invalid payload", "invalid_payload")
			return
		}

		result, err := service.CreatePlayer(r.Context(), ports.PlayerRecordInput{
			GameID:       payload.GameID,
			UserID:       payload.UserID,
			Capabilities: payload.Capabilities,
			Status:       payload.Status,
			IsHost:       payload.IsHost,
		})
		if err != nil {
			if errors.Is(err, services.ErrInvalidPlayerInput) {
				web.RespondError(w, http.StatusBadRequest, "missing fields", "missing_fields")
				return
			}
			web.RespondError(w, http.StatusInternalServerError, "create failed", "create_failed")
			return
		}

		web.RespondJSON(w, http.StatusCreated, result)
	}
}

func NewListPlayersByGameHandler(service ports.PlayerService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		gameID, _ := r.Context().Value(idParamKey).(string)
		result, err := service.ListPlayersByGame(r.Context(), gameID)
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
