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

package main

import (
	"flag"
	"fmt"
	"os"

	"k8s.io/community/generator/sigslackleads"
)

func main() {
	var root string
	var generate bool
	var dryRun bool
	flag.StringVar(&root, "root", ".", "repository root")
	flag.BoolVar(&generate, "generate", false, "update generated usergroups from sigs.yaml")
	flag.BoolVar(&dryRun, "dry-run", false, "validate generation without writing files")
	flag.Parse()
	if flag.NArg() != 0 {
		fail("unexpected positional arguments")
	}
	if dryRun && !generate {
		fail("-dry-run requires -generate")
	}

	if generate {
		if err := sigslackleads.Generate(root, dryRun); err != nil {
			fail("SIG Slack lead generation failed: %v", err)
		}
		_, _ = fmt.Fprintln(os.Stdout, "SIG Slack lead usergroups generated from sigs.yaml")
		return
	}
	if err := sigslackleads.Verify(root); err != nil {
		fail("SIG Slack lead verification failed: %v", err)
	}
	_, _ = fmt.Fprintln(os.Stdout, "SIG Slack lead usergroups match sigs.yaml")
}

func fail(format string, args ...any) {
	_, _ = fmt.Fprintf(os.Stderr, format+"\n", args...)
	os.Exit(1)
}
