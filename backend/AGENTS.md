# SUB-AGENT: BACKEND ENGINEER ⚙️

## Context Específic
El cervell del joc. Gestiona la veritat de la partida i la comunicació amb la IA.
Ha de ser ràpid, concurrent i robust.

## Tech Stack
* **Llenguatge:** Go (Golang) 1.22+.
* **Arquitectura:** Clean Architecture (Handlers -> Services -> Repositories).
* **Comunicació:** HTTP REST + (Futur) WebSockets.

## Estructura de Carpetes (Go Standard Layout)
* `/cmd/server`: El `main.go`. Punt d'entrada.
* `/internal`: Codi privat de l'aplicació.
    * `/game`: Lògica de la partida.
    * `/ai`: Client per parlar amb LLMs.
* `/pkg`: Llibreries compartides.

## Normes de Desenvolupament
1.  **Error Handling:**
    * A Go els errors es retornen, no es llencen excepcions.
    * Sempre comprova `if err != nil`.

2.  **Connexió DB:**
    * Connectem a PocketBase via HTTP API o utilitzant-lo com a llibreria si calgués (en aquest cas, via HTTP API).

3.  **Logs:**
    * Tot ha de quedar registrat (Stdout) per veure-ho a Docker logs.

## Skills Rellevants
* Per estructurar nous endpoints -> `.ai/skills/skill-golang.md`