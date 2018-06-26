# GitHub Management

The Kubernetes project uses Github extensively to store and organize code,
manage issues and documentation, and provide a consistent contributor flow.

With the size and growth of the Kubernetes project, management of our Github
footprint has historically been a challenge. We have created a number of
policies to reduce friction and ease administration of our Github repositories
and organizations. We have also created a number of tools to automate setup and
enforcement of these policies.

## Guides
- [Organization Owners Guide](org-owners-guide.md)
- [Repository Creation Guidelines](kubernetes-repositories.md)
- [Setting up the CNCF CLA Check](setting-up-cla-check.md)

## Project Owned Organizations

The following organizations are currently known to be part of the Kubernetes
project

### Actively used GitHub Organizations

| Name | Description |
| :--: | :---------: |
| [kubernetes](https://github.com/kubernetes) | Core |
| [kubernetes-client](https://github.com/kubernetes-client) | API Client Libraries |
| [kubernetes-csi](https://github.com/kubernetes-csi) | Container Storage Interface Components |
| [kubernetes-incubator](https://github.com/kubernetes-incubator) | Legacy Incubator Projects |
| [kubernetes-retired](https://github.com/kubernetes-retired) | Retired/Archived Projects |
| [kubernetes-security](https://github.com/kubernetes-security) | Private Security Fix Mirror |
| [kubernetes-sigs](https://github.com/kubernetes-sigs) | SIG-related Projects |

### Non-actively used GitHub Organizations

| Name | Description |
| :--: | :---------: |
| [kubernetes-addons](https://github.com/kubernetes-addons) |  |
| [kubernetes-charts](https://github.com/kubernetes-charts) |  |
| [kubernetes-extensions](https://github.com/kubernetes-extensions) |  |
| [kubernetes-federation](https://github.com/kubernetes-federation) |  |
| [kubernetes-graveyard](https://github.com/kubernetes-graveyard) | kubernetes-retired should be used instead going forward |
| [kubernetes-incubator-retired](https://github.com/kubernetes-incubator-retired) | kubernetes-retired should be used instead going forward |
| [kubernetes-providers](https://github.com/kubernetes-providers) |  |
| [kubernetes-sidecars](https://github.com/kubernetes-sidecars) |  |
| [kubernetes-sig-testing](https://github.com/kubernetes-sig-testing) |  |
| [kubernetes-test](https://github.com/kubernetes-test) |  |
| [kubernetes-tools](https://github.com/kubernetes-tools) |  |

Note, this list is subject to change.

There are more organization names that we are squatting on with possible future
intentions. [For more details please see community issue #1407](https://github.com/kubernetes/community/issues/1407).

## Tooling

We have created a number of tools to help with the management of or Github
repositories and organizations:
- [prow](https://git.k8s.io/test-infra/prow): Prow is our system for handling
  GitHub events and commands for Kubernetes. It is comprised of a number of
  modules/plugins. A couple key ones for GitHub management are below, but a full
  list of commands is available [here](https://go.k8s.io/bot-commands)
  - [branchprotector](https://git.k8s.io/test-infra/prow/cmd/branchprotector):
    enforce branch protection settings across an organization
  - [peribolos](https://git.k8s.io/test-infra/prow/cmd/peribolos): Manage Github
    organization and team membership based on a defined YAML configuration
- [label_sync](https://git.k8s.io/test-infra/label_sync): Add, modify, delete,
  and migrate labels across an entire organization based on a defined YAML
  configuration
