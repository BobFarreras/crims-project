# 16 - Auth & Roles (Backend)

## Objectiu
Definir un flux d'autenticacio i rols per controlar permisos amb cookie segura.

## Abast
- Middleware d'autenticacio per extreure userId i role des de la cookie `auth_token`.
- Servei `AuthService` per validar tokens (JWT de PocketBase).
- Endpoints per validar sessio (opcional).

## Decisions
- Auth principal via PocketBase (token JWT).
- Backend valida token i extreu rol.
- La UI no rep el token; només el backend el guarda en cookie HttpOnly.

## Requeriments
1. Middleware valida token (cookie `auth_token`).
2. Si token invalid → 401.
3. Context inclou `userId` i `role`.

## Criteris d'Aceptacio
- Request sense token → 401.
- Token valid → request continua amb context.

## Pla de Tests (Integracio)
1. Middleware sense token → 401.
2. Middleware token invalid → 401.
3. Middleware token valid → context amb userId/role.
