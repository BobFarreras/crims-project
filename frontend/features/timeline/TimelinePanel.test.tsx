import { render, screen } from '@testing-library/react'
import TimelinePanel from './TimelinePanel'

describe('TimelinePanel', () => {
  it('renders title and events', () => {
    render(<TimelinePanel />)

    expect(screen.getByRole('heading', { name: /timeline/i })).toBeInTheDocument()
    expect(screen.getAllByTestId('timeline-item')).toHaveLength(3)
  })
})
