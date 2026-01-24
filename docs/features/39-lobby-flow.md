# 39 - Lobby Flow (Crear i Unir-se)

## Objectiu
Permetre crear una sala amb codi curt i unir-s'hi des del dashboard, amb un flux simple per Solo/Duo/Equip.

## Criteris d'Aceptacio
- `POST /api/lobby/create` retorna `{ game, player }` amb codi de 4 lletres.
- `POST /api/lobby/join` permet unir-se amb `gameCode` + `userId`.
- El dashboard permet crear sala i unir-se amb codi.
- Si falta `gameCode` o `userId`, retorna `400`.

## Disseny Tecnic
- **Backend:** `LobbyService.CreateLobby` crea game `LOBBY` i host player.
- **Frontend:** `lobbyService` consumeix create/join i el dashboard mostra el codi.

## Pla de Tests (Integracio)
1. **Backend create lobby:** retorna codi de 4 lletres i host `isHost=true`.
2. **Backend join lobby:** crea player amb `capabilities`.
3. **Frontend create lobby:** clicar mode crida `lobbyService.createLobby`.
4. **Frontend join lobby:** clicar "Unir-se" crida `lobbyService.joinLobby`.
