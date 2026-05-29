# git-toolkit AGENTS.md

本文件为 AI 编码代理提供项目上下文和指令，遵循 AGENTS.md 开放标准。

## 项目概述

git-toolkit 是一个 Git 工具集，提供 Git 操作封装和 Merge Request 管理功能。

- **语言**: Go 1.21+
- **主要功能**: Git 操作封装、Merge Request 管理
- **支持平台**: Linux (amd64), macOS (amd64/arm64), Windows (amd64)

## 构建命令

### 构建所有平台

```bash
go run scripts/build.go
```

### 构建单个平台

```bash
go run scripts/build.go linux-amd64
go run scripts/build.go darwin-amd64
go run scripts/build.go darwin-arm64
go run scripts/build.go windows-amd64
```

### 清理构建产物

```bash
go run scripts/build.go clean
```

### 本地安装

```bash
go install -mod=vendor
```

## 测试命令

```bash
go test -mod=vendor -v ./...
```

## 依赖管理

### 使用 vendor 模式

项目使用 vendor 目录管理依赖，确保构建可重复。

```bash
# 整理依赖
go mod tidy

# 更新 vendor 目录
go mod vendor

# 使用 vendor 模式构建
go build -mod=vendor
```

### 设置 Go 代理（中国大陆）

```bash
# PowerShell
$env:GOPROXY="https://goproxy.cn,direct"

# Bash
export GOPROXY="https://goproxy.cn,direct"
```

## 项目结构

```
git-toolkit-github/
├── cmd/                    # 命令行应用入口
│   └── root.go            # 根命令定义
├── git/                    # Git 操作封装
├── mr/                     # Merge Request 功能
├── utils/                  # 工具函数
├── scripts/                # 构建和安装脚本
│   ├── build.go           # 核心构建脚本（跨平台）
│   └── installer.sh       # 安装脚本
├── docs/                   # 文档目录
├── dist/                   # 编译输出目录
├── vendor/                 # 依赖目录
├── go.mod
├── go.sum
└── main.go
```

## 代码风格

### 遵循 Google Go Style Guide

- 清晰性：代码易于理解
- 简洁性：避免不必要的复杂性
- 精炼性：代码简洁明了
- 可维护性：易于修改和扩展
- 一致性：风格统一

### 导入顺序

```go
import (
    // 标准库
    "fmt"
    "os"
    
    // 第三方库
    "github.com/spf13/cobra"
    
    // 本地包
    "github.com/tonydeng/git-toolkit/git"
)
```

### 错误处理

```go
// 使用 errors.Wrap 或 fmt.Errorf 包装错误
if err != nil {
    return fmt.Errorf("failed to do something: %w", err)
}
```

### 已弃用的包

- ❌ `io/ioutil` → ✅ `os` 和 `io` 包
  - `ioutil.ReadFile` → `os.ReadFile`
  - `ioutil.TempFile` → `os.CreateTemp`
  - `ioutil.ReadAll` → `io.ReadAll`

## 跨平台编译注意事项

### 环境变量设置

在 PowerShell 中，环境变量必须与命令在同一行执行：

```powershell
# ❌ 错误方式 - 环境变量不会传递
$env:GOOS = "windows"
go build -o output

# ✅ 正确方式 - 环境变量在同一行设置
$env:GOOS="windows"; $env:GOARCH="amd64"; go build -mod=vendor -o output .
```

### 使用 Go 标准库

优先使用 Go 标准库而非外部命令，确保跨平台兼容：

```go
// ❌ 避免使用外部命令
cmd := exec.Command("date", "+%F %T")

// ✅ 使用 Go 标准库
time.Now().Format("2006-01-02 15:04:05")
```

## 边界和约束

### 始终执行

- ✅ 使用 `-mod=vendor` 标志构建
- ✅ 遵循 golang-standards/project-layout 目录结构
- ✅ 使用 Go 标准库替代外部命令
- ✅ 在修改代码后运行测试

### 先询问

- ⚠️ 修改 vendor 目录中的依赖
- ⚠️ 修改 go.mod 中的 Go 版本
- ⚠️ 添加新的第三方依赖

### 永不执行

- ❌ 提交敏感信息（API 密钥、密码等）
- ❌ 使用已弃用的包（如 `io/ioutil`）
- ❌ 在 Windows 下使用 `date` 命令（不可移植）
- ❌ 创建不必要的 wrapper 脚本

## Git 工作流

### 分支命名

- `feature/xxx` - 新功能
- `fix/xxx` - Bug 修复
- `refactor/xxx` - 重构

### 提交信息格式

```
<type>(<scope>): <subject>

<body>

<footer>
```

### 类型

- `feat`: 新功能
- `fix`: Bug 修复
- `refactor`: 重构
- `docs`: 文档更新
- `test`: 测试相关
- `chore`: 构建/工具相关

## 版本信息注入

构建时自动注入版本信息：

```go
// ldflags 参数
-ldflags "-s -w -X 'github.com/tonydeng/git-toolkit/cmd.Version=${VERSION}' \
          -X 'github.com/tonydeng/git-toolkit/cmd.BuildTime=${BUILD_TIME}' \
          -X 'github.com/tonydeng/git-toolkit/cmd.CommitID=${COMMIT_SHA}'"
```

## 常见问题

### BuildTime 显示 unknown

**原因**: 在 Windows 下，`powershell` 命令可能不在 PATH 中。

**解决方案**: 使用 Go 标准库 `time.Now()` 替代外部命令。

### 依赖下载失败

**解决方案**: 设置 Go 代理：

```bash
$env:GOPROXY="https://goproxy.cn,direct"
```

### API 兼容性问题

升级依赖后检查 API 变化，如 `go-gitlab` 的 `GetMergeRequestChanges` 方法签名已改变。

## 参考资料

- [Google Go Style Guide](https://google.github.io/styleguide/go/guide)
- [Uber Go Style Guide](https://github.com/uber-go/guide/blob/master/style.md)
- [golang-standards/project-layout](https://github.com/golang-standards/project-layout)
- [Go Modules 官方文档](https://go.dev/ref/mod)

---

**文件版本**: v1.0.0  
**最后更新**: 2026-05-29
