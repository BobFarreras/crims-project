import LobbyScreen from '@/features/lobby/LobbyScreen'

type LobbyPageProps = {
  params: {
    code: string
  }
}

export default function LobbyPage({ params }: LobbyPageProps) {
  return (
    <div className="min-h-screen bg-amber-50 px-6 py-12 text-zinc-900">
      <div className="mx-auto max-w-5xl">
        <LobbyScreen code={params.code.toUpperCase()} />
      </div>
    </div>
  )
}
