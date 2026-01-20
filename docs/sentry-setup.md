# SENTRY CONFIGURATION & SETUP 📊

## 1. ¿Qué es Sentry?

Sentry es una plataforma de **error tracking** i **performance monitoring** que:
- Captura errors de producció en temps real
- Proporciona stack traces detallats
- Permet veure sessions d'usuaris (replays)
- Monitora performance i traces
- Envia alertes immediates

**Projecte CRIMS:**
- Org: `digitaistudios` (canvia-ho)
- Frontend: `crims-frontend`
- Backend: `crims-backend`

---

## 2. Configuración Inicial

### Paso 1: Crear Cuenta en Sentry

1. Ve a https://sentry.io/signup
2. Crea compte o inicia sesión
3. Crea una **Organización**: `digitaistudios`

### Paso 2: Crear Proyecto Frontend

1. Click "Create Project"
2. Selecciona: **Next.js**
3. Nombre: `crims-frontend`
4. Platform: **JavaScript/TypeScript**
5. Copia el **DSN** (Data Source Name) que te da Sentry

### Paso 3: Crear Proyecto Backend

1. Click "Create Project"
2. Selecciona: **Go**
3. Nombre: `crims-backend`
4. Platform: **Go**
5. Copia el **DSN** (será diferente al del frontend)

---

## 3. Configurar Variables de Entorno

### Frontend (`.env.local`)

```bash
# Sentry DSN para frontend (NEXT_PUBLIC para acceso en browser)
NEXT_PUBLIC_SENTRY_DSN=https://xxx@sentry.io/xxx

# Environment
NEXT_PUBLIC_ENVIRONMENT=development
```

### Backend (`.env.local`)

```bash
# Sentry DSN para backend
SENTRY_DSN=https://xxx@sentry.io/xxx

# Environment
ENVIRONMENT=development
```

### Production (Vercel/VPS)

Añade las variables en el panel de tu hosting:
- **Vercel:** Settings → Environment Variables
- **Docker:** En el `docker-compose.yml` o `.env`

---

## 4. Archivos de Configuración

### Frontend (Next.js)

**Archivos creados:**
- `frontend/sentry.client.config.ts` - Configuración browser
- `frontend/sentry.server.config.ts` - Configuración servidor
- `frontend/sentry.edge.config.ts` - Configuración edge
- `frontend/next.config.ts` - Integración con Next.js
- `frontend/.sentryclirc` - Configuración CLI

**Sample Rates configurados:**
- Traces: 10% (1 de cada 10 peticiones)
- Replays (session): 10%
- Replays (error): 100% (si hay error, grabar todo)

**Filtros de errores:**
- Ignorar errores de "ChunkLoadError" (development)
- Ignorar errores de red
- No enviar en development

### Backend (Go)

**Archivo modificado:**
- `backend/cmd/server/main.go` - Inicialización de Sentry

**Sample Rates configurados:**
- Traces: 10%

**Features activadas:**
- Panic capturing (middleware Sentry)
- Tags globales (app, runtime, framework)
- Flush automático antes de salir

---

## 5. Usar Sentry en el Código

### Frontend (TypeScript/JavaScript)

#### Capturar Error Manualmente

```typescript
import * as Sentry from "@sentry/nextjs";

try {
  const result = await riskyOperation();
} catch (error) {
  // Capturar error en Sentry
  Sentry.captureException(error);
}
```

#### Capturar Mensaje Manualmente

```typescript
// Capturar un mensaje (no exception)
Sentry.captureMessage("Usuario intentó acceso no autorizado", "warning");
```

#### Añadir Contexto (Breadcrumbs)

```typescript
import { setContext } from "@sentry/nextjs";

// Añadir contexto antes de capturar error
setContext("game_state", {
  gameId: "abc-123",
  currentScene: "crime-scene-2",
  playerRole: "DETECTIVE",
  timestamp: Date.now(),
});

// Luego capturar error
Sentry.captureException(error);
```

#### Añadir Tags

```typescript
import { setTag } from "@sentry/nextjs";

// Añadir tags para filtrar errores
setTag("feature", "investigation-board");
setTag("user_role", "ANALYST");

Sentry.captureException(error);
```

#### Capturar Error con Nivel de Severidad

```typescript
Sentry.captureException(error, {
  level: "error", // "fatal", "error", "warning", "log", "info", "debug"
});
```

### Backend (Go)

#### Capturar Error Manualmente

```go
import "github.com/getsentry/sentry-go"

try {
    result, err := riskyOperation()
    if err != nil {
        // Capturar error en Sentry
        sentry.CaptureException(err)
    }
} catch (e) {
    // Capturar panic en Sentry
    sentry.Recover()
}
```

#### Capturar Mensaje Manualmente

```go
// Capturar un mensaje
sentry.CaptureMessage("Jugador intentó acceder a juego inexistente", sentry.LevelWarning)
```

#### Añadir Contexto

```go
import "github.com/getsentry/sentry-go"

// Añadir contexto al scope actual
sentry.ConfigureScope(func(scope *sentry.Scope) {
    scope.SetContext("game", map[string]interface{}{
        "game_id": "abc-123",
        "scene_id": "crime-scene-2",
        "player_count": 4,
    })

    // Añadir tags
    scope.SetTag("feature", "investigation-board")
    scope.SetTag("user_role", "ANALYST")
})

// Capturar error con contexto
sentry.CaptureException(err)
```

#### Capturar con Nivel de Severidad

```go
sentry.CaptureException(err, &sentry.EventHint{
    Level: sentry.LevelError,
})
```

---

## 6. Ver Errores en Sentry

### Dashboard de Sentry

1. Ve a https://sentry.io
2. Selecciona tu organización: `digitaistudios`
3. Selecciona el proyecto: `crims-frontend` o `crims-backend`
4. En la página "Issues", verás todos los errores

### Tipos de Errores

- **Issues:** Errores recurrentes (agrupados)
- **Alerts:** Reglas configuradas para notificaciones
- **Performance:** Traces de rendimiento
- **Replays:** Grabaciones de sesión (Frontend)

### Filtrar Errores

Puedes filtrar por:
- **Environment:** `production`, `staging`, `development`
- **Tags:** `app`, `feature`, `user_role`
- **Date Range:** Última hora, día, semana
- **First Seen:** Nuevos vs antiguos

---

## 7. Crear Alertas

### Alerta por Error Crítico

1. Ve a: **Settings** → **Alerts**
2. Click **Create Alert Rule**
3. Configurar:
   - **Condition:** New Issue created
   - **Environment:** production
   - **Frequency:** Every time
4. Configurar notificaciones:
   - Email, Slack, etc.

### Alerta por Performance Degradation

1. Ve a: **Performance** → **Alerts**
2. Configurar:
   - **Condition:** Response time > 5s
   - **Environment:** production

---

## 8. Debugging con Sentry

### Ver Stack Trace Completo

1. Abre un error en Sentry
2. Verás el **Stack Trace** completo
3. Puedes ver:
   - Línea exacta del error
   - Valores de variables (si están en el scope)
   - Breadcrumbs (acciones previas del usuario)

### Ver Session Replays (Frontend)

1. Abre un error en Sentry
2. Click en "Replay"
3. Verás una grabación de vídeo de la sesión del usuario
4. Puedes ver:
   - Qué hizo el usuario
   - Qué hizo click
   - Qué escribió

### Ver Traces de Performance

1. Ve a: **Performance** → **Transactions**
2. Verás cada petición HTTP
3. Puedes ver:
   - Tiempo total
   - Tiempo de cada componente
   - Queries de DB
   - Latencia de API

---

## 9. Best Practices

### ✅ Buenas Prácticas

1. **Sample Rates Razonables:**
   - Development: 0% (no enviar nada)
   - Staging: 100% (probar todo)
   - Production: 10-20% (ahorrar cuota)

2. **Añadir Contexto Siempre:**
   - Agrega `user_id`, `game_id`, `feature` a cada error

3. **Ignorar Errores No Críticos:**
   - Network errors (los maneja el cliente)
   - ChunkLoadError en development

4. **Configurar Alertas:**
   - Alertas inmediatas para production
   - Alertas resumidas para staging

5. **Revisar Sentry Regularmente:**
   - Al menos 1 vez por semana
   - Priorizar errores más frecuentes

### ❌ Malas Prácticas

1. Enviar datos personales (GDPR):
   ```typescript
   // ❌ MAL
   Sentry.captureException(error, {
     extra: {
       user_password: password, // NO ENVIAR
     }
   })
   ```

2. Enviar información confidencial:
   ```typescript
   // ❌ MAL
   Sentry.captureException(error, {
     extra: {
       api_key: "sk-12345", // NO ENVIAR
     }
   })
   ```

3. Configurar sample rate al 100% en production:
   - Gastará tu cuota de Sentry muy rápido

---

## 10. Troubleshooting

### Error: "DSN is not configured"

**Problema:** No has configurado la variable de entorno.

**Solución:**
```bash
# En .env.local
NEXT_PUBLIC_SENTRY_DSN=https://xxx@sentry.io/xxx

# Reiniciar servidor
pnpm dev
```

### Error: "No events received"

**Problema:** Sentry está recibiendo eventos, pero no los está mostrando.

**Solución:**
1. Verifica que la variable de entorno es correcta
2. Verifica que no estás en "development" (si filtraste por environment)
3. Revisa los filtros en Sentry

### Errores de CORS

**Problema:** Sentry no puede enviar eventos desde el browser.

**Solución:**
1. Ve a: Sentry → Project → Settings → Client Keys (DSN)
2. Verifica que el dominio está en "Allowed Domains"
3. Añade: `http://localhost:3000`, `https://www.crimsdemitjanit.com`

### Errores de Sourcemaps

**Problema:** Los stack traces no tienen los números de línea correctos.

**Solución:**
1. Asegúrate de generar sourcemaps en el build:
   ```typescript
   // next.config.ts
   export default withSentryConfig(nextConfig, {
     hideSourceMaps: false, // Ocultar en production
     tunnelRoute: "/monitoring",
   });
   ```

2. Sube los sourcemaps a Sentry:
   ```bash
   pnpm sentry-upload-sourcemaps
   ```

---

## 11. Costos y Límites

### Plan Gratuito (Sentry)

- **5,000 errors/month**
- **10,000 performance transactions/month**
- **1 GB de attachments**
- **Retención de datos:** 30 días

### Plan Pro

- **100,000 errors/month**
- **1,000,000 performance transactions/month**
- **50 GB de attachments**
- **Retención de datos:** 90 días

### Recomendación para CRIMS

Comienza con el **plan gratuito**:
- Sample rate: 10%
- Debería ser suficiente para MVP
- Monitorea el uso mensual

---

## 12. Recursos

- **Documentación:** https://docs.sentry.io/
- **JavaScript SDK:** https://docs.sentry.io/platforms/javascript/
- **Go SDK:** https://docs.sentry.io/platforms/go/
- **Next.js Integration:** https://docs.sentry.io/platforms/javascript/guides/nextjs/
- **Dashboard:** https://sentry.io/

---

**Última actualización:** 20/01/2025
**Versión:** 1.0
