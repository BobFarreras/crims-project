# Feature: Autenticació i Seguretat del Sistema

## 1. Arquitectura d'Autenticació
El sistema utilitza un flux d'autenticació basat en **HttpOnly Cookies** per complir amb els estàndards OWASP i prevenir atacs XSS (Cross-Site Scripting).

### Components:
* **Frontend (Next.js):** Gestiona la UI i el Middleware de protecció de rutes (llegeix cookie).
* **Backend (Go):** Actua com a Proxy de Seguretat. Valida credencials i gestiona les Cookies.
* **Base de Dades (PocketBase):** Emmagatzema usuaris i valida contrasenyes (Bcrypt).

---

## 2. Fluxos de Dades

### A. Registre (Sign Up)
1.  **UI:** Formulari amb validació client-side (longitud password, format email).
2.  **Petició:** `POST /api/auth/register` (Payload JSON).
3.  **Backend:**
    * Connecta amb PocketBase (`CreateUser`).
    * Retorna `200 OK` si l'usuari es crea.
4.  **Frontend:** Redirigeix automàticament al Login.

### B. Inici de Sessió (Login) - 🔒 SECURE FLOW
1.  **UI:** Formulari `LoginForm` envia `email` i `password`.
2.  **Petició:** `POST /api/auth/login` amb `credentials: 'include'`.
3.  **Backend:**
    * Valida credencials contra PocketBase (`AuthWithPassword`).
    * Rep un JWT Token de PocketBase.
    * **Acció Crítica:** Genera una cookie `auth_token` amb el JWT.
    * Configuració Cookie: `HttpOnly: true`, `SameSite: Lax`, `Path: /`.
4.  **Resposta:** JSON amb `message` i objecte `user` (sense exposar el token).
5.  **Navegador:** Emmagatzema la cookie de forma inaccessible per a JavaScript.

### C. Protecció de Rutes (Middleware)
Un fitxer `middleware.ts` s'executa a cada petició a Next.js:
* **Llegeix la Cookie:** Verifica si existeix `auth_token`.
* **Rutes Privades (`/game/*`):** Si no hi ha cookie -> Redirect a `/login`.
* **Rutes Públiques (`/login`, `/register`):** Si hi ha cookie -> Redirect a `/game/dashboard`.

---

## 3. Mesures de Seguretat Implementades
| Amenaça | Solució Implementada |
| :--- | :--- |
| **XSS (Robatori de Token)** | El token està en una Cookie **HttpOnly**. El JS maliciós no la pot llegir. |
| **CSRF (Falsificació)** | Cookie configurada amb **SameSite=Lax** en entorn same-site. En producció, usar `SameSite=None` + `Secure` per domini separat. |
| **Enumeració d'Usuaris** | Missatges d'error genèrics al Login ("Credencials invàlides"). |
| **Intercepció** | Backend preparat per activar `Secure: true` (només HTTPS) en producció. |

---

## 4. Estructura de Codi Clau
* `frontend/middleware.ts`: El "porter" que vigila les rutes.
* `frontend/features/auth/hooks/useLogin.ts`: Gestió d'estats de UI (sense tocar tokens).
* `backend/internal/adapters/http/auth_handlers.go`: Generació de la Cookie segura.
* `backend/cmd/server/main.go`: Configuració CORS (`AllowCredentials: true`).
