# Kubernetes GitHub Organization Guide

The Kubernetes project leverages multiple GitHub organizations to store and
organize code. This guide contains the details on how to run those organizations
for CNCF compliance and for the guidelines of the community.

## SLOs

The [GitHub Administration Team] will aim to handle requests in the following
time frames:
- Organization invites should be handled within 72 hours of all requirements for
  membership being met (all +1s obtained).
- Repository creation or migration requests should be responded to within 72
  hours of the issue being opened. There may be information required or specific
  requirements that take additional time, but once all requirements are met, the
  repo should be created within 72 hours.
- Security or moderation requests should be handled ASAP, and coverage should be
  provided in multiple time zones and countries.
- All other requests should be responded to within 72 hours of the issue being
  opened. The time to resolve these requests will vary depending on the
  specifics of the request.

If a request is taking longer than the above time frames, or there is a need to
escalate an urgent request, please mention **[@kubernetes/owners]** on the
associated issue for assistance.

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

Due to licensing and CLA issues, prior to transferring software into a
Kubernetes managed organization there is some due diligence that needs to occur.
Please contact the steering committee and CNCF prior to moving any code in.

It is easier to start new code in a Kubernetes organization than it is to
transfer in existing code.

## Team Guidance

### Nomenclature

Each organization should have the following teams:

- teams for each repo `foo`
  - `foo-admins`: granted admin access to the `foo` repo
  - `foo-maintainers`: granted write access to the `foo` repo
  - `foo-reviewers`: granted read access to the `foo` repo; intended to be used
    as a notification mechanism for interested/active contributors for the `foo`
    repo
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
  notification mechanism when requested reviewers are being unresponsive.
  Ideally OWNERS files can be used in lieu of these teams.
- `admins-foo` and `maintainers-foo` teams as used by the kubernetes-incubator
  org. This was a mistake that swapped the usual convention, and we would like
  to rename the team

### Structure and Process

Guidelines on how to create and structure teams are described below:

#### Structure

**Renaming a team**:

To rename a team, add the `previously: <old-team-name>` field to the team
and rename the `name` of the team to the new name.

**Creating a team**:

- Unless a member is part of the [@kubernetes/owners] team, they needed to be
  added to the `members` list in the team. Members of the
  [@kubernetes/owners] team *must* be added to the `maintainers` list because
  of how the GitHub API works.
- The `privacy` of a team *must* be `closed`.

#### Process

A new team can be created or a member can be added to a team by
creating a PR against the [kubernetes/org] repo. The PR must be
approved by the relevant `OWNERS` or the SIG leads.

For example, addition of a member to `foo-maintainers` must be approved
by the `OWNERS` of the repo `foo` or the leads of the SIG associated
with the repo.

## Project Board Guidance

Guidelines for project boards in the Kubernetes GitHub orgs are described below:

- All project boards should be organization-level project boards instead of
repository-level even if the project board is intended to be scoped to a single
repository. It is easier to distribute permissions via org-level project boards
since write access to a repo-level project board requires full write access
to the repo.

- Project Boards *must* have `Public` visibility.

- The default _Organization Member Permission_ is suggested to be `Write` so
that contributors can move cards themselves as they take on work items.
However if the project board needs to be only scoped to a set of people,
the access *must* be granted through a GitHub team, instead of direct
collaborator access.

**NB**: Not all project boards in use today currently follow this guidance.
We are looking to coalesce existing project boards towards this model, and use
this model for all orgs going forward.

## Repository Guidance

Repositories have additional guidelines and requirements, such as the use of CLA
checking on all contributions. For more details on those please see the
[Kubernetes Template
Project](https://github.com/kubernetes/kubernetes-template-project), and the
[Repository Guidelines](kubernetes-repositories.md)

[GitHub Administration Team]:
/github-management/README.md#github-administration-team
[@kubernetes/owners]: https://github.com/orgs/kubernetes/teams/owners
[kubernetes/org]: https://github.com/kubernetes/org
