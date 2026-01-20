---
name: TDD Protocol
trigger: writing new logic, fixing bugs, creating endpoints
scope: global
---

# TDD PHILOSOPHY (RED -> GREEN -> REFACTOR)
**Norma Absoluta:** No escriguis codi d'implementació si no hi ha un test que falli primer.

## 1. Backend (Go)
* **Eina:** `testing` (nativa) + `testify` (asserts).
* **Mocking:** Utilitza interfícies per a la Base de Dades. Mai testegis contra PocketBase real en Unit Tests.
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
* **Eina:** Jest + React Testing Library.
* **Focus:** Testeja comportament, no detalls d'estil.
* **Ubicació:** Carpeta `__tests__` dins de la feature o fitxer `.test.tsx` al costat del component.