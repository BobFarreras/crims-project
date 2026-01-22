import { render, screen, fireEvent, waitFor } from '@testing-library/react'
import LoginFlow from '../components/LoginFlow'
import { authService } from '../services/auth.service'
import { vi } from 'vitest'

// 🔥 FIX: Definim el mock explícitament perquè vi.fn() estigui disponible
vi.mock('../services/auth.service', () => ({
  authService: {
    login: vi.fn(),
  },
}))

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
        message: 'Login successful',
        user: { id: '1', username: 'test', name: 'Test User' }
    })

    render(<LoginFlow />)

    const userInput = screen.getByLabelText(/Email/i)
    const passInput = screen.getByLabelText(/Contrasenya/i, { selector: 'input' })
    const loginBtn = screen.getByRole('button', { name: /Entrar al Cas/i })

    fireEvent.change(userInput, { target: { value: 'detective1@example.com' } })
    fireEvent.change(passInput, { target: { value: 'password123' } })
    fireEvent.click(loginBtn)

    await waitFor(() => {
      expect(authService.login).toHaveBeenCalledWith('detective1@example.com', 'password123')
    })
  })

  it('shows error message on failed login', async () => {
    vi.mocked(authService.login).mockRejectedValue(new Error('Credencials incorrectes. Revisa el correu i la contrasenya.'))

    render(<LoginFlow />)

    const userInput = screen.getByLabelText(/Email/i)
    const passInput = screen.getByLabelText(/Contrasenya/i, { selector: 'input' })
    const loginBtn = screen.getByRole('button', { name: /Entrar al Cas/i })

    fireEvent.change(userInput, { target: { value: 'detective1@example.com' } })
    fireEvent.change(passInput, { target: { value: 'wrongpass' } })
    fireEvent.click(loginBtn)

    expect(await screen.findByText(/Credencials incorrectes\. Revisa el correu i la contrasenya\./i)).toBeInTheDocument()
  })
})
