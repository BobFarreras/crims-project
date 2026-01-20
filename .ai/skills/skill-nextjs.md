# SKILL: Next.js & React Development
**Trigger:** Quan hagis de crear pàgines, components o lògica de client.

1.  **Server vs Client Components:**
    * Per defecte, tot és Server Component.
    * Usa `'use client'` només si necessites `useState`, `useEffect` o events del navegador (clicks).

2.  **Imatges:**
    * Usa sempre `<Image />` de Next.js.
    * Configura `unoptimized` si venen d'una IP externa sense domini (temporal).

3.  **Estructura de Carpetes (App Router):**
    * `app/game/page.tsx` -> La ruta `/game`.
    * `app/game/layout.tsx` -> El layout compartit.