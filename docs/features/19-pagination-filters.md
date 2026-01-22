# 19 - Pagination & Filters (Backend)

## Objectiu
Afegir suport estandard de paginacio i filtratge a les llistes.

## Abast
- Helpers per construir query params (page, perPage, filter).
- Extensio de repositoris per acceptar opcions de listat.

## Requeriments
1. `ListOptions` amb `Page`, `PerPage`, `Filter`.
2. Valors per defecte segurs.
3. Query params afegits a PocketBase.

## Criteris d'Aceptacio
- Listats sense opcions funcionen com ara.
- Amb opcions, el query inclou page/perPage/filter.

## Pla de Tests (TDD)
1. Construccio de query amb defaults.
2. Construccio de query amb params.
