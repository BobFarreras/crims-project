'use client'

import { useState } from 'react'

type LobbyFormProps = {
  roles: string[]
  onSubmit: (payload: { code: string; role: string; userId: string }) => void
}

export default function LobbyForm({ roles, onSubmit }: LobbyFormProps) {
  const [code, setCode] = useState('')
  const [role, setRole] = useState(roles[0] ?? '')
  const [userId, setUserId] = useState('')
  const [error, setError] = useState('')

  const handleSubmit = (event: React.FormEvent<HTMLFormElement>) => {
    event.preventDefault()
    if (code.trim() === '') {
      setError('Code is required')
      return
    }
    setError('')
    onSubmit({ code, role, userId })
  }

  return (
    <form className="flex w-full flex-col gap-5" onSubmit={handleSubmit}>
      <div className="flex flex-col gap-2">
        <label className="text-sm font-semibold uppercase tracking-wide text-zinc-700" htmlFor="game-code">
          Game Code
        </label>
        <input
          id="game-code"
          name="gameCode"
          value={code}
          onChange={(event) => setCode(event.target.value.toUpperCase())}
          placeholder="ABCD"
          className="h-12 rounded-xl border border-zinc-200 bg-white px-4 text-base text-zinc-900 shadow-sm focus:border-amber-400 focus:outline-none"
        />
      </div>

      <div className="flex flex-col gap-2">
        <label className="text-sm font-semibold uppercase tracking-wide text-zinc-700" htmlFor="user-id">
          User ID
        </label>
        <input
          id="user-id"
          name="userId"
          value={userId}
          onChange={(event) => setUserId(event.target.value)}
          placeholder="detective-12"
          className="h-12 rounded-xl border border-zinc-200 bg-white px-4 text-base text-zinc-900 shadow-sm focus:border-amber-400 focus:outline-none"
        />
      </div>

      <div className="flex flex-col gap-2">
        <label className="text-sm font-semibold uppercase tracking-wide text-zinc-700" htmlFor="role">
          Role
        </label>
        <select
          id="role"
          name="role"
          value={role}
          onChange={(event) => setRole(event.target.value)}
          className="h-12 rounded-xl border border-zinc-200 bg-white px-4 text-base text-zinc-900 shadow-sm focus:border-amber-400 focus:outline-none"
        >
          {roles.map((option) => (
            <option key={option} value={option}>
              {option}
            </option>
          ))}
        </select>
      </div>

      {error ? (
        <p className="text-sm font-semibold text-red-600" role="alert" aria-live="polite">
          {error}
        </p>
      ) : null}

      <button
        type="submit"
        className="h-12 rounded-xl bg-amber-500 text-base font-semibold text-white shadow-md transition hover:bg-amber-600"
      >
        Join
      </button>
    </form>
  )
}
