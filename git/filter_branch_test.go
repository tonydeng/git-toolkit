package git

import (
	"os"
	"os/exec"
	"strings"
	"testing"
)

func TestFilterBranch(t *testing.T) {
	oldEmail := "qq@qq.com"
	correctEmail := "user@example.com"
	_ = os.Chdir("..")
	FilterBranch(oldEmail, correctEmail,"","",true,true)
	cmd := exec.Command(Cmd, "log", "--pretty=format:%ae")
	email, _ := cmd.Output()
	elements := strings.Split(string(email), "\n")
	for _, elem := range elements {
		if elem == oldEmail {
			t.Errorf("modify was failed, there is still a %s email \n", elem)
		}
	}
}
