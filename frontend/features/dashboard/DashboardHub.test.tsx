import { render, screen } from '@testing-library/react'
import DashboardHub from './DashboardHub'
import { vi } from 'vitest'

vi.mock('next/navigation', () => ({
  useRouter: () => ({
    push: vi.fn(),
    refresh: vi.fn(),
  }),
}))

describe('DashboardHub', () => {
  it('renders mode selector buttons', () => {
    render(<DashboardHub />)

    expect(screen.getByRole('button', { name: /solo/i })).toBeInTheDocument()
    expect(screen.getByRole('button', { name: /duo/i })).toBeInTheDocument()
    expect(screen.getByRole('button', { name: /equip/i })).toBeInTheDocument()
    expect(screen.getByRole('button', { name: /sala/i })).toBeInTheDocument()
  })

  it('renders profile editor', () => {
    render(<DashboardHub />)

    expect(screen.getByLabelText(/nom del jugador/i)).toBeInTheDocument()
  })

  it('renders active cases and ranking', () => {
    render(<DashboardHub />)

    expect(screen.getByText(/casos actius/i)).toBeInTheDocument()
    expect(screen.getAllByRole('button', { name: /continuar/i }).length).toBeGreaterThan(0)
    expect(screen.getByRole('heading', { name: /ranking/i })).toBeInTheDocument()
  })
})
