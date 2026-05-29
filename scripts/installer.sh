#!/bin/bash
set -e

REPO_HOME="${REPO_HOME:-https://github.com/tonydeng/git-toolkit}"
BRANCH="${BRANCH:-github/golang}"

case "$(uname -s)" in
    Darwin)  PRODUCTION="git-toolkit_darwin_$(uname -m)" ;;
    Linux)   PRODUCTION="git-toolkit_linux_amd64" ;;
    MINGW*|MSYS*|CYGWIN*) PRODUCTION="git-toolkit_windows_amd64.exe" ;;
    *)       echo "Unsupported OS"; exit 1 ;;
esac

URL="$REPO_HOME/raw/$BRANCH/dist/$PRODUCTION"
echo "Downloading: $URL"

if command -v wget >/dev/null; then
    wget -q --show-progress -c "$URL" || exit 1
elif command -v curl >/dev/null; then
    curl -fLO -C - "$URL" || exit 1
else
    echo "Error: wget or curl required"; exit 1
fi

chmod +x "./$PRODUCTION"
"./$PRODUCTION" install
