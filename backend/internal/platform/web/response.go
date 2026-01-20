package web

import (
	"encoding/json"
	"net/http"
)

// ErrorResponse estructura estàndard per a errors JSON
type ErrorResponse struct {
	Status  int    `json:"status"`
	Message string `json:"error"`
	Code    string `json:"code,omitempty"`
}

// RespondJSON envia una resposta JSON correcta
func RespondJSON(w http.ResponseWriter, status int, payload interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(payload)
}

// RespondError envia un error formatat
func RespondError(w http.ResponseWriter, status int, message string, code string) {
	RespondJSON(w, status, ErrorResponse{
		Status:  status,
		Message: message,
		Code:    code,
	})
}
