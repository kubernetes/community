#!/usr/bin/env bash

# Copyright 2022 The Kubernetes Authors.
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

# This script uses pandoc (https://pandoc.org/) to reformat Markdown
# files to enforce those parts of the style guide that can be enforced
# through code.
#
# There are a number of formatting issues which pandoc is not able to
# address as of version 2.14:
#
# * It will wrap lines in the middle of links, which breaks formatting
#   with Hugo.
# * It will also renumber ordered lists, which breaks from the
#   Kubernetes style guide.
# * It alphabetizes by keyword in YAML headers.

PANDOC=$(which pandoc)

if [ ! -x "$PANDOC" ]
then
    echo "Please install pandoc (https://pandoc.org/) before using this script."
    exit 1
fi

TARGET="$1"

if [ ! -f "$TARGET" ]
then
    echo "$TARGET does not exist."
    exit 1
fi

TMPFILE=$(mktemp)

# Run pandoc. Note that "gfm" is "GitHub-flavored Markdown".
$PANDOC -f gfm+pipe_tables -t gfm+pipe_tables \
        --markdown-headings=atx --reference-links \
        --standalone \
        --wrap=auto --columns=80 \
        --toc \
        -o "$TMPFILE" "$TARGET"

if [ 0 -ne $? ]
then
    echo "$PANDOC did not execute successfully."
    exit 1
fi

# Backup the target file and copy the reformatted file over it.
cp -v "$TARGET" "$TARGET~"
cp -v "$TMPFILE" "$TARGET"
rm -v "$TMPFILE"
