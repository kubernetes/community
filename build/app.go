package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"github.com/cbroglie/mustache"
	"gopkg.in/yaml.v2"
)

var (
	sigsYamlFile   = "sigs.yaml"
	templateDir    = "build"
	indexTemplate  = fmt.Sprintf("%s/sig_index.mustache", templateDir)
	listTemplate   = fmt.Sprintf("%s/sig_list.mustache", templateDir)
	sigListOutput  = "sig-list.md"
	sigIndexOutput = "README.md"
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
	Slack       string
	MailingList string `yaml:"mailing_list"`
	GitHubTeam  string `yaml:"github_team"`
}

type Sig struct {
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
	Sigs []Sig
}

func dirExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return true, err
}

func createReadmeFiles(ctx SigEntries) error {
	template, err := mustache.ParseFile(indexTemplate)
	if err != nil {
		return err
	}

	for _, sig := range ctx.Sigs {
		data, err := template.Render(sig)
		if err != nil {
			return err
		}

		exists, err := dirExists(sig.Dir)
		if err != nil {
			return err
		}
		if !exists {
			err = os.Mkdir(sig.Dir, 0755)
			if err != nil {
				return err
			}
		}

		err = ioutil.WriteFile(fmt.Sprintf("%s/%s", sig.Dir, sigIndexOutput), []byte(data), 0644)
		if err != nil {
			return err
		}
	}

	return nil
}

func createListFile(ctx SigEntries) error {
	template, err := mustache.ParseFile(listTemplate)
	if err != nil {
		return err
	}

	data, err := template.Render(ctx)
	if err != nil {
		return err
	}

	err = ioutil.WriteFile(sigListOutput, []byte(data), 0644)
	if err != nil {
		return err
	}

	return nil
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

	err = createReadmeFiles(ctx)
	if err != nil {
		log.Fatal(err)
	}

	err = createListFile(ctx)
	if err != nil {
		log.Fatal(err)
	}
}
