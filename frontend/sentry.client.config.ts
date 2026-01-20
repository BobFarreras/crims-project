// Sentry Client Configuration (Frontend - Browser)
import * as Sentry from "@sentry/nextjs";

Sentry.init({
  dsn: process.env.NEXT_PUBLIC_SENTRY_DSN,

  // Environment
  environment: process.env.NEXT_PUBLIC_ENVIRONMENT || "development",

  // Sample Rate (0.0 = 100%, 1.0 = 0%)
  tracesSampleRate: 0.1, // 10% de traces
  replaysSessionSampleRate: 0.1, // 10% de sessions
  replaysOnErrorSampleRate: 1.0, // 100% de sessions amb error

  // Integrations
  integrations: [
    Sentry.replayIntegration({
      // Opcions de replay
      maskAllText: false,
      blockAllMedia: false,
    }),
    Sentry.browserTracingIntegration(),
  ],

  // Before Send (filtre errors)
  beforeSend(event, hint) {
    // No enviar errors en development (opcional)
    if (process.env.NODE_ENV === "development") {
      console.error("Sentry Error:", event);
      return null;
    }

    // Filtrar errors específics
    if (event.exception) {
      const error = hint.originalException;
      if (error instanceof Error) {
        // No enviar errors de network (ja tenim altres solucions)
        if (error.message.includes("Network Error")) {
          return null;
        }
      }
    }

    return event;
  },

  // Filtre d'errors
  ignoreErrors: [
    "Non-Error promise rejection captured",
    "ResizeObserver loop limit exceeded",
    "ChunkLoadError", // Errors de Webpack en development
  ],

  // Tags per identificar l'aplicació
  initialScope: {
    tags: {
      app: "crims-frontend",
      framework: "nextjs",
    },
  },
});
