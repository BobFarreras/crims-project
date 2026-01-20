// Sentry Server Configuration (API Routes, Server Components)
import * as Sentry from "@sentry/nextjs";

Sentry.init({
  dsn: process.env.SENTRY_DSN,

  // Environment
  environment: process.env.ENVIRONMENT || "development",

  // Sample Rate
  tracesSampleRate: 0.1, // 10% de traces

  // Integrations
  integrations: [
    Sentry.httpIntegration({
      breadcrumbs: true,
      tracing: true,
    }),
  ],

  // Before Send (filtre errors)
  beforeSend(event, hint) {
    // No enviar errors en development (opcional)
    if (process.env.ENVIRONMENT === "development") {
      console.error("Sentry Error (Server):", event);
      return null;
    }

    return event;
  },

  // Tags per identificar l'aplicació
  initialScope: {
    tags: {
      app: "crims-frontend-server",
      framework: "nextjs",
      runtime: "node",
    },
  },
});
