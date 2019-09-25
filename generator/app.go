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

	yaml "gopkg.in/yaml.v3"
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

// FoldedString is a string that will be serialized in FoldedStyle by go-yaml
type FoldedString string

// MarshalYAML customizes how FoldedStrings will be serialized by go-yaml
func (x FoldedString) MarshalYAML() (interface{}, error) {
	return &yaml.Node{
		Kind:  yaml.ScalarNode,
		Style: yaml.FoldedStyle,
		Value: string(x),
	}, nil
}

// Person represents an individual person holding a role in a group.
type Person struct {
	GitHub  string
	Name    string
	Company string
}

// Meeting represents a regular meeting for a group.
type Meeting struct {
	Description   string
	Day           string
	Time          string
	TZ            string
	Frequency     string
	URL           string `yaml:",omitempty"`
	ArchiveURL    string `yaml:"archive_url,omitempty"`
	RecordingsURL string `yaml:"recordings_url,omitempty"`
}

// Contact represents the various contact points for a group.
type Contact struct {
	Slack              string       `yaml:",omitempty"`
	MailingList        string       `yaml:"mailing_list,omitempty"`
	PrivateMailingList string       `yaml:"private_mailing_list,omitempty"`
	GithubTeams        []GithubTeam `yaml:"teams,omitempty"`
}

// GithubTeam represents a specific Github Team.
type GithubTeam struct {
	Name        string
	Description string `yaml:",omitempty"`
}

// Subproject represents a specific subproject owned by the group
type Subproject struct {
	Name        string
	Description string   `yaml:",omitempty"`
	Contact     *Contact `yaml:",omitempty"`
	Owners      []string
	Meetings    []Meeting `yaml:",omitempty"`
}

// LeadershipGroup represents the different groups of leaders within a group
type LeadershipGroup struct {
	Chairs         []Person
	TechnicalLeads []Person `yaml:"tech_leads,omitempty"`
	EmeritusLeads  []Person `yaml:"emeritus_leads,omitempty"`
}

// PrefixToPersonMap returns a map of prefix to persons, useful for iteration over all persons
func (g *LeadershipGroup) PrefixToPersonMap() map[string][]Person {
	return map[string][]Person{
		"chair":         g.Chairs,
		"tech_lead":     g.TechnicalLeads,
		"emeritus_lead": g.EmeritusLeads,
	}
}

// Group represents either a Special Interest Group (SIG) or a Working Group (WG)
type Group struct {
	Dir              string
	Name             string
	MissionStatement FoldedString `yaml:"mission_statement,omitempty"`
	CharterLink      string       `yaml:"charter_link,omitempty"`
	StakeholderSIGs  []string     `yaml:"stakeholder_sigs,omitempty"`
	Label            string
	Leadership       LeadershipGroup `yaml:"leadership"`
	Meetings         []Meeting
	Contact          Contact
	Subprojects      []Subproject `yaml:",omitempty"`
}

// DirName returns the directory that a group's documentation will be
// generated into. It is composed of a prefix (sig for SIGs and wg for WGs),
// and a formatted version of the group's name (in kebab case).
func (g *Group) DirName(prefix string) string {
	return fmt.Sprintf("%s-%s", prefix, strings.ToLower(strings.Replace(g.Name, " ", "-", -1)))
}

// LabelName returns the expected label for a given group
func (g *Group) LabelName(prefix string) string {
	return strings.Replace(g.DirName(prefix), fmt.Sprintf("%s-", prefix), "", 1)
}

// Context is the context for the sigs.yaml file.
type Context struct {
	Sigs          []Group
	WorkingGroups []Group
	UserGroups    []Group
	Committees    []Group
}

func index(groups []Group, predicate func(Group) bool) int {
	for i, group := range groups {
		if predicate(group) {
			return i
		}
	}
	return -1
}

// PrefixToGroupMap returns a map of prefix to groups, useful for iteration over all groups
func (c *Context) PrefixToGroupMap() map[string][]Group {
	return map[string][]Group{
		"sig":       c.Sigs,
		"wg":        c.WorkingGroups,
		"ug":        c.UserGroups,
		"committee": c.Committees,
	}
}

// Sort sorts all lists within the Context struct
func (c *Context) Sort() {
	for _, groups := range c.PrefixToGroupMap() {
		sort.Slice(groups, func(i, j int) bool {
			return groups[i].Dir < groups[j].Dir
		})
		for _, group := range groups {
			sort.Strings(group.StakeholderSIGs)
			for _, people := range [][]Person{
				group.Leadership.Chairs,
				group.Leadership.TechnicalLeads,
				group.Leadership.EmeritusLeads} {
				sort.Slice(people, func(i, j int) bool {
					// This ensure OWNERS / OWNERS_ALIAS files are ordered by github
					return people[i].GitHub < people[j].GitHub
				})
			}
			sort.Slice(group.Meetings, func(i, j int) bool {
				return group.Meetings[i].Description < group.Meetings[j].Description
			})
			sort.Slice(group.Contact.GithubTeams, func(i, j int) bool {
				return group.Contact.GithubTeams[i].Name < group.Contact.GithubTeams[j].Name
			})
			sort.Slice(group.Subprojects, func(i, j int) bool {
				return group.Subprojects[i].Name < group.Subprojects[j].Name
			})
			for _, subproject := range group.Subprojects {
				if subproject.Contact != nil {
					sort.Slice(subproject.Contact.GithubTeams, func(i, j int) bool {
						return subproject.Contact.GithubTeams[i].Name < subproject.Contact.GithubTeams[j].Name
					})
				}
				sort.Strings(subproject.Owners)
				sort.Slice(subproject.Meetings, func(i, j int) bool {
					return subproject.Meetings[i].Description < subproject.Meetings[j].Description
				})
			}
		}
	}
}

// Validate returns a list of errors encountered while validating a Context
func (c *Context) Validate() []error {
	errors := []error{}
	people := make(map[string]Person)
	for prefix, groups := range c.PrefixToGroupMap() {
		for _, group := range groups {
			expectedDir := group.DirName(prefix)
			if expectedDir != group.Dir {
				errors = append(errors, fmt.Errorf("expected dir: %s, got: %s", expectedDir, group.Dir))
			}
			expectedLabel := group.LabelName(prefix)
			if expectedLabel != group.Label {
				errors = append(errors, fmt.Errorf("%s: expected label: %s, got: %s", group.Dir, expectedLabel, group.Label))
			}
			for prefix, persons := range group.Leadership.PrefixToPersonMap() {
				for _, person := range persons {
					if val, ok := people[person.GitHub]; ok {
						if val.Name != person.Name || val.Company != person.Company {
							errors = append(errors, fmt.Errorf("%s: %ss: expected person: %v, got: %v", group.Dir, prefix, val, person))
						}
					} else {
						people[person.GitHub] = person
					}
				}
			}
			if len(group.StakeholderSIGs) != 0 {
				if prefix == "wg" {
					for _, name := range group.StakeholderSIGs {
						if index(c.Sigs, func(g Group) bool { return g.Name == name }) == -1 {
							errors = append(errors, fmt.Errorf("%s: invalid stakeholder sig name %s", group.Dir, name))
						}
					}
				} else {
					errors = append(errors, fmt.Errorf("%s: only WGs may have stakeholder_sigs", group.Dir))
				}
			}
			if prefix == "sig" {
				if group.CharterLink == "" {
					errors = append(errors, fmt.Errorf("%s: has no charter", group.Dir))
				}
				// TODO(spiffxp): is this required though?
				if group.MissionStatement == "" {
					errors = append(errors, fmt.Errorf("%s: has no mission statement", group.Dir))
				}
				if len(group.Subprojects) == 0 {
					errors = append(errors, fmt.Errorf("%s: has no subprojects", group.Dir))
				}
			}
		}
	}
	return errors
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
	"tzUrlEncode": tzURLEncode,
	"trimSpace":   strings.TrimSpace,
}

// tzUrlEncode returns a url encoded string without the + shortcut. This is
// required as the timezone conversion site we are using doesn't recognize + as
// a valid url escape character.
func tzURLEncode(tz string) string {
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

// readSigsYaml decodes yaml stored in a file at path into the
// specified yaml.Node
func readYaml(path string, data interface{}) error {
	file, err := os.Open(path)
	if err != nil {
		return err
	}
	defer file.Close()
	decoder := yaml.NewDecoder(file)
	decoder.KnownFields(true)
	return decoder.Decode(data)
}

// writeSigsYaml writes the specified data to a file at path
// indent is set to 2 spaces
func writeYaml(data interface{}, path string) error {
	file, err := os.Create(path)
	if err != nil {
		return err
	}
	defer file.Close()
	enc := yaml.NewEncoder(file)
	enc.SetIndent(2)
	return enc.Encode(data)
}

func main() {
	yamlPath := filepath.Join(baseGeneratorDir, sigsYamlFile)
	var ctx Context

	err := readYaml(yamlPath, &ctx)
	if err != nil {
		log.Fatal(err)
	}

	ctx.Sort()

	fmt.Printf("Validating %s\n", yamlPath)
	errs := ctx.Validate()
	if len(errs) != 0 {
		for _, err := range errs {
			fmt.Printf("ERROR: %s\n", err.Error())
		}
		os.Exit(1)
	}

	// Write the Context struct back to yaml to enforce formatting
	err = writeYaml(&ctx, yamlPath)
	if err != nil {
		log.Fatal(err)
	}

	for prefix, groups := range ctx.PrefixToGroupMap() {
		err = createGroupReadme(groups, prefix)
		if err != nil {
			log.Fatal(err)
		}
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
