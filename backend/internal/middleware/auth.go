package middleware

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"strings"

	"github.com/golang-jwt/jwt/v5"
)

type contextKey string

const (
	RoleKey   contextKey = "role"
	UserIDKey contextKey = "user_id"
)

// AuthMiddleware valida el JWT de PocketBase.
func AuthMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		// Format esperat: "Bearer <token>"
		parts := strings.Split(authHeader, " ")
		if len(parts) != 2 || parts[0] != "Bearer" {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
		tokenString := parts[1]

		// Validar el token
		claims, err := validateToken(tokenString)
		if err != nil {
			// Token invàlid, caducat o signatura incorrecta
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		// Extreure informació del token
		userID, _ := claims["id"].(string)

		// Intentem llegir el rol si ve al token (custom claim), sinó assumim "USER"
		role, ok := claims["role"].(string)
		if !ok || role == "" {
			role = "USER"
			// Nota: Si necessites rols específics de joc (DETECTIVE, etc.),
			// hauràs de mirar la DB o afegir-ho als claims de PocketBase.
		}

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

// validateToken parseja i valida la signatura del JWT
func validateToken(tokenString string) (jwt.MapClaims, error) {
	secret := os.Getenv("JWT_SECRET")
	if secret == "" {
		// Fallback insegur només per dev si no està configurat, però avisant
		// En producció això hauria de fallar fatalment.
		return nil, fmt.Errorf("JWT_SECRET not configured")
	}

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Validar l'algoritme de signatura (HMAC)
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(secret), nil
	})

	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return claims, nil
	}

	return nil, fmt.Errorf("invalid token")
}
