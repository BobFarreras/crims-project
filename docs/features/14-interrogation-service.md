# 14 - Interrogation Service (Backend)

## Objectiu
Gestionar sessions d'interrogatori i respostes per partida.

## Abast
- Repositori `InterrogationRepository` (ports + adapter PocketBase).
- Servei `InterrogationService` amb validacions.
- Endpoints HTTP per crear i llistar interrogatoris per partida.

## Dades
`Interrogation`:
- `id`
- `gameId`
- `personId`
- `question`
- `answer`
- `tone` (string)

## Requeriments
1. `CreateInterrogation` valida camps obligatoris.
2. `GetInterrogationByID` valida `id`.
3. `ListInterrogationsByGame` valida `gameId`.

## Criteris d'Aceptacio
- POST retorna 201 amb interrogatori creat.
- GET per id retorna 200.
- GET per gameId retorna llista.

## Pla de Tests (TDD)
1. Repo: create OK / error.
2. Repo: list per game OK.
3. Service: invalid inputs.
4. Handlers: status codes i JSON correctes.
