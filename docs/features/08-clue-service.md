# 08 - Clue Service (Backend)

## Objectiu
Afegir la capa de dades i servei per gestionar Clues (pistes) per partida.

## Abast
- Repositori `ClueRepository` (ports + adapter PocketBase).
- Servei `ClueService` amb validacions.
- Endpoints HTTP per crear i llistar clues per partida.

## Dades
`Clue`:
- `id`
- `gameId`
- `type`
- `state`
- `reliability` (0-100)
- `facts` (JSON)

## Requeriments
1. `CreateClue` valida camps obligatoris.
2. `GetClueByID` valida `id`.
3. `ListCluesByGame` valida `gameId`.

## Criteris d'Aceptacio
- POST retorna 201 amb clue creada.
- GET per id retorna 200.
- GET per gameId retorna llista.

## Pla de Tests (TDD)
1. Repo: create OK / error.
2. Repo: list per game OK.
3. Service: invalid inputs.
4. Handlers: status codes i JSON correctes.
