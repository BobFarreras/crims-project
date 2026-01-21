import { render, screen } from '@testing-library/react'
import CaseDashboard from './CaseDashboard'

describe('CaseDashboard', () => {
  it('renders title and summary cards', () => {
    render(<CaseDashboard />)

    expect(screen.getByRole('heading', { name: /case dashboard/i })).toBeInTheDocument()
    expect(screen.getAllByTestId('summary-card')).toHaveLength(3)
  })

  it('renders mock players', () => {
    render(<CaseDashboard />)

    expect(screen.getAllByTestId('player-row')).toHaveLength(3)
  })
})
