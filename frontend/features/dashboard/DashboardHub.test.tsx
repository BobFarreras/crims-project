import { fireEvent, render, screen, waitFor } from '@testing-library/react'
import DashboardHub from './DashboardHub'
import { profileService } from '../profile/services/profile.service'
import { vi } from 'vitest'

vi.mock('next/navigation', () => ({
  useRouter: () => ({
    push: vi.fn(),
    refresh: vi.fn(),
  }),
}))

vi.mock('../profile/services/profile.service', () => ({
  profileService: {
    getProfile: vi.fn(),
    updateProfile: vi.fn(),
  },
}))

describe('DashboardHub', () => {
  it('renders mode selector buttons', () => {
    vi.mocked(profileService.getProfile).mockResolvedValue({ user: { name: 'Detective' } })
    render(<DashboardHub />)

    return screen.findByLabelText(/nom del jugador/i).then(() => {
      expect(screen.getByRole('button', { name: /solo/i })).toBeInTheDocument()
      expect(screen.getByRole('button', { name: /duo/i })).toBeInTheDocument()
      expect(screen.getByRole('button', { name: /equip/i })).toBeInTheDocument()
      expect(screen.getByRole('button', { name: /sala/i })).toBeInTheDocument()
    })
  })

  it('renders profile editor', () => {
    vi.mocked(profileService.getProfile).mockResolvedValue({ user: { name: 'Detective' } })
    render(<DashboardHub />)

    return screen.findByLabelText(/nom del jugador/i).then((input) => {
      expect(input).toBeInTheDocument()
    })
  })

  it('renders active cases and ranking', () => {
    vi.mocked(profileService.getProfile).mockResolvedValue({ user: { name: 'Detective' } })
    render(<DashboardHub />)

    return screen.findByLabelText(/nom del jugador/i).then(() => {
      expect(screen.getByText(/casos actius/i)).toBeInTheDocument()
      expect(screen.getAllByRole('button', { name: /continuar/i }).length).toBeGreaterThan(0)
      expect(screen.getByRole('heading', { name: /ranking/i })).toBeInTheDocument()
    })
  })

  it('calls updateProfile when saving name', async () => {
    vi.mocked(profileService.getProfile).mockResolvedValue({ user: { name: 'Detective' } })
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
})
