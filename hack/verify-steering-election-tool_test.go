/*
Copyright 2025 The Kubernetes Authors.

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
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"testing"
	"time"
)

func TestExtractYAMLHeader(t *testing.T) {
	tests := []struct {
		name    string
		content string
		want    string
		wantErr bool
	}{
		{
			name: "valid YAML header with exactly 61 dashes",
			content: strings.Repeat("-", 61) + "\n" +
				"name: Test User\n" +
				"ID: testuser\n" +
				"info:\n" +
				"  employer: Test Corp\n" +
				"  slack: '@testuser'\n" +
				strings.Repeat("-", 61) + "\n" +
				"## Bio content\n",
			want:    "name: Test User\nID: testuser\ninfo:\n  employer: Test Corp\n  slack: '@testuser'",
			wantErr: false,
		},
		{
			name: "invalid YAML header with wrong number of dashes",
			content: strings.Repeat("-", 60) + "\n" +
				"name: Test User\n" +
				strings.Repeat("-", 60) + "\n",
			want:    "",
			wantErr: true,
		},
		{
			name:    "no YAML header",
			content: "Just some content without header",
			want:    "",
			wantErr: true,
		},
		{
			name: "YAML header with extra spaces",
			content: strings.Repeat("-", 61) + "   \n" +
				"name: Test User\n" +
				"ID: testuser\n" +
				strings.Repeat("-", 61) + "  \n" +
				"## Bio content\n",
			want:    "name: Test User\nID: testuser",
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := extractYAMLHeader(tt.content)
			if (err != nil) != tt.wantErr {
				t.Errorf("extractYAMLHeader() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("extractYAMLHeader() = %q, want %q", got, tt.want)
			}
		})
	}
}

func TestCountWords(t *testing.T) {
	tests := []struct {
		name    string
		content string
		want    int
	}{
		{
			name:    "simple sentence",
			content: "Hello world test",
			want:    3,
		},
		{
			name:    "text with newlines",
			content: "Hello\nworld\ntest",
			want:    3,
		},
		{
			name:    "text with multiple spaces",
			content: "Hello    world     test",
			want:    3,
		},
		{
			name:    "empty content",
			content: "",
			want:    0,
		},
		{
			name:    "only whitespace",
			content: "   \n\t  \n  ",
			want:    0,
		},
		{
			name:    "mixed content with punctuation",
			content: "Hello, world! This is a test.",
			want:    6,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Create a temporary file with the content
			tmpfile, err := os.CreateTemp("", "test")
			if err != nil {
				t.Fatal(err)
			}
			defer os.Remove(tmpfile.Name())

			if _, err := tmpfile.Write([]byte(tt.content)); err != nil {
				t.Fatal(err)
			}
			if err := tmpfile.Close(); err != nil {
				t.Fatal(err)
			}

			got, err := countWords(tmpfile.Name())
			if err != nil {
				t.Errorf("countWords() error = %v", err)
				return
			}
			if got != tt.want {
				t.Errorf("countWords() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestValidateFileNameAndGitHubID(t *testing.T) {
	tests := []struct {
		name        string
		filename    string
		fileContent string
		wantErr     bool
		errContains string
	}{
		{
			name:     "valid filename and matching GitHub ID",
			filename: "candidate-testuser.md",
			fileContent: strings.Repeat("-", 61) + "\n" +
				"name: Test User\n" +
				"ID: testuser\n" +
				"info:\n" +
				"  employer: Test Corp\n" +
				"  slack: '@testuser'\n" +
				strings.Repeat("-", 61) + "\n" +
				"## Bio content\n",
			wantErr: false,
		},
		{
			name:        "invalid filename format",
			filename:    "invalid-format.md",
			fileContent: "dummy content",
			wantErr:     true,
			errContains: "filename must follow format",
		},
		{
			name:     "mismatched GitHub ID",
			filename: "candidate-testuser.md",
			fileContent: strings.Repeat("-", 61) + "\n" +
				"name: Test User\n" +
				"ID: differentuser\n" +
				"info:\n" +
				"  employer: Test Corp\n" +
				"  slack: '@testuser'\n" +
				strings.Repeat("-", 61) + "\n" +
				"## Bio content\n",
			wantErr:     true,
			errContains: "does not match GitHub ID",
		},
		{
			name:        "missing ID field",
			filename:    "candidate-testuser.md",
			fileContent: strings.Repeat("-", 61) + "\nname: Test User\n" + strings.Repeat("-", 61) + "\n",
			wantErr:     true,
			errContains: "missing 'ID' field",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Create a temporary file with the content
			tmpfile, err := os.CreateTemp("", tt.filename)
			if err != nil {
				t.Fatal(err)
			}
			defer os.Remove(tmpfile.Name())

			// Rename to the desired filename
			dir := filepath.Dir(tmpfile.Name())
			targetPath := filepath.Join(dir, tt.filename)
			if err := os.Rename(tmpfile.Name(), targetPath); err != nil {
				t.Fatal(err)
			}
			defer os.Remove(targetPath)

			if err := os.WriteFile(targetPath, []byte(tt.fileContent), 0644); err != nil {
				t.Fatal(err)
			}

			err = validateFileNameAndGitHubID(targetPath)
			if (err != nil) != tt.wantErr {
				t.Errorf("validateFileNameAndGitHubID() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if tt.wantErr && tt.errContains != "" && !strings.Contains(err.Error(), tt.errContains) {
				t.Errorf("validateFileNameAndGitHubID() error = %v, should contain %q", err, tt.errContains)
			}
		})
	}
}

func TestValidateTemplateCompliance(t *testing.T) {
	validContent := strings.Repeat("-", 61) + "\n" +
		"name: Test User\n" +
		"ID: testuser\n" +
		"info:\n" +
		"  employer: Test Corp\n" +
		"  slack: '@testuser'\n" +
		strings.Repeat("-", 61) + "\n" +
		"## What I have done\n" +
		"Some accomplishments\n" +
		"## What I'll do\n" +
		"Future plans\n" +
		"## SIGS\n" +
		"SIG involvement\n" +
		"## Resources About Me\n" +
		"Links and info\n"

	tests := []struct {
		name        string
		content     string
		wantErr     bool
		errContains string
	}{
		{
			name:    "valid template compliance",
			content: validContent,
			wantErr: false,
		},
		{
			name: "missing required field - name",
			content: strings.Repeat("-", 61) + "\n" +
				"ID: testuser\n" +
				"info:\n" +
				"  employer: Test Corp\n" +
				"  slack: '@testuser'\n" +
				strings.Repeat("-", 61) + "\n" +
				"## What I have done\n## What I'll do\n## SIGS\n## Resources About Me\n",
			wantErr:     true,
			errContains: "missing required field: name",
		},
		{
			name: "missing required section",
			content: strings.Repeat("-", 61) + "\n" +
				"name: Test User\n" +
				"ID: testuser\n" +
				"info:\n" +
				"  employer: Test Corp\n" +
				"  slack: '@testuser'\n" +
				strings.Repeat("-", 61) + "\n" +
				"## What I have done\n## What I'll do\n## Resources About Me\n",
			wantErr:     true,
			errContains: "missing required section: SIGs",
		},
		{
			name: "alternative SIGs section name",
			content: strings.Repeat("-", 61) + "\n" +
				"name: Test User\n" +
				"ID: testuser\n" +
				"info:\n" +
				"  employer: Test Corp\n" +
				"  slack: '@testuser'\n" +
				strings.Repeat("-", 61) + "\n" +
				"## What I have done\n## What I'll do\n## SIGs\n## Resources About Me\n",
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Create a temporary file with the content
			tmpfile, err := os.CreateTemp("", "test")
			if err != nil {
				t.Fatal(err)
			}
			defer os.Remove(tmpfile.Name())

			if err := os.WriteFile(tmpfile.Name(), []byte(tt.content), 0644); err != nil {
				t.Fatal(err)
			}

			err = validateTemplateCompliance(tmpfile.Name())
			if (err != nil) != tt.wantErr {
				t.Errorf("validateTemplateCompliance() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if tt.wantErr && tt.errContains != "" && !strings.Contains(err.Error(), tt.errContains) {
				t.Errorf("validateTemplateCompliance() error = %v, should contain %q", err, tt.errContains)
			}
		})
	}
}

func TestFindBioFiles(t *testing.T) {
	// Create a temporary directory structure
	tmpDir, err := os.MkdirTemp("", "elections")
	if err != nil {
		t.Fatal(err)
	}
	defer os.RemoveAll(tmpDir)

	// Create steering directory structure
	steeringDir := filepath.Join(tmpDir, "steering")
	if err := os.MkdirAll(steeringDir, 0755); err != nil {
		t.Fatal(err)
	}

	// Create test files
	currentYear := time.Now().Year()
	testFiles := []struct {
		path      string
		shouldFind bool
	}{
		{filepath.Join(steeringDir, "2025", "candidate-user1.md"), false}, // Before startYear (2026)
		{filepath.Join(steeringDir, "2026", "candidate-user2.md"), true},  // At startYear
		{filepath.Join(steeringDir, fmt.Sprintf("%d", currentYear+1), "candidate-user3.md"), true}, // Future year
		{filepath.Join(steeringDir, "2026", "not-candidate.md"), false},     // Wrong filename format
		{filepath.Join(steeringDir, "2026", "candidate-user4.txt"), false},  // Wrong extension
		{filepath.Join(steeringDir, fmt.Sprintf("%d", currentYear+10), "candidate-user5.md"), false}, // Too far in future
	}

	var expectedFiles []string
	for _, tf := range testFiles {
		dir := filepath.Dir(tf.path)
		if err := os.MkdirAll(dir, 0755); err != nil {
			t.Fatal(err)
		}
		if err := os.WriteFile(tf.path, []byte("dummy content"), 0644); err != nil {
			t.Fatal(err)
		}
		if tf.shouldFind {
			expectedFiles = append(expectedFiles, tf.path)
		}
	}

	// Test findBioFiles
	bioFiles, err := findBioFiles(tmpDir)
	if err != nil {
		t.Errorf("findBioFiles() error = %v", err)
		return
	}

	if len(bioFiles) != len(expectedFiles) {
		t.Errorf("findBioFiles() found %d files, expected %d", len(bioFiles), len(expectedFiles))
	}

	// Check that all expected files are found
	for _, expected := range expectedFiles {
		found := false
		for _, actual := range bioFiles {
			if actual == expected {
				found = true
				break
			}
		}
		if !found {
			t.Errorf("findBioFiles() did not find expected file: %s", expected)
		}
	}
}