package cmd

import (
	"github.com/spf13/cobra"
	"github.com/tonydeng/git-toolkit/utils"
)

var cfgFile string
var RootCmd = &cobra.Command{
	Use:   "git-toolkit",
	Short: "Git工具集",
	Long:  "一个用于提高Git命令使用效率和规范Git Commit Message的工具集",
	Version: utils.GenVersion(Version, BuildTime, CommitID),
	Run: func(cmd *cobra.Command, args []string) {
		_ = cmd.Help()
	},
}

func init() {
	// 添加install & uninstall 命令
	installCmd := NewInstall()
	installCmd.PersistentFlags().StringVar(&installDir, "dir", "/usr/local/bin", "install dir")

	uninstallCmd := NewUninstall()
	uninstallCmd.PersistentFlags().StringVar(&installDir, "dir", "/usr/local/bin/", "install dir")

	//持久标志
	version := utils.GenVersion(Version, BuildTime, CommitID)
	RootCmd.SetVersionTemplate(version)

	// 添加自定义Git命令
	// 安装 & 卸载
	RootCmd.AddCommand(installCmd)
	RootCmd.AddCommand(uninstallCmd)

	// Commit Message相关命令
	RootCmd.AddCommand(NewCi())
	RootCmd.AddCommand(NewCm())

	// 创建分支相关命令
	RootCmd.AddCommand(NewFeatBranch())
	RootCmd.AddCommand(NewFixBranch())
	RootCmd.AddCommand(NewDocsBranch())
	RootCmd.AddCommand(NewStyleBranch())
	RootCmd.AddCommand(NewRefactorBranch())
	RootCmd.AddCommand(NewPerfBranch())
	RootCmd.AddCommand(NewHotFixBranch())
	RootCmd.AddCommand(NewTestBranch())
	RootCmd.AddCommand(NewChoreBranch())

	// 远程仓库相关命令
	RootCmd.AddCommand(NewPs())

	// 变更日志相关命令
	RootCmd.AddCommand(NewClog())

	// 历史提交修改相关命令
	RootCmd.AddCommand(NewRepl())

	// Merge Request相关命令
	RootCmd.AddCommand(NewCmr())
	RootCmd.AddCommand(NewUmr())
	RootCmd.AddCommand(NewDmr())
	RootCmd.AddCommand(NewAmr())
	RootCmd.AddCommand(NewGmr())

	// 版本信息
	RootCmd.AddCommand(NewVersion())
}
