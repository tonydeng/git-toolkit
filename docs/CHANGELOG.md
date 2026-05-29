# CHANGELOG

## git-toolkit-v2.1.1

### FixBugs-2.1.1

- 修复读取版本文件中的版本信息的`bug` #27

## git-toolkit-v2.1

### Features-2.1

- 添加了`git design`命令，创建以`design/`为前缀的架构设计分支 #23
- 添加了`git repl`命令替换历史`commit`的`email`和`author`信息 #22
- `git ci`命令内置针对远程仓库的作者邮箱的检查规则 #21
- `git ci`命令内置自动读取项目版本信息功能 #10

## git-toolkit-v2.0

### Features-2.0

- 使用`Golang`重构了`git ci` #2
- 丰富了`git-toolkit`各个命令的版本信息 #4
- 添加了基于`GitFlow`分支策略模型的创建分支命令 #5
- 添加了`push`当前分支的快捷命令`git ps` #7
- 添加`install`,`uninstall` #3
- 完成了规范持续集成、代码分析以及制品规范 #16
- 将`jira`与`gitlab`相结合，实现提交后`jira`对应的`issue`下面会有评论 #17
- 开发乐高标准文档 #13
- 完善安装文档和使用文档 #15

### FixBugs-2.0

- 修复执行`git-toolkit xx`命令时触发安装流程的`bug` #11
- 修复因为安装路径`/root/`导致普通用户无法使用的`bug` #9
- 修复执行`git-toolkit install`命令不成功的`bug` #12
