# 17 - Role Permissions (Backend)

## Objectiu
Aplicar permisos per rol a endpoints sensibles.

## Abast
- Middleware `RequireRole`.
- Context amb `role` (extret a Auth middleware).
- Errors 403 quan rol no autoritzat.

## Requeriments
1. `RequireRole(roles...)` valida rol en context.
2. Si no hi ha rol → 403.
3. Si rol no permès → 403.

## Criteris d'Aceptacio
- Rol permès → request continua.
- Rol no permès → 403.

## Pla de Tests (TDD)
1. Sense rol → 403.
2. Rol no permès → 403.
3. Rol permès → 200.
