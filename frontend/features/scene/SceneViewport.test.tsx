import { render, screen } from '@testing-library/react'
import SceneViewport from './SceneViewport'

describe('SceneViewport', () => {
  it('renders scene viewport placeholder', () => {
    render(<SceneViewport />)
    expect(screen.getByText(/Scene Viewport/i)).toBeInTheDocument()
  })
})
