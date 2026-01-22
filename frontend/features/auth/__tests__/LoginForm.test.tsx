import { render, screen, fireEvent } from '@testing-library/react'
import LoginForm from '../components/LoginForm'
import { vi } from 'vitest'

describe('LoginForm', () => {
  it('renders email and password fields and login button', () => {
    render(<LoginForm onSubmit={() => {}} />)

    expect(screen.getByLabelText(/Email/i)).toBeInTheDocument()
    expect(screen.getByLabelText(/Contrasenya/i, { selector: 'input' })).toBeInTheDocument()
    expect(screen.getByRole('button', { name: /Entrar al Cas/i })).toBeInTheDocument()
  })

  it('submits values (basic)', () => {
    const onSubmit = vi.fn()
    render(<LoginForm onSubmit={onSubmit} />)

    fireEvent.change(screen.getByLabelText(/Email/i), { target: { value: 'alice@example.com' } })
    fireEvent.change(screen.getByLabelText(/Contrasenya/i, { selector: 'input' }), { target: { value: 'secret' } })
    
    fireEvent.click(screen.getByRole('button', { name: /Entrar al Cas/i }))

    expect(onSubmit).toHaveBeenCalledWith({
      email: 'alice@example.com',
      password: 'secret'
    })
  })
})
