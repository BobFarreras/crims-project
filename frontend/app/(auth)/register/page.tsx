"use client";
import RegisterFlow from '@/features/auth/RegisterFlow'

export default function RegisterPage() {
  return (
    <div className="min-h-screen bg-slate-50 px-6 py-16 text-slate-900">
      <div className="mx-auto max-w-6xl">
        <RegisterFlow />
      </div>
    </div>
  )
}
