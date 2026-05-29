package cmd

import (
	"github.com/spf13/cobra"
	"github.com/tonydeng/git-toolkit/mr"
	"github.com/tonydeng/git-toolkit/utils"
	"log"
)

func NewCmr() *cobra.Command {
	return &cobra.Command{
		Use:     "git-cmr",
		Short:   "create merge request",
		Long:    "create merge request by git-toolkit",
		Version: utils.GenVersion(Version, BuildTime, CommitID),
		Run: func(cmd *cobra.Command, args []string) {
			mergeRequestMap := utils.ReadGitlabConfig("config.yaml")
			token := mergeRequestMap["token"]
			targetProjectID := utils.GetProjectID()
			c := mr.CreateNewMergeRequestOps{
				Token:              token,
				Url:                utils.GetProjectUrl(),
				Title:              mr.InputTitle(),
				Description:        mr.InputDescription(),
				SourceBranch:       utils.GetCurrentBranchName(),
				TargetBranch:       mr.InputTargetBranch(),
				Label:              mr.InputLabel(),
				AssigneeID:         mr.InputAssigneeID(),
				AssigneeIDs:        mr.InputAssigneeIDs(),
				TargetProjectID:    targetProjectID,
				MilestoneID:        mr.InputMilestoneID(),
				RemoveSourceBranch: mr.InputRemoveSourceBranch(),
				Squash:             mr.InputSquash(),
				AllowCollaboration: mr.InputAllowCollaboration(),
			}

			_, _, e := mr.CreateNewMergeRequest(c)
			if e != nil {
				log.Fatalf("The error is {%s}", e)
			}
		},
	}
}
