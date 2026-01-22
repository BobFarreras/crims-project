# 22 - AI Integration (Gemini)

## Objectiu
Integrar Gemini per generar text narratiu (flavor text) sense afectar la logica.

## Abast
- Client HTTP Gemini a `internal/ai/gemini`.
- Servei `AIGenerator` amb una funcio `GenerateNarrative`.
- Config via `GEMINI_API_KEY`.

## Requeriments
1. API key via env `GEMINI_API_KEY`.
2. Timeout configurable.
3. Errors retornats, sense panic.

## Criteris d'Aceptacio
- Sense API key → error clar.
- Resposta ok → text no buit.

## Pla de Tests (TDD)
1. Client sense API key → error.
2. Client respon OK → retorna text.
