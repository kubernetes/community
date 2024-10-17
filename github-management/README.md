# GitHub Management

The Kubernetes project uses Github extensively to store and organize code,
manage issues and documentation, and provide a consistent contributor flow.

With the size and growth of the Kubernetes project, management of our GitHub
footprint has historically been a challenge. We have created a number of
policies to reduce friction and ease administration of our GitHub repositories
and organizations. We have also created a number of tools to automate setup and
enforcement of these policies.

These polices are overseen by the
[GitHub Management subproject](subproject-responsibilities.md) of the Contributor
Experience Special Interest Group.

## Guides
- [Opening a request for assistance with GitHub](opening-a-request.md)
- [Organization Owners Guide](org-owners-guide.md)
- [Repository Creation Guidelines](kubernetes-repositories.md)
- [Setting up the CNCF CLA Check](setting-up-cla-check.md)
- [GitHub Permissions](permissions.md)

## GitHub Administration Team

In order to manage the various organizations that the Kubernetes project owns,
we have a GitHub Administration team that is responsible for carrying out the
various tasks.

This team (**[@kubernetes/owners](https://github.com/orgs/kubernetes/teams/owners)**) is as follows:
* Bob Killen (**[@mrbobbytables](https://github.com/mrbobbytables)**, US Eastern)
* Christoph Blecker (**[@cblecker](https://github.com/cblecker)**, CA Pacific)
* Madhav Jivrajani (**[@MadhavJivrajani](https://github.com/MadhavJivrajani)**, Indian Standard Time)
* Nabarun Pal (**[@palnabarun](https://github.com/palnabarun)**, Indian Standard Time)
* Nikhita Raghunath (**[@nikhita](https://github.com/nikhita)**, Indian Standard Time)
* Priyanka Saggu (**[@Priyankasaggu11929](https://github.com/Priyankasaggu11929)**, Indian Standard Time)

This team is responsible for holding Org Owner privileges over all the active
Kubernetes orgs, and will take action in accordance with our polices and
procedures. All members of this team are subject to the Kubernetes
[security embargo policy].

Nominations to this team will come from the Contributor Experience SIG, and
require confirmation by the Steering Committee before taking effect. Time zones
and country of origin should be considered when selecting membership, to ensure
sufficient after North American business hours and holiday coverage.

### Other roles

#### New Membership Coordinator

New Membership Coordinators help serve as a friendly face to newer, prospective
community members, guiding them through the
[process](new-membership-procedure.md) to request membership to a Kubernetes
GitHub organization.

They also have approval privileges for adding new members to the GitHub config.

Our current coordinators are:
* Mario Jason Braganza (**[jasonbraganza](https://github.com/jasonbraganza)**, Indian Standard Time)

## Project Owned Organizations

The following organizations are currently known to be part of the Kubernetes
project

### Actively used GitHub Organizations

| Name | Description |
| :--: | :---------: |
| [kubernetes](https://github.com/kubernetes) | Core |
| [kubernetes-client](https://github.com/kubernetes-client) | API Client Libraries |
| [kubernetes-csi](https://github.com/kubernetes-csi) | Container Storage Interface Components |
| [kubernetes-nightly](https://github.com/kubernetes-nightly) | Testing org for [publishing-bot](https://github.com/kubernetes/publishing-bot) tooling |
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
| [kubernetes-incubator](https://github.com/kubernetes-incubator) | Earlier used for legacy incubator projects |
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
- [prow](https://sigs.k8s.io/prow/pkg): Prow is our system for handling
  GitHub events and commands for Kubernetes. It is comprised of a number of
  modules/plugins. A couple key ones for GitHub management are below, but a full
  list of commands is available [here](https://go.k8s.io/bot-commands)
  - [branchprotector](https://sigs.k8s.io/prow/cmd/branchprotector):
    enforce branch protection settings across an organization
  - [peribolos](https://sigs.k8s.io/prow/cmd/peribolos): Manage Github
    organization and team membership based on a defined YAML configuration
- [label_sync](https://git.k8s.io/test-infra/label_sync): Add, modify, delete,
  and migrate labels across an entire organization based on a defined YAML
  configuration

[security embargo policy]: https://git.k8s.io/security/private-distributors-list.md#embargo-policy
