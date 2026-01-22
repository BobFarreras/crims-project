# 38 - Perfil jugador (Nom)

## Objectiu
Permetre que el jugador editi el seu nom al Dashboard i guardar-lo a PocketBase.

## Criteris d'Aceptacio
- `GET /api/profile` retorna `{ user }` amb `name`.
- `PUT /api/profile` amb `name` actualitza el nom del jugador.
- Si falta `name`, retorna `400 auth/missing_profile_name`.
- Si no hi ha cookie, retorna `401 auth/missing_session`.

## Disseny Tecnic
- **Backend:** nou handler `ProfileHandler` amb `GetProfile` i `UpdateProfile`.
- **PocketBase:** actualitzar el camp `name` del record `users` autenticat.
- **Frontend:** `profileService` + `DashboardHub` carrega el nom i permet guardar.

## Pla de Tests (Integracio)
1. **Backend update ok:** login + `PUT /api/profile` retorna user amb `name` actualitzat.
2. **Backend missing name:** `PUT /api/profile` amb `name` buit retorna `400`.
3. **Frontend save:** click "Guardar" crida `profileService.updateProfile`.
