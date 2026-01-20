package main

import (
	"log"
	"net/http"
	"os"
	"time"

	// 1. IMPORT INTERN (El teu codi)
	// Li diem "myMiddleware" per no confondre'l amb el de Chi,
	// o simplement usem el nom del paquet "middleware" si l'altre l'anomenem diferent.
	"github.com/digitaistudios/crims-backend/internal/middleware"

	// 2. IMPORT INTERN (La teva utilitat web)
	"github.com/digitaistudios/crims-backend/internal/platform/web"

	// 3. IMPORTS EXTERNS
	"github.com/getsentry/sentry-go"
	sentryhttp "github.com/getsentry/sentry-go/http"
	"github.com/go-chi/chi/v5"
	// ALERTA: Aquí li posem un nom diferent (chimiddleware) per evitar el conflicte!
	chimiddleware "github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
)

func main() {
	// Inicialitzar Sentry per error tracking
	err := sentry.Init(sentry.ClientOptions{
		Dsn:         os.Getenv("SENTRY_DSN"),
		Environment: os.Getenv("ENVIRONMENT"),
		// Sample Rate (10% de traces)
		TracesSampleRate: 0.1,
	})
	if err != nil {
		log.Printf("⚠️  Sentry init failed: %v", err)
	} else {
		// Configurar tags globales
		sentry.ConfigureScope(func(scope *sentry.Scope) {
			scope.SetTag("app", "crims-backend")
			scope.SetTag("runtime", "go")
			scope.SetTag("framework", "chi")
		})
		log.Println("✅ Sentry inicialitzat correctament")
	}
	// Flush events abans de sortir
	defer sentry.Flush(2 * time.Second)

	// Ara "middleware" es refereix a la TEVA carpeta internal/middleware
	logger := middleware.SetupLogger()
	logger.Info("🔌 Inicialitzant Crims de Mitjanit Backend...")

	r := chi.NewRouter()

	// Middleware de Sentry (captura panics)
	sentryHandler := sentryhttp.New(sentryhttp.Options{
		Repanic:         true,
		WaitForDelivery: false,
	})
	r.Use(func(next http.Handler) http.Handler {
		return sentryHandler.Handle(next)
	})

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

	listenErr := http.ListenAndServe(":"+port, r)
	if listenErr != nil {
		logger.Error("❌ Error fatal al servidor", "error", listenErr)
		os.Exit(1)
	}
}
