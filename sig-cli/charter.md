# SIG CLI Charter

This charter adheres to the conventions described in the [Kubernetes Charter README] and uses
the Roles and Organization Management outlined in [sig-governance].

## Scope

The Command Line Interface SIG (SIG CLI) is responsible for kubectl and
related tools. This group focuses on general purpose command line tools and
libraries to interface with Kubernetes API's.

### In scope

#### Transparently Exposing APIs through Declarative and Imperative Workflows

The scope of CLI Tools focuses on enabling declarative and imperative workflows
for invoking kubernetes APIs and authoring Resource Config.  Tools provide
commands for both generalized (e.g. create resource from Resource Config) tasks and
specialized (e.g. drain a node, exec into a container) tasks.

It is the philosophy of the tools developed in SIG CLI to facilitate working
directly with the Kubernetes APIs and Kubernetes style Resources, and to the
extend possible, provide a transparent experience for how commands map to
Kubernetes APIs and Resources.

Building tools that obfuscate the underlying Resources and APIs (e.g. through
generalized templating or DSLs) is an anti-goal of SIG CLI.  Notable examples
of commands that violate this principle: `run`, `expose`, `autoscale`

#### Specific Types of Functionality

SIG CLI develops tools within the Kubernetes project including but not limited
to the following:

- Invoking Kubernetes APIs:
  - Resource APIs - e.g. `create`, `replace`, `delete`, `patch`, `get`, etc
  - SubResource APIs - e.g. `exec`, `attach`, `logs`, `scale`, etc
  - Discovery Service - e.g. `api-versions`, `api-resources
  - Version - e.g. `version`
  - OpenAPI - e.g. `explain`
- Pre and Post processing API Resource Config and API invocation responses
  - `apply` - creates a patch from Resource Config by performing a 3-way diff
  - `get` - formats API responses into columns
- Aggregating multiple API invocations and post processing the results
  - `describe` - performs multiple calls and aggregates results
  - `run --expose` - creates multiple resources
- Collapsing multiple manual steps into a command
  - `edit` - collapses `get`, edit a file, `patch` into a single command
  - `cp` - collapses manual steps for copy a file out of a container
- Generating Kubernetes Resource Config locally or creating Resources remotely
  - From commands + arguments + flags - e.g. `create configmap`, `run`, `expose`
  - From declarative files - Kustomization `configMapGenerator`, `secretMapGenerator`
- Transforming Kubernetes Resource Config locally or patching remotely
  - From commands + arguments + flags - e.g. `annotate`, `set image`, `patch`
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
  - e.g. `--contect`
- Selecting which API group/version to invoke if ambiguous in the context of the command
  - e.g. `get`, `describe`
  
Notably, CLI Tooling may perform operations (generation, transformation, validation)
that are augmented by analyzing collections of Resource Config - e.g. updating
things that reference names of resources if transforming the names of resources.

#### Go Libraries for developing CLIs

SIG CLI develops go libraries for developing CLI tools for working with Kubernetes.
These libraries provide a subset of the libraries used to build the CLI tools
themselves.

#### Enabling Extensibility in Tooling

CLI prefers to develop commands in such a way that they can provide a native
experience for APIs developed as extensions.  This requires a philosophy of
minimizing resource specific functionality and enabling it through data
published by the cluster rather than hard-coding the API data into the tools.
This is a design preference, not a mandate, and should not come at the practical
cost of impacting the UX or functionality of the tool.  SIG CLI owns developing
extension mechanisms in its own tooling.

CLI prefers to develop commands in such a way that enables tools and solutions
developed independently (e.g. outside the SIG, K8S project, etc) to interoperate
with the CLI tools - e.g. through pipes or wrapping / execing.  This is aligned
with the goal of remaining close to the Kubernetes APIs.

#### Code, Binaries and Services

SIG CLI code include general purpose command line tools and binaries for working
with Kubernetes API's. Examples of these binaries include: [kubectl and kustomize].

### Out of scope

- SIG CLI is not responsible for command-line tools built and maintained by other
  SIGs, such as kubeadm, (which is owned by SIG Cluster Lifecycle).
- SIG CLI is not responsible for tools or solutions developed outside of the Kubernetes
  project.
- SIG CLI is not responsible for commands developed as plugins or other extension
  mechanisms outside of SIG CLI.
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
[kubectl and kustomize]: https://github.com/kubernetes/community/blob/master/sig-cli/README.md#subprojects
[Test Playbook]: https://docs.google.com/document/d/1Z3teqtOLvjAtE-eo0G9tjyZbgNc6bMhYGZmOx76v6oM

