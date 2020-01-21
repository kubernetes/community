/*
Copyright 2020 The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package main

import (
	"bytes"
	"testing"
)

const (
	testInputWithCustom = `aliases:
  B:
    - C
    - B
    - A
  A:
    - C
    - B
    - A
## BEGIN CUSTOM CONTENT
  C:
    - B
    - A
## END CUSTOM CONTENT
`
	testInputWithCustomFormatted = `aliases:
  A:
    - A
    - B
    - C
  B:
    - A
    - B
    - C
## BEGIN CUSTOM CONTENT
  C:
    - A
    - B
## END CUSTOM CONTENT
`

	testInputWithCustomWrongPosition = `aliases:
## BEGIN CUSTOM CONTENT
  C:
    - A
    - B
## END CUSTOM CONTENT
  A:
    - A
    - B
    - C
`
	testInputWithComments = `aliases:
# heading comments are ignored
  B:
    - T # commentT
    - S
    - R # commentR
  A:
    - C
    - B
    - A
`
	testInputWithCommentsFormatted = `aliases:
  A:
    - A
    - B
    - C
  B:
    - R # commentR
    - S
    - T # commentT
`
	testInputWithDuplicateOwners = `aliases:
  A:
    - A
    - A
    - B
`
	testInputWithDuplicateOwnersFormatted = `aliases:
  A:
    - A
    - B
`
	testInputWithDuplicateGroup = `aliases:
  A:
    - A
    - B
  A:
    - A
    - B
    - C
`
)

func TestProcessData(t *testing.T) {
	var tests = []struct {
		name           string
		input          string
		f              *flags
		expectedOutput string
		expectedError  bool
	}{
		{
			name:           "valid: with custom content and formating",
			input:          testInputWithCustom,
			expectedOutput: testInputWithCustomFormatted,
			f:              &flags{custom: true},
		},
		{
			name:           "valid: format without custom content",
			input:          testInputWithComments,
			expectedOutput: testInputWithCommentsFormatted,
		},
		{
			name:           "valid: format without custom content",
			input:          testInputWithComments,
			expectedOutput: testInputWithCommentsFormatted,
		},
		{
			name:           "valid: should remove duplicates from owners",
			input:          testInputWithDuplicateOwners,
			expectedOutput: testInputWithDuplicateOwnersFormatted,
		},
		{
			name:          "invalid: custom content requested but not present",
			input:         testInputWithComments,
			f:             &flags{custom: true},
			expectedError: true,
		},
		{
			name:          "invalid: custom content present at wrong position",
			input:         testInputWithCustomWrongPosition,
			f:             &flags{custom: true},
			expectedError: true,
		},
		{
			name:          "invalid: duplicate groups should fail unmarshaling",
			input:         testInputWithDuplicateGroup,
			expectedError: true,
		},
		{
			name:          "invalid: input does not have aliases",
			input:         "",
			expectedError: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.f == nil {
				tt.f = &flags{}
			}
			output, err := processData([]byte(tt.input), tt.f)
			if (err != nil) != tt.expectedError {
				t.Fatalf("expected error: %v, got: %v, error: %v", tt.expectedError, err != nil, err)
			}
			if bytes.Compare(output, []byte(tt.expectedOutput)) != 0 {
				t.Fatalf("expected output does not match\n"+
					"output:\n%s\n"+
					"expectedOutput:\n%s\n",
					output, tt.expectedOutput)
			}
		})
	}
}
