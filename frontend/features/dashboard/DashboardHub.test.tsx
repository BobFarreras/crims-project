import { fireEvent, render, screen, waitFor } from '@testing-library/react'
import DashboardHub from './DashboardHub'
import { profileService } from '../profile/services/profile.service'
import { lobbyService } from '../lobby/services/lobby.service'
import { vi } from 'vitest'

const pushMock = vi.fn()

vi.mock('next/navigation', () => ({
  useRouter: () => ({
    push: pushMock,
    refresh: vi.fn(),
  }),
}))

vi.mock('../profile/services/profile.service', () => ({
  profileService: {
    getProfile: vi.fn(),
    updateProfile: vi.fn(),
  },
}))

vi.mock('../lobby/services/lobby.service', () => ({
  lobbyService: {
    createLobby: vi.fn(),
    joinLobby: vi.fn(),
  },
}))

describe('DashboardHub', () => {
  it('renders mode selector buttons', () => {
    vi.mocked(profileService.getProfile).mockResolvedValue({ user: { id: 'user-1', name: 'Detective' } })
    render(<DashboardHub />)

    return screen.findByLabelText(/nom del jugador/i).then(() => {
      expect(screen.getByRole('button', { name: /solo/i })).toBeInTheDocument()
      expect(screen.getByRole('button', { name: /duo/i })).toBeInTheDocument()
      expect(screen.getByRole('button', { name: /equip/i })).toBeInTheDocument()
      expect(screen.getByRole('button', { name: /sala/i })).toBeInTheDocument()
    })
  })

  it('renders profile editor', () => {
    vi.mocked(profileService.getProfile).mockResolvedValue({ user: { id: 'user-1', name: 'Detective' } })
    render(<DashboardHub />)

    return screen.findByLabelText(/nom del jugador/i).then((input) => {
      expect(input).toBeInTheDocument()
    })
  })

  it('renders active cases and ranking', () => {
    vi.mocked(profileService.getProfile).mockResolvedValue({ user: { id: 'user-1', name: 'Detective' } })
    render(<DashboardHub />)

    return screen.findByLabelText(/nom del jugador/i).then(() => {
      expect(screen.getByText(/casos actius/i)).toBeInTheDocument()
      expect(screen.getAllByRole('button', { name: /continuar/i }).length).toBeGreaterThan(0)
      expect(screen.getByRole('heading', { name: /ranking/i })).toBeInTheDocument()
    })
  })

  it('calls updateProfile when saving name', async () => {
    vi.mocked(profileService.getProfile).mockResolvedValue({ user: { id: 'user-1', name: 'Detective' } })
    vi.mocked(profileService.updateProfile).mockResolvedValue({ user: { name: 'Nova' } })

    render(<DashboardHub />)

    const nameInput = await screen.findByLabelText(/nom del jugador/i)
    const saveButton = screen.getByRole('button', { name: /guardar/i })

    fireEvent.change(nameInput, { target: { value: 'Nova' } })
    fireEvent.click(saveButton)

    await waitFor(() => {
      expect(profileService.updateProfile).toHaveBeenCalledWith('Nova')
    })
  })

  it('calls lobby services when creating and joining', async () => {
    vi.mocked(profileService.getProfile).mockResolvedValue({ user: { id: 'user-1', name: 'Detective' } })
    vi.mocked(lobbyService.createLobby).mockResolvedValue({
      game: { id: 'game-1', code: 'ABCD', state: 'LOBBY', seed: 'seed' },
      player: { id: 'player-1', gameId: 'game-1', userId: 'user-1', capabilities: [], status: 'ONLINE', isHost: true },
    })
    vi.mocked(lobbyService.joinLobby).mockResolvedValue({
      id: 'player-2',
      gameId: 'game-1',
      userId: 'user-1',
      capabilities: [],
      status: 'ONLINE',
      isHost: false,
    })

    render(<DashboardHub />)

    await screen.findByLabelText(/nom del jugador/i)

    fireEvent.click(screen.getByRole('button', { name: /solo/i }))

    await waitFor(() => {
      expect(lobbyService.createLobby).toHaveBeenCalledWith('user-1', ['DETECTIVE', 'FORENSIC', 'ANALYST', 'INTERROGATOR'])
      expect(pushMock).toHaveBeenCalledWith('/game')
    })

    fireEvent.change(screen.getByPlaceholderText(/codi sala/i), { target: { value: 'WXYZ' } })
    fireEvent.click(screen.getByRole('button', { name: /unir-se/i }))

    await waitFor(() => {
      expect(lobbyService.joinLobby).toHaveBeenCalledWith('WXYZ', 'user-1', [])
      expect(pushMock).toHaveBeenCalledWith('/lobby/WXYZ')
    })
  })
})
