BUILD_VERSION	:= $(shell cat version)
BUILD_TIME		:= $(shell date "+%F %T")
COMMIT_SHA1		:= $(shell git rev-parse HEAD)

LDFLAGS := -ldflags "-X 'github.com/tonydeng/git-toolkit/cmd.Version=${BUILD_VERSION}' \
			-X 'github.com/tonydeng/git-toolkit/cmd.BuildTime=${BUILD_TIME}' \
			-X 'github.com/tonydeng/git-toolkit/cmd.CommitID=${COMMIT_SHA1}'"

# Binary name
BINARY_NAME := git-toolkit

# Build directory
DIST_DIR := dist

.PHONY: all clean install build-darwin-amd64 build-darwin-arm64 build-linux-386 build-linux-amd64 build-windows-amd64 build-windows-386

all: build-darwin-amd64 build-darwin-arm64 build-linux-386 build-linux-amd64 build-windows-amd64 build-windows-386

# macOS Intel
build-darwin-amd64:
	GOOS=darwin GOARCH=amd64 go build -mod=vendor $(LDFLAGS) -o $(DIST_DIR)/$(BINARY_NAME)_darwin_amd64

# macOS Apple Silicon
build-darwin-arm64:
	GOOS=darwin GOARCH=arm64 go build -mod=vendor $(LDFLAGS) -o $(DIST_DIR)/$(BINARY_NAME)_darwin_arm64

# Linux 32-bit
build-linux-386:
	GOOS=linux GOARCH=386 go build -mod=vendor $(LDFLAGS) -o $(DIST_DIR)/$(BINARY_NAME)_linux_386

# Linux 64-bit
build-linux-amd64:
	GOOS=linux GOARCH=amd64 go build -mod=vendor $(LDFLAGS) -o $(DIST_DIR)/$(BINARY_NAME)_linux_amd64

# Windows 64-bit
build-windows-amd64:
	GOOS=windows GOARCH=amd64 go build -mod=vendor $(LDFLAGS) -o $(DIST_DIR)/$(BINARY_NAME)_windows_amd64.exe

# Windows 32-bit
build-windows-386:
	GOOS=windows GOARCH=386 go build -mod=vendor $(LDFLAGS) -o $(DIST_DIR)/$(BINARY_NAME)_windows_386.exe

clean:
	rm -rf $(DIST_DIR)

install:
	go install -mod=vendor $(LDFLAGS)

.EXPORT_ALL_VARIABLES:

GO111MODULE = on
GOPROXY = https://goproxy.io
GOSUMDB = sum.golang.google.cn
