# CRIMS de Mitjanit 🕵️‍♂️

[![CI/CD](https://github.com/BobFarreras/crims-project/actions/workflows/ci.yml/badge.svg)](https://github.com/BobFarreras/crims-project/actions/workflows/ci.yml)

Plataforma de joc d'investigació criminal multijugador en temps real (PWA). Els jugadors col·laboren per resoldre misteris connectant pistes, interrogant sospitosos i deduint la veritat.

## 📋 Visió Ràpida

- **Frontend:** Next.js 15 + Tailwind CSS + Framer Motion
- **Backend:** Go (Golang) + Chi Router
- **Database:** PocketBase (BaaS)
- **Architecture:** Monorepo amb arquitectura hexagonal
- **Testing:** Vitest (frontend) + Go Test (backend) + Playwright (E2E)
- **Deployment:** Vercel (frontend) + Docker VPS (backend)

## 🎮 Mecànica de Joc

### Core Loop
1. **Exploració:** Descobrir pistes a l'escena (Hotspots)
2. **Anàlisi:** Processar evidències al laboratori
3. **Deducció:** Connectar pistes al tauler (React Flow)
4. **Interrogatori:** Confrontar sospitosos amb proves
5. **Acusació:** Formular la teoria final i guanyar

### Rols Multijugador
- 🔍 **Detectiu de Camp:** Troba pistes ocultes
- 🔬 **Forense:** Analitza evidències al lab
- 📊 **Analista:** Crea i valida hipòtesis al tauler
- 🎤 **Interrogador:** Detecta mentides i pressiona testimonis

## 🚀 Setup Ràpid

### Prerequisits
- Node.js 20+
- Go 1.25.6
- Docker & Docker Compose
- pnpm 9

### Instal·lació

```bash
# Clonar repositori
git clone https://github.com/BobFarreras/crims-project.git
cd crims-project

# Instal·lar dependències frontend
cd frontend && pnpm install && cd ..

# (Opcional) Setup backend Go (automàtic amb Docker)
cd backend && go mod download && cd ..

# Iniciar serveis locals
docker-compose up -d
```

### Variables d'Entorn

Copia `.env.example` a `.env.local` i configura:

```bash
cp .env.example .env.local
```

**IMPORTANTE:** Usa `.env.local` para desarrollo local (este archivo se ignora en git). Para producción, configura las variables en el hosting (Vercel/Docker).

Variables clau:
- `NEXT_PUBLIC_API_URL` - URL del backend API
- `NEXT_PUBLIC_POCKETBASE_URL` - URL de PocketBase
- `JWT_SECRET` - Secret per JWT authentication
- `OPENAI_API_KEY` - Clau per integració IA (opcional)

## 🏃 Comandes

### Frontend (Next.js)

```bash
cd frontend

# Development
pnpm dev

# Build
pnpm build

# Test
pnpm test
pnpm test:ui

# Lint
pnpm lint
```

### Backend (Go)

```bash
cd backend

# Development
go run ./cmd/server

# Build
go build -o ./bin/server ./cmd/server

# Test
go test ./...
```

### Monorepo

```bash
# Iniciar tots els serveis (Docker)
docker-compose up -d

# Aturar tots els serveis
docker-compose down

# Veure logs
docker-compose logs -f

# Tests units (frontend + backend)
make test-unit

# Tests E2E
make test-e2e
```

## 📁 Estructura del Projecte

```
crims-project/
├── frontend/              # Next.js PWA
│   ├── app/              # App Router
│   ├── features/         # Feature modules
│   ├── components/       # Shared UI components
│   └── lib/              # Utilities & API clients
├── backend/              # Go API
│   ├── cmd/              # Entry points
│   ├── internal/         # Hexagonal architecture
│   │   ├── domain/       # Business logic
│   │   ├── ports/        # Interfaces
│   │   ├── adapters/     # Implementations
│   │   └── services/     # Application logic
│   └── go.mod
├── docs/                 # Documentation
│   ├── architecture/     # System design
│   ├── features/         # Feature specs (TDD)
│   └── deployment.md     # Deployment guide
├── .github/              # CI/CD workflows
├── .ai/                  # AI agent context & skills
└── tests/                # E2E tests
```

## 🧠 Arquitectura

### Frontend (Feature-based)
Cada feature és un mòdul autònom amb components, lògica i tests:
- `lobby/` - Selecció de rols
- `board/` - Tauler de deducció
- `scene/` - Exploració 3D
- `interrogation/` - Diàlegs
- `timeline/` - Editor temporal
- `forensic/` - Eines de laboratori
- `accusation/` - Formulari final

### Backend (Hexagonal)
Separació clara de responsabilitats:
- **Domain:** Regles de negoci pures
- **Ports:** Interfícies (repo, services)
- **Adapters:** Implementacions (HTTP, DB, AI)
- **Middleware:** Cross-cutting concerns (auth, logging)

## 📚 Documentació

- [Deployment Guide](./docs/deployment.md) - Desplegament a producció
- [Git Workflow](./docs/git-workflow.md) - Estratègia de branches (Git Flow)
- [Sentry Setup](./docs/sentry-setup.md) - Configuració d'error tracking
- [Agent Safety](./docs/agent-safety.md) - Mesures de seguretat per agents AI
- [Game Mechanics](./docs/architecture/game-mechanics.md) - Mecàniques de joc
- [Game Logic Engine](./docs/architecture/game-logic-engine.md) - Motor lògic
- [Project Structure](./docs/architecture/project-structure.md) - Estructura detallada
- [Project Phases](./docs/architecture/project-phases.md) - Guia pas a pas del roadmap
- [Features](./docs/features/) - Especificacions de cada feature (TDD)

## 🔧 Desenvolupament

### Git Workflow

Aquest projecte utilitza **Git Flow** adaptat per a monorepos:

```
main              → Producció (sempre estable)
└── develop       → Integració (pre-producció)
    ├── feature/* → Noves funcionalitats
    ├── release/* → Preparació de versions
    ├── hotfix/*  → Correccions urgents
    ├── chore/*   → Tasques tècniques
    └── docs/*    → Documentació
```

**Flux típic:**
1. `git checkout develop && git pull origin develop`
2. `git checkout -b feature/feature-name`
3. Treballar + commitear
4. Crear PR (feature → develop)
5. Merge aprovat + branca esborrada

Veure [Git Workflow](./docs/git-workflow.md) per detalls complets.

### Workflow TDD

1. **Doc:** Crea documentació a `docs/features/X.md`
2. **Test:** Escriu test basat en la doc
3. **Code:** Implementa per passar el test

### Commit Conventions

```
feat: nova funcionalitat
fix: correcció de bug
refactor: refactoring de codi
test: afegir tests
docs: documentació
chore: tasques de manteniment
```

## 🚢 Deployment

### Producció
- **Frontend:** https://www.crimsdemitjanit.com (Vercel)
- **Backend:** https://api.digitaistudios.com (VPS)
- **Database:** https://sspb.digitaistudios.com (VPS)

Veure [Deployment Guide](./docs/deployment.md) per detalls complets.

## 🛡️ Seguretat

- Zero Trust: Validació de tots els inputs
- RBAC: Control d'accés per rols
- JWT Authentication
- Rate limiting a endpoints crítics
- Sanitització d'inputs (XSS prevention)
- Environment variables sensibles no commitejades

Veure [`skill-security.md`](./.ai/skills/skill-security.md) per més detalls.

## 🤝 Contribució

1. Fork el repositori
2. Crea branca de feature: `git checkout -b feature/amazing-feature`
3. Commit canvis: `git commit -m 'feat: add amazing feature'`
4. Push a branca: `git push origin feature/amazing-feature`
5. Obre Pull Request

## 📝 Llicència

Copyright © 2025 DigiTai Studios. Tots els drets reservats.

## 👥 Equip

- **DigiTai Studios** - Development Team

## 📞 Suport

- GitHub Issues: https://github.com/BobFarreras/crims-project/issues
- Email: dev@digitaistudios.com

---

**Built with ❤️ for mystery lovers**
