package http

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/digitaistudios/crims-backend/internal/platform/web"
	"github.com/digitaistudios/crims-backend/internal/ports"
	"github.com/digitaistudios/crims-backend/internal/services"
)

type lobbyJoinRequest struct {
	GameCode string `json:"gameCode"`
	UserID   string `json:"userId"`
	Role     string `json:"role"`
}

func NewLobbyJoinHandler(service ports.LobbyService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var payload lobbyJoinRequest
		if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
			web.RespondError(w, http.StatusBadRequest, "invalid payload", "invalid_payload")
			return
		}

		result, err := service.JoinGame(r.Context(), payload.GameCode, payload.UserID, payload.Role)
		if err != nil {
			if errors.Is(err, services.ErrInvalidLobbyInput) {
				web.RespondError(w, http.StatusBadRequest, "missing fields", "missing_fields")
				return
			}
			web.RespondError(w, http.StatusInternalServerError, "join failed", "join_failed")
			return
		}

		web.RespondJSON(w, http.StatusCreated, result)
	}
}
