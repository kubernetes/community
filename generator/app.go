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
	"context"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"regexp"
	"sort"
	"strconv"
	"strings"
	"text/template"
	"time"

	"k8s.io/enhancements/api"

	"github.com/google/go-github/v32/github"

	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing"
	"github.com/go-git/go-git/v5/storage/memory"

	yaml "gopkg.in/yaml.v3"

	"golang.org/x/mod/semver"
)

const (
	readmeTemplate            = "readme.tmpl"
	listTemplate              = "list.tmpl"
	aliasesTemplate           = "aliases.tmpl"
	liaisonsTemplate          = "liaisons.tmpl"
	headerTemplate            = "header.tmpl"
	annualReportIssueTemplate = "annual-report/github_issue.tmpl"
	annualReportSIGTemplate   = "annual-report/sig_report.tmpl"
	annualReportWGTemplate    = "annual-report/wg_report.tmpl"

	sigsYamlFile     = "sigs.yaml"
	sigListOutput    = "sig-list.md"
	aliasesOutput    = "OWNERS_ALIASES"
	indexFilename    = "README.md"
	liaisonsFilename = "liaisons.md"

	beginCustomMarkdown = "<!-- BEGIN CUSTOM CONTENT -->"
	endCustomMarkdown   = "<!-- END CUSTOM CONTENT -->"
	beginCustomYaml     = "## BEGIN CUSTOM CONTENT"
	endCustomYaml       = "## END CUSTOM CONTENT"

	regexRawGitHubURL = "https://raw.githubusercontent.com/(?P<org>[^/]+)/(?P<repo>[^/]+)/(?P<branch>[^/]+)/(?P<path>.*)"
	regexGitHubURL    = "https://github.com/(?P<org>[^/]+)/(?P<repo>[^/]+)/(blob|tree)/(?P<branch>[^/]+)/(?P<path>.*)"

	// For KEPs automation
	kepURL = "https://storage.googleapis.com/k8s-keps/keps.json"
	// For Subprojects automation
	communityRepoURL = "https://github.com/kubernetes/community.git"
	localRepoPath    = "."
)

var (
	baseGeneratorDir           = ""
	templateDir                = "generator"
	releases                   = Releases{}
	cachedKEPs                 = []api.Proposal{}
	repo                       = &git.Repository{}
	annualReportYear           = time.Time{}
	currentYear                = time.Time{}
	commitFromAnnualReportYear = &plumbing.Hash{}
	commitFromCurrentYear      = &plumbing.Hash{}
)

type Releases struct {
	Latest         string
	LatestMinusOne string
	LatestMinusTwo string
}

// TODO: improve as suggested in https://github.com/kubernetes/community/pull/7038#discussion_r1069456087
func getLastThreeK8sReleases() (Releases, error) {
	ctx := context.Background()
	client := github.NewClient(nil)

	releases, _, err := client.Repositories.ListReleases(ctx, "kubernetes", "kubernetes", nil)
	if err != nil {
		return Releases{}, err
	}
	var result Releases
	for _, release := range releases {
		if release.GetPrerelease() || release.GetDraft() {
			continue
		}
		if result.Latest == "" {
			result.Latest = semver.MajorMinor(release.GetTagName())
			continue
		}
		if result.LatestMinusOne == "" {
			result.LatestMinusOne = semver.MajorMinor(release.GetTagName())
			continue
		}
		if result.LatestMinusTwo == "" {
			result.LatestMinusTwo = semver.MajorMinor(release.GetTagName())
			break
		}
	}

	return result, nil
}

func getReleases() Releases {
	return releases
}

func fetchKEPs() error {
	url, err := url.Parse(kepURL)
	if err != nil {
		return fmt.Errorf("Error parsing url: %v", err)
	}

	req, err := http.NewRequest("GET", url.String(), nil)
	if err != nil {
		return fmt.Errorf("Error creating request: %v", err)
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return fmt.Errorf("Error fetching KEPs: %v", err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("Error reading KEPs body: %v", err)
	}

	err = json.Unmarshal(body, &cachedKEPs)
	if err != nil {
		return fmt.Errorf("Error unmarshalling KEPs: %v", err)
	}
	return nil
}

func stageIfKEPsIsWorkedInReleases(kepMilestone api.Milestone, releases Releases) (api.Stage, bool) {
	if strings.HasSuffix(kepMilestone.Stable, releases.Latest) || strings.HasSuffix(kepMilestone.Stable, releases.LatestMinusOne) || strings.HasSuffix(kepMilestone.Stable, releases.LatestMinusTwo) {
		return api.StableStage, true
	}

	if strings.HasSuffix(kepMilestone.Beta, releases.Latest) || strings.HasSuffix(kepMilestone.Beta, releases.LatestMinusOne) || strings.HasSuffix(kepMilestone.Beta, releases.LatestMinusTwo) {
		return api.BetaStage, true
	}

	if strings.HasSuffix(kepMilestone.Alpha, releases.Latest) || strings.HasSuffix(kepMilestone.Alpha, releases.LatestMinusOne) || strings.HasSuffix(kepMilestone.Alpha, releases.LatestMinusTwo) {
		return api.AlphaStage, true
	}

	return "", false
}

func filterKEPs(owningSig string, releases Releases) (map[string][]api.Proposal, error) {
	// TODO(palnabarun): Hack to allow unprefixed version strings in KEPs.
	// Once all KEPs are updated to use the prefixed version strings, this can be removed.
	// See: https://github.com/kubernetes/community/issues/7213#issuecomment-1484964640
	unPrefixedReleases := Releases{
		Latest:         strings.TrimPrefix(releases.Latest, "v"),
		LatestMinusOne: strings.TrimPrefix(releases.LatestMinusOne, "v"),
		LatestMinusTwo: strings.TrimPrefix(releases.LatestMinusTwo, "v"),
	}

	kepsByStage := make(map[string][]api.Proposal)
	for _, kep := range cachedKEPs {
		if kep.OwningSIG == owningSig {
			stage, ok := stageIfKEPsIsWorkedInReleases(kep.Milestone, unPrefixedReleases)
			if !ok {
				continue
			}
			kepsByStage[string(stage)] = append(kepsByStage[string(stage)], kep)
		}
	}
	return kepsByStage, nil
}

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
	Company string `yaml:"company,omitempty"`
	// NOTE: this isn't displayed in the markdown files by design
	// We collect this info for purposes like the leads@kubernetes.io list
	// You should reach out to SIGs via the group mailinglists, slack, and github
	Email string `yaml:"email,omitempty"`
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
	CalendarURL   string `yaml:"calendar_url,omitempty"`
}

// Contact represents the various contact points for a group.
type Contact struct {
	Slack              string       `yaml:",omitempty"`
	MailingList        string       `yaml:"mailing_list,omitempty"`
	PrivateMailingList string       `yaml:"private_mailing_list,omitempty"`
	GithubTeams        []GithubTeam `yaml:"teams,omitempty"`
	Liaison            Person       `yaml:"liaison,omitempty"`
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
	Leads       []Person  `yaml:",omitempty"`
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

// Owners returns a sorted and de-duped list of owners for a LeadershipGroup
func (g *LeadershipGroup) Owners() []Person {
	o := append(g.Chairs, g.TechnicalLeads...)

	// Sort
	sort.Slice(o, func(i, j int) bool {
		return o[i].GitHub < o[j].GitHub
	})

	// De-dupe
	seen := make(map[string]struct{}, len(o))
	i := 0
	for _, p := range o {
		if _, ok := seen[p.GitHub]; ok {
			continue
		}
		seen[p.GitHub] = struct{}{}
		o[i] = p
		i++
	}
	return o[:i]
}

// Group represents either a Special Interest Group (SIG) or a Working Group (WG)
type Group struct {
	Dir              string
	Prefix           string `yaml:",omitempty"`
	Name             string
	MissionStatement FoldedString `yaml:"mission_statement,omitempty"`
	CharterLink      string       `yaml:"charter_link,omitempty"`
	ReportingWGs     []WGName     `yaml:"-"` // populated by Context#Complete()
	StakeholderSIGs  []SIGName    `yaml:"stakeholder_sigs,omitempty"`
	Label            string
	Leadership       LeadershipGroup `yaml:"leadership"`
	Meetings         []Meeting
	Contact          Contact
	Subprojects      []Subproject              `yaml:",omitempty"`
	KEPs             map[string][]api.Proposal `yaml:",omitempty"`
}

type WGName string

func (n WGName) DirName() string {
	return DirName("wg", string(n))
}

type SIGName string

func (n SIGName) DirName() string {
	return DirName("sig", string(n))
}

// DirName returns the directory that a group's documentation will be
// generated into. It is composed of a prefix (sig for SIGs and wg for WGs),
// and a formatted version of the group's name (in kebab case).
func (g *Group) DirName(prefix string) string {
	return DirName(prefix, g.Name)
}

func DirName(prefix, name string) string {
	return fmt.Sprintf("%s-%s", prefix, strings.ToLower(strings.Replace(name, " ", "-", -1)))
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

// Complete populates derived portions of the Context struct
func (c *Context) Complete() {
	// Copy working group names into ReportingWGs list of their stakeholder sigs
	for _, wg := range c.WorkingGroups {
		for _, stakeholderSIG := range wg.StakeholderSIGs {
			for i, sig := range c.Sigs {
				if sig.Name == string(stakeholderSIG) {
					c.Sigs[i].ReportingWGs = append(c.Sigs[i].ReportingWGs, WGName(wg.Name))
				}
			}
		}
	}
}

// Sort sorts all lists within the Context struct
func (c *Context) Sort() {
	for _, groups := range c.PrefixToGroupMap() {
		sort.Slice(groups, func(i, j int) bool {
			return groups[i].Dir < groups[j].Dir
		})
		for _, group := range groups {
			sort.Slice(group.ReportingWGs, func(i, j int) bool {
				return group.ReportingWGs[i] < group.ReportingWGs[j]
			})
			sort.Slice(group.StakeholderSIGs, func(i, j int) bool {
				return group.StakeholderSIGs[i] < group.StakeholderSIGs[j]
			})
			for _, people := range [][]Person{
				group.Leadership.Chairs,
				group.Leadership.TechnicalLeads,
				group.Leadership.EmeritusLeads} {
				sort.Slice(people, func(i, j int) bool {
					// This ensure OWNERS / OWNERS_ALIASES files are ordered by github
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
				sort.Slice(subproject.Leads, func(i, j int) bool {
					return subproject.Leads[i].GitHub < subproject.Leads[j].GitHub
				})
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
	// github to Person info
	// TODO: this would probably be a better config format? to avoid duplicating
	// people with potentially differing info, versus referring to leads by
	// github handle within each SIG and then keeping this map alongside the SIGs
	// This could break external tooling parsing the file though.
	people := make(map[string]Person)
	reRawGitHubURL := regexp.MustCompile(regexRawGitHubURL)
	reGitHubURL := regexp.MustCompile(regexGitHubURL)
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
						// non-emeritus must have email and company set
						if prefix != "emeritus_lead" {
							// email and company must match across groups
							if val.Email != person.Email {
								errors = append(errors, fmt.Errorf("%s: %s email: %q does not match other entries %q", group.Dir, val.GitHub, val.Email, person.Email))
							}
							if val.Company != person.Company {
								errors = append(errors, fmt.Errorf("%s: %s company: %q does not match other entries %q", group.Dir, val.GitHub, val.Company, person.Company))
							}
						}
						// all entries should have matching github + name, emeritus or not
						if val.Name != person.Name {
							errors = append(errors, fmt.Errorf("%s: %s: expected person: %v, got: %v", group.Dir, prefix, val, person))
						}
					} else if prefix != "emeritus_lead" {
						people[person.GitHub] = person
						// email and company must be set for leads
						if person.Email == "" {
							errors = append(errors, fmt.Errorf("%s: %s: email is empty but should be set", group.Dir, person.GitHub))
						}
						if person.Company == "" {
							errors = append(errors, fmt.Errorf("%s: %s: company is empty but should be set", group.Dir, person.GitHub))
						}
					}
					if person.Name == "" {
						errors = append(errors, fmt.Errorf("%s: %s: name is empty but should be set", group.Dir, person.GitHub))
					}

					if prefix == "emeritus_lead" && person.Company != "" {
						errors = append(errors, fmt.Errorf("%s: emeritus leads should not have company specified; company specified for: %s", group.Dir, person.Name))
					}
				}
			}
			if len(group.ReportingWGs) != 0 {
				if prefix == "sig" {
					for _, name := range group.ReportingWGs {
						if index(c.WorkingGroups, func(g Group) bool { return g.Name == string(name) }) == -1 {
							errors = append(errors, fmt.Errorf("%s: invalid reporting working group name %s", group.Dir, name))
						}
					}
				} else {
					errors = append(errors, fmt.Errorf("%s: only SIGs may have reporting WGs", group.Dir))
				}
			}
			if len(group.StakeholderSIGs) != 0 {
				if prefix == "wg" {
					for _, name := range group.StakeholderSIGs {
						if index(c.Sigs, func(g Group) bool { return g.Name == string(name) }) == -1 {
							errors = append(errors, fmt.Errorf("%s: invalid stakeholder sig name %s", group.Dir, name))
						}
					}
				} else {
					errors = append(errors, fmt.Errorf("%s: only WGs may have stakeholder_sigs", group.Dir))
				}
			} else {
				if prefix == "wg" {
					errors = append(errors, fmt.Errorf("%s: WGs must have stakeholder_sigs", group.Dir))
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
			if prefix != "committee" && prefix != "sig" {
				if len(group.Subprojects) > 0 {
					errors = append(errors, fmt.Errorf("%s: only sigs and committees can own code / have subprojects, found: %v", group.Dir, group.Subprojects))
				}
			}
			for _, subproject := range group.Subprojects {
				if len(subproject.Owners) == 0 {
					errors = append(errors, fmt.Errorf("%s/%s: subproject has no owners", group.Dir, subproject.Name))
				}
				for _, ownerURL := range subproject.Owners {
					if !reRawGitHubURL.MatchString(ownerURL) && !reGitHubURL.MatchString(ownerURL) {
						errors = append(errors, fmt.Errorf("%s/%s: subproject owners should match regexp %s, found: %s", group.Dir, subproject.Name, regexRawGitHubURL, ownerURL))
					}
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
	"tzUrlEncode":                 tzURLEncode,
	"trimSpace":                   strings.TrimSpace,
	"trimSuffix":                  strings.TrimSuffix,
	"githubURL":                   githubURL,
	"orgRepoPath":                 orgRepoPath,
	"now":                         time.Now,
	"lastYear":                    lastYear,
	"toUpper":                     strings.ToUpper,
	"filterKEPs":                  filterKEPs,
	"getReleases":                 getReleases,
	"getCategorizedSubprojects":   getCategorizedSubprojects,
	"getCategorizedWorkingGroups": getCategorizedWorkingGroups,
}

// lastYear returns the last year as a string
func lastYear() string {
	return time.Now().AddDate(-1, 0, 0).Format("2006")
}

// githubURL converts a raw GitHub url (links directly to file contents) into a
// regular GitHub url (links to Code view for file), otherwise returns url untouched
func githubURL(url string) string {
	re := regexp.MustCompile(regexRawGitHubURL)
	mat := re.FindStringSubmatchIndex(url)
	if mat == nil {
		return url
	}
	result := re.ExpandString([]byte{}, "https://github.com/${org}/${repo}/blob/${branch}/${path}", url, mat)
	return string(result)
}

// orgRepoPath converts either
//   - a regular GitHub url of form https://github.com/org/repo/blob/branch/path/to/file
//   - a raw GitHub url of form https://raw.githubusercontent.com/org/repo/branch/path/to/file
//
// to a string of form 'org/repo/path/to/file'
func orgRepoPath(url string) string {
	for _, regex := range []string{regexRawGitHubURL, regexGitHubURL} {
		re := regexp.MustCompile(regex)
		mat := re.FindStringSubmatchIndex(url)
		if mat == nil {
			continue
		}
		result := re.ExpandString([]byte{}, "${org}/${repo}/${path}", url, mat)
		return string(result)
	}
	return url
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

func createAnnualReportIssue(groups []Group, prefix string) error {
	// figure out if the user wants to generate one group
	var selectedGroupName *string
	if envVal, ok := os.LookupEnv("WHAT"); ok {
		selectedGroupName = &envVal
	}

	for _, group := range groups {
		switch prefix {
		case "sig":
			group.Prefix = "sig"
		case "wg":
			group.Prefix = "wg"
		default:
			continue

		}

		outputDir := filepath.Join(baseGeneratorDir, "generator/generated")

		// skip generation if the user specified only one group
		if selectedGroupName != nil && !strings.HasSuffix(group.Dir, *selectedGroupName) {
			fmt.Printf("Skipping %s/%s_%s.md\n", outputDir, lastYear(), group.Dir)
			continue
		}

		fmt.Printf("Generating %s/%s_%s.md\n", outputDir, lastYear(), group.Dir)
		if err := createDirIfNotExists(outputDir); err != nil {
			return err
		}

		outputPath := filepath.Join(outputDir, fmt.Sprintf("%s_%s.md", lastYear(), group.Dir))
		templatePath := filepath.Join(baseGeneratorDir, templateDir, annualReportIssueTemplate)
		if err := writeTemplate(templatePath, outputPath, "", group); err != nil {
			return err
		}
	}

	return nil
}

func createAnnualReport(groups []Group, prefix string) error {
	// figure out if the user wants to generate one group
	var selectedGroupName *string
	var templateFile string
	if envVal, ok := os.LookupEnv("WHAT"); ok {
		selectedGroupName = &envVal
	}

	for _, group := range groups {
		switch prefix {
		case "sig":
			group.Prefix = "sig"
			templateFile = annualReportSIGTemplate
		case "wg":
			group.Prefix = "wg"
			templateFile = annualReportWGTemplate
		default:
			continue

		}

		outputDir := filepath.Join(baseGeneratorDir, group.Dir)

		// skip generation if the user specified only one group
		if selectedGroupName != nil && !strings.HasSuffix(group.Dir, *selectedGroupName) {
			fmt.Printf("Skipping %s/annual-report-%s.md\n", outputDir, lastYear())
			continue
		}

		fmt.Printf("Generating %s/annual-report-%s.md\n", outputDir, lastYear())
		if err := createDirIfNotExists(outputDir); err != nil {
			return err
		}

		outputPath := filepath.Join(outputDir, fmt.Sprintf("annual-report-%s.md", lastYear()))
		templatePath := filepath.Join(baseGeneratorDir, templateDir, templateFile)
		if err := writeTemplate(templatePath, outputPath, "", group); err != nil {
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

// get the first commit on a given date
func getCommitByDate(repo *git.Repository, date time.Time) (*plumbing.Hash, error) {
	// Get the commit iterator
	iterator, err := repo.Log(&git.LogOptions{Order: git.LogOrderCommitterTime})
	if err != nil {
		return nil, err
	}

	// Iterate through the commits
	var commit *plumbing.Hash
	for {
		// Get the next commit
		c, err := iterator.Next()
		if err != nil {
			if err == io.EOF {
				break
			}
			return nil, err
		}

		// Check if the commit date is less than or equal to the specified date
		if c.Committer.When.Before(date) || c.Committer.When.Equal(date) {
			commit = &c.Hash
			break
		}
	}

	return commit, nil
}

// get the "sigs.yaml" file from a given commit
func getFileFromCommit(repo *git.Repository, commit *plumbing.Hash, filename string) ([]byte, error) {
	// Get the commit object
	obj, err := repo.CommitObject(*commit)
	if err != nil {
		return nil, err
	}

	// Get the commit tree
	tree, err := obj.Tree()
	if err != nil {
		return nil, err
	}

	// Get the file from the tree
	entry, err := tree.FindEntry(filename)
	if err != nil {
		return nil, err
	}

	// Get the file content
	file, err := repo.BlobObject(entry.Hash)
	if err != nil {
		return nil, err
	}

	reader, err := file.Reader()
	if err != nil {
		return nil, err
	}
	defer reader.Close()

	content, err := ioutil.ReadAll(reader)
	if err != nil {
		return nil, err
	}

	return content, nil
}

func getSigsYamlFromCommit(repo *git.Repository, commitFromAnnualReportYear, commitFromCurrentYear plumbing.Hash) (Context, Context, error) {
	var annualReportYearSigs, currentYearSigs Context

	annualReportYearSigFile, err := getFileFromCommit(repo, &commitFromAnnualReportYear, sigsYamlFile)
	if err != nil {
		return Context{}, Context{}, err
	}
	err = yaml.Unmarshal([]byte(annualReportYearSigFile), &annualReportYearSigs)
	if err != nil {
		return Context{}, Context{}, err
	}

	currentYearSigFile, err := getFileFromCommit(repo, &commitFromCurrentYear, sigsYamlFile)
	if err != nil {
		return Context{}, Context{}, err
	}
	err = yaml.Unmarshal([]byte(currentYearSigFile), &currentYearSigs)
	if err != nil {
		return Context{}, Context{}, err
	}

	return annualReportYearSigs, currentYearSigs, nil
}

func contains(strlist []string, val string) bool {
	for _, str := range strlist {
		if str == val {
			return true
		}
	}
	return false
}

func getCategorizedSubprojects(dir string) (map[string][]string, error) {
	subprojectsMap := make(map[string][]string)
	// set for the subprojects in the annual year
	annualSubprojects := make(map[string]bool)

	annualReportYearSigs, currentYearSigs, err := getSigsYamlFromCommit(repo, *commitFromAnnualReportYear, *commitFromCurrentYear)
	if err != nil {
		return nil, err
	}

	// iterate over sigs from the annual report year (say 2022)
	for _, sig1 := range annualReportYearSigs.Sigs {
		if sig1.Dir != dir {
			continue
		}
		for _, sub1 := range sig1.Subprojects {
			annualSubprojects[sub1.Name] = true
		}
	}

	// iterate over sigs from the current year (say 2023)
	for _, sig2 := range currentYearSigs.Sigs {
		if sig2.Dir != dir {
			continue
		}
		for _, sub2 := range sig2.Subprojects {
			if annualSubprojects[sub2.Name] {
				subprojectsMap["Continuing"] = append(subprojectsMap["Continuing"], sub2.Name)
				delete(annualSubprojects, sub2.Name)
			} else {
				subprojectsMap["New"] = append(subprojectsMap["New"], sub2.Name)
			}
		}
	}

	for sub := range annualSubprojects {
		subprojectsMap["Retired"] = append(subprojectsMap["Retired"], sub)
	}

	return subprojectsMap, nil
}

func getCategorizedWorkingGroups(dir string) (map[string][]string, error) {
	workingGroupsMap := make(map[string][]string)

	// set for the working groups in the annual year
	annualWGs := make(map[string]bool)
	annualReportYearSigs, currentYearSigs, err := getSigsYamlFromCommit(repo, *commitFromAnnualReportYear, *commitFromCurrentYear)
	if err != nil {
		return nil, err
	}

	annualReportYearSigs.Complete()
	annualReportYearSigs.Sort()
	currentYearSigs.Complete()
	currentYearSigs.Sort()

	// iterate over the ReportingWGs from the annual report year (say 2022)
	for _, sig := range annualReportYearSigs.Sigs {
		if sig.Dir != dir {
			continue
		}
		for _, wg := range sig.ReportingWGs {
			annualWGs[string(wg)] = true
		}
	}

	// iterate over the ReportingWGs from the current year (say 2023)
	for _, sig := range currentYearSigs.Sigs {
		if sig.Dir != dir {
			continue
		}
		for _, newWG := range sig.ReportingWGs {
			if _, ok := annualWGs[string(newWG)]; !ok {
				workingGroupsMap["New"] = append(workingGroupsMap["New"], string(newWG))
			} else {
				workingGroupsMap["Continuing"] = append(workingGroupsMap["Continuing"], string(newWG))
				delete(annualWGs, string(newWG))
			}
		}
	}

	for wg := range annualWGs {
		workingGroupsMap["Retired"] = append(workingGroupsMap["Retired"], string(wg))
	}

	return workingGroupsMap, nil
}

// prep for automated listing of subprojects in the annual report
func prepForAnnualReportGeneration() error {
	intLastYear, err := strconv.Atoi(lastYear())
	if err != nil {
		return err
	}
	annualReportYear = time.Date(intLastYear, time.January, 1, 0, 0, 0, 0, time.UTC)
	currentYear = time.Date(intLastYear+1, time.January, 1, 0, 0, 0, 0, time.UTC)

	repo, err = git.PlainOpen(localRepoPath)
	if err != nil {
		if err == git.ErrRepositoryNotExists {
			repo, err = git.Clone(memory.NewStorage(), nil, &git.CloneOptions{
				URL: communityRepoURL,
			})
		} else {
			return err
		}
	}

	commitFromAnnualReportYear, err = getCommitByDate(repo, annualReportYear)
	if err != nil {
		return err
	}

	commitFromCurrentYear, err = getCommitByDate(repo, currentYear)
	if err != nil {
		return err
	}

	// fetch KEPs and cache them in the keps variable
	err = fetchKEPs()
	if err != nil {
		return err
	}

	releases, err = getLastThreeK8sReleases()
	if err != nil {
		return err
	}

	return nil
}

func main() {
	yamlPath := filepath.Join(baseGeneratorDir, sigsYamlFile)
	var ctx Context

	err := readYaml(yamlPath, &ctx)
	if err != nil {
		log.Fatal(err)
	}

	ctx.Complete()

	ctx.Sort()

	fmt.Printf("Validating %s\n", yamlPath)
	errs := ctx.Validate()
	if len(errs) != 0 {
		for _, err := range errs {
			fmt.Fprintf(os.Stderr, "ERROR: %s\n", err.Error())
		}
		os.Exit(1)
	}

	// Write the Context struct back to yaml to enforce formatting
	err = writeYaml(&ctx, yamlPath)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Generating group READMEs")
	for prefix, groups := range ctx.PrefixToGroupMap() {
		err = createGroupReadme(groups, prefix)
		if err != nil {
			log.Fatal(err)
		}
	}

	if envVal, ok := os.LookupEnv("ANNUAL_REPORT"); ok && envVal == "true" {
		if err := prepForAnnualReportGeneration(); err != nil {
			log.Fatal(err)
		}
		fmt.Println("Generating annual reports")
		for prefix, groups := range ctx.PrefixToGroupMap() {
			err = createAnnualReportIssue(groups, prefix)
			if err != nil {
				log.Fatal(err)
			}
			err = createAnnualReport(groups, prefix)
			if err != nil {
				log.Fatal(err)
			}
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

	fmt.Println("Generating liaisons.md")
	outputPath = filepath.Join(baseGeneratorDir, liaisonsFilename)
	err = writeTemplate(filepath.Join(baseGeneratorDir, templateDir, liaisonsTemplate), outputPath, "markdown", ctx)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Finished generation!")
}
