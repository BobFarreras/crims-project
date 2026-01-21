#!/usr/bin/env bash
set -euo pipefail

SKIP_INSTALL=false
for arg in "$@"; do
  case "$arg" in
    --skip-install)
      SKIP_INSTALL=true
      ;;
  esac
done

echo "==> Frontend"
cd frontend
if [ "$SKIP_INSTALL" = false ]; then
  echo "-> install"
  pnpm install
else
  echo "-> skip install"
fi

echo "-> lint"
pnpm lint

echo "-> test"
pnpm test

echo "-> build"
pnpm build

echo "==> Backend"
cd ../backend
if [ "$SKIP_INSTALL" = false ]; then
  echo "-> install"
  go mod download
else
  echo "-> skip install"
fi

echo "-> test"
go test ./...

echo "-> vet"
go vet ./...

echo "-> build"
go build ./cmd/server

echo "==> Done"
