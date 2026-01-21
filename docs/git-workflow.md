# Git Workflow & Branching Strategy 🌳

## Overview

Aquest projecte utilitza **Git Flow** adaptat per a monorepos Next.js + Go.

```
main              → Producció (només versions estables)
└── develop       → Integració (conté totes les features acabades)
    ├── feature/* → Noves funcionalitats (desenvolupament)
    ├── release/* → Preparació de versions (pre-producció)
    ├── hotfix/*  → Correccions urgents (de producció)
    ├── chore/*   → Tasques tècniques
    └── docs/*    → Documentació
```

---

## 🚨 REGLA D'OR: MAIN ÉS NOMÉS PER RELEASES

⚠️ **IMPORTANT:** `main` conté **ÚNICAMENT** versions estables de producció.

**NUNCA** fas merge directe de `feature/*` a `main`.

**Flux CORRECTE:**
```
feature/* → develop → main (només quan hi ha release)
```

---

## 🌳 Branches Principals

### `main` 🔵
- **Què és:** Codi en producció (només versions estables)
- **Estabilitat:** Sempre estable i desplegable
- **Ús:** Referència final per a releases

**Regles:**
- ❌ NO es treballa directament aquí
- ❌ NO es fa merge directe de `feature/*` a main
- ✔️ Només merges aprovats via Pull Request des de `develop` o `release/*`
- ✔️ Cada merge ha de tenir tag de versió (ex: `v1.0.0`)
- ✔️ Protegit amb branch protection

**Comandes:**
```bash
git checkout main
git pull origin main
git tag v1.0.0
git push origin v1.0.0
```

---

### `develop` 🟢
- **Què és:** Estat "pre-producció" / Integració
- **Estabilitat:** Pot tenir bugs menors
- **Ús:** On es combinen totes les feature/* acabades

**Regles:**
- Les `feature/*` es mergegen aquí
- Ha de compilar i passar tests
- Base per crear `release/*`
- Les `hotfix/*` es mergegen aquí (backport)

**Comandes:**
```bash
# Actualitzar develop
git checkout develop
git pull origin develop
```

---

## 🌱 Branches de Treball

### `feature/*` 🟡
- **Què és:** Desenvolupament d'una funcionalitat concreta
- **Origen:** Es creen des de `develop`
- **Destí:** Es mergegen a `develop` (NO a main)
- **Vida:** S'esborren després del merge

**Nomenclatura:**
```
feature/feature-nom
feature/feature-category

Exemples:
feature/lobby-roles
feature/investigation-board
feature/forensic-tools
feature/auth-jwt
feature/pwa-manifest
```

**Flux complet:**
```bash
# 1. Crear branca des de develop
git checkout develop
git pull origin develop
git checkout -b feature/lobby-roles

# 2. Treballar i commitejar
git add .
git commit -m "feat: implementar selección de roles en el lobby"
git push -u origin feature/lobby-roles

# 3. Crear PR a GitHub (feature/lobby-roles → develop)

# 4. Després del merge (aprovat):
git checkout develop
git pull origin develop
git branch -d feature/lobby-roles
git push origin --delete feature/lobby-roles
```

**Commit naming (IDIOMA: CASTELLANO):**
```
feat: nueva funcionalidad
fix: corrección de bug
refactor: refactorización de código
test: añadir tests
docs: documentación
chore: tareas de mantenimiento
style: formato de código (espacios, punto y coma)
perf: mejora de rendimiento
ci: cambios a CI/CD
```

---

## 🚀 Branches de Preparació

### `release/*` 🟣
- **Què és:** Preparació d'una versió per producció
- **Origen:** Es creen des de `develop`
- **Destí:** Es mergejen a `main` i `develop`
- **Vida:** **No s'esborren** (es manté l'històric de versions)
- **Nomenclatura recomanada:** `release/vX.Y.Z`

**Nomenclatura:**
```
release/VERSION
release/ANIA-MES

Exemples:
release/1.0.0
release/1.1.0
release/2025-01
```

**Què s'hi fa:**
- ✅ Bugfixos finals
- ✅ Canvis petits de configuració
- ✅ Actualització de versionat (package.json)
- ✅ Documentació de release
- ❌ NO features noves

**Documentació de la versió (recomanat):**
- Afegir nota de release a `docs/releases/vX.Y.Z.md`
- Actualitzar `CHANGELOG.md` si existeix
- Actualitzar versio a `frontend/package.json`

**Nota per agents:** L'agent només arriba fins al merge de `develop` a `release/*`. El merge i push a `main` el fa exclusivament l'usuari.

**Flux:**
```bash
# 1. Crear release des de develop
git checkout develop
git pull origin develop
git checkout -b release/1.0.0

# 2. Preparar versió
# - Actualitzar version numbers
# - Crear CHANGELOG.md
# - Fer últims ajustos

# 3. Commit canvis
git add .
git commit -m "chore: preparar release v1.0.0"

# 4. Merge a main i taggear (NOMÉS USUARI)
git checkout main
git merge --no-ff release/1.0.0
git tag -a v1.0.0 -m "Release v1.0.0 - Initial launch"
git push origin main
git push origin v1.0.0

# 5. Merge a develop (backport)
git checkout develop
git merge --no-ff release/1.0.0
git push origin develop

# 6. No esborres release (es guarda la versio)
```

---

## 🔥 Branches d'Emergència

### `hotfix/*` 🔴
- **Què és:** Correcció urgent d'error en producció
- **Origen:** Es creen des de `main`
- **Destí:** Es mergejen a `main` i `develop`
- **Vida:** S'esborren després del merge

**Quan usar-ho:**
- ⚠️ Crash de l'aplicació
- ⚠️ Security vulnerability
- ⚠️ Pèrdua de dades
- ⚠️ Error crític que bloqueja el joc

**Nomenclatura:**
```
hotfix/descripcio-curta

Exemplos:
hotfix/crash-on-login
hotfix/security-jwt-expiry
hotfix/data-loss-investigation
hotfix/database-connection-leak
```

**Flux:**
```bash
# 1. Crear hotfix des de main
git checkout main
git pull origin main
git checkout -b hotfix/crash-on-login

# 2. Arreglar el problema
# (fer canvis necessaris)

# 3. Commit i testar
git add .
git commit -m "fix: corregir fallo en login al usar caracteres especiales"
git push -u origin hotfix/crash-on-login

# 4. Merge a main
git checkout main
git merge --no-ff hotfix/crash-on-login
git tag -a v1.0.1 -m "Hotfix v1.0.1 - Fix login crash"
git push origin main
git push origin v1.0.1

# 5. Merge a develop
git checkout develop
git merge --no-ff hotfix/crash-on-login
git push origin develop

# 6. Esborrar hotfix
git branch -d hotfix/crash-on-login
git push origin --delete hotfix/crash-on-login
```

---

## 📦 Branches Opcionals (segons projecte)

### `chore/*` 🧹
- **Què és:** Tasques tècniques sense funcionalitat d'usuari
- **Origen:** `develop`
- **Destí:** `develop`

**Exemples:**
```
chore/update-dependencies
chore/setup-sentry
chore/refactor-auth-middleware
chore/ci-cd-pipeline
```

### `docs/*` 📚
- **Què és:** Només canvis de documentació
- **Origen:** `develop`
- **Destí:** `develop`

**Exemples:**
```
docs/update-readme
docs/add-deployment-guide
docs/api-documentation
```

### `test/*` 🧪
- **Què és:** Només afegir o millorar tests
- **Origen:** `develop`
- **Destí:** `develop`

**Exemplos:**
```
test/unit-tests-auth
test/e2e-game-flow
test-increase-coverage
```

---

## 🏗️ Estructura Típica Visual

```
main (v1.0.0 - PRODUCCIÓN)
└── develop (integració)
    ├── feature/lobby-roles
    ├── feature/investigation-board
    ├── feature/forensic-tools
    └── feature/auth-jwt
```

---

## 🧭 Quan Usar Quina Branca

| Situació | Branca | Origen | Destí |
|----------|--------|--------|--------|
| Nova funcionalitat (lobby, board, etc.) | `feature/*` | develop | develop |
| Correcció de bug no urgent | `feature/fix-*` | develop | develop |
| Correcció urgent en producció | `hotfix/*` | main | main + develop |
| Preparar versió per producció | `release/*` | develop | main + develop |
| Actualitzar dependències | `chore/*` | develop | develop |
| Documentar | `docs/*` | develop | develop |
| Afegir tests | `test/*` | develop | develop |

---

## 🤖 Para los Agentes AI (OpenCode)

Este documento es la **fuente de verdad** para el workflow Git. Cuando un agente deba:

1. **Crear nueva funcionalidad o lògica:** Crear sempre `feature/nombre-feature` desde develop
2. **Arreglar bug urgente:** Crear `hotfix/descripcion` desde main
3. **Preparar release:** Crear `release/X.Y.Z` desde develop
4. **Comitejar:** Usar prefijos (feat, fix, docs, etc.) en **CASTELLANO**
5. **Fusión:** Siempre vía Pull Request, nunca direct merge
6. **⚠️ L'agent MAI fa merge ni push a `main` (això ho fa l'usuari)**
7. **⚠️ L'agent finalitza la feina quan `develop` es mergeja a `release/*`**
8. **⚠️ NO eliminar branches `release/*` (es mantenen les versions)**

**Comanda automática para agentes:**
```bash
# Crear feature nueva
git checkout develop && git pull origin develop && git checkout -b feature/nombre-feature

# Crear hotfix
git checkout main && git pull origin main && git checkout -b hotfix/descripcion

# Després de merge (aprobat)
git checkout develop && git branch -d feature/nombre-feature && git push origin --delete feature/nombre-feature
```

---

## 🔄 Branch Protection Rules (GitHub)

### Protecció de `main`:
- ✅ Require pull request before merging
- ✅ Require approvals: 1
- ✅ Dismiss stale reviews
- ✅ Require status checks to pass before merging
  - CI/CD pipeline
  - Tests (frontend + backend)
  - Linter
  - Security checks
  - E2E tests
- ✅ Require branches to be up to date before merging
- ✅ Restrict who can push to main
- ❌ Do not allow bypassing the above settings

### Protecció de `develop`:
- ✅ Require pull request before merging
- ✅ Require approvals: 1
- ✅ Require status checks to pass
- ✅ Require branches to be up to date before merging

---

## 🤖 GitHub Actions - Seguridad Adicional

### Solo se Ejecuta en Branches Seguras

Los GitHub Actions **SOLO** se ejecutan en:
- ✅ `main` - Producción
- ✅ `release/*` - Preparación de release

**NO** se ejecutan en:
- ❌ `develop` - Integración (para no romper nada mientras se desarrolla)
- ❌ `feature/*` - Desarrollo de funcionalidades
- ❌ `hotfix/*` - Correcciones urgentes (se ejecutan después de merge a main)

### Ventajas de esta Estrategia

1. **Seguridad:** Los cambios solo se verifican cuando van a producción
2. **Velocidad:** Puedes commitear en `develop` sin esperar el CI/CD
3. **Previene Roturas:** Si el CI/CD falla, el PR no puede mergearse

### ¿Debes Ejecutar Tests Localmente?

**SÍ, SIEMPRE:**
- Antes de CUALQUIER commit
- Sigue el checklist de abajo
- Esto asegura que `develop` esté siempre en estado funcional
- Cuando crees un `release/*` y el CI/CD falle, sabrás exactamente qué arreglar

### Flujo del CI/CD
```
develop → release/1.0.0     → CI/CD ✅
release/1.0.0 → main       → CI/CD ✅
feature/* → develop         → CI/CD ❌ (local tests only)
hotfix/* → main             → CI/CD ✅
```

---

## ✅ Checklist Antes de Hacer Push

Siempre ejecuta este checklist antes de cualquier push:

**Checklist local complet (lint + test + build):** `docs/checklists/local-ci.md`

### Paso 1: Verificación de Seguridad
```bash
# ¿Hay secretos en el commit?
git diff --staged | grep -i "password\|secret\|api_key\|token"
# Si devuelve algo → ABORTAR

# ¿Hay archivos .env?
git diff --staged --name-only | grep "\.env"
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
git diff --staged --name-only | grep -E "node_modules|\.next|build|\.git|\.env"
# Si devuelve algo → RESTORE y RESTART
```

### Paso 6: Verificación de Mensaje de Commit
```bash
# ¿El mensaje sigue las normas?
# Formato: tipo: descripción (CASTELLANO)
# Ejemplos correctos:
✅ feat: implementar selección de roles en el lobby
✅ fix: corregir fallo en login con caracteres especiales
✅ docs: añadir guía de deployment

# Ejemplos incorrectos:
❌ fix typo
❌ update
❌ wip
❌ test changes
```

---

## 📊 Flujo de Decisión para Agentes

```
┌─────────────────────────────────────────┐
│  Usuario solicita operación          │
└──────────────┬──────────────────────┘
               │
               ▼
       ┌───────────────┐
       │ ¿Es documentación? │
       └───────┬───────┘
                 │
         ┌───────────┴───────────┐
         │ Sí              │ No
         ▼                  ▼
┌──────────────────┐   ┌──────────────────────┐
│ Crear PR       │   │ Ejecutar Checklist   │
│ a develop       │   └──────────┬───────────┘
│ (push OK)      │                │
└──────────────────┘                ▼
                          ┌──────────────────────┐
                          │ ¿Todo pasó?      │
                          └──────────┬───────────┘
                                     │
                        ┌──────────────┴──────────────┐
                        │ Sí                  │ No
                        ▼                      ▼
               ┌──────────────────┐   ┌─────────────────┐
               │ Push directo OK  │   │ ❌ ABORTAR      │
               └──────────────────┘   │ Reportar error  │
                                       └─────────────────┘
```

---

## 📋 Resumen de Cambios

### Cambios en esta versión:

1. **Main es ÚNICAMENTE para producción**
   - ❌ NUNCA hacer merge directo de `feature/*` a `main`
   - ✅ `feature/*` → `develop` (cuando la feature está terminada)
   - ✅ `develop` → `main` (SOLAMENTE cuando hay un release)
   - ✅ `release/*` → `main` y `develop` (para preparar versión)
   - ✅ `hotfix/*` → `main` y `develop` (correcciones urgentes)

2. **Flujo de trabajo**
   - `develop` contiene todas las features en desarrollo
   - `main` contiene solo versiones estables
   - Los PRs se hacen a `develop` (no a `main`)
   - `main` se actualiza SOLO desde `develop` o `release/*`

3. **GitHub Actions**
   - Solo se ejecutan en `main` y `release/*`
   - NO se ejecutan en `develop` ni `feature/*`
   - Esto permite desarrollarse más rápido en `develop`
   - Los cambios se validan cuando van a producción

4. **Branch Protection**
   - `main` tiene protección estricta
   - Solo merges aprobados desde `develop` o `release/*`
   - `develop` tiene protección para merges desde `feature/*`

---

## 🔗 Recursos Relacionados

- [Agent Safety](./agent-safety.md) - Medidas de seguridad para agentes AI
- [Deployment Guide](./deployment.md) - Procedimiento de deployment
- [Security Skills](../.ai/skills/skill-security.md) - OWASP y seguridad
- [Sentry Setup](./sentry-setup.md) - Configuración de error tracking

---

## 📞 Suporte

Para dudas sobre el workflow:
- GitHub Issues: https://github.com/BobFarreras/crims-project/issues
- Email: dev@digitaistudios.com

---

**Última actualización:** 20/01/2025
**Versión:** 2.0
