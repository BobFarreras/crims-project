package http

import (
	"context"
	"net/http"

	"github.com/digitaistudios/crims-backend/internal/ports"
	"github.com/go-chi/chi/v5"
)

// RegisterGameRoutes defineix les rutes de Game.
func RegisterGameRoutes(r chi.Router, repo ports.GameRepository) {
	r.Post("/api/games", NewCreateGameHandler(repo))
	r.Get("/api/games/{id}", withPathParam(idParamKey, "id", NewGetGameByIDHandler(repo)))
	r.Get("/api/games/by-code/{code}", withPathParam(codeParamKey, "code", NewGetGameByCodeHandler(repo)))
}

func withPathParam(key contextKey, name string, next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		value := chi.URLParam(r, name)
		ctx := context.WithValue(r.Context(), key, value)
		next.ServeHTTP(w, r.WithContext(ctx))
	}
}
