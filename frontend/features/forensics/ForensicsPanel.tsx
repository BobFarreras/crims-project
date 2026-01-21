'use client'

const analyses = [
  { id: 'f1', clue: 'Fiber sample', result: 'Match', confidence: '92%' },
  { id: 'f2', clue: 'Glass shard', result: 'Inconclusive', confidence: '48%' },
  { id: 'f3', clue: 'Footprint', result: 'Match', confidence: '85%' }
]

export default function ForensicsPanel() {
  return (
    <section className="flex w-full flex-col gap-6">
      <header className="flex flex-col gap-2">
        <span className="text-xs font-semibold uppercase tracking-[0.3em] text-slate-500">Lab</span>
        <h1 className="text-3xl font-semibold text-slate-900">Forensics</h1>
      </header>

      <div className="grid gap-4 md:grid-cols-3">
        {analyses.map((analysis) => (
          <article
            key={analysis.id}
            data-testid="forensic-card"
            className="rounded-2xl border border-slate-200 bg-white/90 p-4 shadow-sm"
          >
            <p className="text-xs font-semibold uppercase tracking-[0.2em] text-slate-500">{analysis.clue}</p>
            <p className="mt-3 text-xl font-semibold text-slate-900">{analysis.result}</p>
            <p className="mt-1 text-sm text-emerald-600">Confidence {analysis.confidence}</p>
          </article>
        ))}
      </div>
    </section>
  )
}
