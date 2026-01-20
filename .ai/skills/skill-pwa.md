---
name: PWA Manifest & Service Worker
trigger: creating PWA manifest, service worker, configuring installation
scope: frontend, mobile
---

# PWA MANIFEST & SERVICE WORKER SKILL 📱

## 1. ¿Qué es una PWA?

**PWA** (Progressive Web App) és una aplicació web que:
- Funciona offline
- Es pot instal·lar en el dispositiu (com app nadiua)
- Té icona a la pantalla d'inici
- Pot rebre notificacions push
- S'actualitza automàticament en segon pla

**Plataformes suportades:**
- ✅ **Desktop:** Chrome, Edge, Firefox, Safari, Opera
- ✅ **Mòbil:** Android (Chrome), iOS (Safari)
- ✅ **Multiplataforma:** React Native, Ionic, Capacitor

---

## 2. Estrategia de Multiplataforma

### Enfoc Actual (Web Reactiva)
```
Web Reactiva (Next.js) + PocketBase (Backend)
    ↓
    PWA Manifest + Service Worker
    ↓
Instal·lació al navegador
    ↓
Funcionament offline limitat (cache d'assets)
```

### Futur (Multiplataforma Nadiua)
```
Web (PWA Manifest + SW)
    ↓
    React Native (iOS/Android)
    ↓
Instal·lació a App Store / Google Play
    ↓
API compartida (Go + PocketBase)
```

**Per què NO implementar ara Android/iOS:**
- El joc és molt complex (realtime multiplayer)
- Requereix molt més temps de desenvolupament
- La PWA web ja proveeix l'experiència bàsica
- En un futur, es pot evolucionar a apps nadius amb Kotlin
- **NO és urgent:** La web ja funciona bé en mòbils

---

## 3. PWA Manifest (frontend/public/manifest.json)

### Configuració Bàsica

**Què és:** Un fitxer JSON que descriu l'app.

**Exemple per CRIMS:**
```json
{
  "name": "CRIMS de Mitjanit",
  "short_name": "CRIMS",
  "description": "Plataforma de joc d'investigació criminal multijugador en temps real",
  "start_url": "/",
  "display": "standalone",
  "background_color": "#1a1a2e",
  "theme_color": "#3b82f6",
  "orientation": "portrait",
  "icons": [
    {
      "src": "/icons/icon-72x72.png",
      "sizes": "72x72",
      "type": "image/png",
      "purpose": "any maskable"
    },
    {
      "src": "/icons/icon-96x96.png",
      "sizes": "96x96",
      "type": "image/png",
      "purpose": "any maskable"
    },
    {
      "src": "/icons/icon-128x128.png",
      "sizes": "128x128",
      "type": "image/png",
      "purpose": "any maskable"
    },
    {
      "src": "/icons/icon-144x144.png",
      "sizes": "144x144",
      "type": "image/png",
      "purpose": "any maskable"
    },
    {
      "src": "/icons/icon-152x152.png",
      "sizes": "152x152",
      "type": "image/png",
      "purpose": "any maskable"
    },
    {
      "src": "/icons/icon-192x192.png",
      "sizes": "192x192",
      "type": "image/png",
      "purpose": "any maskable"
    },
    {
      "src": "/icons/icon-384x384.png",
      "sizes": "384x384",
      "type": "image/png",
      "purpose": "any maskable"
    },
    {
      "src": "/icons/icon-512x512.png",
      "sizes": "512x512",
      "type": "image/png",
      "purpose": "any maskable"
    },
    {
      "src": "/icons/maskable_icon.png",
      "sizes": "any",
      "type": "image/png",
      "purpose": "maskable"
    }
  ],
  "shortcuts": [
    {
      "name": "Jugar",
      "short_name": "Jugar",
      "description": "Iniciar una nova partida",
      "url": "/lobby",
      "icons": [
        {
          "src": "/icons/shortcut-jugar.png",
          "sizes": "96x96"
        }
      ]
    }
  ],
  "screenshots": [
    {
      "src": "/screenshots/desktop.png",
      "sizes": "1280x720",
      "type": "image/png",
      "form_factor": "wide"
    },
    {
      "src": "/screenshots/mobile.png",
      "sizes": "750x1334",
      "type": "image/png",
      "form_factor": "narrow"
    }
  ],
  "categories": ["games", "social", "entertainment"]
}
```

### Camps Importants

| Camp | Descripció | Valor CRIMS |
|-------|------------|--------------|
| `name` | Nom complet de l'aplicació | CRIMS de Mitjanit |
| `short_name` | Nom curt per a la pantalla | CRIMS |
| `description` | Descripció per la store | Plataforma de joc d'investigació... |
| `start_url` | URL d'inici | / |
| `display` | Com es mostra | standalone (sense barra del navegador) |
| `background_color` | Color de fons | #1a1a2e (fosc) |
| `theme_color` | Color de la barra | #3b82f6 (blau CRIMS) |
| `orientation` | Orientació | portrait (vertical) |
| `icons` | Icones | Diferents mides (72-512px) |
| `categories` | Categories | games, social, entertainment |
| `shortcuts` | Shortcuts | Accions ràpides (Jugar, etc.) |
| `screenshots` | Captures de pantalla | Per la store |

---

## 4. Service Worker (frontend/public/sw.js)

### Què és?

El **Service Worker** és un script que s'executa en background (separ del main thread).

**Funcions:**
- Caching d'assets per funcionar offline
- Interceptar peticions network
- Sync de dades en background
- Notificacions push (opcional)

### Implementació Bàsica

```javascript
// frontend/public/sw.js
const CACHE_NAME = 'crims-cache-v1';
const STATIC_CACHE = 'crims-static-v1';

// Assets que sempre hem de cachear
const ASSETS_TO_CACHE = [
  '/',
  '/lobby',
  '/board',
  '/scene',
  '/interrogation',
  '/timeline',
  '/forensic',
  '/accusation',
  '/manifest.json',
  '/icons',
];

// Instal·lar Service Worker
self.addEventListener('install', (event) => {
  console.log('📱 Service Worker instal·lat');

  // Precachear assets estàtics
  event.waitUntil(
    caches.open(STATIC_CACHE).then((cache) => {
      return cache.addAll(ASSETS_TO_CACHE.map((url) => new Request(url, { cache: 'reload' })));
    })
  );

  // Forçar als clients antics a actualitzar
  self.clients.claim();
});

// Activar Service Worker
self.addEventListener('activate', (event) => {
  console.log('📱 Service Worker activat');
  event.waitUntil(
    caches.keys().then((cacheNames) => {
      return Promise.all(
        cacheNames.map((cacheName) => {
          if (cacheName !== STATIC_CACHE && cacheName !== CACHE_NAME) {
            console.log('🗑️  Eliminant cache antic:', cacheName);
            return caches.delete(cacheName);
          }
        })
      );
    })
  );
});

// Interceptar peticions (caching)
self.addEventListener('fetch', (event) => {
  event.respondWith(
    caches.match(event.request).then((response) => {
      // Si hi ha cache, retornar-lo
      if (response) {
        return response;
      }

      // Si no hi ha cache, fer fetch
      return fetch(event.request).then((networkResponse) => {
        // Cachejar peticions GET d'assets estàtics
        if (
          event.request.method === 'GET' &&
          ASSETS_TO_CACHE.includes(new URL(event.request.url).pathname)
        ) {
          const responseToCache = networkResponse.clone();
          caches.open(CACHE_NAME).then((cache) => {
            cache.put(event.request, responseToCache);
          });
        }

        return networkResponse;
      });
    })
  );
});

// Gestió de cache quan hi ha actualització
self.addEventListener('message', (event) => {
  if (event.data === 'SKIP_WAITING') {
    self.skipWaiting();
  }
});
```

---

## 5. Integració amb Next.js 15

### Registra el Service Worker

**A `app/layout.tsx`:**
```typescript
import Script from 'next/script';

export default function RootLayout({ children }: { children: React.ReactNode }) {
  return (
    <html lang="ca">
      <head>
        <link rel="manifest" href="/manifest.json" />
        <title>CRIMS de Mitjanit</title>
        <Script
          id="register-sw"
          strategy="beforeInteractive"
        >
          {`
            if ('serviceWorker' in navigator) {
              window.addEventListener('load', () => {
                navigator.serviceWorker.register('/sw.js', {
                  scope: '/',
                  updateViaCache: 'all',
                }).then((registration) => {
                  console.log('✅ Service Worker registrat:', registration.scope);
                  
                  // Comprovar actualitzacions
                  registration.addEventListener('updatefound', () => {
                    console.log('🔄 Service Worker actualitzat');
                  });
                  
                  // Forçar actualització
                  setInterval(() => {
                    registration.update();
                  }, 60 * 60 * 1000); // Cada hora
                });
              });
            }
          `}
        </Script>
        <meta name="theme-color" content="#3b82f6" />
        <meta name="viewport" content="width=device-width, initial-scale=1" />
      </head>
      <body>
        {children}
      </body>
    </html>
  );
}
```

### Configuració de Next.js

**A `next.config.ts`:**
```typescript
const nextConfig: NextConfig = {
  images: {
    remotePatterns: [
      {
        protocol: 'https',
        hostname: 'sspb.digitaistudios.com',
        pathname: '/api/files/**',
      },
    ],
  },
};
```

---

## 6. Eines i Recursos

### Crear Icones de PWA

**Eines online:**
- [PWA Asset Generator](https://www.pwabuilder.com/imageGenerator) - Generar totes les mides
- [RealFaviconGenerator](https://realfavicongenerator.net/) - Generar favicon i PWA icons
- [Favicon.io](https://favicon.io/) - Generar favicon ràpid

**Dimensions necessàries:**
- 72x72 px (Android)
- 96x96 px (iOS)
- 128x128 px (iPad)
- 144x144 px (Android HDPI)
- 152x152 px (Windows)
- 192x192 px (Mac)
- 384x384 px (Windows HDPI)
- 512x512 px (Windows Ultra HDPI)

**Tipus de fitxer:**
- `png` - Recomanat per la major compatibilitat
- `favicon.ico` - Per navegadors antics
- `svg` - Per icones de qualitat (opcional)

### Provar PWA

**Eines:**
- [Lighthouse](https://developers.google.com/web/tools/lighthouse) - Audit de PWA
- [PWA Builder](https://www.pwabuilder.com/) - Generar PWA
- [Manifest Validator](https://manifest-validator.appspot.com/) - Validar manifest.json

**Provar en dispositiu real:**
- Chrome DevTools → Application → Service Workers
- Safari DevTools → Cache → Service Workers
- Android: Chrome Remote Debugging
- iOS: Safari Web Inspector

---

## 7. Implementació Pas a Pas

### PAS 1: Crear Icones

1. Utilitza [PWA Asset Generator](https://www.pwabuilder.com/imageGenerator)
2. Carrega el logotip de CRIMS
3. Genera totes les mides (72-512px)
4. Descarrega el ZIP amb totes les icones
5. Descomprimeix a `frontend/public/icons/`

```bash
# Crear carpeta d'icones
mkdir -p frontend/public/icons

# Copiar les icones generades al directori
# (Copia els fitxers del ZIP descarregat)
```

### PAS 2: Crear manifest.json

Crea el fitxer `frontend/public/manifest.json` amb l'estructura de dalt.

### PAS 3: Crear Service Worker

Crea el fitxer `frontend/public/sw.js` amb el codi de dalt.

### PAS 4: Actualitzar layout.tsx

Afegeix:
- `<link rel="manifest" href="/manifest.json" />`
- Component `<Script>` per registrar el Service Worker
- Meta tags per colors de tema

### PAS 5: Crear Captures de Pantalla

Genera captures per la store:
1. Versió desktop (1280x720)
2. Versió mòbil (750x1334)
3. Guarda a `frontend/public/screenshots/`

### PAS 6: Provar Localment

```bash
cd frontend
pnpm dev
```

Obre Chrome DevTools → Application i revisa:
- Manifest carregat correctament
- Service Worker registrat
- Caches funcionant
- No errors a la consola

### PAS 7: Audit amb Lighthouse

```bash
# Instalar Lighthouse CLI
pnpm add -D @lhci/cli

# Audit de PWA
lhci autor --view=pwa http://localhost:3000
```

Ha de passar:
- ✅ PWA: 90-100
- ✅ Installable: Sí
- ✅ Manifest: Sí
- ✅ Service Worker: Sí
- ✅ Offline: Sí

---

## 8. Funcionalitats Avançades (Opcionals)

### Background Sync

**Què és:** El Service Worker pot sincronitzar dades en background sense que l'usuari estigui a l'app.

**Implementació:**
```javascript
// frontend/public/sw.js
const SYNC_INTERVAL = 5 * 60 * 1000; // 5 minuts

self.addEventListener('sync', (event) => {
  const { url, headers } = event;
  
  fetch(url, { headers })
    .then(response => response.json())
    .then(data => {
      // Enviar a pocketbase
      fetch('https://sspb.digitaistudios.com/api/sync', {
        method: 'POST',
        headers: {
          'Content-Type': 'application/json',
          'Authorization': \`Bearer \${headers.get('authorization')}\`,
        },
        body: JSON.stringify({
          gameId: data.gameId,
          lastSync: data.lastSync,
        }),
      });
    });
});

// Programar sync
self.setInterval(() => {
  self.registration.showNotification('Sincronitzant...', {
    body: 'Actualitzant dades del joc',
    icon: '/icons/icon-96x96.png',
  });
}, SYNC_INTERVAL);
```

### Notificacions Push

**Què és:** Enviar notificacions als usuaris (esdeveniments del joc, torn, etc.).

**Implementació:**
```javascript
// frontend/public/sw.js
self.addEventListener('push', (event) => {
  const { data } = event;
  
  self.registration.showNotification('Nova actualització!', {
    body: data.message,
    icon: '/icons/icon-96x96.png',
    badge: '/icons/badge.png',
    tag: 'game-update',
  });
});
```

### Offline Strategy

**Què és:** Definir què passa quan l'usuari està offline.

**Estrategies:**
```javascript
// Cache-First Strategy
self.addEventListener('fetch', (event) => {
  event.respondWith(
    caches.match(event.request).then((response) => {
      // Primer mirar al cache, després a la xarxa
      if (response) {
        return response;
      }
      
      // Si no hi ha cache, anar a la xarxa
      return fetch(event.request);
    })
  );
});

// Network-First Strategy (per API)
self.addEventListener('fetch', (event) => {
  if (event.request.url.includes('api/')) {
    // Per peticions API, anar primer a la xarxa
    return fetch(event.request);
  }
  
  // Per assets estàtics, usar cache-first
  return caches.match(event.request).then((response) => {
    return response || fetch(event.request);
  });
});
```

---

## 9. Troubleshooting

### Service Worker no s'actualitza

**Problema:** El Service Worker es queda amb una versió antiga.

**Solució:**
```javascript
// En app/layout.tsx
registration.addEventListener('updatefound', () => {
  console.log('🔄 Service Worker actualitzat');
  // Notificar a l'usuari (opcional)
  if ('serviceWorker' in navigator) {
    navigator.serviceWorker.controller.postMessage('SKIP_WAITING');
  }
});

// En sw.js
self.addEventListener('message', (event) => {
  if (event.data === 'SKIP_WAITING') {
    self.skipWaiting();
  }
});
```

### Cache no funciona offline

**Problema:** L'aplicació no funciona offline.

**Solució:**
1. Verificar que l'aplicació és cachejada
2. Utilitzar DevTools → Application → Cache Storage
3. Verificar què assets estan al cache
4. Provar en mode offline (Desactivar internet)
5. Revisar els logs del Service Worker

### No s'instal·la com a PWA

**Problema:** L'usuari no pot instal·lar l'aplicació.

**Solucions:**
1. **Manifest JSON:** Verificar que és JSON vàlid
2. **Service Worker:** Verificar que s'ha registrat correctament
3. **HTTPS:** Requisit per PWA (funciona a localhost amb flags)
4. **Icons:** Verificar totes les mides estan presents
5. **Scope:** Verificar el scope del Service Worker

### Manifest no es detecta

**Problema:** Chrome no detecta el manifest.

**Solució:**
1. Asegurar que el `<link rel="manifest">` està al `<head>`
2. Verificar que el manifest.json està a `/public/` (no `/app/`)
3. Revisar la sintaxi del JSON (comas i cometes)
4. Obre Chrome DevTools → Application → Manifest
5. Veure si hi ha errors

---

## 10. Best Practices

### ✅ DO'S

1. **Cachejar només el necessari**
   - No cachejar API calls dinàmiques
   - No cachejar dades personals dels usuaris
   - Limitar la mida del cache (50MB màx)

2. **Actualització progressiva**
   - El Service Worker s'ha d'actualitzar automàticament
   - Utilitzar `updateViaCache: 'all'`
   - Notificar als usuaris quan hi ha actualització

3. **Provar en dispositius reals**
   - Emulatoris no són 100% fiables
   - Provar en mòbils Android i iOS
   - Verificar que funciona offline

4. **Audit regularment**
   - Utilitzar Lighthouse mensualment
   - Provar el PWA score (>90)
   - Corregir problemes immediatament

5. **Monitorització**
   - Utilitzar Google Analytics
   - Utilitzar Sentry per errors del Service Worker
   - Monitoritzar cache hits/misses

### ❌ DON'TS

1. **No utilitzar caches excessives**
   - No cachejar tots els assets
   - El Service Worker ha de ser lleuger
   - No utilitzar `*` en `caches.open()`

2. **No actualitzar el Service Worker**
   - Les versions antigues causen bugs
   - Implementar sempre actualitzacions automàtiques

3. **No provar offline**
   - La PWA ha de funcionar offline
   - Provar en mode avió després de provar en mode terra

4. **No utilitzar HTTPS en desenvolupament**
   - El Service Worker requereix HTTPS (excepte localhost)
   - Provar amb `http://localhost:3000` però preparar-se per HTTPS

---

## 11. Checklist per a PWA Producció

### Antes del desplegament

- [ ] Totes les icones generades (72-512px)
- [ ] manifest.json creat i validat
- [ ] Service Worker implementat i provat
- [ ] Layout actualitzat amb `<link rel="manifest">`
- [ ] Service Worker registrat correctament
- [ ] Funcionament offline provat
- [ ] Lighthouse audit >90
- [ ] Captures de pantalla creades
- [ ] Notificacions push (opcional)
- [ ] HTTPS configurat en producció

### Després del desplegament

- [ ] Manifest es detectat correctament
- [ ] Service Worker s'actualitza
- [ ] App s'instal·la correctament
- [ ] Funciona offline
- [ ] No errors a la consola
- [ ] Lighthouse audit >90 en producció

---

## 12. Recursos

### Documentació Oficial
- [PWA Documentation](https://web.dev/progressive-web-apps/)
- [Next.js PWA Guide](https://nextjs.org/docs/app/building-your-application/optimizing/production)
- [Service Worker API](https://developer.mozilla.org/en-US/docs/Web/API/Service_Worker_API)
- [PWA Best Practices](https://web.dev/pwa/)

### Eines
- [PWA Asset Generator](https://www.pwabuilder.com/imageGenerator)
- [Manifest Validator](https://manifest-validator.appspot.com/)
- [Lighthouse](https://developers.google.com/web/tools/lighthouse)
- [React PWA](https://www.pwabuilder.com/)

---

## 13. Futur: Multiplataforma Nadiua

### Què necessita

Per evolucionar a apps nadius (iOS/Android):

**Backend:**
- ✅ API compartida (Go + PocketBase) - JA FET
- ✅ Auth compartida (JWT) - JA FET

**Frontend (Nou):**
- React Native amb **Expo** (recomanat)
- Navegació compartida entre PWA i app nadiua
- Autenticació compartida
- Sincronització de dades

**Per què NO ara:**
- Duplicació innecessària de treball
- Cost de desenvolupament molt més alt
- El projecte ja és complex (multiplayer realtime)
- El PWA web ja proveeix experiència excel·lent

**Quan fer-ho:**
1. Quan el PWA web estigui madura (6-12 mesos)
2. Quan tengui una base d'usuaris sòlida
3. Quan tingui resources per 2 desenvolupaments paral·lels
4. Quan el producte tingui demandes de versió nadiu

**Plataformes a considerar:**
- **Expo + React Native** (iOS + Android) - Recomanat
- **Kotlin Multiplataforma** - Alternativa avançada
- **Ionic** - Menys codi natiu, però més ràpid

---

**Última actualització:** 20/01/2025
**Versió:** 1.0
**Enfoc:** PWA Web Reactiva (No multiplataforma nadiua actualment)
