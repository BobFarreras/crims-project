"use client";
import { useState } from 'react'

type Props = {
  // Ara acceptem tots els camps necessaris
  onSubmit: (payload: { username: string; email: string; password: string; passwordConfirm: string }) => void
  isLoading?: boolean
}

export default function RegisterForm({ onSubmit, isLoading }: Props) {
  const [formData, setFormData] = useState({
    username: '',
    email: '',
    password: '',
    passwordConfirm: ''
  })

  const handleChange = (e: React.ChangeEvent<HTMLInputElement>) => {
    setFormData({ ...formData, [e.target.name]: e.target.value })
  }

  const handleSubmit = (e: React.FormEvent<HTMLFormElement>) => {
    e.preventDefault()
    onSubmit(formData)
  }

  return (
    <form onSubmit={handleSubmit} className="flex w-full max-w-md flex-col gap-4">
      <div className="flex flex-col gap-1">
        <label className="text-sm font-bold" htmlFor="username">Username</label>
        <input name="username" placeholder="detective_01" value={formData.username} onChange={handleChange} className="h-10 rounded border px-3" required />
      </div>
      
      <div className="flex flex-col gap-1">
        <label className="text-sm font-bold" htmlFor="email">Email</label>
        <input name="email" type="email" placeholder="det@crims.com" value={formData.email} onChange={handleChange} className="h-10 rounded border px-3" required />
      </div>

      <div className="flex flex-col gap-1">
        <label className="text-sm font-bold" htmlFor="password">Password</label>
        <input name="password" type="password" value={formData.password} onChange={handleChange} className="h-10 rounded border px-3" required />
      </div>

      <div className="flex flex-col gap-1">
        <label className="text-sm font-bold" htmlFor="passwordConfirm">Confirm Password</label>
        <input name="passwordConfirm" type="password" value={formData.passwordConfirm} onChange={handleChange} className="h-10 rounded border px-3" required />
      </div>

      <button 
        type="submit" 
        disabled={isLoading}
        className="mt-2 h-10 rounded bg-zinc-900 text-white font-bold hover:bg-zinc-700 disabled:opacity-50"
      >
        {isLoading ? 'Registering...' : 'Start Investigation'}
      </button>
    </form>
  )
}