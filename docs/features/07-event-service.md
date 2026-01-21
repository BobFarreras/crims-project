# 07 - Event Service (Backend)

## Objectiu
Afegir la capa de dades i servei per gestionar Events (timeline) per partida.

## Abast
- Repositori `EventRepository` (ports + adapter PocketBase).
- Servei `EventService` amb validacions.
- Endpoints HTTP per crear i llistar events per partida.

## Dades
`Event`:
- `id`
- `gameId`
- `timestamp`
- `locationId`
- `participants` (llista d'IDs)

## Requeriments
1. `CreateEvent` valida camps obligatoris.
2. `GetEventByID` valida `id`.
3. `ListEventsByGame` valida `gameId`.

## Criteris d'Aceptacio
- POST retorna 201 amb event creat.
- GET per id retorna 200.
- GET per gameId retorna llista.

## Pla de Tests (TDD)
1. Repo: create OK / error.
2. Repo: list per game OK.
3. Service: invalid inputs.
4. Handlers: status codes i JSON correctes.
