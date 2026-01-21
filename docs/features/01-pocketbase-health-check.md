# 01 - PocketBase Health Check (Backend)

## Objectiu
Exposar un endpoint de salut per validar la connexio amb PocketBase en runtime.

## Abast
- Endpoint HTTP `/api/health`.
- Retorna estat del sistema i estat de PocketBase.
- Usa el port `PocketBaseClient` per desacoblament.

## Requeriments
1. Handler a `internal/adapters/http`.
2. Resposta JSON consistent.
3. Si PocketBase no respon, retornar `503`.

## Criteris d'Aceptacio
- Si `Ping()` retorna OK, resposta `200` amb `{ status: "healthy", pocketbase: "ok" }`.
- Si `Ping()` falla, resposta `503` amb error JSON i codi `pocketbase_unavailable`.

## Pla de Tests (TDD)
1. **Ping OK** → status 200 + JSON correcte.
2. **Ping error** → status 503 + ErrorResponse amb codi.
