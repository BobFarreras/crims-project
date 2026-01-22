package http

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/digitaistudios/crims-backend/internal/platform/web"
	"github.com/digitaistudios/crims-backend/internal/ports"
	"github.com/digitaistudios/crims-backend/internal/services"
)

type createPersonRequest struct {
	GameID        string `json:"gameId"`
	Name          string `json:"name"`
	OfficialStory string `json:"officialStory"`
	TruthStory    string `json:"truthStory"`
	Stress        int    `json:"stress"`
	Credibility   int    `json:"credibility"`
}

func NewCreatePersonHandler(service ports.PersonService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var payload createPersonRequest
		if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
			web.RespondError(w, http.StatusBadRequest, "invalid payload", "invalid_payload")
			return
		}

		result, err := service.CreatePerson(r.Context(), ports.PersonRecordInput{
			GameID:        payload.GameID,
			Name:          payload.Name,
			OfficialStory: payload.OfficialStory,
			TruthStory:    payload.TruthStory,
			Stress:        payload.Stress,
			Credibility:   payload.Credibility,
		})
		if err != nil {
			if errors.Is(err, services.ErrInvalidPersonInput) {
				web.RespondError(w, http.StatusBadRequest, "missing fields", "missing_fields")
				return
			}
			web.RespondError(w, http.StatusInternalServerError, "create failed", "create_failed")
			return
		}

		web.RespondJSON(w, http.StatusCreated, result)
	}
}

func NewListPersonsByGameHandler(service ports.PersonService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		gameID, _ := r.Context().Value(idParamKey).(string)
		result, err := service.ListPersonsByGame(r.Context(), gameID)
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
