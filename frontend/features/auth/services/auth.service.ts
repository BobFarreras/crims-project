// frontend/features/auth/services/auth.service.ts

// Tipus per la resposta (pots moure-ho a types/auth.types.ts)
export type LoginResponse = {
  token: string;
  user: {
    id: string;
    username: string;
    role: string;
  };
};

const API_URL = process.env.NEXT_PUBLIC_API_URL || '';

export const authService = {
  /**
   * Realitza la petició de login al backend
   */
  async login(username: string, password: string): Promise<LoginResponse> {
    const res = await fetch(`${API_URL}/api/auth/login`, {
      method: 'POST',
      headers: { 
        'Content-Type': 'application/json' 
      },
      body: JSON.stringify({ username, password }),
    });

    if (!res.ok) {
      // Aquí podries gestionar errors 401, 500, etc.
      throw new Error('Credencials incorrectes o error del servidor');
    }

    return await res.json();
  }
};