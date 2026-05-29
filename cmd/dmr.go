package cmd

import (
	"github.com/spf13/cobra"
	"github.com/tonydeng/git-toolkit/mr"
	"github.com/tonydeng/git-toolkit/utils"
	"log"
)

func NewDmr() *cobra.Command {
	return &cobra.Command{
		Use:     "git-dmr",
		Short:   "delete merge request",
		Long:    "delete merge request by git-toolkit",
		Version: utils.GenVersion(Version, BuildTime, CommitID),
		Run: func(cmd *cobra.Command, args []string) {
			mergeRequestMap := utils.ReadGitlabConfig("config.yaml")
			token := mergeRequestMap["token"]
			targetProjectID := utils.GetProjectID()
			d := mr.DeleteNewMergeRequestOps{
				Token:           token,
				Url:             utils.GetProjectUrl(),
				MergeRequestID:  mr.InputMergeRequestID(),
				TargetProjectID: targetProjectID,
			}
			_, e := mr.DeleteNewMergeRequest(d)
			if e != nil {
				log.Fatalf("The error is {%s}", e)
			}
		},
	}
}
