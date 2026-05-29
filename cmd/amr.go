package cmd

import (
	"github.com/spf13/cobra"
	"github.com/tonydeng/git-toolkit/mr"
	"github.com/tonydeng/git-toolkit/utils"
	"log"
)

func NewAmr() *cobra.Command {
	return &cobra.Command{
		Use:     "git-amr",
		Short:   "accept merge request",
		Long:    "accept merge request by git-toolkit",
		Version: utils.GenVersion(Version, BuildTime, CommitID),
		Run: func(cmd *cobra.Command, args []string) {
			mergeRequestMap := utils.ReadGitlabConfig("config.yaml")
			token := mergeRequestMap["token"]
			targetProjectID := utils.GetProjectID()
			a := mr.AcceptNewMergeRequestOps{
				Token:                     token,
				Url:                       utils.GetProjectUrl(),
				MergeRequestID:            mr.InputMergeRequestID(),
				TargetProjectID:           targetProjectID,
				MergeCommitMessage:        mr.InputMergeCommitMessage(),
				SquashCommitMessage:       mr.InputSquashCommitMessage(),
				Squash:                    mr.InputSquash(),
				ShouldRemoveSourceBranch:  mr.InputRemoveSourceBranch(),
				MergeWhenPipelineSucceeds: mr.InputMergeWhenPipelineSucceeds(),
				Sha:                       utils.GetSHA(),
			}
			_, _, e := mr.AcceptNewMergeRequest(a)
			if e != nil {
				log.Fatalf("The error is {%s}", e)
			}
		},
	}
}
