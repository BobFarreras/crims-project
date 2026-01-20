# PROJECT CONTEXT: CRIMS DE MITJANIT 🕵️‍♂️ (ROOT)

## Visió General
Plataforma de joc interactiu d'investigació criminal multijugador en temps real.
L'objectiu és oferir una experiència immersiva (PWA) amb sincronització, narrativa generada per IA i multimèdia.

## 🛑 PROTOCOL STRICTE: TDD & DOCS
Abans de generar qualsevol codi d'implementació, has de seguir aquest ordre:
1.  **PHASE 1 - DOC:** Existeix el fitxer `/docs/features/X.md`? Si no, crea'l seguint `skill-documentation.md`.
2.  **PHASE 2 - TEST:** Crea el fitxer de test (`_test.go` o `.test.tsx`) basat en la documentació. Verifica que falla.
3.  **PHASE 3 - CODE:** Ara sí, genera el codi per passar el test.

**Si l'usuari demana codi directament, ATURA'T i demana permís per crear primer el pla de tests.**

## Arquitectura del Sistema
El projecte és un **Monorepo** dividit en dues àrees clares:

1.  **FRONTEND (`/frontend`)**:
    * Stack: Next.js 15, Tailwind, Framer Motion.
    * Responsabilitat: UI, Animacions, Àudio (Howler.js), PWA logic.
    * **⚠️ NORMA:** Si la tasca és visual o d'interacció usuari, LLEGEIX `/frontend/AGENTS.md`.

2.  **BACKEND (`/backend`)**:
    * Stack: Go (Golang), WebSockets/HTTP.
    * Responsabilitat: Lògica de joc, Connexió amb IA (OpenAI), Gestió d'estats.
    * **⚠️ NORMA:** Si la tasca és de lògica, dades o servidors, LLEGEIX `/backend/AGENTS.md`.

3.  **DADES**:
    * Stack: PocketBase (al VPS).
    * Responsabilitat: Auth, Persistència de dades, Fitxers.

## Flux de Treball (Skills Triggers)
Quan facis una tasca, verifica si s'aplica alguna d'aquestes habilitats i llegeix-la:
* **Crear rutes o components visuals?** -> Llegeix `.ai/skills/skill-nextjs.md`
* **Crear endpoints o lògica de servidor?** -> Llegeix `.ai/skills/skill-golang.md`
* **Modificar la Base de Dades?** -> Llegeix `.ai/skills/skill-pocketbase.md`

## Normes Globals
* **Idioma:** Tot el codi en Anglès. Comentaris i documentació en Català o Anglès.
* **URLs:**
    * Frontend Prod: `https://www.crimsdemitjanit.com`
    * Backend API: `https://api.digitaistudios.com`
    * PocketBase: `https://sspb.digitaistudios.com`

## 🎮 Game Design Context
El projecte no és una web estàtica, és un **Sistema de Deducció basat en Grafs**.
* **Core Loop:** Explorar -> Connectar Nodes (Tauler) -> Validar Hipòtesis -> Acusar.
* **Multiplayer:** Co-op asimètric (Rols: Forense, Detectiu, Analista, Interrogador).
* **AI:** Actua com a "Dungeon Master" assistit (genera flavor text, però no decideix la veritat lògica).

## 🗺️ Mapa de Context (On he d'anar?)
Llegeix el fitxer indicat segons la tasca que hagis de fer:

| Tipus de Tasca | Fitxer de Context (LLEGEIX-ME) |
| :--- | :--- |
| **Frontend / UI / PWA** | `frontend/AGENTS.md` |
| **Backend / API / Lògica** | `backend/AGENTS.md` |
| **Mecàniques & Regles de Joc** | `docs/architecture/game-mechanics.md` |
| **Motor Lògic & Estats** | `docs/architecture/game-logic-engine.md` |
| **Arquitectura Tècnica** | `.ai/context/architecture.md` |