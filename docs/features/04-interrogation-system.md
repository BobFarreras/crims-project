# Feature 04: Interrogation System
$f04 = @"
# Feature 04: Interrogation System 🗣️
## 1. User Story
**Com a** Interrogador,
**Vull** parlar amb un sospitós i presentar proves,
**Per tal de** detectar contradiccions i fer-lo confessar.

## 2. Criteris d'Acceptació
- [ ] Arbre de diàleg navegable.
- [ ] Opció 'Present Evidence' obre l'inventari.
- [ ] Si la prova contradiu l'afirmació -> `Stress` puja.
- [ ] Si `Stress` > Threshold -> Desbloqueja 'Breakdown'.

## 3. Disseny Tècnic
* **Model:** `DialogueNode` { text, responses[], stress_effect }.
* **Logic:** `CheckContradiction(statementID, evidenceID)`.

## 4. Integration Plan
- [ ] `TestPresentEvidence_ValidContradiction_IncreasesStress`
- [ ] `TestDialogue_LockedOptions_UnlockWithFacts`
"@
Set-Content -Path "docs\features\04-interrogation-system.md" -Value $f04 -Encoding UTF8