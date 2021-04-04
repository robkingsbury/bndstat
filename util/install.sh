#!/bin/bash

# When tagging a commit, the git tag has to be faked so that the tagged commit
# includes a README with the correct tag. Other than preparing a tagged commit,
# providing a tag from the cmdline should not be used.
version=""
tags=""
if [[ ${1} != "" ]]; then
  echo "Forcing tag to ${1}"
  version=${1}
  tags=${1}
fi

# Get the latest semver tag to set the version
if [[ ${version} == "" ]]; then
  version=$(git tag --sort=-v:refname | head -1)
fi
echo "Version: ${version}"

# Get current commit's short hash
commit=$(git rev-parse --short HEAD)
echo "Current commit: ${commit}"

# Get any tags for the current commit if they exist
if [[ ${tags} == "" ]]; then
  tags=$(git tag --points-at ${commit})
fi
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
