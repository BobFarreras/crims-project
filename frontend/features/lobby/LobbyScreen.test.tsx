import { render, screen } from '@testing-library/react'
import { vi } from 'vitest'
import LobbyScreen from './LobbyScreen'
import { lobbyService } from './services/lobby.service'

vi.mock('./services/lobby.service', () => ({
  lobbyService: {
    getGameByCode: vi.fn(),
  },
}))

describe('LobbyScreen', () => {
  it('renders lobby code and status', async () => {
    vi.mocked(lobbyService.getGameByCode).mockResolvedValue({
      id: 'game-1',
      code: 'ABCD',
      state: 'LOBBY',
      seed: 'seed',
    })

    render(<LobbyScreen code="ABCD" />)

    expect(await screen.findByText(/codi sala/i)).toBeInTheDocument()
    expect(screen.getByText(/ABCD/i)).toBeInTheDocument()
    expect(screen.getByText(/LOBBY/i)).toBeInTheDocument()
  })

  it('renders error when lobby not found', async () => {
    vi.mocked(lobbyService.getGameByCode).mockRejectedValue(new Error('not found'))

    render(<LobbyScreen code="ZZZZ" />)

    expect(await screen.findByText(/no hem trobat aquesta sala/i)).toBeInTheDocument()
  })
})
