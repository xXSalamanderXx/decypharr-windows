param(
    [string]$Version = "dev",
    [string]$Channel = "stable"
)

Set-StrictMode -Version Latest
$root = Split-Path -Parent $MyInvocation.MyCommand.Definition
Push-Location $root

# Ensure go is available
if (-not (Get-Command go -ErrorAction SilentlyContinue)) {
    Write-Error "Go not found in PATH"
    exit 1
}

$env:CGO_ENABLED = "0"

# Adjust module path below if your module path differs.
$ldflags = "-w -s -X github.com/sirrobot01/decypharr/pkg/version.Version=$Version -X github.com/sirrobot01/decypharr/pkg/version.Channel=$Channel -trimpath"

$OutDir = Join-Path $root "dist"
if (-Not (Test-Path $OutDir)) {
    New-Item -ItemType Directory -Path $OutDir | Out-Null
}

$exe = Join-Path $OutDir "decypharr.exe"
Write-Host "Building $exe ..."

go build -o $exe -ldflags $ldflags

if ($LASTEXITCODE -ne 0) {
    Write-Error "Go build failed with exit code $LASTEXITCODE"
    Pop-Location
    exit $LASTEXITCODE
}

Write-Host "Built $exe successfully"
Pop-Location
