/*
Copyright 2025 The Kubernetes Authors.

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
	"os"
	"path/filepath"
	"regexp"
	"strings"
	"time"

	yaml "gopkg.in/yaml.v3"
)

const (
	maxWordCount         = 450
	recommendedWordCount = 300
	startYear            = 2026
)

var yamlHeaderRegex = regexp.MustCompile(`(?s)^-{61}\s*\n(.*?)\n-{61}\s*\n`)

type ValidationError struct {
	File    string
	Message string
}

type CandidateHeader struct {
	Name string   `yaml:"name"`
	ID   string   `yaml:"ID"`
	Info InfoData `yaml:"info"`
}

type ElectionConfig struct {
	ShowCandidateFields []string `yaml:"show_candidate_fields"`
}

// TODO: Have Elekto accept both mapping and sequence forms for candidate info
// so compatibility is enforced at the application boundary as well as here.
type InfoData map[string]string

func (i *InfoData) UnmarshalYAML(value *yaml.Node) error {
	switch value.Kind {
	case yaml.MappingNode:
		var info map[string]string
		if err := value.Decode(&info); err != nil {
			return err
		}
		*i = InfoData(info)
		return nil
	case yaml.SequenceNode:
		var entries []map[string]string
		if err := value.Decode(&entries); err != nil {
			return err
		}
		info := make(InfoData)
		for _, entry := range entries {
			for field, fieldValue := range entry {
				info[field] = fieldValue
			}
		}
		*i = info
		return nil
	default:
		return fmt.Errorf("info must be a mapping or sequence")
	}
}

func main() {
	if len(os.Args) < 2 {
		fmt.Fprintf(os.Stderr, "Usage: %s <elections-path>\n", os.Args[0])
		os.Exit(1)
	}

	electionsPath := os.Args[1]

	bioFiles, err := findBioFiles(electionsPath)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error finding bio files: %v\n", err)
		os.Exit(1)
	}

	var errors []ValidationError

	for _, bioFile := range bioFiles {
		// Check word count
		wordCount, err := countWords(bioFile)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error counting words in %s: %v\n", bioFile, err)
			continue
		}

		if wordCount > maxWordCount {
			errors = append(errors, ValidationError{
				File:    bioFile,
				Message: fmt.Sprintf("has %d words", wordCount),
			})
		}

		// Check filename format and GitHub ID matching
		if err := validateFileNameAndGitHubID(bioFile); err != nil {
			errors = append(errors, ValidationError{
				File:    bioFile,
				Message: err.Error(),
			})
		}

		candidateFields, err := loadShowCandidateFields(bioFile)
		if err != nil {
			errors = append(errors, ValidationError{
				File:    bioFile,
				Message: err.Error(),
			})
			continue
		}

		// Check template compliance
		if err := validateTemplateCompliance(bioFile, candidateFields); err != nil {
			errors = append(errors, ValidationError{
				File:    bioFile,
				Message: err.Error(),
			})
		}
	}

	if len(errors) > 0 {
		for _, err := range errors {
			fmt.Printf("%s: %s\n", err.File, err.Message)
		}

		separator := strings.Repeat("=", 68)
		fmt.Printf("\n%s\n", separator)
		fmt.Printf("%d invalid Steering Committee election bio(s) detected.\n", len(errors))
		fmt.Printf("Bios should be limited to around %d words, excluding headers.\n", recommendedWordCount)
		fmt.Printf("Bios must follow the nomination template and filename format.\n")
		fmt.Printf("%s\n", separator)
		os.Exit(1)
	}
}

func loadShowCandidateFields(filename string) ([]string, error) {
	electionPath := filepath.Join(filepath.Dir(filename), "election.yaml")
	content, err := os.ReadFile(electionPath)
	if err != nil {
		return nil, fmt.Errorf("error reading election configuration: %v", err)
	}

	var config ElectionConfig
	if err := yaml.Unmarshal(content, &config); err != nil {
		return nil, fmt.Errorf("error parsing election configuration: %v", err)
	}

	return config.ShowCandidateFields, nil
}

// findBioFiles finds all steering committee candidate bio files from 2025 forward
func findBioFiles(electionsPath string) ([]string, error) {
	var bioFiles []string

	// Walk through elections/steering directory specifically
	steeringPath := filepath.Join(electionsPath, "steering")
	err := filepath.Walk(steeringPath, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		// Skip if not a regular file
		if !info.Mode().IsRegular() {
			return nil
		}

		// Check if it's a candidate bio file
		if !strings.HasPrefix(info.Name(), "candidate-") || !strings.HasSuffix(info.Name(), ".md") {
			return nil
		}

		// Only include elections from startYear forward
		pathLower := strings.ToLower(path)
		includeFile := false
		currentYear := time.Now().Year()
		// Check from startYear to a few years ahead of current year
		endYear := currentYear + 5
		for year := startYear; year <= endYear; year++ {
			yearStr := fmt.Sprintf("/%d/", year)
			if strings.Contains(pathLower, yearStr) {
				includeFile = true
				break
			}
		}

		if !includeFile {
			return nil
		}

		bioFiles = append(bioFiles, path)
		return nil
	})

	return bioFiles, err
}

// countWords counts the number of words in a file
func countWords(filename string) (int, error) {
	content, err := os.ReadFile(filename)
	if err != nil {
		return 0, err
	}

	content = yamlHeaderRegex.ReplaceAll(content, nil)

	// Split by whitespace and count non-empty strings
	words := regexp.MustCompile(`\s+`).Split(string(content), -1)
	count := 0
	for _, word := range words {
		if strings.TrimSpace(word) != "" {
			count++
		}
	}

	return count, nil
}

// validateFileNameAndGitHubID checks if filename matches format candidate-$username.md
// and if the username matches the GitHub ID in the document header
func validateFileNameAndGitHubID(filename string) error {
	// Extract filename from path
	base := filepath.Base(filename)

	// Check filename format: candidate-*.md
	candidateRegex := regexp.MustCompile(`^candidate-([a-zA-Z0-9_-]+)\.md$`)
	matches := candidateRegex.FindStringSubmatch(base)
	if len(matches) != 2 {
		return fmt.Errorf("filename must follow format 'candidate-username.md'")
	}

	expectedUsername := matches[1]

	// Read file content to extract GitHub ID
	content, err := os.ReadFile(filename)
	if err != nil {
		return fmt.Errorf("error reading file: %v", err)
	}

	// Extract YAML header and parse it
	yamlHeader, err := extractYAMLHeader(string(content))
	if err != nil {
		return fmt.Errorf("error extracting YAML header: %v", err)
	}

	// Parse the YAML to get the ID field
	var header map[string]interface{}
	if err := yaml.Unmarshal([]byte(yamlHeader), &header); err != nil {
		return fmt.Errorf("error parsing YAML header: %v", err)
	}

	idValue, exists := header["ID"]
	if !exists {
		return fmt.Errorf("missing 'ID' field in header")
	}

	actualUsername, ok := idValue.(string)
	if !ok {
		return fmt.Errorf("'ID' field must be a string")
	}

	actualUsername = strings.TrimSpace(actualUsername)
	if actualUsername != expectedUsername {
		return fmt.Errorf("filename username '%s' does not match GitHub ID '%s' in header", expectedUsername, actualUsername)
	}

	return nil
}

// validateTemplateCompliance checks if the bio follows the required template structure
func validateTemplateCompliance(filename string, candidateFields []string) error {
	content, err := os.ReadFile(filename)
	if err != nil {
		return fmt.Errorf("error reading file: %v", err)
	}

	contentStr := string(content)

	// Extract YAML header between dashes
	yamlHeader, err := extractYAMLHeader(contentStr)
	if err != nil {
		return fmt.Errorf("error extracting YAML header: %v", err)
	}

	// Parse YAML header
	var header CandidateHeader
	if err := yaml.Unmarshal([]byte(yamlHeader), &header); err != nil {
		return fmt.Errorf("invalid YAML header format: %v", err)
	}

	// Validate required fields
	if header.Name == "" {
		return fmt.Errorf("missing required field: name")
	}
	if header.ID == "" {
		return fmt.Errorf("missing required field: ID")
	}
	for _, field := range candidateFields {
		if header.Info[field] == "" {
			return fmt.Errorf("missing required field: info.%s", field)
		}
	}

	// Check for required sections
	requiredSections := map[string][]string{
		"SIGs":               {"## SIGS", "## SIGs"},
		"What I have done":   {"## What I have done"},
		"What I'll do":       {"## What I'll do"},
		"Resources About Me": {"## Resources About Me"},
	}

	for sectionName, alternatives := range requiredSections {
		found := false
		for _, section := range alternatives {
			if strings.Contains(contentStr, section) {
				found = true
				break
			}
		}
		if !found {
			return fmt.Errorf("missing required section: %s", sectionName)
		}
	}

	return nil
}

// extractYAMLHeader extracts the YAML content between dash separators
// The format requires exactly 61 dashes for consistency
func extractYAMLHeader(content string) (string, error) {
	// Find the YAML header between exactly 61 dashes
	matches := yamlHeaderRegex.FindStringSubmatch(content)
	if len(matches) != 2 {
		return "", fmt.Errorf("could not find YAML header between dashes")
	}
	return matches[1], nil
}
