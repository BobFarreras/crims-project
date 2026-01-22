import { render, screen, fireEvent, waitFor } from '@testing-library/react'
import LoginFlow from '../components/LoginFlow' // Assegura't que la ruta és correcta
import { authService } from '../services/auth.service'
import { vi } from 'vitest'

// 🔥 FIX: Definim el mock explícitament perquè vi.fn() estigui disponible
vi.mock('../services/auth.service', () => ({
  authService: {
    login: vi.fn(),
  },
}))

// Mock del Router
vi.mock('next/navigation', () => ({
  useRouter: () => ({
    push: vi.fn(),
  }),
}))

describe('LoginFlow', () => {
  it('renders login form and submits successfully', async () => {
    // Ara sí, com que l'hem definit com a vi.fn() a dalt, tenim mockResolvedValue
    vi.mocked(authService.login).mockResolvedValue({
        message: 'Login successful',
        user: { id: '1', username: 'test', name: 'Test User' }
    })

    render(<LoginFlow />)

    const userInput = screen.getByLabelText(/Email/i)
    // 🔥 FIX: També aquí especifiquem selector input
    const passInput = screen.getByLabelText(/Contrasenya/i, { selector: 'input' })
    const loginBtn = screen.getByRole('button', { name: /Entrar al Cas/i })

    fireEvent.change(userInput, { target: { value: 'detective1@example.com' } })
    fireEvent.change(passInput, { target: { value: 'password123' } })
    fireEvent.click(loginBtn)

    await waitFor(() => {
      expect(authService.login).toHaveBeenCalledWith('detective1@example.com', 'password123')
    })
  })
})
