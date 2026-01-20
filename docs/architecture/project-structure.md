# PROJECT STRUCTURE MAP 🗺️

## BACKEND (Go) - Arquitectura Hexagonal
### L'objectiu és aïllar el domini (regles del joc) de la tecnologia (HTTP, PocketBase). 

```text
/backend
├── /cmd/server
│   └── main.go              # Entry point (Injecta dependències i arrenca)
│
├── /internal
│   ├── /domain              # ⚠️ NUCLI PUR (Business Logic & Models)
│   │   ├── models.go        # Structs (Clue, Game, Player)
│   │   ├── errors.go        # Errors de negoci (ErrGameFull, ErrUnauthorized)
│   │   └── logic.go         # Algorismes purs (Scoring, Truth tables)
│   │
│   ├── /ports               # ⚠️ CONTRACTES (Interfaces)
│   │   ├── repository.go    # DB Interfaces (GameRepo, PlayerRepo)
│   │   └── service.go       # Logic Interfaces (GameService)
│   │
│   ├── /services            # 🧠 LÒGICA D'APLICACIÓ (Use Cases)
│   │   ├── game_service.go  # Orquestració (ValidateMove, CheckGates)
│   │   └── ai_service.go    # Prompt Engineering logic
│   │
│   ├── /middleware          # 🛡️ SEGURETAT (Cross-cutting concerns)
│   │   ├── auth.go          # Validació JWT/Session & RBAC
│   │   └── logger.go        # Request logging & Monitoring
│   │
│   └── /adapters            # 🔌 CONECTORS (Implementacions)
│       ├── /http            # API Handlers (Gin/StdLib)
│       │   ├── router.go    # Mapeig de rutes
│       │   └── handlers.go  # Serialització JSON i Validació Input
│       │
│       ├── /repo_pb         # PocketBase Driver
│       │   ├── client.go    # Client HTTP intern
│       │   └── game_repo.go # Implementació de ports.GameRepository
│       │
│       └── /ai_openai       # OpenAI Driver

```

## FRONTEND (Next.js) - Feature-based
### Organitzem el codi per funcionalitat (Feature), no per tipus tècnic.
/frontend
├── /app                     # Routing Layer (App Router)
│   ├── /(game)              # Layout de Joc (PWA Style - No UI browser)
│   │   ├── lobby/page.tsx
│   │   ├── board/page.tsx
│   │   └── ...
│   ├── /(auth)              # Layout d'Autenticació
│   │   └── login/page.tsx
│   └── /api                 # API Proxy (Opcional)
│
├── /features                # 🧩 MÒDULS AUTÒNOMS (Coherent amb docs/features/)
│   ├── /board               # Feature 02
│   │   ├── /components      # UI específica (Nodes, Canvas)
│   │   ├── /logic           # React Flow hooks & State
│   │   └── /__tests__       # Unit Tests (Jest/Vitest)
│   ├── /scene               # Feature 03
│   │   ├── /components      # (Viewport3D, InventoryBar)
│   │   └── /__tests__
│   ├── /lobby               # Feature 01
│   ├── /interrogation       # Feature 04
│   ├── /timeline            # Feature 05
│   ├── /forensic            # Feature 06
│   └── /accusation          # Feature 07
│
├── /lib                     # Shared Kernel
│   ├── /core                # Utilities pures i Types globals
│   └── /infra               # Clients externs
│       ├── api-client.ts    # Fetch wrapper amb Auth Header
│       └── pocketbase.ts    # Auth wrapper (Login/Logout logic)

## GLOBAL & TESTS
## On viuen els tests d'integració i la configuració.

/
├── Makefile                 # Automatització (run-all, test-unit)
├── AGENTS.md                # Orquestrador d'IA
├── /tests
│   └── /e2e                 # Playwright/Cypress Specs (Tests Integració)
│       └── game_flow.spec.ts
├── /.ai                     # Context i Skills per a la IA
└── /docs                    # Documentació viva del projecte