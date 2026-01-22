package main

import (
	"context"
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

// --- MOCKS FOR DISABLED POCKETBASE (Mantinguts per seguretat) ---
type disabledPocketBaseClient struct{ err error }

func (d disabledPocketBaseClient) Ping(ctx context.Context) error { return d.err }
func (d disabledPocketBaseClient) CreateUser(username, email, password, passwordConfirm, name string) error {
	return d.err
} // NOU MÈTODE
func (d disabledPocketBaseClient) AuthWithPassword(identity, password string) (*ports.AuthResponse, error) {
	return nil, d.err
}

type disabledGameRepository struct{ err error }

func (d disabledGameRepository) CreateGame(ctx context.Context, input ports.GameRecordInput) (ports.GameRecord, error) {
	return ports.GameRecord{}, d.err
}
func (d disabledGameRepository) GetGameByID(ctx context.Context, id string) (ports.GameRecord, error) {
	return ports.GameRecord{}, d.err
}
func (d disabledGameRepository) GetGameByCode(ctx context.Context, code string) (ports.GameRecord, error) {
	return ports.GameRecord{}, d.err
}

type disabledPlayerRepository struct{ err error }

func (d disabledPlayerRepository) CreatePlayer(ctx context.Context, input ports.PlayerRecordInput) (ports.PlayerRecord, error) {
	return ports.PlayerRecord{}, d.err
}
func (d disabledPlayerRepository) GetPlayerByID(ctx context.Context, id string) (ports.PlayerRecord, error) {
	return ports.PlayerRecord{}, d.err
}
func (d disabledPlayerRepository) ListPlayersByGame(ctx context.Context, gameID string) ([]ports.PlayerRecord, error) {
	return nil, d.err
}

type disabledEventRepository struct{ err error }

func (d disabledEventRepository) CreateEvent(ctx context.Context, input ports.EventRecordInput) (ports.EventRecord, error) {
	return ports.EventRecord{}, d.err
}
func (d disabledEventRepository) GetEventByID(ctx context.Context, id string) (ports.EventRecord, error) {
	return ports.EventRecord{}, d.err
}
func (d disabledEventRepository) ListEventsByGame(ctx context.Context, gameID string) ([]ports.EventRecord, error) {
	return nil, d.err
}

type disabledClueRepository struct{ err error }

func (d disabledClueRepository) CreateClue(ctx context.Context, input ports.ClueRecordInput) (ports.ClueRecord, error) {
	return ports.ClueRecord{}, d.err
}
func (d disabledClueRepository) GetClueByID(ctx context.Context, id string) (ports.ClueRecord, error) {
	return ports.ClueRecord{}, d.err
}
func (d disabledClueRepository) ListCluesByGame(ctx context.Context, gameID string) ([]ports.ClueRecord, error) {
	return nil, d.err
}

type disabledPersonRepository struct{ err error }

func (d disabledPersonRepository) CreatePerson(ctx context.Context, input ports.PersonRecordInput) (ports.PersonRecord, error) {
	return ports.PersonRecord{}, d.err
}
func (d disabledPersonRepository) GetPersonByID(ctx context.Context, id string) (ports.PersonRecord, error) {
	return ports.PersonRecord{}, d.err
}
func (d disabledPersonRepository) ListPersonsByGame(ctx context.Context, gameID string) ([]ports.PersonRecord, error) {
	return nil, d.err
}

type disabledHypothesisRepository struct{ err error }

func (d disabledHypothesisRepository) CreateHypothesis(ctx context.Context, input ports.HypothesisRecordInput) (ports.HypothesisRecord, error) {
	return ports.HypothesisRecord{}, d.err
}
func (d disabledHypothesisRepository) GetHypothesisByID(ctx context.Context, id string) (ports.HypothesisRecord, error) {
	return ports.HypothesisRecord{}, d.err
}
func (d disabledHypothesisRepository) ListHypothesesByGame(ctx context.Context, gameID string) ([]ports.HypothesisRecord, error) {
	return nil, d.err
}

type disabledAccusationRepository struct{ err error }

func (d disabledAccusationRepository) CreateAccusation(ctx context.Context, input ports.AccusationRecordInput) (ports.AccusationRecord, error) {
	return ports.AccusationRecord{}, d.err
}
func (d disabledAccusationRepository) GetAccusationByID(ctx context.Context, id string) (ports.AccusationRecord, error) {
	return ports.AccusationRecord{}, d.err
}
func (d disabledAccusationRepository) ListAccusationsByGame(ctx context.Context, gameID string) ([]ports.AccusationRecord, error) {
	return nil, d.err
}

type disabledForensicRepository struct{ err error }

func (d disabledForensicRepository) CreateAnalysis(ctx context.Context, input ports.ForensicRecordInput) (ports.ForensicRecord, error) {
	return ports.ForensicRecord{}, d.err
}
func (d disabledForensicRepository) GetAnalysisByID(ctx context.Context, id string) (ports.ForensicRecord, error) {
	return ports.ForensicRecord{}, d.err
}
func (d disabledForensicRepository) ListAnalysesByGame(ctx context.Context, gameID string) ([]ports.ForensicRecord, error) {
	return nil, d.err
}

type disabledTimelineRepository struct{ err error }

func (d disabledTimelineRepository) CreateEntry(ctx context.Context, input ports.TimelineRecordInput) (ports.TimelineRecord, error) {
	return ports.TimelineRecord{}, d.err
}
func (d disabledTimelineRepository) GetEntryByID(ctx context.Context, id string) (ports.TimelineRecord, error) {
	return ports.TimelineRecord{}, d.err
}
func (d disabledTimelineRepository) ListEntriesByGame(ctx context.Context, gameID string) ([]ports.TimelineRecord, error) {
	return nil, d.err
}

type disabledInterrogationRepository struct{ err error }

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
	// 1. Càrrega d'entorn
	err := godotenv.Load("../.env.local")
	if err != nil {
		err = godotenv.Load(".env")
		if err != nil {
			log.Printf("⚠️  Warning: No s'ha pogut carregar .env: %v", err)
		} else {
			log.Println("✅ Carregat .env")
		}
	} else {
		log.Println("✅ Carregat .env.local")
	}

	// 2. Configuració
	cfg, cfgErr := config.Load()
	if cfgErr != nil {
		log.Fatalf("❌ Error configuracio: %v", cfgErr)
	}

	// 3. Sentry
	sentryErr := sentry.Init(sentry.ClientOptions{
		Dsn:              os.Getenv("SENTRY_DSN"),
		Environment:      cfg.Environment,
		TracesSampleRate: 0.1,
	})
	if sentryErr != nil {
		log.Printf("⚠️  Sentry init failed: %v", sentryErr)
	} else {
		log.Println("✅ Sentry inicialitzat correctament")
	}
	defer sentry.Flush(2 * time.Second)

	logger := middleware.SetupLogger()
	logger.Info("🔌 Inicialitzant Crims de Mitjanit Backend...")

	// 4. Repositoris
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
		lobbyService            ports.LobbyService
	)

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

	// 5. Serveis
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

	// 🔥 NOU: Inicialitzar AuthHandler
	authHandler := apihttp.NewAuthHandler(pocketBaseClient)

	// 6. Router
	r := chi.NewRouter()

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

	// 🔥 IMPORTANT: REGISTRAR RUTES D'AUTH AQUÍ
	// Això farà que /api/auth/register funcioni
	apihttp.RegisterAuthRoutes(r, authHandler)

	// Rutes de l'API v1
	apihttp.RegisterAPIV1Routes(r, func(r chi.Router) {
		r.Get("/status", func(w http.ResponseWriter, r *http.Request) {
			web.RespondJSON(w, http.StatusOK, map[string]string{
				"system": "Crims Backend",
				"status": "healthy",
			})
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
	})

	// Debug Sentry
	r.Get("/api/test-sentry/debug", func(w http.ResponseWriter, r *http.Request) {
		web.RespondJSON(w, http.StatusOK, map[string]bool{"sentry_active": os.Getenv("SENTRY_DSN") != ""})
	})

	logger.Info("🚀 Servidor escoltant", "port", cfg.Port, "url", "http://localhost:"+cfg.Port)
	if err := http.ListenAndServe(":"+cfg.Port, r); err != nil {
		logger.Error("❌ Error fatal al servidor", "error", err)
		os.Exit(1)
	}
}
