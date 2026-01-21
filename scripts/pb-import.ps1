param(
  [string]$EnvFile = "",
  [string]$ImportFile = "",
  [switch]$Debug
)

$ErrorActionPreference = "Stop"

if (-not $EnvFile) {
  $EnvFile = Join-Path $PSScriptRoot "..\.env.local"
}

if (-not $ImportFile) {
  $ImportFile = Join-Path $PSScriptRoot "..\docs\db\pocketbase-import.json"
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
if (-not $env:PB_ADMIN_TOKEN) { throw "PB_ADMIN_TOKEN not set" }

$pbUri = [System.Uri]$env:PB_URL
$baseUrl = "$($pbUri.Scheme)://$($pbUri.Host)"
if (-not $pbUri.IsDefaultPort) {
  $baseUrl = "${baseUrl}:$($pbUri.Port)"
}

if ($Debug) {
  Write-Host "PB_URL set: " -NoNewline; Write-Host ([bool]$env:PB_URL)
  Write-Host "PB_ADMIN_TOKEN set: " -NoNewline; Write-Host ([bool]$env:PB_ADMIN_TOKEN)
  Write-Host "Import file: $ImportFile"
}

$headers = @{ Authorization = "Bearer $env:PB_ADMIN_TOKEN" }
$collections = Invoke-RestMethod -Method Get -Uri "$baseUrl/api/collections" -Headers $headers
$importRoot = Get-Content $ImportFile -Raw | ConvertFrom-Json
$importData = $importRoot.collections
if (-not $importData) { throw "Import file missing collections array" }

foreach ($col in $importData) {
  $existing = $collections.items | Where-Object { $_.name -eq $col.name }
  if ($existing -and -not $existing.system) {
    Invoke-RestMethod -Method Delete -Uri "$baseUrl/api/collections/$($existing.id)" -Headers $headers | Out-Null
  }
}

$idMap = @{}
foreach ($col in $importData) {
  $body = @{ name = $col.name; type = $col.type } | ConvertTo-Json
  $created = Invoke-RestMethod -Method Post -Uri "$baseUrl/api/collections" -Headers $headers -ContentType "application/json" -Body $body
  if ($col.id) {
    $idMap[$col.id] = $created.id
  }
}

if ($Debug) {
  Write-Host "ID map:"; $idMap.GetEnumerator() | ForEach-Object { Write-Host "  $($_.Key) -> $($_.Value)" }
}

foreach ($col in $importData) {
  $targetId = if ($col.id) { $idMap[$col.id] } else { ($collections.items | Where-Object { $_.name -eq $col.name }).id }
  $fields = @()
  foreach ($field in $col.schema) {
    $options = $field.options
    if ($field.type -eq "relation") {
      if (-not $options) { throw "Missing relation options in $($col.name).$($field.name)" }
      if (-not $options.collectionId) { throw "Missing collectionId in $($col.name).$($field.name)" }
      if ($idMap[$options.collectionId]) {
        $options.collectionId = $idMap[$options.collectionId]
      } else {
        throw "Unknown collectionId $($options.collectionId) in $($col.name).$($field.name)"
      }
    }
    $fieldPayload = @{ name = $field.name; type = $field.type; required = $field.required }
    if ($field.id) { $fieldPayload["id"] = $field.id }
    if ($options) {
      $options.PSObject.Properties | ForEach-Object {
        $fieldPayload[$_.Name] = $_.Value
      }
    }
    $fields += $fieldPayload
  }
  if ($Debug -and $col.name -eq "players") {
    Write-Host "Players payload:"; (@{ name = $col.name; type = $col.type; fields = $fields } | ConvertTo-Json -Depth 10)
  }
  $body = @{ name = $col.name; type = $col.type; fields = $fields } | ConvertTo-Json -Depth 10
  Write-Host "Updating collection: $($col.name)"
  try {
    Invoke-RestMethod -Method Patch -Uri "$baseUrl/api/collections/$targetId" -Headers $headers -ContentType "application/json" -Body $body | Out-Null
  } catch {
    Write-Host "Failed on: $($col.name)"
    throw
  }
}

Write-Host "PocketBase import completed."
