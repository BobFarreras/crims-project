"use client";
import { useRouter } from 'next/navigation'
import LoginForm from './LoginForm'
// Uses real API endpoints now (Next.js API routes) for login
import { useState } from 'react'

export default function LoginFlow() {
  const router = useRouter()
  const [error, setError] = useState<string | null>(null)

  const onSubmit = async (payload: { username: string; password: string }) => {
    try {
      // Call the real login API (mock backend in this dev setup)
      const res = await fetch('/api/auth/login', {
        method: 'POST',
        headers: { 'Content-Type': 'application/json' },
        body: JSON.stringify({ username: payload.username, password: payload.password }),
      })
      if (!res.ok) {
        throw new Error('Login failed')
      }
      const data = await res.json()
      if (data.ok || data.token) {
        // On success, navigate to dashboard. The token is stored as HttpOnly cookie by the backend.
        router.push('/game/dashboard')
      } else {
        throw new Error('Login failed')
      }
    } catch {
      setError('Login failed')
    }
  }

  return (
    <div>
      {error ? <p className="text-sm font-semibold text-red-600" aria-live="polite">{error}</p> : null}
      <LoginForm onSubmit={onSubmit} />
    </div>
  )
}
