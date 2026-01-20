# GAME MECHANICS (User Experience) 🕵️‍♂️

## 1. Core Loop (El Flux del Jugador)
El joc no és lineal, el jugador salta entre aquestes fases:

### A. Exploració (L'Escena)
* **Objectiu:** Trobar Hotspots.
* **Accions:** "Mirar", "Recollir" (Inventory), "Eina Forense" (UV/Lupa).
* **Feedback:** Pistes entren a l'inventari com a `DISCOVERED`.

### B. Anàlisi (El Laboratori)
* **Objectiu:** Convertir `DISCOVERED` -> `ANALYZED`.
* **Mecànica:** Minijocs o temps d'espera on s'extreuen "EvidenceFacts" (ex: trobar una empremta parcial).
* **Resultat:** La pista guanya fiabilitat i revela dades ocultes.

### C. Deducció (El Tauler)
* **Objectiu:** Connectar punts.
* **Mecànica:** Drag & Drop.
    * Connectar Pista ↔ Sospitós.
    * Crear Hipòtesi (Node agrupador).
* **Feedback:** Visualització de fils (Vermell=Contradicció, Verd=Suport).

### D. Interrogatori (La Confrontació)
* **Mecànica:** Arbre de diàleg amb estat d'ànim (`Stress` meter).
* **Acció "Press":** Presentar una prova que contradiu el testimoni.
* **Resultat:** Si l'estrès puja massa, el testimoni pot tancar-se o confessar (Breakdown).

## 2. Rols Multijugador (Co-op Asimètric)
Cada jugador té superpoders únics:

| Rol | Habilitat Especial | Bonus |
| :--- | :--- | :--- |
| **Detectiu de Camp** | Veu Hotspots ocults a l'escena 3D | Velocitat exploració |
| **Forense** | Pot fer l'acció "Analitzar" al Lab | +Fiabilitat pistes |
| **Analista** | Pot crear Hipòtesis al Tauler | Detecta contradiccions auto. |
| **Interrogador** | Desbloqueja opcions de diàleg "Pressió" | Detecta mentides (Stress) |

## 3. Flux Final (Acusació)
Per guanyar, cal omplir el formulari final:
1.  **Qui:** Sospitós.
2.  **Per què:** Mòbil.
3.  **Amb què:** Arma/Prova clau.
4.  **Quan:** Timeline coherent.