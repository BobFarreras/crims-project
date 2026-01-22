import { useState } from 'react';
import { useRouter } from 'next/navigation';
import { authService } from '../services/auth.service';

export function useLogin() {
  const router = useRouter();
  const [isLoading, setIsLoading] = useState(false);
  const [error, setError] = useState<string | null>(null);

  const login = async (username: string, password: string) => {
    setIsLoading(true);
    setError(null);
    
    try {
      // 1. Cridem al servei (lògica backend)
      await authService.login(username, password);
      
      // 2. Si tot va bé, redirigim
      router.push('/game/dashboard');
    } catch (err) {
      // 3. Gestionem errors
      setError(err instanceof Error ? err.message : 'Error en iniciar sessió');
    } finally {
      setIsLoading(false);
    }
  };

  return { login, isLoading, error };
}