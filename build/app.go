package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"time"

	"github.com/cbroglie/mustache"
	"gopkg.in/yaml.v2"
)

var (
	sigsYamlFile   = "sigs.yaml"
	templateDir    = "build"
	indexTemplate  = fmt.Sprintf("%s/sig_index.mustache", templateDir)
	listTemplate   = fmt.Sprintf("%s/sig_list.mustache", templateDir)
	footerTemplate = fmt.Sprintf("%s/footer.mustache", templateDir)
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

type Page struct {
	LastGenerated string
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

func constructFooter() (data string, err error) {
	template, err := mustache.ParseFile(footerTemplate)
	if err != nil {
		return
	}

	ctx := Page{LastGenerated: time.Now().Format("Mon Jan 2 2006 15:04:05")}

	data, err = template.Render(ctx)
	if err != nil {
		return
	}

	err = ioutil.WriteFile(sigListOutput, []byte(data), 0644)
	if err != nil {
		return
	}

	return
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

		footer, err := constructFooter()
		if err != nil {
			return err
		}

		filePath := fmt.Sprintf("%s/%s", sig.Dir, sigIndexOutput)
		err = ioutil.WriteFile(filePath, []byte(data+footer), 0644)
		if err != nil {
			return err
		}

		fmt.Printf("Generated %s\n", filePath)
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

	footer, err := constructFooter()
	if err != nil {
		return err
	}

	err = ioutil.WriteFile(sigListOutput, []byte(data+footer), 0644)
	if err != nil {
		return err
	}

	fmt.Printf("Generated %s\n", sigListOutput)
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

	fmt.Println("Finished generation!")
}
