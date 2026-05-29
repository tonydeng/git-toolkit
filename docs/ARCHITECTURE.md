# Git Toolkit 项目架构说明

## 项目概述

Git Toolkit 是一个用于简化 Git 工作流程的命令行工具集，支持 GitLab 集成，提供 Merge Request 管理、分支操作、代码日志等功能。

## 目录结构

```
git-toolkit-github/
├── cmd/                    # 命令行应用入口
│   ├── root.go            # 根命令定义
│   ├── amr.go             # Accept MR 命令
│   ├── branch.go          # 分支操作命令
│   ├── ci.go              # CI 相关命令
│   ├── clog.go            # 变更日志命令
│   ├── cm.go              # Commit 命令
│   ├── cmr.go             # Create MR 命令
│   ├── dmr.go             # Delete MR 命令
│   ├── gmr.go             # Get MR 命令
│   ├── umr.go             # Update MR 命令
│   ├── install.go         # 安装命令
│   ├── uninstall.go       # 卸载命令
│   ├── ps.go              # Push 命令
│   ├── repl.go            # REPL 模式命令
│   └── version.go         # 版本信息命令
├── git/                    # Git 操作封装
│   ├── git.go             # Git 核心操作
│   ├── ci.go              # CI 集成
│   ├── clog.go            # 变更日志生成
│   ├── cm.go              # Commit 操作
│   └── filter_branch.go   # 分支过滤操作
├── mr/                     # Merge Request 功能
│   ├── merge_request.go   # MR 核心逻辑
│   └── mr_prompt.go       # MR 交互提示
├── utils/                  # 工具函数
│   ├── common.go          # 通用工具函数
│   ├── config.go          # 配置管理
│   ├── install.go         # 安装工具
│   ├── uninstall.go       # 卸载工具
│   └── version.go         # 版本管理
├── scripts/                # 构建和安装脚本
│   ├── build.ps1          # Windows 构建脚本
│   └── installer.sh       # Linux/macOS 安装脚本
├── docs/                   # 文档目录
│   ├── README.md          # 项目说明
│   ├── Install.md         # 安装指南
│   ├── Use.md             # 使用指南
│   ├── CHANGELOG.md       # 变更日志
│   └── images/            # 图片资源
├── dist/                   # 编译输出目录
├── vendor/                 # 依赖目录
├── go.mod                  # Go 模块定义
├── go.sum                  # 依赖校验文件
├── Makefile               # 构建脚本
├── main.go                # 程序入口
└── version                # 版本号文件
```

## 代码组织规范

### 1. 包设计原则

#### cmd 包
- 每个文件对应一个 CLI 子命令
- 使用 Cobra 框架定义命令
- 命令逻辑应尽量简洁，业务逻辑委托给其他包

#### git 包
- 封装 Git 命令行操作
- 提供 Git 相关的高级抽象
- 处理 Git 命令的执行和输出解析

#### mr 包
- 处理 Merge Request 相关业务逻辑
- 与 GitLab API 交互
- 提供用户交互提示

#### utils 包
- 提供通用工具函数
- 配置管理
- 版本信息管理

### 2. 命名规范

遵循 Google Go Style Guide 和 Uber Go Style Guide：

| 类型 | 规范 | 示例 |
|------|------|------|
| 包名 | 小写、简洁、无下划线 | `git`, `mr`, `utils` |
| 导出函数 | MixedCaps | `GetMergeRequest`, `CreateBranch` |
| 私有函数 | mixedCaps | `parseConfig`, `validateInput` |
| 常量 | MixedCaps 或全大写 | `MaxRetries`, `DEFAULT_TIMEOUT` |
| 接口 | 动词+er 后缀 | `Reader`, `Writer`, `ConfigLoader` |
| 错误变量 | Err 前缀 | `ErrNotFound`, `ErrInvalidInput` |

### 3. 错误处理

```go
// 推荐：使用 errors.New 或 fmt.Errorf
var ErrNotFound = errors.New("resource not found")

// 推荐：错误包装使用 %w
if err := doSomething(); err != nil {
    return fmt.Errorf("failed to process: %w", err)
}

// 推荐：显式检查错误
result, err := process()
if err != nil {
    return nil, err
}
```

## 依赖管理策略

### Go Modules

项目使用 Go Modules 进行依赖管理（Go 1.21+）：

```go
module github.com/tonydeng/git-toolkit

go 1.21

require (
    github.com/spf13/cobra v1.8.0
    github.com/xanzy/go-gitlab v0.52.0
    github.com/mitchellh/go-homedir v1.1.0
    gopkg.in/yaml.v2 v2.4.0
)
```

### Vendor 模式

使用 vendor 目录确保构建可重复性：

```bash
# 创建 vendor 目录
go mod vendor

# 使用 vendor 构建
go build -mod=vendor
```

### 依赖更新

```bash
# 更新所有依赖到最新版本
go get -u ./...

# 更新特定依赖
go get github.com/spf13/cobra@latest

# 整理依赖
go mod tidy
```

## 构建流程说明

### Makefile 构建

```bash
# 构建所有平台（仅 64 位）
make all

# 构建特定平台
make build-linux-amd64
make build-windows-amd64
make build-darwin-amd64
make build-darwin-arm64

# 清理构建产物
make clean
```

### Windows PowerShell 构建

```powershell
# 构建所有平台
.\scripts\build.ps1 -UseVendor

# 构建特定平台
.\scripts\build.ps1 -Platform windows-amd64

# 清理后构建
.\scripts\build.ps1 -Clean -UseVendor
```

### 版本信息注入

构建时自动注入版本信息：

```bash
# 通过 Makefile
make all  # 自动从 version 文件读取版本号

# 手动指定版本
go build -ldflags "-X 'github.com/tonydeng/git-toolkit/cmd.Version=v2.1.1'"
```

## 最佳实践参考

### Google Go Style Guide 核心原则

1. **清晰性 (Clarity)**: 代码的目的和理由对读者清晰
2. **简洁性 (Simplicity)**: 以最简单的方式实现目标
3. **精炼性 (Concision)**: 高信噪比，突出重要细节
4. **可维护性 (Maintainability)**: 易于正确修改
5. **一致性 (Consistency)**: 与代码库保持一致

### Uber Go Style Guide 关键建议

1. **错误处理**: 显式检查错误，使用 `errors.Is` 和 `errors.As`
2. **接口设计**: 小接口优先，接收接口返回具体类型
3. **并发安全**: 使用 `sync.Mutex` 零值，避免复制包含锁的结构体
4. **Channel 使用**: 大小为 1 或无缓冲
5. **时间处理**: 使用 `time.Time` 和 `time.Duration`

## 参考资料

1. [Google Go Style Guide](https://google.github.io/styleguide/go/guide)
2. [Uber Go Style Guide](https://github.com/uber-go/guide/blob/master/style.md)
3. [golang-standards/project-layout](https://github.com/golang-standards/project-layout)
4. [Go Modules 官方文档](https://go.dev/wiki/Modules)
5. [Effective Go](https://go.dev/doc/effective_go)
