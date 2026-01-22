'use client'

import { useState } from 'react'

type LobbyFormProps = {
  capabilities: string[]
  onSubmit: (payload: { code: string; capabilities: string[]; userId: string }) => void
}

export default function LobbyForm({ capabilities, onSubmit }: LobbyFormProps) {
  const [code, setCode] = useState('')
  const [selectedCapabilities, setSelectedCapabilities] = useState<string[]>([])
  const [userId, setUserId] = useState('')
  const [error, setError] = useState('')

  const handleSubmit = (event: React.FormEvent<HTMLFormElement>) => {
    event.preventDefault()
    if (code.trim() === '') {
      setError('Code is required')
      return
    }
    setError('')
    onSubmit({ code, capabilities: selectedCapabilities, userId })
  }

  const toggleCapability = (capability: string) => {
    setSelectedCapabilities((current) =>
      current.includes(capability)
        ? current.filter((item) => item !== capability)
        : [...current, capability]
    )
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

      <fieldset className="flex flex-col gap-3" aria-label="Capabilities">
        <legend className="text-sm font-semibold uppercase tracking-wide text-zinc-700">
          Capabilities
        </legend>
        <div className="grid gap-2 sm:grid-cols-2">
          {capabilities.map((option) => (
            <label key={option} className="flex items-center gap-3 rounded-xl border border-zinc-200 bg-white px-4 py-3 text-sm font-semibold text-zinc-800 shadow-sm">
              <input
                type="checkbox"
                name="capabilities"
                value={option}
                checked={selectedCapabilities.includes(option)}
                onChange={() => toggleCapability(option)}
                className="h-4 w-4 rounded border-zinc-300 text-amber-500 focus:ring-amber-500"
              />
              <span>{option}</span>
            </label>
          ))}
        </div>
      </fieldset>

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
