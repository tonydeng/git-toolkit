package git

import (
	"strings"

	"github.com/tonydeng/git-toolkit/utils"
)

type CommitType string

type RepoType string

const Cmd = "git"

const (
	FEAT     CommitType = "feat"
	FIX      CommitType = "fix"
	DOCS     CommitType = "docs"
	STYLE    CommitType = "style"
	REFACTOR CommitType = "refactor"
	TEST     CommitType = "test"
	CHORE    CommitType = "chore"
	PERF     CommitType = "perf"
	HOTFIX   CommitType = "hotfix"
	EXIT     CommitType = "exit"
)

// git emoji参考 https://gitmoji.carloscuesta.me/
var CommitEmoji = map[CommitType]string{
	FEAT:     ":sparkles:",
	FIX:      ":bug:",
	DOCS:     ":pencil:",
	STYLE:    ":art:",
	REFACTOR: ":hammer:",
	TEST:     ":white_check_mark:",
	CHORE:    ":wrench:",
	PERF:     ":zap:",
	HOTFIX:   ":ambulance:",
}

const CommitTpl = `{{ .Emoji }} {{ .Type }}({{ .Scope }}): {{ .Subject }}

{{ .Body }}

{{ .Footer }}

{{ .Sob }}
`

const CommitMessagePattern = `^(?:fixup!\s*)?(\w*)(\(([\w\$\.\*/-].*)\))?\: (.*)|^Merge\ branch(.*)`

func CheckGitProject() {
	utils.ExecNoOut(Cmd, "rev-parse", "--show-toplevel")
}

func CheckStageFiles() bool {
	output := utils.ExecRtOut(Cmd, "diff", "--cached", "--name-only")
	return strings.TrimSpace(output) != ""
}

func CheckLastCommitInfo() *[]string {
	title := utils.ExecRtOut(Cmd, "log", "-1", "--pretty=format:%s")
	desc := utils.ExecRtOut(Cmd, "log", "-1", "--pretty=format:%b")
	return &[]string{title, desc}
}

func Checkout(prefix CommitType, branch string) {
	utils.Exec(Cmd, "checkout", "-b", string(prefix)+"/"+branch)
}

func GetCurrentBranch() string {
	// 1.8+ git symbolic-ref --short HEAD
	return strings.TrimSpace(utils.ExecRtOut(Cmd, "rev-parse", "--abbrev-ref", "HEAD"))
}

func Push() {
	utils.Exec(Cmd, "push", "origin", strings.TrimSpace(GetCurrentBranch()))
}
