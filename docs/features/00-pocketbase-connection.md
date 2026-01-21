# 00 - PocketBase Connection (Backend)

## Objectiu
Definir una connexio robusta cap a PocketBase via HTTP API, amb separacio clara entre domini i infraestructura, pensada per a futurs canvis (Supabase).

## Abast
- Backend Go: client HTTP per PocketBase.
- Configuracio via variables d'entorn.
- Errors manejats de forma explicita.
- No inclou endpoints de joc encara, nomes la capa d'acces.

## Requeriments
1. Client PocketBase encapsulat a `adapters/repo_pb`.
2. Interface de repositori a `internal/ports` per desacoblar el backend del BaaS.
3. Temps d'espera configurable (timeout).
4. Logs d'errors i temps de resposta.
5. Config per entorn: `PB_URL` i `PB_ADMIN_EMAIL`/`PB_ADMIN_PASSWORD` si cal auth.

## Decisions d'Arquitectura
- L'adapter només parla HTTP (no SDK intern).
- El domini no depen de PocketBase.
- El client retornara errors tipats (no panic).

## Dades Minimes
No calen coleccions per provar la connexio. La prova es pot fer amb health check / collection list.

## Criteris d'Aceptacio
- Si `PB_URL` es buit, el client falla amb error clar.
- Si el servidor no respon, el client retorna error amb timeout.
- Si el servidor respon, el client retorna OK.
- Els logs inclouen latencia i error si n'hi ha.

## Pla de Tests (TDD)

### Unit Tests (Go)
1. **Client init sense PB_URL** → retorna error `ErrMissingPBUrl`.
2. **Health check OK** → client retorna OK amb HTTP 200 mockejat.
3. **Health check timeout** → client retorna error de timeout.
4. **HTTP 500** → client retorna error amb codi i missatge.

### Integracio (opcional)
- Si `PB_URL` apunta a PocketBase local, `Ping()` retorna OK.

## Notes de Seguretat
- No commitejar credencials.
- Usar `SENTRY_AUTH_TOKEN` i variables d'entorn per secrets.
