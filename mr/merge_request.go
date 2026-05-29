package mr

import (
	"github.com/xanzy/go-gitlab"
	"log"
)

func CreateNewMergeRequest(ops CreateNewMergeRequestOps) (*gitlab.MergeRequest, *gitlab.Response, error) {

	client, err := gitlab.NewClient(ops.Token, gitlab.WithBaseURL(ops.Url))
	if err != nil {
		log.Fatal(err)
	}
	c := &gitlab.CreateMergeRequestOptions{
		Title:              gitlab.String(ops.Title),
		Description:        gitlab.String(ops.Description),
		SourceBranch:       gitlab.String(ops.SourceBranch),
		TargetBranch:       gitlab.String(ops.TargetBranch),
		Labels:             ops.Label,
		AssigneeID:         gitlab.Int(ops.AssigneeID),
		AssigneeIDs:        ops.AssigneeIDs,
		TargetProjectID:    gitlab.Int(ops.TargetProjectID),
		MilestoneID:        gitlab.Int(ops.MilestoneID),
		RemoveSourceBranch: gitlab.Bool(ops.RemoveSourceBranch),
		Squash:             gitlab.Bool(ops.Squash),
		AllowCollaboration: gitlab.Bool(ops.AllowCollaboration),
	}
	mr, response, e := client.MergeRequests.CreateMergeRequest(ops.TargetProjectID, c)
	return mr, response, e
}

func UpdateNewMergeRequest(ops UpdateNewMergeRequestOps) (*gitlab.MergeRequest, *gitlab.Response, error) {

	client, err := gitlab.NewClient(ops.Token, gitlab.WithBaseURL(ops.Url))
	if err != nil {
		log.Fatal(err)
	}
	u := &gitlab.UpdateMergeRequestOptions{
		Title:              gitlab.String(ops.Title),
		Description:        gitlab.String(ops.Description),
		TargetBranch:       gitlab.String(ops.TargetBranch),
		AssigneeID:         gitlab.Int(ops.AssigneeID),
		AssigneeIDs:        ops.AssigneeIDs,
		Labels:             ops.Label,
		AddLabels:          ops.AddLabels,
		RemoveLabels:       ops.RemoveLabels,
		MilestoneID:        gitlab.Int(ops.MilestoneID),
		StateEvent:         gitlab.String(ops.StateEvent),
		RemoveSourceBranch: gitlab.Bool(ops.RemoveSourceBranch),
		Squash:             gitlab.Bool(ops.Squash),
		DiscussionLocked:   gitlab.Bool(ops.DiscussionLocked),
		AllowCollaboration: gitlab.Bool(ops.AllowCollaboration),
	}
	mr, response, e := client.MergeRequests.UpdateMergeRequest(ops.TargetProjectID, ops.MergeRequestID, u)
	return mr, response, e
}

func DeleteNewMergeRequest(ops DeleteNewMergeRequestOps) (*gitlab.Response, error) {

	client, err := gitlab.NewClient(ops.Token, gitlab.WithBaseURL(ops.Url))
	if err != nil {
		log.Fatal(err)
	}
	response, e := client.MergeRequests.DeleteMergeRequest(ops.TargetProjectID, ops.MergeRequestID)
	return response, e
}

func AcceptNewMergeRequest(ops AcceptNewMergeRequestOps) (*gitlab.MergeRequest, *gitlab.Response, error) {

	client, err := gitlab.NewClient(ops.Token, gitlab.WithBaseURL(ops.Url))
	if err != nil {
		log.Fatal(err)
	}
	a := &gitlab.AcceptMergeRequestOptions{
		MergeCommitMessage:        gitlab.String(ops.MergeCommitMessage),
		SquashCommitMessage:       gitlab.String(ops.SquashCommitMessage),
		Squash:                    gitlab.Bool(ops.Squash),
		ShouldRemoveSourceBranch:  gitlab.Bool(ops.ShouldRemoveSourceBranch),
		MergeWhenPipelineSucceeds: gitlab.Bool(ops.MergeWhenPipelineSucceeds),
		SHA:                       gitlab.String(ops.Sha),
	}
	mergeRequest, response, e := client.MergeRequests.AcceptMergeRequest(ops.TargetProjectID, ops.MergeRequestID, a)
	return mergeRequest, response, e
}

func GetMergeRequestChanges(ops GetMergeRequestChangesOps) (*gitlab.MergeRequest, *gitlab.Response, error) {

	client, err := gitlab.NewClient(ops.Token, gitlab.WithBaseURL(ops.Url))
	if err != nil {
		log.Fatal(err)
	}
	opt := &gitlab.GetMergeRequestChangesOptions{}
	mergeRequest, response, e := client.MergeRequests.GetMergeRequestChanges(ops.TargetProjectID, ops.MergeRequestID, opt)
	return mergeRequest, response, e
}
