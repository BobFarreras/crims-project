package main

import (
	"net/http"
	"os"

	// 1. IMPORT INTERN (El teu codi)
	// Li diem "myMiddleware" per no confondre'l amb el de Chi,
	// o simplement usem el nom del paquet "middleware" si l'altre l'anomenem diferent.
	"github.com/digitaistudios/crims-backend/internal/middleware"

	// 2. IMPORT INTERN (La teva utilitat web)
	"github.com/digitaistudios/crims-backend/internal/platform/web"

	// 3. IMPORTS EXTERNS
	"github.com/go-chi/chi/v5"
	// ALERTA: Aquí li posem un nom diferent (chimiddleware) per evitar el conflicte!
	chimiddleware "github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
)

func main() {
	// Ara "middleware" es refereix a la TEVA carpeta internal/middleware
	logger := middleware.SetupLogger()
	logger.Info("🔌 Inicialitzant Crims de Mitjanit Backend...")

	r := chi.NewRouter()

	// Ara "chimiddleware" es refereix a la llibreria externa
	r.Use(chimiddleware.Recoverer)

	// Usem el teu middleware propi
	r.Use(middleware.RequestLogger(logger))

	r.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"http://localhost:3000"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: true,
		MaxAge:           300,
	}))

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("🕵️‍♂️ Backend Operatiu amb Logs i Seguretat"))
	})

	r.Get("/api/status", func(w http.ResponseWriter, r *http.Request) {
		// Usem el teu paquet "web" que hem arreglat al Pas 1
		status := map[string]string{
			"system":  "Crims Backend",
			"status":  "healthy",
			"version": "0.1.0-alpha",
		}
		web.RespondJSON(w, http.StatusOK, status)
	})

	port := "8080"
	logger.Info("🚀 Servidor escoltant", "port", port, "url", "http://localhost:"+port)

	err := http.ListenAndServe(":"+port, r)
	if err != nil {
		logger.Error("❌ Error fatal al servidor", "error", err)
		os.Exit(1)
	}
}
