# 17 - Capability Permissions (Backend)

## Objectiu
Aplicar permisos per rol a endpoints sensibles.

## Abast
- Middleware `RequireCapability`.
- Context amb `role` (extret a Auth middleware).
- Errors 403 quan rol no autoritzat.

## Requeriments
1. `RequireCapability(capabilities...)` valida capacitats en context.
2. Si no hi ha rol → 403.
3. Si rol no permès → 403.

## Criteris d'Aceptacio
- Rol permès → request continua.
- Rol no permès → 403.

## Pla de Tests (Integracio)
1. Sense capacitat → 403.
2. Capacitat no permesa → 403.
3. Capacitat permesa → 200.
