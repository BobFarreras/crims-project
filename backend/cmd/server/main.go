package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/joho/godotenv"

	// 1. IMPORT INTERN (El teu codi)
	// Li diem "myMiddleware" per no confondre'l amb el de Chi,
	// o simplement usem el nom del paquet "middleware" si l'altre l'anomenem diferent.
	apihttp "github.com/digitaistudios/crims-backend/internal/adapters/http"
	"github.com/digitaistudios/crims-backend/internal/adapters/repo_pb"
	"github.com/digitaistudios/crims-backend/internal/middleware"
	"github.com/digitaistudios/crims-backend/internal/platform/config"
	"github.com/digitaistudios/crims-backend/internal/ports"
	"github.com/digitaistudios/crims-backend/internal/services"

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

type disabledPocketBaseClient struct {
	err error
}

func (d disabledPocketBaseClient) Ping(ctx context.Context) error {
	return d.err
}

type disabledGameRepository struct {
	err error
}

func (d disabledGameRepository) CreateGame(ctx context.Context, input ports.GameRecordInput) (ports.GameRecord, error) {
	return ports.GameRecord{}, d.err
}

func (d disabledGameRepository) GetGameByID(ctx context.Context, id string) (ports.GameRecord, error) {
	return ports.GameRecord{}, d.err
}

func (d disabledGameRepository) GetGameByCode(ctx context.Context, code string) (ports.GameRecord, error) {
	return ports.GameRecord{}, d.err
}

type disabledPlayerRepository struct {
	err error
}

func (d disabledPlayerRepository) CreatePlayer(ctx context.Context, input ports.PlayerRecordInput) (ports.PlayerRecord, error) {
	return ports.PlayerRecord{}, d.err
}

func (d disabledPlayerRepository) GetPlayerByID(ctx context.Context, id string) (ports.PlayerRecord, error) {
	return ports.PlayerRecord{}, d.err
}

func (d disabledPlayerRepository) ListPlayersByGame(ctx context.Context, gameID string) ([]ports.PlayerRecord, error) {
	return nil, d.err
}

type disabledEventRepository struct {
	err error
}

func (d disabledEventRepository) CreateEvent(ctx context.Context, input ports.EventRecordInput) (ports.EventRecord, error) {
	return ports.EventRecord{}, d.err
}

func (d disabledEventRepository) GetEventByID(ctx context.Context, id string) (ports.EventRecord, error) {
	return ports.EventRecord{}, d.err
}

func (d disabledEventRepository) ListEventsByGame(ctx context.Context, gameID string) ([]ports.EventRecord, error) {
	return nil, d.err
}

type disabledClueRepository struct {
	err error
}

func (d disabledClueRepository) CreateClue(ctx context.Context, input ports.ClueRecordInput) (ports.ClueRecord, error) {
	return ports.ClueRecord{}, d.err
}

func (d disabledClueRepository) GetClueByID(ctx context.Context, id string) (ports.ClueRecord, error) {
	return ports.ClueRecord{}, d.err
}

func (d disabledClueRepository) ListCluesByGame(ctx context.Context, gameID string) ([]ports.ClueRecord, error) {
	return nil, d.err
}

type disabledPersonRepository struct {
	err error
}

func (d disabledPersonRepository) CreatePerson(ctx context.Context, input ports.PersonRecordInput) (ports.PersonRecord, error) {
	return ports.PersonRecord{}, d.err
}

func (d disabledPersonRepository) GetPersonByID(ctx context.Context, id string) (ports.PersonRecord, error) {
	return ports.PersonRecord{}, d.err
}

func (d disabledPersonRepository) ListPersonsByGame(ctx context.Context, gameID string) ([]ports.PersonRecord, error) {
	return nil, d.err
}

type disabledHypothesisRepository struct {
	err error
}

func (d disabledHypothesisRepository) CreateHypothesis(ctx context.Context, input ports.HypothesisRecordInput) (ports.HypothesisRecord, error) {
	return ports.HypothesisRecord{}, d.err
}

func (d disabledHypothesisRepository) GetHypothesisByID(ctx context.Context, id string) (ports.HypothesisRecord, error) {
	return ports.HypothesisRecord{}, d.err
}

func (d disabledHypothesisRepository) ListHypothesesByGame(ctx context.Context, gameID string) ([]ports.HypothesisRecord, error) {
	return nil, d.err
}

type disabledAccusationRepository struct {
	err error
}

func (d disabledAccusationRepository) CreateAccusation(ctx context.Context, input ports.AccusationRecordInput) (ports.AccusationRecord, error) {
	return ports.AccusationRecord{}, d.err
}

func (d disabledAccusationRepository) GetAccusationByID(ctx context.Context, id string) (ports.AccusationRecord, error) {
	return ports.AccusationRecord{}, d.err
}

func (d disabledAccusationRepository) ListAccusationsByGame(ctx context.Context, gameID string) ([]ports.AccusationRecord, error) {
	return nil, d.err
}

type disabledForensicRepository struct {
	err error
}

func (d disabledForensicRepository) CreateAnalysis(ctx context.Context, input ports.ForensicRecordInput) (ports.ForensicRecord, error) {
	return ports.ForensicRecord{}, d.err
}

func (d disabledForensicRepository) GetAnalysisByID(ctx context.Context, id string) (ports.ForensicRecord, error) {
	return ports.ForensicRecord{}, d.err
}

func (d disabledForensicRepository) ListAnalysesByGame(ctx context.Context, gameID string) ([]ports.ForensicRecord, error) {
	return nil, d.err
}

type disabledTimelineRepository struct {
	err error
}

func (d disabledTimelineRepository) CreateEntry(ctx context.Context, input ports.TimelineRecordInput) (ports.TimelineRecord, error) {
	return ports.TimelineRecord{}, d.err
}

func (d disabledTimelineRepository) GetEntryByID(ctx context.Context, id string) (ports.TimelineRecord, error) {
	return ports.TimelineRecord{}, d.err
}

func (d disabledTimelineRepository) ListEntriesByGame(ctx context.Context, gameID string) ([]ports.TimelineRecord, error) {
	return nil, d.err
}

type disabledInterrogationRepository struct {
	err error
}

func (d disabledInterrogationRepository) CreateInterrogation(ctx context.Context, input ports.InterrogationRecordInput) (ports.InterrogationRecord, error) {
	return ports.InterrogationRecord{}, d.err
}

func (d disabledInterrogationRepository) GetInterrogationByID(ctx context.Context, id string) (ports.InterrogationRecord, error) {
	return ports.InterrogationRecord{}, d.err
}

func (d disabledInterrogationRepository) ListInterrogationsByGame(ctx context.Context, gameID string) ([]ports.InterrogationRecord, error) {
	return nil, d.err
}

func main() {
	// Carregar variables d'entorn des de .env.local (o .env)
	// El fitxer ha d'estar a l'arrel del projecte (../ des de backend/)
	err := godotenv.Load("../.env.local")
	if err != nil {
		// Si no es pot carregar .env.local, intentar .env
		err = godotenv.Load(".env")
		if err != nil {
			log.Printf("⚠️  Warning: No s'ha pogut carregar .env.local o .env: %v", err)
			log.Println("⚠️  Utilitzant variables d'entorn del sistema")
		} else {
			log.Println("✅ Carregat .env (a l'arrel del projecte)")
		}
	} else {
		log.Println("✅ Carregat .env.local (a l'arrel del projecte)")
	}

	// Carregar configuracio
	cfg, cfgErr := config.Load()
	if cfgErr != nil {
		log.Fatalf("❌ Error configuracio: %v", cfgErr)
	}

	// Inicialitzar Sentry per error tracking
	sentryErr := sentry.Init(sentry.ClientOptions{
		Dsn:         os.Getenv("SENTRY_DSN"),
		Environment: cfg.Environment,
		// Sample Rate (10% de traces)
		TracesSampleRate: 0.1,
	})
	if sentryErr != nil {
		log.Printf("⚠️  Sentry init failed: %v", sentryErr)
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

	var pocketBaseClient ports.PocketBaseClient
	var gameRepository ports.GameRepository
	var playerRepository ports.PlayerRepository
	var eventRepository ports.EventRepository
	var clueRepository ports.ClueRepository
	var personRepository ports.PersonRepository
	var hypothesisRepository ports.HypothesisRepository
	var accusationRepository ports.AccusationRepository
	var forensicRepository ports.ForensicRepository
	var timelineRepository ports.TimelineRepository
	var interrogationRepository ports.InterrogationRepository
	var lobbyService ports.LobbyService
	pbClient, pbErr := repo_pb.NewClient(repo_pb.Config{
		BaseURL: cfg.PocketBaseURL,
		Timeout: cfg.PocketBaseTimeout,
	})
	if pbErr != nil {
		logger.Warn("PocketBase client disabled", "error", pbErr)
		pocketBaseClient = disabledPocketBaseClient{err: pbErr}
		gameRepository = disabledGameRepository{err: pbErr}
		playerRepository = disabledPlayerRepository{err: pbErr}
		eventRepository = disabledEventRepository{err: pbErr}
		clueRepository = disabledClueRepository{err: pbErr}
		personRepository = disabledPersonRepository{err: pbErr}
		hypothesisRepository = disabledHypothesisRepository{err: pbErr}
		accusationRepository = disabledAccusationRepository{err: pbErr}
		forensicRepository = disabledForensicRepository{err: pbErr}
		timelineRepository = disabledTimelineRepository{err: pbErr}
		interrogationRepository = disabledInterrogationRepository{err: pbErr}
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
	lobbyService = services.NewLobbyService(gameRepository, playerRepository)

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

	r.Get("/api/health", apihttp.NewHealthHandler(pocketBaseClient))
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

	// ===============================
	// DEBUG SENTRY CONFIGURATION
	// ===============================

	// Endpoint per verificar que Sentry està ben configurat
	r.Get("/api/test-sentry/debug", func(w http.ResponseWriter, r *http.Request) {
		debugInfo := map[string]interface{}{
			"dsn_configured": false,
			"dsn_value":      "",
			"environment":    "",
			"sentry_init":    false,
		}

		// Check DSN
		dsn := os.Getenv("SENTRY_DSN")
		if dsn != "" {
			debugInfo["dsn_configured"] = true
			debugInfo["dsn_value"] = dsn[:min(len(dsn), 30)] + "..." // Mostrar només els primers 30 caràcters
		}

		// Check Environment
		env := os.Getenv("ENVIRONMENT")
		if env != "" {
			debugInfo["environment"] = env
		}

		// Check si Sentry està inicialitzat
		// Això és un check simple - no podem veure l'estat intern de Sentry
		debugInfo["sentry_init"] = (dsn != "")

		web.RespondJSON(w, http.StatusOK, debugInfo)
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

	logger.Info("🚀 Servidor escoltant", "port", cfg.Port, "url", "http://localhost:"+cfg.Port)

	listenErr := http.ListenAndServe(":"+cfg.Port, r)
	if listenErr != nil {
		logger.Error("❌ Error fatal al servidor", "error", listenErr)
		os.Exit(1)
	}
}
