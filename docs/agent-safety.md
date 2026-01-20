# 🛡️ Medidas de Seguridad para Agentes AI (OpenCode)

## 🚨 Propósito

Este documento definece las medidas de seguridad que DEBEN seguir los agentes AI (como OpenCode) al realizar operaciones en el repositorio para evitar romper la aplicación en producción.

---

## ⚠️ Principios Fundamentales

### 1. Never Push Directly to Main
**PROHIBIDO:** Nunca hacer `git push origin main` directamente sin aprobación.

**CORRECTO:**
- Trabajar en `feature/*`, `hotfix/*`, `release/*`, `chore/*`, `docs/*`
- Crear Pull Request para aprobación
- Esperar aprobación humana antes del merge

### 2. Branch Protection Rules
Antes de CUALQUIER operación que afecte código crítico:

✅ **Verificar que:**
- CI/CD está pasando (tests, lint, build)
- No hay errores de compilación
- No hay tests fallando
- El código formatea correctamente
- No hay warnings críticos

❌ **DETENER si:**
- Tests failing
- Build errors
- Lint errors
- Missing dependencies
- Security vulnerabilities

### 3. Pre-commit Verification
Antes de hacer commit, ejecutar:

```bash
# Frontend
cd frontend && pnpm lint && pnpm test -- run

# Backend
cd backend && go test ./... && go vet ./...

# Si algún comando falla → NO COMMIT
```

### 4. Dry-run Mode
Antes de hacer push:

```bash
# 1. Verificar cambios
git status
git diff --staged

# 2. Verificar que no hay secretos
git diff --staged | grep -i "password\|secret\|api_key\|token"

# 3. Verificar archivos peligrosos
git diff --staged --name-only | grep -E "\.env|\.pem|\.key"

# 4. Si hay secretos o archivos peligrosos → ABORTAR
```

---

## 🔒 Niveles de Seguridad

### Nivel 1: Documentación (Bajo Riesgo) 🟢
Operaciones permitidas:
- Crear/actualizar archivos en `docs/`
- Modificar README
- Actualizar skills en `.ai/`

**Procedimiento:**
1. Crear `branch: docs/descripcion`
2. Commitear cambios
3. Crear PR a `develop`
4. **Puede hacer push directo sin aprobación previa** (es solo documentación)

### Nivel 2: Tests y Configuración (Riesgo Medio) 🟡
Operaciones permitidas:
- Añadir/actualizar tests
- Modificar `Makefile`
- Actualizar `package.json` (versiones menores)
- Modificar configs (tsconfig, eslint, etc.)

**Procedimiento:**
1. Crear `branch: test/nombre` o `chore/nombre`
2. Ejecutar tests locally: `make test-unit`
3. **Solo puede hacer push si todos los tests pasan**
4. Crear PR a `develop`
5. Esperar aprobación para merge

### Nivel 3: Funcionalidades (Riesgo Alto) 🟠
Operaciones permitidas:
- Implementar nuevas features
- Modificar lógica de negocio
- Cambios en `backend/` o `frontend/features/`

**Procedimiento:**
1. **VERIFICAR TDD PRIMERO:**
   - ¿Existe `docs/features/XX.md`? → Si no, PEDIR PERMISO
   - ¿Existe el test? → Si no, PEDIR PERMISO
2. Crear `branch: feature/nombre`
3. Implementar siguiendo TDD
4. Ejecutar `make test-unit` → **Si falla, NO COMMIT**
5. Ejecutar `pnpm lint` → **Si falla, NO COMMIT**
6. Crear PR a `develop`
7. **NO puede hacer push directamente** → PEDIR REVISIÓN

### Nivel 4: Hotfix en Producción (Riesgo Crítico) 🔴
Operaciones permitidas:
- Correcciones urgentes en `main`
- Security patches
- Critical bugs

**Procedimiento:**
1. **PEDIR PERMISO EXPLÍCITO AL USUARIO**
2. Crear `branch: hotfix/descripcion` DESDE `main`
3. Implementar fix mínimo
4. Ejecutar **TODOS** los tests → **Si falla, ABORTAR**
5. Crear PR a `main` Y a `develop`
6. **NO puede hacer push** → PEDIR REVISIÓN Y APROBACIÓN HUMANA

---

## 🚫 Prohibiciones Absolutas

### ❌ NUNCA hagas:
1. `git push origin main` directamente (solo release/hotfix aprobados)
2. Commit de `.env` files
3. Commit de passwords, API keys, tokens
4. Commit de archivos `.pem`, `.key`, certificates
5. `git push --force` en branches compartidos
6. Commit con mensaje vacío o sin sentido
7. Borrar archivos sin confirmación
8. Modificar `package.json` o `go.mod` sin verificar compatibilidad
9. Commit de `node_modules` o carpetas de build
10. Modificar archivos de configuración crítica sin revisión

---

## 🤖 GitHub Actions - Cuando se Ejecutan los Tests

### Regla de Oro:
**Los GitHub Actions SOLO se ejecutan en `main` y `release/*` branches.**

Esto significa:
- ✅ Si haces push a `main` → CI/CD se ejecuta (tests, lint, build)
- ✅ Si haces push a `release/*` → CI/CD se ejecuta (tests, lint, build)
- ❌ Si haces push a `develop` → CI/CD NO se ejecuta (ahorra tiempo)
- ❌ Si haces push a `feature/*` → CI/CD NO se ejecuta (ahorra tiempo)

### Ventajas para Agentes AI:
1. **Seguridad:** Los cambios solo se verifican cuando van a producción
2. **Velocidad:** Puedes commitear en `develop` sin esperar el CI/CD
3. **Previene Roturas:** Si el CI/CD falla, el PR no puede mergearse

### ¿Debes Ejecutar Tests Localmente?

**SÍ, SIEMPRE:**
- Antes de CUALQUIER commit
- Sigue el checklist de abajo

**POR QUÉ:**
- Aunque el CI/CD solo corra en `main` y `release`, los tests deben pasar localmente
- Esto asegura que `develop` esté siempre en estado funcional
- Cuando crees un `release/*` y el CI/CD falle, sabrás exactamente qué arreglar

---

## ✅ Checklist Antes de Hacer Push

Siempre ejecutar este checklist antes de cualquier push:

### Paso 1: Verificación de Seguridad
```bash
# ¿Hay secretos en el commit?
git diff --cached | grep -i "password\|secret\|api_key\|token\|jwt"
# Si devuelve algo → ABORTAR

# ¿Hay archivos de env?
git diff --cached --name-only | grep "\.env"
# Si devuelve algo → ABORTAR
```

### Paso 2: Verificación de Tests
```bash
# Frontend
cd frontend && pnpm test -- run
# Si tests fallan → ABORTAR

# Backend
cd backend && go test ./...
# Si tests fallan → ABORTAR
```

### Paso 3: Verificación de Lint
```bash
# Frontend
cd frontend && pnpm lint
# Si hay errores → ABORTAR

# Backend (si hay linter)
cd backend && go vet ./...
# Si hay errores → ABORTAR
```

### Paso 4: Verificación de Build
```bash
# Frontend
cd frontend && pnpm build
# Si falla → ABORTAR

# Backend
cd backend && go build ./cmd/server
# Si falla → ABORTAR
```

### Paso 5: Verificación de Archivos
```bash
# ¿Hay archivos incorrectos commiteados?
git diff --cached --name-only | grep -E "node_modules|\.next|build|\.git|\.env"
# Si devuelve algo → RESTORE y RESTART
```

### Paso 6: Verificación de Mensaje de Commit
```bash
# ¿El mensaje sigue las normas?
# Formato: tipo: descripción (CASTELLANO)
# Ejemplos correctos:
✅ feat: implementar selección de roles en el lobby
✅ fix: corregir error en autenticación JWT
✅ docs: añadir guía de deployment

# Ejemplos incorrectos:
❌ fix typo
❌ update
❌ wip
❌ test changes
```

---

## 🤖 Protocolo para Agentes AI

### Cuando el Usuario Solicite Algo Destructivo

**Situación:** El usuario pide eliminar archivos, borrar carpetas, etc.

**Protocolo:**
1. ⚠️ **ALERTA AL USUARIO:** "Esto eliminará archivos. ¿Estás seguro?"
2. Mostrar qué se eliminará: `git status`
3. Esperar confirmación explícita del usuario
4. Solo después, proceder con la operación

### Cuando el Usuario Solicite Hacer Commit/Push

**Situación:** El usuario pide hacer commit y push de cambios.

**Protocolo:**
1. ⚠️ **VERIFICAR CHECKLIST:** Ejecutar los 6 pasos
2. Si algún paso falla → **DETENERSE** y reportar el problema
3. Si todo pasa:
   - Si es documentación (docs/*) → Push directo OK
   - Si es código → "He hecho el commit. ¿Quieres que cree el PR o esperas revisión?"

### Cuando el Usuario Solicite Implementar Feature

**Situación:** El usuario pide implementar una nueva funcionalidad.

**Protocolo:**
1. ⚠️ **VERIFICAR TDD:**
   - ¿Existe `docs/features/XX.md`?
   - Si NO → "Antes de implementar, necesito crear la documentación. ¿Te parece bien?"
   - Si SÍ → Continuar con el siguiente paso
2. ⚠️ **VERIFICAR TEST:**
   - ¿Existe el test?
   - Si NO → "Antes de implementar, necesito crear el test. ¿Te parece bien?"
   - Si SÍ → Implementar el código

### Cuando Hay Errores en Tests

**Situación:** Tests están fallando después de un cambio.

**Protocolo:**
1. ❌ **NO HACER COMMIT**
2. Reportar qué tests fallan: `pnpm test -- run`
3. Solicitar permiso para:
   - a) Arreglar los tests
   - b) Deshacer los cambios
4. Esperar confirmación del usuario

---

## 📋 Flujo de Decisión para Agentes

```
┌─────────────────────────────────────────┐
│  Usuario solicita operación            │
└──────────────┬────────────────────────┘
               │
               ▼
       ┌───────────────┐
       │ ¿Es documentación? │
       └───────┬───────┘
               │
     ┌─────────┴─────────┐
     │ Sí              │ No
     ▼                  ▼
┌─────────────┐   ┌──────────────────────┐
│ Crear PR   │   │ Ejecutar Checklist    │
│ a develop  │   │ (6 pasos)             │
│ Push OK    │   └──────────┬───────────┘
└─────────────┘              │
                             ▼
                    ┌──────────────────┐
                    │ ¿Todo pasó?      │
                    └────────┬─────────┘
                             │
              ┌──────────────┴──────────────┐
              │ Sí                      │ No
              ▼                          ▼
     ┌─────────────────┐      ┌─────────────────┐
     │ Crear PR        │      │ ❌ ABORTAR       │
     │ Solicitar review │      │ Reportar error  │
     └─────────────────┘      └─────────────────┘
```

---

## 🚨 Procedimiento de Emergencia

### Si Rompes Algo en Producción

**SITUACIÓN CRÍTICA:** Has hecho un cambio que rompe producción.

**PROCEDIMIENTO:**
1. ❌ **NO hacer más cambios**
2. ⚠️ **ALERTAR INMEDIATAMENTE** al usuario:
   - "He detectado que mis cambios han causado un problema. Detente todo."
3. Verificar qué está roto:
   ```bash
   git log --oneline -5
   git diff HEAD~1
   ```
4. Deshacer cambios:
   ```bash
   git revert HEAD
   git push origin main
   ```
5. Reportar qué falló y por qué

### Si Commiteas Secretos

**SITUACIÓN CRÍTICA:** Has commiteado passwords, API keys, etc.

**PROCEDIMIENTO:**
1. ❌ **NO hacer push**
2. ⚠️ **ALERTAR INMEDIATAMENTE** al usuario
3. Eliminar el commit:
   ```bash
   git reset --soft HEAD~1
   git checkout -- archivo_con_secreto
   git commit -m "fix: remover secreto accidental"
   ```
4. Rotar las credenciales comprometidas (informar al usuario)

---

## 🔗 Recursos Relacionados

- [Git Workflow](./git-workflow.md) - Estrategia de branches
- [Security Skills](../.ai/skills/skill-security.md) - OWASP y seguridad
- [Deployment Guide](./deployment.md) - Procedimiento de deployment

---

**Última actualización:** 20/01/2025
**Versión:** 1.0
