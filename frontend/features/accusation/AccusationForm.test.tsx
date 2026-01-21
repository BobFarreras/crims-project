import { render, screen } from '@testing-library/react'
import AccusationForm from './AccusationForm'

describe('AccusationForm', () => {
  it('renders title and selects', () => {
    render(<AccusationForm />)

    expect(screen.getByRole('heading', { name: /accusation/i })).toBeInTheDocument()
    expect(screen.getAllByRole('combobox')).toHaveLength(3)
  })
})
