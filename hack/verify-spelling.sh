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

export KUBE_ROOT=$(dirname "${BASH_SOURCE}")/..
export GO111MODULE=on
export GOPROXY="${GOPROXY:-https://proxy.golang.org}"

# Pick out version of misspell from go.mod
go mod download
misspell="$(go list -m -f '{{.Dir}}' github.com/client9/misspell)"

# Spell checking
# All the skipping files are defined in hack/.spelling_failures
skipping_file="${KUBE_ROOT}/hack/.spelling_failures"
failing_packages=$(echo `cat ${skipping_file}` | sed "s| | -e |g")
git ls-files | grep -v -e ${failing_packages} | xargs go run "${misspell}/cmd/misspell" -i "" -error -o stderr
