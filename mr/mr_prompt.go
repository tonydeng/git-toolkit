package mr

import (
	"errors"
	"github.com/mritd/promptx"
	"github.com/tonydeng/git-toolkit/utils"
	"strconv"
	"strings"
)

type CreateNewMergeRequestOps struct {
	Token              string
	Url                string
	Title              string
	Description        string
	SourceBranch       string
	TargetBranch       string
	Label              []string
	AssigneeID         int
	AssigneeIDs        []int
	TargetProjectID    int
	MilestoneID        int
	RemoveSourceBranch bool
	Squash             bool
	AllowCollaboration bool
}

type UpdateNewMergeRequestOps struct {
	Token              string
	MergeRequestID     int
	Url                string
	Title              string
	Description        string
	TargetBranch       string
	Label              []string
	AddLabels          []string
	RemoveLabels       []string
	AssigneeID         int
	AssigneeIDs        []int
	TargetProjectID    int
	MilestoneID        int
	StateEvent         string
	RemoveSourceBranch bool
	Squash             bool
	AllowCollaboration bool
	DiscussionLocked   bool
}

type DeleteNewMergeRequestOps struct {
	Token           string
	Url             string
	MergeRequestID  int
	TargetProjectID int
}

type AcceptNewMergeRequestOps struct {
	Token                     string
	Url                       string
	MergeRequestID            int
	TargetProjectID           int
	MergeCommitMessage        string
	SquashCommitMessage       string
	Squash                    bool
	ShouldRemoveSourceBranch  bool
	MergeWhenPipelineSucceeds bool
	Sha                       string
}

type GetMergeRequestChangesOps struct {
	Token           string
	Url             string
	MergeRequestID  int
	TargetProjectID int
}

func InputTitle() string {
	p := promptx.NewDefaultPrompt(func(line []rune) error {
		if strings.TrimSpace(string(line)) == "" {
			return errors.New("Title is empty!\n")
		} else {
			return nil
		}
	}, "Title(请添加本次MR提交的标题)")
	return strings.TrimSpace(p.Run())
}

func InputDescription() string {
	p := promptx.NewDefaultPrompt(func(line []rune) error {
		if strings.TrimSpace(string(line)) == "" {
			return errors.New("Description is empty!\n")
		} else {
			return nil
		}
	}, "Description(请添加简短的、必要的对本次MR的描述)")
	return strings.TrimSpace(p.Run())
}

func InputTargetBranch() string {
	p := promptx.NewDefaultPrompt(func(line []rune) error {
		if strings.TrimSpace(string(line)) == "" {
			return errors.New("TargetBranch is empty!\n")
		} else {
			return nil
		}
	}, "TargetBranch(请添加本次MR的目标分支)")
	return strings.TrimSpace(p.Run())
}

func InputLabel() []string {
	p := promptx.NewDefaultPrompt(func(line []rune) error {
		if strings.TrimSpace(string(line)) == "" {
			return errors.New("Label is empty!\n")
		} else {
			return nil
		}
	}, "Label(请添加本次MR的标签)")
	return strings.Split(strings.TrimSpace(p.Run()), " ")
}

func InputAssigneeID() int {
	p := promptx.NewDefaultPrompt(func(line []rune) error {
		if strings.TrimSpace(string(line)) == "" {
			return errors.New("AssigneeID is empty!\n")
		} else {
			return nil
		}
	}, "AssigneeID(请添加本次MR的分配者)")
	username := strings.TrimSpace(p.Run())
	userInfo := utils.FetchUserInfo(username)
	return userInfo[0].ID
}

func InputAssigneeIDs() []int {
	p := promptx.NewDefaultPrompt(func(line []rune) error {
		return nil
	}, "AssigneeIDs(请添加本次MR的多个分配者)")
	usernames := strings.Split(strings.TrimSpace(p.Run()), " ")
	intIDs := make([]int, len(usernames))
	for index, username := range usernames {
		userInfo := utils.FetchUserInfo(username)
		intIDs[index] = userInfo[0].ID
	}
	return intIDs
}

func InputMilestoneID() int {
	p := promptx.NewDefaultPrompt(func(line []rune) error {
		return nil
	}, "MilestoneID(请添加本次MR的Milestone ID)")
	milestoneID, _ := strconv.Atoi(strings.TrimSpace(p.Run()))
	return milestoneID
}

func InputRemoveSourceBranch() bool {
	p := promptx.NewDefaultPrompt(func(line []rune) error {
		return nil
	}, "RemoveSourceBranch(请确定本次MR是否要移除源分支)")
	removeSourceBranch, _ := strconv.ParseBool(strings.TrimSpace(p.Run()))
	return removeSourceBranch
}

func InputSquash() bool {
	p := promptx.NewDefaultPrompt(func(line []rune) error {
		return nil
	}, "Squash(请确定本次MR是否要压缩commit)")
	squash, _ := strconv.ParseBool(strings.TrimSpace(p.Run()))
	return squash
}

func InputAllowCollaboration() bool {
	p := promptx.NewDefaultPrompt(func(line []rune) error {
		return nil
	}, "AllowCollaboration(请确定本次MR是否要采用合作模式)")
	allowCollaboration, _ := strconv.ParseBool(strings.TrimSpace(p.Run()))
	return allowCollaboration
}

func InputAddLabels() []string {
	p := promptx.NewDefaultPrompt(func(line []rune) error {
		return nil
	}, "AddLabels(请添加本次MR要新增的标签)")
	return strings.Split(strings.TrimSpace(p.Run()), " ")
}

func InputRemoveLabels() []string {
	p := promptx.NewDefaultPrompt(func(line []rune) error {
		return nil
	}, "RemoveLabels(请添加本次MR要删除的标签)")
	return strings.Split(strings.TrimSpace(p.Run()), " ")
}

func InputStateEvent() string {
	p := promptx.NewDefaultPrompt(func(line []rune) error {
		if strings.TrimSpace(string(line)) == "" {
			return errors.New("StateEvent is empty!\n")
		} else {
			return nil
		}
	}, "StateEvent(请添加本次MR状态，reopen or close)")
	return strings.TrimSpace(p.Run())
}

func InputDiscussionLocked() bool {
	p := promptx.NewDefaultPrompt(func(line []rune) error {
		return nil
	}, "DiscussionLocked(请确定本次MR是否要开启讨论权限)")
	discussionLocked, _ := strconv.ParseBool(strings.TrimSpace(p.Run()))
	return discussionLocked
}

func InputMergeRequestID() int {
	p := promptx.NewDefaultPrompt(func(line []rune) error {
		if strings.TrimSpace(string(line)) == "" {
			return errors.New("MergeRequestID is empty!\n")
		} else {
			return nil
		}
	}, "MergeRequestID(请添加MR的MergeRequestID)")
	mergeRequestID, _ := strconv.Atoi(strings.TrimSpace(p.Run()))
	return mergeRequestID
}

func InputMergeCommitMessage() string {
	p := promptx.NewDefaultPrompt(func(line []rune) error {
		if strings.TrimSpace(string(line)) == "" {
			return errors.New("MergeCommitMessage is empty!\n")
		} else {
			return nil
		}
	}, "MergeCommitMessage(请添加合并MR的备注)")
	return strings.TrimSpace(p.Run())
}

func InputSquashCommitMessage() string {
	p := promptx.NewDefaultPrompt(func(line []rune) error {
		return nil
	}, "SquashCommitMessage(请添加压缩的Commit的备注)")
	return strings.TrimSpace(p.Run())
}

func InputMergeWhenPipelineSucceeds() bool {
	p := promptx.NewDefaultPrompt(func(line []rune) error {
		return nil
	}, "MergeWhenPipelineSucceeds(请确定当流水线成功时MR是否会进行合并)")
	mergeWhenPipelineSucceeds, _ := strconv.ParseBool(strings.TrimSpace(p.Run()))
	return mergeWhenPipelineSucceeds
}
