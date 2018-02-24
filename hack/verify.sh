#!/bin/bash

CRT_DIR=$(pwd)
VERIFY_TEMP=$(mktemp -d 2>/dev/null || mktemp -d -t k8s-community.XXXXXX)
WORKING_DIR=${VERIFY_TEMP}/src/testgendocs
GOPATH=${VERIFY_TEMP}
mkdir -p ${WORKING_DIR}

function cleanup {
  rm -rf "${VERIFY_TEMP}"
}
trap cleanup EXIT

cp -r sigs.yaml sig-* wg-* Makefile generator ${WORKING_DIR}/

cd ${WORKING_DIR}
make 1>/dev/null

mismatches=0
break=$(printf "=%.0s" $(seq 1 68))

for file in $(ls ${CRT_DIR}/sig-*/README.md ${CRT_DIR}/wg-*/README.md ${CRT_DIR}/sig-list.md ${CRT_DIR}/OWNERS_ALIASES); do
  real=${file#$CRT_DIR/}
  if ! diff -q ${file} ${WORKING_DIR}/${real} &>/dev/null; then
    echo "${file} does not match ${WORKING_DIR}/${real}";
    mismatches=$((mismatches+1))
  fi;
done

if [ ${mismatches} -gt "0" ]; then
  echo ""
  echo ${break}
  noun="mismatch was"
  if [ ${mismatches} -gt "0" ]; then
    noun="mismatches were"
  fi
  echo "${mismatches} ${noun} detected."
  echo "Do not manually edit sig-list.md or README.md files inside the sig folders."
  echo "Instead make your changes to sigs.yaml and then run \`make\`.";
  echo ${break}
  exit 1;
fi

exit 0
