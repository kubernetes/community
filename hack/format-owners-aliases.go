/*
Copyright 2020 The Kubernetes Authors.

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
	"bytes"
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"sort"
	"strings"

	"github.com/pkg/errors"
	"github.com/pmezard/go-difflib/difflib"
	yaml "gopkg.in/yaml.v3"
)

const (
	aliasesPrefix = "aliases:\n"
	customStart   = "## BEGIN CUSTOM CONTENT"
	customEnd     = "## END CUSTOM CONTENT"
)

type flags struct {
	diff   bool
	write  bool
	custom bool
}

type nodes []yaml.Node

func (n nodes) Len() int {
	return len(n)
}

func (n nodes) Swap(i, j int) {
	n[i], n[j] = n[j], n[i]
}

func (n nodes) Less(i, j int) bool {
	return n[i].Value < n[j].Value
}

type aliases struct {
	Aliases map[string]nodes
}

func (a *aliases) marshalYAML() []byte {
	// add the keys to a string slice
	keys := make([]string, len(a.Aliases))
	i := 0
	for k := range a.Aliases {
		keys[i] = k
		// remove dupes and sort the list of OWNERS
		a.Aliases[k] = mergeDupes(a.Aliases[k])
		sort.Sort(a.Aliases[k])
		i++
	}
	sort.Strings(keys)

	// write YAML
	const indent = "  "
	b := bytes.Buffer{}
	b.WriteString(aliasesPrefix)

	for _, k := range keys {
		b.WriteString(indent + k + ":\n")
		for _, v := range a.Aliases[k] {
			line := indent + indent + "- " + v.Value
			// there seems to be a bug in yaml.v3 where a LineComment
			// from a parent Node ends up on a child Node,
			// thus comments on group lines (e.g. "some-sig: # comment")
			// is not supported.
			if len(v.LineComment) > 0 {
				line += " " + v.LineComment
			}
			line += "\n"
			b.WriteString(line)
		}
	}
	return b.Bytes()
}

func mergeDupes(s nodes) nodes {
	seen := make(map[string]struct{}, len(s))
	i := 0
	for _, v := range s {
		val := v.Value
		if _, ok := seen[val]; ok {
			continue
		}
		seen[val] = struct{}{}
		s[i] = v
		i++
	}
	return s[:i]
}

func printErrorAndExit(err error) {
	fmt.Println(err)
	os.Exit(1)
}

func newDecoderFromBytes(b []byte) *yaml.Decoder {
	d := yaml.NewDecoder(bytes.NewReader(b))
	d.KnownFields(true)
	return d
}

func processData(original []byte, f *flags) ([]byte, error) {
	var custom []byte
	regular := original[:]

	// look for custom content if requested
	if f.custom {
		customErr := errors.Errorf("custom content must be located at the end of the file "+
			"enclosed between \"%s\" and \"%s\\n\"", customStart, customEnd)
		str := string(original)
		idxStart := strings.Index(str, customStart)
		idxEnd := strings.Index(str, customEnd)
		if idxStart == -1 {
			return nil, customErr
		}
		if idxEnd != len(str)-len(customEnd)-1 {
			return nil, customErr
		}
		regular = regular[:idxStart]
		custom = original[idxStart : idxEnd+len(customEnd)+1]
		// prepend 'aliases:\n' to the custom content
		custom = append([]byte(aliasesPrefix), custom...)
	}

	// unmarshal the YAMLs
	aRegular := aliases{}
	d := newDecoderFromBytes(regular)
	if err := d.Decode(&aRegular); err != nil {
		return nil, err
	}
	sorted := aRegular.marshalYAML()

	if len(custom) > 0 {
		aCustom := aliases{}
		d := newDecoderFromBytes(custom)
		if err := d.Decode(&aCustom); err != nil {
			return nil, err
		}
		sortedCustom := aCustom.marshalYAML()

		// trim the aliases prefix from the custom YAML
		sortedCustom = bytes.TrimPrefix(sortedCustom, []byte(aliasesPrefix))

		// concatenate the regular and custom content
		sorted = append(sorted, []byte(customStart)...)
		sorted = append(sorted, []byte("\n")...)
		sorted = append(sorted, sortedCustom...)
		sorted = append(sorted, []byte(customEnd)...)
		sorted = append(sorted, []byte("\n")...)
	}

	return sorted, nil
}

func processFile(filePath string, f *flags) error {
	info, err := os.Stat(filePath)
	if err != nil {
		return err
	}

	original, err := ioutil.ReadFile(filePath)
	if err != nil {
		return err
	}

	sorted, err := processData(original, f)
	if err != nil {
		return err
	}

	if bytes.Compare(original, sorted) != 0 {
		// if writing in place was requested, write and return
		if f.write {
			if err := ioutil.WriteFile(filePath, sorted, info.Mode()); err != nil {
				return err
			}
			return nil
		}

		// check if a diff is requested
		if f.diff {
			diff := difflib.UnifiedDiff{
				A:        difflib.SplitLines(string(original)),
				B:        difflib.SplitLines(string(sorted)),
				FromFile: "Original",
				ToFile:   "Formatted",
				Context:  3,
			}
			text, err := difflib.GetUnifiedDiffString(diff)
			if err != nil {
				return errors.Wrapf(err, "could not obtain a unified diff for file %q", filePath)
			}
			fmt.Println(text)
		} else {
			fmt.Println(string(sorted))
		}
		return errors.Errorf("the file %q is not formatted", filePath)
	}
	return nil
}

func main() {
	if len(os.Args) < 2 {
		printErrorAndExit(errors.New("Usage: format-owners-aliases <flags> <OWNERS_ALIASES_FILE>"))
	}

	f := flags{}
	flag.BoolVar(&f.custom, "c", false, "look for custom content in the file")
	flag.BoolVar(&f.diff, "d", false, "write a diff instead of the whole formatted output")
	flag.BoolVar(&f.write, "w", false, "write result to (source) file instead of stdout")
	flag.Parse()

	// treat the first non-flag as the file path to process
	var filePath string
	for _, f := range os.Args[1:] {
		if !strings.HasPrefix(f, "-") {
			filePath = f
			break
		}
	}

	if err := processFile(filePath, &f); err != nil {
		printErrorAndExit(err)
	}
}
