package cmd

import (
	"github.com/spf13/cobra"
	"github.com/tonydeng/git-toolkit/git"
	"github.com/tonydeng/git-toolkit/utils"
	"strings"
)

var flag bool

func NewRepl() *cobra.Command {
	newRepl := &cobra.Command{
		Use:     "git-repl",
		Short:   "replace the email and author",
		Long:    "replace the email and author in the historical commits",
		Version: utils.GenVersion(Version, BuildTime, CommitID),
		Args:    cobra.MinimumNArgs(2),
		Aliases: []string{"git-repl"},
		Run: func(cmd *cobra.Command, args []string) {
			l := len(args)
			flagForceOverWrite, _ := cmd.Flags().GetBool("flag")
			if strings.Index(args[0], ".com") != -1 {
				for i := 0; i < l-1; i++ {
					git.FilterBranch(args[i], args[l-1], "", "", true, flagForceOverWrite)
				}
			} else {
				for i := 0; i < l-1; i++ {
					git.FilterBranch(args[l-1], "", args[0], args[i], false, flagForceOverWrite)
				}
			}

		},
	}
	newRepl.Flags().BoolVarP(&flag, "flag", "f", false, "force overwriting the backup")
	return newRepl
}
