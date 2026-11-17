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
	"os"
	"path/filepath"
	"slices"
	"strings"
	"testing"
)

func TestExpectedGroup(t *testing.T) {
	group := sig{Dir: "sig-example", Name: "Example"}
	group.Contact.Slack = "sig-example"
	group.Leadership.Chairs = []person{{GitHub: "Zulu"}, {GitHub: "alpha"}}
	group.Leadership.TechnicalLeads = []person{{GitHub: "alpha"}, {GitHub: "Beta"}}
	users := map[string]string{"alpha": "U000000001", "Beta": "U000000002", "Zulu": "U000000003"}
	channels := map[string]bool{"sig-example": false}

	name, members, err := expectedGroup(group, users, userKeyIndex(users), channels)
	if err != nil {
		t.Fatal(err)
	}
	if name != "sig-example-leads" {
		t.Fatalf("name = %q", name)
	}
	want := []string{"alpha", "Beta", "Zulu"}
	if !slices.Equal(members, want) {
		t.Fatalf("members = %q, want %q", members, want)
	}
}

func TestExpectedGroupUsesSIGSlackFields(t *testing.T) {
	group := sig{Dir: "sig-windows", Name: "Windows"}
	group.Contact.Slack = "sig-windows"
	group.Contact.SlackUsergroup = "sig-windows-leads-group"
	group.Contact.SlackUsergroupChannels = []string{"windows-alerts", "sig-windows"}
	group.Leadership.Chairs = []person{{GitHub: "lead"}}
	users := map[string]string{"lead": "U000000001"}
	channels := map[string]bool{
		"sig-windows":             false,
		"sig-windows-leads":       false,
		"sig-windows-leads-group": false,
		"windows-alerts":          false,
	}

	_, _, err := expectedGroup(group, users, userKeyIndex(users), channels)
	if err == nil || !strings.Contains(err.Error(), "collides with a channel") {
		t.Fatalf("err = %v", err)
	}
	delete(channels, "sig-windows-leads-group")
	name, _, err := expectedGroup(group, users, userKeyIndex(users), channels)
	if err != nil {
		t.Fatal(err)
	}
	if name != "sig-windows-leads-group" {
		t.Fatalf("name = %q", name)
	}
	wantChannels := []string{"windows-alerts", "sig-windows"}
	if got := requiredChannels(group); !slices.Equal(got, wantChannels) {
		t.Fatalf("required channels = %q, want %q", got, wantChannels)
	}
}

func TestExpectedGroupRejectsSharedSlackIDForDistinctActiveLeaders(t *testing.T) {
	group := sig{Dir: "sig-example", Name: "Example"}
	group.Contact.Slack = "sig-example"
	group.Leadership.Chairs = []person{{GitHub: "one"}, {GitHub: "two"}}
	users := map[string]string{"one": "U000000001", "two": "U000000001"}
	_, _, err := expectedGroup(group, users, userKeyIndex(users), map[string]bool{"sig-example": false})
	if err == nil || !strings.Contains(err.Error(), "share Slack user ID") {
		t.Fatalf("err = %v", err)
	}
}

func TestVerifyGroups(t *testing.T) {
	group := sig{Dir: "sig-example", Name: "Example"}
	group.Contact.Slack = "sig-example"
	group.Contact.SlackUsergroupDescription = "Leads for SIG Example."
	group.Contact.SlackUsergroupChannels = []string{"chairs-and-techleads"}
	group.Leadership.Chairs = []person{{GitHub: "lead"}}
	users := map[string]string{"lead": "U000000001", "unrelated": "U000000002"}
	actual := map[string]usergroup{
		"sig-example-leads": {
			Name:        "sig-example-leads",
			LongName:    "SIG Example Leads",
			Description: "Leads for SIG Example.",
			Channels:    []string{"chairs-and-techleads", "sig-example"},
			Members:     []string{"lead"},
		},
	}
	channels := map[string]bool{"sig-example": false, "chairs-and-techleads": false}
	if err := verifyGroups([]sig{group}, users, channels, actual); err != nil {
		t.Fatal(err)
	}

	badLong := actual["sig-example-leads"]
	badLong.LongName = "Old Example Leads"
	actual["sig-example-leads"] = badLong
	if err := verifyGroups([]sig{group}, users, channels, actual); err == nil || !strings.Contains(err.Error(), "expected \"SIG Example Leads\"") {
		t.Fatalf("err = %v", err)
	}

	badChannel := actual["sig-example-leads"]
	badChannel.LongName = "SIG Example Leads"
	badChannel.Channels = []string{"sig-example", "missing"}
	actual["sig-example-leads"] = badChannel
	if err := verifyGroups([]sig{group}, users, channels, actual); err == nil || !strings.Contains(err.Error(), "channel \"missing\" is not configured") {
		t.Fatalf("err = %v", err)
	}

	badMembers := actual["sig-example-leads"]
	badMembers.Channels = []string{"sig-example", "chairs-and-techleads"}
	badMembers.Members = []string{}
	actual["sig-example-leads"] = badMembers
	if err := verifyGroups([]sig{group}, users, channels, actual); err == nil || !strings.Contains(err.Error(), "members are") {
		t.Fatalf("err = %v", err)
	}
}

func TestVerifyGroupsRejectsMissingAndExtraGroups(t *testing.T) {
	group := sig{Dir: "sig-example", Name: "Example"}
	group.Contact.Slack = "sig-example"
	group.Leadership.Chairs = []person{{GitHub: "lead"}}
	users := map[string]string{"lead": "U000000001"}
	channels := map[string]bool{"sig-example": false}
	if err := verifyGroups([]sig{group}, users, channels, nil); err == nil || !strings.Contains(err.Error(), "missing generated usergroup") {
		t.Fatalf("err = %v", err)
	}
	actual := map[string]usergroup{
		"sig-old-leads": {Name: "sig-old-leads", LongName: "Old", Description: "Old", Members: []string{"lead"}},
	}
	if err := verifyGroups([]sig{group}, users, channels, actual); err == nil || !strings.Contains(err.Error(), "does not match an active SIG") {
		t.Fatalf("err = %v", err)
	}
}

func TestLoadManagedGroupsUsesGeneratedFile(t *testing.T) {
	slackDir := writeSlackLayout(t)
	generated := `usergroups:
  - name: sig-example-leads
    long_name: SIG Example Leads
    description: Active SIG Example leads.
    channels:
      - sig-example
    members:
      - lead
`
	if err := os.WriteFile(filepath.Join(slackDir, generatedFileName), []byte(generated), 0o600); err != nil {
		t.Fatal(err)
	}
	groups, err := loadManagedGroups(slackDir)
	if err != nil {
		t.Fatal(err)
	}
	if got := groups["sig-example-leads"]; got.Name != "sig-example-leads" {
		t.Fatalf("groups = %#v", groups)
	}
}

func TestLoadManagedGroupsRejectsDuplicateHandles(t *testing.T) {
	slackDir := writeSlackLayout(t)
	generated := `usergroups:
  - name: duplicate
    long_name: Duplicate
    description: Duplicate
    members: []
`
	if err := os.WriteFile(filepath.Join(slackDir, generatedFileName), []byte(generated), 0o600); err != nil {
		t.Fatal(err)
	}
	manual := `usergroups:
  - name: duplicate
    long_name: Duplicate
    description: Duplicate
    members: []
`
	if err := os.WriteFile(filepath.Join(slackDir, "usergroups.yaml"), []byte(manual), 0o600); err != nil {
		t.Fatal(err)
	}
	if _, err := loadManagedGroups(slackDir); err == nil || !strings.Contains(err.Error(), "declared in both") {
		t.Fatalf("err = %v", err)
	}
}

func TestGenerateGroupsUpdatesGeneratedFileOnly(t *testing.T) {
	root, slackDir := writeGeneratorFixture(t)
	manualPath := filepath.Join(slackDir, "usergroups.yaml")
	manual := []byte(`usergroups:
  - name: unrelated
    long_name: Unrelated
    description: Keep this exact group.
    members:
      - unrelated
`)
	if err := os.WriteFile(manualPath, manual, 0o600); err != nil {
		t.Fatal(err)
	}
	generatedPath := filepath.Join(slackDir, generatedFileName)
	original := []byte(`usergroups:
  - name: sig-example-leads
    long_name: SIG Example Leads
    description: Existing generated group.
    channels:
      - sig-example
    members:
      - stale
      - Zulu # This comment is generated away
`)
	if err := os.WriteFile(generatedPath, original, 0o600); err != nil {
		t.Fatal(err)
	}

	if err := verify(root); err == nil {
		t.Fatal("verify accepted stale generated content")
	}
	if unchanged, err := os.ReadFile(generatedPath); err != nil || !slices.Equal(unchanged, original) {
		t.Fatalf("verify changed generated file: %v", err)
	}
	if err := generateGroups(root, false); err != nil {
		t.Fatal(err)
	}
	if err := verify(root); err != nil {
		t.Fatal(err)
	}
	generated, err := os.ReadFile(generatedPath)
	if err != nil {
		t.Fatal(err)
	}
	text := string(generated)
	for _, want := range []string{"description: Existing generated group.", "- extra-channel", "- alpha\n", "- Zulu\n"} {
		if !strings.Contains(text, want) {
			t.Fatalf("generated file does not contain %q:\n%s", want, text)
		}
	}
	if strings.Contains(text, "- stale") || strings.Contains(text, "generated away") || strings.Contains(text, "# Retained") {
		t.Fatalf("non-generated list data remains:\n%s", text)
	}
	if unchanged, err := os.ReadFile(manualPath); err != nil || !slices.Equal(unchanged, manual) {
		t.Fatalf("generation changed manual file: %v", err)
	}
	if err := os.WriteFile(generatedPath, append(generated, []byte("# manual drift\n")...), 0o600); err != nil {
		t.Fatal(err)
	}
	if err := verify(root); err == nil || !strings.Contains(err.Error(), "is stale") {
		t.Fatalf("verify accepted comment drift: %v", err)
	}
	if err := generateGroups(root, false); err != nil {
		t.Fatal(err)
	}
	if second, err := os.ReadFile(generatedPath); err != nil || !slices.Equal(second, generated) {
		t.Fatalf("second generation was not idempotent: %v", err)
	}
}

func TestGenerateGroupsValidatesBeforeWriting(t *testing.T) {
	root, slackDir := writeGeneratorFixture(t)
	groupsPath := filepath.Join(slackDir, generatedFileName)
	original := []byte(`usergroups:
  - name: sig-example-leads
    long_name: SIG Example Leads
    description: Existing group.
    channels:
      - sig-example
      - extra-channel
    members:
      - old
`)
	if err := os.WriteFile(groupsPath, original, 0o600); err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(filepath.Join(slackDir, "users.yaml"), []byte("users: {}\n"), 0o600); err != nil {
		t.Fatal(err)
	}
	if err := generateGroups(root, false); err == nil {
		t.Fatal("generation accepted a missing Slack identity")
	}
	if actual, err := os.ReadFile(groupsPath); err != nil || !slices.Equal(actual, original) {
		t.Fatalf("generation wrote before validation: %v", err)
	}
}

func TestVerifyGroupsRejectsInvalidSlackIDsAndAmbiguousKeys(t *testing.T) {
	group := sig{Dir: "sig-example", Name: "Example"}
	group.Contact.Slack = "sig-example"
	group.Leadership.Chairs = []person{{GitHub: "lead"}}
	actual := map[string]usergroup{
		"sig-example-leads": {Name: "sig-example-leads", LongName: "SIG Example Leads", Channels: []string{"sig-example"}, Members: []string{"lead"}},
	}
	for _, test := range []struct {
		name  string
		users map[string]string
		want  string
	}{
		{name: "invalid", users: map[string]string{"lead": "not-a-slack-id"}, want: "invalid Slack ID"},
		{name: "case ambiguous", users: map[string]string{"lead": "U000000001", "Lead": "U000000002"}, want: "ambiguous users.yaml keys"},
	} {
		t.Run(test.name, func(t *testing.T) {
			if err := verifyGroups([]sig{group}, test.users, map[string]bool{"sig-example": false}, actual); err == nil || !strings.Contains(err.Error(), test.want) {
				t.Fatalf("err = %v", err)
			}
		})
	}
}

func writeSlackLayout(t *testing.T) string {
	t.Helper()
	slackDir := filepath.Join(t.TempDir(), slackConfigDir)
	if err := os.MkdirAll(slackDir, 0o700); err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(filepath.Join(slackDir, "usergroups.yaml"), []byte("usergroups: []\n"), 0o600); err != nil {
		t.Fatal(err)
	}
	return slackDir
}

func writeGeneratorFixture(t *testing.T) (string, string) {
	t.Helper()
	root := t.TempDir()
	slackDir := filepath.Join(root, slackConfigDir)
	if err := os.MkdirAll(slackDir, 0o700); err != nil {
		t.Fatal(err)
	}
	files := map[string]string{
		filepath.Join(root, "sigs.yaml"): `sigs:
- dir: sig-example
  name: Example
  leadership:
    chairs:
    - github: Zulu
    tech_leads:
    - github: alpha
  contact:
    slack: sig-example
    slack_usergroup_description: Existing generated group.
    slack_usergroup_channels:
    - extra-channel
`,
		filepath.Join(slackDir, "users.yaml"): `users:
  alpha: U000000001
  unrelated: U000000002
  Zulu: U000000003
`,
		filepath.Join(slackDir, "channels.yaml"): `channels:
- name: extra-channel
- name: sig-example
`,
	}
	for path, content := range files {
		if err := os.WriteFile(path, []byte(content), 0o600); err != nil {
			t.Fatal(err)
		}
	}
	return root, slackDir
}

func userKeyIndex(users map[string]string) map[string][]string {
	result := map[string][]string{}
	for login := range users {
		key := strings.ToLower(login)
		result[key] = append(result[key], login)
	}
	return result
}
