'use client'

const history = [
  { id: 'h1', question: 'Where were you?', answer: 'At home.' },
  { id: 'h2', question: 'Did you see anyone?', answer: 'No one.' },
  { id: 'h3', question: 'Why the delay?', answer: 'Traffic.' }
]

export default function InterrogationPanel() {
  return (
    <section className="flex w-full flex-col gap-6">
      <header className="flex flex-col gap-2">
        <span className="text-xs font-semibold uppercase tracking-[0.3em] text-slate-500">Interview</span>
        <h1 className="text-3xl font-semibold text-slate-900">Interrogation</h1>
      </header>

      <div className="grid gap-6 lg:grid-cols-[1fr_280px]">
        <div className="rounded-3xl border border-slate-200 bg-white/90 p-6 shadow-lg">
          <div className="text-xs font-semibold uppercase tracking-[0.3em] text-amber-600">Current Question</div>
          <h2 className="mt-3 text-2xl font-semibold text-slate-900">Where were you?</h2>
          <p className="mt-4 text-base leading-7 text-slate-700">"At home. I never left the apartment that night."</p>
        </div>

        <aside className="rounded-3xl border border-slate-200 bg-white/90 p-6 shadow-lg">
          <h3 className="text-sm font-semibold uppercase tracking-[0.3em] text-slate-500">History</h3>
          <div className="mt-4 flex flex-col gap-3">
            {history.map((item) => (
              <div key={item.id} data-testid="history-item" className="rounded-xl border border-slate-100 bg-slate-50 p-4">
                <p className="text-sm font-semibold text-slate-900">{item.question}</p>
                <p className="mt-2 text-xs text-slate-600">{item.answer}</p>
              </div>
            ))}
          </div>
        </aside>
      </div>
    </section>
  )
}
