export function getCookie(name: string): string | null {
  if (typeof document === 'undefined') return null
  const v = `; ${document.cookie}`
  const parts = v.split(`; ${name}=`)
  if (parts.length === 2) return parts.pop()!.split(';').shift() || null
  return null
}
