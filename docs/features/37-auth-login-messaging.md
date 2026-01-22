# 37 - Auth Login Messaging & Safety

## Objectiu
Millorar el missatge d'error de login al frontend i eliminar qualsevol log de credencials al backend.

## Criteris d'Aceptacio
- Cap log mostra contrasenyes o credencials en clar.
- El login retorna `400` amb codi `auth/missing_credentials` si falten dades.
- El frontend mostra missatges clars per errors de credencials o servidor.

## Disseny Tecnic
- **Backend:** validar `email` i `password` abans d'autenticar; eliminar `fmt.Printf` de login.
- **Frontend:** `authService.login` interpreta codi HTTP i llança errors amigables.
- **UI:** `LoginFlow` mostra el missatge retornat pel hook.

## Pla de Tests (Integracio)
1. **Backend missing credentials:** `POST /api/auth/login` amb `email` buit retorna `400` + `auth/missing_credentials`.
2. **Frontend error messaging:** quan `authService.login` falla amb `Credencials incorrectes`, es mostra el missatge.
