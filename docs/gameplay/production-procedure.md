# Procediment de Produccio del Joc

## 1. Objectiu
Definir un procediment clar per construir el joc com una app mobile-first, amb UI animada, experient 3D i llogica de joc escalable.

## 2. Prioritat de desenvolupament (ordre de produccio)
1. **Dashboard d'entrada (Hub)**
   - Seleccio de mode: Solo, Duo, Equip, Sala existent.
   - Perfil: nom del jugador, avatar, preferencies basiques.
   - Ranking i historials.
   - Entrada rapida a casos actius.
2. **Lobby i Matchmaking**
   - Crear sala, unir-se amb codi, assignar capacitats.
   - Estat de preparacio i start.
3. **Core Loop jugable**
   - Escena (exploracio 3D).
   - Laboratori (analisi).
   - Tauler (deduccio).
   - Interrogatori.
4. **Acusacio i resolucio**
   - Formulari final, puntuacio, feedback.
5. **Metajoc**
   - Progressio, achievements, estadistiques, rankings.

## 3. Dashboard (Hub) - Quins blocs ha de tenir
- **Mode de joc**: botons clars per Solo, Duo, Equip, Unir-se a sala.
- **Casos actius**: targetes amb estat, temps restant, participants.
- **Perfil**: nom editable, avatar, id curt.
- **Ranking**: top setmanal i personal best.
- **Accesos rapids**: continuar ultima partida, veure tutorial, configuracio.

## 4. Disseny mobile-first i app-like
- Layout en targetes grans, una columna en mobil.
- Accions principals sempre visibles a la part baixa.
- Animacions suaus (transicions de panell, cards, hover tactil).
- Estat clar: loading, sincronitzacio, errors.

## 5. Contingut 3D (escena)
- **Motor recomanat**: Three.js + React Three Fiber.
- **Pipeline d'assets**: models low-poly, textures lleugeres, hotspots amb glow.
- **Interaccio**: tocs curts per seleccionar, llarg per accions.
- **Optimitzacio**: LOD, textures compresses, limit de poligons.

## 6. Procediment de produccio (fases)
1. **Spec**
   - Documentacio a `docs/features/XX.md`.
   - Criteris d'acceptacio.
2. **Tests d'integracio**
   - Crear tests abans d'implementar.
   - Validar contractes API i UI.
3. **Implementacio**
   - Backend: serveis i adaptadors.
   - Frontend: components, rutes i estats.
4. **Verificacio**
   - `go test ./...` i `pnpm test`.
   - `pnpm build` per UI.
5. **Iteracio**
   - Feedback intern.
   - Ajustos visuals i de UX.

## 7. Decisions pendents
- Definir rules de PocketBase per crear games i players en test.
- Definir format de ranking (global, setmanal, per rol/capacitat).
- Definir presets de modes (solo/duo/equip) i repartiment de capacitats.
