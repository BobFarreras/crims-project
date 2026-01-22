# 09 - Person Service (Backend)

## Objectiu
Afegir la capa de dades i servei per gestionar persones (testimonis i sospitosos) per partida.

## Abast
- Repositori `PersonRepository` (ports + adapter PocketBase).
- Servei `PersonService` amb validacions.
- Endpoints HTTP per crear i llistar persones per partida.

## Dades
`Person`:
- `id`
- `gameId`
- `name`
- `officialStory`
- `truthStory`
- `stress` (0-100)
- `credibility` (0-100)

## Requeriments
1. `CreatePerson` valida camps obligatoris.
2. `GetPersonByID` valida `id`.
3. `ListPersonsByGame` valida `gameId`.

## Criteris d'Aceptacio
- POST retorna 201 amb persona creada.
- GET per id retorna 200.
- GET per gameId retorna llista.

## Pla de Tests (Integracio)
1. Repo: create OK / error.
2. Repo: list per game OK.
3. Service: invalid inputs.
4. Handlers: status codes i JSON correctes.
