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
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestNonExistentDirIsCreated(t *testing.T) {
	dir := "/tmp/nonexistent"
	err := createDirIfNotExists(dir)
	if err != nil {
		t.Fatalf("Received error creating dir: %v", err)
	}
	if !pathExists(dir) {
		t.Fatalf("%s should exist", dir)
	}
}

func TestExistentDirNotCreated(t *testing.T) {
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
		content, err := getExistingContent(c.path, "markdown")
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
	baseGeneratorDir = "generated"
	templateDir = "../../generator"

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
		fileFormat   string
		expected     string
	}{
		{
			templatePath: "./testdata/non_existent_template.tmpl",
			expectErr:    true,
			fileFormat:   "markdown",
		},
		{
			templatePath: "./testdata/example.tmpl",
			outputPath:   "/tmp/non_existing_path.md",
			expectErr:    false,
			fileFormat:   "markdown",
			data:         map[string]string{"Message": "Hello!"},
			expected:     "Hello!",
		},
		{
			templatePath: "./testdata/example.tmpl",
			outputPath:   "./testdata/example.md",
			expectErr:    false,
			fileFormat:   "markdown",
			data:         map[string]string{"Message": "Hello!"},
			expected:     customContent,
		},
	}

	for _, c := range cases {
		err := writeTemplate(c.templatePath, c.outputPath, c.fileFormat, c.data)
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

func TestCreateGroupReadmes(t *testing.T) {
	baseGeneratorDir = "generated"
	templateDir = "../../generator"
	const groupType = "sig"

	groups := []Group{}
	for _, n := range []string{"Foo", "Bar"} {
		g := Group{Name: n}
		g.Dir = g.DirName(groupType)
		groups = append(groups, g)
	}

	err := createGroupReadme(groups, groupType)
	if err != nil {
		t.Fatal(err)
	}

	for _, group := range groups {
		path := filepath.Join(baseGeneratorDir, group.DirName(groupType), "README.md")
		if !pathExists(path) {
			t.Fatalf("%s should exist", path)
		}
	}
}

func TestReadmesAreSkipped(t *testing.T) {
	baseGeneratorDir = "generated"
	templateDir = "../../generator"

	os.Setenv("SIG", "sig-foo")

	groups := []Group{
		{Name: "Foo"},
		{Name: "Bar"},
	}

	err := createGroupReadme(groups, "sig")
	if err != nil {
		t.Fatal(err)
	}

	for _, group := range groups[1:] {
		path := filepath.Join(baseGeneratorDir, group.DirName("sig"), "README.md")
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
	baseGeneratorDir = "generated"
	templateDir = "../../generator"

	err := copyFile("testdata/sigs.yaml", "generated/sigs.yaml")
	if err != nil {
		t.Fatalf("Error received: %v", err)
	}

	main()

	expectedDirs := []string{"sig-foo", "sig-bar", "wg-baz"}
	for _, ed := range expectedDirs {
		path := filepath.Join(baseGeneratorDir, ed, "README.md")
		if !pathExists(path) {
			t.Fatalf("%s should exist", path)
		}
	}
}

func TestGitHubURL(t *testing.T) {
	cases := []struct {
		name     string
		url      string
		expected string
	}{
		{
			name:     "kubernetes-sigs root raw github url",
			url:      "https://raw.githubusercontent.com/kubernetes-sigs/boskos/main/OWNERS",
			expected: "https://github.com/kubernetes-sigs/boskos/blob/main/OWNERS",
		},
		{
			name:     "kubernetes non-root raw github url",
			url:      "https://raw.githubusercontent.com/kubernetes/kubernetes/main/test/OWNERS",
			expected: "https://github.com/kubernetes/kubernetes/blob/main/test/OWNERS",
		},
		{
			name:     "kubernetes github url should be unchanged",
			url:      "https://github.com/kubernetes/kubernetes/blob/main/test/OWNERS",
			expected: "https://github.com/kubernetes/kubernetes/blob/main/test/OWNERS",
		},
		{
			name:     "non-github url should be unchanged",
			url:      "https://viewsource.com/github/kubernetes/community/generator/app.go",
			expected: "https://viewsource.com/github/kubernetes/community/generator/app.go",
		},
	}
	for _, c := range cases {
		actual := githubURL(c.url)
		if actual != c.expected {
			t.Errorf("FAIL %s: got: '%s' but expected: '%s'", c.name, actual, c.expected)
		}
	}
}

func TestOrgRepoPath(t *testing.T) {
	cases := []struct {
		name     string
		url      string
		expected string
	}{
		{
			name:     "kubernetes-sigs root raw github url",
			url:      "https://raw.githubusercontent.com/kubernetes-sigs/boskos/main/OWNERS",
			expected: "kubernetes-sigs/boskos/OWNERS",
		},
		{
			name:     "kubernetes non-root raw github url",
			url:      "https://raw.githubusercontent.com/kubernetes/kubernetes/main/test/OWNERS",
			expected: "kubernetes/kubernetes/test/OWNERS",
		},
		{
			name:     "kubernetes github url",
			url:      "https://github.com/kubernetes/kubernetes/blob/main/test/OWNERS",
			expected: "kubernetes/kubernetes/test/OWNERS",
		},
		{
			name:     "non-github url should be unchanged",
			url:      "https://viewsource.com/github/kubernetes/community/generator/app.go",
			expected: "https://viewsource.com/github/kubernetes/community/generator/app.go",
		},
	}
	for _, c := range cases {
		actual := orgRepoPath(c.url)
		if actual != c.expected {
			t.Errorf("FAIL %s: got: '%s' but expected: '%s'", c.name, actual, c.expected)
		}
	}
}
