#!/usr/bin/env bash

# Copyright 2018 The Kubernetes Authors.
#
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
#
#     http://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.

set -o errexit
set -o nounset
set -o pipefail

CRT_DIR=$(pwd)
VERIFY_TEMP=$(mktemp -d 2>/dev/null || mktemp -d -t k8s-community.XXXXXX)
WORKING_DIR="${VERIFY_TEMP}/src/test-generate-icons-drawio"
mkdir -p "${WORKING_DIR}"

function cleanup {
  rm -rf "${VERIFY_TEMP}"
}
trap cleanup EXIT

git archive --format=tar "$(git write-tree)" | (cd "${WORKING_DIR}" && tar xf -)

cd "${WORKING_DIR}"
make generate-icons-drawio 1>/dev/null

mismatches=0
break=$(printf "=%.0s" $(seq 1 68))

diff=$(diff -rq ${CRT_DIR}/icons/tools/draw.io/ ${WORKING_DIR}/icons/tools/draw.io/ || true)

if [[ ! -z $diff ]]; then
  echo "$diff"
  echo ""
  echo "${break}"
  mismatches=$(echo "$diff" | wc -l | sed 's/^ *//')
  noun="mismatch was"
  if [ ${mismatches} -gt "1" ]; then
    noun="mismatches were"
  fi
  echo "${mismatches} ${noun} detected."
  echo "Do not manually edit draw.io icon library files."
  echo "When icon SVGs are updated, run 'make generate-icons-drawio', and then"
  echo "commit your changes to generated draw.io icon library files."
  echo "${break}"
  exit 1
else
  exit 0
fi
