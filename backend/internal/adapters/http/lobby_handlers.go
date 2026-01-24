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
	GameCode     string   `json:"gameCode"`
	UserID       string   `json:"userId"`
	Capabilities []string `json:"capabilities"`
}

type lobbyCreateRequest struct {
	UserID       string   `json:"userId"`
	Capabilities []string `json:"capabilities"`
}

func NewLobbyJoinHandler(service ports.LobbyService, pbClient ports.PocketBaseClient) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		authResp, token, authErr := extractAuth(r, pbClient)
		if authErr != nil {
			web.RespondError(w, authErr.Status, authErr.Message, authErr.Code)
			return
		}

		var payload lobbyJoinRequest
		if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
			web.RespondError(w, http.StatusBadRequest, "invalid payload", "invalid_payload")
			return
		}

		userID := authResp.Record.ID
		if payload.UserID != "" && payload.UserID != userID {
			web.RespondError(w, http.StatusForbidden, "user mismatch", "auth/user_mismatch")
			return
		}

		ctx := ports.WithAuthToken(r.Context(), token)
		result, err := service.JoinGame(ctx, payload.GameCode, userID, payload.Capabilities)
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

func NewLobbyCreateHandler(service ports.LobbyService, pbClient ports.PocketBaseClient) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		authResp, token, authErr := extractAuth(r, pbClient)
		if authErr != nil {
			web.RespondError(w, authErr.Status, authErr.Message, authErr.Code)
			return
		}

		var payload lobbyCreateRequest
		if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
			web.RespondError(w, http.StatusBadRequest, "invalid payload", "invalid_payload")
			return
		}

		userID := authResp.Record.ID
		if payload.UserID != "" && payload.UserID != userID {
			web.RespondError(w, http.StatusForbidden, "user mismatch", "auth/user_mismatch")
			return
		}

		ctx := ports.WithAuthToken(r.Context(), token)

		result, err := service.CreateLobby(ctx, userID, payload.Capabilities)
		if err != nil {
			if errors.Is(err, services.ErrInvalidLobbyInput) {
				web.RespondError(w, http.StatusBadRequest, "missing fields", "missing_fields")
				return
			}
			web.RespondError(w, http.StatusInternalServerError, "create failed", "create_failed")
			return
		}

		web.RespondJSON(w, http.StatusCreated, result)
	}
}

type authError struct {
	Status  int
	Message string
	Code    string
}

func extractAuth(r *http.Request, pbClient ports.PocketBaseClient) (*ports.AuthResponse, string, *authError) {
	cookie, err := r.Cookie("auth_token")
	if err != nil || cookie.Value == "" {
		return nil, "", &authError{Status: http.StatusUnauthorized, Message: "missing session", Code: "auth/missing_session"}
	}

	authResp, err := pbClient.RefreshAuth(cookie.Value)
	if err != nil {
		return nil, "", &authError{Status: http.StatusUnauthorized, Message: "invalid session", Code: "auth/invalid_session"}
	}

	return authResp, cookie.Value, nil
}
