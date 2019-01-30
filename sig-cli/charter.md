# SIG CLI Charter

This charter adheres to the conventions described in the [Kubernetes Charter README] and uses
the Roles and Organization Management outlined in [sig-governance].

## Scope

### In scope

The Command Line Interface SIG (SIG CLI) is responsible for kubectl and
related tools that are developed by the SIG.  See kubectl
[design principles](design-principles.md) for the focus of kubectl functionality.

This group focuses on command line tooling for working with Kubernetes APIs and Resource
Config.  This includes both generalized tooling for working with Resources, Resource Config
and Resource Types (e.g. using resource / object metadata, duck-typing, openapi, discovery, 
scale and status subresources, etc) as well as tooling for working with specific Kubernetes
APIs (e.g. `logs`, `exec`, `create configmap`).

The scope includes both low level tooling that may be used by scripts or higher-level
solutions, as well as higher level porcelain to reduce user friction for simple,
common, difficult or important tasks performed by users (e.g. `edit`, `top`, `wait`).  The
scope also includes publishing a subset the libraries which were used to develop the tooling
itself.

The decision whether to publish specific functionality as part of kubectl,
as a separate tool, as a kubectl extension, or as a library are technical
decisions made by the SIG and owners of the code under development.

#### Examples of Functionality In Kubectl

Following are examples of functionality that is in kubectl.

- Invoking Kubernetes APIs:
  - Resource APIs - e.g. `create`, `replace`, `delete`, `patch`, `get`, etc
  - SubResource APIs - e.g. `exec`, `attach`, `logs`, `scale`, etc
  - Discovery Service - e.g. `api-versions`, `api-resources`
  - Version - e.g. `version`
  - OpenAPI - e.g. `explain`
- Pre and Post processing API Resource Config, API Requests and API Responses
  - `apply` - creates a patch from Resource Config by performing a 3-way diff
  - `get` - formats API responses into columns
- Aggregating multiple API Responses and post processing them
  - `describe` - performs multiple calls and aggregates results
  - `run --expose` - creates multiple resources
- Collapsing multiple manual steps into a command
  - `edit` - collapses `get`, edit a file, `patch` into a single command
  - `cp` - collapses manual steps for copy a file out of a container
- Generating Kubernetes Resource Config locally or creating Resources remotely
  - From commands + arguments + flags - e.g. `create configmap`, `run`, `expose`
  - From declarative files - Kustomization `configMapGenerator`, `secretMapGenerator`
- Transforming Kubernetes Resource Config locally or patching remotely
  - From commands + arguments + flags - e.g. `annotate`, `set image`, `patch`, etc
  - From declarative files - e.g. Kustomization `commonAnnotations`, `images`, `patches`
  - From flags - e.g. `-n` - set the namespace of Resources
- Blocking on propagation of an event or change to the cluster
  - e.g. `wait`, `rollout`
- Defining or referencing a collection of Resource Config
  - e.g. `-f` can reference a url, a file containing multiple Resources, or a directory
  - e.g. `-R` can traverse a directory recursively
  - e.g. Kustomization `bases` can refer to other Kustomizations
  - e.g. Kustomization `resources` can refer to Resource Config files
- Configure how to talk to a specific cluster from the cli
  - e.g. `config`
  - e.g. `--context`
- Selecting which API group/version to invoke if ambiguous in the context of the command
  - e.g. `get`, `describe`
  
Notably, CLI Tooling also performs operations (generation, transformation, validation)
that work against collections of Resources or Resource Config - e.g. updating
things that reference names of resources if transforming the names of resources.

### Out of scope

- SIG CLI is not responsible for tools developed outside of the
  SIG (even if they are part of the broader Kubernetes project).
- SIG CLI is not responsible for kubectl subcommands developed outside of the
  SIG (even if they are developed through kubectl extension mechanisms).
- SIG CLI is not responsible for defining the Kubernetes API or Resource Types
  that it interfaces with (which is owned by SIG Apps and SIG API Machinery).
- SIG CLI is not responsible for integrating with *specific tools* or APIs
  developed outside of the Kubernetes project.

## Roles and Organization Management

SIG CLI adheres to the Roles and Organization Management outlined in [sig-governance]
and opts-in to updates and modifications to [sig-governance].

### Deviations from [sig-governance]

- In addition to Technical Leads, SIG CLI defines Emeritus Leads. These former
  SIG CLI leaders *SHOULD* be available to provide historical perspective and
  domain knowledge.
- SIG CLI defines the role of Test Health Maintainer. Contributors who have
  successfully completed one test on-call rotation within the last six months as
  shown in the test on-call schedule of the [Test Playbook] are included in this
  group. Test Health Maintainers are SIG CLI Members.

### Subproject Creation

Option 1: by [SIG Technical Leads](https://github.com/kubernetes/community/blob/master/committee-steering/governance/sig-governance.md#tech-lead)


[Kubernetes Charter README]: https://github.com/kubernetes/community/blob/master/committee-steering/governance/README.md
[sig-governance]: https://github.com/kubernetes/community/blob/master/committee-steering/governance/sig-governance.md
[README]: https://github.com/kubernetes/community/blob/master/sig-cli/README.md
[Test Playbook]: https://docs.google.com/document/d/1Z3teqtOLvjAtE-eo0G9tjyZbgNc6bMhYGZmOx76v6oM

