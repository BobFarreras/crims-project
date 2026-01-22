# PocketBase Schema (Project)

Aquest document defineix les col·leccions i camps que hem creat al backend.

## Col·leccions

### games
- `code` (text, unique)
- `state` (text)
- `seed` (text)

### players
- `gameId` (relation -> games, single)
- `userId` (text)
- `role` (text)
- `status` (text)
- `isHost` (bool)

### events
- `gameId` (relation -> games, single)
- `timestamp` (text)
- `locationId` (text)
- `participants` (relation -> persons, multiple)

### clues
- `gameId` (relation -> games, single)
- `type` (text)
- `state` (text)
- `reliability` (number)
- `facts` (json)

### persons
- `gameId` (relation -> games, single)
- `name` (text)
- `officialStory` (text)
- `truthStory` (text)
- `stress` (number)
- `credibility` (number)

### hypotheses
- `gameId` (relation -> games, single)
- `title` (text)
- `strengthScore` (number)
- `status` (text)
- `nodeIds` (json)

### accusations
- `gameId` (relation -> games, single)
- `playerId` (relation -> players, single)
- `suspectId` (relation -> persons, single)
- `motiveId` (text)
- `evidenceId` (relation -> clues, single)
- `verdict` (text)

### forensics
- `gameId` (relation -> games, single)
- `clueId` (relation -> clues, single)
- `result` (text)
- `confidence` (number)
- `status` (text)

### timeline
- `gameId` (relation -> games, single)
- `timestamp` (text)
- `title` (text)
- `description` (text)
- `eventId` (relation -> events, single)

### interrogations
- `gameId` (relation -> games, single)
- `personId` (relation -> persons, single)
- `question` (text)
- `answer` (text)
- `tone` (text)

## Indexos recomanats
- `games.code` (unique)
- `players.gameId`
- `events.gameId`
- `clues.gameId`
- `persons.gameId`
- `hypotheses.gameId`
- `accusations.gameId`
- `forensics.gameId`
- `timeline.gameId`
- `interrogations.gameId`
