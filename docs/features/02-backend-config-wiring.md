# 02 - Backend Config + Wiring

## Objectiu
Centralitzar la configuracio del backend per evitar valors hardcoded i facilitar entorns futurs.

## Abast
- Crear un paquet `internal/platform/config`.
- Carregar variables d'entorn amb defaults segurs.
- Usar aquesta config a `cmd/server` per construir dependències.

## Requeriments
1. Config struct amb valors:
   - `Port`
   - `Environment`
   - `PocketBaseURL`
   - `PocketBaseTimeout`
2. Defaults:
   - `Port = 8080`
   - `Environment = development`
   - `PocketBaseTimeout = 5s`
3. Si `PB_TIMEOUT` es invalid → error.
4. No fer panic; retornar error.

## Criteris d'Aceptacio
- `Load()` retorna defaults quan no hi ha envs.
- `Load()` retorna error si `PB_TIMEOUT` no es pot parsejar.
- `cmd/server` usa `config.Load()` per inicialitzar el client de PocketBase.

## Pla de Tests (TDD)
1. **Sense envs** → defaults correctes.
2. **PB_TIMEOUT invalid** → error.
3. **PB_TIMEOUT valid** → duration correcta.
