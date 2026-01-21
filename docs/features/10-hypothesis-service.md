# 10 - Hypothesis Service (Backend)

## Objectiu
Gestionar hipotesis del tauler de deduccio per partida.

## Abast
- Repositori `HypothesisRepository` (ports + adapter PocketBase).
- Servei `HypothesisService` amb validacions.
- Endpoints HTTP per crear i llistar hipotesis per partida.

## Dades
`Hypothesis`:
- `id`
- `gameId`
- `title`
- `strengthScore`
- `status`
- `nodeIds` (llista d'IDs)

## Requeriments
1. `CreateHypothesis` valida camps obligatoris.
2. `GetHypothesisByID` valida `id`.
3. `ListHypothesesByGame` valida `gameId`.

## Criteris d'Aceptacio
- POST retorna 201 amb hipotesi creada.
- GET per id retorna 200.
- GET per gameId retorna llista.

## Pla de Tests (TDD)
1. Repo: create OK / error.
2. Repo: list per game OK.
3. Service: invalid inputs.
4. Handlers: status codes i JSON correctes.
