#!/bin/bash

if [[ ${1} == "" ]]; then
  echo "Tag needs to be provided"
  exit 1
fi

echo "[tagit] Running gofmt"
pushd ..
find . -name \*.go -exec gofmt -l -w -s {} \;
popd

echo "[tagit] Generating README"
./genreadme.sh ${1}

echo "[tagit] Commiting changes and tagging as ${1}"
git add ..
git commit -m "Tagging ${1}"
git tag ${1}
