#!/bin/bash

# Script per probar Sentry - Frontend i Backend

echo "🧪 Script de Proves de Sentry"
echo "================================"
echo ""

# Check si .env.local existeix
if [ ! -f ".env.local" ]; then
    echo "❌ ERROR: .env.local NO existeix!"
    echo "Copia: cp .env.example .env.local"
    echo "I configura els DSN de Sentry"
    exit 1
fi

echo "✅ .env.local trobat"

# Check si els DSN estan configurats
if grep -q "NEXT_PUBLIC_SENTRY_DSN=https://" .env.local; then
    echo "✅ Frontend DSN configurat"
else
    echo "❌ ERROR: NEXT_PUBLIC_SENTRY_DSN NO configurat a .env.local"
    exit 1
fi

if grep -q "SENTRY_DSN=https://" .env.local; then
    echo "✅ Backend DSN configurat"
else
    echo "❌ ERROR: SENTRY_DSN NO configurat a .env.local"
    exit 1
fi

echo ""
echo "================================"
echo "🚀 Iniciant servidors..."
echo ""

# Iniciar backend en background
echo "📦 Iniciant Backend (Go)..."
cd backend
go run ./cmd/server > ../backend.log 2>&1 &
BACKEND_PID=$!
echo "✅ Backend iniciat (PID: $BACKEND_PID)"
echo "📦 Backend URL: http://localhost:8080"
echo ""

# Esperar que backend estigui a punt
echo "⏳ Esperant que backend estigui a punt..."
sleep 3

# Test debug de backend
echo ""
echo "================================"
echo "🔍 Test 1: Debug de configuració Backend"
echo "================================"
curl -s http://localhost:8080/api/test-sentry/debug | jq '.'
echo ""

# Test d'errors de backend
echo "================================"
echo "🔍 Test 2: Errors Backend"
echo "================================"
echo "Test 2.1: Error manual"
curl -s http://localhost:8080/api/test-sentry/error1 | jq '.'
echo ""

echo "Test 2.2: Error amb context"
curl -s http://localhost:8080/api/test-sentry/error2 | jq '.'
echo ""

echo "Test 2.3: Error amb nivell de severitat"
curl -s http://localhost:8080/api/test-sentry/error3 | jq '.'
echo ""

echo "Test 2.4: Mensaje"
curl -s http://localhost:8080/api/test-sentry/message | jq '.'
echo ""

echo ""
echo "================================"
echo "🌐 Iniciant Frontend (Next.js)..."
echo "================================"
cd ../frontend
pnpm dev > ../frontend.log 2>&1 &
FRONTEND_PID=$!
echo "✅ Frontend iniciat (PID: $FRONTEND_PID)"
echo "🌐 Frontend URL: http://localhost:3000"
echo ""
echo "================================"
echo "📋 Comandes per probar:"
echo "================================"
echo ""
echo "Test Frontend:"
echo "  1. Obre el navegador: http://localhost:3000/test-sentry"
echo "  2. Revisa 'Estat de Configuración' (hauria de ser tot ✅)"
echo "  3. Fes clic en els botons de test (1-6)"
echo "  4. Revisa els logs de sota (haurien de dir 'Capturat a Sentry')"
echo ""
echo "Test Backend (ja executat):"
echo "  ✅ Ja hem fet els tests automàtics (veure sota)"
echo ""
echo "Verificar en Sentry:"
echo "  1. Ves a: https://sentry.io"
echo "  2. Inicia sessió"
echo "  3. Selecciona: digitaistudios (organització)"
echo "  4. Selecciona: crims-frontend o crims-backend"
echo "  5. Hauries de veure els errors generats"
echo ""
echo "================================"
echo "📝 Logs:"
echo "  Backend: tail -f backend.log"
echo "  Frontend: tail -f frontend.log"
echo ""
echo "================================"
echo "🛑 Per aturar: Ctrl+C"
echo "================================"
echo ""

# Esperar que l'usuari premi Ctrl+C
wait $FRONTEND_PID

# Aturar tot en sortir
echo ""
echo "🛑 Aturant servidors..."
kill $BACKEND_PID 2>/dev/null
kill $FRONTEND_PID 2>/dev/null
echo "✅ Servidors aturats"
