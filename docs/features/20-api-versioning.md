# 20 - API Versioning

## Objectiu
Versionar l'API per garantir compatibilitat amb clients actuals i futurs.

## Abast
- Prefix `/api/v1` per totes les rutes publiques.
- Helper de registre de rutes versionades.

## Requeriments
1. Totes les rutes existents passen a `/api/v1`.
2. Sense prefix, la ruta retorna 404.

## Criteris d'Aceptacio
- `/api/v1/health` respon.
- `/api/health` no existeix.

## Pla de Tests (TDD)
1. Registrar rutes amb helper afegeix prefix `/api/v1`.
