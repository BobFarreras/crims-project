"use client";
import { useState } from 'react'

type Props = {
  onSubmit: (payload: { username: string; password: string }) => void
}

export default function LoginForm({ onSubmit }: Props) {
  const [username, setUsername] = useState('')
  const [password, setPassword] = useState('')

  const handleSubmit = (e: React.FormEvent<HTMLFormElement>) => {
    e.preventDefault()
    onSubmit({ username, password })
  }

  return (
    <form onSubmit={handleSubmit} className="flex w-full max-w-md flex-col gap-3">
      <label className="text-sm font-semibold" htmlFor="username">Username</label>
      <input
        id="username"
        value={username}
        onChange={(e) => setUsername(e.target.value)}
        className="h-10 rounded border px-3"
      />
      <label className="text-sm font-semibold" htmlFor="password">Password</label>
      <input
        id="password"
        type="password"
        value={password}
        onChange={(e) => setPassword(e.target.value)}
        className="h-10 rounded border px-3"
      />
      <button type="submit" className="h-10 rounded bg-amber-500 text-white font-semibold">Login</button>
    </form>
  )
}
