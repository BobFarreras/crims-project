# Feature 07: Accusation Flow
$f07 = @"
# Feature 07: Accusation Flow ⚖️
## 1. User Story
**Com a** equip,
**Vull** presentar les meves conclusions finals,
**Per tal de** tancar el cas i veure si hem guanyat.

## 2. Criteris d'Acceptació
- [ ] Bloquejat fins que `GATE_ACCUSATION` sigui true.
- [ ] Formulari: Qui, Mòbil, Arma.
- [ ] Feedback final amb puntuació.

## 3. Disseny Tècnic
* **Backend:** `SubmitAccusation()` compara amb `CaseTruth`.
* **Scoring:** Algorisme de precisió (Pistes útils vs Soroll).

## 4. Integration Plan
- [ ] `TestAccusation_Gate_Locked_If_Hypothesis_Weak`
- [ ] `TestScore_Calculation_PerfectMatch`
"@
Set-Content -Path "docs\features\07-accusation-flow.md" -Value $f07 -Encoding UTF8