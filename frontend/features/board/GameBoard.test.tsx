import { render, screen } from '@testing-library/react'
import GameBoard from './GameBoard'

describe('GameBoard', () => {
  it('renders board layout', () => {
    render(<GameBoard />)

    expect(screen.getByRole('heading', { name: /case board/i })).toBeInTheDocument()
    expect(screen.getByTestId('board-area')).toBeInTheDocument()
  })

  it('shows mock nodes', () => {
    render(<GameBoard />)

    expect(screen.getAllByTestId('board-node')).toHaveLength(3)
  })
})
