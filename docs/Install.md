# 下载 & 安装

## 下载

- [macOS](https://github.com/tonydeng/git-toolkit/releases/download/v2.1.1/git-toolkit_darwin_amd64)
- [Linux 64](https://github.com/tonydeng/git-toolkit/releases/download/v2.1.1/git-toolkit_linux_amd64)
- [Linux 32](https://github.com/tonydeng/git-toolkit/releases/download/v2.1.1/git-toolkit_linux_386)

## 查看版本

```bash
./git-toolkit_darwin_amd64 version
```

当你通过上面命令查看到如下信息时，说明下载成功。

```bash
 ██████╗ ██╗████████╗    ████████╗ ██████╗  ██████╗ ██╗     ██╗  ██╗██╗████████╗
██╔════╝ ██║╚══██╔══╝    ╚══██╔══╝██╔═══██╗██╔═══██╗██║     ██║ ██╔╝██║╚══██╔══╝
██║  ███╗██║   ██║          ██║   ██║   ██║██║   ██║██║     █████╔╝ ██║   ██║
██║   ██║██║   ██║          ██║   ██║   ██║██║   ██║██║     ██╔═██╗ ██║   ██║
╚██████╔╝██║   ██║          ██║   ╚██████╔╝╚██████╔╝███████╗██║  ██╗██║   ██║
 ╚═════╝ ╚═╝   ╚═╝          ╚═╝    ╚═════╝  ╚═════╝ ╚══════╝╚═╝  ╚═╝╚═╝   ╚═╝


Name: git-toolkit
Version: v2.1.1
Author: Tony Deng <wolf.deng@gmail.com>
Arch: darwin/amd64
BuildTime: 2020-06-24 12:26:49
```

## 安装

> 以下的安装以`mac`为例，其他操作系统场景类似

1.在终端输入如下命令修改权限

```bash
chmod 755 git-toolkit_darwin_amd64
```

2.查看支持的基本命令

```bash
./git-toolkit_darwin_amd64
```

```bash
一个用于提高Git命令使用效率和规范Git Commit Message的工具集

Usage:
  git-toolkit [flags]
  git-toolkit [command]

Available Commands:
  chore       Create chore branch
  design      Create design branch
  docs        Create docs branch
  feat        Create feature branch
  fix         Create fix branch
  git-amr     accept merge request
  git-ci      交互式输入 commit message
  git-clog    Git Change logs Output
  git-cm      Check the commit message specification
  git-cmr     create merge request
  git-dmr     delete merge request
  git-gmr     get merge request changes
  git-ps      Push local branch to remote git server
  git-repl    replace the email and author
  git-umr     update merge request
  help        Help about any command
  hotfix      Create hotfix branch
  install     Install git-toolkit
  perf        Create perf branch
  refactor    Create refactor branch
  style       Create style branch
  test        Create test branch
  uninstall   Uninstall git-toolkit
  version     Print version

Flags:
  -h, --help   help for git-toolkit

Use "git-toolkit [command] --help" for more information about a command.
```

3.安装git-toolkit

```bash
./git-tool_darwin_amd64 install
```

4.查看是否安装成功

当看到如下信息，说明安装成功

```bash
👉 remove /usr/local/git-toolkit
📥 mkdir /usr/local/git-toolkit/hooks
📥 copy file /usr/local/git-toolkit/git-toolkit
📥 install symbolic /usr/local/bin/git-ci
📥 install symbolic /usr/local/bin/git-cm
📥 install symbolic /usr/local/bin/git-feat
📥 install symbolic /usr/local/bin/git-fix
📥 install symbolic /usr/local/bin/git-design
📥 install symbolic /usr/local/bin/git-docs
📥 install symbolic /usr/local/bin/git-style
📥 install symbolic /usr/local/bin/git-refactor
📥 install symbolic /usr/local/bin/git-test
📥 install symbolic /usr/local/bin/git-chore
📥 install symbolic /usr/local/bin/git-pref
📥 install symbolic /usr/local/bin/git-hotfix
📥 install symbolic /usr/local/bin/git-ps
📥 install symbolic /usr/local/bin/git-repl
📥 install symbolic /usr/local/bin/git-clog
📥 install symbolic /usr/local/bin/git-cmr
📥 install symbolic /usr/local/bin/git-umr
📥 install symbolic /usr/local/bin/git-dmr
📥 install symbolic /usr/local/bin/git-amr
📥 install symbolic /usr/local/bin/git-gmr
📥 config set core.hooksPath /usr/local/git-toolkit/hooks
```

5.卸载

```bash
./git-toolkit_darwin_amd64 uninstall
```
