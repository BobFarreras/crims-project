import { useState } from 'react';
import { useRouter } from 'next/navigation';
import { authService } from '../services/auth.service';

export function useLogin() {
  const router = useRouter();
  const [isLoading, setIsLoading] = useState(false);
  const [error, setError] = useState<string | null>(null);

  const login = async (username: string, password: string) => {
    // 🛡️ Protecció contra 'undefined'
    if (!username || !password) {
      console.error("❌ Error: Dades incompletes al useLogin", { username, password });
      setError("Error intern: Falten dades. Refresca la pàgina.");
      return;
    }

    setIsLoading(true);
    setError(null);
    
    try {
      // Debug per veure que arriba bé
      console.log("🚀 Fent login amb:", { user: username, passLength: password.length });
      
      await authService.login(username, password);
      
      router.push('/game/dashboard');
    } catch (err) {
      console.error("❌ Error al login:", err);
      // Missatge genèric de seguretat
      setError('Credencials incorrectes.');
    } finally {
      setIsLoading(false);
    }
  };

  return { login, isLoading, error };
}