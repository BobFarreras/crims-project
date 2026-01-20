// Sentry Edge Configuration (Edge Functions, Middleware)
import * as Sentry from "@sentry/nextjs";

Sentry.init({
  dsn: process.env.SENTRY_DSN,

  // Environment
  environment: process.env.NEXT_PUBLIC_ENVIRONMENT || "development",

  // Sample Rate (més baix per Edge)
  tracesSampleRate: 0.05, // 5% de traces

  // Before Send (filtre errors)
  beforeSend(event, hint) {
    // PERMITIR ENVIAR ERRORS EN DEVELOPMENT PER FER TESTS
    // Comentat temporalment per verificar que funciona
    // if (process.env.NODE_ENV === "development") {
    //   console.error("Sentry Error (Edge):", event);
    //   return null;
    // }

    return event;
  },

  // Tags per identificar l'aplicació
  initialScope: {
    tags: {
      app: "crims-frontend-edge",
      framework: "nextjs",
      runtime: "edge",
    },
  },
});
