#!/usr/bin/env bash

# Copyright 2020 The Kubernetes Authors.
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

root=$(dirname "${BASH_SOURCE[@]}")/..
export KUBE_ROOT=$root
export GO111MODULE=on
export GOPROXY="${GOPROXY:-https://proxy.golang.org}"

a="${KUBE_ROOT}/hack/format-owners-aliases.go"
b="${KUBE_ROOT}/hack/format-owners-aliases_test.go"
go vet "$a" "$b"
go test "$a" "$b" -count=1
