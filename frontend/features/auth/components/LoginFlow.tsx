"use client";
import { useLogin } from '../hooks/useLogin';
import LoginForm from '../components/LoginForm';
import { AlertCircle } from 'lucide-react';

export default function LoginFlow() {
  const { login, isLoading, error } = useLogin();

  return (
    <div>
      {error && (
        <div className="mb-4 flex items-center gap-2 rounded-lg bg-red-50 p-3 text-sm font-medium text-red-600 animate-in fade-in slide-in-from-top-1">
          <AlertCircle size={18} />
          {error}
        </div>
      )}
      
      {/* ✅ CANVI: Reben 'data' (objecte) i l'obrim per passar-lo al hook */}
      <LoginForm 
        onSubmit={(data) => login(data.email, data.password)} 
        isLoading={isLoading}
      />
    </div>
  );
}
