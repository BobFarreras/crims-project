import { render, screen, fireEvent, waitFor } from '@testing-library/react'
import { describe, it, expect, vi } from 'vitest'
import LoginFlow from './LoginFlow'

vi.mock('../../lib/infra/auth-client', () => ({
  login: vi.fn(() => Promise.resolve({ token: 'tok', userId: 'u1' }))
}))
vi.mock('next/navigation', () => ({ useRouter: () => ({ push: vi.fn() }) }))

describe('LoginFlow', () => {
  it('renders login form and submits successfully', async () => {
    render(<LoginFlow />)
    const userInput = screen.getByLabelText(/Username/i)
    const passInput = screen.getByLabelText(/Password/i)
    const loginBtn = screen.getByRole('button', { name: /Login/i })

    fireEvent.change(userInput, { target: { value: 'alice' } })
    fireEvent.change(passInput, { target: { value: 'secret' } })
    fireEvent.click(loginBtn)

    // Expect login to be called; navigation is mocked
    await waitFor(() => expect(require('../../lib/infra/auth-client').login).toHaveBeenCalled())
  })
})
