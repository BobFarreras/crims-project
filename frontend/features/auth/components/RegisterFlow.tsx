"use client";
import { useRegister } from '../hooks/useRegister'; // Importem el Hook nou
import RegisterForm from '../components/RegisterForm'; // Assegura't que la ruta és correcta

export default function RegisterFlow() {
  // Usem el hook que vam crear abans, que crida a authService -> Go Backend
  const { register, error, isLoading } = useRegister();

  return (
    <div>
      {error ? (
        <div className="mb-4 rounded bg-red-100 p-3 text-sm font-bold text-red-600">
          {error}
        </div>
      ) : null}
      
      {/* Passem les dades directament a la funció register del hook */}
      <RegisterForm 
        onSubmit={(data) => register(data)} 
        isLoading={isLoading}
      />
    </div>
  )
}