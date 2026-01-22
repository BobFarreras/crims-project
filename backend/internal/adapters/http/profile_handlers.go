package http

import (
	"encoding/json"
	"net/http"

	"github.com/digitaistudios/crims-backend/internal/platform/web"
	"github.com/digitaistudios/crims-backend/internal/ports"
)

type ProfileHandler struct {
	pbClient ports.PocketBaseClient
}

type profileUpdateRequest struct {
	Name string `json:"name"`
}

func NewProfileHandler(pbClient ports.PocketBaseClient) *ProfileHandler {
	return &ProfileHandler{pbClient: pbClient}
}

func (h *ProfileHandler) HandleGetProfile(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie("auth_token")
	if err != nil || cookie.Value == "" {
		web.RespondError(w, http.StatusUnauthorized, "missing session", "auth/missing_session")
		return
	}

	authResp, err := h.pbClient.RefreshAuth(cookie.Value)
	if err != nil {
		web.RespondError(w, http.StatusUnauthorized, "invalid session", "auth/invalid_session")
		return
	}

	name := authResp.Record.Name
	if name == "" {
		name = authResp.Record.Username
	}

	response := map[string]interface{}{
		"user": map[string]string{
			"id":       authResp.Record.ID,
			"username": authResp.Record.Username,
			"name":     name,
		},
	}

	web.RespondJSON(w, http.StatusOK, response)
}

func (h *ProfileHandler) HandleUpdateProfile(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie("auth_token")
	if err != nil || cookie.Value == "" {
		web.RespondError(w, http.StatusUnauthorized, "missing session", "auth/missing_session")
		return
	}

	var payload profileUpdateRequest
	if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
		web.RespondError(w, http.StatusBadRequest, "invalid request body", "auth/bad_request")
		return
	}
	if payload.Name == "" {
		web.RespondError(w, http.StatusBadRequest, "missing name", "auth/missing_profile_name")
		return
	}

	authResp, err := h.pbClient.RefreshAuth(cookie.Value)
	if err != nil {
		web.RespondError(w, http.StatusUnauthorized, "invalid session", "auth/invalid_session")
		return
	}

	if err := h.pbClient.UpdateUserName(cookie.Value, authResp.Record.ID, payload.Name); err != nil {
		if statusErr, ok := err.(interface{ Status() int }); ok {
			switch statusErr.Status() {
			case http.StatusUnauthorized, http.StatusForbidden:
				web.RespondError(w, statusErr.Status(), "not authorized", "auth/profile_update_forbidden")
				return
			case http.StatusNotFound:
				web.RespondError(w, http.StatusNotFound, "user not found", "auth/profile_not_found")
				return
			}
		}
		web.RespondError(w, http.StatusInternalServerError, "update failed", "auth/profile_update_failed")
		return
	}

	updated, err := h.pbClient.RefreshAuth(cookie.Value)
	if err != nil {
		web.RespondError(w, http.StatusUnauthorized, "invalid session", "auth/invalid_session")
		return
	}

	name := updated.Record.Name
	if name == "" {
		name = updated.Record.Username
	}

	response := map[string]interface{}{
		"user": map[string]string{
			"id":       updated.Record.ID,
			"username": updated.Record.Username,
			"name":     name,
		},
	}

	web.RespondJSON(w, http.StatusOK, response)
}
