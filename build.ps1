<#
.SYNOPSIS
    Build script for git-toolkit - Cross-platform compilation with version injection.

.DESCRIPTION
    This script compiles git-toolkit for multiple platforms with version information
    injected via ldflags. Supports both PowerShell 5.1 and PowerShell 7.

.PARAMETER Platform
    Specify a single platform to build (e.g., "windows/amd64").
    If not specified, builds all platforms.

.PARAMETER OutputDir
    Output directory for compiled binaries. Default is "dist".

.PARAMETER UseVendor
    Use vendor directory for dependencies (-mod=vendor).

.PARAMETER Clean
    Clean the output directory before building.

.EXAMPLE
    .\build.ps1
    Build all platforms.

.EXAMPLE
    .\build.ps1 -Platform "windows/amd64"
    Build only for Windows amd64.

.EXAMPLE
    .\build.ps1 -Clean -UseVendor
    Clean output directory and build using vendor dependencies.
#>

[CmdletBinding()]
param(
    [Parameter(Mandatory = $false)]
    [string]$Platform,

    [Parameter(Mandatory = $false)]
    [string]$OutputDir = "dist",

    [Parameter(Mandatory = $false)]
    [switch]$UseVendor,

    [Parameter(Mandatory = $false)]
    [switch]$Clean
)

# Error action preference
$ErrorActionPreference = "Stop"

# Project root directory
$ScriptDir = Split-Path -Parent $MyInvocation.MyCommand.Path
$ProjectRoot = $ScriptDir

# Define all supported platforms
$AllPlatforms = @(
    @{OS = "darwin"; Arch = "amd64"},
    @{OS = "darwin"; Arch = "arm64"},
    @{OS = "linux"; Arch = "386"},
    @{OS = "linux"; Arch = "amd64"},
    @{OS = "windows"; Arch = "amd64"},
    @{OS = "windows"; Arch = "386"}
)

# Package path for ldflags
$PackagePath = "github.com/tonydeng/git-toolkit/cmd"

# Binary name
$BinaryName = "git-toolkit"

#region Functions

function Write-Info {
    param([string]$Message)
    Write-Host "[INFO] $Message" -ForegroundColor Cyan
}

function Write-Success {
    param([string]$Message)
    Write-Host "[OK] $Message" -ForegroundColor Green
}

function Write-ErrorMsg {
    param([string]$Message)
    Write-Host "[ERROR] $Message" -ForegroundColor Red
}

function Get-Version {
    $versionFile = Join-Path $ProjectRoot "version"
    if (Test-Path $versionFile) {
        $version = (Get-Content $versionFile -Raw).Trim()
        Write-Info "Version: $version"
        return $version
    }
    else {
        Write-ErrorMsg "Version file not found: $versionFile"
        exit 1
    }
}

function Get-BuildTime {
    # Format: YYYY-MM-DD HH:MM:SS
    $buildTime = Get-Date -Format "yyyy-MM-dd HH:mm:ss"
    Write-Info "Build Time: $buildTime"
    return $buildTime
}

function Get-CommitID {
    try {
        Push-Location $ProjectRoot
        $commitID = git rev-parse HEAD 2>$null
        if ($LASTEXITCODE -eq 0 -and $commitID) {
            $commitID = $commitID.Trim()
            Write-Info "Commit ID: $commitID"
            return $commitID
        }
        else {
            Write-Warning "Unable to get git commit ID, using 'unknown'"
            return "unknown"
        }
    }
    catch {
        Write-Warning "Git not available, using 'unknown' for commit ID"
        return "unknown"
    }
    finally {
        Pop-Location
    }
}

function Test-GoInstalled {
    try {
        $goVersion = go version 2>$null
        if ($LASTEXITCODE -eq 0) {
            Write-Info "Go version: $goVersion"
            return $true
        }
    }
    catch {
        # Ignore
    }
    Write-ErrorMsg "Go is not installed or not in PATH"
    return $false
}

function New-OutputDirectory {
    param([string]$Dir)

    $fullPath = Join-Path $ProjectRoot $Dir

    if ($Clean -and (Test-Path $fullPath)) {
        Write-Info "Cleaning output directory: $fullPath"
        Remove-Item -Path $fullPath -Recurse -Force
    }

    if (-not (Test-Path $fullPath)) {
        New-Item -Path $fullPath -ItemType Directory -Force | Out-Null
        Write-Info "Created output directory: $fullPath"
    }

    return $fullPath
}

function Build-Platform {
    param(
        [string]$OS,
        [string]$Arch,
        [string]$Version,
        [string]$BuildTime,
        [string]$CommitID,
        [string]$OutputPath,
        [bool]$UseVendorFlag
    )

    $outputFile = Join-Path $OutputPath "${BinaryName}_${OS}_${Arch}"

    # Add .exe suffix for Windows
    if ($OS -eq "windows") {
        $outputFile = "$outputFile.exe"
    }

    Write-Info "Building for ${OS}/${Arch}..."

    # Set environment variables for cross-compilation
    $env:GOOS = $OS
    $env:GOARCH = $Arch

    # Build ldflags
    $ldflags = "-X '${PackagePath}.Version=${Version}' " +
               "-X '${PackagePath}.BuildTime=${BuildTime}' " +
               "-X '${PackagePath}.CommitID=${CommitID}'"

    # Build go build arguments
    $buildArgs = @("build")

    # Add vendor flag if requested
    if ($UseVendorFlag) {
        $buildArgs += "-mod=vendor"
    }

    $buildArgs += "-ldflags"
    $buildArgs += $ldflags
    $buildArgs += "-o"
    $buildArgs += $outputFile
    $buildArgs += "."

    try {
        Push-Location $ProjectRoot

        $process = Start-Process -FilePath "go" `
            -ArgumentList $buildArgs `
            -NoNewWindow `
            -Wait `
            -PassThru

        if ($process.ExitCode -eq 0) {
            Write-Success "Built: $outputFile"
            return $true
        }
        else {
            Write-ErrorMsg "Failed to build for ${OS}/${arch}"
            return $false
        }
    }
    catch {
        Write-ErrorMsg "Build error for ${OS}/${Arch}: $_"
        return $false
    }
    finally {
        Pop-Location
        # Reset environment variables
        $env:GOOS = $null
        $env:GOARCH = $null
    }
}

#endregion

#region Main

Write-Host ""
Write-Host "========================================" -ForegroundColor Yellow
Write-Host "  git-toolkit Build Script" -ForegroundColor Yellow
Write-Host "========================================" -ForegroundColor Yellow
Write-Host ""

# Check Go installation
if (-not (Test-GoInstalled)) {
    exit 1
}

# Get version information
$version = Get-Version
$buildTime = Get-BuildTime
$commitID = Get-CommitID

# Check vendor directory
$vendorDir = Join-Path $ProjectRoot "vendor"
$useVendorFlag = $false
if ($UseVendor) {
    if (Test-Path $vendorDir) {
        $useVendorFlag = $true
        Write-Info "Using vendor directory for dependencies"
    }
    else {
        Write-Warning "Vendor directory not found, ignoring -UseVendor flag"
    }
}

# Create output directory
$outputPath = New-OutputDirectory -Dir $OutputDir

# Determine platforms to build
$platformsToBuild = @()
if ($Platform) {
    # Parse single platform specification
    if ($Platform -match "^(\w+)/(\w+)$") {
        $platformsToBuild = @(@{OS = $Matches[1]; Arch = $Matches[2]})
    }
    else {
        Write-ErrorMsg "Invalid platform format: $Platform. Expected format: os/arch (e.g., windows/amd64)"
        exit 1
    }
}
else {
    $platformsToBuild = $AllPlatforms
}

Write-Info "Building $($platformsToBuild.Count) platform(s)..."
Write-Host ""

# Build each platform
$successCount = 0
$failureCount = 0

foreach ($platform in $platformsToBuild) {
    $result = Build-Platform `
        -OS $platform.OS `
        -Arch $platform.Arch `
        -Version $version `
        -BuildTime $buildTime `
        -CommitID $commitID `
        -OutputPath $outputPath `
        -UseVendorFlag $useVendorFlag

    if ($result) {
        $successCount++
    }
    else {
        $failureCount++
    }
}

Write-Host ""
Write-Host "========================================" -ForegroundColor Yellow
Write-Host "  Build Summary" -ForegroundColor Yellow
Write-Host "========================================" -ForegroundColor Yellow
Write-Host "  Successful: $successCount" -ForegroundColor Green
Write-Host "  Failed:     $failureCount" -ForegroundColor $(if ($failureCount -gt 0) { "Red" } else { "Green" })
Write-Host "  Output:     $outputPath" -ForegroundColor Cyan
Write-Host ""

if ($failureCount -gt 0) {
    exit 1
}

exit 0

#endregion
