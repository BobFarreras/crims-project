# Feature 03: Scene Exploration
$f03 = @"
# Feature 03: Scene Exploration 🔍
## 1. User Story
**Com a** Detectiu de Camp,
**Vull** explorar una escena 3D/2D i clicar en Hotspots,
**Per tal de** trobar pistes amagades.

## 2. Criteris d'Acceptació
- [ ] Navegació per l'escena (Pan/Zoom).
- [ ] Els Hotspots canvien de cursor al passar per sobre.
- [ ] Clicar un Hotspot afegeix la pista a l'inventari (`DISCOVERED`).
- [ ] El rol 'Detectiu' veu hotspots que altres no veuen.

## 3. Disseny Tècnic
* **Frontend:** Imatge interactiva amb coordenades absolutes (%).
* **Data:** JSON `scene_config` amb llista de `{id, x, y, required_role}`.

## 4. TDD Plan
- [ ] `TestCollectClue_AddsToInventory`
- [ ] `TestGetScene_FiltersHotspotsByRole`
"@
Set-Content -Path "docs\features\03-scene-exploration.md" -Value $f03 -Encoding UTF8