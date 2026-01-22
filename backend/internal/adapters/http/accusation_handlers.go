package http

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/digitaistudios/crims-backend/internal/platform/web"
	"github.com/digitaistudios/crims-backend/internal/ports"
	"github.com/digitaistudios/crims-backend/internal/services"
)

type createAccusationRequest struct {
	GameID     string `json:"gameId"`
	PlayerID   string `json:"playerId"`
	SuspectID  string `json:"suspectId"`
	MotiveID   string `json:"motiveId"`
	EvidenceID string `json:"evidenceId"`
	Verdict    string `json:"verdict"`
}

func NewCreateAccusationHandler(service ports.AccusationService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var payload createAccusationRequest
		if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
			web.RespondError(w, http.StatusBadRequest, "invalid payload", "invalid_payload")
			return
		}

		result, err := service.CreateAccusation(r.Context(), ports.AccusationRecordInput{
			GameID:     payload.GameID,
			PlayerID:   payload.PlayerID,
			SuspectID:  payload.SuspectID,
			MotiveID:   payload.MotiveID,
			EvidenceID: payload.EvidenceID,
			Verdict:    payload.Verdict,
		})
		if err != nil {
			if errors.Is(err, services.ErrInvalidAccusationInput) {
				web.RespondError(w, http.StatusBadRequest, "missing fields", "missing_fields")
				return
			}
			web.RespondError(w, http.StatusInternalServerError, "create failed", "create_failed")
			return
		}

		web.RespondJSON(w, http.StatusCreated, result)
	}
}

func NewListAccusationsByGameHandler(service ports.AccusationService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		gameID, _ := r.Context().Value(idParamKey).(string)
		result, err := service.ListAccusationsByGame(r.Context(), gameID)
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
