param(
  [switch]$SkipInstall
)

$ErrorActionPreference = "Stop"

Write-Host "==> Frontend"
Set-Location frontend
if (-not $SkipInstall) {
  Write-Host "-> install"
  pnpm install
} else {
  Write-Host "-> skip install"
}

Write-Host "-> lint"
pnpm lint

Write-Host "-> test"
pnpm test

Write-Host "-> build"
pnpm build

Write-Host "==> Backend"
Set-Location ..\backend
if (-not $SkipInstall) {
  Write-Host "-> install"
  go mod download
} else {
  Write-Host "-> skip install"
}

Write-Host "-> test"
go test ./...

Write-Host "-> vet"
go vet ./...

Write-Host "-> build"
go build ./cmd/server

Write-Host "==> Done"
