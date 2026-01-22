import { fireEvent, render, screen } from '@testing-library/react'
import LobbyForm from './LobbyForm'

const roles = ['DETECTIVE', 'FORENSIC', 'ANALYST', 'INTERROGATOR']

describe('LobbyForm', () => {
  it('renders fields', () => {
    render(<LobbyForm roles={roles} onSubmit={() => undefined} />)

    expect(screen.getByLabelText(/game code/i)).toBeInTheDocument()
    expect(screen.getByLabelText(/role/i)).toBeInTheDocument()
    expect(screen.getByRole('button', { name: /join/i })).toBeInTheDocument()
  })

  it('shows error when code is empty', () => {
    render(<LobbyForm roles={roles} onSubmit={() => undefined} />)

    fireEvent.click(screen.getByRole('button', { name: /join/i }))

    expect(screen.getByText(/code is required/i)).toBeInTheDocument()
  })
})
