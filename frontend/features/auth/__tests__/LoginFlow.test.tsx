import { render, screen, fireEvent, waitFor } from '@testing-library/react'
import { describe, it, expect, vi } from 'vitest'
import LoginFlow from './LoginFlow'

vi.stubGlobal('fetch', async () => {
  return {
    ok: true,
    json: async () => ({ ok: true, token: 'mock' }),
  } as Response
})
const pushMock = vi.fn()
vi.mock('next/navigation', () => ({ useRouter: () => ({ push: pushMock }) }))

describe('LoginFlow', () => {
  it('renders login form and submits successfully', async () => {
    render(<LoginFlow />)
    const userInput = screen.getByLabelText(/Username/i)
    const passInput = screen.getByLabelText(/Password/i)
    const loginBtn = screen.getByRole('button', { name: /Login/i })

    fireEvent.change(userInput, { target: { value: 'alice' } })
    fireEvent.change(passInput, { target: { value: 'secret' } })
    fireEvent.click(loginBtn)

    // Expect navigation to be called after login
    await waitFor(() => expect(pushMock).toHaveBeenCalled())
  })
})
