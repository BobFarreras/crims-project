# Feature 02: Investigation Board (Tauler de Deducció) 🕵️‍♂️

## 1. User Story
**Com a** grup de detectius (rol Analista liderant),
**Vull** un tauler digital interactiu on organitzar Pistes (`Clues`) i crear Hipòtesis (`Hypotheses`),
**Per tal de** connectar evidències visualment i validar si la nostra teoria té prou suport o contradiccions abans d'acusar.

## 2. Criteris d'Acceptació (Checklist)
- [ ] **Nodes:** Es visualitzen Pistes (amb foto), Persones i Events Temporals.
- [ ] **Hipòtesis:** Es poden crear nodes agrupador "Hipòtesi" on arrossegar pistes a dins.
- [ ] **Connexions (Edges):**
    - Fil Verd (`SUPPORTS`): Connecta Pista -> Hipòtesi.
    - Fil Vermell (`CONTRADICTS`): Connecta Pista -> Coartada/Persona.
    - Fil Gris (`RELATED`): Neutre.
- [ ] **Calculadora de Veritat:** El tauler mostra visualment la força d'una hipòtesi (WEAK/PLAUSIBLE/SOLID) basant-se en les connexions.
- [ ] **Realtime:** Moviments i connexions es sincronitzen <100ms entre jugadors.

## 3. Disseny Tècnic

### A. Frontend (React Flow)
* **Custom Nodes:**
    * `ClueNode`: Mostra imatge i estat (Analyzed/Verified).
    * `HypothesisNode`: Contenidor (Parent Node). Canvia de color segons l'estat (Gris -> Groc -> Verd).
* **Interacció:**
    * `onConnect`: Obre un petit menú contextual per triar tipus de relació (Supports/Contradicts).
* **Capability Gating:**
    * Només la capacitat **Analyst** pot crear Hipòtesis i validar-les.
    * La resta de jugadors poden moure nodes i proposar connexions (que queden "pendents" fins que l'Analyst confirma - Opcional V2).

### B. Backend (Go API)
* **Endpoint:** `POST /api/board/connect`
    * Payload: `{ source: "clue_1", target: "hypo_A", type: "SUPPORTS" }`
    * Lògica:
        1. Crea el registre a DB `board_edges`.
        2. Recalcula l'estat de la Hipòtesi (`CalculateHypothesisStrength`).
        3. Si la força canvia (ex: passa a SOLID), emet event Realtime `HYPOTHESIS_UPDATED`.

### C. Algorisme de Força (Go Service)
```go
func CalculateStrength(hypoID string) string {
    score := 0
    // Cada pista verificada que suporta suma punts
    score += countSupportingClues(hypoID) * 10
    // Cada contradicció resta molt
    score -= countContradictions(hypoID) * 20
    
    if score > 50 { return "SOLID" }
    if score > 20 { return "PLAUSIBLE" }
    return "WEAK"
}
```

### D. Base de Dades (PocketBase)

* **Col·lecció `hypotheses`:**
    * `game_id`, `statement`, `status` (`WEAK`/`PLAUSIBLE`/`SOLID`), `score`.
* **Col·lecció `board_edges`:**
    * `source_id`, `target_id`, `type` (`SUPPORTS`/`CONTRADICTS`).

---

### 4. Pla de Tests (Integracio)

**Backend (Go)** - `internal/game/board_test.go`
- [ ] `TestConnect_Supports_IncreasesScore`: Connectar una pista a una hipòtesi ha de pujar el seu score.
- [ ] `TestConnect_Contradicts_DecreasesScore`: Una contradicció ha de baixar dràsticament el score.
- [ ] `TestStatus_Change`: Verificar que si el score passa de 20, l'estat canvia a `PLAUSIBLE` automàticament.

**Frontend (React)** - `__tests__/BoardLogic.test.tsx`
- [ ] `should render hypothesis node`: Comprovar que es pinta el contenidor.
- [ ] `should color code edges`: Verificar que `type='CONTRADICTS'` pinta la línia vermella.
