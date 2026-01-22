import LoginForm from '@/features/auth/LoginForm'

export default function LoginPage() {
  return (
    <div className="min-h-screen bg-slate-50 px-6 py-16 text-slate-900">
      <div className="mx-auto max-w-6xl">
        <LoginForm onSubmit={() => undefined} />
      </div>
    </div>
  )
}
