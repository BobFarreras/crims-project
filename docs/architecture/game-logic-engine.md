# GAME LOGIC ENGINE (System Architecture) ⚙️

## 1. State Machine (Màquina d'Estats Global)
Variable: `CASE_STATE` (Enum).

1.  `BRIEFING`: Generació de `SeedLeads`.
2.  `INVESTIGATION`: Estat obert (Loop principal).
    * Permet accés a rutes `/scene`, `/board`, `/lab`.
3.  `ACCUSATION_PHASE`:
    * Trigger: Es compleix `GATE_ACCUSATION`.
    * Efecte: Read-only en pistes. Només permet votacions.
4.  `RESOLUTION`: Càlcul de score i fi de partida.

## 2. Data Models (Structs & DB Schema)

### A. Entities
* **Clue:** `{ id, type, reliability (0-100), state (DISCOVERED/ANALYZED), facts[] }`
* **Person:** `{ id, officialStory, truthStory, stress (0-100), credibility }`
* **Event:** `{ id, timestamp, locationId, participants[] }`
* **Player:** `{ id, userId, gameId, capabilities[], status, isHost }`

### B. Logic Structures (Nodes de Deducció)
* **Hypothesis:**
    * `strengthScore`: Calculat dinàmicament (`suport - contradiccions`).
    * `status`: Enum (`WEAK` < 20, `PLAUSIBLE` < 50, `SOLID` > 50).
* **Contradiction:**
    * `severity`: Enum (`MINOR`, `MAJOR`, `CRITICAL`).
    * `resolved`: Bool.

## 3. Progression System (Gating Logic)
El backend avalua aquestes regles a cada acció (`CheckGates()`):

* **GATE_SCENE_2:**
    * `IF count(AnalyzedClues) >= 5 AND count(Contradictions) >= 1`
    * `THEN Unlock(Location_2)`
* **GATE_ACCUSATION:**
    * `IF count(SolidHypothesis) >= 1 AND count(CriticalContradictions) == 0`
    * `THEN AllowState(ACCUSATION_PHASE)`

## 4. AI Constraint Rules (Anti-Hallucination)
La IA (LLM) s'invoca només com a transformador de text.
* **Context:** Injectar sempre `KnownFacts` i `CharacterProfile`.
* **Límit:** Mai pot inventar un `NewFact` que no estigui a la DB.

## 5. Multiplayer Sync Logic
* **State Source:** PocketBase és la única font de veritat (Single Source of Truth).
* **Conflict Resolution:** Last-write-wins per a moviments de tauler. Votació majoritària per a canvis de fase (`NextPhase`).
