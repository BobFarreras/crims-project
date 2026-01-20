"use client";

import * as Sentry from "@sentry/nextjs";
import { useState, useEffect, useRef } from "react";

export default function TestSentryPage() {
  const [logs, setLogs] = useState<string[]>([]);
  const [dsnConfigured, setDsnConfigured] = useState(false);
  const [envConfigured, setEnvConfigured] = useState(false);
  const hasInitialized = useRef(false);

  const addLog = (message: string) => {
    const timestamp = new Date().toLocaleTimeString();
    setLogs(prev => [`[${timestamp}] ${message}`, ...prev]);
  };

  const checkConfig = () => {
    // Evitar doble execució
    if (hasInitialized.current) return;

    addLog("=== CHECKING CONFIGURATION ===");

    const dsn = process.env.NEXT_PUBLIC_SENTRY_DSN;
    if (dsn && dsn.includes("sentry.io")) {
      addLog(`✅ DSN Configurat: ${dsn.substring(0, 20)}...`);
      setDsnConfigured(true);
    } else {
      addLog(`❌ DSN NO configurat o invàlid`);
      setDsnConfigured(false);
    }

    const env = process.env.NEXT_PUBLIC_ENVIRONMENT;
    if (env) {
      addLog(`✅ Environment: ${env}`);
      setEnvConfigured(true);
    } else {
      addLog(`❌ Environment NO configurat`);
      setEnvConfigured(false);
    }

    addLog(`=== END CHECKING ===`);
    hasInitialized.current = true;
  };

  const testError1 = () => {
    if (!dsnConfigured) {
      alert("❌ DSN NO configurat! Revisa la configuració.");
      return;
    }

    addLog("Test 1: Error manual (try/catch)");

    try {
      throw new Error("Test Error 1: Error de JavaScript manual");
    } catch (error) {
      Sentry.captureException(error as Error);
      addLog("✅ Error 1: Capturat a Sentry");
    }
  };

  const testError2 = () => {
    if (!dsnConfigured) {
      alert("❌ DSN NO configurat! Revisa la configuració.");
      return;
    }

    addLog("Test 2: Error asíncrono");

    Promise.reject(new Error("Test Error 2: Promise reject"))
      .catch((error) => {
        Sentry.captureException(error as Error);
        addLog("✅ Error 2: Capturat a Sentry");
      });
  };

  const testError3 = () => {
    if (!dsnConfigured) {
      alert("❌ DSN NO configurat! Revisa la configuració.");
      return;
    }

    addLog("Test 3: Error con context y tags");

    Sentry.setContext("test_context", {
      test_type: "manual_trigger",
      user_action: "click_test_button",
    });

    Sentry.setTag("test_type", "manual_error");

    const error = new Error("Test Error 3: Error amb context");
    Sentry.captureException(error);
    addLog("✅ Error 3: Capturat a Sentry");
  };

  const testMessage = () => {
    if (!dsnConfigured) {
      alert("❌ DSN NO configurat! Revisa la configuració.");
      return;
    }

    addLog("Test 4: Capturar mensaje");

    Sentry.captureMessage("Test Message: Usuari ha fet alguna cosa");
    addLog("✅ Mensaje 4: Capturat a Sentry");
  };

  const testBreadcrumbs = () => {
    if (!dsnConfigured) {
      alert("❌ DSN NO configurat! Revisa la configuració.");
      return;
    }

    addLog("Test 5: Error con breadcrumbs");

    Sentry.addBreadcrumb({
      message: "Usuari ha fet clic al test",
      category: "user",
      level: "info",
    });
    addLog("✅ Breadcrumb 5: Afegit");

    setTimeout(() => {
      try {
        throw new Error("Test Error 5: Error després de breadcrumb");
      } catch (error) {
        Sentry.captureException(error as Error);
        addLog("✅ Error 5: Capturat a Sentry");
      }
    }, 100);
  };

  const testNetwork = async () => {
    if (!dsnConfigured) {
      alert("❌ DSN NO configurat! Revisa la configuració.");
      return;
    }

    addLog("Test 6: Error de xarxa (fetch)");

    try {
      const response = await fetch("http://localhost:9999/api/test", {
        method: "POST",
      });

      if (!response.ok) {
        throw new Error(`HTTP ${response.status}`);
      }
    } catch (error) {
      Sentry.captureException(error as Error);
      addLog(`✅ Error 6: Capturat a Sentry (${(error as Error).message})`);
    }
  };

  const clearLogs = () => {
    setLogs([]);
  };

  useEffect(() => {
    checkConfig();
  }, []);

  return (
    <div className="min-h-screen bg-gray-100 p-8">
      <div className="max-w-4xl mx-auto">
        <h1 className="text-3xl font-bold mb-6">🧪 Test de Sentry - Frontend (v3)</h1>

        <div className="bg-white rounded-lg shadow-md p-6 mb-6">
          <h2 className="text-xl font-semibold mb-4">Estat de Configuración</h2>

          <div className="grid grid-cols-2 gap-4">
            <div className={`p-4 rounded ${dsnConfigured ? 'bg-green-50 border-l-4 border-green-500' : 'bg-red-50 border-l-4 border-red-500'}`}>
              <p className="font-semibold">
                {dsnConfigured ? '✅ DSN Configurat' : '❌ DSN NO Configurat'}
              </p>
              {dsnConfigured && (
                <p className="text-sm text-gray-600 mt-1">
                  {process.env.NEXT_PUBLIC_SENTRY_DSN?.substring(0, 30)}...
                </p>
              )}
            </div>

            <div className={`p-4 rounded ${envConfigured ? 'bg-green-50 border-l-4 border-green-500' : 'bg-red-50 border-l-4 border-red-500'}`}>
              <p className="font-semibold">
                {envConfigured ? `✅ Environment: ${process.env.NEXT_PUBLIC_ENVIRONMENT}` : '❌ Environment NO Configurat'}
              </p>
            </div>
          </div>
        </div>

        <div className="bg-white rounded-lg shadow-md p-6 mb-6">
          <div className="flex justify-between items-center mb-4">
            <h2 className="text-xl font-semibold">Logs de Ejecución</h2>
            <button
              onClick={clearLogs}
              className="text-sm bg-gray-200 hover:bg-gray-300 px-3 py-1 rounded"
            >
              Netejar Logs
            </button>
          </div>

          <div className="bg-gray-900 text-green-400 rounded p-4 font-mono text-sm h-48 overflow-y-auto">
            {logs.length === 0 ? (
              <p className="text-gray-500">Cap log encara...</p>
            ) : (
              logs.map((log, index) => (
                <div key={index} className="mb-1">
                  {log}
                </div>
              ))
            )}
          </div>
        </div>

        <div className="bg-white rounded-lg shadow-md p-6 mb-6">
          <h2 className="text-xl font-semibold mb-4">Tests d'Errors</h2>

          <div className="grid grid-cols-2 gap-4">
            <button
              onClick={testError1}
              disabled={!dsnConfigured}
              className="bg-red-500 hover:bg-red-600 disabled:opacity-50 disabled:cursor-not-allowed text-white font-bold py-3 px-4 rounded transition duration-300"
            >
              1. Error Manual (try/catch)
            </button>

            <button
              onClick={testError2}
              disabled={!dsnConfigured}
              className="bg-orange-500 hover:bg-orange-600 disabled:opacity-50 disabled:cursor-not-allowed text-white font-bold py-3 px-4 rounded transition duration-300"
            >
              2. Error Asíncrono (Promise reject)
            </button>

            <button
              onClick={testError3}
              disabled={!dsnConfigured}
              className="bg-yellow-500 hover:bg-yellow-600 disabled:opacity-50 disabled:cursor-not-allowed text-white font-bold py-3 px-4 rounded transition duration-300"
            >
              3. Error amb Context i Tags
            </button>

            <button
              onClick={testMessage}
              disabled={!dsnConfigured}
              className="bg-blue-500 hover:bg-blue-600 disabled:opacity-50 disabled:cursor-not-allowed text-white font-bold py-3 px-4 rounded transition duration-300"
            >
              4. Capturar Missatge
            </button>

            <button
              onClick={testBreadcrumbs}
              disabled={!dsnConfigured}
              className="bg-purple-500 hover:bg-purple-600 disabled:opacity-50 disabled:cursor-not-allowed text-white font-bold py-3 px-4 rounded transition duration-300"
            >
              5. Error amb Breadcrumbs
            </button>

            <button
              onClick={testNetwork}
              disabled={!dsnConfigured}
              className="bg-pink-500 hover:bg-pink-600 disabled:opacity-50 disabled:cursor-not-allowed text-white font-bold py-3 px-4 rounded transition duration-300"
            >
              6. Error de Xarxa (fetch)
            </button>
          </div>
        </div>

        <div className="bg-white rounded-lg shadow-md p-6">
          <h2 className="text-xl font-semibold mb-4">Com Verificar</h2>

          <ol className="list-decimal list-inside space-y-3 text-gray-700">
            <li><strong>Pas 1:</strong> Revisa "Estat de Configuración" - hauria de ser tot ✅</li>
            <li><strong>Pas 2:</strong> Fes clic en un dels botons de test</li>
            <li><strong>Pas 3:</strong> Mira els logs de sota (haurien de dir "Capturat a Sentry")</li>
            <li><strong>Pas 4:</strong> Ves al dashboard de Sentry: <a href="https://sentry.io" target="_blank" rel="noreferrer" className="text-blue-500 hover:underline">https://sentry.io</a></li>
            <li><strong>Pas 5:</strong> Selecciona l'organització: <strong>digitaistudios</strong></li>
            <li><strong>Pas 6:</strong> Selecciona el projecte: <strong>crims-frontend</strong></li>
          </ol>
        </div>
      </div>
    </div>
  );
}


    // Check Environment
    const env = process.env.NEXT_PUBLIC_ENVIRONMENT;
    if (env) {
      addLog(`✅ Environment: ${env}`);
      setEnvConfigured(true);
    } else {
      addLog(`❌ Environment NO configurat`);
      setEnvConfigured(false);
    }

    addLog(`=== END CHECKING ===`);
  };

  const testError1 = () => {
    if (!dsnConfigured) {
      alert("❌ DSN NO configurat! Revisa la configuració.");
      return;
    }

    addLog("Test 1: Error manual (try/catch)");

    try {
      throw new Error("Test Error 1: Error de JavaScript manual");
    } catch (error) {
      Sentry.captureException(error as Error);
      addLog("✅ Error 1: Capturat a Sentry");
    }
  };

  const testError2 = () => {
    if (!dsnConfigured) {
      alert("❌ DSN NO configurat! Revisa la configuració.");
      return;
    }

    addLog("Test 2: Error asíncrono");

    Promise.reject(new Error("Test Error 2: Promise reject"))
      .then(() => {
        addLog("❌ Promise no hauria de resolt");
      })
      .catch((error) => {
        Sentry.captureException(error as Error);
        addLog("✅ Error 2: Capturat a Sentry");
      });
  };

  const testError3 = () => {
    if (!dsnConfigured) {
      alert("❌ DSN NO configurat! Revisa la configuració.");
      return;
    }

    addLog("Test 3: Error con context y tags");

    Sentry.setContext("test_context", {
      test_type: "manual_trigger",
      user_action: "click_test_button",
      timestamp: Date.now(),
    });

    Sentry.setTag("test_type", "manual_error");

    const error = new Error("Test Error 3: Error amb context");
    Sentry.captureException(error);
    addLog("✅ Error 3: Capturat a Sentry");
  };

  const testMessage = () => {
    if (!dsnConfigured) {
      alert("❌ DSN NO configurat! Revisa la configuració.");
      return;
    }

    addLog("Test 4: Capturar mensaje");

    Sentry.captureMessage("Test Message: Usuari ha fet alguna cosa", "warning");
    addLog("✅ Mensaje 4: Capturat a Sentry");
  };

  const testBreadcrumbs = () => {
    if (!dsnConfigured) {
      alert("❌ DSN NO configurat! Revisa la configuració.");
      return;
    }

    addLog("Test 5: Error con breadcrumbs");

    Sentry.addBreadcrumb({
      message: "Usuari ha fet clic al test",
      category: "user",
      level: "info",
    });
    addLog("✅ Breadcrumb 5: Afegit");

    setTimeout(() => {
      try {
        throw new Error("Test Error 5: Error després de breadcrumb");
      } catch (error) {
        Sentry.captureException(error as Error);
        addLog("✅ Error 5: Capturat a Sentry");
      }
    }, 100);
  };

  const testNetwork = async () => {
    if (!dsnConfigured) {
      alert("❌ DSN NO configurat! Revisa la configuració.");
      return;
    }

    addLog("Test 6: Error de xarxa (fetch)");

    try {
      // Intentar fer fetch a un endpoint que no existeix
      const response = await fetch("http://localhost:9999/api/test", {
        method: "POST",
      });

      if (!response.ok) {
        throw new Error(`HTTP ${response.status}`);
      }
    } catch (error) {
      Sentry.captureException(error as Error);
      addLog(`✅ Error 6: Capturat a Sentry (${(error as Error).message})`);
    }
  };

  const clearLogs = () => {
    setLogs([]);
  };

  // Check configuration on mount
  useState(() => {
    checkConfig();
  });

  return (
    <div className="min-h-screen bg-gray-100 p-8">
      <div className="max-w-4xl mx-auto">
        <h1 className="text-3xl font-bold mb-6">🧪 Test de Sentry - Frontend (v2)</h1>

        {/* Configuration Status */}
        <div className="bg-white rounded-lg shadow-md p-6 mb-6">
          <h2 className="text-xl font-semibold mb-4">Estat de Configuración</h2>

          <div className="grid grid-cols-2 gap-4">
            <div className={`p-4 rounded ${dsnConfigured ? 'bg-green-50 border-l-4 border-green-500' : 'bg-red-50 border-l-4 border-red-500'}`}>
              <p className="font-semibold">
                {dsnConfigured ? '✅ DSN Configurat' : '❌ DSN NO Configurat'}
              </p>
              {dsnConfigured && (
                <p className="text-sm text-gray-600 mt-1">
                  {process.env.NEXT_PUBLIC_SENTRY_DSN?.substring(0, 30)}...
                </p>
              )}
            </div>

            <div className={`p-4 rounded ${envConfigured ? 'bg-green-50 border-l-4 border-green-500' : 'bg-red-50 border-l-4 border-red-500'}`}>
              <p className="font-semibold">
                {envConfigured ? `✅ Environment: ${process.env.NEXT_PUBLIC_ENVIRONMENT}` : '❌ Environment NO Configurat'}
              </p>
            </div>
          </div>
        </div>

        {/* Logs */}
        <div className="bg-white rounded-lg shadow-md p-6 mb-6">
          <div className="flex justify-between items-center mb-4">
            <h2 className="text-xl font-semibold">Logs de Execució</h2>
            <button
              onClick={clearLogs}
              className="text-sm bg-gray-200 hover:bg-gray-300 px-3 py-1 rounded"
            >
              Netejar Logs
            </button>
          </div>

          <div className="bg-gray-900 text-green-400 rounded p-4 font-mono text-sm h-48 overflow-y-auto">
            {logs.length === 0 ? (
              <p className="text-gray-500">Cap log encara...</p>
            ) : (
              logs.map((log, index) => (
                <div key={index} className="mb-1">
                  {log}
                </div>
              ))
            )}
          </div>
        </div>

        {/* Tests */}
        <div className="bg-white rounded-lg shadow-md p-6 mb-6">
          <h2 className="text-xl font-semibold mb-4">Tests d'Errors</h2>

          <div className="grid grid-cols-2 gap-4">
            <button
              onClick={testError1}
              disabled={!dsnConfigured}
              className="bg-red-500 hover:bg-red-600 disabled:opacity-50 disabled:cursor-not-allowed text-white font-bold py-3 px-4 rounded transition duration-300"
            >
              1. Error Manual (try/catch)
            </button>

            <button
              onClick={testError2}
              disabled={!dsnConfigured}
              className="bg-orange-500 hover:bg-orange-600 disabled:opacity-50 disabled:cursor-not-allowed text-white font-bold py-3 px-4 rounded transition duration-300"
            >
              2. Error Asíncrono (Promise reject)
            </button>

            <button
              onClick={testError3}
              disabled={!dsnConfigured}
              className="bg-yellow-500 hover:bg-yellow-600 disabled:opacity-50 disabled:cursor-not-allowed text-white font-bold py-3 px-4 rounded transition duration-300"
            >
              3. Error amb Context i Tags
            </button>

            <button
              onClick={testMessage}
              disabled={!dsnConfigured}
              className="bg-blue-500 hover:bg-blue-600 disabled:opacity-50 disabled:cursor-not-allowed text-white font-bold py-3 px-4 rounded transition duration-300"
            >
              4. Capturar Missatge
            </button>

            <button
              onClick={testBreadcrumbs}
              disabled={!dsnConfigured}
              className="bg-purple-500 hover:bg-purple-600 disabled:opacity-50 disabled:cursor-not-allowed text-white font-bold py-3 px-4 rounded transition duration-300"
            >
              5. Error amb Breadcrumbs
            </button>

            <button
              onClick={testNetwork}
              disabled={!dsnConfigured}
              className="bg-pink-500 hover:bg-pink-600 disabled:opacity-50 disabled:cursor-not-allowed text-white font-bold py-3 px-4 rounded transition duration-300"
            >
              6. Error de Xarxa (fetch)
            </button>
          </div>
        </div>

        {/* Instructions */}
        <div className="bg-white rounded-lg shadow-md p-6">
          <h2 className="text-xl font-semibold mb-4">Com Verificar</h2>

          <ol className="list-decimal list-inside space-y-3 text-gray-700">
            <li><strong>Pas 1:</strong> Fes clic en "Netejar Logs" per esborrar logs anteriors</li>
            <li><strong>Pas 2:</strong> Revisa "Estat de Configuración" - hauria de ser tot ✅</li>
            <li><strong>Pas 3:</strong> Fes clic en un dels botons de test</li>
            <li><strong>Pas 4:</strong> Mira els logs de sota (haurien de dir "Capturat a Sentry")</li>
            <li><strong>Pas 5:</strong> Ves al dashboard de Sentry: <a href="https://sentry.io" target="_blank" rel="noreferrer" className="text-blue-500 hover:underline">https://sentry.io</a></li>
            <li><strong>Pas 6:</strong> Selecciona l'organització: <strong>digitaistudios</strong></li>
            <li><strong>Pas 7:</strong> Selecciona el projecte: <strong>crims-frontend</strong></li>
            <li><strong>Pas 8:</strong> Hauries de veure l'error que has generat</li>
          </ol>

          <div className="mt-6 bg-yellow-50 border-l-4 border-yellow-500 p-4">
            <p className="font-semibold text-yellow-800">⚠️  Si NO veus errors a Sentry:</p>
            <ul className="mt-2 text-sm text-yellow-700 list-disc list-inside">
              <li>Revisa que el DSN és correcte a .env.local</li>
              <li>Revisa que .env.local està a l'ARREL del projecte</li>
              <li>Revisa que tens connexió a internet</li>
              <li>Revisa el navegador (F12 → Console) per veure si hi ha errors</li>
            </ul>
          </div>
        </div>
      </div>
    </div>
  );
}
