export type LobbyCreateResponse = {
  game: {
    id: string;
    code: string;
    state: string;
    seed: string;
  };
  player: {
    id: string;
    gameId: string;
    userId: string;
    capabilities: string[];
    status: string;
    isHost: boolean;
  };
};

export type LobbyJoinResponse = {
  id: string;
  gameId: string;
  userId: string;
  capabilities: string[];
  status: string;
  isHost: boolean;
};

export type GameByCodeResponse = {
  id: string;
  code: string;
  state: string;
  seed: string;
};

const API_URL = process.env.NEXT_PUBLIC_API_URL || 'http://localhost:8080';

export const lobbyService = {
  async createLobby(userId: string, capabilities: string[]): Promise<LobbyCreateResponse> {
    const res = await fetch(`${API_URL}/api/lobby/create`, {
      method: 'POST',
      headers: { 'Content-Type': 'application/json' },
      credentials: 'include',
      body: JSON.stringify({ userId, capabilities }),
    });

    if (!res.ok) {
      throw new Error('No hem pogut crear la sala.');
    }

    return await res.json();
  },

  async joinLobby(gameCode: string, userId: string, capabilities: string[]): Promise<LobbyJoinResponse> {
    const res = await fetch(`${API_URL}/api/lobby/join`, {
      method: 'POST',
      headers: { 'Content-Type': 'application/json' },
      credentials: 'include',
      body: JSON.stringify({ gameCode, userId, capabilities }),
    });

    if (!res.ok) {
      throw new Error('No hem pogut unir-nos a la sala.');
    }

    return await res.json();
  },

  async getGameByCode(code: string): Promise<GameByCodeResponse> {
    const res = await fetch(`${API_URL}/api/games/by-code/${code}`, {
      method: 'GET',
      credentials: 'include',
    });

    if (!res.ok) {
      throw new Error('No hem trobat aquesta sala.');
    }

    return await res.json();
  },
};
