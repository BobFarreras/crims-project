# 05 - Game Service (Backend)

## Objectiu
Afegir una capa de servei per orquestrar la logica de creacio i consulta de partides, separant domini i adaptadors.

## Abast
- Servei `GameService` a `internal/services`.
- Validacio basica d'inputs.
- Delegacio a `ports.GameRepository`.

## Requeriments
1. `CreateGame` valida `code`, `state`, `seed`.
2. `GetGameByID` valida `id`.
3. `GetGameByCode` valida `code`.
4. Errors clars per inputs invalids.

## Criteris d'Aceptacio
- Entrades invalides → error immediat (sense cridar repositori).
- Entrades valides → crida repositori i retorna resultat.

## Pla de Tests (TDD)
1. **CreateGame invalid** → error per `code/state/seed` buits.
2. **CreateGame OK** → delega al repositori.
3. **GetGameByID invalid** → error.
4. **GetGameByCode invalid** → error.
