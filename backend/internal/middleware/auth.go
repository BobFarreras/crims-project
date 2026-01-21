package middleware

import (
	"context"
	"net/http"
	"strings"
)

type contextKey string

const RoleKey contextKey = "role"

// AuthMiddleware valida que existeix token d'autenticacio.
func AuthMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Header.Get("Authorization") == "" {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		role := r.Header.Get("X-Role")
		ctx := context.WithValue(r.Context(), RoleKey, role)
		next.ServeHTTP(w, r.WithContext(ctx))
	}
}

// RequireRole garanteix que el rol esta autoritzat.
func RequireRole(roles ...string) func(http.Handler) http.Handler {
	allowed := map[string]struct{}{}
	for _, role := range roles {
		allowed[strings.ToUpper(role)] = struct{}{}
	}

	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
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
