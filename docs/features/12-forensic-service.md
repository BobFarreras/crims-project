# 12 - Forensic Service (Backend)

## Objectiu
Gestionar l'analisi forense de pistes per partida.

## Abast
- Repositori `ForensicRepository` (ports + adapter PocketBase).
- Servei `ForensicService` amb validacions.
- Endpoints HTTP per crear i llistar analisis per partida.

## Dades
`ForensicAnalysis`:
- `id`
- `gameId`
- `clueId`
- `result` (string)
- `confidence` (0-100)
- `status` (PENDING | DONE)

## Requeriments
1. `CreateAnalysis` valida camps obligatoris.
2. `GetAnalysisByID` valida `id`.
3. `ListAnalysesByGame` valida `gameId`.

## Criteris d'Aceptacio
- POST retorna 201 amb analisi creada.
- GET per id retorna 200.
- GET per gameId retorna llista.

## Pla de Tests (TDD)
1. Repo: create OK / error.
2. Repo: list per game OK.
3. Service: invalid inputs.
4. Handlers: status codes i JSON correctes.
