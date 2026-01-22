import { render, screen } from '@testing-library/react'
import CaseDashboard from './CaseDashboard'
import { vi } from 'vitest'

// 🔥 FIX 1: Simulem el router de Next.js
vi.mock('next/navigation', () => ({
  useRouter: () => ({
    push: vi.fn(),
    refresh: vi.fn(),
  }),
}))

describe('CaseDashboard', () => {
  it('renders title and summary cards', () => {
    render(<CaseDashboard />)
    
    expect(screen.getByText(/Case Dashboard/i)).toBeInTheDocument()
    // Comprovem que hi ha 3 targetes de resum
    const summaryCards = screen.getAllByTestId('summary-card')
    expect(summaryCards).toHaveLength(3)
  })

  it('renders mock players', () => {
    render(<CaseDashboard />)
    
    expect(screen.getByText(/Active Players/i)).toBeInTheDocument()
    // Comprovem que hi ha 3 jugadors
    const playerRows = screen.getAllByTestId('player-row')
    expect(playerRows).toHaveLength(3)
  })
})