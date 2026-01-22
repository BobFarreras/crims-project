package http

import (
	"encoding/json"
	"net/http"

	"github.com/digitaistudios/crims-backend/internal/platform/web"
	"github.com/digitaistudios/crims-backend/internal/ports"
)

type AuthHandler struct {
	pbClient ports.PocketBaseClient
}

func NewAuthHandler(pbClient ports.PocketBaseClient) *AuthHandler {
	return &AuthHandler{
		pbClient: pbClient,
	}
}

type RegisterRequest struct {
	Username        string `json:"username"`
	Email           string `json:"email"`
	Password        string `json:"password"`
	PasswordConfirm string `json:"passwordConfirm"`
	Name            string `json:"name"`
}

func (h *AuthHandler) HandleRegister(w http.ResponseWriter, r *http.Request) {
	var req RegisterRequest

	// ⚠️ CRÍTIC: Faltava aquesta part per llegir les dades del frontend!
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		// Passem el 4t argument (codi d'error)
		web.RespondError(w, http.StatusBadRequest, "invalid request body", "auth/bad_request")
		return
	}

	// Validació bàsica
	if req.Password != req.PasswordConfirm {
		web.RespondError(w, http.StatusBadRequest, "passwords do not match", "auth/password_mismatch")
		return
	}

	// ✅ IMPLEMENTACIÓ REAL
	err := h.pbClient.CreateUser(req.Username, req.Email, req.Password, req.PasswordConfirm, req.Name)
	if err != nil {
		// En producció, podries mirar si és error de validació o de servidor
		web.RespondError(w, http.StatusBadRequest, "registration failed: "+err.Error(), "auth/registration_failed")
		return
	}

	web.RespondJSON(w, http.StatusOK, map[string]string{"message": "User registered successfully"})
}

func (h *AuthHandler) HandleLogin(w http.ResponseWriter, r *http.Request) {
	// Implementació futura
	web.RespondJSON(w, http.StatusOK, map[string]string{"token": "fake-jwt-token"})
}
