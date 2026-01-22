# 15 - Lobby Service (Backend)

## Objectiu
Gestionar l'entrada i preparacio de jugadors dins una partida.

## Abast
- Servei `LobbyService` per unir jugadors, assignar rol i iniciar partida.
- Reutilitza `GameRepository` i `PlayerRepository`.
- Endpoints HTTP per unir-se a la partida i llistar jugadors.

## Dades
`LobbyJoin`:
- `gameCode`
- `userId`
- `role`

## Requeriments
1. `JoinGame` valida codi i userId.
2. `JoinGame` crea player associat a game.
3. `ListPlayers` valida `gameId`.

## Criteris d'Aceptacio
- POST retorna 201 amb player creat.
- GET retorna llista de players de la partida.

## Pla de Tests (Integracio)
1. Service: invalid input.
2. Service: join OK delega a repos.
3. Handler: status codes correctes.
