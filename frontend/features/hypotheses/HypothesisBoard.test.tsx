import { render, screen } from '@testing-library/react'
import HypothesisBoard from './HypothesisBoard'

describe('HypothesisBoard', () => {
  it('renders title and cards', () => {
    render(<HypothesisBoard />)

    expect(screen.getByRole('heading', { name: /hypotheses/i })).toBeInTheDocument()
    expect(screen.getAllByTestId('hypothesis-card')).toHaveLength(3)
  })
})
