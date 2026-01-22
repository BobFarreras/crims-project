"use client";
import { useState } from 'react';
import { Eye, EyeOff, AlertCircle, CheckCircle2 } from 'lucide-react';

type Props = {
  onSubmit: (payload: { username: string; email: string; password: string; passwordConfirm: string }) => void;
  isLoading?: boolean;
};

export default function RegisterForm({ onSubmit, isLoading }: Props) {
  // Estats del formulari
  const [formData, setFormData] = useState({
    username: '',
    email: '',
    password: '',
    passwordConfirm: ''
  });

  // Estats de UI
  const [showPassword, setShowPassword] = useState(false);
  const [clientError, setClientError] = useState<string | null>(null);

  const handleChange = (e: React.ChangeEvent<HTMLInputElement>) => {
    setFormData({ ...formData, [e.target.name]: e.target.value });
    // Netejar error quan l'usuari escriu
    if (clientError) setClientError(null);
  };

  const handleSubmit = (e: React.FormEvent<HTMLFormElement>) => {
    e.preventDefault();
    
    // 🛡️ 1. VALIDACIONS DE SEGURETAT (CLIENT-SIDE)
    // Això estalvia peticions al servidor i millora la UX
    if (formData.username.includes(' ')) {
      setClientError("L'usuari no pot contenir espais.");
      return;
    }
    if (formData.password.length < 8) {
      setClientError("La contrasenya ha de tenir almenys 8 caràcters per seguretat.");
      return;
    }
    if (formData.password !== formData.passwordConfirm) {
      setClientError("Les contrasenyes no coincideixen.");
      return;
    }

    // Si tot està bé, enviem al backend
    onSubmit(formData);
  };

  return (
    <form onSubmit={handleSubmit} className="flex w-full max-w-md flex-col gap-5 bg-white p-6 rounded-2xl shadow-sm border border-zinc-100">
      
      {/* USERNAME */}
      <div className="flex flex-col gap-1.5">
        <label className="text-xs font-bold uppercase tracking-wider text-zinc-500" htmlFor="username">
          Nom de Detectiu (Username)
        </label>
        <input 
          id="username"
          name="username" 
          placeholder="detective_boby" 
          value={formData.username} 
          onChange={handleChange} 
          className="h-12 rounded-xl border border-zinc-200 bg-zinc-50 px-4 transition focus:border-amber-500 focus:bg-white focus:outline-none" 
          required 
        />
      </div>
      
      {/* EMAIL */}
      <div className="flex flex-col gap-1.5">
        <label className="text-xs font-bold uppercase tracking-wider text-zinc-500" htmlFor="email">
          Correu Electrònic
        </label>
        <input 
          id="email"
          name="email" 
          type="email" 
          placeholder="contact@agency.com" 
          value={formData.email} 
          onChange={handleChange} 
          className="h-12 rounded-xl border border-zinc-200 bg-zinc-50 px-4 transition focus:border-amber-500 focus:bg-white focus:outline-none" 
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
            name="password" 
            // 👁️ Aquí fem la màgia de mostrar/amagar
            type={showPassword ? "text" : "password"} 
            value={formData.password} 
            onChange={handleChange} 
            className="h-12 w-full rounded-xl border border-zinc-200 bg-zinc-50 px-4 pr-12 transition focus:border-amber-500 focus:bg-white focus:outline-none" 
            required 
          />
          {/* Botó de l'ull */}
          <button
            type="button"
            onClick={() => setShowPassword(!showPassword)}
            className="absolute right-3 top-1/2 -translate-y-1/2 text-zinc-400 hover:text-zinc-600"
            aria-label={showPassword ? "Amagar contrasenya" : "Mostrar contrasenya"}
          >
            {showPassword ? <EyeOff size={20} /> : <Eye size={20} />}
          </button>
        </div>
        {/* Helper text */}
        <p className="text-xs text-zinc-400">Mínim 8 caràcters.</p>
      </div>

      {/* CONFIRM PASSWORD */}
      <div className="flex flex-col gap-1.5">
        <label className="text-xs font-bold uppercase tracking-wider text-zinc-500" htmlFor="passwordConfirm">
          Confirmar Contrasenya
        </label>
        <input 
          id="passwordConfirm"
          name="passwordConfirm" 
          type="password" 
          value={formData.passwordConfirm} 
          onChange={handleChange} 
          className={`h-12 rounded-xl border bg-zinc-50 px-4 transition focus:bg-white focus:outline-none ${
            formData.passwordConfirm && formData.password !== formData.passwordConfirm 
              ? 'border-red-300 focus:border-red-500' 
              : 'border-zinc-200 focus:border-amber-500'
          }`}
          required 
        />
      </div>

      {/* 🚨 ZONA D'ERRORS (FEEDBACK) */}
      {clientError && (
        <div className="flex items-center gap-2 rounded-lg bg-red-50 p-3 text-sm font-medium text-red-600">
          <AlertCircle size={18} />
          {clientError}
        </div>
      )}

      {/* SUBMIT BUTTON */}
      <button 
        type="submit" 
        disabled={isLoading}
        className="mt-2 flex h-12 items-center justify-center gap-2 rounded-xl bg-zinc-900 text-base font-bold text-white shadow-lg transition hover:bg-zinc-800 hover:shadow-xl disabled:cursor-not-allowed disabled:opacity-70"
      >
        {isLoading ? (
          <span className="animate-pulse">Processant...</span>
        ) : (
          <>
            <CheckCircle2 size={18} />
            Crear Compte
          </>
        )}
      </button>
    </form>
  );
}