#!/bin/bash

if [[ ${1} == "" ]]; then
  echo "Tag needs to be provided"
  exit 1
fi

# Check the format of the tag.
TAG_REGEX='^v[0-9]+\.[0-9]+\.[0-9]+$'
if ! [[ ${1} =~ ${TAG_REGEX} ]]; then
  echo "Tag needs to be in v##.##.## format"
  exit 1
fi

echo "[tagit] Running gofmt"
echo -n "pushd: "
pushd ..
find . -name \*.go -exec gofmt -l -w -s {} \;
echo -n "pop: "
popd

echo
echo "[tagit] Generating README"
./genreadme.sh ${1}

echo
echo "[tagit] Commiting changes and tagging as ${1}"
git add ..
git commit -m "Tagging ${1}"
git tag ${1}

echo
echo "[tagit] Pushing"
git push
git push origin ${1}
