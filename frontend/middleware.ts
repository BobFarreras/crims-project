import { NextResponse } from 'next/server';
import type { NextRequest } from 'next/server';

// Rutes que requereixen estar loguejat (zona de joc)
const PROTECTED_ROUTES = ['/game'];

// Rutes on NO pots estar si ja tens sessió (Login/Register)
const AUTH_ROUTES = ['/login', '/register', '/']; // Afegeixo '/' per redirigir al dashboard si ja estàs loguejat

export function middleware(request: NextRequest) {
  // 🔥 LLEGIM LA COOKIE HTTPONLY (El servidor sí que la pot veure)
  const token = request.cookies.get('auth_token')?.value;
  const { pathname } = request.nextUrl;

  console.log(`📡 Middleware revisant: ${pathname} | Token present: ${!!token}`);

  // 1. PROTECCIÓ DE RUTES PRIVADES
  // Si intentes entrar a /game/... i no tens la cookie -> FORA
  if (PROTECTED_ROUTES.some(route => pathname.startsWith(route))) {
    if (!token) {
      console.log("⛔ Accés denegat. Redirigint a Login.");
      return NextResponse.redirect(new URL('/login', request.url));
    }
  }

  // 2. REDIRECCIÓ D'USUARIS LOGUEJATS
  // Si ja tens cookie i vas a /login -> Cap al Dashboard
  if (AUTH_ROUTES.some(route => pathname === route)) {
    if (token) {
      console.log("✅ Usuari ja loguejat. Redirigint a Dashboard.");
      return NextResponse.redirect(new URL('/game/dashboard', request.url));
    }
  }

  return NextResponse.next();
}

// Configurem a quines rutes s'executa
export const config = {
  matcher: [
    '/game/:path*', // Totes les rutes de joc
    '/login', 
    '/register',
    '/'             // La Home també la vigilem
  ],
};