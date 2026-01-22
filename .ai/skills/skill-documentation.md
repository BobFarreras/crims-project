---
name: Feature Documentation
trigger: starting new feature, planning architecture
scope: global
---

# DOCUMENTATION FIRST STRATEGY

Abans de demanar codi, crea un fitxer de especificació a `/docs/features/NOM-FEATURE.md`.

## Plantilla de Feature (`FEATURE_TEMPLATE.md`)
1.  **User Story:** "Com a [usuari], vull [acció] per a [benefici]".
2.  **Criteris d'Acceptació:** Llista de punts exactes per considerar la feina acabada (Checklist).
3.  **Technical Design:**
    * Endpoint (API): `POST /api/game/join`
    * Dades (JSON): `{ "code": "XY99" }`
    * Canvis a BD: Taula `games`, camp `players`.
4.  **Casos de Test (Integration Plan):**
    * [ ] Cas feliç (tot va bé).
    * [ ] Cas error (codi incorrecte).
    * [ ] Cas límit (sala plena).
