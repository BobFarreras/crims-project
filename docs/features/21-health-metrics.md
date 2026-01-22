# 21 - Health & Metrics

## Objectiu
Exposar endpoints de salut i metrics basiques.

## Abast
- `/api/v1/health` (ja existeix)
- `/api/v1/metrics` amb counters basics

## Requeriments
1. Response JSON amb uptime i counts.
2. Endpoint només per observabilitat.

## Criteris d'Aceptacio
- `/api/v1/metrics` respon 200 amb JSON.

## Pla de Tests (Integracio)
1. Handler retorna 200 i json.
