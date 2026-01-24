'use client'

import { useEffect, useState } from 'react'
import { useRouter } from 'next/navigation'
import LogoutButton from '@/features/auth/components/LogoutButton'
import { profileService } from '@/features/profile/services/profile.service'
import { lobbyService } from '@/features/lobby/services/lobby.service'
import { motion } from 'framer-motion'

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
  const [playerName, setPlayerName] = useState('')
  const [isSaving, setIsSaving] = useState(false)
  const [profileError, setProfileError] = useState('')
  const [userId, setUserId] = useState('')
  const [lobbyCode, setLobbyCode] = useState('')
  const [joinCode, setJoinCode] = useState('')
  const [lobbyError, setLobbyError] = useState('')
  const [isLobbyLoading, setIsLobbyLoading] = useState(false)

  useEffect(() => {
    let isMounted = true

    profileService
      .getProfile()
      .then((data) => {
        if (isMounted) {
          setPlayerName(data.user.name || 'Detective')
          setUserId(data.user.id || '')
        }
      })
      .catch(() => {
        if (isMounted) {
          setPlayerName('Detective')
        }
      })

    return () => {
      isMounted = false
    }
  }, [])

  const handleSaveProfile = async () => {
    setIsSaving(true)
    setProfileError('')

    try {
      const response = await profileService.updateProfile(playerName.trim())
      setPlayerName(response.user.name || playerName)
    } catch (error) {
      const message = error instanceof Error ? error.message : 'No hem pogut guardar el perfil.'
      setProfileError(message)
    } finally {
      setIsSaving(false)
    }
  }

  const handleCreateLobby = async (mode: string) => {
    if (!userId) {
      setLobbyError('No tenim usuari actiu. Torna a iniciar sessio.')
      return
    }

    setIsLobbyLoading(true)
    setLobbyError('')

    try {
      const response = await lobbyService.createLobby(userId, mode === 'solo' ? ['DETECTIVE', 'FORENSIC', 'ANALYST', 'INTERROGATOR'] : [])
      setLobbyCode(response.game.code)
      setJoinCode(response.game.code)
      if (mode === 'solo') {
        router.push('/game')
      } else {
        router.push(`/lobby/${response.game.code}`)
      }
    } catch (error) {
      const message = error instanceof Error ? error.message : 'No hem pogut crear la sala.'
      setLobbyError(message)
    } finally {
      setIsLobbyLoading(false)
    }
  }

  const handleJoinLobby = async () => {
    if (!userId) {
      setLobbyError('No tenim usuari actiu. Torna a iniciar sessio.')
      return
    }

    if (!joinCode.trim()) {
      setLobbyError('Introdueix el codi de sala.')
      return
    }

    setIsLobbyLoading(true)
    setLobbyError('')

    try {
      const code = joinCode.trim()
      await lobbyService.joinLobby(code, userId, [])
      router.push(`/lobby/${code}`)
    } catch (error) {
      const message = error instanceof Error ? error.message : 'No hem pogut unir-nos a la sala.'
      setLobbyError(message)
    } finally {
      setIsLobbyLoading(false)
    }
  }

  return (
    <section className="relative flex w-full flex-col gap-6 overflow-hidden p-6">
      <div className="pointer-events-none absolute -top-32 right-10 h-64 w-64 rounded-full bg-amber-200/40 blur-3xl" />
      <div className="pointer-events-none absolute -bottom-32 left-0 h-72 w-72 rounded-full bg-red-200/30 blur-3xl" />

      <motion.header
        initial={{ opacity: 0, y: 12 }}
        animate={{ opacity: 1, y: 0 }}
        transition={{ duration: 0.6 }}
        className="flex flex-col gap-3 sm:flex-row sm:items-start sm:justify-between"
      >
        <div className="flex flex-col gap-2">
          <p className="text-xs font-semibold uppercase tracking-[0.35em] text-amber-700">Crims Hub</p>
          <h1 className="text-3xl font-semibold text-zinc-900">Tria com vols jugar</h1>
          <p className="text-sm text-zinc-600">Selecciona el mode, ajusta el perfil i entra al cas.</p>
        </div>
        <div className="self-start">
          <LogoutButton />
        </div>
      </motion.header>

      <div className="grid gap-6 lg:grid-cols-[1.2fr_0.8fr]">
        <div className="flex flex-col gap-6">
          <motion.div
            initial={{ opacity: 0, y: 18 }}
            animate={{ opacity: 1, y: 0 }}
            transition={{ duration: 0.6, delay: 0.05 }}
            className="rounded-[28px] border border-amber-100 bg-white/90 p-6 shadow-[0_20px_40px_rgba(31,41,55,0.12)]"
          >
            <div className="flex items-center justify-between">
              <div>
                <h2 className="text-sm font-semibold uppercase tracking-[0.3em] text-amber-600">Crear sala</h2>
                <p className="mt-2 text-sm text-zinc-600">
                  Genera un codi privat per reunir l'equip. El lobby és l'espai de preparació abans del cas.
                </p>
              </div>
              <div className="hidden h-12 w-12 items-center justify-center rounded-2xl bg-amber-100 text-sm font-bold text-amber-700 sm:flex">
                HUB
              </div>
            </div>

            <div className="mt-5 grid gap-3 sm:grid-cols-2">
              {modes.map((mode) => (
                <button
                  key={mode.id}
                  type="button"
                  aria-label={mode.label}
                  onClick={() => (mode.id === 'join' ? undefined : handleCreateLobby(mode.id))}
                  className="group flex flex-col gap-2 rounded-2xl border border-zinc-200 bg-white px-4 py-4 text-left text-sm font-semibold text-zinc-900 shadow-sm transition hover:-translate-y-0.5 hover:border-amber-400 hover:bg-amber-50"
                >
                  <span className="text-base font-semibold text-zinc-900">{mode.label}</span>
                  <span className="text-xs text-zinc-600">{mode.description}</span>
                  {mode.id !== 'join' ? (
                    <span
                      className="mt-2 text-[10px] font-semibold uppercase tracking-[0.2em] text-amber-600"
                      aria-hidden="true"
                    >
                      Crear sala
                    </span>
                  ) : null}
                </button>
              ))}
            </div>

            <div className="mt-5 grid gap-4 lg:grid-cols-[1fr_0.9fr]">
              <div className="rounded-2xl border border-amber-100 bg-amber-50/70 p-4">
                <p className="text-xs font-semibold uppercase tracking-[0.3em] text-amber-600">Unir-se a sala</p>
                <p className="mt-1 text-xs text-zinc-600">
                  Escriu el codi de 4 lletres per entrar a una sala activa.
                </p>
                <div className="mt-3 flex items-center gap-2">
                  <input
                    value={joinCode}
                    onChange={(event) => setJoinCode(event.target.value.toUpperCase())}
                    placeholder="Codi sala"
                    className="h-10 flex-1 rounded-xl border border-amber-200 bg-white px-3 text-sm font-semibold text-zinc-800"
                  />
                  <button
                    type="button"
                    onClick={handleJoinLobby}
                    disabled={isLobbyLoading}
                    className="h-10 rounded-xl bg-zinc-900 px-4 text-xs font-semibold uppercase tracking-wider text-white disabled:opacity-60"
                  >
                    Unir-se
                  </button>
                </div>
                {lobbyError ? (
                  <p className="mt-2 text-xs font-semibold text-red-600" role="alert">
                    {lobbyError}
                  </p>
                ) : null}
              </div>

              <div className="rounded-2xl border border-zinc-200 bg-white p-4">
                <p className="text-xs font-semibold uppercase tracking-[0.3em] text-zinc-500">Codi actual</p>
                <div className="mt-2 flex items-center justify-between rounded-xl border border-dashed border-amber-200 bg-amber-50 px-3 py-2">
                  <span className="text-lg font-semibold text-zinc-900">{lobbyCode || '----'}</span>
                  <span className="text-[10px] font-semibold uppercase tracking-[0.2em] text-amber-600">Compartir</span>
                </div>
                <p className="mt-2 text-xs text-zinc-600">Aquest codi connecta tot l'equip al mateix cas.</p>
              </div>
            </div>
          </motion.div>

          <motion.div
            initial={{ opacity: 0, y: 18 }}
            animate={{ opacity: 1, y: 0 }}
            transition={{ duration: 0.6, delay: 0.1 }}
            className="rounded-[28px] border border-zinc-200 bg-white/90 p-6 shadow-[0_18px_38px_rgba(31,41,55,0.1)]"
          >
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
          </motion.div>
        </div>

        <aside className="flex flex-col gap-6">
          <motion.div
            initial={{ opacity: 0, y: 18 }}
            animate={{ opacity: 1, y: 0 }}
            transition={{ duration: 0.6, delay: 0.15 }}
            className="rounded-[28px] border border-zinc-200 bg-white/90 p-6 shadow-[0_18px_38px_rgba(31,41,55,0.1)]"
          >
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
              {profileError ? (
                <p className="text-xs font-semibold text-red-600" role="alert">
                  {profileError}
                </p>
              ) : null}
              <button
                type="button"
                onClick={handleSaveProfile}
                disabled={isSaving}
                className="h-10 rounded-xl bg-amber-500 text-xs font-semibold uppercase tracking-wider text-white transition hover:bg-amber-600 disabled:opacity-60"
              >
                {isSaving ? 'Guardant...' : 'Guardar'}
              </button>
              <div className="flex items-center gap-3 rounded-2xl border border-amber-100 bg-amber-50 px-3 py-2">
                <div className="flex h-10 w-10 items-center justify-center rounded-full bg-amber-200 text-sm font-bold text-amber-900">
                  {(playerName || 'D').charAt(0).toUpperCase()}
                </div>
                <div>
                  <p className="text-xs font-semibold uppercase tracking-widest text-amber-700">Alias</p>
                  <p className="text-sm font-semibold text-zinc-900">{playerName || 'Detective'}</p>
                </div>
              </div>
            </div>
          </motion.div>

          <motion.div
            initial={{ opacity: 0, y: 18 }}
            animate={{ opacity: 1, y: 0 }}
            transition={{ duration: 0.6, delay: 0.2 }}
            className="rounded-[28px] border border-zinc-200 bg-white/90 p-6 shadow-[0_18px_38px_rgba(31,41,55,0.1)]"
          >
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
          </motion.div>
        </aside>
      </div>
    </section>
  )
}
