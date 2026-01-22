# 13 - Timeline Service (Backend)

## Objectiu
Gestionar entrades de timeline per partida, basades en events i notes.

## Abast
- Repositori `TimelineRepository` (ports + adapter PocketBase).
- Servei `TimelineService` amb validacions.
- Endpoints HTTP per crear i llistar entrades per partida.

## Dades
`TimelineEntry`:
- `id`
- `gameId`
- `timestamp`
- `title`
- `description`
- `eventId` (opcional)

## Requeriments
1. `CreateEntry` valida camps obligatoris.
2. `GetEntryByID` valida `id`.
3. `ListEntriesByGame` valida `gameId`.

## Criteris d'Aceptacio
- POST retorna 201 amb entrada creada.
- GET per id retorna 200.
- GET per gameId retorna llista.

## Pla de Tests (Integracio)
1. Repo: create OK / error.
2. Repo: list per game OK.
3. Service: invalid inputs.
4. Handlers: status codes i JSON correctes.
