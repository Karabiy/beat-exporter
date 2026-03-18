#!/bin/bash

set -e

if [[ -z "$GITHUB_WORKSPACE" ]]; then
  GITHUB_WORKSPACE=$(pwd)
  echo "Setting up GITHUB_WORKSPACE to current directory: ${GITHUB_WORKSPACE}"
fi

if [[ -z "$GITHUB_ACTOR" ]]; then
  GITHUB_ACTOR=$(whoami)
  echo "Setting up GITHUB_ACTOR to current user: ${GITHUB_ACTOR}"
fi

GITVERSION=$(git describe --tags --always)
GITBRANCH=$(git branch | grep \* | cut -d ' ' -f2)
GITREVISION=$(git log -1 --oneline | cut -d ' ' -f1)
TIME=$(date +%FT%T%z)
LDFLAGS="-s -X github.com/prometheus/common/version.Version=${GITVERSION} \
-X github.com/prometheus/common/version.Revision=${GITREVISION} \
-X github.com/prometheus/common/version.Branch=master \
-X github.com/prometheus/common/version.BuildUser=${GITHUB_ACTOR} \
-X github.com/prometheus/common/version.BuildDate=${TIME}"

build() {
    local OS=$1 ARCH=$2
    echo "Building ${OS}/${ARCH} with version: ${GITVERSION}, revision: ${GITREVISION}, buildUser: ${GITHUB_ACTOR}"
    local EXT=""
    if [[ $OS == "windows" ]]; then EXT=".exe"; fi
    CGO_ENABLED=0 GOOS=${OS} GOARCH=${ARCH} go build -ldflags "${LDFLAGS}" -tags 'netgo static_build' -a -o ".build/${OS}-${ARCH}/beat-exporter${EXT}"
}

# darwin: 386 dropped in Go 1.15
for ARCH in "amd64" "arm64"; do build darwin "$ARCH"; done

# linux: all three
for ARCH in "amd64" "386" "arm64"; do build linux "$ARCH"; done

# windows: 386 dropped in recent Go
for ARCH in "amd64" "arm64"; do build windows "$ARCH"; done
