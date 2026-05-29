package git

import (
	"fmt"
	"github.com/tonydeng/git-toolkit/utils"
	"strings"
	"testing"
)

func TestVerifyRepositoryAndEmail(t *testing.T) {
	emails := []string{"qq.com", "example.com"}
	VerifyRepositoryAndEmail("github.com", emails)
	companyRepository := "github.com"
	repositoryOrigin := "https://github.com/tonydeng/git-toolkit.git"
	repositoryOriginArray := strings.Split(string(repositoryOrigin), "/")
	repository := repositoryOriginArray[2]
	repository = utils.TrimString(repository, ".") 
	flag := companyRepository == repository
	fmt.Println(flag)

	fmt.Println(strings.Contains(repositoryOrigin,companyRepository))
}
