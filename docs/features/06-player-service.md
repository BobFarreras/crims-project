# 06 - Player Service (Backend)

## Objectiu
Afegir una capa de servei per crear i consultar players dins una partida.

## Abast
- Repositori `PlayerRepository` (ports + adapter PocketBase).
- Servei `PlayerService` (validacions).
- Endpoints HTTP per crear i obtenir players.

## Dades
`Player`:
- `id`
- `gameId`
- `userId`
- `role`
- `status`
- `isHost`

## Requeriments
1. `CreatePlayer` valida camps obligatoris.
2. `GetPlayerByID` valida `id`.
3. `ListPlayersByGame` valida `gameId`.
4. Errors clars per inputs invalids.

## Criteris d'Aceptacio
- POST retorna 201 amb player creat.
- GET per id retorna 200.
- GET per gameId retorna llista.

## Pla de Tests (TDD)
1. Repo: create OK / error.
2. Repo: list per game OK.
3. Service: invalid inputs.
4. Handlers: status codes i JSON correctes.
