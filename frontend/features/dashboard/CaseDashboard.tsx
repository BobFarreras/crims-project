'use client'

// 1. Importem el botó de Logout que hem creat
import LogoutButton from '@/features/auth/components/LogoutButton';

const summary = [
  { id: 'status', label: 'Status', value: 'Investigation' },
  { id: 'clues', label: 'Clues', value: '12' },
  { id: 'hypotheses', label: 'Hypotheses', value: '3' }
]

const players = [
  { id: 'p1', name: 'Detective Mara', capabilities: ['DETECTIVE'] },
  { id: 'p2', name: 'Forensic Leo', capabilities: ['FORENSIC', 'ANALYST'] },
  { id: 'p3', name: 'Analyst Quin', capabilities: ['ANALYST'] }
]

export default function CaseDashboard() {
  return (
    <section className="flex w-full flex-col gap-6 p-6"> {/* Afegit padding p-6 per que respiri */}
      
      {/* 2. HEADER AMB LOGOUT */}
      <header className="flex items-center justify-between">
        <div>
          <p className="text-xs font-semibold uppercase tracking-[0.3em] text-slate-500">Case Overview</p>
          <h1 className="text-3xl font-semibold text-slate-900">Case Dashboard</h1>
        </div>
        
        {/* Agrupem el badge Live i el botó de Logout */}
        <div className="flex items-center gap-4">
          <span className="hidden rounded-full bg-emerald-100 px-3 py-1 text-xs font-semibold text-emerald-700 sm:block">
            Live
          </span>
          <LogoutButton />
        </div>
      </header>

      <div className="grid gap-4 md:grid-cols-3">
        {summary.map((item) => (
          <article
            key={item.id}
            data-testid="summary-card"
            className="rounded-2xl border border-slate-200 bg-white/90 p-4 shadow-sm transition hover:shadow-md"
          >
            <p className="text-xs font-semibold uppercase tracking-[0.2em] text-slate-500">{item.label}</p>
            <p className="mt-3 text-2xl font-semibold text-slate-900">{item.value}</p>
          </article>
        ))}
      </div>

      <div className="rounded-3xl border border-slate-200 bg-white/90 p-6 shadow-lg">
        <header className="mb-4 flex items-center justify-between">
             <h2 className="text-sm font-semibold uppercase tracking-[0.3em] text-slate-500">Active Players</h2>
             <span className="text-xs font-bold text-slate-400">3 Online</span>
        </header>
       
        <div className="flex flex-col gap-3">
          {players.map((player) => (
            <div
              key={player.id}
              data-testid="player-row"
              className="flex items-center justify-between rounded-xl border border-slate-100 bg-slate-50 px-4 py-3 transition hover:bg-slate-100"
            >
              <div className="flex items-center gap-3">
                {/* Petit avatar placeholder */}
                <div className="flex h-8 w-8 items-center justify-center rounded-full bg-slate-200 text-xs font-bold text-slate-600">
                    {player.name.charAt(0)}
                </div>
                <div>
                    <p className="text-sm font-bold text-slate-900">{player.name}</p>
                    <p className="text-[10px] font-bold uppercase tracking-wider text-amber-600">
                      {player.capabilities.join(' + ')}
                    </p>
                </div>
              </div>
              <div className="h-2 w-2 rounded-full bg-emerald-500 animate-pulse" title="Online"></div>
            </div>
          ))}
        </div>
      </div>
    </section>
  )
}
