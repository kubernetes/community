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
	"log"
	"os"
	"path/filepath"
	"sort"
	"strings"
	"text/template"

	"gopkg.in/yaml.v2"
)

var (
	readmeTemplate = "readme.tmpl"
	listTemplate   = "list.tmpl"
	headerTemplate = "header.tmpl"

	sigsYamlFile  = "sigs.yaml"
	sigListOutput = "sig-list.md"
	indexFilename = "README.md"
	baseOutputDir = "generated"

	githubTeamNames = []string{"misc", "test-failures", "bugs", "feature-requests", "proposals", "pr-reviews", "api-reviews"}
	beginMarker     = "<!-- BEGIN CUSTOM CONTENT -->"
	endMarker       = "<!-- END CUSTOM CONTENT -->"
)

// Lead represents a lead engineer for a particular group. There are usually
// 2 per group.
type Lead struct {
	Name    string
	Company string
	GitHub  string
}

// Meeting represents a regular meeting for a group.
type Meeting struct {
	Day       string
	UTC       string
	PST       string
	Frequency string
}

// Contact represents the various contact points for a group.
type Contact struct {
	Slack            string
	MailingList      string `yaml:"mailing_list"`
	FullGitHubTeams  bool   `yaml:"full_github_teams"`
	GithubTeamPrefix string `yaml:"github_team_prefix"`
	GithubTeamNames  []string
}

// Group represents either a Special Interest Group (SIG) or a Working Group (WG)
type Group struct {
	Name                 string
	Dir                  string
	MissionStatement     string `yaml:"mission_statement"`
	Leads                []Lead
	Meetings             []Meeting
	MeetingURL           string `yaml:"meeting_url"`
	MeetingArchiveURL    string `yaml:"meeting_archive_url"`
	MeetingRecordingsURL string `yaml:"meeting_recordings_url"`
	Contact              Contact
}

// DirName returns the directory that a group's documentation will be
// generated into. It is composed of a prefix (sig for SIGs and wg for WGs),
// and a formatted version of the group's name (in kebab case).
func (e *Group) DirName(prefix string) string {
	return fmt.Sprintf("%s-%s", prefix, strings.ToLower(strings.Replace(e.Name, " ", "-", -1)))
}

// SetupGitHubTeams will iterate over all the possible teams available to a
// group (these are defined by the Kubernetes organization) and populate a
// list using the group's prefix.
func (e *Group) SetupGitHubTeams(prefix string) {
	ghPrefix := e.Contact.GithubTeamPrefix
	if ghPrefix == "" {
		ghPrefix = e.DirName(prefix)
	}

	for _, gtn := range githubTeamNames {
		e.Contact.GithubTeamNames = append(e.Contact.GithubTeamNames, fmt.Sprintf("%s-%s", ghPrefix, gtn))
	}
}

// Context is the context for the sigs.yaml file.
type Context struct {
	Sigs          []Group
	WorkingGroups []Group
}

func pathExists(path string) bool {
	_, err := os.Stat(path)
	return err == nil
}

func createDirIfNotExists(path string) error {
	if !pathExists(path) {
		return os.MkdirAll(path, 0755)
	}
	return nil
}

func getExistingContent(path string) (string, error) {
	capture := false
	var captured []string

	// NOTE: For some reason using bufio.Scanner with existing file pointer prepends
	// a bunch of null ^@ characters, so using to ioutil.ReadFile instead.
	content, err := ioutil.ReadFile(path)
	if err != nil {
		return "", err
	}

	for _, line := range strings.Split(string(content), "\n") {
		if strings.Contains(line, endMarker) {
			capture = false
		}
		if capture {
			captured = append(captured, line)
		}
		if strings.Contains(line, beginMarker) {
			capture = true
		}
	}

	return strings.Join(captured, "\n"), nil
}

func writeTemplate(templatePath, outputPath string, data interface{}) error {
	// set up template
	t, err := template.ParseFiles(templatePath, headerTemplate)
	if err != nil {
		return err
	}

	// create if not exists
	if !pathExists(outputPath) {
		_, err = os.Create(outputPath)
		if err != nil {
			return err
		}
	}

	// open file and truncate
	f, err := os.OpenFile(outputPath, os.O_RDWR, 0644)
	if err != nil {
		return err
	}
	defer f.Close()

	// get any existing content
	content, err := getExistingContent(outputPath)
	if err != nil {
		return err
	}

	// ensure file is empty
	f.Truncate(0)

	// generated content
	err = t.Execute(f, data)
	if err != nil {
		return err
	}

	// custom content block
	writeCustomContentBlock(f, content)

	return nil
}

func writeCustomContentBlock(f *os.File, content string) {
	lines := []string{beginMarker, "\n", content, "\n", endMarker, "\n"}
	for _, line := range lines {
		f.Write([]byte(line))
	}
}

func createGroupReadme(groups []Group, prefix string) error {
	// figure out if the user wants to generate one group
	var selectedGroupName *string
	if envVal, ok := os.LookupEnv("WHAT"); ok {
		selectedGroupName = &envVal
	}

	for _, group := range groups {
		group.Dir = group.DirName(prefix)
		// skip generation if the user specified only one group
		if selectedGroupName != nil && strings.HasSuffix(group.Dir, *selectedGroupName) == false {
			fmt.Printf("Skipping %s/README.md\n", group.Dir)
			continue
		}

		fmt.Printf("Generating %s/README.md\n", group.Dir)

		outputDir := filepath.Join(baseOutputDir, group.Dir)
		if err := createDirIfNotExists(outputDir); err != nil {
			return err
		}

		group.SetupGitHubTeams(prefix)

		outputPath := filepath.Join(outputDir, indexFilename)
		readmePath := fmt.Sprintf("%s_%s", prefix, readmeTemplate)
		if err := writeTemplate(readmePath, outputPath, group); err != nil {
			return err
		}
	}

	return nil
}

func main() {
	yamlData, err := ioutil.ReadFile(filepath.Join(baseOutputDir, sigsYamlFile))
	if err != nil {
		log.Fatal(err)
	}

	var ctx Context
	err = yaml.Unmarshal(yamlData, &ctx)
	if err != nil {
		log.Fatal(err)
	}

	sort.Slice(ctx.Sigs, func(i, j int) bool {
		return ctx.Sigs[i].Name <= ctx.Sigs[j].Name
	})

	sort.Slice(ctx.WorkingGroups, func(i, j int) bool {
		return ctx.WorkingGroups[i].Name <= ctx.WorkingGroups[j].Name
	})

	err = createGroupReadme(ctx.Sigs, "sig")
	if err != nil {
		log.Fatal(err)
	}

	err = createGroupReadme(ctx.WorkingGroups, "wg")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Generating sig-list.md")
	outputPath := filepath.Join(baseOutputDir, sigListOutput)
	err = writeTemplate(listTemplate, outputPath, ctx)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Finished generation!")
}
