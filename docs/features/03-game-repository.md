# 03 - Game Repository (PocketBase)

## Objectiu
Crear el repositori de `Game` per persistir i recuperar partides a PocketBase, desacoblat via ports.

## Abast
- Interface `GameRepository` al port.
- Implementacio PocketBase a `adapters/repo_pb`.
- Operacions minimals: `CreateGame`, `GetGameByID`, `GetGameByCode`.

## Dades
`Game` (record):
- `id` (string)
- `code` (string, unique)
- `state` (string)
- `seed` (string)

## Requeriments
1. `CreateGame` crea record a PocketBase.
2. `GetGameByID` recupera per `id`.
3. `GetGameByCode` filtra per `code` (unique).
4. Errors clars per status no esperats.

## Criteris d'Aceptacio
- POST correcte retorna `Game` amb `id`.
- GET per id retorna `Game`.
- GET per code retorna `Game` o error si no existeix.

## Pla de Tests (TDD)
1. **CreateGame OK** → POST amb payload correcte i resposta 200.
2. **CreateGame HTTP 500** → retorna error `ErrUnexpectedStatus`.
3. **GetGameByID OK** → GET correcte i parse de resposta.
4. **GetGameByCode OK** → GET amb filtre correcte.
