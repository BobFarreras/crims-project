---
name: PocketBase Database Operations
trigger: creating collections, modifying schemas, CRUD operations, realtime subscriptions
scope: backend, database
---

# POCKETBASE DATABASE SKILL 💾

## 1. ¿Qué es PocketBase?

PocketBase es un BaaS (Backend as a Service) open-source escrito en Go que incluye:
- **Database:** Embedded SQLite (fichero único)
- **Authentication:** JWT-based, roles, OAuth
- **Realtime:** WebSockets nativos
- **File Storage:** Integrado, con resizes automáticos
- **Admin Panel:** UI web para gestionar datos
- **REST API:** Automática para todas las colecciones

**URLs del Proyecto:**
- Development: `http://localhost:8090`
- Production: `https://sspb.digitaistudios.com`
- Admin Panel: `https://sspb.digitaistudios.com/_/`

---

## 2. Estructura de Datos del Juego

### Colecciones Principales

#### `games`
Tabla principal de partidas.

**Campos:**
```json
{
  "id": "string (UUID)",
  "code": "string (unique, 4 chars)",          // Ej: "ABCD"
  "status": "enum (LOBBY, INVESTIGATION, ACCUSATION, RESOLUTION)",
  "created_by": "relation (players)",           // Creador de la sala
  "created_at": "datetime",
  "updated_at": "datetime"
}
```

**API Rules:**
- `Create`: `id = @request.auth.id` (solo usuarios autenticados)
- `Read`: `id = @request.auth.id` or `game_id = @request.data.game_id`
- `Update`: `created_by = @request.auth.id`
- `Delete`: `created_by = @request.auth.id`

---

#### `players`
Jugadores conectados a una partida.

**Campos:**
```json
{
  "id": "string (UUID)",
  "game_id": "relation (games)",               // Partida actual
  "user_id": "string (external)",              // ID del usuario (PB auth)
  "role": "enum (DETECTIVE, FORENSIC, ANALYST, INTERROGATOR)",
  "status": "enum (CONNECTED, DISCONNECTED, READY)",
  "score": "number (default 0)",
  "created_at": "datetime",
  "updated_at": "datetime"
}
```

**API Rules:**
- `Create`: `id = @request.auth.id`
- `Read`: `game_id = @request.data.game_id`
- `Update`: `id = @request.auth.id`
- `Delete`: `id = @request.auth.id`

---

#### `clues`
Pistas descubiertas en la escena.

**Campos:**
```json
{
  "id": "string (UUID)",
  "game_id": "relation (games)",
  "type": "enum (PHYSICAL, DIGITAL, TESTIMONY, DOCUMENT)",
  "name": "string",
  "description": "text",
  "location_id": "string",
  "image_url": "string",
  "reliability": "number (0-100)",            // Fiabilidad de la pista
  "state": "enum (DISCOVERED, ANALYZED, VERIFIED)",
  "hidden_truth": "text (encrypted)",          // Solo visible al final
  "facts": "array of objects",                 // Hechos derivados
  "created_at": "datetime"
}
```

**API Rules:**
- `Create`: Solo backend (via servicio)
- `Read`: `game_id = @request.data.game_id` AND `state != DISCOVERED`
- `Update`: Solo backend (via servicio)
- `Delete`: Solo backend (via servicio)

**⚠️ SECURITY:** El campo `hidden_truth` NUNCA debe enviarse al frontend hasta que la partida termine.

---

#### `persons`
Personajes/NPCs del juego.

**Campos:**
```json
{
  "id": "string (UUID)",
  "game_id": "relation (games)",
  "name": "string",
  "role_in_case": "string",                   // "Suspect", "Witness", "Victim"
  "official_story": "text",                    // Lo que dice al jugador
  "truth_story": "text",                       // La verdad (oculta)
  "stress_level": "number (0-100)",            // Estado durante interrogatorio
  "credibility": "number (0-100)",            // Qué tanto confiar
  "avatar_url": "string",
  "created_at": "datetime"
}
```

**API Rules:**
- `Create`: Solo backend (seed de casos)
- `Read`: `game_id = @request.data.game_id` (NO enviar `truth_story`)
- `Update`: Solo backend (durante juego)
- `Delete`: Solo backend

---

#### `hypotheses`
Hipótesis creadas en el tauler.

**Campos:**
```json
{
  "id": "string (UUID)",
  "game_id": "relation (games)",
  "created_by": "relation (players)",
  "statement": "text",                        // "El asesino es X con arma Y"
  "status": "enum (WEAK, PLAUSIBLE, SOLID)", // Calculado automáticamente
  "score": "number",                          // Fuerza numérica
  "evidence_count": "number (default 0)",
  "created_at": "datetime",
  "updated_at": "datetime"
}
```

**API Rules:**
- `Create`: `id = @request.auth.id` AND `role = ANALYST`
- `Read`: `game_id = @request.data.game_id`
- `Update`: `created_by = @request.auth.id`
- `Delete`: `created_by = @request.auth.id`

---

#### `board_edges`
Conexiones entre elementos del tauler.

**Campos:**
```json
{
  "id": "string (UUID)",
  "game_id": "relation (games)",
  "source_id": "string",                       // Puede ser clue_id o person_id
  "target_id": "string",                       // Puede ser hypothesis_id o person_id
  "type": "enum (SUPPORTS, CONTRADICTS, RELATED)",
  "created_by": "relation (players)",
  "created_at": "datetime"
}
```

**API Rules:**
- `Create`: `id = @request.auth.id`
- `Read`: `game_id = @request.data.game_id`
- `Update`: Solo backend (recalcular hipótesis)
- `Delete`: `created_by = @request.auth.id`

---

## 3. Operaciones CRUD desde Backend (Go)

### Conexión a PocketBase

Usa el cliente oficial de Go:

```go
import "github.com/pocketbase/pocketbase"

// Inicializar cliente
pb := pocketbase.New("path/to/pb_data")

// O usar HTTP API (más flexible para producción)
client := pocketbase.NewClient("http://localhost:8090")
```

### Crear Registro

```go
import "github.com/pocketbase/pocketbase/models"

// Crear un juego
game := &models.Record{
    CollectionId: "games",
    Data: map[string]any{
        "code":   "ABCD",
        "status": "LOBBY",
    },
}

err := pb.DB().Create(game)
```

### Leer Registro

```go
// Buscar por ID
game, err := pb.DB().FindById("games", id)

// Buscar con filtros
games, err := pb.DB().FindRecordsByFilter(
    "games",
    "status = 'LOBBY'",
    "-created_at", // Orden descendente
    10,          // Límite
    0,           // Offset
)
```

### Actualizar Registro

```go
game.Data["status"] = "INVESTIGATION"
err := pb.DB().Update(game)
```

### Borrar Registro

```go
err := pb.DB().Delete(game)
```

---

## 4. Realtime Subscriptions

### Backend (Go) - Escuchar Cambios

```go
import "github.com/pocketbase/pocketbase/apis/websocket"

// Suscribirse a cambios en una colección
pb.OnRecordAfterCreateRequest("players").Add(func(e *core.RecordEvent) error {
    // Enviar WebSocket a todos los jugadores del juego
    pb.WebSocket().Broadcast([]byte(fmt.Sprintf(`{
        "type": "PLAYER_JOINED",
        "data": %s
    }`, e.Record))
    return nil
})
```

### Frontend (JavaScript) - Suscribirse

```javascript
import PocketBase from 'pocketbase';

const pb = new PocketBase('http://localhost:8090');

// Suscribirse a cambios en jugadores
pb.collection('players').subscribe('*', (e) => {
    console.log('Cambio en jugadores:', e);

    // e.action: 'create', 'update', 'delete'
    // e.record: el registro modificado

    if (e.action === 'create') {
        // Nuevo jugador se unió
        updatePlayerList(e.record);
    }
});
```

### Patrón de Realtime para Multijugador

```javascript
// Suscribirse a TODOS los cambios de una partida
const subscribeToGame = (gameId) => {
    // Jugadores
    pb.collection('players').subscribe(`game_id = "${gameId}"`, (e) => {
        handlePlayerChange(e);
    });

    // Pistas
    pb.collection('clues').subscribe(`game_id = "${gameId}"`, (e) => {
        handleClueChange(e);
    });

    // Hipótesis
    pb.collection('hypotheses').subscribe(`game_id = "${gameId}"`, (e) => {
        handleHypothesisChange(e);
    });
};

// Cancelar suscripción
const unsubscribe = () => {
    pb.collection('players').unsubscribe();
    pb.collection('clues').unsubscribe();
    pb.collection('hypotheses').unsubscribe();
};
```

---

## 5. API Rules de Seguridad (CRÍTICO)

### Reglas Básicas

**Nunca enviar datos sensibles:**

```yaml
# collection: persons
# rule: List
id != "" AND (game_id = @request.data.game_id)
```

**Excluir campos de respuestas:**

```go
// En el backend, antes de enviar al frontend
type PersonPublic struct {
    ID             string `json:"id"`
    Name           string `json:"name"`
    OfficialStory  string `json:"official_story"`
    // NO enviar TruthStory ni HiddenTruth
}

func PersonToPublic(p *models.Record) PersonPublic {
    return PersonPublic{
        ID:            p.ID,
        Name:          p.Data["name"].(string),
        OfficialStory: p.Data["official_story"].(string),
    }
}
```

### Validación de Permisos

```go
// Verificar que el usuario pertenece al juego
func validateGameAccess(userID, gameID string) error {
    player, err := pb.DB().FindRecordsByFilter(
        "players",
        fmt.Sprintf("user_id = '%s' AND game_id = '%s'", userID, gameID),
        "", 0, 0,
    )
    if len(player) == 0 {
        return errors.New("not authorized")
    }
    return nil
}
```

---

## 6. Migraciones y Schema Management

### Crear una Colección

**Opción A: Via Admin Panel (Fácil)**
1. Ir a `http://localhost:8090/_/`
2. Settings → Collections → New Collection
3. Configurar campos
4. Guardar

**Opción B: Via API (Programático)**
```go
collection := &models.Collection{
    Name: "games",
    Schema: []SchemaField{
        {Name: "code", Type: "text", Required: true},
        {Name: "status", Type: "select", Options: []string{"LOBBY", "INVESTIGATION", "ACCUSATION", "RESOLUTION"}},
    },
}

err := pb.DB().CreateCollection(collection)
```

### Migración de Datos

```go
// Ejemplo: Añadir campo nuevo a registros existentes
migrateGames := func() error {
    games, _ := pb.DB().FindRecordsByFilter("games", "", "", 0, 0)

    for _, game := range games {
        if _, ok := game.Data["new_field"]; !ok {
            game.Data["new_field"] = "default_value"
            pb.DB().Update(game)
        }
    }
    return nil
}
```

---

## 7. File Upload (Evidencias)

### Subir Imagen desde Frontend

```javascript
const formData = new FormData();
formData.append('file', fileInput.files[0]);
formData.append('game_id', gameId);

const record = await pb.collection('evidence').create(formData);
console.log('Archivo subido:', record);
```

### Procesar Imagen (Resize/Thumbnail)

PocketBase lo hace automáticamente si configuras en el Admin Panel:
- Settings → Collections → Field Type: File
- Configurar max size
- Configurar thumbnail sizes (ej: 100x100)

---

## 8. Autenticación

### Registro

```javascript
// Frontend
await pb.collection('users').create({
    email: 'user@example.com',
    password: 'secure_password',
    passwordConfirm: 'secure_password',
});
```

### Login

```javascript
// Frontend
const authData = await pb.collection('users').authWithPassword(
    'user@example.com',
    'secure_password'
);

// authData.token → JWT token
// authData.record → User data
```

### Usar Token en Backend

```go
// Middleware de autenticación
func AuthMiddleware(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        token := r.Header.Get("Authorization")

        // Verificar token con PocketBase
        record, err := pb.DB().FindAuthRecordByToken(
            token,
            pb.Settings().RecordAuth.Token.Duration,
        )

        if err != nil {
            response.Error(w, "Unauthorized", http.StatusUnauthorized)
            return
        }

        // Añadir user al contexto
        ctx := context.WithValue(r.Context(), "user", record)
        next.ServeHTTP(w, r.WithContext(ctx))
    })
}
```

---

## 9. Queries Avanzadas

### Joins (Relaciones)

```go
// Obtener juego con sus jugadores
games, _ := pb.DB().FindRecordsByFilter(
    "games",
    "id = 'abc-123'",
    "", 0, 0,
)

for _, game := range games {
    // Cargar jugadores relacionados
    players, _ := pb.DB().FindRecordsByFilter(
        "players",
        fmt.Sprintf("game_id = '%s'", game.ID),
        "", 0, 0,
    )
}
```

### Agregaciones

```go
// Contar cuántos jugadores hay en un juego
count, _ := pb.DB().Count("players", "game_id = 'abc-123'")
```

### Full-Text Search

```json
// Añadir índice de búsqueda en Admin Panel
{
    "name": "fts_clues",
    "type": "search",
    "fields": ["name", "description"]
}

// Buscar
results, _ := pb.DB().FindRecordsByFilter(
    "clues",
    "fts_clues ~ 'pistol'",
    "", 10, 0,
)
```

---

## 10. Backup y Restore

### Backup Manual

```bash
# Entrar al contenedor Docker
docker exec -it crims-pocketbase sh

# Copiar base de datos
cp /pb_data/data.db /pb_data/backup_$(date +%Y%m%d).db
```

### Backup Automático

```go
// Función de backup programado
func backupDatabase(pb *pocketbase.PocketBase) {
    ticker := time.NewTicker(24 * time.Hour)
    go func() {
        for range ticker.C {
            dbPath := pb.DataDir() + "/data.db"
            backupPath := pb.DataDir() + "/backup_" + time.Now().Format("20060102") + ".db"
            os.CopyFile(dbPath, backupPath, 0644)
        }
    }()
}
```

### Restore

```bash
# Parar PocketBase
docker-compose stop pocketbase

# Restaurar backup
docker cp backup_data.db crims-pocketbase:/pb_data/data.db

# Iniciar PocketBase
docker-compose start pocketbase
```

---

## 11. Performance Tips

### 1. Usar Índices

Crea índices en campos que se usan mucho en filtros:
- `game_id` (usado en casi todas las queries)
- `user_id`
- `status`

### 2. Paginar Resultados

```go
// Siempre usar límite
games, err := pb.DB().FindRecordsByFilter(
    "games",
    "",
    "-created_at",
    20, // Límite
    offset,
)
```

### 3. Cachear Consultas

```go
import "github.com/patrickmn/go-cache"

var cache = go_cache.New(5*time.Minute, 10*time.Minute)

func getGame(id string) (*models.Record, error) {
    // Intentar caché
    if cached, found := cache.Get("game_" + id); found {
        return cached.(*models.Record), nil
    }

    // Consultar DB
    game, err := pb.DB().FindById("games", id)
    if err != nil {
        return nil, err
    }

    // Guardar en caché
    cache.Set("game_" + id, game, go_cache.DefaultExpiration)
    return game, nil
}
```

---

## 12. Common Errors & Solutions

### Error: "Token expired"

**Solución:** Refrescar token
```javascript
await pb.collection('users').authRefresh();
```

### Error: "Record not found"

**Solución:** Verificar ID y permisos
```go
// En backend
game, err := pb.DB().FindById("games", id)
if err != nil {
    return fmt.Errorf("game not found: %w", err)
}
```

### Error: "Not authorized to create"

**Solución:** Verificar API Rules
- Estás autenticado?
- La regla permite crear?

### Error: "Database locked"

**Solución:** Solo un proceso puede escribir a la vez
- Asegúrate de no tener múltiples instancias de PocketBase corriendo
- Usa transacciones para operaciones complejas

---

## 13. Recursos

- **Documentación Oficial:** https://pocketbase.io/docs/
- **API Reference:** https://pocketbase.io/docs/js-overview
- **Admin Panel:** http://localhost:8090/_/
- **GitHub:** https://github.com/pocketbase/pocketbase

---

**Última actualización:** 20/01/2025
**Versión:** 1.0
