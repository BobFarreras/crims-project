"use client"
import { Cormorant_Garamond, Crimson_Text } from 'next/font/google'
import LobbyForm from '@/features/lobby/LobbyForm'

const headingFont = Cormorant_Garamond({ subsets: ['latin'], weight: ['600', '700'] })
const bodyFont = Crimson_Text({ subsets: ['latin'], weight: ['400', '600'] })

import Link from 'next/link'
import { useEffect } from 'react'
import { getCookie } from '@/utils/cookie'

export default function Home() {
  // Si hay sesión, redirigir al dashboard (cliente)
  useEffect(() => {
    const token = getCookie('token')
    if (token) {
      window.location.assign('/game/dashboard')
    }
  }, [])
  return (
    <div className={`${bodyFont.className} relative min-h-screen bg-amber-50 text-zinc-900`}>
      <div className="bg-linear-to-b from-amber-100 to-amber-200 absolute inset-0 -z-10 opacity-60" />
      <main className="relative mx-auto flex min-h-screen w-full max-w-6xl flex-col gap-12 p-6 md:flex-row md:items-center md:justify-between md:p-12">
        <section className="flex-1 flex flex-col justify-center gap-6 md:max-w-lg">
          <span className="text-xs font-semibold uppercase tracking-[0.4em] text-amber-700">Crims de Mitjanit</span>
          <h1 className={`${headingFont.className} text-4xl font-semibold leading-tight text-zinc-900 sm:text-5xl`}>Join the investigation and shape the night.</h1>
          <p className="text-lg leading-7 text-zinc-700">Entrar un código de juego, elegir un rol y sumergirse en un caso en vivo. Cada pista cambia la historia.</p>
          <div className="flex gap-4">
            <Link href="/login">
              <a className="inline-flex items-center justify-center rounded-xl bg-amber-600 px-6 py-3 text-white font-semibold shadow hover:bg-amber-700">Login</a>
            </Link>
            <Link href="/register">
              <a className="inline-flex items-center justify-center rounded-xl bg-zinc-900 px-6 py-3 text-white font-semibold shadow hover:bg-zinc-700">Register</a>
            </Link>
          </div>
        </section>
        <section className="w-full max-w-md rounded-3xl border border-amber-100 bg-white/90 p-6 shadow-xl md:ml-0 md:flex-1">
          <LobbyForm
            roles={["DETECTIVE", "FORENSIC", "ANALYST", "INTERROGATOR"]}
            onSubmit={() => undefined}
          />
        </section>
      </main>
    </div>
  )
}
