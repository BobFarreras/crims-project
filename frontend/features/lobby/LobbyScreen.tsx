'use client'

import { useEffect, useState } from 'react'
import { lobbyService } from './services/lobby.service'

type LobbyScreenProps = {
  code: string
}

export default function LobbyScreen({ code }: LobbyScreenProps) {
  const [state, setState] = useState<'loading' | 'ready' | 'error'>('loading')
  const [gameState, setGameState] = useState('')

  useEffect(() => {
    let isMounted = true

    lobbyService
      .getGameByCode(code)
      .then((game) => {
        if (isMounted) {
          setGameState(game.state)
          setState('ready')
        }
      })
      .catch(() => {
        if (isMounted) {
          setState('error')
        }
      })

    return () => {
      isMounted = false
    }
  }, [code])

  if (state === 'error') {
    return (
      <section className="rounded-3xl border border-red-200 bg-red-50 p-6 text-red-700">
        <h1 className="text-xl font-semibold">No hem trobat aquesta sala</h1>
        <p className="mt-2 text-sm">Comprova el codi i torna-ho a provar.</p>
      </section>
    )
  }

  if (state === 'loading') {
    return (
      <section className="rounded-3xl border border-amber-200 bg-amber-50 p-6 text-amber-700">
        <h1 className="text-xl font-semibold">Carregant sala...</h1>
        <p className="mt-2 text-sm">Estem sincronitzant els jugadors.</p>
      </section>
    )
  }

  return (
    <section className="rounded-3xl border border-zinc-200 bg-white/90 p-6 shadow-lg">
      <div className="flex items-center justify-between">
        <div>
          <p className="text-xs font-semibold uppercase tracking-[0.3em] text-amber-600">Codi sala</p>
          <h1 className="mt-2 text-3xl font-semibold text-zinc-900">{code}</h1>
        </div>
        <div className="rounded-full bg-amber-100 px-3 py-1 text-[10px] font-semibold uppercase tracking-wider text-amber-700">
          {gameState}
        </div>
      </div>
      <p className="mt-4 text-sm text-zinc-600">Comparteix aquest codi amb l'equip. La partida començarà quan estigueu llestos.</p>
    </section>
  )
}
