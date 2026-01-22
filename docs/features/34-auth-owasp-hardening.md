# 34 - Auth OWASP Hardening

## Objectiu
Endurir el flux d'autenticacio amb cookies segures, validacio de token i CORS configurables per complir OWASP.

## Abast
- Cookies d'autenticacio amb `Secure` i `SameSite=None` en entorns cross-site.
- CORS basat en `CORS_ALLOWED_ORIGINS` (no hardcoded).
- Eliminar logs de credencials en clar.
- Validar tokens contra PocketBase abans de permetre rutes protegides.
- Endpoint `/api/auth/session` per validar sessio i retornar usuari.

## Disseny Tecnic
- **Config:**
  - `ENVIRONMENT=production` activa cookie `Secure`.
  - `AUTH_COOKIE_SAMESITE=none|lax|strict` per controlar SameSite.
  - `CORS_ALLOWED_ORIGINS` ja suporta llista separada per comes.
- **Backend Auth:**
  - `AuthHandler` rep configuracio de cookie.
  - `AuthMiddleware` llegeix `auth_token` de la cookie i valida via PocketBase (`auth-refresh`).
  - `/api/auth/session` retorna `{ user }` o `401` si la sessio es invalida.
- **Frontend:**
  - `authService.login` manté `credentials: 'include'`.
  - `authService.logout` segueix POST.

## Criteris d'Aceptacio
1. Login genera cookie amb `HttpOnly`, `Secure` segons entorn i `SameSite` configurable.
2. CORS permet orígens configurats per entorn (prod + local).
3. Cap log de credencials en clar al backend.
4. `/api/auth/session` retorna usuari si el token es valid i `401` si no.

## Pla de Tests (Integracio)
1. **Login cookie flags (prod):** en `ENVIRONMENT=production`, resposta de login inclou `Secure` i `SameSite=None`.
2. **Session valid:** login + `GET /api/auth/session` retorna `200` amb user.
3. **Session invalid:** sense cookie retorna `401`.
