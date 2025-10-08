# ==================================================================
# FusionPDF - Build and Run Script
# Compatible: Windows PowerShell 5+ / PowerShell Core 7+
# ==================================================================

Write-Host "=== FusionPDF Build Script ===" -ForegroundColor Cyan

# --------------------------
# Step 1: Check if Go is installed
# --------------------------
$goVersion = & go version 2>$null
if (-not $goVersion) {
    Write-Host "Go is not installed. Installing Go 1.25.1..." -ForegroundColor Yellow

    $goInstaller = "$env:TEMP\go_installer.msi"
    Invoke-WebRequest -Uri "https://go.dev/dl/go1.25.1.windows-amd64.msi" -OutFile $goInstaller -UseBasicParsing

    Write-Host "Running Go 1.25.1 installer..." -ForegroundColor Yellow
    Start-Process msiexec.exe -ArgumentList "/i `"$goInstaller`" /quiet /norestart" -Wait

    Remove-Item $goInstaller -Force
    Write-Host "Go 1.25.1 installed successfully. Please restart PowerShell if it's the first installation." -ForegroundColor Green
}
else {
    Write-Host "Go detected: $goVersion" -ForegroundColor Green
}

# --------------------------
# Step 2: Check if project exists
# --------------------------
if (-not (Test-Path "go.mod")) {
    Write-Host "Error: go.mod not found. Run this script from the project root." -ForegroundColor Red
    exit 1
}

# --------------------------
# Step 3: Install dependencies
# --------------------------
Write-Host "`nInstalling dependencies..." -ForegroundColor Yellow
go mod tidy
if ($LASTEXITCODE -ne 0) {
    Write-Host "Dependency installation failed." -ForegroundColor Red
    exit 1
}
Write-Host "Dependencies installed successfully." -ForegroundColor Green

# --------------------------
# Step 4: Build the binary
# --------------------------
$buildPath = "FusionPDF.exe"
Write-Host "`nBuilding FusionPDF..." -ForegroundColor Yellow
go build -ldflags="-s -w -H=windowsgui" -o $buildPath .
if ($LASTEXITCODE -ne 0) {
    Write-Host "Build failed. Check for errors above." -ForegroundColor Red
    exit 1
}
Write-Host "Build succeeded: $buildPath" -ForegroundColor Green

# --------------------------
# Step 5: Run the application
# --------------------------
Write-Host "`nLaunching FusionPDF..." -ForegroundColor Cyan
Start-Process -FilePath ".\FusionPDF.exe"

Write-Host "`n=== Build & Launch complete ===" -ForegroundColor Green
