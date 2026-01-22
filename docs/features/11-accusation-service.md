# 11 - Accusation Service (Backend)

## Objectiu
Gestionar el flux d'acusacio final per partida.

## Abast
- Repositori `AccusationRepository` (ports + adapter PocketBase).
- Servei `AccusationService` amb validacions.
- Endpoints HTTP per crear i obtenir acusacions per partida.

## Dades
`Accusation`:
- `id`
- `gameId`
- `playerId`
- `suspectId`
- `motiveId`
- `evidenceId`
- `verdict` (string, opcional)

## Requeriments
1. `CreateAccusation` valida camps obligatoris.
2. `GetAccusationByID` valida `id`.
3. `ListAccusationsByGame` valida `gameId`.

## Criteris d'Aceptacio
- POST retorna 201 amb acusacio creada.
- GET per id retorna 200.
- GET per gameId retorna llista.

## Pla de Tests (TDD)
1. Repo: create OK / error.
2. Repo: list per game OK.
3. Service: invalid inputs.
4. Handlers: status codes i JSON correctes.
