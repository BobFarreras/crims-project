package http

import (
	"context"
	"net/http"

	"github.com/digitaistudios/crims-backend/internal/ports"
	"github.com/go-chi/chi/v5"
)

// RegisterAPIV1Routes encapsula totes les rutes a /api/v1.
func RegisterAPIV1Routes(r chi.Router, register func(r chi.Router)) {
	r.Route("/api/v1", func(r chi.Router) {
		register(r)
	})
}

// 🔥 NOU: Registre de rutes d'autenticació
func RegisterAuthRoutes(r chi.Router, handler *AuthHandler) {
	r.Route("/api/auth", func(r chi.Router) {
		r.Post("/register", handler.HandleRegister)
		r.Post("/login", handler.HandleLogin)
		r.Post("/logout", handler.HandleLogout)
	})
}

func RegisterMetricsRoutes(r chi.Router) {
	r.Get("/metrics", NewMetricsHandler())
}

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

func RegisterEventRoutes(r chi.Router, service ports.EventService) {
	r.Post("/api/events", NewCreateEventHandler(service))
	r.Get("/api/games/{id}/events", withPathParam(idParamKey, "id", NewListEventsByGameHandler(service)))
}

func RegisterClueRoutes(r chi.Router, service ports.ClueService) {
	r.Post("/api/clues", NewCreateClueHandler(service))
	r.Get("/api/games/{id}/clues", withPathParam(idParamKey, "id", NewListCluesByGameHandler(service)))
}

func RegisterPersonRoutes(r chi.Router, service ports.PersonService) {
	r.Post("/api/persons", NewCreatePersonHandler(service))
	r.Get("/api/games/{id}/persons", withPathParam(idParamKey, "id", NewListPersonsByGameHandler(service)))
}

func RegisterHypothesisRoutes(r chi.Router, service ports.HypothesisService) {
	r.Post("/api/hypotheses", NewCreateHypothesisHandler(service))
	r.Get("/api/games/{id}/hypotheses", withPathParam(idParamKey, "id", NewListHypothesesByGameHandler(service)))
}

func RegisterAccusationRoutes(r chi.Router, service ports.AccusationService) {
	r.Post("/api/accusations", NewCreateAccusationHandler(service))
	r.Get("/api/games/{id}/accusations", withPathParam(idParamKey, "id", NewListAccusationsByGameHandler(service)))
}

func RegisterForensicRoutes(r chi.Router, service ports.ForensicService) {
	r.Post("/api/forensics", NewCreateForensicHandler(service))
	r.Get("/api/games/{id}/forensics", withPathParam(idParamKey, "id", NewListForensicsByGameHandler(service)))
}

func RegisterTimelineRoutes(r chi.Router, service ports.TimelineService) {
	r.Post("/api/timeline", NewCreateTimelineHandler(service))
	r.Get("/api/games/{id}/timeline", withPathParam(idParamKey, "id", NewListTimelineByGameHandler(service)))
}

func RegisterInterrogationRoutes(r chi.Router, service ports.InterrogationService) {
	r.Post("/api/interrogations", NewCreateInterrogationHandler(service))
	r.Get("/api/games/{id}/interrogations", withPathParam(idParamKey, "id", NewListInterrogationsByGameHandler(service)))
}

func RegisterLobbyRoutes(r chi.Router, service ports.LobbyService) {
	r.Post("/api/lobby/join", NewLobbyJoinHandler(service))
}

func withPathParam(key contextKey, name string, next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		value := chi.URLParam(r, name)
		ctx := context.WithValue(r.Context(), key, value)
		next.ServeHTTP(w, r.WithContext(ctx))
	}
}
