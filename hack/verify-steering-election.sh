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
shopt -s extglob

export KUBE_ROOT=$(dirname "${BASH_SOURCE}")/..

# exclude bios before 2025 since some of them have more than 300 words
STEERING_ELECTION_BIOS="${KUBE_ROOT}/elections/steering/!(2017|2018|2019|2020|2021|2022|2023|2024)/candidate-*.md"

invalid_bios=0
break=$(printf "=%.0s" $(seq 1 68))

for bio in ${STEERING_ELECTION_BIOS} ; do
  [[ -f $bio ]] || continue
  word_count=$(wc -w < "$bio")
  if [[ ${word_count} -gt "450" ]]; then
    echo "${bio} has ${word_count} words."
    invalid_bios=$((invalid_bios+1))
  fi
done

if [[ ${invalid_bios} -gt "0" ]]; then
  echo ""
  echo "${break}"
  echo "${invalid_bios} invalid Steering Committee election bio(s) detected."
  echo "Bios should be limited to around 300 words, excluding headers."
  echo "${break}"
  exit 1;
fi
