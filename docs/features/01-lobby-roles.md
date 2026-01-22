# Feature 01: Lobby & Roles
$f01 = @"
# Feature 01: Lobby & Role Selection 👥
## 1. User Story
**Com a** jugador,
**Vull** unir-me a una sala mitjançant un codi i triar un rol (Forense, Analista...),
**Per tal de** jugar cooperativament amb habilitats úniques.

## 2. Criteris d'Acceptació
- [ ] Es pot crear una sala i retorna un codi de 4 lletres.
- [ ] Altres jugadors poden unir-se amb el codi.
- [ ] No es pot repetir un rol que ja està agafat.
- [ ] La partida comença quan l'Host prem 'Start'.

## 3. Disseny Tècnic
* **DB:** Col·lecció `games` (code, status) i `players` (role, game_id).
* **Realtime:** Subscripció a `players(game_id)` per veure qui entra.

## 4. Integration Plan
- [ ] `TestCreateLobby_ReturnsCode`
- [ ] `TestJoinLobby_Success`
- [ ] `TestJoinLobby_RoleTaken_ReturnsError`
"@
Set-Content -Path "docs\features\01-lobby-roles.md" -Value $f01 -Encoding UTF8