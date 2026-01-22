"use client";
import { useRouter } from 'next/navigation'
import { useState } from 'react'
import RegisterForm from './RegisterForm'

export default function RegisterFlow() {
  const router = useRouter()
  const [error, setError] = useState<string | null>(null)

  const onSubmit = async (payload: { username: string; password: string }) => {
    try {
      const res = await fetch('/api/auth/register', {
        method: 'POST',
        headers: { 'Content-Type': 'application/json' },
        body: JSON.stringify(payload),
      })
      if (!res.ok) throw new Error('Register failed')
      // After successful registration, navigate to login
      router.push('/login')
  } catch {
      setError('Register failed')
    }
  }

  return (
    <div>
      {error ? <p className="text-sm font-semibold text-red-600" aria-live="polite">{error}</p> : null}
      <RegisterForm onSubmit={onSubmit} />
    </div>
  )
}
