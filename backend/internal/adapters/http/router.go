package http

import (
	"context"
	"net/http"

	"github.com/digitaistudios/crims-backend/internal/ports"
	"github.com/go-chi/chi/v5"
)

// RegisterGameRoutes defineix les rutes de Game.
func RegisterGameRoutes(r chi.Router, service ports.GameService) {
	r.Post("/api/games", NewCreateGameHandler(service))
	r.Get("/api/games/{id}", withPathParam(idParamKey, "id", NewGetGameByIDHandler(service)))
	r.Get("/api/games/by-code/{code}", withPathParam(codeParamKey, "code", NewGetGameByCodeHandler(service)))
}

func RegisterPlayerRoutes(r chi.Router, service ports.PlayerService) {
	r.Post("/api/players", NewCreatePlayerHandler(service))
	r.Get("/api/games/{id}/players", withPathParam(idParamKey, "id", NewListPlayersByGameHandler(service)))
}

func withPathParam(key contextKey, name string, next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		value := chi.URLParam(r, name)
		ctx := context.WithValue(r.Context(), key, value)
		next.ServeHTTP(w, r.WithContext(ctx))
	}
}
