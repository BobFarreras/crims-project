param(
  [string]$EnvFile = "",
  [switch]$Debug
)

$ErrorActionPreference = "Stop"

if (-not $EnvFile) {
  $EnvFile = Join-Path $PSScriptRoot "..\.env.local"
}

if (Test-Path $EnvFile) {
  Get-Content $EnvFile | ForEach-Object {
    if ($_ -match '^[\s]*([A-Z0-9_]+)[\s]*=[\s]*(.*)[\s]*$') {
      $name = $matches[1]
      $value = $matches[2].Trim('"')
      if (-not (Get-Item -Path "Env:$name" -ErrorAction SilentlyContinue)) {
        Set-Item -Path "Env:$name" -Value $value
      }
    }
  }
}

if (-not $env:PB_URL) { throw "PB_URL not set" }
$pbEmail = $env:PB_SUPERUSER_EMAIL
$pbPassword = $env:PB_SUPERUSER_PASSWORD
$pbToken = $env:PB_ADMIN_TOKEN
if (-not $pbEmail) { $pbEmail = $env:PB_EMAIL }
if (-not $pbPassword) { $pbPassword = $env:PB_PASSWORD }

if ($Debug) {
  Write-Host "PB_URL set: " -NoNewline; Write-Host ([bool]$env:PB_URL)
  Write-Host "PB_SUPERUSER_EMAIL set: " -NoNewline; Write-Host ([bool]$env:PB_SUPERUSER_EMAIL)
  Write-Host "PB_SUPERUSER_PASSWORD set: " -NoNewline; Write-Host ([bool]$env:PB_SUPERUSER_PASSWORD)
  Write-Host "PB_EMAIL set: " -NoNewline; Write-Host ([bool]$env:PB_EMAIL)
  Write-Host "PB_PASSWORD set: " -NoNewline; Write-Host ([bool]$env:PB_PASSWORD)
  Write-Host "PB_ADMIN_TOKEN set: " -NoNewline; Write-Host ([bool]$env:PB_ADMIN_TOKEN)
}

if (-not $pbToken) {
  if (-not $pbEmail) { throw "PB_SUPERUSER_EMAIL/PB_EMAIL not set" }
  if (-not $pbPassword) { throw "PB_SUPERUSER_PASSWORD/PB_PASSWORD not set" }
}

$pbUri = [System.Uri]$env:PB_URL
$baseUrl = "$($pbUri.Scheme)://$($pbUri.Host)"
if (-not $pbUri.IsDefaultPort) {
  $baseUrl = "${baseUrl}:$($pbUri.Port)"
}

$headers = @{}
if ($pbToken) {
  $headers = @{ Authorization = "Bearer $pbToken" }
} else {
  $authBody = @{ identity = $pbEmail; password = $pbPassword } | ConvertTo-Json

  function Invoke-Auth($endpoint) {
    return Invoke-RestMethod -Method Post -Uri "$baseUrl$endpoint" -ContentType "application/json" -Body $authBody
  }

  try {
    $auth = Invoke-Auth "/api/collections/_superusers/auth-with-password"
  } catch {
    if ($_.Exception.Response.StatusCode.value__ -eq 404) {
      $auth = Invoke-Auth "/api/admins/auth-with-password"
    } else {
      throw
    }
  }
  $headers = @{ Authorization = "Bearer $($auth.token)" }
}

function Get-Collections {
  return (Invoke-RestMethod -Method Get -Uri "$baseUrl/api/collections" -Headers $headers).items
}

function Ensure-Collection($name) {
  $existing = (Get-Collections) | Where-Object { $_.name -eq $name }
  if ($existing) { return $existing.id }

  $body = @{ name = $name; type = "base"; schema = @(); indexes = @() } | ConvertTo-Json
  $created = Invoke-RestMethod -Method Post -Uri "$baseUrl/api/collections" -Headers $headers -ContentType "application/json" -Body $body
  return $created.id
}

$collections = @("games","players","events","clues","persons","hypotheses","accusations","forensics","timeline","interrogations")
$ids = @{}
foreach ($name in $collections) {
  $ids[$name] = Ensure-Collection $name
}

function Update-Collection($name, $schema, $indexes = @()) {
  $id = $ids[$name]
  $body = @{ name = $name; type = "base"; schema = $schema; indexes = $indexes } | ConvertTo-Json -Depth 6
  Invoke-RestMethod -Method Patch -Uri "$baseUrl/api/collections/$id" -Headers $headers -ContentType "application/json" -Body $body | Out-Null
}

$gamesSchema = @(
  @{ name = "code"; type = "text"; required = $true; options = @{ min = 1; max = 50; pattern = "" } },
  @{ name = "state"; type = "text"; required = $true; options = @{ min = 1; max = 50; pattern = "" } },
  @{ name = "seed"; type = "text"; required = $true; options = @{ min = 1; max = 200; pattern = "" } }
)

$playersSchema = @(
  @{ name = "gameId"; type = "relation"; required = $true; options = @{ collectionId = $ids["games"]; cascadeDelete = $true; maxSelect = 1 } },
  @{ name = "userId"; type = "text"; required = $true; options = @{ min = 1; max = 100; pattern = "" } },
  @{ name = "role"; type = "text"; required = $true; options = @{ min = 1; max = 30; pattern = "" } },
  @{ name = "status"; type = "text"; required = $true; options = @{ min = 1; max = 30; pattern = "" } },
  @{ name = "isHost"; type = "bool"; required = $true }
)

$eventsSchema = @(
  @{ name = "gameId"; type = "relation"; required = $true; options = @{ collectionId = $ids["games"]; cascadeDelete = $true; maxSelect = 1 } },
  @{ name = "timestamp"; type = "text"; required = $true; options = @{ min = 1; max = 50; pattern = "" } },
  @{ name = "locationId"; type = "text"; required = $true; options = @{ min = 1; max = 100; pattern = "" } },
  @{ name = "participants"; type = "relation"; required = $false; options = @{ collectionId = $ids["persons"]; cascadeDelete = $false; maxSelect = 999 } }
)

$cluesSchema = @(
  @{ name = "gameId"; type = "relation"; required = $true; options = @{ collectionId = $ids["games"]; cascadeDelete = $true; maxSelect = 1 } },
  @{ name = "type"; type = "text"; required = $true; options = @{ min = 1; max = 50; pattern = "" } },
  @{ name = "state"; type = "text"; required = $true; options = @{ min = 1; max = 50; pattern = "" } },
  @{ name = "reliability"; type = "number"; required = $true; options = @{ min = 0; max = 100; noDecimal = $true } },
  @{ name = "facts"; type = "json"; required = $false; options = @{ maxSize = 200000 } }
)

$personsSchema = @(
  @{ name = "gameId"; type = "relation"; required = $true; options = @{ collectionId = $ids["games"]; cascadeDelete = $true; maxSelect = 1 } },
  @{ name = "name"; type = "text"; required = $true; options = @{ min = 1; max = 100; pattern = "" } },
  @{ name = "officialStory"; type = "text"; required = $true; options = @{ min = 1; max = 5000; pattern = "" } },
  @{ name = "truthStory"; type = "text"; required = $true; options = @{ min = 1; max = 5000; pattern = "" } },
  @{ name = "stress"; type = "number"; required = $true; options = @{ min = 0; max = 100; noDecimal = $true } },
  @{ name = "credibility"; type = "number"; required = $true; options = @{ min = 0; max = 100; noDecimal = $true } }
)

$hypothesesSchema = @(
  @{ name = "gameId"; type = "relation"; required = $true; options = @{ collectionId = $ids["games"]; cascadeDelete = $true; maxSelect = 1 } },
  @{ name = "title"; type = "text"; required = $true; options = @{ min = 1; max = 200; pattern = "" } },
  @{ name = "strengthScore"; type = "number"; required = $true; options = @{ min = 0; max = 100; noDecimal = $true } },
  @{ name = "status"; type = "text"; required = $true; options = @{ min = 1; max = 50; pattern = "" } },
  @{ name = "nodeIds"; type = "json"; required = $false; options = @{ maxSize = 200000 } }
)

$accusationsSchema = @(
  @{ name = "gameId"; type = "relation"; required = $true; options = @{ collectionId = $ids["games"]; cascadeDelete = $true; maxSelect = 1 } },
  @{ name = "playerId"; type = "relation"; required = $true; options = @{ collectionId = $ids["players"]; cascadeDelete = $false; maxSelect = 1 } },
  @{ name = "suspectId"; type = "relation"; required = $true; options = @{ collectionId = $ids["persons"]; cascadeDelete = $false; maxSelect = 1 } },
  @{ name = "motiveId"; type = "text"; required = $true; options = @{ min = 1; max = 100; pattern = "" } },
  @{ name = "evidenceId"; type = "relation"; required = $true; options = @{ collectionId = $ids["clues"]; cascadeDelete = $false; maxSelect = 1 } },
  @{ name = "verdict"; type = "text"; required = $false; options = @{ min = 0; max = 50; pattern = "" } }
)

$forensicsSchema = @(
  @{ name = "gameId"; type = "relation"; required = $true; options = @{ collectionId = $ids["games"]; cascadeDelete = $true; maxSelect = 1 } },
  @{ name = "clueId"; type = "relation"; required = $true; options = @{ collectionId = $ids["clues"]; cascadeDelete = $false; maxSelect = 1 } },
  @{ name = "result"; type = "text"; required = $true; options = @{ min = 1; max = 200; pattern = "" } },
  @{ name = "confidence"; type = "number"; required = $true; options = @{ min = 0; max = 100; noDecimal = $true } },
  @{ name = "status"; type = "text"; required = $true; options = @{ min = 1; max = 20; pattern = "" } }
)

$timelineSchema = @(
  @{ name = "gameId"; type = "relation"; required = $true; options = @{ collectionId = $ids["games"]; cascadeDelete = $true; maxSelect = 1 } },
  @{ name = "timestamp"; type = "text"; required = $true; options = @{ min = 1; max = 50; pattern = "" } },
  @{ name = "title"; type = "text"; required = $true; options = @{ min = 1; max = 200; pattern = "" } },
  @{ name = "description"; type = "text"; required = $true; options = @{ min = 1; max = 5000; pattern = "" } },
  @{ name = "eventId"; type = "relation"; required = $false; options = @{ collectionId = $ids["events"]; cascadeDelete = $false; maxSelect = 1 } }
)

$interrogationsSchema = @(
  @{ name = "gameId"; type = "relation"; required = $true; options = @{ collectionId = $ids["games"]; cascadeDelete = $true; maxSelect = 1 } },
  @{ name = "personId"; type = "relation"; required = $true; options = @{ collectionId = $ids["persons"]; cascadeDelete = $false; maxSelect = 1 } },
  @{ name = "question"; type = "text"; required = $true; options = @{ min = 1; max = 1000; pattern = "" } },
  @{ name = "answer"; type = "text"; required = $true; options = @{ min = 1; max = 5000; pattern = "" } },
  @{ name = "tone"; type = "text"; required = $true; options = @{ min = 1; max = 50; pattern = "" } }
)

Update-Collection "games" $gamesSchema
Update-Collection "players" $playersSchema
Update-Collection "events" $eventsSchema
Update-Collection "clues" $cluesSchema
Update-Collection "persons" $personsSchema
Update-Collection "hypotheses" $hypothesesSchema
Update-Collection "accusations" $accusationsSchema
Update-Collection "forensics" $forensicsSchema
Update-Collection "timeline" $timelineSchema
Update-Collection "interrogations" $interrogationsSchema

Write-Host "PocketBase schema applied."
