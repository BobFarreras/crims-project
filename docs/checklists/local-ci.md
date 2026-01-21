# ✅ Checklist Local (Lint + Test + Build)

Aquest checklist s'ha d'executar **abans de qualsevol push** a `develop` o `release/*`.

**Scripts automàtics (recomanat):**
- `scripts/local-ci.sh` (macOS/Linux)
- `scripts/local-ci.ps1` (Windows PowerShell)

Executar tot el flux (macOS/Linux):
```bash
bash scripts/local-ci.sh
```

Executar sense reinstal.lar dependències (macOS/Linux):
```bash
bash scripts/local-ci.sh --skip-install
```

Executar tot el flux (Windows PowerShell):
```powershell
./scripts/local-ci.ps1
```

Executar sense reinstal.lar dependències (Windows PowerShell):
```powershell
./scripts/local-ci.ps1 -SkipInstall
```

## 1) Frontend (Next.js)

```bash
cd frontend
pnpm install
pnpm lint
pnpm test
pnpm build
```

## 2) Backend (Go)

```bash
cd backend
go mod download
go test ./...
go vet ./...
go build ./cmd/server
```

## 3) Validacions Finals

```bash
git status -sb
git diff --staged
```

**Regla:** si qualsevol pas falla, **NO fer commit ni push**.
