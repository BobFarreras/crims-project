---
name: Testing Skill
trigger: creating test files, running tests, mocking data
scope: development, testing
---

# TESTING SKILL 🧪

## 1. ¿Qué es Testing?

Testing és el procés de verificar que el codi funciona correctament segons les especificacions.

**Tipus de Testing:**
- **Unit Tests:** Proven unitats individuals (funcions, components)
- **Integration Tests:** Proven que diferents mòduls funcionen junts
- **E2E Tests:** Proven el flux complet de l'aplicació com un usuari
- **Manual Testing:** Proves manuals (QA, UAT)

---

## 2. Frontend Testing (Next.js + Vitest)

### Unit Tests amb Vitest

**Què és:** Provar components i funcions individualment.

**Comandes:**
```bash
cd frontend
pnpm test -- run    # Executar tots els tests
pnpm test -- ui      # Veure results en UI
pnpm test -- file  # Provar un sol fitxer
```

**Estructura de Test:**
```typescript
// Example: features/lobby/__tests__/roles.test.tsx
import { render, screen } from '@testing-library/react';
import { describe, it, expect, vi } from 'vitest';

describe('Lobby - Roles', () => {
  it('debería mostrar 4 roles disponibles', () => {
    const roles = ['DETECTIVE', 'FORENSIC', 'ANALISTA', 'INTERROGADOR'];
    const { getByText } = render(<LobbyRoles roles={roles} />);
    
    roles.forEach(role => {
      expect(getByText(role)).toBeInTheDocument();
    });
  });

  it('debería deshabilitar rol ya seleccionado', () => {
    const selectedRole = 'DETECTIVE';
    const { getByText } = render(<LobbyRoles selectedRole={selectedRole} />);
    
    expect(getByText('DETECTIVE')).toBeDisabled();
  });
});
```

### Component Testing amb React Testing Library

**Què és:** Provar components React simulant interaccions d'usuari.

**Comandes:**
```bash
pnpm add -D @testing-library/react @testing-library/jest-dom
```

**Best Practices:**
```typescript
// ✅ BONO: Provar comportament, no implementación
import { render, screen, fireEvent, waitFor } from '@testing-library/react';

test('usuario puede seleccionar rol', async () => {
  render(<LobbyRoles />);
  
  // Simular clic
  fireEvent.click(screen.getByRole('button', { name: /detective/i }));
  
  // Esperar actualización
  await waitFor(() => {
    expect(screen.getByText(/rol seleccionado/i)).toBeInTheDocument();
  });
});

// ❌ MALO: Probar implementación interna
test('useState funciona correctamente', () => {
  // No hacer esto! Es probar el comportamiento, no cómo funciona useState
});
```

### Integration Tests

**Què és:** Provar que diferents mòduls funcionen junts.

**Example:**
```typescript
// features/lobby/__tests__/lobby-integration.test.tsx
import { render, screen, waitFor } from '@testing-library/react';
import Lobby from '@/features/lobby';
import { PocketBaseProvider } from '@/lib/pocketbase';

describe('Lobby - Integración con PocketBase', () => {
  it('debería unirse a una sala existente', async () => {
    render(
      <PocketBaseProvider>
        <Lobby gameId="abc-123" />
      </PocketBaseProvider>
    );
    
    // Simular input de código
    fireEvent.change(screen.getByLabelText(/código de sala/i), {
      target: { value: 'ABCD' }
    });
    
    // Simular clic en unirse
    fireEvent.click(screen.getByRole('button', { name: /unirse/i }));
    
    // Esperar que el usuario se una a la sala
    await waitFor(() => {
      expect(screen.getByText(/bienvenido/i)).toBeInTheDocument();
    });
  });
});
```

---

## 3. Backend Testing (Go)

### Unit Tests

**Què és:** Provar funcions i lògica de negoci individualment.

**Estructura de Test:**
```go
// Example: internal/services/game_service_test.go
package services_test

import (
    "testing"
    "github.com/digitaistudios/crims-backend/internal/services"
    "github.com/digitaistudios/crims-backend/internal/domain"
)

func TestCreateGame_ValidInput_ReturnsGameID(t *testing.T) {
    // Arrange
    mockRepo := new MockGameRepository()
    gameService := NewGameService(mockRepo)
    
    // Act
    game, err := gameService.Create("ABCD", "LOBBY")
    
    // Assert
    if err != nil {
        t.Fatalf("Expected no error, got %v", err)
    }
    if game.Code != "ABCD" {
        t.Fatalf("Expected code ABCD, got %s", game.Code)
    }
}
```

### Table-Driven Tests

**Què és:** Definir múltiples casos de prova en una taula.

**Example:**
```go
func TestCreateGame_InvalidCodes(t *testing.T) {
    tests := []struct {
        name     string
        code     string
        expected error
    }{
        {"empty code", "", errors.ErrInvalidGameCode},
        {"too short", "AB", errors.ErrInvalidGameCode},
        {"too long", "ABCDE", errors.ErrInvalidGameCode},
        {"invalid chars", "A1C2", errors.ErrInvalidGameCode},
    }
    
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            mockRepo := new MockGameRepository()
            gameService := NewGameService(mockRepo)
            
            _, err := gameService.Create(tt.code, "LOBBY")
            
            if err != tt.expectedError {
                t.Errorf("Expected error %v, got %v", tt.expectedError, err)
            }
        })
    }
}
```

### Mocking Dependencies

**Què és:** Simular dependències com PocketBase per provar sense connexió real.

**Example:**
```go
// internal/adapters/repo/mock_game_repo_test.go
package adapters_test

import (
    "testing"
    "github.com/stretchr/testify/mock"
    "github.com/digitaistudios/crims-backend/internal/domain"
    "github.com/digitaistudios/crims-backend/internal/ports"
)

type MockGameRepository struct {
    mock.Mock
}

func (m *MockGameRepository) Create(game *domain.Game) (*domain.Game, error) {
    args := m.Called(game)
    return &domain.Game{
        ID:     "test-123",
        Code:   game.Code,
        Status:  game.Status,
    }, nil
}

func TestGameService_CreateGame_SavesToRepository(t *testing.T) {
    // Create mock
    mockRepo := new(MockGameRepository)
    
    // Create service with mock
    gameService := NewGameService(mockRepo)
    
    // Call function
    game, err := gameService.Create("ABCD", "LOBBY")
    
    // Verify mock was called
    mockRepo.AssertCalled(t, "Create", 1)
    
    // Verify result
    assert.NoError(t, err)
    assert.Equal(t, "ABCD", game.Code)
}
```

---

## 4. E2E Testing (Playwright)

**Què és:** Provar el flux complet de l'aplicació en un navegador real (automatitzat).

**Comandes:**
```bash
cd frontend
pnpm add -D @playwright/test
npx playwright install --with-deps
pnpm playwright test
```

**Test Structure:**
```typescript
// tests/e2e/game-flow.spec.ts
import { test, expect } from '@playwright/test';

test.describe('Flujo completo del juego', () => {
  test.beforeEach(async ({ page }) => {
    // Login antes de cada test
    await page.goto('http://localhost:3000/login');
    await page.fill('input[name="email"]', 'user@example.com');
    await page.fill('input[name="password"]', 'password');
    await page.click('button[type="submit"]');
    await page.waitForURL('**/lobby');
  });

  test('lobby: unirse a sala y seleccionar rol', async ({ page }) => {
    await page.goto('http://localhost:3000/lobby');
    
    // Ingresar código
    await page.fill('input[name="gameCode"]', 'ABCD');
    await page.click('button:has-text("Unirse")');
    
    // Esperar selección de rol
    await expect(page.locator('.role-card').first()).toBeVisible();
    
    // Seleccionar rol
    await page.click('.role-card:has-text("Detective de Camp")');
    
    // Verificar que el rol está seleccionado
    await expect(page.locator('text=Tu rol:')).toContainText('Detective de Camp');
  });

  test('investigación: explorar escena y encontrar pistas', async ({ page }) => {
    await page.goto('http://localhost:3000/scene');
    
    // Hacer clic en hotspot
    await page.click('.hotspot[data-clue="pistol"]');
    
    // Verificar que la pista fue encontrada
    await expect(page.locator('.clue-inventory')).toContainText('Pistola');
  });

  test('tauler: conectar pistas y crear hipótesis', async ({ page }) => {
    await page.goto('http://localhost:3000/board');
    
    // Arrastrar pista al tauler
    const clueNode = page.locator('.node[data-type="clue"]');
    const hypothesisNode = page.locator('.node[data-type="hypothesis"]');
    
    await clueNode.dragTo(hypothesisNode);
    
    // Verificar conexión
    await expect(page.locator('.edge')).toBeVisible();
  });
});
```

---

## 5. Testing Best Practices

### ✅ DO'S

1. **Testear comportamiento, no implementación**
   - Probar QUÉ hace el componente, no CÓMO lo hace internamente

2. **Tests independientes**
   - Cada test debe poder ejecutarse solo
   - No dependen de otros tests
   - No hay orden específico de ejecución

3. **Tests rápidos**
   - Los tests deben ejecutarse en segundos, no minutos
   - Evitar `sleep` innecesario

4. **Tests mantenibles**
   - Si cambias la lógica del componente, el test no debería romperse
   - Los tests deben describir el comportamiento esperado, no la implementación

5. **Tests descriptivos**
   - Nombres de test claros: `test('usuario puede seleccionar rol', ...)`
   - Mensajes de error útiles: `expected: "Debe haber 4 roles", got: 3`
   - Organización clara: `describe('Lobby', ...)` + `describe('Roles', ...)`

6. **Mocking apropjat**
   - Moquear dependències externes (PocketBase, API)
   - Provar casos límit (errors, casos edge)
   - No hacer llamadas reales a bases de datos

7. **Coverage**
   - Buscar alta cobertura (> 80%)
   - Probar casos exitosos y de error
   - Provar caminos alternativos en el código

### ❌ DON'TS

1. **No testear código de terceras**
   - No probar que React renderiza correctamente (es problema de React)
   - No probar que Go compila (es problema del compilador)

2. **No tests con asserts frágiles**
   - Evitar `expect(true).toBe(true)` sin lógica adicional
   - Verificar algo específico del resultado

3. **No tests que dependen del tiempo**
   - Evitar `Date.now()` en asserts
   - Evitar sleeps innecesarios

4. **No tests con duplicación**
   - Evitar copiar y pegar código en tests
   - Reutilizar lógica en funciones auxiliares

5. **No tests que prueban múltiples cosas**
   - Un test debe probar UNA cosa específicamente
   - Evitar "test the entire application flow"

---

## 6. Testing Checklist

### Antes de commitear código:

**Frontend:**
- [ ] Unit tests pasan (`pnpm test -- run`)
- [ ] Component tests pasan (`pnpm test -- ui`)
- [ ] Linter no tiene errores (`pnpm lint`)
- [ ] No hay errores de TypeScript (`pnpm tsc --noEmit`)
- [ ] E2E tests pasan (opcionals)

**Backend:**
- [ ] Unit tests pasan (`go test ./...`)
- [ ] Go vet no tiene errores (`go vet ./...`)
- [ ] No hay errores de compilación (`go build ./cmd/server`)
- [ ] Coverage > 80% (`go test -cover ./...`)

**General:**
- [ ] No hay secrets commiteados
- [ ] No hay archivos .env en el repo
- [ ] Los tests son deterministas (pasan siempre)

---

## 7. Integration-First (Test-Driven Development)

### Flujo Completo de Integration-First:

```
1. DOCUMENTACIÓN
   → Crear `docs/features/XX.md`
   → Definir User Story
   → Definir Criterios de Aceptación

2. TEST FALLIDO (RED)
   → Escribir test de integración basado en la documentación
   → Ejecutar: test DEBE fallar
   → Verificar que falla: ❌

3. IMPLEMENTACIÓN
   → Escribir el código MÍNIMO para pasar el test
   → Evitar mocks si el flujo puede validarse con adaptadores reales

4. TEST PASADO (GREEN)
   → Ejecutar el test de nuevo
   → Verificar que pasa: ✅

5. REFACTORING (Opcional)
   → Mejorar el código sin cambiar el comportamiento
   → Mantener los tests pasando

6. COMMIT
   → Commit con mensaje: `feat: implementar [feature]`
```

### Ejemplo Práctico de Integration-First:

```typescript
// PASO 1: Documentación (docs/features/investigation-board.md)
// User Story: "Como jugador, quiero conectar pistas en el tauler..."

// PASO 2: Test FALLIDO (features/board/__tests__/board.test.tsx)
import { render, screen, fireEvent } from '@testing-library/react';

describe('Tauler - Conexión de pistas', () => {
  it('debería crear una conexión entre pista e hipótesis', () => {
    const { container } = render(<InvestigationBoard />);
    
    const clueNode = screen.getByTestId('clue-1');
    const hypothesisNode = screen.getByTestId('hypothesis-1');
    
    // Simular arrastrar
    fireEvent.dragStart(clueNode);
    fireEvent.dragEnter(hypothesisNode);
    fireEvent.drop(hypothesisNode);
    
    // Verificar conexión - DEBE FALLAR (feature no implementada aún)
    expect(container).toHaveClass('has-connection');
  });
});

// Resultado: ❌ TEST FALLIDO (Expected to find class...)

// PASO 3: Implementación MÍNIMA
const handleConnect = (clueId, hypothesisId, type) => {
  const newConnection = {
    id: `edge-${Date.now()}`,
    source: clueId,
    target: hypothesisId,
    type: type,
  };
  
  // Guardar conexión
  setConnections(prev => [...prev, newConnection]);
};

// PASO 4: Test PASADO ✅
// El test ahora pasa
```

---

## 8. Mocking Patterns

### Mock de PocketBase

```go
// internal/adapters/repo/mock_pocketbase_test.go
package repo_test

import (
    "testing"
    "github.com/stretchr/testify/mock"
)

func NewMockPocketBaseClient() *MockPocketBaseClient {
    mock := new(MockPocketBaseClient)
    
    // Configurar comportamiento mock
    mock.On("Collection", "games").Return(mockGameCollection)
    mock.On("Collection", "players").Return(mockPlayerCollection)
    
    return mock
}
```

### Mock de HTTP Requests

```typescript
// features/lobby/__tests__/mocks/handlers.ts
export const mockPocketBaseClient = {
  collection: vi.fn(),
  getList: vi.fn(),
  getFirstListItem: vi.fn(),
  create: vi.fn(),
  update: vi.fn(),
  delete: vi.fn(),
};

// Configurar mock en test
mockPocketBaseClient.collection('games').mockResolvedValue([
  { id: 'game-1', code: 'ABCD', status: 'LOBBY' }
]);
```

---

## 9. Debugging Tests

### Frontend Debugging

**Ver tests en UI:**
```bash
pnpm test -- ui
```

**Depurar un test específic:**
```bash
pnpm test --reporter=verbose -- testsMatchingPattern="lobby"
```

### Backend Debugging

**Executar tests amb verbositat:**
```bash
go test -v ./services
```

**Executar un test específic:**
```bash
go test -v ./services -run TestCreateGame_ValidInput
```

**Prints en tests:**
```go
func TestSomething(t *testing.T) {
    t.Log("Comenzando test...")
    // ... código de test ...
    t.Log("Test completado")
}
```

---

## 10. Performance Testing

### Load Testing (Frontend)

**Què és:** Provar que l'aplicació soporta molts usuaris simultànies.

**Herramentes:**
- k6
- JMeter
- Artillery

**Exemple simple:**
```bash
# Simular 100 usuarios durante 30 segundos
k6 run --vus 100 --duration 30s http://localhost:3000
```

### Stress Testing (Backend)

**Què és:** Provar que l'API resisteix moltes peticions.

**Herramentes:**
- Vegeta (Go)
- Apache Bench
- wrk

**Exemple simple:**
```bash
# 1000 peticiones concurrentes
wrk -t 12 -c 1000 http://localhost:8080/api/games
```

---

## 11. Recursos

- [Vitest Documentation](https://vitest.dev/)
- [Go Testing Guide](https://go.dev/doc/tutorial/add-a-test)
- [React Testing Library](https://testing-library.com/docs/react-testing-library/intro/)
- [Playwright Documentation](https://playwright.dev/)
- [Test-Driven Development](./tdd.md)

---

**Última actualización:** 20/01/2025
**Versión:** 1.0
