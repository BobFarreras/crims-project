# Frontend - CRIMS de Mitjanit

This is a [Next.js](https://next.js) project bootstrapped with [`create-next-app`](https://nextjs.org/docs/app/api-reference/cli/create-next-app).

## 🚀 Getting Started

First, run the development server:

```bash
pnpm dev
```

Open [http://localhost:3000](http://localhost:3000) with your browser to see the result.

## 🧪 Testing Sentry Configuration

To verify that Sentry is correctly configured, follow these steps:

### Step 1: Check Environment Variables

1. Make sure `.env.local` exists in the project root (NOT in `frontend/` folder)
2. Verify these variables are set:
   - `NEXT_PUBLIC_SENTRY_DSN` - Frontend DSN from Sentry
   - `NEXT_PUBLIC_ENVIRONMENT` - Should be "development" for testing

### Step 2: Run Development Server

```bash
cd frontend
pnpm dev
```

### Step 3: Open Test Page

Open your browser and navigate to:
```
http://localhost:3000/test-sentry
```

### Step 4: Check Configuration Status

On the test page, look at "Estat de Configuración":
- ✅ **DSN Configurat** - Good!
- ✅ **Environment: development** - Good!

If you see ❌ **DSN NO Configurat**:
1. Check that `.env.local` exists in project root
2. Verify `NEXT_PUBLIC_SENTRY_DSN` is set
3. Restart the dev server (`pnpm dev`)

### Step 5: Run Tests

Click on each test button in order:

1. **1. Error Manual (try/catch)**
   - Click the button
   - Check logs: Should say "✅ Error 1: Capturat a Sentry"

2. **2. Error Asíncrono (Promise reject)**
   - Click the button
   - Check logs: Should say "✅ Error 2: Capturat a Sentry"

3. **3. Error amb Context i Tags**
   - Click the button
   - Check logs: Should say "✅ Error 3: Capturat a Sentry"

4. **4. Capturar Missatge**
   - Click the button
   - Check logs: Should say "✅ Mensaje 4: Capturat a Sentry"

5. **5. Error amb Breadcrumbs**
   - Click the button
   - Check logs: Should say "✅ Breadcrumb 5: Afegit" + "✅ Error 5: Capturat a Sentry"

6. **6. Error de Xarxa (fetch)**
   - Click the button
   - Check logs: Should say "✅ Error 6: Capturat a Sentry (HTTP 404)"

### Step 6: Verify in Sentry Dashboard

1. Go to: [https://sentry.io](https://sentry.io)
2. Login
3. Select: **digitaistudios** (organization)
4. Select: **crims-frontend** (project)
5. You should see 6 new errors (one for each test button)

## 🐛 Troubleshooting

### ❌ "DSN NO Configurat" on test page

**Problem:** Environment variables not loaded.

**Solution:**
```bash
# Check if .env.local exists
ls -la ../.env.local

# Check if variables are set
cat ../.env.local | grep NEXT_PUBLIC_SENTRY_DSN
```

Make sure `.env.local` is in the **project root**, not in `frontend/` folder.

### ❌ Errors not appearing in Sentry

**Possible causes:**

1. **Network issues:**
   - Check you have internet connection
   - Try pinging Sentry: `ping o4510557342400512.ingest.de.sentry.io`

2. **Wrong DSN:**
   - Copy the DSN exactly from Sentry (no extra spaces)
   - Verify it's the **frontend DSN**, not the backend one

3. **Environment filters:**
   - In Sentry, check filters are set to "All" environments
   - Not just "production"

4. **Sample rate:**
   - With 10% sample rate, some errors might not be captured
   - Check logs - they should still say "Capturat a Sentry"

### ❌ Browser Console Errors

If you see errors in browser console (F12):

**"Hydration failed" error:**
- This is normal during development
- It's because React re-renders
- It's not related to Sentry

**"DSN is not configured" error:**
- Check `.env.local` file
- Restart dev server

## 📝 Learn More

To learn more about Sentry configuration, see:
- [Sentry Setup Guide](../docs/sentry-setup.md) - Complete Sentry setup documentation
- [Sentry Documentation](https://docs.sentry.io/platforms/javascript/) - Official Sentry docs

## 🎯 Next Steps

Once Sentry is working correctly:

1. **Test Backend:**
   ```bash
   cd backend
   go run ./cmd/server
   ```
   Then visit:
   - http://localhost:8080/api/test-sentry/error1
   - http://localhost:8080/api/test-sentry/error2
   - etc.

2. **Configure Alerts:**
   - Go to Sentry → Settings → Alerts
   - Create alerts for new errors

3. **Deploy to Production:**
   - Change `NEXT_PUBLIC_ENVIRONMENT=production` in `.env.local`
   - Now Sentry will capture real user errors

---

**Built with ❤️ for CRIMS de Mitjanit**
