// Auth client (mock/basic)
export async function login(username: string, password: string): Promise<{ token: string; userId: string }>{
  // Mock delay to simulate network
  await new Promise((r) => setTimeout(r, 200));
  if (!username || !password) {
    throw new Error('Invalid credentials')
  }
  return { token: 'mock-token-' + username, userId: 'user-' + username }
}
