# SIG CLI Charter

This charter adheres to the conventions described in the [Kubernetes Charter README] and uses
the Roles and Organization Management outlined in [sig-governance].

## Scope

The Command Line Interface SIG (SIG CLI) is responsible for kubectl and
related tools. This group focuses on command line tools and
libraries to interface with Kubernetes API's.

### In scope

SIG CLI section in [sigs.yaml]

#### Code, Binaries and Services

SIG CLI code include command line tools and binaries for working with Kubernetes
API's. Examples of these binaries include: kubectl and kustomize.

### Out of scope

SIG CLI is NOT responsible for defining the Kubernetes API that it
interfaces with. The Kubernetes API is the responsibility of SIG API Machinery.

## Roles and Organization Management

SIG CLI adheres to the Roles and Organization Management outlined in [sig-governance]
and opts-in to updates and modifications to [sig-governance].

### Deviations from [sig-governance]

- In addition to Technical Leads, SIG CLI defines Emeritus Leads. These former
  SIG CLI leaders *SHOULD* be available to provide historical perspective and
  knowledge.
- SIG CLI defines the role of Test Health Maintainer. Contributors who have
  successfully completed one test on-call rotation within the last six months as
  shown in the test on-call schedule of the [Test Playbook] are included in this
  group. Test Health Maintainers are SIG CLI Members.

### Subproject Creation

Option 1: by [SIG Technical Leads](https://github.com/kubernetes/community/blob/master/committee-steering/governance/sig-governance.md#L100)


[sig-governance]: https://github.com/kubernetes/community/blob/master/committee-steering/governance/sig-governance.md
[sigs.yaml]: https://github.com/kubernetes/community/blob/master/sigs.yaml#L487
[Kubernetes Charter README]: https://github.com/kubernetes/community/blob/master/committee-steering/governance/README.md
[Test Playbook]: https://docs.google.com/document/d/1Z3teqtOLvjAtE-eo0G9tjyZbgNc6bMhYGZmOx76v6oM

