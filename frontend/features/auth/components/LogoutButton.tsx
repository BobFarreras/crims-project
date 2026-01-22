"use client";
import { useRouter } from 'next/navigation';
import { LogOut } from 'lucide-react';
import { authService } from '../services/auth.service';
import { useState } from 'react';

export default function LogoutButton() {
  const router = useRouter();
  const [isLoading, setIsLoading] = useState(false);

  const handleLogout = async () => {
    setIsLoading(true);
    try {
      // 1. Diem al backend que mati la cookie
      await authService.logout();
      
      // 2. Refresquem el router per netejar la caché de Next.js
      router.refresh();
      
      // 3. Ens n'anem al login
      router.push('/login');
    } catch (error) {
      console.error("Error al sortir:", error);
    } finally {
      setIsLoading(false);
    }
  };

  return (
    <button
      onClick={handleLogout}
      disabled={isLoading}
      className="flex items-center gap-2 rounded-lg border border-red-200 px-4 py-2 text-sm font-semibold text-red-600 transition hover:bg-red-50 hover:text-red-700 disabled:opacity-50"
    >
      <LogOut size={16} />
      {isLoading ? "Sortint..." : "Tancar Sessió"}
    </button>
  );
}