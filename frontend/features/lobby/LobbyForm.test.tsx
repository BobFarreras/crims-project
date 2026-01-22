import { fireEvent, render, screen } from '@testing-library/react'
import LobbyForm from './LobbyForm'

const capabilities = ['DETECTIVE', 'FORENSIC', 'ANALYST', 'INTERROGATOR']

describe('LobbyForm', () => {
  it('renders fields', () => {
    render(<LobbyForm capabilities={capabilities} onSubmit={() => undefined} />)

    expect(screen.getByLabelText(/game code/i)).toBeInTheDocument()
    expect(screen.getByRole('group', { name: /capabilities/i })).toBeInTheDocument()
    expect(screen.getByRole('button', { name: /join/i })).toBeInTheDocument()
  })

  it('shows error when code is empty', () => {
    render(<LobbyForm capabilities={capabilities} onSubmit={() => undefined} />)

    fireEvent.click(screen.getByRole('button', { name: /join/i }))

    expect(screen.getByText(/code is required/i)).toBeInTheDocument()
  })
})
