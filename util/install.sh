#!/bin/bash

# Get the latest semver tag to set the version
version=$(git tag --sort=-v:refname | head -1)
echo "Version: ${version}"

# Get current commit's short hash
commit=$(git rev-parse --short HEAD)
echo "Current commit: ${commit}"

# Get any tags for the current commit if they exist
tags=$(git tag --points-at ${commit})
echo "Current Commit Tags: ${tags}"

# What time is it?
compile_time=$(date)
echo "Compile time: ${compile_time}"

# What's the machine that we're on
host=$(uname -n)
echo "Build host: ${host}"

cd ..
echo "Compiling and installing bndstat ..."
go install -ldflags="\
  -X 'main.buildHost=${host}' \
  -X 'main.commit=${commit}' \
  -X 'main.compileTime=${compile_time}'\
  -X 'main.tags=${tags}' \
  -X 'main.version=${version}' \
"
