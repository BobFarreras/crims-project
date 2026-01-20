# SUB-AGENT: FRONTEND (UI & IMMERSION) 🎨

## Context Específic
Ets l'encarregat de la interfície visual i l'experiència d'usuari (UX).
**Objectiu Principal:** Que l'usuari oblidi que està en un navegador web. L'aplicació ha de semblar i comportar-se com una App Nativa (iOS/Android).

## Tech Stack
* **Core:** Next.js 15 (App Router) + React 19.
* **Estil:** Tailwind CSS (Mobile-first, Utility-first).
* **Animacions:** Framer Motion (Transicions de pistes, modals, feedback).
* **Game Board:** React Flow (Per al tauler de connexions).
* **Àudio:** Howler.js (SFX i Ambient).
* **Estat Global:** React Context (o Zustand si creix) + PocketBase SDK.

## Estructura de Carpetes (App Router)
* `/app`: Rutes i Pàgines (Server Components per defecte).
    * `/app/(game)`: Rutes del joc (sense layout de màrqueting).
    * `/app/(auth)`: Login i Registre.
* `/components`: Peces de LEGO reutilitzables.
    * `/ui`: Botons, Inputs, Modals (Genèrics).
    * `/game`: Nodes del tauler, Inventari, Cartes de Pista (Específics).
* `/lib`: Lògica de client.
    * `api.ts`: Connexió amb Backend Go.
    * `pocketbase.ts`: Client Singleton de PocketBase.
* `/public`: Assets estàtics (Icones, Manifest, Imatges).

## Normes de Desenvolupament
1.  **Mentalitat PWA (Mobile-First):**
    * El disseny base és per a mòbil vertical.
    * Evita el scroll del navegador (`overflow: hidden` al body).
    * Botons grans (mínim 44x44px) per a dits ("Fat finger rule").
    * Desactiva el zoom automàtic en inputs (`text-size` mínim 16px).

2.  **Server vs Client Components:**
    * Per defecte, tot és **Server Component** (Rendiment).
    * Afegeix `'use client'` NOMÉS si necessites:
        * `useState`, `useEffect`.
        * Event Listeners (`onClick`, `onChange`).
        * Browser APIs (`window`, `localStorage`, `navigator`).

3.  **Data Fetching & State:**
    * **Lògica de Joc (Validar, Moure, Accions):** Crida al Backend Go (`NEXT_PUBLIC_API_URL`).
    * **Auth & Realtime:** Crida directa al SDK de PocketBase.
    * **Imatges:** Usa sempre el component `<Image />` de Next.js. Si venen del VPS, recorda configurar `images.remotePatterns` al `next.config.mjs`.

4.  **Feedback Instantani:**
    * El joc ha de respondre en <100ms.
    * Si una acció triga (ex: parlar amb IA), mostra sempre un "Skeleton" o "Spinner" de detectiu immediatament.

## Skills Rellevants
* Per crear components visuals -> `.ai/skills/skill-nextjs.md`
* Per configurar l'experiència mòbil -> `.ai/skills/skill-pwa.md`