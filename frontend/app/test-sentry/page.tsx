"use client";

import * as Sentry from "@sentry/nextjs";

export default function TestSentryPage() {
  const testError1 = () => {
    try {
      // Error de JavaScript típic
      throw new Error("Test Error 1: Error de JavaScript manual");
    } catch (error) {
      Sentry.captureException(error);
      alert("Error capturat! Mira Sentry Dashboard");
    }
  };

  const testError2 = () => {
    // Error asíncrono
    Promise.reject(new Error("Test Error 2: Promise reject"))
      .catch((error) => {
        Sentry.captureException(error);
        alert("Error asíncron capturat! Mira Sentry Dashboard");
      });
  };

  const testError3 = () => {
    // Error amb context
    Sentry.setContext("test_context", {
      test_type: "manual_trigger",
      user_action: "click_test_button",
      timestamp: Date.now(),
    });

    Sentry.setTag("test_type", "manual_error");

    Sentry.captureException(new Error("Test Error 3: Error amb context"));
    alert("Error amb context capturat! Mira Sentry Dashboard");
  };

  const testMessage = () => {
    // Capturar un missatge (no error)
    Sentry.captureMessage("Test Message: Usuari ha fet alguna cosa", "warning");
    alert("Missatge capturat! Mira Sentry Dashboard");
  };

  const testBreadcrumbs = () => {
    // Crear breadcrumbs (rastre d'accions)
    Sentry.addBreadcrumb({
      message: "Usuari ha fet clic al test",
      category: "user",
      level: "info",
    });

    setTimeout(() => {
      throw new Error("Test Error 4: Error després de breadcrumb");
    }, 100);
  };

  return (
    <div className="min-h-screen bg-gray-100 p-8">
      <div className="max-w-2xl mx-auto">
        <h1 className="text-3xl font-bold mb-6">🧪 Test de Sentry - Frontend</h1>

        <div className="bg-white rounded-lg shadow-md p-6 mb-6">
          <h2 className="text-xl font-semibold mb-4">Informació</h2>
          <div className="bg-blue-50 border-l-4 border-blue-500 p-4">
            <p className="text-sm">
              <strong>Environment:</strong> {process.env.NEXT_PUBLIC_ENVIRONMENT || "No configurat"}<br />
              <strong>Sentry DSN:</strong> {process.env.NEXT_PUBLIC_SENTRY_DSN ? "Configurat ✅" : "No configurat ❌"}
            </p>
          </div>
        </div>

        <div className="bg-white rounded-lg shadow-md p-6 mb-6">
          <h2 className="text-xl font-semibold mb-4">Tests d'Errors</h2>

          <div className="space-y-4">
            <button
              onClick={testError1}
              className="w-full bg-red-500 hover:bg-red-600 text-white font-bold py-3 px-4 rounded transition duration-300"
            >
              1. Error Manual (try/catch)
            </button>

            <button
              onClick={testError2}
              className="w-full bg-orange-500 hover:bg-orange-600 text-white font-bold py-3 px-4 rounded transition duration-300"
            >
              2. Error Asíncrono (Promise reject)
            </button>

            <button
              onClick={testError3}
              className="w-full bg-yellow-500 hover:bg-yellow-600 text-white font-bold py-3 px-4 rounded transition duration-300"
            >
              3. Error amb Context i Tags
            </button>

            <button
              onClick={testMessage}
              className="w-full bg-blue-500 hover:bg-blue-600 text-white font-bold py-3 px-4 rounded transition duration-300"
            >
              4. Capturar Missatge (no error)
            </button>

            <button
              onClick={testBreadcrumbs}
              className="w-full bg-purple-500 hover:bg-purple-600 text-white font-bold py-3 px-4 rounded transition duration-300"
            >
              5. Error amb Breadcrumbs
            </button>
          </div>
        </div>

        <div className="bg-white rounded-lg shadow-md p-6">
          <h2 className="text-xl font-semibold mb-4">Com Verificar</h2>
          <ol className="list-decimal list-inside space-y-2 text-gray-700">
            <li>Fes clic en un dels botons de dalt</li>
            <li>Ves al dashboard de Sentry: <a href="https://sentry.io" target="_blank" rel="noreferrer" className="text-blue-500 hover:underline">https://sentry.io</a></li>
            <li>Selecciona l'organització: <strong>digitaistudios</strong></li>
            <li>Selecciona el projecte: <strong>crims-frontend</strong></li>
            <li>Hauries de veure l'error que has generat</li>
          </ol>
        </div>
      </div>
    </div>
  );
}
