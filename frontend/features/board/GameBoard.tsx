'use client'

const mockNodes = [
  { id: 'clue-1', title: 'Broken Glass', type: 'Clue' },
  { id: 'person-1', title: 'Witness A', type: 'Person' },
  { id: 'hyp-1', title: 'Primary Theory', type: 'Hypothesis' }
]

export default function GameBoard() {
  return (
    <section className="flex w-full flex-col gap-6">
      <header className="flex flex-col gap-2">
        <span className="text-xs font-semibold uppercase tracking-[0.3em] text-slate-500">
          Investigation
        </span>
        <h1 className="text-3xl font-semibold text-slate-900">Case Board</h1>
      </header>

      <div className="grid gap-6 lg:grid-cols-[1fr_280px]">
        <div
          data-testid="board-area"
          className="min-h-[420px] rounded-3xl border border-slate-200 bg-white/80 p-6 shadow-lg"
        >
          <div className="grid gap-4 sm:grid-cols-2">
            {mockNodes.map((node) => (
              <article
                key={node.id}
                data-testid="board-node"
                className="rounded-2xl border border-slate-100 bg-slate-50 p-4 shadow-sm"
              >
                <p className="text-xs font-semibold uppercase tracking-[0.2em] text-amber-600">
                  {node.type}
                </p>
                <h2 className="mt-2 text-lg font-semibold text-slate-900">{node.title}</h2>
              </article>
            ))}
          </div>
        </div>

        <aside className="rounded-3xl border border-slate-200 bg-white/80 p-6 shadow-lg">
          <h3 className="text-sm font-semibold uppercase tracking-[0.3em] text-slate-500">Filters</h3>
          <div className="mt-4 flex flex-col gap-3">
            <button className="rounded-xl border border-amber-200 bg-amber-100 px-4 py-2 text-sm font-semibold text-amber-900">
              All Nodes
            </button>
            <button className="rounded-xl border border-slate-200 bg-white px-4 py-2 text-sm font-semibold text-slate-600">
              Clues
            </button>
            <button className="rounded-xl border border-slate-200 bg-white px-4 py-2 text-sm font-semibold text-slate-600">
              Persons
            </button>
          </div>
        </aside>
      </div>
    </section>
  )
}
