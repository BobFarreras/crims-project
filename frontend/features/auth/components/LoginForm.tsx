"use client";
import { useState } from 'react';
import { Eye, EyeOff, LogIn } from 'lucide-react';

// ✅ CANVI: Ara esperem un objecte (igual que a RegisterForm)
type LoginFormData = {
  username: string;
  password: string;
};

type Props = {
  onSubmit: (data: LoginFormData) => void;
  isLoading?: boolean;
};

export default function LoginForm({ onSubmit, isLoading }: Props) {
  const [username, setUsername] = useState('');
  const [password, setPassword] = useState('');
  const [showPassword, setShowPassword] = useState(false);

  const handleSubmit = (e: React.FormEvent<HTMLFormElement>) => {
    e.preventDefault();
    // ✅ CANVI: Enviem un objecte compacte
    onSubmit({ username, password });
  };

  return (
    <form onSubmit={handleSubmit} className="flex w-full max-w-md flex-col gap-5 bg-white p-6 rounded-2xl shadow-sm border border-zinc-100">
      
      {/* USERNAME */}
      <div className="flex flex-col gap-1.5">
        <label className="text-xs font-bold uppercase tracking-wider text-zinc-500" htmlFor="username">
          Usuari o Email
        </label>
        <input 
          id="username"
          name="username" 
          value={username} 
          onChange={(e) => setUsername(e.target.value)} 
          className="h-12 rounded-xl border border-zinc-200 bg-zinc-50 px-4 transition focus:border-amber-500 focus:bg-white focus:outline-none" 
          placeholder="detective_boby"
          required 
        />
      </div>

      {/* PASSWORD */}
      <div className="flex flex-col gap-1.5">
        <label className="text-xs font-bold uppercase tracking-wider text-zinc-500" htmlFor="password">
          Contrasenya
        </label>
        <div className="relative">
          <input 
            id="password"
            type={showPassword ? "text" : "password"} 
            value={password} 
            onChange={(e) => setPassword(e.target.value)} 
            className="h-12 w-full rounded-xl border border-zinc-200 bg-zinc-50 px-4 pr-12 transition focus:border-amber-500 focus:bg-white focus:outline-none" 
            placeholder="••••••••"
            required 
          />
          <button
            type="button"
            onClick={() => setShowPassword(!showPassword)}
            className="absolute right-3 top-1/2 -translate-y-1/2 text-zinc-400 hover:text-zinc-600"
            aria-label={showPassword ? "Amagar contrasenya" : "Mostrar contrasenya"}
          >
            {showPassword ? <EyeOff size={20} /> : <Eye size={20} />}
          </button>
        </div>
      </div>

      <button 
        type="submit" 
        disabled={isLoading}
        className="mt-2 flex h-12 items-center justify-center gap-2 rounded-xl bg-amber-600 text-base font-bold text-white shadow-lg transition hover:bg-amber-700 hover:shadow-xl disabled:cursor-not-allowed disabled:opacity-70"
      >
        {isLoading ? (
          <span className="animate-pulse">Iniciant sessió...</span>
        ) : (
          <>
            <LogIn size={18} />
            Entrar al Cas
          </>
        )}
      </button>
    </form>
  );
}