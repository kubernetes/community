/*
Copyright 2026 The Kubernetes Authors.

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

package sigslackleads

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"slices"
	"strings"

	"gopkg.in/yaml.v3"
)

const (
	slackConfigDir    = "communication/slack-config"
	generatedFileName = "sigslackleads.yaml"
)

type person struct {
	GitHub string `yaml:"github"`
}

type sig struct {
	Dir        string `yaml:"dir"`
	Name       string `yaml:"name"`
	Leadership struct {
		Chairs         []person `yaml:"chairs"`
		TechnicalLeads []person `yaml:"tech_leads"`
	} `yaml:"leadership"`
	Contact struct {
		Slack                     string   `yaml:"slack"`
		SlackUsergroup            string   `yaml:"slack_usergroup,omitempty"`
		SlackUsergroupDescription string   `yaml:"slack_usergroup_description,omitempty"`
		SlackUsergroupChannels    []string `yaml:"slack_usergroup_channels,omitempty"`
	} `yaml:"contact"`
}

type usergroup struct {
	Name        string   `yaml:"name"`
	LongName    string   `yaml:"long_name"`
	Description string   `yaml:"description"`
	Channels    []string `yaml:"channels,omitempty"`
	Members     []string `yaml:"members"`
	External    bool     `yaml:"external,omitempty"`
}

type slackFile struct {
	Users           map[string]string `yaml:"users,omitempty"`
	Channels        []channel         `yaml:"channels,omitempty"`
	Usergroups      []usergroup       `yaml:"usergroups,omitempty"`
	ChannelTemplate any               `yaml:"channel_template,omitempty"`
	Restrictions    any               `yaml:"restrictions,omitempty"`
}

type channel struct {
	Name     string `yaml:"name"`
	ID       string `yaml:"id,omitempty"`
	Archived bool   `yaml:"archived,omitempty"`
}

// Verify checks that sigslackleads.yaml matches sigs.yaml and the Slack config.
func Verify(root string) error {
	return verify(root)
}

// Generate updates sigslackleads.yaml from sigs.yaml. dryRun validates the
// generated result without replacing the file.
func Generate(root string, dryRun bool) error {
	return generateGroups(root, dryRun)
}

func verify(root string) error {
	sigs, users, channels, managed, err := load(root)
	if err != nil {
		return err
	}
	if err := verifyGroups(sigs, users, channels, managed); err != nil {
		return err
	}
	_, ordered, err := expectedGroups(sigs, users, channels)
	if err != nil {
		return err
	}
	want, err := marshalGenerated(ordered)
	if err != nil {
		return err
	}
	filename := filepath.Join(root, slackConfigDir, generatedFileName)
	got, err := os.ReadFile(filename)
	if err != nil {
		return err
	}
	if !bytes.Equal(got, want) {
		return fmt.Errorf("%s is stale; run make generate", filename)
	}
	return nil
}

func generateGroups(root string, dryRun bool) error {
	sigs, users, channels, _, err := load(root)
	if err != nil {
		return err
	}
	projected, ordered, err := expectedGroups(sigs, users, channels)
	if err != nil {
		return err
	}
	if err := verifyGroups(sigs, users, channels, projected); err != nil {
		return err
	}
	content, err := marshalGenerated(ordered)
	if err != nil {
		return err
	}
	filename := filepath.Join(root, slackConfigDir, generatedFileName)
	if dryRun {
		return nil
	}
	info, err := os.Stat(filename)
	if err != nil {
		return err
	}
	return os.WriteFile(filename, content, info.Mode().Perm())
}

func expectedGroups(sigs []sig, users map[string]string, channels map[string]bool) (map[string]usergroup, []usergroup, error) {
	userKeys := map[string][]string{}
	for login := range users {
		folded := strings.ToLower(login)
		userKeys[folded] = append(userKeys[folded], login)
	}
	projected := make(map[string]usergroup, len(sigs))
	ordered := make([]usergroup, 0, len(sigs))
	for _, groupSIG := range sigs {
		name, members, err := expectedGroup(groupSIG, users, userKeys, channels)
		if err != nil {
			return nil, nil, err
		}
		group := usergroup{
			Name:        name,
			LongName:    "SIG " + groupSIG.Name + " Leads",
			Description: groupDescription(groupSIG),
			Channels:    requiredChannels(groupSIG),
			Members:     members,
		}
		projected[name] = group
		ordered = append(ordered, group)
	}
	return projected, ordered, nil
}

func marshalGenerated(groups []usergroup) ([]byte, error) {
	var content bytes.Buffer
	content.WriteString("# This file is generated from sigs.yaml.\n")
	content.WriteString("# Update leadership or Slack usergroup fields in sigs.yaml, then run make generate.\n\n")
	encoder := yaml.NewEncoder(&content)
	encoder.SetIndent(2)
	if err := encoder.Encode(slackFile{Usergroups: groups}); err != nil {
		return nil, err
	}
	if err := encoder.Close(); err != nil {
		return nil, err
	}
	return content.Bytes(), nil
}

func load(root string) ([]sig, map[string]string, map[string]bool, map[string]usergroup, error) {
	sigs, err := loadSIGs(filepath.Join(root, "sigs.yaml"))
	if err != nil {
		return nil, nil, nil, nil, err
	}
	users, err := loadUsers(filepath.Join(root, slackConfigDir, "users.yaml"))
	if err != nil {
		return nil, nil, nil, nil, err
	}
	channels, err := loadChannels(filepath.Join(root, slackConfigDir))
	if err != nil {
		return nil, nil, nil, nil, err
	}
	managed, err := loadManagedGroups(filepath.Join(root, slackConfigDir))
	if err != nil {
		return nil, nil, nil, nil, err
	}
	return sigs, users, channels, managed, nil
}

func loadSIGs(filename string) ([]sig, error) {
	raw, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	var data struct {
		SIGs []sig `yaml:"sigs"`
	}
	if err := decodeYAMLOpen(raw, &data); err != nil {
		return nil, fmt.Errorf("parse %s: %w", filename, err)
	}
	if len(data.SIGs) == 0 {
		return nil, fmt.Errorf("%s has no SIGs", filename)
	}
	return data.SIGs, nil
}

func loadUsers(filename string) (map[string]string, error) {
	raw, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	var data slackFile
	if err := decodeYAML(raw, &data); err != nil {
		return nil, fmt.Errorf("parse %s: %w", filename, err)
	}
	if len(data.Users) == 0 {
		return nil, fmt.Errorf("%s has no users", filename)
	}
	return data.Users, nil
}

func loadChannels(dir string) (map[string]bool, error) {
	channels := map[string]bool{}
	err := filepath.WalkDir(dir, func(filename string, entry os.DirEntry, walkErr error) error {
		if walkErr != nil {
			return walkErr
		}
		if entry.IsDir() || filepath.Ext(entry.Name()) != ".yaml" {
			return nil
		}
		raw, err := os.ReadFile(filename)
		if err != nil {
			return err
		}
		var data slackFile
		if err := decodeYAML(raw, &data); err != nil {
			return fmt.Errorf("parse %s: %w", filename, err)
		}
		for _, channel := range data.Channels {
			if channel.Name == "" {
				return fmt.Errorf("%s contains a channel without a name", filename)
			}
			if _, exists := channels[channel.Name]; exists {
				return fmt.Errorf("channel %q is declared more than once", channel.Name)
			}
			channels[channel.Name] = channel.Archived
		}
		return nil
	})
	if err != nil {
		return nil, err
	}
	return channels, nil
}

func loadManagedGroups(dir string) (map[string]usergroup, error) {
	generatedPath := filepath.Join(dir, generatedFileName)
	generatedRaw, err := os.ReadFile(generatedPath)
	if err != nil {
		return nil, err
	}
	var generated slackFile
	if err := decodeYAML(generatedRaw, &generated); err != nil {
		return nil, fmt.Errorf("parse %s: %w", generatedPath, err)
	}
	groups := make(map[string]usergroup, len(generated.Usergroups))
	for _, group := range generated.Usergroups {
		if group.Name == "" {
			return nil, fmt.Errorf("%s contains a usergroup without a name", generatedPath)
		}
		if _, exists := groups[group.Name]; exists {
			return nil, fmt.Errorf("generated usergroup %q is declared more than once", group.Name)
		}
		groups[group.Name] = group
	}

	seenHandles := map[string]string{}
	err = filepath.WalkDir(dir, func(filename string, entry os.DirEntry, walkErr error) error {
		if walkErr != nil {
			return walkErr
		}
		if entry.IsDir() || filepath.Ext(entry.Name()) != ".yaml" {
			return nil
		}
		raw, err := os.ReadFile(filename)
		if err != nil {
			return err
		}
		var data slackFile
		if err := decodeYAML(raw, &data); err != nil {
			return fmt.Errorf("parse %s: %w", filename, err)
		}
		rel, err := filepath.Rel(dir, filename)
		if err != nil {
			return err
		}
		for _, group := range data.Usergroups {
			if group.Name == "" {
				return fmt.Errorf("%s contains a usergroup without a name", filename)
			}
			if previous, exists := seenHandles[group.Name]; exists {
				return fmt.Errorf("usergroup %q is declared in both %s and %s", group.Name, previous, filepath.ToSlash(rel))
			}
			seenHandles[group.Name] = filepath.ToSlash(rel)
		}
		return nil
	})
	if err != nil {
		return nil, err
	}
	return groups, nil
}

func verifyGroups(sigs []sig, users map[string]string, channels map[string]bool, actual map[string]usergroup) error {
	var problems []string
	userKeys := map[string][]string{}
	for login, id := range users {
		folded := strings.ToLower(login)
		userKeys[folded] = append(userKeys[folded], login)
		if !isSlackUserID(id) {
			problems = append(problems, fmt.Sprintf("users.yaml entry %q has invalid Slack ID %q", login, id))
		}
	}
	for _, got := range actual {
		if got.External {
			problems = append(problems, fmt.Sprintf("%s: generated SIG group must not be external", got.Name))
		}
	}

	expectedNames := map[string]string{}
	for _, groupSIG := range sigs {
		expectedName, members, err := expectedGroup(groupSIG, users, userKeys, channels)
		if err != nil {
			problems = append(problems, err.Error())
			continue
		}
		if previous, exists := expectedNames[expectedName]; exists && previous != groupSIG.Dir {
			problems = append(problems, fmt.Sprintf("usergroup %q is expected by both %s and %s", expectedName, previous, groupSIG.Dir))
		}
		expectedNames[expectedName] = groupSIG.Dir
		got, ok := actual[expectedName]
		if !ok {
			problems = append(problems, fmt.Sprintf("missing generated usergroup %q for %s", expectedName, groupSIG.Dir))
			continue
		}
		if got.Name != expectedName {
			problems = append(problems, fmt.Sprintf("%s: handle is %q, expected %q", groupSIG.Dir, got.Name, expectedName))
		}
		wantLongName := "SIG " + groupSIG.Name + " Leads"
		if got.LongName != wantLongName {
			problems = append(problems, fmt.Sprintf("%s: long_name is %q, expected %q", groupSIG.Dir, got.LongName, wantLongName))
		}
		if want := groupDescription(groupSIG); got.Description != want {
			problems = append(problems, fmt.Sprintf("%s: description is %q, expected %q", groupSIG.Dir, got.Description, want))
		}
		for _, channel := range got.Channels {
			archived, exists := channels[channel]
			if !exists {
				problems = append(problems, fmt.Sprintf("%s: channel %q is not configured", groupSIG.Dir, channel))
			} else if archived {
				problems = append(problems, fmt.Sprintf("%s: channel %q is archived", groupSIG.Dir, channel))
			}
		}
		if want := requiredChannels(groupSIG); !slices.Equal(got.Channels, want) {
			problems = append(problems, fmt.Sprintf("%s: channels are %q, expected %q", groupSIG.Dir, got.Channels, want))
		}
		if !slices.Equal(got.Members, members) {
			problems = append(problems, fmt.Sprintf("%s: members are %q, expected %q", groupSIG.Dir, got.Members, members))
		}
	}
	for name := range actual {
		if expectedNames[name] == "" {
			problems = append(problems, fmt.Sprintf("generated usergroup %q does not match an active SIG", name))
		}
	}
	if len(problems) != 0 {
		slices.Sort(problems)
		return errors.New(strings.Join(slices.Compact(problems), "\n"))
	}
	return nil
}

func expectedGroup(group sig, users map[string]string, userKeys map[string][]string, channels map[string]bool) (string, []string, error) {
	if group.Dir == "" || group.Name == "" || group.Contact.Slack == "" {
		return "", nil, fmt.Errorf("SIG has incomplete dir/name/contact.slack data")
	}
	if archived, exists := channels[group.Contact.Slack]; !exists {
		return "", nil, fmt.Errorf("%s: contact channel %q is not configured", group.Dir, group.Contact.Slack)
	} else if archived {
		return "", nil, fmt.Errorf("%s: contact channel %q is archived", group.Dir, group.Contact.Slack)
	}
	seenLogins := map[string]bool{}
	seenIDs := map[string]string{}
	members := []string{}
	all := slices.Concat(group.Leadership.Chairs, group.Leadership.TechnicalLeads)
	for _, leader := range all {
		login := strings.TrimSpace(leader.GitHub)
		if login == "" {
			return "", nil, fmt.Errorf("%s: active leader has no GitHub login", group.Dir)
		}
		folded := strings.ToLower(login)
		matches := userKeys[folded]
		if len(matches) == 0 {
			return "", nil, fmt.Errorf("%s: %q is missing from users.yaml", group.Dir, login)
		}
		if len(matches) != 1 {
			return "", nil, fmt.Errorf("%s: %q has ambiguous users.yaml keys %q", group.Dir, login, matches)
		}
		key := matches[0]
		id := users[key]
		if !isSlackUserID(id) {
			return "", nil, fmt.Errorf("%s: %q has invalid Slack user ID %q", group.Dir, key, id)
		}
		if previous, exists := seenIDs[id]; exists && previous != key {
			return "", nil, fmt.Errorf("%s: active leaders %q and %q share Slack user ID %q", group.Dir, previous, key, id)
		}
		seenIDs[id] = key
		if !seenLogins[key] {
			seenLogins[key] = true
			members = append(members, key)
		}
	}
	if len(members) == 0 {
		return "", nil, fmt.Errorf("%s: no active leaders", group.Dir)
	}
	slices.SortFunc(members, func(a, b string) int { return strings.Compare(strings.ToLower(a), strings.ToLower(b)) })
	name := group.Contact.SlackUsergroup
	if name == "" {
		name = group.Dir + "-leads"
	}
	if _, collides := channels[name]; collides {
		return "", nil, fmt.Errorf("%s: usergroup handle %q collides with a channel", group.Dir, name)
	}
	return name, members, nil
}

func groupDescription(group sig) string {
	if group.Contact.SlackUsergroupDescription != "" {
		return group.Contact.SlackUsergroupDescription
	}
	return "Active chairs and technical leads for SIG " + group.Name + "."
}

func requiredChannels(group sig) []string {
	channels := slices.Clone(group.Contact.SlackUsergroupChannels)
	if !slices.Contains(channels, group.Contact.Slack) {
		channels = append(channels, group.Contact.Slack)
	}
	return channels
}

func isSlackUserID(id string) bool {
	if len(id) < 9 || (id[0] != 'U' && id[0] != 'W') {
		return false
	}
	for _, char := range id[1:] {
		if (char < 'A' || char > 'Z') && (char < '0' || char > '9') {
			return false
		}
	}
	return true
}

func decodeYAML(raw []byte, out any) error {
	return decodeYAMLMode(raw, out, true)
}

func decodeYAMLOpen(raw []byte, out any) error {
	return decodeYAMLMode(raw, out, false)
}

func decodeYAMLMode(raw []byte, out any, knownFields bool) error {
	dec := yaml.NewDecoder(bytes.NewReader(raw))
	dec.KnownFields(knownFields)
	if err := dec.Decode(out); err != nil {
		return err
	}
	var trailing any
	if err := dec.Decode(&trailing); !errors.Is(err, io.EOF) {
		return errors.New("multiple YAML documents are not allowed")
	}
	return nil
}
