package http

import (
	"encoding/json"
	"net/http"

	"github.com/digitaistudios/crims-backend/internal/adapters/repo_pb"
	"github.com/digitaistudios/crims-backend/internal/platform/web"
	"github.com/digitaistudios/crims-backend/internal/ports"
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

func NewCreateGameHandler(repo ports.GameRepository) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var payload createGameRequest
		if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
			web.RespondError(w, http.StatusBadRequest, "invalid payload", "invalid_payload")
			return
		}
		if payload.Code == "" || payload.State == "" || payload.Seed == "" {
			web.RespondError(w, http.StatusBadRequest, "missing fields", "missing_fields")
			return
		}

		result, err := repo.CreateGame(r.Context(), ports.GameRecordInput{
			Code:  payload.Code,
			State: payload.State,
			Seed:  payload.Seed,
		})
		if err != nil {
			web.RespondError(w, http.StatusInternalServerError, "create failed", "create_failed")
			return
		}

		web.RespondJSON(w, http.StatusCreated, result)
	}
}

func NewGetGameByIDHandler(repo ports.GameRepository) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id, _ := r.Context().Value(idParamKey).(string)
		if id == "" {
			web.RespondError(w, http.StatusBadRequest, "missing id", "missing_id")
			return
		}

		result, err := repo.GetGameByID(r.Context(), id)
		if err != nil {
			web.RespondError(w, http.StatusInternalServerError, "not found", "game_not_found")
			return
		}

		web.RespondJSON(w, http.StatusOK, result)
	}
}

func NewGetGameByCodeHandler(repo ports.GameRepository) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		code, _ := r.Context().Value(codeParamKey).(string)
		if code == "" {
			web.RespondError(w, http.StatusBadRequest, "missing code", "missing_code")
			return
		}

		result, err := repo.GetGameByCode(r.Context(), code)
		if err != nil {
			if err == repo_pb.ErrRecordNotFound {
				web.RespondError(w, http.StatusNotFound, "not found", "game_not_found")
				return
			}
			web.RespondError(w, http.StatusInternalServerError, "not found", "game_not_found")
			return
		}

		web.RespondJSON(w, http.StatusOK, result)
	}
}
