# 🧱 Data Model (PocketBase Draft)

Draft inicial de col·leccions i camps. Objectiu: base mínima, escalable i preparada per canviar a Supabase si cal.

## Principis
- PocketBase és la font de veritat (Single Source of Truth).
- Separar domini (backend) de persistència (adapters).
- Camps estables + versionables per a migracions futures.
- Disseny optimitzat per lectures frequents i latencia baixa.

## Col·leccions

### 1) Game
**Per què:** Conté l'estat global d'una partida.

**Camps suggerits:**
- `code` (string, unique): codi d'accés (ex: ABCD).
- `state` (string): `BRIEFING | INVESTIGATION | ACCUSATION_PHASE | RESOLUTION`.
- `seed` (string): seed base per generar narrativa.
- `createdAt` (timestamp)
- `updatedAt` (timestamp)

**Relacions:**
- 1 Game → N Players
- 1 Game → N Clues
- 1 Game → N Events

### 2) Player
**Per què:** Estat i rol de cada jugador dins d'una partida.

**Camps suggerits:**
- `gameId` (relation → Game)
- `userId` (string): ID d'auth (PocketBase auth o extern).
- `role` (string): `DETECTIVE | FORENSIC | ANALYST | INTERROGATOR`.
- `status` (string): `ONLINE | OFFLINE | DISCONNECTED`.
- `isHost` (bool)
- `createdAt` (timestamp)
- `updatedAt` (timestamp)

### 3) Clue
**Per què:** Pistes descobertes i el seu estat d'anàlisi.

**Camps suggerits:**
- `gameId` (relation → Game)
- `type` (string): `OBJECT | DOCUMENT | TESTIMONY | TRACE`.
- `state` (string): `DISCOVERED | ANALYZED`.
- `reliability` (int 0-100)
- `facts` (json): llista d'EvidenceFacts.
- `createdAt` (timestamp)
- `updatedAt` (timestamp)

### 4) Person
**Per què:** Personatges i testimonis amb estats de credibilitat.

**Camps suggerits:**
- `gameId` (relation → Game)
- `name` (string)
- `officialStory` (text)
- `truthStory` (text)
- `stress` (int 0-100)
- `credibility` (int 0-100)
- `createdAt` (timestamp)
- `updatedAt` (timestamp)

### 5) Event
**Per què:** Fets temporals per a la timeline i coherència.

**Camps suggerits:**
- `gameId` (relation → Game)
- `timestamp` (string/ISO)
- `locationId` (string)
- `participants` (relation[] → Person)
- `createdAt` (timestamp)
- `updatedAt` (timestamp)

## Notes d'escalabilitat
- Relacions sempre via `gameId` per filtrar ràpid.
- Camps d'estat amb enums controlats per backend.
- `facts` com JSON per evolució sense migracions costoses.
- Indexar `gameId`, `code`, i `userId` per consultes rapides.
- Evitar joins profunds: carregar entitats per lots (batch).

## Preparat per Supabase
- Models equivalents a taules SQL sense canvis semàntics.
- `gameId` com FK.
- `facts` com JSONB.

## Indexos recomanats (PocketBase)
- `Game.code` (unique)
- `Player.gameId`
- `Player.userId`
- `Clue.gameId`
- `Person.gameId`
- `Event.gameId`

## Access patterns (rendiment)
- Lobby: carregar `Game` per `code` i llistar `Player` per `gameId`.
- Board: carregar `Clue` per `gameId` i paginar si cal.
- Timeline: carregar `Event` per `gameId` ordenat per `timestamp`.
