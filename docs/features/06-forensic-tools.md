# Feature 06: Forensic Tools
$f06 = @"
# Feature 06: Forensic Tools 🔬
## 1. User Story
**Com a** Forense,
**Vull** aplicar filtres (UV, Espectrograma) a una pista `DISCOVERED`,
**Per tal de** extreure'n informació oculta (`ANALYZED`).

## 2. Criteris d'Acceptació
- [ ] UI de Laboratori amb eines seleccionables.
- [ ] Eina correcta sobre pista correcta -> Èxit.
- [ ] Eina incorrecta -> Feedback 'No s'ha trobat res'.

## 3. Disseny Tècnic
* **Frontend:** Canvas filters (CSS brightness/contrast/invert).
* **Backend:** `AnalyzeClue(clueID, toolID)` -> actualitza estat.

## 4. TDD Plan
- [ ] `TestAnalyze_CorrectTool_UpdatesState`
- [ ] `TestAnalyze_WrongTool_NoChange`
"@
Set-Content -Path "docs\features\06-forensic-tools.md" -Value $f06 -Encoding UTF8