package http

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/digitaistudios/crims-backend/internal/adapters/repo_pb"
	"github.com/digitaistudios/crims-backend/internal/platform/web"
	"github.com/digitaistudios/crims-backend/internal/ports"
	"github.com/digitaistudios/crims-backend/internal/services"
)

type contextKey string

const (
	idParamKey   contextKey = "gameID"
	codeParamKey contextKey = "gameCode"
)

type createGameRequest struct {
	Code  string `json:"code"`
	State string `json:"state"`
	Seed  string `json:"seed"`
}

func NewCreateGameHandler(service ports.GameService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var payload createGameRequest
		if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
			web.RespondError(w, http.StatusBadRequest, "invalid payload", "invalid_payload")
			return
		}

		result, err := service.CreateGame(r.Context(), ports.GameRecordInput{
			Code:  payload.Code,
			State: payload.State,
			Seed:  payload.Seed,
		})
		if err != nil {
			if errors.Is(err, services.ErrInvalidGameInput) {
				web.RespondError(w, http.StatusBadRequest, "missing fields", "missing_fields")
				return
			}
			web.RespondError(w, http.StatusInternalServerError, "create failed", "create_failed")
			return
		}

		web.RespondJSON(w, http.StatusCreated, result)
	}
}

func NewGetGameByIDHandler(service ports.GameService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id, _ := r.Context().Value(idParamKey).(string)

		result, err := service.GetGameByID(r.Context(), id)
		if err != nil {
			if errors.Is(err, services.ErrMissingGameID) {
				web.RespondError(w, http.StatusBadRequest, "missing id", "missing_id")
				return
			}
			if errors.Is(err, repo_pb.ErrRecordNotFound) {
				web.RespondError(w, http.StatusNotFound, "not found", "game_not_found")
				return
			}
			web.RespondError(w, http.StatusInternalServerError, "not found", "game_not_found")
			return
		}

		web.RespondJSON(w, http.StatusOK, result)
	}
}

func NewGetGameByCodeHandler(service ports.GameService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		code, _ := r.Context().Value(codeParamKey).(string)

		result, err := service.GetGameByCode(r.Context(), code)
		if err != nil {
			if errors.Is(err, services.ErrMissingGameCode) {
				web.RespondError(w, http.StatusBadRequest, "missing code", "missing_code")
				return
			}
			if errors.Is(err, repo_pb.ErrRecordNotFound) {
				web.RespondError(w, http.StatusNotFound, "not found", "game_not_found")
				return
			}
			web.RespondError(w, http.StatusInternalServerError, "not found", "game_not_found")
			return
		}

		web.RespondJSON(w, http.StatusOK, result)
	}
}
