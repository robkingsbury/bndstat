#!/bin/bash

template="README.md.template"
readme="../README.md"

# Define function to execute cmds and save output in a var, with the cmd itself
# and markdown backticks included in the output.
output=""
cmdOutput() {
  local cmd=${1}

  output='```\n'
  output+="\$ "
  output+=$(echo ${cmd})
  output+="\n"
  output+=$(${cmd} 2>&1)
  output+='\n```\n'
}

# Define function to iterate line by line over a string, substituting the
# third input string for the second when encountered. Output is saved in the
# defined var.
subout=""
substitute() {
  local template="${1}"
  local to_be_replaced="${2}"
  local replacement="${3}"

  subout=""
  while IFS= read -r line; do
    if [ "${line}" == "${to_be_replaced}" ];
    then
      subout+="${replacement}"
    else
      subout+="${line}"
      subout+="\n"
    fi
  done <<< "$template"
}

# Rebuild the binary
echo "Building the binary ..."
pushd ..
go install
popd

t=$(cat ${template})

echo "Generating help ..."
cmd="bndstat --help"
cmdOutput "${cmd}"
substitute "${t}" "HELPOUTPUT" "${output}"
t=$(echo -e "${subout}")

echo "Generating example one ..."
cmd="bndstat 3 5"
cmdOutput "${cmd}"
substitute "${t}" "EXAMPLEONE" "${output}"
t=$(echo -e "${subout}")

echo "Generating example two ..."
cmd="bndstat --devices=eth1,eth2 --interval=1 --count=5"
cmdOutput "${cmd}"
substitute "${t}" "EXAMPLETWO" "${output}"
t=$(echo -e "${subout}")

echo "Generating debug example ..."
cmd="bndstat --logtostderr --v=2 --count=1"
cmdOutput "${cmd}"
substitute "${t}" "DEBUGEXAMPLE" "${output}"
t=$(echo -e "${subout}")

echo "Generating version ..."
cmd="bndstat --version"
cmdOutput "${cmd}"
substitute "${t}" "VERSION" "${output}"
t=$(echo -e "${subout}")

echo -e "${t}" > ${readme}
