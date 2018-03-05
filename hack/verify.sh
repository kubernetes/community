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

KUBE_ROOT=$(dirname "${BASH_SOURCE}")/..

# Some useful colors.
if [[ -z "${color_start-}" ]]; then
  declare -r color_start="\033["
  declare -r color_red="${color_start}0;31m"
  declare -r color_yellow="${color_start}0;33m"
  declare -r color_green="${color_start}0;32m"
  declare -r color_norm="${color_start}0m"
fi

# Excluded check patterns are always skipped.
EXCLUDED_PATTERNS=(
  "verify-all.sh"                # this script calls the make rule and would cause a loop
  "verify-*-dockerized.sh"       # Don't run any scripts that intended to be run dockerized
  )

EXCLUDED_CHECKS=$(ls ${EXCLUDED_PATTERNS[@]/#/${KUBE_ROOT}\/hack\/} 2>/dev/null || true)

function is-excluded {
  for e in ${EXCLUDED_CHECKS[@]}; do
    if [[ $1 -ef "$e" ]]; then
      return
    fi
  done
  return 1
}

function run-cmd {
  if ${SILENT}; then
    "$@" &> /dev/null
  else
    "$@"
  fi
}

# Collect Failed tests in this Array , initialize it to nil
FAILED_TESTS=()

function print-failed-tests {
  echo -e "========================"
  echo -e "${color_red}FAILED TESTS${color_norm}"
  echo -e "========================"
  for t in ${FAILED_TESTS[@]}; do
      echo -e "${color_red}${t}${color_norm}"
  done
}

function run-checks {
  local -r pattern=$1
  local -r runner=$2

  local t
  for t in $(ls ${pattern})
  do
    local check_name="$(basename "${t}")"
    if is-excluded "${t}" ; then
      echo "Skipping ${check_name}"
      continue
    fi
    echo -e "Verifying ${check_name}"
    local start=$(date +%s)
    run-cmd "${runner}" "${t}" && tr=$? || tr=$?
    local elapsed=$(($(date +%s) - ${start}))
    if [[ ${tr} -eq 0 ]]; then
      echo -e "${color_green}SUCCESS${color_norm}  ${check_name}\t${elapsed}s"
    else
      echo -e "${color_red}FAILED${color_norm}   ${check_name}\t${elapsed}s"
      ret=1
      FAILED_TESTS+=(${t})
    fi
  done
}

SILENT=false

while getopts ":s" opt; do
  case ${opt} in
    s)
      SILENT=true
      ;;
    \?)
      echo "Invalid flag: -${OPTARG}" >&2
      exit 1
      ;;
  esac
done

if ${SILENT} ; then
  echo "Running in silent mode, run without -s if you want to see script logs."
fi

ret=0
run-checks "${KUBE_ROOT}/hack/verify-*.sh" bash

if [[ ${ret} -eq 1 ]]; then
    print-failed-tests
fi
exit ${ret}

# ex: ts=2 sw=2 et filetype=sh
