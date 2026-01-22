'use client'

const hypotheses = [
  { id: 'h1', title: 'Harbor escape', status: 'PLAUSIBLE', score: 68 },
  { id: 'h2', title: 'Inside job', status: 'FRINGE', score: 42 },
  { id: 'h3', title: 'False alibi', status: 'STRONG', score: 82 }
]

export default function HypothesisBoard() {
  return (
    <section className="flex w-full flex-col gap-6">
      <header className="flex flex-col gap-2">
        <span className="text-xs font-semibold uppercase tracking-[0.3em] text-slate-500">Case Board</span>
        <h1 className="text-3xl font-semibold text-slate-900">Hypotheses</h1>
      </header>

      <div className="grid gap-4 md:grid-cols-3">
        {hypotheses.map((item) => (
          <article
            key={item.id}
            data-testid="hypothesis-card"
            className="rounded-2xl border border-slate-200 bg-white/90 p-4 shadow-sm"
          >
            <p className="text-xs font-semibold uppercase tracking-[0.2em] text-amber-600">{item.status}</p>
            <h2 className="mt-2 text-lg font-semibold text-slate-900">{item.title}</h2>
            <p className="mt-4 text-sm text-slate-600">Strength {item.score}%</p>
          </article>
        ))}
      </div>
    </section>
  )
}
