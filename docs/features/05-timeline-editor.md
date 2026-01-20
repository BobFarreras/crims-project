# Feature 05: Timeline Editor
$f05 = @"
# Feature 05: Timeline Editor ⏱️
## 1. User Story
**Com a** Analista,
**Vull** ordenar events en una línia temporal,
**Per tal de** trobar forats en les coartades.

## 2. Criteris d'Acceptació
- [ ] Drag & Drop d'events a slots horaris.
- [ ] Validació visual: Si dos events del mateix sospitós se solapen en llocs diferents -> Error (Vermell).
- [ ] Gate: No pots acusar sense timeline coherent.

## 3. Disseny Tècnic
* **Backend:** Validació lògica `ValidateTimeline(events[])`.

## 4. TDD Plan
- [ ] `TestTimeline_Overlap_ReturnsConflict`
- [ ] `TestTimeline_ImpossibleTravelTime`
"@
Set-Content -Path "docs\features\05-timeline-editor.md" -Value $f05 -Encoding UTF8