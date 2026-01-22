# 23 - Integration Tests (PocketBase)

## Objectiu
Validar que el backend funciona contra PocketBase real.

## Abast
- Tests d'integracio amb dades reals.
- Usa `PB_URL` i `PB_ADMIN_TOKEN`.

## Requeriments
1. Si no hi ha envs, fer `t.Skip`.
2. Crear un `game` real i recuperar-lo.
3. Netejar el registre creat.

## Criteris d'Aceptacio
- Test passa amb PocketBase configurat.
- No trenca entorns sense envs.
