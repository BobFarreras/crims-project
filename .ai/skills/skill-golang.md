# SKILL: Go Backend Development
**Trigger:** Quan hagis de crear lògica de backend o APIs.

1.  **JSON Responses:**
    * Usa sempre `json.NewEncoder(w).Encode(data)`.
    * Defineix les structs amb tags: `json:"nom_camp"`.

2.  **Injecció de Dependències:**
    * No usis globals. Passa la connexió de DB o el client AI a través dels structs del servei.

3.  **Noms:**
    * Usa `CamelCase` per a coses exportades (públiques).
    * Usa `camelCase` per a coses privades.