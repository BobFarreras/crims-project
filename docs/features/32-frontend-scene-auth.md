# 32 - Frontend Scene + Auth UI (Inicial)

## Objectiu
- Implementar la UI mínima per a Scene (exploració 3D simulada) i una UI d Autenticació bàsica (Login) en el frontend.

## Abast
- Nova ruta de joc: `/game/scene` amb una vista de SceneViewport (mock).
- Nova ruta d'Autenticació: `/login` dins l’espai `(auth)` (layout existent) amb un LoginForm bàsic.
- Components modulars sota `frontend/features/scene` i `frontend/features/auth`.
- Tests unitats per SceneViewport i LoginForm (renderització bàsica).

## Requeriments
- UI simplificada i responsive per mobile/desktop.
- SceneViewport amb contingut de placeholder i estil coherent amb la resta de UI.
- LoginForm amb dos camps (username, password) i botó de login.
- Tests que fallen abans de la implementació i passen després.

## Criteris d'Aceptació
- Renderització del títol i dels components SceneViewport i LoginForm.
- SceneViewport mostra el text de mock utilitzat i és present a la pàgina `/game/scene`.
- LoginForm mostra dos camps i un botó, i és possible interaccionar amb ells.

## Pla de Tests (TDD)
- 1) Documentar feature a `docs/features/32-frontend-scene-auth.md` (ja fet).
- 2) Escriviu tests per SceneViewport i LoginForm (falsos al principi).
- 3) Implementar el codi fins passar tots els tests.

## Pla d'Implementació
- Afegeix la ruta `/game/scene` i el component `SceneViewport` (mock).
- Afegeix la ruta `/login` dins `(auth)` i el component `LoginForm` (form bàsic).
- Afegeix tests per ambdós components i implementa el codi per passar-los.

## Impacte i Riscos
- Impacte mínim en l’arquitectura actual; elimina dependències externes.
- L’UI és mock/placeholder, pensada per iterar cap a frontend 3D real en fases posteriors.

## Treball Futur
- Integrar SceneViewport amb un motor 3D (Three.js) quan sigui viable.
- Substituir LoginForm mock per autenticació real (PocketBase/Auth).
