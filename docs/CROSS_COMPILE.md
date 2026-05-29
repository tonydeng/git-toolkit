# Git Toolkit 交叉编译使用说明

## 概述

Go 语言原生支持交叉编译，可以在一个平台上编译出其他平台的可执行文件。本文档介绍如何为 Git Toolkit 进行多平台交叉编译。

## 环境变量说明

交叉编译主要通过两个环境变量控制：

| 变量 | 说明 | 示例值 |
|------|------|--------|
| `GOOS` | 目标操作系统 | `linux`, `windows`, `darwin` |
| `GOARCH` | 目标架构 | `amd64`, `arm64`, `386` |

## 支持的平台

### 常用平台组合

| 平台 | GOOS | GOARCH | 说明 |
|------|------|--------|------|
| Linux 64-bit | linux | amd64 | 服务器常用 |
| Linux 32-bit | linux | 386 | 旧系统支持 |
| Linux ARM64 | linux | arm64 | ARM 服务器 |
| Windows 64-bit | windows | amd64 | Windows 服务器 |
| Windows 32-bit | windows | 386 | 旧版 Windows |
| macOS Intel | darwin | amd64 | Intel Mac |
| macOS Apple Silicon | darwin | arm64 | M1/M2 Mac |

### 查看所有支持的平台

```bash
go tool dist list
```

## 编译方法

### 方法一：使用 Makefile（推荐）

```bash
# 编译所有平台
make all

# 编译特定平台
make build-linux-amd64
make build-windows-amd64
make build-darwin-arm64
```

### 方法二：使用 PowerShell（Windows）

```powershell
# 编译所有平台
.\build.ps1 -BuildAll

# 编译特定平台
.\build.ps1 -Platform windows-amd64

# 使用 vendor 模式编译
.\build.ps1 -BuildAll -UseVendor
```

### 方法三：直接使用 go build

#### Linux/macOS

```bash
# 编译 Linux amd64
GOOS=linux GOARCH=amd64 go build -o dist/git-toolkit_linux_amd64

# 编译 Windows amd64
GOOS=windows GOARCH=amd64 go build -o dist/git-toolkit_windows_amd64.exe

# 编译 macOS arm64
GOOS=darwin GOARCH=arm64 go build -o dist/git-toolkit_darwin_arm64
```

#### Windows PowerShell

```powershell
# 编译 Linux amd64
$env:GOOS="linux"; $env:GOARCH="amd64"; go build -o dist/git-toolkit_linux_amd64

# 编译 Windows amd64
$env:GOOS="windows"; $env:GOARCH="amd64"; go build -o dist/git-toolkit_windows_amd64.exe

# 编译 macOS arm64
$env:GOOS="darwin"; $env:GOARCH="arm64"; go build -o dist/git-toolkit_darwin_arm64
```

#### Windows CMD

```cmd
# 编译 Linux amd64
set GOOS=linux
set GOARCH=amd64
go build -o dist/git-toolkit_linux_amd64

# 编译 Windows amd64
set GOOS=windows
set GOARCH=amd64
go build -o dist/git-toolkit_windows_amd64.exe
```

## 版本信息注入

编译时可以注入版本信息：

```bash
# 读取版本号
VERSION=$(shell cat version)
BUILD_TIME=$(shell date "+%F %T")
COMMIT=$(shell git rev-parse HEAD)

# 注入版本信息
LDFLAGS="-X 'github.com/tonydeng/git-toolkit/cmd.Version=${VERSION}' \
         -X 'github.com/tonydeng/git-toolkit/cmd.BuildTime=${BUILD_TIME}' \
         -X 'github.com/tonydeng/git-toolkit/cmd.Commit=${COMMIT}'"

# 编译
go build -ldflags "$LDFLAGS" -o dist/git-toolkit
```

## 使用 Vendor 模式

Vendor 模式将所有依赖打包到项目中，实现离线编译：

```bash
# 1. 创建 vendor 目录
go mod vendor

# 2. 使用 vendor 编译
go build -mod=vendor -o dist/git-toolkit

# 或在 Makefile 中使用
make build GOFLAGS="-mod=vendor"
```

## CGO 注意事项

### 禁用 CGO

如果项目不依赖 C 库，建议禁用 CGO 以实现纯静态编译：

```bash
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o dist/git-toolkit
```

### 启用 CGO

如果需要 CGO（如使用 SQLite），交叉编译会更复杂：

```bash
# 需要安装对应平台的 C 交叉编译工具链
# 例如在 Linux 上编译 Windows 程序
sudo apt-get install mingw-w64
CC=x86_64-w64-mingw32-gcc GOOS=windows GOARCH=amd64 CGO_ENABLED=1 go build
```

## 构建脚本示例

### Makefile 完整示例

```makefile
# 变量定义
GOCMD=go
GOBUILD=$(GOCMD) build
BINARY_NAME=git-toolkit
DIST_DIR=dist
VERSION=$(shell cat version)
BUILD_TIME=$(shell date "+%F %T")
COMMIT_SHA1=$(shell git rev-parse HEAD)

LDFLAGS=-ldflags "-X 'github.com/tonydeng/git-toolkit/cmd.Version=${VERSION}' \
                  -X 'github.com/tonydeng/git-toolkit/cmd.BuildTime=${BUILD_TIME}' \
                  -X 'github.com/tonydeng/git-toolkit/cmd.Commit=${COMMIT_SHA1}'"

.PHONY: all clean build build-linux build-windows build-darwin

all: build-linux build-windows build-darwin

build:
	$(GOBUILD) $(LDFLAGS) -o $(DIST_DIR)/$(BINARY_NAME)

build-linux:
	GOOS=linux GOARCH=amd64 $(GOBUILD) $(LDFLAGS) -o $(DIST_DIR)/$(BINARY_NAME)_linux_amd64

build-windows:
	GOOS=windows GOARCH=amd64 $(GOBUILD) $(LDFLAGS) -o $(DIST_DIR)/$(BINARY_NAME)_windows_amd64.exe

build-darwin:
	GOOS=darwin GOARCH=amd64 $(GOBUILD) $(LDFLAGS) -o $(DIST_DIR)/$(BINARY_NAME)_darwin_amd64
	GOOS=darwin GOARCH=arm64 $(GOBUILD) $(LDFLAGS) -o $(DIST_DIR)/$(BINARY_NAME)_darwin_arm64

clean:
	rm -rf $(DIST_DIR)
```

### PowerShell 完整示例

```powershell
param(
    [switch]$BuildAll,
    [switch]$Clean,
    [switch]$UseVendor,
    [string]$Platform
)

$BinaryName = "git-toolkit"
$DistDir = "dist"
$Version = Get-Content version
$BuildTime = Get-Date -Format "yyyy-MM-dd HH:mm:ss"
$Commit = git rev-parse HEAD

$LdFlags = "-X 'github.com/tonydeng/git-toolkit/cmd.Version=$Version' " + 
           "-X 'github.com/tonydeng/git-toolkit/cmd.BuildTime=$BuildTime' " +
           "-X 'github.com/tonydeng/git-toolkit/cmd.Commit=$Commit'"

$GoFlags = if ($UseVendor) { "-mod=vendor" } else { "" }

function Build-Platform {
    param($GOOS, $GOARCH, $Suffix)
    
    $env:GOOS = $GOOS
    $env:GOARCH = $GOARCH
    $output = "$DistDir/${BinaryName}_$GOOS_$GOARCH$Suffix"
    
    Write-Host "Building for $GOOS/$GOARCH..."
    go build $GoFlags -ldflags $LdFlags -o $output
    
    $env:GOOS = ""
    $env:GOARCH = ""
}

if ($Clean -and (Test-Path $DistDir)) {
    Remove-Item -Recurse -Force $DistDir
}

New-Item -ItemType Directory -Force -Path $DistDir | Out-Null

if ($BuildAll) {
    Build-Platform "linux" "amd64" ""
    Build-Platform "linux" "386" ""
    Build-Platform "windows" "amd64" ".exe"
    Build-Platform "windows" "386" ".exe"
    Build-Platform "darwin" "amd64" ""
    Build-Platform "darwin" "arm64" ""
} elseif ($Platform) {
    $parts = $Platform -split "-"
    $Suffix = if ($parts[0] -eq "windows") { ".exe" } else { "" }
    Build-Platform $parts[0] $parts[1] $Suffix
} else {
    Build-Platform $env:GOOS $env:GOARCH ""
}

Write-Host "Build complete!"
```

## 常见问题

### 1. 编译后文件过大

使用 `-s -w` 标志去除调试信息：

```bash
go build -ldflags "-s -w" -o dist/git-toolkit
```

### 2. Windows 下编译 Linux 程序失败

确保设置了正确的环境变量：

```powershell
$env:CGO_ENABLED=0
$env:GOOS="linux"
$env:GOARCH="amd64"
```

### 3. macOS arm64 程序在 Intel Mac 上无法运行

这是正常的，arm64 程序只能在 Apple Silicon Mac 上运行。需要同时编译 amd64 和 arm64 版本。

### 4. 静态链接问题

对于完全静态链接（无动态库依赖）：

```bash
CGO_ENABLED=0 go build -a -installsuffix cgo -o dist/git-toolkit
```

## 验证编译结果

```bash
# 检查文件类型
file dist/git-toolkit_linux_amd64

# 检查 Windows 文件
file dist/git-toolkit_windows_amd64.exe

# 在对应平台运行
./dist/git-toolkit_linux_amd64 version
```

## 参考资料

1. [Go Cross Compilation](https://go.dev/doc/install/source#environment)
2. [Go Wiki: WindowsCrossCompiling](https://go.dev/wiki/WindowsCrossCompiling)
3. [Go Modules](https://go.dev/wiki/Modules)
