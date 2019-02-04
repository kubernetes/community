# SIG CLI Charter

This charter adheres to the conventions described in the [Kubernetes Charter README] and uses
the Roles and Organization Management outlined in [sig-governance].

## Scope

### In scope

The Command Line Interface SIG (SIG CLI) is responsible for kubectl, as well as accompanying
libraries and documentation.  See kubectl [design principles](design-principles.md) for the
focus of kubectl functionality.

**Note:** Definition of kubectl may include commands developed by SIG CLI as kubectl plugins.

Kubectl is a dynamic + resource-oriented CLI, and reference implementation for interacting
with the API, as well as a basic tool for declarative and imperative management.

SIG CLI focuses on command line tooling for working with Kubernetes APIs and Resource Config.
This includes both generalized tooling for working with Resources, Resource Config
and Resource Types (e.g. using resource / object metadata, duck-typing, openapi, discovery, 
scale and status subresources, etc), as well as tooling for working with specific Kubernetes
APIs (e.g. `logs`, `exec`, `create configmap`).

The scope includes both low level tooling that may be used by things like scripts,
as well as higher level porcelain to reduce user friction for simple,
common, difficult or important tasks performed by users.  The
scope also includes publishing a subset the libraries which were used to develop the tooling
itself.

The decision whether to publish specific functionality as part of kubectl,
as a separate tool, as a kubectl extension, or as a library are technical
decisions made by the SIG and owners of the code under development.

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

