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
	"net/url"
	"os"
	"path/filepath"
	"sort"
	"strings"
	"text/template"

	"gopkg.in/yaml.v2"
)

const (
	readmeTemplate  = "readme.tmpl"
	listTemplate    = "list.tmpl"
	aliasesTemplate = "aliases.tmpl"
	headerTemplate  = "header.tmpl"

	sigsYamlFile  = "sigs.yaml"
	sigListOutput = "sig-list.md"
	aliasesOutput = "OWNERS_ALIASES"
	indexFilename = "README.md"

	beginCustomMarkdown = "<!-- BEGIN CUSTOM CONTENT -->"
	endCustomMarkdown   = "<!-- END CUSTOM CONTENT -->"
	beginCustomYaml     = "## BEGIN CUSTOM CONTENT"
	endCustomYaml       = "## END CUSTOM CONTENT"
)

var (
	baseGeneratorDir = ""
	templateDir      = "generator"
)

// Person represents an individual person holding a role in a group.
type Person struct {
	Name    string
	Company string
	GitHub  string
}

// Meeting represents a regular meeting for a group.
type Meeting struct {
	Description   string
	Day           string
	Time          string
	TZ            string
	Frequency     string
	URL           string
	ArchiveURL    string `yaml:"archive_url"`
	RecordingsURL string `yaml:"recordings_url"`
}

// Contact represents the various contact points for a group.
type Contact struct {
	Slack       string
	MailingList string        `yaml:"mailing_list"`
	GithubTeams []GithubTeams `yaml:"teams"`
}

// GithubTeams represents a specific Github Team.
type GithubTeams struct {
	Name        string
	Description string
}

// Subproject represents a specific subproject owned by the group
type Subproject struct {
	Name        string
	Description string
	Owners      []string
	Meetings    []Meeting
}

// LeadershipGroup represents the different groups of leaders within a group
type LeadershipGroup struct {
	Chairs         []Person
	TechnicalLeads []Person `yaml:"tech_leads"`
	EmeritusLeads  []Person `yaml:"emeritus_leads"`
}

// Group represents either a Special Interest Group (SIG) or a Working Group (WG)
type Group struct {
	Name             string
	Dir              string
	MissionStatement string `yaml:"mission_statement"`
	Label            string
	Leadership       LeadershipGroup `yaml:"leadership"`
	Meetings         []Meeting
	Contact          Contact
	Subprojects      []Subproject
}

// DirName returns the directory that a group's documentation will be
// generated into. It is composed of a prefix (sig for SIGs and wg for WGs),
// and a formatted version of the group's name (in kebab case).
func (e *Group) DirName(prefix string) string {
	return fmt.Sprintf("%s-%s", prefix, strings.ToLower(strings.Replace(e.Name, " ", "-", -1)))
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

func getExistingContent(path string, fileFormat string) (string, error) {
	capture := false
	var captured []string

	beginMarker := ""
	endMarker := ""
	switch fileFormat {
	case "markdown":
		beginMarker = beginCustomMarkdown
		endMarker = endCustomMarkdown
	case "yaml":
		beginMarker = beginCustomYaml
		endMarker = endCustomYaml
	case "":
		return "", nil
	}

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

var funcMap = template.FuncMap{
	"tzUrlEncode": tzUrlEncode,
}

// tzUrlEncode returns a url encoded string without the + shortcut. This is
// required as the timezone conversion site we are using doesn't recognize + as
// a valid url escape character.
func tzUrlEncode(tz string) string {
	return strings.Replace(url.QueryEscape(tz), "+", "%20", -1)
}

func writeTemplate(templatePath, outputPath string, fileFormat string, data interface{}) error {
	// set up template
	t, err := template.New(filepath.Base(templatePath)).
		Funcs(funcMap).
		ParseFiles(templatePath, filepath.Join(baseGeneratorDir, templateDir, headerTemplate))
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
	content, err := getExistingContent(outputPath, fileFormat)
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

	writeCustomContentBlock(f, content, fileFormat)

	return nil
}

func writeCustomContentBlock(f *os.File, content string, fileFormat string) {
	beginMarker := ""
	endMarker := ""
	switch fileFormat {
	case "markdown":
		beginMarker = beginCustomMarkdown
		endMarker = endCustomMarkdown
	case "yaml":
		beginMarker = beginCustomYaml
		endMarker = endCustomYaml
	case "":
		return
	}

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

		outputDir := filepath.Join(baseGeneratorDir, group.Dir)
		if err := createDirIfNotExists(outputDir); err != nil {
			return err
		}

		outputPath := filepath.Join(outputDir, indexFilename)
		readmePath := filepath.Join(baseGeneratorDir, templateDir, fmt.Sprintf("%s_%s", prefix, readmeTemplate))
		if err := writeTemplate(readmePath, outputPath, "markdown", group); err != nil {
			return err
		}
	}

	return nil
}

func main() {
	yamlData, err := ioutil.ReadFile(filepath.Join(baseGeneratorDir, sigsYamlFile))
	if err != nil {
		log.Fatal(err)
	}

	var ctx Context
	err = yaml.Unmarshal(yamlData, &ctx)
	if err != nil {
		log.Fatal(err)
	}

	sort.Slice(ctx.Sigs, func(i, j int) bool {
		return strings.ToLower(ctx.Sigs[i].Name) <= strings.ToLower(ctx.Sigs[j].Name)
	})

	sort.Slice(ctx.WorkingGroups, func(i, j int) bool {
		return strings.ToLower(ctx.WorkingGroups[i].Name) <= strings.ToLower(ctx.WorkingGroups[j].Name)
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
	outputPath := filepath.Join(baseGeneratorDir, sigListOutput)
	err = writeTemplate(filepath.Join(baseGeneratorDir, templateDir, listTemplate), outputPath, "markdown", ctx)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Generating OWNERS_ALIASES")
	outputPath = filepath.Join(baseGeneratorDir, aliasesOutput)
	err = writeTemplate(filepath.Join(baseGeneratorDir, templateDir, aliasesTemplate), outputPath, "yaml", ctx)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Finished generation!")
}
