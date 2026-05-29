# git-toolkit

> 人类懒惰的本性和不满足的本性是驱使科技发展的源泉......

本工具集包含几个部分，自定义命令，`Hook`脚本，以及配置模板

## 下载 & 安装

### 下载

- [macOS](https://github.com/tonydeng/git-toolkit/raw/github/golang/dist/git-toolkit_darwin_amd64)
- [Linux 64](https://github.com/tonydeng/git-toolkit/raw/github/golang/dist/git-toolkit_linux_amd64)
- [Linux 32](https://github.com/tonydeng/git-toolkit/raw/github/golang/dist/git-toolkit_linux_386)

### 查看版本

```bash
./git-toolkit version
```

### 安装

```bash
./git-toolkit install
```

## 使用介绍

### 自定义命令

| 命令 | 描述 |
| -- | -- |
| **branch** | 创建分支的相关命令 |
| `feat` | 创建`feat`分支，接受一个字符串，并创建一个`feat`分支，分支名称格式为 `feat/xxx` |
| `fix` | 创建`fix`分支，接受一个字符串，并创建一个`fix`分支，分支名称格式为 `fix/xxx` |
| `docs` | 创建`docs`分支，接受一个字符串，并创建一个`docs`分支，分支名称格式为 `docs/xxx` |
| `style` | 创建`style`分支，接受一个字符串，并创建一个`style`分支，分支名称格式为 `style/xxx` |
| `refactor` | 创建`refactor`分支，接受一个字符串，并创建一个`refactor`分支，分支名称格式为 `refactor/xxx` |
| `chore` | 创建`chore`分支，接受一个字符串，并创建一个`chore`分支，分支名称格式为 `chore/xxx` |
| `perf` | 创建`perf`分支，接受一个字符串，并创建一个`perf`分支，分支名称格式为 `perf/xxx` |
| `design` | 创建`design`分支，接受一个字符串，并创建一个`design`分支，分支名称格式为 `design/xxx` |
| `hotfix` | 创建`hotfix`分支(通常用于对`master`紧急修复)，接受一个字符串，并创建一个`hotfix`分支，分支名称格式为 `hotfix/xxx` |
| **commit** | 提交的相关命令 |
| `ci` | 提供交互式`git commit`的命令，用于定制统一的`commit message` |
| `cm` | 检查`commit message`是否符合规范 |
| **push** | 推送的相关命令 |
| `ps` | 等同于`git push origin current_branch` |
| **CHANGELOG** | 日志的相关命令 |
| `clog` | 输出`git`的`change`日志 |
| **MergeRequest** | `Merge Request`的相关命令（需要配置 GitLab） |
| `cmr` | 创建指定项目的`MR` |
| `umr` | 更新指定项目的某个`MR` |
| `dmr` | 删除指定项目的某个`MR` |
| `amr` | 合并指定项目的某个`MR` |
| `gmr` | 获取指定项目的`MR`引起的改变 |
| **其他** | 其他命令 |
| `repl` | 替换历史提交中的邮箱或作者信息 |

### 支持使用场景

- 规范`Git`提交记录
- 生成`Changelog`
- 管理 GitLab Merge Request

## 环境要求

### 操作系统

目前只支持`Macos`、`Linux`

## 文档

- [安装说明](./docs/Install.md)
- [使用说明](./docs/Use.md)
- [变更日志](./docs/CHANGELOG.md)
- [Git Hooks](./docs/hooks.md)

## 配置

在使用 Merge Request 相关命令前，需要在项目根目录创建 `config.yaml` 文件：

```yaml
gitlab:
  url: https://your-gitlab-url
  token: your-personal-access-token
```

## 卸载

```bash
./git-toolkit uninstall
```

## License

MIT License
