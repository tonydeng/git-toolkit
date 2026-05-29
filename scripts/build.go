package main

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"time"
)

const (
	BinaryName = "git-toolkit"
	DistDir    = "dist"
)

var platforms = []Platform{
	{OS: "linux", Arch: "amd64", Suffix: ""},
	{OS: "darwin", Arch: "amd64", Suffix: ""},
	{OS: "darwin", Arch: "arm64", Suffix: ""},
	{OS: "windows", Arch: "amd64", Suffix: ".exe"},
}

type Platform struct {
	OS     string
	Arch   string
	Suffix string
}

func main() {
	args := os.Args[1:]
	
	if len(args) > 0 {
		switch args[0] {
		case "clean":
			clean()
			return
		case "help":
			printHelp()
			return
		}
	}
	
	if len(args) == 1 && args[0] != "all" {
		buildSinglePlatform(args[0])
	} else {
		buildAllPlatforms()
	}
}

func printHelp() {
	fmt.Println("Usage: go run scripts/build.go [command] [platform]")
	fmt.Println()
	fmt.Println("Commands:")
	fmt.Println("  all       Build all platforms (default)")
	fmt.Println("  clean     Clean dist directory")
	fmt.Println("  help      Show this help message")
	fmt.Println()
	fmt.Println("Platforms:")
	fmt.Println("  linux-amd64")
	fmt.Println("  darwin-amd64")
	fmt.Println("  darwin-arm64")
	fmt.Println("  windows-amd64")
	fmt.Println()
	fmt.Println("Examples:")
	fmt.Println("  go run scripts/build.go              # Build all platforms")
	fmt.Println("  go run scripts/build.go all          # Build all platforms")
	fmt.Println("  go run scripts/build.go linux-amd64  # Build only linux-amd64")
	fmt.Println("  go run scripts/build.go clean        # Clean dist directory")
}

func clean() {
	if _, err := os.Stat(DistDir); os.IsNotExist(err) {
		fmt.Println("dist directory does not exist")
		return
	}
	
	fmt.Printf("Cleaning %s...\n", DistDir)
	if err := os.RemoveAll(DistDir); err != nil {
		fmt.Printf("Error cleaning dist directory: %v\n", err)
		os.Exit(1)
	}
	fmt.Println("Clean complete!")
}

func buildSinglePlatform(platform string) {
	parts := strings.Split(platform, "-")
	if len(parts) != 2 {
		fmt.Printf("Invalid platform format: %s\n", platform)
		fmt.Println("Use: os-arch (e.g., linux-amd64, windows-amd64)")
		os.Exit(1)
	}
	
	osName, arch := parts[0], parts[1]
	suffix := ""
	if osName == "windows" {
		suffix = ".exe"
	}
	
	p := Platform{OS: osName, Arch: arch, Suffix: suffix}
	if err := build(p); err != nil {
		fmt.Printf("Build failed: %v\n", err)
		os.Exit(1)
	}
}

func buildAllPlatforms() {
	fmt.Println("Building all 64-bit platforms...")
	
	if err := os.MkdirAll(DistDir, 0755); err != nil {
		fmt.Printf("Error creating dist directory: %v\n", err)
		os.Exit(1)
	}
	
	success := true
	for _, p := range platforms {
		if err := build(p); err != nil {
			fmt.Printf("  -> Failed: %v\n", err)
			success = false
		}
	}
	
	fmt.Println()
	if success {
		fmt.Println("Build complete!")
	} else {
		fmt.Println("Build completed with errors!")
		os.Exit(1)
	}
}

func build(p Platform) error {
	output := filepath.Join(DistDir, fmt.Sprintf("%s_%s_%s%s", BinaryName, p.OS, p.Arch, p.Suffix))
	
	fmt.Printf("Building for %s/%s...\n", p.OS, p.Arch)
	
	version := getVersion()
	commit := getCommit()
	buildTime := getBuildTime()
	
	ldflags := fmt.Sprintf("-s -w -X 'github.com/tonydeng/git-toolkit/cmd.Version=%s' -X 'github.com/tonydeng/git-toolkit/cmd.BuildTime=%s' -X 'github.com/tonydeng/git-toolkit/cmd.CommitID=%s'", version, buildTime, commit)
	
	cmd := exec.Command("go", "build", "-mod=vendor", "-ldflags", ldflags, "-o", output, ".")
	cmd.Env = append(os.Environ(),
		fmt.Sprintf("GOOS=%s", p.OS),
		fmt.Sprintf("GOARCH=%s", p.Arch),
	)
	
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	
	if err := cmd.Run(); err != nil {
		return err
	}
	
	info, err := os.Stat(output)
	if err != nil {
		return err
	}
	
	sizeMB := float64(info.Size()) / 1024 / 1024
	fmt.Printf("  -> %s (%.2f MB)\n", output, sizeMB)
	
	return nil
}

func getVersion() string {
	data, err := os.ReadFile("version")
	if err != nil {
		return "dev"
	}
	return strings.TrimSpace(string(data))
}

func getCommit() string {
	cmd := exec.Command("git", "rev-parse", "HEAD")
	output, err := cmd.Output()
	if err != nil {
		return "unknown"
	}
	return strings.TrimSpace(string(output))
}

func getBuildTime() string {
	return time.Now().Format("2006-01-02 15:04:05")
}
