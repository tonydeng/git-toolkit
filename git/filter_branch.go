package git

import (
	"fmt"
	"github.com/tonydeng/git-toolkit/utils"
)

func FilterBranch(oldEmail string, correctEmail string, oldName string, correctName string, flag bool, flagFalseOverWrite bool) {
	shell := fmt.Sprintf(`
OLD_EMAIL="%s"
CORRECT_EMAIL="%s"
if [ "$GIT_COMMITTER_EMAIL" = "$OLD_EMAIL" ]
then
export GIT_COMMITTER_EMAIL="$CORRECT_EMAIL"
fi
if [ "$GIT_AUTHOR_EMAIL" = "$OLD_EMAIL" ]
then
export GIT_AUTHOR_EMAIL="$CORRECT_EMAIL"
fi
`, oldEmail, correctEmail)
	nameShell := fmt.Sprintf(`
OLD_NAME="%s"
CORRECT_NAME="%s"
OLD_EMAIL="%s"
if [ "$GIT_AUTHOR_NAME" = "$OLD_NAME" ] && [ "$GIT_AUTHOR_EMAIL" = "$OLD_EMAIL" ]
then
export GIT_AUTHOR_NAME="$CORRECT_NAME"
fi
`, oldName, correctName, oldEmail)
	if flagFalseOverWrite {
		if flag {
			utils.Exec(Cmd, "filter-branch", "-f", "--env-filter", shell, "--tag-name-filter", "cat", "--", "--branches", "--tags")
		} else {
			utils.Exec(Cmd, "filter-branch", "-f", "--env-filter", nameShell, "--tag-name-filter", "cat", "--", "--branches", "--tags")
		}
	} else {
		if flag {
			utils.Exec(Cmd, "filter-branch", "--env-filter", shell, "--tag-name-filter", "cat", "--", "--branches", "--tags")
		} else {
			utils.Exec(Cmd, "filter-branch", "--env-filter", nameShell, "--tag-name-filter", "cat", "--", "--branches", "--tags")
		}
	}

}
