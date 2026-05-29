package git

import (
	"fmt"
	"github.com/tonydeng/git-toolkit/utils"
	"os"
	"strconv"
	"strings"
)

type Author struct {
	Name         string
	Email        string
	CommitNumber int
}

type CurrentTag struct {
	index          int
	commitIdSha1   string
	projectVersion string
	version        string
}

type CommitInfo struct {
	AuthorInfo    string
	CommitDetails string
}

func GetRemoteUrl() string {
	origin := strings.TrimSpace(utils.ExecRtOut(Cmd, "remote", "get-url", "origin"))
	runes := []rune(origin)
	origin = string(runes[0:strings.LastIndex(origin, ".")])
	if strings.HasPrefix(origin, "git@") {
		_arr := strings.Split(strings.Split(origin, "@")[1], ":")
		origin = "https://" + _arr[0] + "/" + _arr[1]
	}
	return origin
}

func GetOldestAndNewestTags() (string, string) {
	allTags := GetAllTag()
	oldestTag := strings.Split(strings.TrimSpace(allTags[len(allTags)-1]), "|")[1]
	newestTag := strings.Split(strings.TrimSpace(allTags[0]), "|")[1]
	return oldestTag, newestTag
}

func GetTagRange(tagV string) string {
	allTags := GetAllTag()
	oldestTag, newestTag := GetOldestAndNewestTags()
	if strings.Index(tagV, "..") == -1 {
		tagV = SetTag(allTags, tagV)
	} else if strings.Index(tagV, "..") == 0 && strings.Split(tagV, "..")[1] != oldestTag {
		tagV = oldestTag + tagV
	} else if strings.LastIndex(tagV, "..") == len(tagV)-2 && strings.Split(tagV, "..")[0] != newestTag {
		tagV = tagV + newestTag
	} else if strings.Index(tagV, "..") == 0 && strings.Split(tagV, "..")[1] == oldestTag {
		tagV = oldestTag
	}
	return tagV
}

func GetAllAuthor(tagV string) []Author {
	tagV = GetTagRange(tagV)
	args := []string{"shortlog", "-s", "-n", "-e", "--no-merges", tagV}
	var authors []Author
	info := utils.MustExecRtOut(Cmd, args...)

	lines := strings.Split(info, "\n")
	for _, line := range lines {
		if len(strings.TrimSpace(line)) > 0 {
			author := GetAuthor(line)
			authors = append(authors, author)
		}
	}
	return authors
}

func GetAuthor(line string) Author {
	authorInfo := strings.Fields(strings.TrimSpace(line))
	emailIndex := len(authorInfo) - 1

	email := strings.TrimSpace(authorInfo[emailIndex])
	email = utils.SubString(email, 1, len(email)-1)
	author := strings.Join(authorInfo[1:emailIndex], " ")
	number, err := strconv.Atoi(authorInfo[0])

	if err != nil {
		fmt.Printf("%s parse commit number %s error!\n", line, authorInfo[0])
	}

	return Author{
		Name:         author,
		Email:        email,
		CommitNumber: number,
	}

}

func GetProjectName(tagV string) string {
	projectName := utils.ExecRtOut(Cmd, "remote", "get-url", "origin")
	runes := []rune(projectName)
	projectName = string(runes[0:strings.LastIndex(projectName, ".")])
	projectName = strings.Split(projectName, "/")[len(strings.Split(projectName, "/"))-1]
	return "# " + projectName + "-" + tagV + "\n\n"
}

func GenAuthorInfo(author Author) string {
	name := author.Name
	email := author.Email
	commitNumber := strconv.Itoa(author.CommitNumber)
	return "## [" + name + "](" + email + ")(" + commitNumber + ")\n"
}

func GenChangeLog(tagV string) []string {
	repoUrl := GetRemoteUrl()
	tagV = GetTagRange(tagV)
	authors := GetAllAuthor(tagV)

	var changelogs []string
	changelogs = append(changelogs, GetProjectName(tagV))
	for i, author := range authors {
		commitInfo := CommitInfo{
			AuthorInfo: GenAuthorInfo(author),
		}
		commitInfo.CommitDetails = utils.MustExecRtOut(Cmd, "log", "--author="+author.Email,
			"--no-merges", tagV,
			"--pretty=format:* %s (%ad) [view commit %h]("+repoUrl+"/commit/%h)",
			"--date=format:%Y-%m-%d %H:%M:%S")

		changelogs = append(changelogs, commitInfo.AuthorInfo, "\n")

		if i != len(authors)-1 {
			changelogs = append(changelogs, commitInfo.CommitDetails, "\n", "\n")
		} else {
			changelogs = append(changelogs, commitInfo.CommitDetails, "\n")
		}

	}
	return changelogs
}

func GetAllTag() []string {
	allTags := strings.TrimSpace(utils.ExecRtOut(Cmd, "for-each-ref",
		`--sort=version:refname`, `--format='%(objectname)|%(refname:short)'`, "refs/tags/"))
	var tags []string
	lines := strings.Split(allTags, "\n")
	for _, line := range lines {
		line = utils.SubString(line, 1, len(line)-1)
		tags = append(tags, line)
	}
	return utils.ReverseArray(tags)
}

func GetTag(tags []string, tag string) CurrentTag {
	var output CurrentTag
	if strings.Index(tag, "-") != -1 {
		tag = utils.SubString(tag, strings.LastIndex(tag, "-")+1, len(tag))
	}
	ver := strings.Split(tag, "|")[len(strings.Split(tag, "|"))-1]
	for index, elem := range tags {
		commitId := utils.SubString(elem, strings.LastIndex(elem, "|")+1, len(elem))
		flag := ver == commitId
		if flag {
			output.index = index
			output.commitIdSha1 = utils.SubString(elem, 0, strings.Index(elem, "|"))
			output.projectVersion = utils.SubString(elem, strings.Index(elem, "|")+1, len(elem))
			output.version = utils.SubString(elem, strings.LastIndex(elem, "|")+1, len(elem))
		}
	}
	return output
}

func SetTag(tags []string, tag string) string {
	var beforeTag CurrentTag
	var versionParam string
	currentTag := GetTag(tags, tag)
	index := currentTag.index
	index = index + 1
	l := len(tags)
	if index < l {
		beforeTag = GetTag(tags, tags[index])
		versionParam = beforeTag.version + ".." + currentTag.version
	} else {
		versionParam = currentTag.version
	}
	return versionParam
}

func Help() string {
	output := fmt.Sprintf(`
%s
     git-clog - Git Change Logs Output
%s
     git clog [<options>] [<revision range>] [[--] <path>...]
%s
     Summarizes git log output in a format suitable for inclusion in release announcements. Each commit will be grouped by author and title.
%s
     -o, --output
         The changelog output to a file.
     -v, --version
         Specify the output one version.
     -h, --help
         Prints the synopsis and a list of the most commonly used commands.
`,
		"NAME", "SYNOPSIS", "DESCRIPTION", "OPTIONS")
	return output
}

func VerifyGitVersion() {
	version := strings.TrimSpace(utils.ExecRtOut(Cmd, "--version"))
	arr := strings.Split(version, " ")
	arr = strings.Split(arr[2], ".")
	version = arr[0] + "." + arr[1]
	ver, _ := strconv.ParseFloat(version, 64)
	if ver < 2.1 {
		fmt.Printf("The current Git version is %f, less than 2.1,please upgrade to git 2.1 above", ver)
		os.Exit(1)
	}
}
