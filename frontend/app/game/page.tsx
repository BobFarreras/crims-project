import GameBoard from '@/features/board/GameBoard'

export default function GamePage() {
  return (
    <div className="min-h-screen bg-slate-50 px-6 py-16 text-slate-900">
      <div className="mx-auto max-w-6xl">
        <GameBoard />
      </div>
    </div>
  )
}
