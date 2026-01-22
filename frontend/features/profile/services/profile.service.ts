export type ProfileResponse = {
  user: {
    id: string;
    username: string;
    name: string;
  };
};

const API_URL = process.env.NEXT_PUBLIC_API_URL || 'http://localhost:8080';

export const profileService = {
  async getProfile(): Promise<ProfileResponse> {
    const res = await fetch(`${API_URL}/api/profile`, {
      method: 'GET',
      credentials: 'include',
    });

    if (!res.ok) {
      throw new Error('No hem pogut carregar el perfil.');
    }

    return await res.json();
  },

  async updateProfile(name: string): Promise<ProfileResponse> {
    const res = await fetch(`${API_URL}/api/profile`, {
      method: 'PUT',
      headers: {
        'Content-Type': 'application/json',
      },
      credentials: 'include',
      body: JSON.stringify({ name }),
    });

    if (!res.ok) {
      if (res.status === 400) {
        throw new Error('El nom no pot estar buit.');
      }
      throw new Error('No hem pogut guardar el perfil.');
    }

    return await res.json();
  },
};
