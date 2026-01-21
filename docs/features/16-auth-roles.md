# 16 - Auth & Roles (Backend)

## Objectiu
Definir un flux d'autenticacio i rols per controlar permisos.

## Abast
- Middleware d'autenticacio per extreure userId i role.
- Servei `AuthService` per validar tokens.
- Endpoints per validar sessio (opcional).

## Decisions
- Auth principal via PocketBase (token JWT).
- Backend valida token i extreu rol.

## Requeriments
1. Middleware valida token (header Authorization: Bearer).
2. Si token invalid → 401.
3. Context inclou `userId` i `role`.

## Criteris d'Aceptacio
- Request sense token → 401.
- Token valid → request continua amb context.

## Pla de Tests (TDD)
1. Middleware sense token → 401.
2. Middleware token invalid → 401.
3. Middleware token valid → context amb userId/role.
