package git

import (
	"fmt"
	"github.com/tonydeng/git-toolkit/utils"
	"strconv"
	"strings"
	"testing"
)

func TestGetRemoteUrl(t *testing.T) {
	url := "git@github.com:tonydeng/git-toolkit.git"
	fmt.Printf("%s %d\n", url, strings.Index(url, "https://"))

	if strings.HasPrefix(url, "git@") {
		arr := strings.Split(strings.Split(url, "@")[1], ":")
		url = "https://" + arr[0] + "/" + arr[1]
	}
	fmt.Printf("%s\n", url)

	fmt.Println(GetRemoteUrl())

}

func TestGetAllAuthor(t *testing.T) {

	authors := GetAllAuthor("v2.0")
	fmt.Println(authors)

	for _, author := range authors {
		fmt.Printf("email %s \n", author.Email)
		fmt.Printf("CommitNumber %d \n", author.CommitNumber)
		fmt.Printf("Name %s \n", author.Name)
	}

	args := []string{"shortlog", "-s", "-n", "-e", "--no-merges"}
	shortLog := utils.MustExecRtOut(Cmd, args...)

	fmt.Printf("%s\n", shortLog)
}

func TestGenChangelog(t *testing.T) {
	repoUrl := GetRemoteUrl()
	fmt.Println(repoUrl)

	authors := GetAllAuthor("v2.1")
	for _, author := range authors {
		fmt.Printf("name:'%s' email:'%s' number:'%v'\n", author.Name, author.Email, author.CommitNumber)
	}
	for _, author := range authors {
		log := utils.MustExecRtOut(Cmd, "log", "--author='"+author.Email+"'",
			"--no-merges", "v2.1",
			"--pretty=format:'* %s (%ad) [view commit %h]("+repoUrl+"/commit/%h)",
			"--date=format:'%Y-%m-%d %H:%M:%S'")

		arg := "--author='" + author.Email + "' --no-merges 'v2.1' --pretty=format:'* %s (%ad) [view commit %h](https://github.com/tonydeng/git-toolkit/commit/%h)' --date=format:'%Y-%m-%d %H:%M:%S'"

		fmt.Println(log)

		output := utils.MustExecRtOut(Cmd, "log", arg)
		fmt.Println(output)

	}
	fmt.Println(GenChangeLog("v2.1"))
	utils.GenChangeLogSaveToFile("Change.md",GenChangeLog("v2.1"))
}

func TestHelp(t *testing.T) {
	tag := "abc"

	fmt.Println(tag + "def")
	fmt.Println(tag)
	output := Help()
	fmt.Println(output)
	if output == "" {
		t.Errorf("Help: something is wrong {%s}", output)
	}
}

func TestGetAllTag(t *testing.T) {
	result := GetAllTag()
	fmt.Println(result)
}

func TestGetTag(t *testing.T) {
	tags := []string{"1b6ce61f9b608f1d1c4b066482e3421c59e1c45c|v2.1",
	"fb92cd16cc324bd041a7c508b5f4b810d5158909|v2.1.1"}
	tag := "1b6ce61f9b608f1d1c4b066482e3421c59e1c45c|v2.1"
	current := GetTag(tags, tag)
	fmt.Println(current.commitIdSha1)
	fmt.Println(current.version)
	if current.version != "v2.1" {
		t.Errorf("GetTag: something is wrong with version id {%s} \n", current.version)
	}
}

func TestSetTag(t *testing.T) {
	tags := []string{"fb92cd16cc324bd041a7c508b5f4b810d5158909|v2.1.1",
		"1b6ce61f9b608f1d1c4b066482e3421c59e1c45c|v2.1",
		"3d811f57abdc203a76028d719ef289e8e7a9c31c|v2.0"}
	tag1 := "3d811f57abdc203a76028d719ef289e8e7a9c31c|v2.0"
	param1 := SetTag(tags, tag1)
	fmt.Println(param1)
	if param1 != "v2.0" {
		t.Errorf("SetTag: something is wrong with param1 {%s} \n", param1)
	}

	tag2 := "1b6ce61f9b608f1d1c4b066482e3421c59e1c45c|v2.1"
	param2 := SetTag(tags, tag2)
	fmt.Println(param2)
	if param2 != "v2.0..v2.1" {
		t.Errorf("SetTag: something is wrong with param2 {%s} \n", param2)
	}

	tag3 := "fb92cd16cc324bd041a7c508b5f4b810d5158909|v2.1.1"
	param3 := SetTag(tags, tag3)
	fmt.Println(param3)
	if param3 != "v2.1..v2.1.1" {
		t.Errorf("SetTag: something is wrong with param3 {%s} \n", param3)
	}
}

func TestGetAuthor(t *testing.T) {
	line := " 169  Joey <user@example.com>  "
	author := GetAuthor(line)
	fmt.Println(author)
	fmt.Println(author.Name)
	fmt.Println(author.CommitNumber)
	fmt.Println(author.Email)
	if author.Name != "Joey" {
		t.Errorf("GetAuthor: something is wrong with author.name {%s} \n", author.Name)
	}
}

func TestVerifyGitVersion(t *testing.T) {
	VerifyGitVersion()
	version := strings.TrimSpace(utils.ExecRtOut(Cmd, "--version"))
	arr := strings.Split(version, " ")
	arr = strings.Split(arr[2], ".")
	version = arr[0] + "." + arr[1]
	ver, _ := strconv.ParseFloat(version, 64)
	fmt.Println(strings.Index("v2..", ".."))
	if ver < 2.1 {
		t.Errorf("TestVerifyGitVersion:something is wrong with ver{%f} \n", ver)
	} else {
		fmt.Println("ver >= 2.1")
	}

}

func TestGetOldestAndNewestTags(t *testing.T) {
	old, newest := GetOldestAndNewestTags()
	fmt.Println(old)
	fmt.Println(newest)
}

func TestGetTagRange(t *testing.T) {
	tags := []string{"..v2.1", "v2.1..", "v2.1"}
	for _, str := range tags {
		fmt.Println(GetTagRange(str))
		fmt.Println(strings.Split("..v2.1", "..")[1])
	}
}

func TestGetProjectName(t *testing.T) {
	projectName := GetProjectName("v2.1")
	fmt.Println(projectName)
	if projectName == "" {
		t.Errorf("TestGetProjectName:Something is wrong with projectName:{%s}", projectName)
	}
}

func TestGenAuthorInfo(t *testing.T) {
	author:=Author{
		Name:         "me",
		Email:        "163",
		CommitNumber: 0,
	}
	authorInfo:=GenAuthorInfo(author)
	fmt.Println(authorInfo)
	if authorInfo=="" {
		t.Errorf("TestGenAuthorInfo:Something is wrong with authorInfo:{%s}", authorInfo)
	}
}
