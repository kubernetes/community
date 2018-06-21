# Kubernetes Github Organization Guide

The Kubernetes project leverages multiple GitHub organizations to store and
organize code. This guide contains the details on how to run those organizations
for CNCF compliance and for the guidelines of the community.

## Organization Naming

Kubernetes managed organizations should be in the form of `kubernetes-[thing]`.
For example, [kubernetes-client](https://github.com/kubernetes-client) where the
API clients are housed.

Prior to creating an organization please contact the steering committee for
direction and approval.

Note: The CNCF, as part of the Linux Foundation, holds the trademark on the
Kubernetes name. All GitHub organizations with Kubernetes in the name should be
managed by the Kubernetes project or use a different name.

## Transferring Outside Code Into A Kubernetes Organization

Due to licensing and CLA issues, prior to transferring software into a Kubernetes
managed organization there is some due diligence that needs to occur. Please
contact the steering committee and CNCF prior to moving any code in.

It is easier to start new code in a Kubernetes organization than it is to
transfer in existing code.

## Current Organizations In Use

The following organizations are currently known to be part of the Kubernetes
project:

### Actively used GitHub Organizations:
* [kubernetes](https://github.com/kubernetes)
* [kubernetes-client](https://github.com/kubernetes-client)
* [kubernetes-csi](https://github.com/kubernetes-csi)
* [kubernetes-incubator](https://github.com/kubernetes-incubator)
* [kubernetes-retired](https://github.com/kubernetes-retired)
* [kubernetes-security](https://github.com/kubernetes-security)
* [kubernetes-sig-testing](https://github.com/kubernetes-sig-testing)
* [kubernetes-sigs](https://github.com/kubernetes-sigs)

### Non-actively used GitHub Organizations:
* [kubernetes-addons](https://github.com/kubernetes-addons)
* [kubernetes-charts](https://github.com/kubernetes-charts)
* [kubernetes-extensions](https://github.com/kubernetes-extensions)
* [kubernetes-federation](https://github.com/kubernetes-federation)
* [kubernetes-graveyard](https://github.com/kubernetes-graveyard) †
* [kubernetes-incubator-retired](https://github.com/kubernetes-incubator-retired)
* [kubernetes-providers](https://github.com/kubernetes-providers)
* [kubernetes-sidecars](https://github.com/kubernetes-sidecars)
* [kubernetes-test](https://github.com/kubernetes-test)
* [kubernetes-tools](https://github.com/kubernetes-tools)

† kubernetes-retired should be used instead of kubernetes-graveyard going forward.

Note, this list is subject to change.

There are more organization names that we are squatting on with possible future
intentions. [For more details please see community issue #1407](https://github.com/kubernetes/community/issues/1407).

## Team Guidance

Each organization should have the following teams:

- teams for each repo `foo`
  - `foo-admins`: granted admin access to the `foo` repo
  - `foo-maintainers`: granted write access to the `foo` repo
  - `foo-reviewers`: granted read access to the `foo` repo; intended to be used as
    a notification mechanism for interested/active contributors for the `foo` repo
- a `bots` team
  - should contain bots such as @k8s-ci-robot and @thelinuxfoundation that are
    necessary for org and repo automation
- an `owners` team
  - should be populated by everyone who has `owner` privileges to the org
  - gives users the opportunity to ping owners as a group rather than having to
    search for individuals

**NB**: Not all organizations in use today currently follow this team guidance.
We are looking to coalesce existing teams towards this model, and use this model
for all orgs going forward.  Notable discrepancies at the moment:

- `foo-reviewers` teams are considered a historical subset of
  `kubernetes-sig-foo-pr-reviews` teams and are intended mostly as a fallback
  notification mechanism when requested reviewers are being unresponsive.  Ideally
  OWNERS files can be used in lieu of these teams.
- `admins-foo` and `maintainers-foo` teams as used by the kubernetes-incubator
  org. This was a mistake that swapped the usual convention, and we would like
  to rename the team

## Repository Guidance

Repositories have additional guidelines and requirements, such as the use of
CLA checking on all contributions. For more details on those please see the
[Kubernetes Template Project](https://github.com/kubernetes/kubernetes-template-project).
