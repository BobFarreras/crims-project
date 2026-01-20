# Git Workflow & Branching Strategy 🌳

## Overview

Aquest projecte utilitza **Git Flow** adaptat per a monorepos Next.js + Go.

```
main              → Producció (sempre estable)
└── develop       → Integració (pre-producció)
    ├── feature/* → Noves funcionalitats
    ├── release/* → Preparació de versions
    ├── hotfix/*  → Correccions urgents
    ├── chore/*   → Tasques tècniques
    └── docs/*    → Documentació
```

---

## 🌳 Branches Principals

### `main` 🔵
- **Què és:** Codi en producció
- **Estabilitat:** Sempre estable i desplegable
- **Ús:** Referència final per a releases

**Regles:**
- ❌ NO es treballa directament aquí
- ✔️ Només merges aprovats via Pull Request
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
- Les feature/* es mergegen aquí
- Ha de compilar i passar tests
- Base per crear release/*

**Comandes:**
```bash
# Crear develop (només primer cop)
git checkout -b develop

# Actualitzar develop
git checkout develop
git pull origin develop
```

---

## 🌱 Branches de Treball

### `feature/*` 🟡
- **Què és:** Desenvolupament d'una funcionalitat concreta
- **Origen:** Es creen des de `develop`
- **Destí:** Es mergegen a `develop`
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

# 2. Treballar i commitear
git add .
git commit -m "feat: implement role selection in lobby"
git push -u origin feature/lobby-roles

# 3. Crear PR a GitHub (feature/lobby-roles → develop)

# 4. Després del merge (aprovat):
git checkout develop
git pull origin develop
git branch -d feature/lobby-roles
git push origin --delete feature/lobby-roles
```

**Commit naming:**
```
feat: nova funcionalitat
fix: correcció de bug
refactor: refactoring de codi
test: afegir tests
docs: documentació
chore: tasques de manteniment
style: format de codi (espais, punt i coma)
perf: millora de rendiment
ci: canvis a CI/CD
```

---

## 🚀 Branches de Preparació

### `release/*` 🟣
- **Què és:** Preparació d'una versió per producció
- **Origen:** Es creen des de `develop`
- **Destí:** Es mergejan a `main` i `develop`
- **Vida:** S'esborren després del merge

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
git commit -m "chore: prepare release v1.0.0"

# 4. Merge a main i taggear
git checkout main
git merge --no-ff release/1.0.0
git tag -a v1.0.0 -m "Release v1.0.0 - Initial launch"
git push origin main
git push origin v1.0.0

# 5. Merge a develop (backport)
git checkout develop
git merge --no-ff release/1.0.0
git push origin develop

# 6. Esborrar release
git branch -d release/1.0.0
git push origin --delete release/1.0.0
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

Exemples:
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
git commit -m "fix: crash on login when using special characters"
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

## 📦 Branches Opcionals

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

**Exemples:**
```
test/unit-tests-auth
test/e2e-game-flow
test-increase-coverage
```

---

## 🏗️ Visual Flow Diagram

```
                     main (v1.0.0)
                         │
                         │ release/1.0.0
                         │
                      develop
                        │  │
        ┌───────────────┼──┴───────────────┐
        │               │                  │
 feature/lobby  feature/board  feature/forensic
        │               │                  │
        └───────────────┼──────────────────┘
                        │
                        └─→ (merge a develop)
```

**Hotfix flow:**
```
      main (v1.0.0)  →  main (v1.0.1)
          │                 ↑
          │                 │
      hotfix/crash ────────┘
          │
          └─→ (backport a develop)
```

---

## 🧭 Quan Usar Quina Branca?

| Situació | Branca | Origen | Destí |
|----------|--------|--------|-------|
| Nova funcionalitat (lobby, board, etc.) | `feature/*` | develop | develop |
| Correcció de bug no urgent | `feature/fix-*` | develop | develop |
| Preparar versió per producció | `release/X.Y.Z` | develop | main + develop |
| Error crític en producció | `hotfix/*` | main | main + develop |
| Actualitzar dependències | `chore/*` | develop | develop |
| Afegir tests | `test/*` | develop | develop |
| Documentar | `docs/*` | develop | develop |
| Refactoring | `chore/refactor-*` | develop | develop |

---

## 🤖 Per als Agents AI (OpenCode)

Aquest document és la **font de veritat** per al workflow Git. Quan un agent hagi de:

1. **Crear nova funcionalitat:** Crear `feature/feature-nom` des de develop
2. **Arreglar bug urgent:** Crear `hotfix/descripcio` des de main
3. **Preparar release:** Crear `release/X.Y.Z` des de develop
4. **Comitejar:** Usar prefixos (feat, fix, docs, etc.)
5. **Fusió:** Sempre via Pull Request, mai direct merge

**Comanda automàtica per agents:**
```bash
# Crear feature nova
git checkout develop && git pull origin develop && git checkout -b feature/feature-name

# Crear hotfix
git checkout main && git pull origin main && git checkout -b hotfix/description

# Després de merge (aprobat)
git checkout develop && git branch -d feature/feature-name && git push origin --delete feature/feature-name
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
- ✅ Require branches to be up to date before merging
- ❌ Do not allow bypassing the above settings

### Protecció de `develop`:
- ✅ Require pull request before merging
- ✅ Require approvals: 1
- ✅ Require status checks to pass

---

## 📌 Bones Pràctiques

### ✔️ SEMPRE:
- Usar branch protection a main/develop
- Fer Pull Requests, no direct merge
- Reviews de codi abans del merge
- Commits petits i descriptius
- CI/CD ha de passar abans del merge
- Esborrar branches després del merge
- Tags a cada release a main

### ❌ MAI:
- Commitejar directament a main
- Commit "fix typo", "update", etc. (més descriptiu)
- Pujar secrets o .env files
- Deixar branches antigues al remote
- Push force a branches compartides
- Ignorar warnings de linter

---

## 📅 Exemple de Projecte Real

**Fase 1: Setup** (main: v0.1.0)
```bash
feature/setup-monorepo
feature/configure-ci-cd
feature/add-base-documentation
→ Merge to develop
→ Release v0.1.0 → main
```

**Fase 2: MVP** (main: v1.0.0)
```bash
feature/lobby-roles
feature/investigation-board
feature/scene-exploration
feature/auth-system
feature/pwa-manifest
→ Merge to develop
→ Release v1.0.0 → main
```

**Fase 3: Post-MVP** (main: v1.1.0)
```bash
feature/forensic-tools
feature/interrogation-system
feature/timeline-editor
feature/sentry-integration
→ Merge to develop
→ Release v1.1.0 → main
```

**Emergència:**
```bash
hotfix/login-crash (des de main)
→ Merge a main → v1.0.1
→ Backport a develop
```

---

## 🛠️ Aliases Útils (opcional)

Afegir a `~/.gitconfig`:
```bash
[alias]
    co = checkout
    br = branch
    st = status
    ci = commit
    fe = "!f() { git checkout develop && git pull origin develop && git checkout -b feature/$@; }; f"
    hf = "!f() { git checkout main && git pull origin main && git checkout -b hotfix/$@; }; f"
    merge-feature = "!f() { git checkout develop && git merge --no-ff feature/$@; }; f"
    done = "!f() { git checkout develop && git branch -d $@ && git push origin --delete $@; }; f"
```

Ús:
```bash
git fe lobby-roles      # Crear feature/lobby-roles
git merge-feature lobby-roles
git done feature/lobby-roles
```

---

## 📞 Suport

Per dubtes sobre el workflow:
- GitHub Issues: https://github.com/BobFarreras/crims-project/issues
- Email: dev@digitaistudios.com

---

**Última actualització:** 20/01/2025
