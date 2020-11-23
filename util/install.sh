#!/bin/bash

# Get current commit's short hash
commit=$(git rev-parse --short HEAD)
echo "Current commit: ${commit}"

# Get any tags if they exist
tags=$(git tag --points-at ${commit})
echo "Tags: ${tags}"

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
"
