# 40 - Pantalla Lobby

## Objectiu
Mostrar un lobby pre-joc amb el codi de sala, jugadors connectats i accions d'espera abans d'entrar al cas.

## Criteris d'Aceptacio
- URL ` /lobby/{code}` mostra el codi de sala i estat de partida.
- Si el codi no existeix, es mostra error clar.
- En mode Solo, es navega directament a `/game`.
- En mode Duo/Equip, es navega a `/lobby/{code}`.

## Disseny Tecnic
- **Frontend:** `LobbyScreen` com a component client amb `lobbyService.getGameByCode`.
- **Backend:** es reutilitza `GET /api/games/by-code/{code}`.
- **Navegacio:** `DashboardHub` redirigeix segons mode.

## Pla de Tests (Integracio)
1. **Dashboard create solo:** redirigeix a `/game`.
2. **Dashboard create multi:** redirigeix a `/lobby/{code}`.
3. **Lobby screen:** mostra el codi i estat quan `getGameByCode` respon.
4. **Lobby screen error:** mostra error si codi no existeix.
