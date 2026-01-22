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

const API_URL = process.env.NEXT_PUBLIC_API_URL || 'http://localhost:8080';

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
  },

  /**
   * Registra un nou usuari a PocketBase
   */
  async register(payload: { username: string; email: string; password: string; passwordConfirm: string }) {

    // 2. IMPORTANT: Fem servir la variable API_URL al principi
    // Fixa't que no hi ha barra '/' entre API_URL i /api perquè normalment l'URL base no en porta
    const res = await fetch(`${API_URL}/api/auth/register`, {
      method: 'POST',
      headers: { 'Content-Type': 'application/json' },
      body: JSON.stringify(payload),
    });

    if (!res.ok) {
      const errorData = await res.json().catch(() => ({}));
      throw new Error(errorData.error || 'Error en el registre');
    }

    return await res.json();
  }



};