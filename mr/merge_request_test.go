package mr

import (
	"fmt"
	"github.com/tonydeng/git-toolkit/utils"
	"testing"
)

func TestCreateNewMergeRequest(t *testing.T) {
	c := CreateNewMergeRequestOps{
		Token:              "token",
		Url:                utils.GetProjectUrl(),
		Title:              "t",
		Description:        "d",
		SourceBranch:       utils.GetCurrentBranchName(),
		TargetBranch:       "master",
		Label:              []string{"feature", "document"},
		AssigneeID:         1440,
		AssigneeIDs:        []int{},
		TargetProjectID:    1516,
		MilestoneID:        1,
		RemoveSourceBranch: false,
		Squash:             false,
		AllowCollaboration: false,
	}
	mergeRequest, response, err := CreateNewMergeRequest(c)
	fmt.Printf("The mergeRequest is {%s}", mergeRequest)
	fmt.Println(response)
	fmt.Println(err)
}

func TestUpdateMergeRequest(t *testing.T) {
	u := UpdateNewMergeRequestOps{
		Token:              "token",
		MergeRequestID:     28,
		Url:                utils.GetProjectUrl(),
		Title:              "d",
		Description:        "d",
		TargetBranch:       "master",
		Label:              []string{"feature"},
		AddLabels:          []string{},
		RemoveLabels:       []string{},
		AssigneeID:         1440,
		AssigneeIDs:        []int{},
		TargetProjectID:    1516,
		MilestoneID:        1,
		StateEvent:         "reopen",
		RemoveSourceBranch: false,
		Squash:             false,
		AllowCollaboration: false,
		DiscussionLocked:   false,
	}
	mergeRequest, response, err := UpdateNewMergeRequest(u)
	fmt.Printf("The mergeRequest is {%s}", mergeRequest)
	fmt.Println(response)
	fmt.Println(err)
}

func TestDeleteNewMergeRequest(t *testing.T) {
	d := DeleteNewMergeRequestOps{
		Token:           "token",
		Url:             utils.GetProjectUrl(),
		MergeRequestID:  28,
		TargetProjectID: 1516,
	}
	response, err := DeleteNewMergeRequest(d)
	fmt.Println(response)
	fmt.Println(err)
}

func TestAcceptNewMergeRequest(t *testing.T) {
	a := AcceptNewMergeRequestOps{
		Token:                     "token",
		Url:                       utils.GetProjectUrl(),
		MergeRequestID:            28,
		TargetProjectID:           1516,
		MergeCommitMessage:        "accept mr",
		SquashCommitMessage:       "",
		Squash:                    false,
		ShouldRemoveSourceBranch:  false,
		MergeWhenPipelineSucceeds: false,
		Sha:                       utils.GetSHA(),
	}
	mergeRequest, response, err := AcceptNewMergeRequest(a)
	fmt.Printf("The mergeRequest is {%s}", mergeRequest)
	fmt.Println(response)
	fmt.Println(err)
}

func TestGetMergeRequestChanges(t *testing.T) {
	g := GetMergeRequestChangesOps{
		Token:           "token",
		Url:             utils.GetProjectUrl(),
		MergeRequestID:  28,
		TargetProjectID: 1516,
	}
	mergeRequest, response, err := GetMergeRequestChanges(g)
	fmt.Printf("The mergeRequest is {%s}", mergeRequest)
	fmt.Println(response)
	fmt.Println(err)
}
