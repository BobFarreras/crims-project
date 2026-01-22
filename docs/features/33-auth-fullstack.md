# 33 - Auth Full-Stack (Frontend + Backend + BD)  

Objetiu
- Implementar un flux d’autenticació autèntic i lligat al backend i la base de dades, amb FDD/TDD.

Abast
- Phase 1 (DOC): Documentar l’estat d’auth, endpoints i esquemes a `docs/features/33-auth-fullstack.md`.
- Phase 2 (TEST): Crear tests de unitat i integració per l’auth (frontend + backend mock).
- Phase 3 (CODE): Implementar el login real amb un backend Go (Go) i BD (PocketBase) o una mock API si no és possible encara, i connectar amb frontend via `frontend/lib/infra/auth-client.ts`.

Estructura de l’arquitectura
- Frontend: UI de Login (LoginFlow + LoginForm), gestió de token al localStorage i navegació al flux del joc.
- Backend: API d’autenticació (POST /auth/login) que retorna token i dades d’usuari.
- BD: PocketBase o alternativa, per emmagatzemar usuaris i sessions.

Plan de desenvolupament (TDD)
- Documentar feature a docs/features/33-auth-fullstack.md (amb fases).
- Escriure tests (unit tests de LoginFlow i LoginForm; tests d’integració simulant backend en mock).
- Implementar backend API i connexió frontend (auth-client.ts) i tests que fallen, després corregir fins passar.

Regles i Consideracions
- Seguretat: no exposem secret keys; token emmagatzemat correctament; protegim rutes amb RBAC si cal.
- UX: pantalla de login amigable, animacions lleus per a mobile (transicions, loading spinners).
- Compatibilitat: mantenir compatibilitat amb l’estructura de feature-based i reuse de components.
