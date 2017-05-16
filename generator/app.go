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
	"time"

	"gopkg.in/yaml.v2"
)

var (
	sigsYamlFile    = "sigs.yaml"
	templateDir     = "generator"
	indexTemplate   = filepath.Join(templateDir, "sig_index.tmpl")
	listTemplate    = filepath.Join(templateDir, "sig_list.tmpl")
	headerTemplate  = filepath.Join(templateDir, "header.tmpl")
	footerTemplate  = filepath.Join(templateDir, "footer.tmpl")
	sigListOutput   = "sig-list.md"
	sigIndexOutput  = "README.md"
	githubTeamNames = []string{"misc", "test-failures", "bugs", "feature-requests", "proposals", "pr-reviews", "api-reviews"}
)

type Lead struct {
	Name    string
	Company string
	GitHub  string
}

type Meeting struct {
	Day       string
	UTC       string
	PST       string
	Frequency string
}

type Contact struct {
	Slack            string
	MailingList      string `yaml:"mailing_list"`
	FullGitHubTeams  bool   `yaml:"full_github_teams"`
	GithubTeamPrefix string `yaml:"github_team_prefix"`
	GithubTeamNames  []string
}

type Sig struct {
	Page
	Name              string
	Dir               string
	MissionStatement  string `yaml:"mission_statement"`
	Leads             []Lead
	Meetings          []Meeting
	MeetingURL        string `yaml:"meeting_url"`
	MeetingArchiveURL string `yaml:"meeting_archive_url"`
	Contact           Contact
}

type SigEntries struct {
	Page
	Sigs []Sig
}

type Page struct {
	LastGenerated string
}

func (slice SigEntries) Len() int {
	return len(slice.Sigs)
}

func (slice SigEntries) Less(i, j int) bool {
	return slice.Sigs[i].Name < slice.Sigs[j].Name
}

func (slice SigEntries) Swap(i, j int) {
	slice.Sigs[i], slice.Sigs[j] = slice.Sigs[j], slice.Sigs[i]
}

func createDirIfNotExists(path string) error {
	_, err := os.Stat(path)
	if os.IsNotExist(err) {
		fmt.Printf("%s directory does not exist, creating\n", path)
		return os.Mkdir(path, 0755)
	}
	return nil
}

func writeTemplate(templatePath, outputPath string, data interface{}) error {
	// set up template
	t, err := template.ParseFiles(templatePath, headerTemplate, footerTemplate)
	if err != nil {
		return err
	}

	// open file and truncate
	f, err := os.OpenFile(outputPath, os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	defer f.Close()
	f.Truncate(0)

	// write template output to file
	err = t.Execute(f, data)
	if err != nil {
		return err
	}

	fmt.Printf("Generated %s\n", outputPath)
	return nil
}

func lastGenerated() string {
	return time.Now().Format("Mon Jan 2 2006 15:04:05")
}

func createReadmeFiles(ctx SigEntries) error {
	for _, sig := range ctx.Sigs {
		dirName := fmt.Sprintf("sig-%s", strings.ToLower(strings.Replace(sig.Name, " ", "-", -1)))

		createDirIfNotExists(dirName)

		sig.LastGenerated = lastGenerated()

		prefix := sig.Contact.GithubTeamPrefix
		if prefix == "" {
			prefix = dirName
		}

		for _, gtn := range githubTeamNames {
			sig.Contact.GithubTeamNames = append(sig.Contact.GithubTeamNames, fmt.Sprintf("%s-%s", prefix, gtn))
		}

		outputPath := fmt.Sprintf("%s/%s", dirName, sigIndexOutput)
		if err := writeTemplate(indexTemplate, outputPath, sig); err != nil {
			return err
		}
	}

	return nil
}

func createListFile(ctx SigEntries) error {
	ctx.LastGenerated = lastGenerated()
	return writeTemplate(listTemplate, sigListOutput, ctx)
}

func main() {
	yamlData, err := ioutil.ReadFile(sigsYamlFile)
	if err != nil {
		log.Fatal(err)
	}

	var ctx SigEntries
	err = yaml.Unmarshal(yamlData, &ctx)
	if err != nil {
		log.Fatal(err)
	}

	sort.Sort(ctx)

	err = createReadmeFiles(ctx)
	if err != nil {
		log.Fatal(err)
	}

	err = createListFile(ctx)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Finished generation!")
}
