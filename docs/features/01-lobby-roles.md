# Feature 01: Lobby & Roles
$f01 = @"
# Feature 01: Lobby & Role Selection 👥
## 1. User Story
**Com a** jugador,
**Vull** unir-me a una sala mitjançant un codi i triar capacitats (Forense, Analista...),
**Per tal de** jugar cooperativament sense límit d'un sol rol.

## 2. Criteris d'Acceptació
- [ ] Es pot crear una sala i retorna un codi de 4 lletres.
- [ ] Altres jugadors poden unir-se amb el codi.
- [ ] Les capacitats es poden repetir si hi ha més jugadors (assistents).
- [ ] Un jugador pot tenir més d'una capacitat.
- [ ] La partida comença quan l'Host prem 'Start'.

## 3. Disseny Tècnic
* **DB:** Col·lecció `games` (code, status) i `players` (capabilities[], game_id).
* **Realtime:** Subscripció a `players(game_id)` per veure qui entra.

## 4. Integration Plan
- [ ] `TestCreateLobby_ReturnsCode`
- [ ] `TestJoinLobby_Success`
- [ ] `TestJoinLobby_AllowsMultipleCapabilities`
"@
Set-Content -Path "docs\features\01-lobby-roles.md" -Value $f01 -Encoding UTF8
