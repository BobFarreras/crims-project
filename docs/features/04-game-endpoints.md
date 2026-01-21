# 04 - Game Endpoints (Backend)

## Objectiu
Exposar endpoints minimals per crear i obtenir partides.

## Endpoints
1. **POST /api/games**
   - Input: `{ code, state, seed }`
   - Output: Game record

2. **GET /api/games/{id}**
   - Output: Game record

3. **GET /api/games/by-code/{code}**
   - Output: Game record

## Requeriments
- Handlers a `internal/adapters/http`.
- Validacio basica d'input.
- Errors JSON consistents (`web.RespondError`).

## Criteris d'Aceptacio
- POST retorna 201 amb JSON de la partida creada.
- GET per id retorna 200.
- GET per code retorna 200 o 404 si no existeix.

## Pla de Tests (TDD)
1. **POST OK** → 201 + JSON.
2. **POST invalid** → 400.
3. **GET id OK** → 200 + JSON.
4. **GET code not found** → 404.
