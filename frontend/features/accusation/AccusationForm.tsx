'use client'

const suspects = ['Witness A', 'Dock Worker', 'Harbor Chief']
const motives = ['Debt', 'Jealousy', 'Cover-up']
const evidences = ['Glass shard', 'Fingerprints', 'Alibi gap']

export default function AccusationForm() {
  return (
    <section className="flex w-full flex-col gap-6">
      <header className="flex flex-col gap-2">
        <span className="text-xs font-semibold uppercase tracking-[0.3em] text-slate-500">Final Phase</span>
        <h1 className="text-3xl font-semibold text-slate-900">Accusation</h1>
      </header>

      <div className="rounded-3xl border border-slate-200 bg-white/90 p-6 shadow-lg">
        <div className="grid gap-4">
          <label className="text-xs font-semibold uppercase tracking-[0.2em] text-slate-500">
            Suspect
            <select className="mt-2 h-12 w-full rounded-xl border border-slate-200 bg-white px-4 text-slate-900">
              {suspects.map((item) => (
                <option key={item}>{item}</option>
              ))}
            </select>
          </label>

          <label className="text-xs font-semibold uppercase tracking-[0.2em] text-slate-500">
            Motive
            <select className="mt-2 h-12 w-full rounded-xl border border-slate-200 bg-white px-4 text-slate-900">
              {motives.map((item) => (
                <option key={item}>{item}</option>
              ))}
            </select>
          </label>

          <label className="text-xs font-semibold uppercase tracking-[0.2em] text-slate-500">
            Evidence
            <select className="mt-2 h-12 w-full rounded-xl border border-slate-200 bg-white px-4 text-slate-900">
              {evidences.map((item) => (
                <option key={item}>{item}</option>
              ))}
            </select>
          </label>
        </div>

        <button className="mt-6 h-12 w-full rounded-xl bg-rose-600 text-sm font-semibold text-white shadow-md transition hover:bg-rose-700">
          Submit Accusation
        </button>
      </div>
    </section>
  )
}
