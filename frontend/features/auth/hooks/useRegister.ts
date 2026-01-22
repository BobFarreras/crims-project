// frontend/features/auth/hooks/useRegister.ts
import { useState } from 'react';
import { useRouter } from 'next/navigation';
import { authService } from '../services/auth.service';

export function useRegister() {
  const router = useRouter();
  const [isLoading, setIsLoading] = useState(false);
  const [error, setError] = useState<string | null>(null);

  const register = async (data: { 
    username: string; 
    email: string; 
    password: string; 
    passwordConfirm: string 
  }) => {
    setIsLoading(true);
    setError(null);
    try {
      await authService.register(data);
      // Si tot va bé, redirigim al login o directament al dashboard
      router.push('/login?registered=true');
    } catch (err) {
      setError(err instanceof Error ? err.message : 'Error desconegut');
    } finally {
      setIsLoading(false);
    }
  };

  return { register, isLoading, error };
}