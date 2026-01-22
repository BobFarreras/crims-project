package middleware

import (
	"context"
	"net/http"
	"strings"

	"github.com/digitaistudios/crims-backend/internal/ports"
)

type contextKey string

const (
	RoleKey   contextKey = "role"
	UserIDKey contextKey = "user_id"
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
		role := "USER"

		// Posem la informació segura al context
		ctx := context.WithValue(r.Context(), UserIDKey, userID)
		ctx = context.WithValue(ctx, RoleKey, role)

		next.ServeHTTP(w, r.WithContext(ctx))
	}
}

// RequireRole garanteix que el rol (validat pel token) està autoritzat.
func RequireRole(roles ...string) func(http.Handler) http.Handler {
	allowed := map[string]struct{}{}
	for _, role := range roles {
		allowed[strings.ToUpper(role)] = struct{}{}
	}

	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			// Ara el rol ve del context segur (JWT), no de la capçalera HTTP
			role, _ := r.Context().Value(RoleKey).(string)
			if role == "" {
				w.WriteHeader(http.StatusForbidden)
				return
			}
			if _, ok := allowed[strings.ToUpper(role)]; !ok {
				w.WriteHeader(http.StatusForbidden)
				return
			}

			next.ServeHTTP(w, r)
		})
	}
}
