package http

import (
	"encoding/json"
	"fmt"
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

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (h *AuthHandler) HandleLogin(w http.ResponseWriter, r *http.Request) {
	var req LoginRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		web.RespondError(w, http.StatusBadRequest, "invalid request body", "auth/bad_request")
		return
	}
	fmt.Printf("🔍 LOGIN INTENT: Email='%s' Pass='%s'\n", req.Email, req.Password)
	// Crida a PocketBase
	authResp, err := h.pbClient.AuthWithPassword(req.Email, req.Password)
	if err != nil {
		web.RespondError(w, http.StatusUnauthorized, "invalid credentials", "auth/invalid_credentials")
		return
	}

	// 🔥 IMPLEMENTACIÓ COOKIE SEGURA (OWASP)
	http.SetCookie(w, &http.Cookie{
		Name:  "auth_token",   // Nom de la cookie
		Value: authResp.Token, // El token JWT
		Path:  "/",            // Disponible a tota l'app

		// 🛡️ SEGURETAT MÀXIMA
		HttpOnly: true,                 // JS no la pot llegir (Anti-XSS)
		SameSite: http.SameSiteLaxMode, // Protecció CSRF bàsica

		// ⚠️ EN PRODUCCIÓ: Posa Secure: true (només HTTPS).
		// En localhost (HTTP), ha de ser false o el navegador la bloquejarà.
		Secure: false,

		MaxAge: 7 * 24 * 60 * 60, // 7 dies (en segons)
	})

	// Resposta JSON neta (sense token visible)
	response := map[string]interface{}{
		"message": "Login successful",
		"user": map[string]string{
			"id":       authResp.Record.ID,
			"username": authResp.Record.Username,
			"name":     authResp.Record.Name,
		},
	}

	web.RespondJSON(w, http.StatusOK, response)
}

// HandleLogout tanca la sessió invalidant la cookie
func (h *AuthHandler) HandleLogout(w http.ResponseWriter, r *http.Request) {
	// 🔥 TÈCNICA SEGURA: Sobreescriure la cookie amb caducitat immediata
	http.SetCookie(w, &http.Cookie{
		Name:     "auth_token",
		Value:    "",  // Valor buit
		Path:     "/", // Ha de coincidir amb el Path original!
		HttpOnly: true,
		SameSite: http.SameSiteLaxMode,
		Secure:   false, // Recorda: true en producció (HTTPS)
		MaxAge:   -1,    // 💀 Aixo mata la cookie a l'instant
	})

	web.RespondJSON(w, http.StatusOK, map[string]string{"message": "Logged out successfully"})
}
