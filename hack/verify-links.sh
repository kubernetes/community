#!/usr/bin/env bash

# Copyright 2019 The Kubernetes Authors.
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

KUBE_TEMP=/tmp # TODO someone please tell me where the output should be written
go install ./vendor/k8s.io/test-infra/linkcheck/
OUTPUT="${KUBE_TEMP}"/linkcheck-output
cleanup() {
	rm -rf "${OUTPUT}"
}
trap "cleanup" EXIT SIGINT
mkdir -p "$OUTPUT"

root="$(git rev-parse --show-toplevel)"
ROOTS=("${root}")
found_invalid=false
for root in "${ROOTS[@]}"; do
  linkcheck "--root-dir=${root}" 2> "${OUTPUT}/error" 1> "${OUTPUT}/output" && ret=0 || ret=$?
  if [[ $ret -eq 1 ]]; then
    found_invalid=true
  fi
  if [[ $ret -gt 1 ]]; then
    echo "Error running linkcheck"
    exit 1
  fi
done

if [ ${found_invalid} = true ]; then
  exit 1
fi

# ex: ts=2 sw=2 et filetype=sh
