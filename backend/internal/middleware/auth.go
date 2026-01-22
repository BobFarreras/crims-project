package middleware

import (
	"context"
	"net/http"
	"strings"

	"github.com/digitaistudios/crims-backend/internal/ports"
)

type contextKey string

const (
	CapabilityKey contextKey = "capabilities"
	UserIDKey     contextKey = "user_id"
)

// AuthMiddleware valida el JWT de PocketBase via cookie.
func AuthMiddleware(pbClient ports.PocketBaseClient, next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		cookie, err := r.Cookie("auth_token")
		if err != nil || cookie.Value == "" {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		authResp, err := pbClient.RefreshAuth(cookie.Value)
		if err != nil {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		userID := authResp.Record.ID
		capabilities := []string{}

		// Posem la informació segura al context
		ctx := context.WithValue(r.Context(), UserIDKey, userID)
		ctx = context.WithValue(ctx, CapabilityKey, capabilities)

		next.ServeHTTP(w, r.WithContext(ctx))
	}
}

// RequireCapability garanteix que la capacitat està autoritzada.
func RequireCapability(capabilities ...string) func(http.Handler) http.Handler {
	allowed := map[string]struct{}{}
	for _, capability := range capabilities {
		allowed[strings.ToUpper(capability)] = struct{}{}
	}

	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			stored, _ := r.Context().Value(CapabilityKey).([]string)
			if len(stored) == 0 {
				w.WriteHeader(http.StatusForbidden)
				return
			}
			allowedMatch := false
			for _, capability := range stored {
				if _, ok := allowed[strings.ToUpper(capability)]; ok {
					allowedMatch = true
					break
				}
			}
			if !allowedMatch {
				w.WriteHeader(http.StatusForbidden)
				return
			}

			next.ServeHTTP(w, r)
		})
	}
}
