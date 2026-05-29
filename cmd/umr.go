package cmd

import (
	"github.com/spf13/cobra"
	"github.com/tonydeng/git-toolkit/mr"
	"github.com/tonydeng/git-toolkit/utils"
	"log"
)

func NewUmr() *cobra.Command {
	return &cobra.Command{
		Use:     "git-umr",
		Short:   "update merge request",
		Long:    "update merge request by git-toolkit",
		Version: utils.GenVersion(Version, BuildTime, CommitID),
		Run: func(cmd *cobra.Command, args []string) {
			mergeRequestMap := utils.ReadGitlabConfig("config.yaml")
			token := mergeRequestMap["token"]
			targetProjectID := utils.GetProjectID()
			u := mr.UpdateNewMergeRequestOps{
				Token:              token,
				MergeRequestID:     mr.InputMergeRequestID(),
				Url:                utils.GetProjectUrl(),
				Title:              mr.InputTitle(),
				Description:        mr.InputDescription(),
				TargetBranch:       mr.InputTargetBranch(),
				Label:              mr.InputLabel(),
				AddLabels:          mr.InputAddLabels(),
				RemoveLabels:       mr.InputRemoveLabels(),
				AssigneeID:         mr.InputAssigneeID(),
				AssigneeIDs:        mr.InputAssigneeIDs(),
				TargetProjectID:    targetProjectID,
				MilestoneID:        mr.InputMilestoneID(),
				StateEvent:         mr.InputStateEvent(),
				RemoveSourceBranch: mr.InputRemoveSourceBranch(),
				Squash:             mr.InputSquash(),
				AllowCollaboration: mr.InputAllowCollaboration(),
				DiscussionLocked:   mr.InputDiscussionLocked(),
			}

			_, _, e := mr.UpdateNewMergeRequest(u)
			if e != nil {
				log.Fatalf("The error is {%s}", e)
			}
		},
	}
}
