/*
Copyright 2017 The Kubernetes Authors.

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
	"io/ioutil"
	"os"
	"path/filepath"
	"reflect"
	"strings"
	"testing"
)

func TestNonExistantDirIsCreated(t *testing.T) {
	dir := "/tmp/nonexistent"
	err := createDirIfNotExists(dir)
	if err != nil {
		t.Fatalf("Received error creating dir: %v", err)
	}
	if !pathExists(dir) {
		t.Fatalf("%s should exist", dir)
	}
}

func TestExistantDirNotCreated(t *testing.T) {
	dir := "./testdata"
	err := createDirIfNotExists(dir)
	if err != nil {
		t.Fatalf("Received error creating dir: %v", err)
	}
}

func TestGetExistingData(t *testing.T) {
	cases := []struct {
		path      string
		expected  string
		expectErr bool
	}{
		{
			path:      "./testdata/custom_content.md",
			expected:  "FOO BAR BAZ",
			expectErr: false,
		},
		{
			path:      "./testdata/no_custom_content.md",
			expected:  "",
			expectErr: false,
		},
		{
			path:      "./testdata/foo.md",
			expected:  "",
			expectErr: true,
		},
	}

	for _, c := range cases {
		content, err := getExistingContent(c.path)
		if err != nil && c.expectErr == false {
			t.Fatalf("Received unexpected error for %s: %v", c.path, err)
		}
		if err == nil && c.expectErr == true {
			t.Fatalf("Expected error for %s but received none", c.path)
		}
		if content != c.expected {
			t.Fatalf("Expected %s but received %s", c.expected, content)
		}
	}
}

func TestWriteTemplate(t *testing.T) {
	customContent := `
<!-- BEGIN CUSTOM CONTENT -->
Example
custom
content!

<!-- END CUSTOM CONTENT -->
`

	cases := []struct {
		templatePath string
		outputPath   string
		data         map[string]string
		expectErr    bool
		expected     string
	}{
		{
			templatePath: "./testdata/non_existent_template.tmpl",
			expectErr:    true,
		},
		{
			templatePath: "./testdata/example.tmpl",
			outputPath:   "/tmp/non_existing_path.md",
			expectErr:    false,
			data:         map[string]string{"Message": "Hello!"},
			expected:     "Hello!",
		},
		{
			templatePath: "./testdata/example.tmpl",
			outputPath:   "./testdata/example.md",
			expectErr:    false,
			data:         map[string]string{"Message": "Hello!"},
			expected:     customContent,
		},
	}

	for _, c := range cases {
		err := writeTemplate(c.templatePath, c.outputPath, c.data)
		if err != nil && c.expectErr == false {
			t.Fatalf("Received unexpected error for %s: %v", c.templatePath, err)
		}
		if c.expectErr {
			if err == nil {
				t.Fatalf("Expected error for %s but received none", c.templatePath)
			}
			continue
		}
		content, err := ioutil.ReadFile(c.outputPath)
		if err != nil {
			t.Fatalf("%s should exist", c.outputPath)
		}
		if strings.Contains(string(content), c.expected) == false {
			t.Fatalf("%s was not found in %s", c.expected, c.outputPath)
		}
	}
}

func TestGroupDirName(t *testing.T) {
	group := Group{Name: "Foo Bar"}
	if group.DirName("sig") != "sig-foo-bar" {
		t.Fatal("DirName incorrect")
	}
}

func TestSetupGithubTeams(t *testing.T) {
	group := Group{Name: "Foo Bar"}
	group.SetupGitHubTeams("sig")

	var expected []string
	for _, ght := range githubTeamNames {
		expected = append(expected, fmt.Sprintf("sig-foo-bar-%s", ght))
	}

	if !reflect.DeepEqual(group.Contact.GithubTeamNames, expected) {
		t.Fatalf("%v does not match %v", group.Contact.GithubTeamNames, expected)
	}
}

func TestCustomPrefixSetupGithubTeams(t *testing.T) {
	group := Group{Contact: Contact{GithubTeamPrefix: "foo"}}
	group.SetupGitHubTeams("")

	var expected []string
	for _, ght := range githubTeamNames {
		expected = append(expected, fmt.Sprintf("foo-%s", ght))
	}

	if !reflect.DeepEqual(group.Contact.GithubTeamNames, expected) {
		t.Fatalf("%v does not match %v", group.Contact.GithubTeamNames, expected)
	}
}

func TestCreateGroupReadmes(t *testing.T) {
	groups := []Group{
		Group{Name: "Foo"},
		Group{Name: "Bar"},
	}

	err := createGroupReadme(groups, "sig")
	if err != nil {
		t.Fatal(err)
	}

	for _, group := range groups {
		path := filepath.Join(baseOutputDir, group.DirName("sig"), "README.md")
		if !pathExists(path) {
			t.Fatalf("%s should exist", path)
		}
	}
}

func TestReadmesAreSkipped(t *testing.T) {
	os.Setenv("SIG", "sig-foo")

	groups := []Group{
		Group{Name: "Foo"},
		Group{Name: "Bar"},
	}

	err := createGroupReadme(groups, "sig")
	if err != nil {
		t.Fatal(err)
	}

	for _, group := range groups[1:] {
		path := filepath.Join(baseOutputDir, group.DirName("sig"), "README.md")
		if !pathExists(path) {
			t.Fatalf("%s should exist", path)
		}
	}

	os.Setenv("SIG", "")
}

func copyFile(src, dst string) error {
	// Read all content of src to data
	data, err := ioutil.ReadFile(src)
	if err != nil {
		return err
	}
	// Write data to dst
	err = ioutil.WriteFile(dst, data, 0644)
	if err != nil {
		return err
	}
	return nil
}

func TestFullGeneration(t *testing.T) {
	err := copyFile("testdata/sigs.yaml", "generated/sigs.yaml")
	if err != nil {
		t.Fatalf("Error received: %v", err)
	}

	main()

	expectedDirs := []string{"sig-foo", "sig-bar", "wg-baz"}
	for _, ed := range expectedDirs {
		path := filepath.Join(baseOutputDir, ed, "README.md")
		if !pathExists(path) {
			t.Fatalf("%s should exist", path)
		}
	}
}
