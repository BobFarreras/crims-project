import { Cormorant_Garamond, Crimson_Text } from 'next/font/google'
import LobbyForm from '@/features/lobby/LobbyForm'

const headingFont = Cormorant_Garamond({ subsets: ['latin'], weight: ['600', '700'] })
const bodyFont = Crimson_Text({ subsets: ['latin'], weight: ['400', '600'] })

export default function Home() {
  return (
    <div className={`${bodyFont.className} relative min-h-screen bg-amber-50 text-zinc-900`}>
      <div className="pointer-events-none absolute inset-0 bg-[radial-gradient(circle_at_top,_rgba(251,191,36,0.28),_transparent_60%)]" />
      <div className="pointer-events-none absolute right-0 top-0 h-64 w-64 rounded-full bg-amber-200/50 blur-3xl" />
      <main className="relative mx-auto flex min-h-screen w-full max-w-6xl flex-col gap-10 px-6 py-20 lg:flex-row lg:items-center lg:justify-between lg:px-12">
        <section className="flex flex-col gap-6 lg:max-w-xl">
          <span className="text-xs font-semibold uppercase tracking-[0.4em] text-amber-700">
            Crims de Mitjanit
          </span>
          <h1 className={`${headingFont.className} text-4xl font-semibold leading-tight text-zinc-900 sm:text-5xl`}>
            Join the investigation and shape the night.
          </h1>
          <p className="text-lg leading-7 text-zinc-700">
            Enter a game code, pick a role, and step into a live case. Every clue you connect changes the story.
          </p>
          <div className="rounded-2xl border border-amber-200 bg-white/90 p-5 shadow-lg">
            <div className="text-sm uppercase tracking-[0.3em] text-amber-600">Live Session</div>
            <div className="mt-2 text-2xl font-semibold text-zinc-900">Ready for suspects.</div>
          </div>
        </section>

        <section className="w-full max-w-md rounded-3xl border border-amber-100 bg-white/90 p-6 shadow-xl backdrop-blur">
          <LobbyForm
            roles={['DETECTIVE', 'FORENSIC', 'ANALYST', 'INTERROGATOR']}
            onSubmit={() => undefined}
          />
        </section>
      </main>
    </div>
  )
}
