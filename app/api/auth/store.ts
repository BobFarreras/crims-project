type User = { username: string; password: string };

// Simple in-memory user store (for dev/prototype). Persists while server runs.
const users: Map<string, string> = new Map(); // username -> password

export function addUser(username: string, password: string): boolean {
  if (users.has(username)) return false;
  users.set(username, password);
  return true;
}

export function validateUser(username: string, password: string): boolean {
  const pass = users.get(username);
  return pass !== undefined && pass === password;
}
