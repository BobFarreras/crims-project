# 24 - Frontend Lobby UI

## Objectiu
Crear una pantalla inicial per unir-se a una partida amb codi i rol.

## Abast
- Pantalla inicial `/` amb formulari.
- Camps: game code, role selector, user id (placeholder).
- Botó "Join" sense crida real (mock per ara).

## Requeriments
1. UI responsive (mobile + desktop).
2. Formulari amb validacio bàsica.
3. Components a `frontend/features/lobby`.

## Criteris d'Aceptacio
- Renderitza títol, inputs i botó.
- Valida que el codi no estigui buit.

## Pla de Tests (Integracio)
1. Renderitza els camps principals.
2. Mostra error quan el codi es buit.
