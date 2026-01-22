# 35 - Role Capabilities Model

## Objectiu
Permetre partides de 1 a 8 jugadors utilitzant capacitats de rol (no rols unics) per evitar bloqueig de joc i permetre mode solo.

## User Story
**Com a** jugador,
**Vull** jugar una partida amb 1, 2 o mes jugadors sense perdre habilitats clau,
**Per tal de** completar el cas sense esperes innecessaries.

## Criteris d'Aceptacio
- Un jugador pot tenir multiples capacitats (ex: Detectiu + Analista).
- Les accions especials depenen de capacitats, no del rol unic.
- El lobby permet assignar capacitats en modes 1/2/4/8.
- El backend valida accions per capacitats (no per rol unic).
- El frontend mostra accions disponibles segons capacitats.

## Disseny Tecnic
- **Model Player:** canviar `role` per `capabilities[]` (string array).
- **Lobby:** assignacio flexible segons mode de partida.
- **Middleware/Guardes:** `RequireCapability(...)` en lloc de `RequireRole(...)`.
- **UI:** els botons d'accio comproven capacitats (ex: `canAnalyze`).

## Pla de Tests (Integracio)
1. **Lobby assigna capacitats:** crear partida i assignar 4 capacitats a 1 jugador.
2. **Accio permesa:** jugador amb capacitat fa accio especial (200).
3. **Accio denegada:** jugador sense capacitat rep 403.
4. **Mode 8 jugadors:** 4 principals + assistents sense capacitats especials.
