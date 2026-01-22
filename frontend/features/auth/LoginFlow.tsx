"use client";
import { useRouter } from 'next/navigation'
import LoginForm from './LoginForm'
import { login as mockLogin } from '@/lib/infra/auth-client'
import { useState } from 'react'

export default function LoginFlow() {
  const router = useRouter()
  const [error, setError] = useState<string | null>(null)

  const onSubmit = async (payload: { username: string; password: string }) => {
    try {
      // Call authentication client (mocked for now)
      const res = await mockLogin(payload.username, payload.password)
      // Persist token if needed (localStorage example)
      if (typeof window !== 'undefined' && res?.token) {
        localStorage.setItem('token', res.token)
      }
      // Redirect to game dashboard after login
      router.push('/game/dashboard')
    } catch (e) {
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
