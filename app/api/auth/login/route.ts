import { NextRequest, NextResponse } from 'next/server'
import { validateUser } from '@/app/api/auth/store'
import crypto from 'crypto'

function base64url(input: Buffer | string) {
  return Buffer.from(input).toString('base64').replace(/=+$/g, '').replace(/\+/g, '-').replace(/\//g, '_')
}

function generateToken(username: string, secret: string): string {
  const header = { alg: 'HS256', typ: 'JWT' }
  const payload = { sub: username, iat: Math.floor(Date.now() / 1000), exp: Math.floor(Date.now() / 1000) + 3600 }
  const headerB64 = base64url(Buffer.from(JSON.stringify(header)))
  const payloadB64 = base64url(Buffer.from(JSON.stringify(payload)))
  const unsigned = `${headerB64}.${payloadB64}`
  const signature = crypto.createHmac('sha256', secret).update(unsigned).digest()
  const sigB64 = base64url(signature)
  return `${unsigned}.${sigB64}`
}

export async function POST(req: NextRequest) {
  const body = await req.json()
  const { username, password } = body
  const secret = process.env.AUTH_SECRET || 'dev-secret'
  if (!username || !password) {
    return NextResponse.json({ ok: false, error: 'missing_credentials' }, { status: 400 })
  }
  const valid = validateUser(username, password)
  if (!valid) {
    return NextResponse.json({ ok: false, error: 'invalid_credentials' }, { status: 401 })
  }
  const token = generateToken(username, secret)
  const res = NextResponse.json({ ok: true, token, username })
  res.headers.append('Set-Cookie', `token=${token}; HttpOnly; Path=/; Max-Age=3600; Secure; SameSite=Lax`)
  return res
}

export const runtime = 'edge'
