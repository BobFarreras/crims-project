import { render, screen, fireEvent, waitFor } from '@testing-library/react'
import LoginFlow from '../components/LoginFlow'
import { authService } from '../services/auth.service'
import { vi } from 'vitest'

// Mock del servei d'auth
vi.mock('../../services/auth.service')

// 🔥 FIX 4: Mock del Router (necessari pel useLogin)
vi.mock('next/navigation', () => ({
  useRouter: () => ({
    push: vi.fn(),
  }),
}))

describe('LoginFlow', () => {
  it('renders login form and submits successfully', async () => {
    // Simulem que el login va bé
    vi.mocked(authService.login).mockResolvedValue({ 
        token: 'fake-jwt', 
        user: { id: '1', username: 'test', role: 'admin' } 
    })

    render(<LoginFlow />)

    // 🔥 FIX 5: Textos en català
    const userInput = screen.getByLabelText(/Usuari o Email/i)
    const passInput = screen.getByLabelText(/Contrasenya/i)
    const loginBtn = screen.getByRole('button', { name: /Entrar al Cas/i })

    fireEvent.change(userInput, { target: { value: 'detective1' } })
    fireEvent.change(passInput, { target: { value: 'password123' } })
    fireEvent.click(loginBtn)

    await waitFor(() => {
      // Ara comprovem que crida al servei amb els arguments separats
      expect(authService.login).toHaveBeenCalledWith('detective1', 'password123')
    })
  })
})