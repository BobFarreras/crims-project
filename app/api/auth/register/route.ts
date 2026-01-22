import { NextRequest, NextResponse } from 'next/server'
import { addUser } from '@/app/api/auth/store'

export async function POST(req: NextRequest) {
  const body = await req.json()
  const { username, password } = body
  if (!username || !password) {
    return NextResponse.json({ ok: false, error: 'missing_credentials' }, { status: 400 })
  }
  const ok = addUser(username, password)
  if (!ok) {
    return NextResponse.json({ ok: false, error: 'user_exists' }, { status: 409 })
  }
  // Simple success response with a token (will be set by login path in typical flow)
  return NextResponse.json({ ok: true, username })
}

export const runtime = 'edge'
