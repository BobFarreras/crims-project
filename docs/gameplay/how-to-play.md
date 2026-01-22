# Com es juga a Crims de Mitjanit

## 1. Visio general
Crims de Mitjanit es un joc cooperatiu de deduccio criminal. Els jugadors comparteixen pistes, construeixen hipotesis i desbloquegen fases fins arribar a l'acusacio final. L'experiencia esta pensada com una app jugable, amb UI immersiva, animacions i feedback immediat.

## 2. Entrada a la partida (Lobby)
- **Host crea sala** i rep un codi de 4 lletres.
- **Jugadors s'uneixen** amb el codi.
- **Assignacio de capacitats**: els rols son capacitats, no limitacions rigides.
- **Inici**: la partida comenca quan l'host prem Start.

## 3. Capacitats i rols
Les habilitats del joc estan agrupades en 4 capacitats. Un jugador pot tenir una o diverses capacitats segons el mode.

- **Detectiu de Camp**: explora l'escena i troba hotspots ocults.
- **Forense**: analitza pistes al laboratori per pujar fiabilitat.
- **Analista**: organitza hipotesis al tauler i detecta contradiccions.
- **Interrogador**: gestiona interrogatoris i pressio del testimoni.

Les capacitats donen **bonus** (info extra, accions exclusives o velocitat), pero tots els jugadors poden participar a totes les fases.

## 4. Modes de joc (1 a 8 jugadors)
- **Solo (1 jugador)**: el jugador te les 4 capacitats.
- **Duo (2 jugadors)**: cada jugador te 2 capacitats o un rol dual.
- **Equip (3-4 jugadors)**: capacitats repartides 1 a 1.
- **Grup gran (5-8 jugadors)**: 4 capacitats principals + assistents.

Els assistents poden fer accions universals (explorar, connectar pistes, votar) pero no tenen bonus de rol.

## 5. Flux general de partida (Core Loop)
El jugador salta entre aquests espais, segons necessitats de la investigacio:

1. **Escena (Exploracio)**
   - Objectiu: descobrir hotspots i recollir pistes.
   - Resultat: pistes en estat `DISCOVERED`.

2. **Laboratori (Analisi)**
   - Objectiu: convertir `DISCOVERED` a `ANALYZED`.
   - Resultat: nova informacio i augment de fiabilitat.

3. **Tauler (Deduccio)**
   - Objectiu: connectar pistes, persones i hipotesis.
   - Resultat: hipotesis amb força (WEAK, PLAUSIBLE, SOLID) i contradiccions visibles.

4. **Interrogatori (Confrontacio)**
   - Objectiu: exposar contradiccions i elevar l'estres del testimoni.
   - Resultat: desbloqueig de confessions o bloqueig de respostes.

## 6. Progressio i fases (State Machine)
La partida es regeix per una maquina d'estats:

- **BRIEFING**: generacio de pistes base i context.
- **INVESTIGATION**: loop principal (escena, lab, tauler, interrogatori).
- **ACCUSATION_PHASE**: quan es compleixen les condicions de gate.
- **RESOLUTION**: calcul final de puntuacio i tancament de cas.

### Gates principals
- **GATE_SCENE_2**: desbloqueja nova localitzacio quan hi ha prou pistes i contradiccions.
- **GATE_ACCUSATION**: activa fase d'acusacio quan hi ha hipotesi solida i cap contradiccio critica.

## 7. Acusacio final
El formulari final demana:
1. **Qui** (sospitos).
2. **Per que** (mobil).
3. **Amb que** (prova clau o arma).
4. **Quan** (timeline coherent).

La puntuacio depen de la coherencia global i de la qualitat de les proves.

## 8. Sincronitzacio multijugador
- **Font de veritat**: PocketBase.
- **Conflictes**: last-write-wins al tauler; votacio majoritaria per canvis de fase.
- **Temps real**: subscripcio a canvis de jugadors, pistes i estat.

## 9. IA com a narrador
La IA nomes transforma text amb les dades reals del cas. No pot inventar fets nous i sempre treballa sobre `KnownFacts` i `CharacterProfile`.

## 10. Experiencia visual i 3D
La UI ha de comportar-se com una app (mobile-first), amb animacions, transicions i feedback visual immediat.

### Proposta per escena 3D
- **Recomanat**: Three.js amb React Three Fiber per integracio directa amb Next.js.
- **Alternatives**: Babylon.js (escenes complexes), PlayCanvas (editor visual).

L'escena 3D es un espai d'exploracio on els hotspots destaquen amb llums, partícules o vibracions subtils.

## 11. Resum
El joc es un flux de deduccio iteratiu: explorar, analitzar, connectar, confrontar i acusar. Cada rol aporta una part critica de la investigacio, i el sistema de gates assegura que la narrativa avanci quan la logica es solida.
