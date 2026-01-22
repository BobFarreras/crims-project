import { useState } from 'react';
import { useRouter } from 'next/navigation';
import { authService } from '../services/auth.service';

export function useLogin() {
  const router = useRouter();
  const [isLoading, setIsLoading] = useState(false);
  const [error, setError] = useState<string | null>(null);

  const login = async (email: string, password: string) => {
    // 🛡️ Protecció contra 'undefined'
    if (!email || !password) {
      console.error("❌ Error: Dades incompletes al useLogin", { email, password });
      setError("Error intern: Falten dades. Refresca la pàgina.");
      return;
    }

    setIsLoading(true);
    setError(null);
    
    try {
      // Debug per veure que arriba bé
      console.log("🚀 Fent login amb:", { email, passLength: password.length });
      
      await authService.login(email, password);
      
      router.push('/game/dashboard');
    } catch (err) {
      console.error("❌ Error al login:", err);
      const message = err instanceof Error ? err.message : 'No hem pogut iniciar sessio.';
      setError(message);
    } finally {
      setIsLoading(false);
    }
  };

  return { login, isLoading, error };
}
