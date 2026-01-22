'use client'

import { useState } from 'react'
import { useRouter } from 'next/navigation'
import LogoutButton from '@/features/auth/components/LogoutButton'

const modes = [
  { id: 'solo', label: 'Solo', description: 'Un jugador amb totes les capacitats.' },
  { id: 'duo', label: 'Duo', description: 'Co-op ràpid en parella.' },
  { id: 'team', label: 'Equip', description: '4 rols principals en equip.' },
  { id: 'join', label: 'Sala', description: 'Entrar amb codi.' },
]

const activeCases = [
  { id: 'case-1', title: 'El Silenci del Port', status: 'Investigation', players: 3 },
  { id: 'case-2', title: 'La Nit del Ferro', status: 'Briefing', players: 1 },
]

const ranking = [
  { id: 'rank-1', name: 'Mara', score: 980 },
  { id: 'rank-2', name: 'Leo', score: 910 },
  { id: 'rank-3', name: 'Quin', score: 870 },
]

export default function DashboardHub() {
  const router = useRouter()
  const [playerName, setPlayerName] = useState('Detective')

  return (
    <section className="flex w-full flex-col gap-6 p-6">
      <header className="flex flex-col gap-3 sm:flex-row sm:items-start sm:justify-between">
        <div className="flex flex-col gap-2">
          <p className="text-xs font-semibold uppercase tracking-[0.35em] text-amber-700">Crims Hub</p>
          <h1 className="text-3xl font-semibold text-zinc-900">Tria com vols jugar</h1>
          <p className="text-sm text-zinc-600">Selecciona el mode, ajusta el perfil i entra al cas.</p>
        </div>
        <div className="self-start">
          <LogoutButton />
        </div>
      </header>

      <div className="grid gap-6 lg:grid-cols-[1.2fr_0.8fr]">
        <div className="flex flex-col gap-6">
          <div className="rounded-3xl border border-amber-100 bg-white/90 p-5 shadow-lg">
            <h2 className="text-sm font-semibold uppercase tracking-[0.3em] text-amber-600">Modes</h2>
            <div className="mt-4 grid gap-3 sm:grid-cols-2">
              {modes.map((mode) => (
                <button
                  key={mode.id}
                  type="button"
                  onClick={() => router.push(mode.id === 'join' ? '/lobby' : `/game/${mode.id}`)}
                  className="flex flex-col gap-2 rounded-2xl border border-zinc-200 bg-white px-4 py-4 text-left text-sm font-semibold text-zinc-900 shadow-sm transition hover:border-amber-400 hover:bg-amber-50"
                >
                  <span className="text-base font-semibold text-zinc-900">{mode.label}</span>
                  <span className="text-xs text-zinc-600">{mode.description}</span>
                </button>
              ))}
            </div>
          </div>

          <div className="rounded-3xl border border-zinc-200 bg-white/90 p-5 shadow-lg">
            <h2 className="text-sm font-semibold uppercase tracking-[0.3em] text-zinc-500">Casos actius</h2>
            <div className="mt-4 flex flex-col gap-3">
              {activeCases.map((item) => (
                <article
                  key={item.id}
                  className="flex flex-col gap-2 rounded-2xl border border-zinc-100 bg-zinc-50 px-4 py-3"
                >
                  <div className="flex items-center justify-between">
                    <h3 className="text-sm font-semibold text-zinc-900">{item.title}</h3>
                    <span className="rounded-full bg-amber-100 px-3 py-1 text-[10px] font-semibold uppercase tracking-wider text-amber-700">
                      {item.status}
                    </span>
                  </div>
                  <p className="text-xs text-zinc-500">{item.players} jugadors en línia</p>
                  <button
                    type="button"
                    className="mt-1 inline-flex h-9 items-center justify-center rounded-xl bg-zinc-900 text-xs font-semibold uppercase tracking-wider text-white"
                  >
                    Continuar
                  </button>
                </article>
              ))}
            </div>
          </div>
        </div>

        <aside className="flex flex-col gap-6">
          <div className="rounded-3xl border border-zinc-200 bg-white/90 p-5 shadow-lg">
            <h2 className="text-sm font-semibold uppercase tracking-[0.3em] text-zinc-500">Perfil</h2>
            <div className="mt-4 flex flex-col gap-3">
              <label className="text-xs font-semibold uppercase tracking-[0.2em] text-zinc-500" htmlFor="player-name">
                Nom del jugador
              </label>
              <input
                id="player-name"
                name="playerName"
                value={playerName}
                onChange={(event) => setPlayerName(event.target.value)}
                className="h-11 rounded-xl border border-zinc-200 bg-white px-3 text-sm text-zinc-900 focus:border-amber-400 focus:outline-none"
              />
              <div className="flex items-center gap-3 rounded-2xl border border-amber-100 bg-amber-50 px-3 py-2">
                <div className="flex h-10 w-10 items-center justify-center rounded-full bg-amber-200 text-sm font-bold text-amber-900">
                  {playerName.charAt(0).toUpperCase()}
                </div>
                <div>
                  <p className="text-xs font-semibold uppercase tracking-widest text-amber-700">Alias</p>
                  <p className="text-sm font-semibold text-zinc-900">{playerName || 'Detective'}</p>
                </div>
              </div>
            </div>
          </div>

          <div className="rounded-3xl border border-zinc-200 bg-white/90 p-5 shadow-lg">
            <div className="flex items-center justify-between">
              <h2 className="text-sm font-semibold uppercase tracking-[0.3em] text-zinc-500">Ranking</h2>
              <span className="text-xs font-semibold text-amber-600">Setmanal</span>
            </div>
            <div className="mt-4 flex flex-col gap-2">
              {ranking.map((item, index) => (
                <div key={item.id} className="flex items-center justify-between rounded-xl border border-zinc-100 bg-zinc-50 px-3 py-2">
                  <div className="flex items-center gap-3">
                    <span className="text-xs font-semibold text-zinc-400">#{index + 1}</span>
                    <span className="text-sm font-semibold text-zinc-900">{item.name}</span>
                  </div>
                  <span className="text-xs font-semibold text-zinc-600">{item.score}</span>
                </div>
              ))}
            </div>
            <button
              type="button"
              className="mt-4 w-full rounded-xl border border-amber-200 bg-amber-50 py-2 text-xs font-semibold uppercase tracking-wider text-amber-700"
            >
              Veure ranking complet
            </button>
          </div>
        </aside>
      </div>
    </section>
  )
}
