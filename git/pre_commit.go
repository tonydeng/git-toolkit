package git

import (
	"fmt"
	"github.com/tonydeng/git-toolkit/utils"
	"os"
	"os/exec"
	"strings"
)

func VerifyRepositoryAndEmail(companyRepository string, companyEmail []string) {

	res, _ := exec.Command(Cmd, "remote").Output()
	repos := strings.Split(string(res), "\n")
	var flag bool
	for _, repo := range repos {
		command := exec.Command(Cmd, "remote", "get-url", repo)
		repositoryOrigin, _ := command.Output()
		flag := strings.Contains(string(repositoryOrigin), companyRepository)
		if flag {
			break
		}
	}

	command := exec.Command(Cmd, "config", "user.email")
	out, _ := command.Output()
	emailArray := strings.Split(string(out), "\n")
	email := emailArray[0]
	email = utils.TrimString(email, "@") 

	if flag {
		if !utils.IsContain(companyEmail, email) {
			fmt.Printf("When you push code to %s repository, your email should be a formal OA email!, but your email is %s \n", companyRepository, email)
			os.Exit(1)
		}
	}
}
