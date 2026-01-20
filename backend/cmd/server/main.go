package main

import (
	"fmt"
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

	// ===============================
	// TESTS DE SENTRY (BACKEND)
	// ===============================

	// Test 1: Error manual amb captureException
	r.Get("/api/test-sentry/error1", func(w http.ResponseWriter, r *http.Request) {
		// Error manual
		err := fmt.Errorf("Test Error 1: Error de Go manual")
		sentry.CaptureException(err)
		web.RespondJSON(w, http.StatusOK, map[string]string{
			"message": "Error capturat! Mira Sentry Dashboard",
			"dsn":     os.Getenv("SENTRY_DSN"),
		})
	})

	// Test 2: Error amb context
	r.Get("/api/test-sentry/error2", func(w http.ResponseWriter, r *http.Request) {
		// Añadir contexto
		sentry.ConfigureScope(func(scope *sentry.Scope) {
			scope.SetContext("test_context", map[string]interface{}{
				"test_type": "manual_trigger",
				"endpoint":  "/api/test-sentry/error2",
				"timestamp": time.Now().Unix(),
			})

			scope.SetTag("test_type", "manual_error")
			scope.SetTag("backend", "go")
		})

		// Capturar error amb context
		err := fmt.Errorf("Test Error 2: Error amb context")
		sentry.CaptureException(err)

		web.RespondJSON(w, http.StatusOK, map[string]string{
			"message": "Error amb context capturat! Mira Sentry Dashboard",
		})
	})

	// Test 3: Error amb nivell de severitat
	r.Get("/api/test-sentry/error3", func(w http.ResponseWriter, r *http.Request) {
		err := fmt.Errorf("Test Error 3: Error amb nivell de severitat")

		sentry.WithScope(func(scope *sentry.Scope) {
			scope.SetLevel(sentry.LevelError)
			sentry.CaptureException(err)
		})

		web.RespondJSON(w, http.StatusOK, map[string]string{
			"message": "Error amb nivell de severitat capturat! Mira Sentry Dashboard",
		})
	})

	// Test 4: Panic (middleware hauria de capturar-ho)
	r.Get("/api/test-sentry/panic", func(w http.ResponseWriter, r *http.Request) {
		// Això hauria de ser capturat pel middleware de Sentry
		panic("Test Error 4: Panic intentional")
	})

	// Test 5: Capturar missatge (no error)
	r.Get("/api/test-sentry/message", func(w http.ResponseWriter, r *http.Request) {
		sentry.CaptureMessage("Test Message: Alguna cosa ha passat")
		web.RespondJSON(w, http.StatusOK, map[string]string{
			"message": "Missatge capturat! Mira Sentry Dashboard",
		})
	})

	port := "8080"
	logger.Info("🚀 Servidor escoltant", "port", port, "url", "http://localhost:"+port)

	listenErr := http.ListenAndServe(":"+port, r)
	if listenErr != nil {
		logger.Error("❌ Error fatal al servidor", "error", listenErr)
		os.Exit(1)
	}
}
