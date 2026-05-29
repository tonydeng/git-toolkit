package cmd

import (
	"github.com/spf13/cobra"
	"github.com/tonydeng/git-toolkit/mr"
	"github.com/tonydeng/git-toolkit/utils"
	"log"
)

func NewGmr() *cobra.Command {
	return &cobra.Command{
		Use:     "git-gmr",
		Short:   "get merge request changes",
		Long:    "get merge request changes by git-toolkit",
		Version: utils.GenVersion(Version, BuildTime, CommitID),
		Run: func(cmd *cobra.Command, args []string) {
			mergeRequestMap := utils.ReadGitlabConfig("config.yaml")
			token := mergeRequestMap["token"]
			targetProjectID := utils.GetProjectID()
			g := mr.GetMergeRequestChangesOps{
				Token:           token,
				Url:             utils.GetProjectUrl(),
				MergeRequestID:  mr.InputMergeRequestID(),
				TargetProjectID: targetProjectID,
			}
			_, _, e := mr.GetMergeRequestChanges(g)
			if e != nil {
				log.Fatalf("The error is {%s}", e)
			}
		},
	}
}
