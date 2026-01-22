# 36 - Dashboard Hub (Entrada de Joc)

## Objectiu
Crear el hub principal on el jugador tria el mode de joc, gestiona el perfil i accedeix a casos actius. Ha de ser mobile-first i app-like.

## User Story
**Com a** jugador,
**Vull** un dashboard clar per escollir mode, personalitzar el nom i veure el meu ranking,
**Per tal de** començar a jugar sense friccio.

## Criteris d'Aceptacio
- El dashboard mostra modes: Solo, Duo, Equip i Unir-se a sala.
- El jugador pot editar el seu nom i avatar basics.
- Es mostren casos actius i un botó de continuar.
- Es mostra ranking setmanal i marca personal.
- UI en una columna a mobil i dos blocs a desktop.

## Disseny Tecnic
- **Frontend:** component `DashboardHub` a `frontend/features/dashboard`.
- **Seccions:** `ModeSelector`, `ProfileCard`, `ActiveCases`, `RankingPanel`.
- **Dades:** mock local inicial + endpoint futur `GET /api/profile` i `GET /api/rankings`.

## Pla de Tests (Integracio)
1. **Render modes:** mostra botons de mode (Solo/Duo/Equip/Sala).
2. **Perfil editable:** input de nom visible i editable.
3. **Casos actius:** almenys una targeta de cas amb CTA.
4. **Ranking visible:** titol de ranking i llista d'entrades.
