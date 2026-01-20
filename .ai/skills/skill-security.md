---
name: OWASP & Security Hardening
trigger: creating endpoints, handling user input, authentication
scope: global
---

# SECURITY FIRST PROTOCOL (OWASP Top 10)

## 1. Principi de "Zero Trust" (Backend)
Mai confiïs en el que envia el Frontend. El client pot ser manipulat (ex: un jugador hackejant el JS per veure pistes ocultes).
* **Validació:** Tota dada d'entrada (JSON) s'ha de validar amb una llibreria estricta (ex: `go-playground/validator`).
* **Sanitització:** Neteja strings per evitar XSS abans de guardar-los, encara que React ja protegeix bastant.

## 2. Autenticació i Autorització (RBAC)
No n'hi ha prou amb estar loguejat. Has de tenir el **Rol** correcte.
* **Bad:** `if user != nil { return secret }`
* **Good:** `if user.Role != "FORENSIC" { return ErrorForbidden }`
* **Broken Object Level Authorization (BOLA):** Verifica sempre que l'usuari pertany a la `game_id` que està intentant tocar.

## 3. Injecció (SQL/NoSQL)
* Mai concatenis strings per fer consultes a PocketBase o SQL.
* Usa sempre paràmetres vinculats (Prepared Statements) o els mètodes segurs de l'ORM/SDK.

## 4. Rate Limiting & DoS
* Els WebSockets/SSE són cars. Protegeix els endpoints de "spam" d'accions (ex: moure una fitxa 100 cops per segon).

## 5. Dades Sensibles
* Mai retornis el camp `hidden_truth` d'un NPC al frontend fins que no s'hagi desbloquejat lògicament.
* Si ho envies al JSON (encara que no ho pintis), un hacker ho veurà a la pestanya Network.