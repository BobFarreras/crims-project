'use client';

// ✅ Importem des de l'índex de la feature (Bona pràctica!)
import { LoginForm, useLogin } from '@/features/auth'; 

export default function LoginPage() {
  // Utilitzem el hook de la feature
  const { login, error, isLoading } = useLogin();

  return (
    <div className="flex min-h-screen flex-col items-center justify-center bg-gray-50 p-4">
      <div className="w-full max-w-md space-y-8">
        <div className="text-center">
          <h2 className="text-3xl font-bold tracking-tight text-gray-900">
            Accés a Crims
          </h2>
        </div>

        {/* Mostrem error si n'hi ha */}
        {error && (
          <div className="rounded-md bg-red-50 p-4 text-sm text-red-700">
            {error}
          </div>
        )}

        {/* El component UI rep la funció del hook */}
        <LoginForm 
          onSubmit={({ username, password }) => login(username, password)}
          // Si el teu LoginForm suporta loading visual, passa-li aquí:
          // isLoading={isLoading} 
        />
      </div>
    </div>
  );
}