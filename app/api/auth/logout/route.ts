import { NextResponse } from 'next/server'

export async function POST() {
  const res = NextResponse.json({ ok: true })
  res.headers.set('Set-Cookie', `token=; HttpOnly; Path=/; Max-Age=0; Secure; SameSite=Lax`)
  return res
}

export const runtime = 'edge'
