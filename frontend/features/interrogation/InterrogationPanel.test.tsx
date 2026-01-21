import { render, screen } from '@testing-library/react'
import InterrogationPanel from './InterrogationPanel'

describe('InterrogationPanel', () => {
  it('renders title and main question', () => {
    render(<InterrogationPanel />)

    expect(screen.getByRole('heading', { name: /interrogation/i })).toBeInTheDocument()
    expect(screen.getByText(/where were you/i)).toBeInTheDocument()
  })

  it('renders history entries', () => {
    render(<InterrogationPanel />)

    expect(screen.getAllByTestId('history-item')).toHaveLength(3)
  })
})
