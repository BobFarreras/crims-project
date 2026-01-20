# DEPLOYMENT GUIDE 🚀

## 1. Overview

### Production Infrastructure
```
┌─────────────────────────────────────────────────────────┐
│                      CLOUD INFRA                         │
├─────────────────────────────────────────────────────────┤
│  Frontend: Vercel                                        │
│  → URL: https://www.crimsdemitjanit.com                  │
│  → Framework: Next.js 15 (SSG + ISR)                     │
│  → Edge Functions: API Proxy (opcional)                  │
├─────────────────────────────────────────────────────────┤
│  Backend API: VPS (Ubuntu)                               │
│  → URL: https://api.digitaistudios.com                   │
│  → Stack: Go + Chi Router                                │
│  → Deployment: Docker Container                          │
├─────────────────────────────────────────────────────────┤
│  Database: VPS (Same as API)                             │
│  → URL: https://sspb.digitaistudios.com                  │
│  → Stack: PocketBase (BaaS)                              │
│  → Deployment: Docker Container                         │
└─────────────────────────────────────────────────────────┘
```

### Development Infrastructure
```
Local Machine
├─ Frontend: localhost:3000 (Next.js dev server)
├─ Backend: localhost:8080 (Go dev server)
└─ DB: localhost:8090 (PocketBase)
```

---

## 2. Prerequisites

### Required Tools
- **Git** (version control)
- **Node.js 20+** (frontend)
- **Go 1.25.6** (backend)
- **Docker & Docker Compose** (VPS deployment)
- **pnpm 9** (package manager - recommended)

### Required Accounts
- **Vercel** (Frontend hosting)
- **Domain provider** (DNS management)
- **VPS provider** (Backend + DB)
- **GitHub** (CI/CD + repository)
- **Sentry** (Error tracking - opcional però recomanat)

---

## 3. Environment Variables

### Root `.env.example`
```bash
# ===============================
# FRONTEND ENVIRONMENT
# ===============================
NEXT_PUBLIC_API_URL=https://api.digitaistudios.com
NEXT_PUBLIC_POCKETBASE_URL=https://sspb.digitaistudios.com
NEXT_PUBLIC_SENTRY_DSN=https://xxx@sentry.io/xxx

# ===============================
# BACKEND ENVIRONMENT
# ===============================
PORT=8080
ENVIRONMENT=production

# PocketBase Connection
PB_URL=https://sspb.digitaistudios.com
PB_ADMIN_EMAIL=admin@example.com
PB_ADMIN_PASSWORD=change_this_password

# OpenAI (Optional - AI Dungeon Master)
OPENAI_API_KEY=sk-...

# JWT Authentication
JWT_SECRET=your_super_secret_key_change_this
JWT_EXPIRATION=24h

# Sentry
SENTRY_DSN=https://xxx@sentry.io/xxx
SENTRY_ENVIRONMENT=production

# Rate Limiting
RATE_LIMIT_REQUESTS=100
RATE_LIMIT_WINDOW=1m

# ===============================
# POCKETBASE ENVIRONMENT
# ===============================
# Set these in PocketBase Admin Panel or docker-compose
# PB_ADMIN_EMAIL
# PB_ADMIN_PASSWORD
```

---

## 4. Frontend Deployment (Vercel)

### 4.1 Initial Setup

1. **Install Vercel CLI:**
```bash
pnpm add -g vercel
```

2. **Login to Vercel:**
```bash
vercel login
```

3. **Configure project:**
```bash
cd frontend
vercel
```

Follow the wizard:
- Project name: `crims-frontend`
- Framework preset: Next.js
- Root directory: `./frontend`

### 4.2 Environment Variables in Vercel

Go to: **Vercel Dashboard** → **Settings** → **Environment Variables**

Add these variables:
```
NEXT_PUBLIC_API_URL = https://api.digitaistudios.com
NEXT_PUBLIC_POCKETBASE_URL = https://sspb.digitaistudios.com
NEXT_PUBLIC_SENTRY_DSN = https://xxx@sentry.io/xxx
```

**Important:** Select **Production**, **Preview**, and **Development** environments.

### 4.3 Custom Domain

1. Go to **Settings** → **Domains**
2. Add: `www.crimsdemitjanit.com`
3. Vercel will provide DNS records. Add these to your domain provider.

### 4.4 Build Configuration

The project uses Next.js standard build. Vercel will automatically:
- Run `pnpm install`
- Run `pnpm build`
- Start `pnpm start`

**Build logs:** Check at **Vercel Dashboard** → **Deployments**

### 4.5 Troubleshooting Vercel

**Build fails?**
- Check `next.config.ts` is correct
- Verify environment variables are set
- Check `tsconfig.json` paths

**Deploy slow?**
- Enable Edge Functions for API routes
- Use `next/image` for images (optimization)
- Enable ISR (Incremental Static Regeneration)

---

## 5. Backend Deployment (VPS + Docker)

### 5.1 VPS Initial Setup

Connect to your VPS via SSH:
```bash
ssh user@your-vps-ip
```

Update system:
```bash
sudo apt update && sudo apt upgrade -y
```

Install Docker:
```bash
curl -fsSL https://get.docker.com -o get-docker.sh
sudo sh get-docker.sh
sudo usermod -aG docker $USER
```

Install Docker Compose:
```bash
sudo curl -L "https://github.com/docker/compose/releases/latest/download/docker-compose-$(uname -s)-$(uname -m)" -o /usr/local/bin/docker-compose
sudo chmod +x /usr/local/bin/docker-compose
```

### 5.2 Docker Configuration

Create `docker-compose.yml` in VPS:
```yaml
version: '3.8'

services:
  # ===============================
  # POCKETBASE DATABASE
  # ===============================
  pocketbase:
    image: ghcr.io/muchobien/pocketbase:latest
    container_name: crims-pocketbase
    ports:
      - "8090:8090"
    volumes:
      - pb_data:/pb_data
      - pb_public:/pb_public
    restart: unless-stopped
    environment:
      - TZ=Europe/Madrid

  # ===============================
  # BACKEND API (GO)
  # ===============================
  backend:
    build:
      context: ./backend
      dockerfile: Dockerfile
    container_name: crims-backend
    ports:
      - "8080:8080"
    depends_on:
      - pocketbase
    restart: unless-stopped
    environment:
      - PORT=8080
      - ENVIRONMENT=production
      - PB_URL=http://pocketbase:8090
      - PB_ADMIN_EMAIL=${PB_ADMIN_EMAIL}
      - PB_ADMIN_PASSWORD=${PB_ADMIN_PASSWORD}
      - JWT_SECRET=${JWT_SECRET}
      - SENTRY_DSN=${SENTRY_DSN}
      - SENTRY_ENVIRONMENT=production

volumes:
  pb_data:
  pb_public:
```

### 5.3 Backend Dockerfile

Create `backend/Dockerfile`:
```dockerfile
# Multi-stage build for smaller image
FROM golang:1.25.6-alpine AS builder

WORKDIR /app

# Install dependencies
COPY go.mod go.sum ./
RUN go mod download

# Copy source code
COPY . .

# Build binary
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main ./cmd/server

# Final stage
FROM alpine:latest

WORKDIR /root/

# Install ca-certificates for HTTPS
RUN apk --no-cache add ca-certificates

COPY --from=builder /app/main .

# Expose port
EXPOSE 8080

# Run binary
CMD ["./main"]
```

### 5.4 Deploy Backend to VPS

**Option A: Git Clone (Simple)**
```bash
# On VPS
cd /opt
git clone https://github.com/your-repo/crims-project.git
cd crims-project
```

**Option B: GitHub Actions (CI/CD)**
See section 7.

**Start services:**
```bash
docker-compose up -d --build
```

**Check logs:**
```bash
docker-compose logs -f backend
docker-compose logs -f pocketbase
```

---

## 6. PocketBase Configuration

### 6.1 Initial Admin Setup

1. Access PocketBase Admin: `https://sspb.digitaistudios.com/_/`
2. Create admin account
3. Update `.env` with credentials

### 6.2 Import Schema

Run the migration script (create one in `/backend/migrations`):
```bash
docker exec -it crims-pocketbase pb_migrate up
```

Or import JSON schema via Admin Panel manually.

### 6.3 API Rules Configuration

**Critical security settings:**

| Collection | Rules |
|-----------|-------|
| `games` | `id = @request.auth.id` (players can only see their games) |
| `players` | `id = @request.auth.id` or `game_id = @request.data.game_id` |
| `clues` | `public read` (but filter by discovery state) |
| `hypotheses` | `game_id = @request.data.game_id` |

---

## 7. CI/CD Pipeline (GitHub Actions)

### 7.1 Existing CI Pipeline

The project already has a basic CI pipeline at `.github/workflows/ci.yml`:
- Runs tests on push/PR
- Lints frontend
- Builds backend

### 7.2 CD Pipeline (Auto-deploy)

Create `.github/workflows/cd.yml`:
```yaml
name: Crims CD Pipeline

on:
  push:
    branches: [ "main" ]
  workflow_dispatch:

jobs:
  deploy-frontend:
    runs-on: ubuntu-latest
    steps:
    - uses: actions/checkout@v4
    - uses: amondnet/vercel-action@v25
      with:
        vercel-token: ${{ secrets.VERCEL_TOKEN }}
        vercel-org-id: ${{ secrets.VERCEL_ORG_ID }}
        vercel-project-id: ${{ secrets.VERCEL_PROJECT_ID }}
        working-directory: ./frontend
        vercel-args: '--prod'

  deploy-backend:
    runs-on: ubuntu-latest
    steps:
    - uses: actions/checkout@v4
    - name: Deploy to VPS
      uses: appleboy/ssh-action@master
      with:
        host: ${{ secrets.VPS_HOST }}
        username: ${{ secrets.VPS_USER }}
        key: ${{ secrets.VPS_SSH_KEY }}
        script: |
          cd /opt/crims-project
          git pull origin main
          docker-compose up -d --build
          docker-compose exec -T backend ./main
```

### 7.3 Required GitHub Secrets

Add these to **Repository** → **Settings** → **Secrets and variables** → **Actions**:

```
VERCEL_TOKEN
VERCEL_ORG_ID
VERCEL_PROJECT_ID
VPS_HOST
VPS_USER
VPS_SSH_KEY
```

---

## 8. Monitoring & Logging

### 8.1 Sentry Integration (Error Tracking)

**Frontend Setup:**
```bash
cd frontend
pnpm add @sentry/nextjs
pnpm sentry-wizard -i nextjs
```

This will create:
- `sentry.client.config.ts`
- `sentry.server.config.ts`
- `sentry.edge.config.ts`
- `.sentryclirc`

**Backend Setup:**
```bash
cd backend
go get github.com/getsentry/sentry-go
```

Add to `main.go`:
```go
import "github.com/getsentry/sentry-go"

func main() {
    err := sentry.Init(sentry.ClientOptions{
        Dsn:        os.Getenv("SENTRY_DSN"),
        Environment: os.Getenv("SENTRY_ENVIRONMENT"),
    })
    if err != nil {
        log.Fatalf("Sentry init failed: %v", err)
    }
    defer sentry.Flush(2 * time.Second)
}
```

### 8.2 Application Logging (Go)

The project uses `slog` (structured logging). Logs are currently stdout.

**For production:**
- Use Docker logs: `docker-compose logs -f backend`
- Or forward to centralized log service (Loki, ELK)

### 8.3 Health Checks

Add health endpoint to backend (`/health`):
```go
func HealthCheck(w http.ResponseWriter, r *http.Request) {
    w.WriteHeader(http.StatusOK)
    w.Write([]byte("OK"))
}
```

Monitor with UptimeRobot or similar.

---

## 9. SSL/HTTPS Configuration

### 9.1 Vercel (Frontend)
Automatically handles HTTPS. No action needed.

### 9.2 VPS (Backend + PocketBase)

**Option A: Nginx Reverse Proxy (Recommended)**
```nginx
server {
    listen 443 ssl http2;
    server_name api.digitaistudios.com;

    ssl_certificate /etc/letsencrypt/live/api.digitaistudios.com/fullchain.pem;
    ssl_certificate_key /etc/letsencrypt/live/api.digitaistudios.com/privkey.pem;

    location / {
        proxy_pass http://localhost:8080;
        proxy_set_header Host $host;
        proxy_set_header X-Real-IP $remote_addr;
    }
}

server {
    listen 443 ssl http2;
    server_name sspb.digitaistudios.com;

    ssl_certificate /etc/letsencrypt/live/sspb.digitaistudios.com/fullchain.pem;
    ssl_certificate_key /etc/letsencrypt/live/sspb.digitaistudios.com/privkey.pem;

    location / {
        proxy_pass http://localhost:8090;
        proxy_set_header Host $host;
        proxy_set_header X-Real-IP $remote_addr;
    }
}
```

**Install Certbot:**
```bash
sudo apt install certbot python3-certbot-nginx
sudo certbot --nginx -d api.digitaistudios.com -d sspb.digitaistudios.com
```

---

## 10. Backup & Recovery

### 10.1 PocketBase Data Backup

**Automated backup script** (`/opt/backup-pocketbase.sh`):
```bash
#!/bin/bash
DATE=$(date +%Y%m%d_%H%M%S)
BACKUP_DIR="/opt/backups/pocketbase"

mkdir -p $BACKUP_DIR
docker exec crims-pocketbase tar czf - /pb_data | gzip > $BACKUP_DIR/pb_data_$DATE.tar.gz

# Keep last 30 days
find $BACKUP_DIR -name "pb_data_*.tar.gz" -mtime +30 -delete
```

**Add to crontab:**
```bash
crontab -e
# Run daily at 3 AM
0 3 * * * /opt/backup-pocketbase.sh
```

### 10.2 Restore from Backup

```bash
# Stop services
docker-compose down

# Restore data
docker run --rm -v pb_data:/data -v /opt/backups:/backup alpine tar xzf /backup/pb_data_20240120_030000.tar.gz -C /

# Start services
docker-compose up -d
```

---

## 11. Security Checklist

### 11.1 Production Security

- [ ] Change all default passwords
- [ ] Enable firewall (UFW): `sudo ufw allow 80,443,22`
- [ ] Disable root SSH login: `PermitRootLogin no` in `/etc/ssh/sshd_config`
- [ ] Use SSH keys only (no password auth)
- [ ] Update system regularly: `sudo apt update && sudo apt upgrade`
- [ ] Configure rate limiting on API
- [ ] Enable CORS only for trusted domains
- [ ] Rotate secrets periodically
- [ ] Monitor for vulnerabilities (Snyk, Dependabot)

### 11.2 Environment Variables Security

- [ ] Never commit `.env` files
- [ ] Use `.env.example` as template
- [ ] Use strong passwords (32+ chars)
- [ ] Different secrets per environment (dev/staging/prod)
- [ ] Use GitHub Secrets for CI/CD

---

## 12. Rollback Procedure

### 12.1 Frontend Rollback (Vercel)

1. Go to **Vercel Dashboard** → **Deployments**
2. Find previous successful deployment
3. Click "Promote to Production"

### 12.2 Backend Rollback

**Option A: Git Revert**
```bash
# On VPS
cd /opt/crims-project
git revert HEAD
docker-compose up -d --build
```

**Option B: Previous Tag**
```bash
git checkout tags/v1.0.0
docker-compose up -d --build
```

---

## 13. Performance Optimization

### 13.1 Frontend (Next.js)

- Enable ISR for static pages
- Use `next/image` for images
- Minimize bundle size (analyze with `@next/bundle-analyzer`)
- Enable compression (gzip/brotli)
- Use CDN for assets

### 13.2 Backend (Go)

- Use connection pooling to PocketBase
- Implement caching (Redis) for frequent queries
- Optimize database queries
- Use profiling: `go tool pprof`

---

## 14. Troubleshooting

### 14.1 Common Issues

**Issue: Frontend cannot connect to backend**
- Check CORS configuration
- Verify environment variables
- Check VPS firewall rules

**Issue: PocketBase connection refused**
- Verify Docker containers are running: `docker ps`
- Check logs: `docker-compose logs pocketbase`
- Verify ports are exposed correctly

**Issue: Build fails in CI/CD**
- Check GitHub Actions logs
- Verify all environment variables are set
- Check Go module version compatibility

**Issue: High memory usage**
- Profile application: `go tool pprof`
- Check for memory leaks
- Optimize database queries

---

## 15. Maintenance Schedule

| Task | Frequency | Responsibility |
|------|-----------|----------------|
| System updates | Monthly | DevOps |
| Dependency updates | Weekly | Developer |
| Security scans | Weekly | Security |
| Backup verification | Monthly | DevOps |
| Log review | Weekly | Developer |
| Performance audit | Quarterly | Team |

---

## 16. Quick Reference

### Useful Commands

```bash
# Frontend
cd frontend && pnpm dev          # Start dev server
cd frontend && pnpm build        # Build for production
cd frontend && pnpm lint         # Run linter
cd frontend && pnpm test         # Run tests

# Backend
cd backend && go run ./cmd/server  # Start dev server
cd backend && go test ./...       # Run tests
cd backend && go build ./cmd/server # Build binary

# Docker
docker-compose up -d              # Start all services
docker-compose down               # Stop all services
docker-compose logs -f [service]  # View logs
docker-compose exec [service] sh  # Access container shell

# VPS
ssh user@vps-ip                   # Connect to VPS
docker ps                         # List running containers
docker stats                       # Monitor container resources
```

### Important URLs

- Frontend: https://www.crimsdemitjanit.com
- Backend API: https://api.digitaistudios.com
- PocketBase Admin: https://sspb.digitaistudios.com/_/
- Sentry: https://sentry.io (project dashboard)
- Vercel: https://vercel.com (project dashboard)

---

## 17. Support & Contact

For deployment issues:
- Check GitHub Issues: https://github.com/your-repo/crims-project/issues
- Email: dev@digitaistudios.com
- Slack: #crims-deployment channel
