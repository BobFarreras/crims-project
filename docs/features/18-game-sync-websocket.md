# 18 - Game State Sync (WebSocket)

## Objectiu
Permetre sincronitzacio en temps real de l'estat del joc.

## Abast
- Endpoint WebSocket `/ws/game/{gameId}`.
- Hub per broadcast a tots els clients de la partida.
- Missatges JSON amb tipus i payload.

## Requeriments
1. Connexio WebSocket acceptada amb `gameId`.
2. Broadcast a tots els clients del mateix `gameId`.
3. Desconnexio segura.

## Criteris d'Aceptacio
- Connecta i manté connexio.
- Broadcast arriba als clients del mateix gameId.

## Pla de Tests (Integracio)
1. Hub registra i elimina clients.
2. Broadcast envia a tots els clients del mateix gameId.
