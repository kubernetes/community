---
kep-number: 8
title: Kustomize
authors:
  - "@pwittrock"
  - "@monopole"
owning-sig: sig-cli
participating-sigs:
  - sig-cli
reviewers:
  - "@droot"
approvers:
  - "@maciej"
editor: "@droot"
creation-date: 2018-05-5
last-updated: 2018-05-5
status: implemented
see-also:
  - n/a
replaces:
  - kinflate # Old name for kustomize
superseded-by:
  - n/a
---

# Kustomize

## Table of Contents

- [Kustomize](#kustomize)
   - [Table of Contents](#table-of-contents)
   - [Summary](#summary)
   - [Motivation](#motivation)
      - [Goals](#goals)
      - [Non-Goals](#non-goals)
   - [Proposal](#proposal)
      - [Implementation Details/Notes/Constraints [optional]](#implementation-detailsnotesconstraints-optional)
      - [Risks and Mitigations](#risks-and-mitigations)
      - [Risks of Not Having a Solution](#risks-of-not-having-a-solution)
   - [Graduation Criteria](#graduation-criteria)
   - [Implementation History](#implementation-history)
   - [Drawbacks](#drawbacks)
   - [Alternatives](#alternatives)
   - [FAQ](#faq)

## Summary

Declarative specification of Kubernetes objects is the recommended way to manage Kubernetes
production workloads, however gaps in the kubectl tooling force users to write their own scripting and
tooling to augment the declarative tools with preprocessing transformations.
While most of theser transformations already exist as imperative kubectl commands, they are not natively accessible
from a declarative workflow.
 
This KEP describes how `kustomize` addresses this problem by providing a declarative format for users to access
the imperative kubectl commands they are already familiar natively from declarative workflows.

## Motivation

The kubectl command provides a cli for:

- accessing the Kubernetes apis through json or yaml configuration
- porcelain commands for generating and transforming configuration off of commandline flags.

Examples:

- Generate a configmap or secret from a text or binary file
  - `kubectl create configmap`, `kubectl create secret`
  - Users can manage their configmaps and secrets text and binary files

- Create or update fields that cut across other fields and objects
  - `kubectl label`, `kubectl annotate`
  - Users can add and update labels for all objects composing an application
  
- Transform an existing declarative configuration without forking it
  - `kubectl patch`
  - Users may generate multiple variations of the same workload

- Transform live resources arbitrarily without auditing
  - `kubectl edit`

To create a Secret from a binary file, users must first base64 encode the binary file and then create a Secret yaml
config from the resulting data.  Because the source of truth is actually the binary file, not the config,
users must write scripting and tooling to keep the 2 sources consistent.

Instead, users should be able to access the simple, but necessary, functionality available in the imperative
kubectl commands from their declarative workflow.

#### Long standing issues

Kustomize addresses a number of long standing issues in kubectl.

- Declarative enumeration of multiple files [kubernetes/kubernetes#24649]
- Declarative configmap and secret creation: [kubernetes/kubernetes#24744], [kubernetes/kubernetes#30337]
- Configmap rollouts: [kubernetes/kubernetes#22368]
  - [Example in kustomize](https://github.com/kubernetes/kubectl/blob/master/cmd/kustomize/demos/helloWorld.md#how-this-works-with-kustomize)
- Name/label scoping and safer pruning: [kubernetes/kubernetes#1698]
  - [Example in kustomize](https://github.com/kubernetes/kubectl/blob/master/cmd/kustomize/demos/breakfast.md#demo-configure-breakfast)
- Template-free add-on customization: [kubernetes/kubernetes#23233]
  - [Example in kustomize](https://github.com/kubernetes/kubectl/blob/master/cmd/kustomize/demos/helloWorld.md#staging-kustomization)

### Goals

- Declarative support for defining ConfigMaps and Secrets generated from binary and text files
- Declarative support for adding or updating cross-cutting fields
  - labels & selectors
  - annotations
  - names (as transformation of the original name)
- Declarative support for applying patches to transform arbitrary fields
  - use strategic-merge-patch format
- Ease of integration with CICD systems that maintain configuration in a version control repository
  as a single source of truth, and take action (build, test, deploy, etc.) when that truth changes (gitops).

### Non-Goals

#### Exposing every imperative kubectl command in a declarative fashion

The scope of kustomize is limited only to functionality gaps that would otherwise prevent users from
defining their workloads in a purely declarative manner (e.g. without writing scripts to perform pre-processing
or linting).  Commands such as `kubectl run`, `kubectl create deployment` and `kubectl edit` are unnecessary
in a declarative workflow because a Deployment can easily be managed as declarative config.

#### Providing a simpler facade on top of the Kubernetes APIs

The community has developed a number of facades in front of the Kubernetes APIs using
templates or DSLs.  Attempting to provide an alternative interface to the Kubernetes API is
a non-goal.  Instead the focus is on:

- Facilitating simple cross-cutting transformations on the raw config that would otherwise require other tooling such
  as *sed*
- Generating configuration when the source of truth resides elsewhere
- Patching existing configuration with transformations

## Proposal

### Capabilities

**Note:** This proposal has already been implemented in `github.com/kubernetes/kubectl`.

Define a new meta config format called *kustomization.yaml*.

#### *kustomization.yaml* will allow users to reference config files

- Path to config yaml file (similar to `kubectl apply -f <file>`)
- Urls to config yaml file (similar to `kubectl apply -f <url>`)
- Path to *kustomization.yaml* file (takes the output of running kustomize)

#### *kustomization.yaml* will allow users to generate configs from files

- ConfigMap (`kubectl create configmap`)
- Secret (`kubectl create secret`)

#### *kustomization.yaml* will allow users to apply transformations to configs

- Label (`kubectl label`)
- Annotate (`kubectl annotate`)
- Strategic-Merge-Patch (`kubectl patch`)
- Name-Prefix

### UX

Kustomize will also contain subcommands to facilitate authoring *kustomization.yaml*.

#### Edit

The edit subcommands will allow users to modify the *kustomization.yaml* through cli commands containing
helpful messaging and documentation.

- Add ConfigMap - like `kubectl create configmap` but declarative in *kustomization.yaml*
- Add Secret - like `kubectl create secret` but declarative in *kustomization.yaml*
- Add Resource - adds a file reference to *kustomization.yaml*
- Set NamePrefix - adds NamePrefix declaration to *kustomization.yaml*

#### Diff

The diff subcommand will allow users to see a diff of the original and transformed configuration files

- Generated config (configmap) will show the files as created
- Transformations (name prefix) will show the files as modified

### Implementation Details/Notes/Constraints [optional]

Kustomize has already been implemented in the `github.com/kubernetes/kubectl` repo, and should be moved to a
separate repo for the subproject.

Kustomize was initially developed as its own cli, however once it has matured, it should be published
as a subcommand of kubectl or as a statically linked plugin.  It should also be more tightly integrated with apply.

- Create the *kustomize* sig-cli subproject and update sigs.yaml
- Move the existing kustomize code from `github.com/kubernetes/kubectl` to `github.com/kubernetes-sigs/kustomize`

### Risks and Mitigations


### Risks of Not Having a Solution

By not providing a viable option for working directly with Kubernetes APIs as json or
yaml config, we risk the ecosystem becoming fragmented with various bespoke API facades.
By ensuring the raw Kubernetes API json or yaml is a usable approach for declaratively
managing applications, even tools that do not use the Kubernetes API as their native format can
better work with one another through transformation to a common format.

## Graduation Criteria

- Dogfood kustomize by either:
  - moving one or more of our own (OSS Kubernetes) services to it.
  - getting user feedback from one or more mid or large application deployments using kustomize.
- Publish kustomize as a subcommand of kubectl.

## Implementation History

kustomize was implemented in the kubectl repo before subprojects became a first class thing in Kubernetes.
The code has been fully implemented, but it must be moved to a proper location.

## Drawbacks


## Alternatives

1. Users write their own bespoke scripts to generate and transform the config before it is applied.
2. Users don't work with the API directly, and use or develop DSLs for interacting with Kubernetes.

## FAQs
