# 📘 Project Phases (Frontend + Backend + Data)

Guia pas a pas per construir el projecte de manera segura i incremental, seguint tests d'integracio i les normes d'AGENTS.

## 0) Preparacio
- Definir versio objectiu (ex: `v1.0.1`) i crear `release/vX.Y.Z`.
- Crear nota de release: `docs/releases/vX.Y.Z.md`.
- Actualitzar `frontend/package.json` amb la versio.
- Executar checklist local: `docs/checklists/local-ci.md`.

## 1) Fase Base (Infra + Data)
**Objectiu:** Tenir serveis minims funcionant i connectats.
- PocketBase al VPS (Portainer) com a font de veritat.
- Definir coleccions i camps basics (Clue, Person, Event, Game, Player).
- Crear una capa d'acces a dades (backend `adapters/repo_pb`).
- Definir ports de repositori per permetre canvi futur (PocketBase → Supabase).

## 2) Fase Backend Core (Go)
**Objectiu:** Motor logic i API minima.
- Definir models i errors en `backend/internal/domain`.
- Definir ports en `backend/internal/ports`.
- Implementar serveis en `backend/internal/services`.
- API HTTP basica a `backend/internal/adapters/http`.
- Contractes d'API i DTOs estables (per clients futurs).

**Integration-first obligatori:**
1. Documentar feature a `docs/features/X.md`.
2. Crear tests d'integracio (`_test.go`) i verificar que fallen.
3. Implementar fins a passar els tests.

## 3) Fase Frontend Core (Next.js)
**Objectiu:** Flux minim de joc i UI base.
- App Router amb layouts `/app/(game)` i `/app/(auth)`.
- Features modulars a `/frontend/features`.
- Integracio amb PocketBase i API (`/frontend/lib/infra`).
- Evitar logica de domini al client; el backend valida tot.

**Integration-first obligatori:**
1. Documentar feature a `docs/features/X.md`.
2. Crear tests d'integracio (`.test.tsx`) i verificar que fallen.
3. Implementar fins a passar els tests.

## 4) Fase Multiplayer + Sync
**Objectiu:** Estat compartit i temps real.
- WebSocket/HTTP al backend.
- Sincronitzacio amb PocketBase (single source of truth).
- Resolucio de conflictes (last-write-wins / votacions).
- Events i payloads versionats per compatibilitat futura.

## 5) Fase Gameplay (Loop complet)
**Objectiu:** Core loop complet (Scene → Lab → Board → Interrogation → Accusation).
- Aplicar regles de `docs/architecture/game-mechanics.md`.
- Aplicar motor logic de `docs/architecture/game-logic-engine.md`.

## 6) Fase Qualitat (QA + Stabilization)
**Objectiu:** Seguretat, tests, i estabilitat.
- Ampliar tests (integracio + unit si cal).
- Revisar lint i build.
- Revisar secrets i configuracio `.env.local`.
- Mantenir separacio de responsabilitats (ports/adapters).
- Revisar que el client sigui substituible (Kotlin Multiplatform futur).

## 7) Fase Release
**Objectiu:** Preparar versio i tancar.
- Crear `release/vX.Y.Z` desde `develop`.
- Fixes finals + documentacio de release.
- Merge a `main` (nomes usuari).

## Regles d'Agent (resum)
- Sempre nova branch per feature o logica.
- Merge només via PR a `develop`.
- L'agent no fa merge ni push a `main`.
- Les branches `release/*` no s'esborren.
