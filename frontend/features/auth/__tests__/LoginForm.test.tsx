import { render, screen, fireEvent } from '@testing-library/react'
import LoginForm from '../components/LoginForm'
import { vi } from 'vitest'

describe('LoginForm', () => {
  it('renders username and password fields and login button', () => {
    render(<LoginForm onSubmit={() => {}} />)

    // 🔥 FIX 2: Textos en català
    expect(screen.getByLabelText(/Usuari o Email/i)).toBeInTheDocument()
    expect(screen.getByLabelText(/Contrasenya/i)).toBeInTheDocument()
    expect(screen.getByRole('button', { name: /Entrar al Cas/i })).toBeInTheDocument()
  })

  it('submits values (basic)', () => {
    const onSubmit = vi.fn()
    render(<LoginForm onSubmit={onSubmit} />)

    // 🔥 FIX 3: Emplenem els camps buscant pel text català
    fireEvent.change(screen.getByLabelText(/Usuari o Email/i), { target: { value: 'Alice' } })
    fireEvent.change(screen.getByLabelText(/Contrasenya/i), { target: { value: 'secret' } })
    
    fireEvent.click(screen.getByRole('button', { name: /Entrar al Cas/i }))

    // Comprovem que s'envia l'objecte correcte
    expect(onSubmit).toHaveBeenCalledWith({
      username: 'Alice',
      password: 'secret'
    })
  })
})