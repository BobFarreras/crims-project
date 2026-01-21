'use client'

const summary = [
  { id: 'status', label: 'Status', value: 'Investigation' },
  { id: 'clues', label: 'Clues', value: '12' },
  { id: 'hypotheses', label: 'Hypotheses', value: '3' }
]

const players = [
  { id: 'p1', name: 'Detective Mara', role: 'DETECTIVE' },
  { id: 'p2', name: 'Forensic Leo', role: 'FORENSIC' },
  { id: 'p3', name: 'Analyst Quin', role: 'ANALYST' }
]

export default function CaseDashboard() {
  return (
    <section className="flex w-full flex-col gap-6">
      <header className="flex items-center justify-between">
        <div>
          <p className="text-xs font-semibold uppercase tracking-[0.3em] text-slate-500">Case Overview</p>
          <h1 className="text-3xl font-semibold text-slate-900">Case Dashboard</h1>
        </div>
        <span className="rounded-full bg-emerald-100 px-3 py-1 text-xs font-semibold text-emerald-700">Live</span>
      </header>

      <div className="grid gap-4 md:grid-cols-3">
        {summary.map((item) => (
          <article
            key={item.id}
            data-testid="summary-card"
            className="rounded-2xl border border-slate-200 bg-white/90 p-4 shadow-sm"
          >
            <p className="text-xs font-semibold uppercase tracking-[0.2em] text-slate-500">{item.label}</p>
            <p className="mt-3 text-2xl font-semibold text-slate-900">{item.value}</p>
          </article>
        ))}
      </div>

      <div className="rounded-3xl border border-slate-200 bg-white/90 p-6 shadow-lg">
        <h2 className="text-sm font-semibold uppercase tracking-[0.3em] text-slate-500">Active Players</h2>
        <div className="mt-4 flex flex-col gap-3">
          {players.map((player) => (
            <div
              key={player.id}
              data-testid="player-row"
              className="flex items-center justify-between rounded-xl border border-slate-100 bg-slate-50 px-4 py-3"
            >
              <div>
                <p className="text-base font-semibold text-slate-900">{player.name}</p>
                <p className="text-xs font-semibold uppercase tracking-[0.2em] text-amber-600">{player.role}</p>
              </div>
              <span className="text-xs font-semibold text-slate-500">Online</span>
            </div>
          ))}
        </div>
      </div>
    </section>
  )
}
