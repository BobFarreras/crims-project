# SYSTEM ARCHITECTURE

## Entorns
* **Producció (VPS):**
    * Frontend: Vercel (`www.crimsdemitjanit.com`)
    * Backend API: VPS (`api.digitaistudios.com`) -> Docker Container Go
    * DB: VPS (`sspb.digitaistudios.com`) -> Docker Container PocketBase
* **Local:**
    * Frontend: `localhost:3000`
    * Backend: `localhost:8080`
    * DB: `localhost:8090`

## Flux de Dades Multijugador
1.  Usuari fa acció -> Frontend envia POST a Backend Go.
2.  Backend Go valida i actualitza PocketBase.
3.  PocketBase dispara event Realtime.
4.  Tots els Frontends reben l'event i s'actualitzen.