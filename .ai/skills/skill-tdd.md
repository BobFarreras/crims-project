---
name: Integration-First Protocol
trigger: writing new logic, fixing bugs, creating endpoints
scope: global
---

# INTEGRATION-FIRST (RED -> GREEN -> REFACTOR)
**Norma Absoluta:** No escriguis codi d'implementació si no hi ha un test d'integracio que falli primer.

## 1. Backend (Go)
* **Eina:** `testing` (nativa).
* **Integracio:** Prioritza tests contra adaptadors reals (HTTP/PocketBase) amb env vars. Evita mocks si el test valida contractes.
* **Ubicació:** El test va al costat del fitxer.
    * Codi: `internal/game/service.go`
    * Test: `internal/game/service_test.go`
* **Format:**
    ```go
    func TestCreateGame(t *testing.T) {
        // 1. Arrange (Prepara dades)
        // 2. Act (Executa funció)
        // 3. Assert (Verifica resultat)
    }
    ```

## 2. Frontend (Next.js)
* **Eina:** Vitest + React Testing Library.
* **Focus:** Testeja fluxos complets (rutes, components i serveis), no detalls d'estil.
* **Ubicació:** Carpeta `__tests__` dins de la feature o fitxer `.test.tsx` al costat del component.
