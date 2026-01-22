package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/joho/godotenv"

	apihttp "github.com/digitaistudios/crims-backend/internal/adapters/http"
	"github.com/digitaistudios/crims-backend/internal/adapters/repo_pb"
	"github.com/digitaistudios/crims-backend/internal/middleware"
	"github.com/digitaistudios/crims-backend/internal/platform/config"
	"github.com/digitaistudios/crims-backend/internal/platform/web"
	"github.com/digitaistudios/crims-backend/internal/ports"
	"github.com/digitaistudios/crims-backend/internal/services"

	"github.com/getsentry/sentry-go"
	sentryhttp "github.com/getsentry/sentry-go/http"
	"github.com/go-chi/chi/v5"
	chimiddleware "github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
)

func main() {
	// 1. Càrrega de Configuració i Entorn
	loadEnv()
	cfg, cfgErr := config.Load()
	if cfgErr != nil {
		log.Fatalf("❌ Error configuracio: %v", cfgErr)
	}

	// 2. Inicialització de Sentry
	initSentry(cfg)
	defer sentry.Flush(2 * time.Second)

	// 3. Setup Logger
	logger := middleware.SetupLogger()
	logger.Info("🔌 Inicialitzant Crims de Mitjanit Backend...")

	// 4. Inicialització de Repositoris (PocketBase)
	var (
		pocketBaseClient        ports.PocketBaseClient
		gameRepository          ports.GameRepository
		playerRepository        ports.PlayerRepository
		eventRepository         ports.EventRepository
		clueRepository          ports.ClueRepository
		personRepository        ports.PersonRepository
		hypothesisRepository    ports.HypothesisRepository
		accusationRepository    ports.AccusationRepository
		forensicRepository      ports.ForensicRepository
		timelineRepository      ports.TimelineRepository
		interrogationRepository ports.InterrogationRepository
	)

	pbClient, pbErr := repo_pb.NewClient(repo_pb.Config{
		BaseURL: cfg.PocketBaseURL,
		Timeout: cfg.PocketBaseTimeout,
	})

	if pbErr != nil {
		logger.Warn("PocketBase client disabled", "error", pbErr)
		// Usem les implementacions "Disabled" (netejades del main)
		pocketBaseClient = repo_pb.DisabledPocketBaseClient{Err: pbErr}
		gameRepository = repo_pb.DisabledGameRepository{Err: pbErr}
		playerRepository = repo_pb.DisabledPlayerRepository{Err: pbErr}
		eventRepository = repo_pb.DisabledEventRepository{Err: pbErr}
		clueRepository = repo_pb.DisabledClueRepository{Err: pbErr}
		personRepository = repo_pb.DisabledPersonRepository{Err: pbErr}
		hypothesisRepository = repo_pb.DisabledHypothesisRepository{Err: pbErr}
		accusationRepository = repo_pb.DisabledAccusationRepository{Err: pbErr}
		forensicRepository = repo_pb.DisabledForensicRepository{Err: pbErr}
		timelineRepository = repo_pb.DisabledTimelineRepository{Err: pbErr}
		interrogationRepository = repo_pb.DisabledInterrogationRepository{Err: pbErr}
	} else {
		pocketBaseClient = pbClient
		gameRepository = repo_pb.NewGameRepository(pbClient)
		playerRepository = repo_pb.NewPlayerRepository(pbClient)
		eventRepository = repo_pb.NewEventRepository(pbClient)
		clueRepository = repo_pb.NewClueRepository(pbClient)
		personRepository = repo_pb.NewPersonRepository(pbClient)
		hypothesisRepository = repo_pb.NewHypothesisRepository(pbClient)
		accusationRepository = repo_pb.NewAccusationRepository(pbClient)
		forensicRepository = repo_pb.NewForensicRepository(pbClient)
		timelineRepository = repo_pb.NewTimelineRepository(pbClient)
		interrogationRepository = repo_pb.NewInterrogationRepository(pbClient)
	}

	// 5. Inicialització de Serveis
	gameService := services.NewGameService(gameRepository)
	playerService := services.NewPlayerService(playerRepository)
	eventService := services.NewEventService(eventRepository)
	clueService := services.NewClueService(clueRepository)
	personService := services.NewPersonService(personRepository)
	hypothesisService := services.NewHypothesisService(hypothesisRepository)
	accusationService := services.NewAccusationService(accusationRepository)
	forensicService := services.NewForensicService(forensicRepository)
	timelineService := services.NewTimelineService(timelineRepository)
	interrogationService := services.NewInterrogationService(interrogationRepository)
	lobbyService := services.NewLobbyService(gameRepository, playerRepository)

	// 6. Setup Router
	r := chi.NewRouter()

	// Middlewares
	r.Use(sentryhttp.New(sentryhttp.Options{Repanic: true}).Handle)
	r.Use(chimiddleware.Recoverer)
	r.Use(middleware.RequestLogger(logger))
	r.Use(cors.Handler(cors.Options{
		AllowedOrigins:   cfg.AllowedOrigins,
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: true,
		MaxAge:           300,
	}))

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("🕵️‍♂️ Backend Operatiu amb Logs i Seguretat"))
	})

	apihttp.RegisterAPIV1Routes(r, func(r chi.Router) {
		r.Get("/status", func(w http.ResponseWriter, r *http.Request) {
			status := map[string]string{
				"system":  "Crims Backend",
				"status":  "healthy",
				"version": "0.1.0-alpha",
			}
			web.RespondJSON(w, http.StatusOK, status)
		})
		r.Get("/health", apihttp.NewHealthHandler(pocketBaseClient))
		apihttp.RegisterMetricsRoutes(r)
		apihttp.RegisterGameRoutes(r, gameService)
		apihttp.RegisterPlayerRoutes(r, playerService)
		apihttp.RegisterEventRoutes(r, eventService)
		apihttp.RegisterClueRoutes(r, clueService)
		apihttp.RegisterPersonRoutes(r, personService)
		apihttp.RegisterHypothesisRoutes(r, hypothesisService)
		apihttp.RegisterAccusationRoutes(r, accusationService)
		apihttp.RegisterForensicRoutes(r, forensicService)
		apihttp.RegisterTimelineRoutes(r, timelineService)
		apihttp.RegisterInterrogationRoutes(r, interrogationService)
		apihttp.RegisterLobbyRoutes(r, lobbyService)

		// Debug Sentry (Opcional, pots treure-ho si no ho vols aquí)
		registerSentryDebugRoutes(r)
	})

	logger.Info("🚀 Servidor escoltant", "port", cfg.Port, "url", "http://localhost:"+cfg.Port)
	if err := http.ListenAndServe(":"+cfg.Port, r); err != nil {
		logger.Error("❌ Error fatal al servidor", "error", err)
		os.Exit(1)
	}
}

// Helpers per mantenir el main net

func loadEnv() {
	err := godotenv.Load("../.env.local")
	if err != nil {
		err = godotenv.Load(".env")
		if err != nil {
			log.Printf("⚠️  Warning: No s'ha pogut carregar .env.local o .env: %v", err)
		} else {
			log.Println("✅ Carregat .env")
		}
	} else {
		log.Println("✅ Carregat .env.local")
	}
}

func initSentry(cfg config.Config) {
	err := sentry.Init(sentry.ClientOptions{
		Dsn:              os.Getenv("SENTRY_DSN"),
		Environment:      cfg.Environment,
		TracesSampleRate: 0.1,
	})
	if err != nil {
		log.Printf("⚠️  Sentry init failed: %v", err)
	} else {
		sentry.ConfigureScope(func(scope *sentry.Scope) {
			scope.SetTag("app", "crims-backend")
		})
		log.Println("✅ Sentry inicialitzat")
	}
}

func registerSentryDebugRoutes(r chi.Router) {
	r.Get("/test-sentry/debug", func(w http.ResponseWriter, r *http.Request) {
		web.RespondJSON(w, http.StatusOK, map[string]bool{"sentry_active": os.Getenv("SENTRY_DSN") != ""})
	})
	r.Get("/test-sentry/error", func(w http.ResponseWriter, r *http.Request) {
		sentry.CaptureException(fmt.Errorf("Test Error Manual"))
		web.RespondJSON(w, http.StatusOK, map[string]string{"msg": "Error enviat"})
	})
}
