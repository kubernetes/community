---
kep-number: 31
title: Data Driven Commands for Kubectl
authors:
  - "@pwittrock"
owning-sig: sig-cli
participating-sigs:
reviewers:
  - "@soltysh"
  - "@juanvallejo"
  - "@seans3 "
approvers:
  - "@soltysh"
editor: TBD
creation-date: 2018-11-13
last-updated: 2018-11-13
status: provisional
see-also:
replaces:
superseded-by:
---

# data driven commands

## Table of Contents

* [Table of Contents](#table-of-contents)
* [Summary](#summary)
* [Motivation](#motivation)
    * [Goals](#goals)
    * [Non-Goals](#non-goals)
* [Proposal](#proposal)
    * [Implementation Details](#implementation-details)
    * [Risks and Mitigations](#risks-and-mitigations)
* [Graduation Criteria](#graduation-criteria)
* [Alternatives](#alternatives)

## Summary

Many Kubectl commands make requests to specific Resource endpoints.  The request bodies are populated by flags
provided by the user.

Examples:

- `create <resource>`
- `set <field> <resource>`
- `logs`

Although these commands are compiled into the kubectl binary, their workflow is similar to a form on a webpage and
could be complete driven by the server providing the client with the request (endpoint + body) and a set of flags
to populate the request body.

Publishing commands as data from the server addresses cli integration with API extensions as well as client-server
version skew.

## Motivation

Kubectl provides a number of commands to simplify working with Kubernetes by making requests to
Resources and SubResources.  These requests are mostly static, with fields filled in by user
supplied flags.  Today the commands are compiled into the client, which as the following challenges:

- Extension APIs cannot be compiled into the client
- Version-Skewed clients (old client) may be missing commands for new APIs or send outdated requests
- Version-Skewed clients (new client) may have commands for APIs that are not present in the server or expose
  fields not present in older API versions

### Goals

Allow client commands that make a single request to a specific resource and output the result to be data driven
from the server.

- Address cli support for extension APIs
- Address user experience for version skewed clients

### Non-Goals

Allow client commands that have complex client-side logic to be data driven.

- Require a TTY
- Are Agnostic to Specific Resources

## Proposal

Define a format for publishing simple cli commands as data.  CLI commands would be limited to:

- Sending one or more requests to Resource or SubResource Endpoints
- Populating requests from command line flags and arguments
- Writing output populated from the Responses

**Proof of Concept:** [cnctl](https://github.com/pwittrock/kubectl/tree/cnctl/cmd/cnctl)

Instructions to run PoC:

- `go run ./main.go` (no commands show up)
- `kubectl apply` the `cli_v1alpha1_clitestresource.yaml` (add the CRD with the commands)
- `go run ./main.go` (create command shows up)
- `go run ./main create deployment -h` (view create command help)
- `go run ./main create deploy --image nginx --name nginx` (create a deployment)
- `kubectl get deployments`

### Implementation Details

**Publishing Data:**

Alpha: No apimachinery changes required

- Alpha: publish extension Resource Commands as an annotation on CRDs.
- Alpha: publish core Resource Commands as openapi extension.

Beta: apimachinery changes required

- Beta: publish extension Resource Commands a part of the CRD Spec.
- Beta: publish core Resource Commands from new endpoint (like *swagger.json*)

**Data Command Structure:**

```go
/*
Copyright 2018 The Kubernetes Authors.

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

package v1alpha1

// ResourceCommand defines a command that is dynamically defined as an annotation on a CRD
type ResourceCommand struct {
	// Command is the cli Command
	Command Command `json:"command"`

	// Requests are the requests the command will send to the apiserver.
	// +optional
	Requests []ResourceRequest `json:"requests,omitempty"`

	// Output is a go-template used write the command output.  It may reference values specified as flags using
	// {{index .Flags.Strings "flag-name"}}, {{index .Flags.Ints "flag-name"}}, {{index .Flags.Bools "flag-name"}},
	// {{index .Flags.Floats "flag-name"}}.
	//
	// It may also reference values from the responses that were saved using saveResponseValues
	// - {{index .Responses.Strings "response-value-name"}}.
	//
	// Example:
	// 		deployment.apps/{{index .Responses.Strings "responsename"}} created
	//
	// +optional
	Output string `json:"output,omitempty"`
}

type ResourceOperation string

const (
	CREATE_RESOURCE ResourceOperation = "Create"
	UPDATE_RESOURCE                   = "Update"
	DELETE_RESOURCE                   = "Delete"
	GET_RESOURCE                      = "Get"
	PATCH_RESOURCE                    = "Patch"
)

type ResourceRequest struct {
	// Group is the API group of the request endpoint
	//
	// Example: apps
	Group string `json:"group"`

	// Version is the API version of the request endpoint
	//
	// Example: v1
	Version string `json:"version"`

	// Resource is the API resource of the request endpoint
	//
	// Example: deployments
	Resource string `json:"resource"`

	// Operation is the type of operation to perform for the request.  One of: Create, Update, Delete, Get, Patch
	Operation ResourceOperation `json:"operation"`

	// BodyTemplate is a go-template for the request Body.  It may reference values specified as flags using
	// {{index .Flags.Strings "flag-name"}}, {{index .Flags.Ints "flag-name"}}, {{index .Flags.Bools "flag-name"}},
	// {{index .Flags.Floats "flag-name"}}
	//
	// Example:
	//      apiVersion: apps/v1
	//      kind: Deployment
	//      metadata:
	//        name: {{index .Flags.Strings "name"}}
	//        namespace: {{index .Flags.Strings "namespace"}}
	//        labels:
	//          app: nginx
	//      spec:
	//        replicas: {{index .Flags.Ints "replicas"}}
	//        selector:
	//          matchLabels:
	//            app: {{index .Flags.Strings "name"}}
	//        template:
	//          metadata:
	//            labels:
	//              app: {{index .Flags.Strings "name"}}
	//          spec:
	//            containers:
	//            - name: {{index .Flags.Strings "name"}}
	//              image: {{index .Flags.Strings "image"}}
	//
	// +optional
	BodyTemplate string `json:"bodyTemplate,omitempty"`

	// SaveResponseValues are values read from the response and saved in {{index .Responses.Strings "flag-name"}}.
	// They may be used in the ResourceCommand.Output go-template.
	//
	// Example:
	//		- name: responsename
	//        jsonPath: "{.metadata.name}"
	//
	// +optional
	SaveResponseValues []ResponseValue `json:"saveResponseValues,omitempty"`
}

// Flag defines a cli flag that should be registered and available in request / output templates
type Flag struct {
	Type FlagType `json:"type"`

	Name string `json:"name"`

	Description string `json:"description"`

	// +optional
	StringValue string `json:"stringValue,omitempty"`

	// +optional
	StringSliceValue []string `json:"stringSliceValue,omitempty"`

	// +optional
	BoolValue bool `json:"boolValue,omitempty"`

	// +optional
	IntValue int32 `json:"intValue,omitempty"`

	// +optional
	FloatValue float64 `json:"floatValue,omitempty"`
}

// ResponseValue defines a value that should be parsed from a response and available in output templates
type ResponseValue struct {
	Name     string `json:"name"`
	JsonPath string `json:"jsonPath"`
}

type FlagType string

const (
	STRING       FlagType = "String"
	BOOL                  = "Bool"
	FLOAT                 = "Float"
	INT                   = "Int"
	STRING_SLICE          = "StringSlice"
)

type Command struct {
	// Use is the one-line usage message.
	Use string `json:"use"`

	// Path is the path to the sub-command.  Omit if the command is directly under the root command.
	// +optional
	Path []string `json:"path,omitempty"`

	// Short is the short description shown in the 'help' output.
	// +optional
	Short string `json:"short,omitempty"`

	// Long is the long message shown in the 'help <this-command>' output.
	// +optional
	Long string `json:"long,omitempty"`

	// Example is examples of how to use the command.
	// +optional
	Example string `json:"example,omitempty"`

	// Deprecated defines, if this command is deprecated and should print this string when used.
	// +optional
	Deprecated string `json:"deprecated,omitempty"`

	// Flags are the command line flags.
	//
	// Example:
	// 		  - name: namespace
	//    		type: String
	//    		stringValue: "default"
	//    		description: "deployment namespace"
	//
	// +optional
	Flags []Flag `json:"flags,omitempty"`

	// SuggestFor is an array of command names for which this command will be suggested -
	// similar to aliases but only suggests.
	SuggestFor []string `json:"suggestFor,omitempty"`

	// Aliases is an array of aliases that can be used instead of the first word in Use.
	Aliases []string `json:"aliases,omitempty"`

	// Version defines the version for this command. If this value is non-empty and the command does not
	// define a "version" flag, a "version" boolean flag will be added to the command and, if specified,
	// will print content of the "Version" variable.
	// +optional
	Version string `json:"version,omitempty"`
}

// ResourceCommandList contains a list of Commands
type ResourceCommandList struct {
	Items []ResourceCommand `json:"items"`
}
```

**Example Command:**

```go
# Set Label: "cli.sigs.k8s.io/cli.v1alpha1.CommandList": ""
# Set Annotation: "cli.sigs.k8s.io/cli.v1alpha1.CommandList": '<json>'
---
items:
- command:
    path:
    - "create" # Command is a subcommand of this path
    use: "deployment" # Command use
    aliases: # Command alias'
    - "deploy"
    - "deployments"
    short: Create a deployment with the specified name.
    long: Create a deployment with the specified name.
    example: |
        # Create a new deployment named my-dep that runs the busybox image.
        kubectl create deployment --name my-dep --image=busybox
    flags:
    - name: name
      type: String
      stringValue: ""
      description: deployment name
    - name: image
      type: String
      stringValue: ""
      description: Image name to run.
    - name: replicas
      type: Int
      intValue: 1
      description: Image name to run.
    - name: namespace
      type: String
      stringValue: "default"
      description: deployment namespace
  requests:
  - group: apps
    version: v1
    resource: deployments
    operation: Create
    bodyTemplate: |
      apiVersion: apps/v1
      kind: Deployment
      metadata:
        name: {{index .Flags.Strings "name"}}
        namespace: {{index .Flags.Strings "namespace"}}
        labels:
          app: nginx
      spec:
        replicas: {{index .Flags.Ints "replicas"}}
        selector:
          matchLabels:
            app: {{index .Flags.Strings "name"}}
        template:
          metadata:
            labels:
              app: {{index .Flags.Strings "name"}}
          spec:
            containers:
            - name: {{index .Flags.Strings "name"}}
              image: {{index .Flags.Strings "image"}}
    saveResponseValues:
    - name: responsename
      jsonPath: "{.metadata.name}"
  output: |
    deployment.apps/{{index .Responses.Strings "responsename"}} created
```

### Risks and Mitigations

- Command name collisions: CRD publishes command that conflicts with another command
  - Initially require the resource name to be the command name (e.g. `create foo`, `set image foo`)
- Approach is hard to maintain, complex, etc
  - Initially restrict to only `create` commands, get feedback
  

## Graduation Criteria

- Simple commands for Core Resources have been migrated to be data driven
- In use by high profile extension APIs - e.g. Istio
- Published as first class item for Extension and Core Resources

## Alternatives

- Use plugins for these cases
  - Still suffer from version skew
  - Require users to download and install binaries
  - Hard to keep in sync with set of Resources for each cluster
- Don't support cli commands for Extension Resources