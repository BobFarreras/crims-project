import { render, screen } from '@testing-library/react'
import ForensicsPanel from './ForensicsPanel'

describe('ForensicsPanel', () => {
  it('renders title and cards', () => {
    render(<ForensicsPanel />)

    expect(screen.getByRole('heading', { name: /forensics/i })).toBeInTheDocument()
    expect(screen.getAllByTestId('forensic-card')).toHaveLength(3)
  })
})
