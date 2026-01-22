"use client";
import { useState } from 'react'

type Props = {
  onSubmit: (payload: { username: string; password: string }) => void
}

export default function RegisterForm({ onSubmit }: Props) {
  const [username, setUsername] = useState('')
  const [password, setPassword] = useState('')
  const [error, setError] = useState('')

  const handleSubmit = async (e: React.FormEvent<HTMLFormElement>) => {
    e.preventDefault()
    if (!username || !password) {
      setError('Username and password are required')
      return
    }
    setError('')
    onSubmit({ username, password })
  }

  return (
    <form className="flex w-full max-w-md flex-col gap-3" onSubmit={handleSubmit}>
      <label className="text-sm font-semibold" htmlFor="reg-username">Username</label>
      <input id="reg-username" value={username} onChange={(e) => setUsername(e.target.value)} className="h-10 rounded border px-3" />
      <label className="text-sm font-semibold" htmlFor="reg-password">Password</label>
      <input id="reg-password" type="password" value={password} onChange={(e) => setPassword(e.target.value)} className="h-10 rounded border px-3" />
      {error ? <span className="text-sm text-red-600">{error}</span> : null}
      <button className="mt-2 h-10 rounded bg-emerald-500 text-white" type="submit">Register</button>
    </form>
  )
}
