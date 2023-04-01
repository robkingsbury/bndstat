#!/bin/bash

if [[ ${1} == "" ]]; then
  echo "Tag needs to be provided"
  exit 1
fi

echo Fixing formatting if needed
pushd ..
find . -name \*.go -exec gofmt -l -w -s {} \;
popd

./genreadme.sh ${1}
git add ../README.md
git commit -m "Tagging ${1}"
git tag ${1}
