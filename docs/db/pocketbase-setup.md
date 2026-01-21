# PocketBase Setup (REST)

## Prerequisits
- `.env.local` amb:
  - `PB_URL`
  - `PB_SUPERUSER_EMAIL` i `PB_SUPERUSER_PASSWORD`
  - o `PB_ADMIN_TOKEN` (recomanat)

## Executar (Windows PowerShell)
```powershell
./scripts/pb-setup.ps1
```

## Import amb reemplaç (recomanat)
```powershell
./scripts/pb-import.ps1
```

Debug (comprova que llegeix variables):
```powershell
./scripts/pb-setup.ps1 -Debug
```

## Resultat
- Crea/actualitza totes les col·leccions.
- Configura relacions i camps segons `docs/db/pocketbase-schema.md`.

## Notes
- El script no imprimeix secrets.
- Si una col·leccio ja existeix, la reutilitza.
