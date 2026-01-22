'use client'

const events = [
  { id: 't1', time: '21:05', title: 'Victim arrives at Harbor' },
  { id: 't2', time: '21:20', title: 'Witness hears a scream' },
  { id: 't3', time: '21:45', title: 'Security alarm triggered' }
]

export default function TimelinePanel() {
  return (
    <section className="flex w-full flex-col gap-6">
      <header className="flex flex-col gap-2">
        <span className="text-xs font-semibold uppercase tracking-[0.3em] text-slate-500">Timeline</span>
        <h1 className="text-3xl font-semibold text-slate-900">Timeline</h1>
      </header>

      <div className="rounded-3xl border border-slate-200 bg-white/90 p-6 shadow-lg">
        <div className="mb-6 rounded-2xl border border-amber-200 bg-amber-50 px-4 py-3 text-sm font-semibold text-amber-900">
          Current time: 22:10
        </div>
        <div className="flex flex-col gap-4">
          {events.map((event) => (
            <div key={event.id} data-testid="timeline-item" className="flex items-start gap-4">
              <div className="text-xs font-semibold text-slate-500">{event.time}</div>
              <div className="rounded-xl border border-slate-100 bg-slate-50 px-4 py-3 text-sm text-slate-800 shadow-sm">
                {event.title}
              </div>
            </div>
          ))}
        </div>
      </div>
    </section>
  )
}
