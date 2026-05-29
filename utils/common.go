package utils

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"os/exec"
	"os/user"
	"path/filepath"
	"runtime"
	"strings"
)

const Cmd = "git"

var GitToolkitHome string
var InstallPath string
var HooksPath string
var GitCMHookPath string
var CurrentPath string

func init() {

	var err error

	home := "/usr/local"
	CheckAndExit(err)

	GitToolkitHome = filepath.Join(home, "git-toolkit")
	InstallPath = filepath.Join(GitToolkitHome, "git-toolkit")
	HooksPath = filepath.Join(GitToolkitHome, "hooks")
	GitCMHookPath = filepath.Join(HooksPath, "commit-msg")

	CurrentPath, err = exec.LookPath(os.Args[0])

	CheckAndExit(err)
}

func IskErr(err error) bool {
	if err != nil {
		fmt.Println(err)
		return false
	}
	return true
}

func CheckAndExit(err error) {
	if !IskErr(err) {
		os.Exit(1)
	}
}

func MustExec(name string, arg ...string) {
	cmd := exec.Command(name, arg...)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Run()
}

func Exec(name string, arg ...string) {
	cmd := exec.Command(name, arg...)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	CheckAndExit(cmd.Run())
}

func ExecRtOut(name string, arg ...string) string {
	cmd := exec.Command(name, arg...)
	cmd.Stdin = os.Stdin
	cmd.Stderr = os.Stderr

	b, err := cmd.Output()

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	return string(b)
}

func ExecNoOut(name string, arg ...string) {
	cmd := exec.Command(name, arg...)
	cmd.Stderr = os.Stderr
	CheckAndExit(cmd.Run())
}

func TryExec(name string, arg ...string) error {
	cmd := exec.Command(name, arg...)
	return cmd.Run()
}

func Root() bool {
	u, err := user.Current()
	CheckAndExit(err)
	return u.Uid == "0" || u.Gid == "0"
}

func OSEditInput() string {
	f, err := os.CreateTemp("", "git-toolkit")
	CheckAndExit(err)

	defer func() {
		_ = f.Close()
		_ = os.Remove(f.Name())
	}()

	bom := []byte{0xef, 0xbb, 0xbf}
	_, err = f.Write(bom)
	CheckAndExit(err)

	editor := "vim"
	if runtime.GOOS == "windows" {
		editor = "notepad"
	}

	if v := os.Getenv("VISUAL"); v != "" {
		editor = v
	} else if e := os.Getenv("EDITOR"); e != "" {
		editor = e
	}

	Exec(editor, f.Name())

	raw, err := os.ReadFile(f.Name())
	CheckAndExit(err)
	input := string(bytes.TrimPrefix(raw, bom))

	return input
}

func CheckOS() {
	if runtime.GOOS != "linux" && runtime.GOOS != "darwin" {
		fmt.Println("Platform not support!")
		os.Exit(1)
	}
}

func GenBinPaths(dir string) []string {
	return []string{
		filepath.Join(dir, "git-ci"),
		filepath.Join(dir, "git-cm"),
		filepath.Join(dir, "git-feat"),
		filepath.Join(dir, "git-fix"),
		filepath.Join(dir, "git-docs"),
		filepath.Join(dir, "git-style"),
		filepath.Join(dir, "git-refactor"),
		filepath.Join(dir, "git-test"),
		filepath.Join(dir, "git-chore"),
		filepath.Join(dir, "git-perf"),
		filepath.Join(dir, "git-hotfix"),
		filepath.Join(dir, "git-ps"),
		filepath.Join(dir, "git-cmr"),
		filepath.Join(dir, "git-umr"),
		filepath.Join(dir, "git-dmr"),
		filepath.Join(dir, "git-amr"),
		filepath.Join(dir, "git-gmr"),
	}
}

func IsContain(items []string, item string) bool {
	for _, eachItem := range items {
		if eachItem == item {
			return true
		}
	}
	return false
}

func MustExecRtOut(name string, arg ...string) string {
	cmd := exec.Command(name, arg...)
	cmd.Stdin = os.Stdin
	cmd.Stderr = os.Stderr
	b, _ := cmd.Output()
	return string(b)
}

func TrimString(str string, sign string) string {
	arr := strings.Split(str, sign)
	var result strings.Builder
	if len(arr) <= 1 {
		return str
	}
	for i := 1; i < len(arr); i++ {
		result.WriteString(arr[i])
		if i != len(arr)-1 {
			result.WriteString(sign)
		}
	}
	return result.String()
}

func SubString(source string, start int, end int) string {
	var r = []rune(source)
	length := len(r)
	if start < 0 || end > length || start > end {
		return ""
	}

	var substring = ""
	for i := start; i < end; i++ {
		substring += string(r[i])
	}

	return substring
}

func ReverseArray(array []string) []string {
	for i, j := 0, len(array)-1; i < j; i, j = i+1, j-1 {
		array[i], array[j] = array[j], array[i]
	}
	return array
}

func FileExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsExist(err) {
		return true, nil
	} else {
		return false, err
	}
}

func GenChangeLogSaveToFile(path string, content []string) {
	var f *os.File
	var err error
	fileExistsFlag, _ := FileExists(path)
	if fileExistsFlag {
		f, _ = os.OpenFile(path, os.O_APPEND|os.O_WRONLY, 0777)
	} else {
		dir, filename := filepath.Split(path)
		e := os.MkdirAll(dir, 0766)
		if e != nil {
			fmt.Errorf("something is wrong with {%s}", path)
		}
		t := dir + filename
		f, err = os.Create(t)
		if err != nil {
			fmt.Println(err)
		}
	}
	defer f.Close()
	for _, info := range content {
		_, err = io.WriteString(f, info)
		if err != nil {
			log.Fatal(err)
		}
	}
}

func ReadVersionFile(filename string) string {
	file, err := os.Open(filename)
	if err != nil {
		fmt.Printf("There is a error %s \n", err)
	}
	defer file.Close()
	line := bufio.NewReader(file)
	content, _, err := line.ReadLine()
	if err != nil {
		fmt.Printf("there is a error %s \n", err)
		os.Exit(1)
	}
	return string(content)
}

func GetCurrentBranchName() string {
	currentBranchName := ExecRtOut(Cmd, "rev-parse", "--abbrev-ref", "HEAD")
	return strings.TrimSpace(currentBranchName)
}

func GetSHA() string {
	sha := ExecRtOut(Cmd, "rev-parse", "HEAD")
	return sha
}

func GetProjectUrl() string {
	command := exec.Command(Cmd, "remote", "get-url", "origin")
	repositoryOrigin, _ := command.Output()
	repositoryOriginArray := strings.Split(string(repositoryOrigin), "/")
	repository := repositoryOriginArray[2]
	repositoryUrl := "https://" + repository + "/api/v4"
	return repositoryUrl
}

func GetProjectName() string {
	dir := ExecRtOut(Cmd, "rev-parse", "--show-toplevel")
	_, projectName := filepath.Split(dir)
	return projectName
}

type UserInfo struct {
	ID        int    `json:"id"`
	Name      string `json:"name"`
	Username  string `json:"username"`
	State     string `json:"state"`
	AvatarUrl string `json:"avatar_url"`
	WebUrl    string `json:"web_url"`
}

type ProjectInfo struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

func FetchUserInfo(username string) []UserInfo {
	url := GetProjectUrl() + "/users?username=" + username
	var userInfoByte []byte
	if response, err := http.Get(url); err != nil {
		fmt.Println(err)
	} else {
		if userInfoByte, err = io.ReadAll(response.Body); err != nil {
			fmt.Println(err)
		}
	}
	var u []UserInfo
	err := json.Unmarshal(userInfoByte, &u)
	if err != nil {
		log.Fatal(err)
	}
	return u
}

func FetchProjectInfo() []ProjectInfo {
	projectName := strings.TrimSpace(GetProjectName())
	url := GetProjectUrl() + "/projects?search=" + projectName
	var projectInfoByte []byte
	if response, err := http.Get(url); err != nil {
		fmt.Println(err)
	} else {
		if projectInfoByte, err = io.ReadAll(response.Body); err != nil {
			fmt.Println(err)
		}
	}
	var p []ProjectInfo
	err := json.Unmarshal(projectInfoByte, &p)
	if err != nil {
		log.Fatal(err)
	}
	return p
}

func GetProjectID() int {
	projectInfo := FetchProjectInfo()
	var targetProjectID int
	for _, project := range projectInfo {
		if project.Name == strings.TrimSpace(GetProjectName()) {
			targetProjectID = project.ID
		}
	}
	return targetProjectID
}
