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
If needed, please contact the steering committee and CNCF prior to moving any
code in.

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

The process for creating and removing new repositories is detailed below.

### Creating Repositories

#### Non-staging repositories

For non-staging repositories, suggestions on how to create
a new repository are described below.

  * Ensure that the repo creation request has appropriate approvals
  as per the rules mentioned above.
  * Using the organization and repository name mentioned in the repo creation
  request, create a new repo with:
    * The [kubernetes-template-project] repo as the [template repo]
    * `Description` as mentioned in the repo creation request
  * Clone the newly created repo locally.
  * Make the following changes:
    * If the request references a team to be listed in the `OWNERS`
    file, update the `OWNERS_ALIASES` file to remove the steering-committee
    alias and add a new alias for the team with members populated as per the
    GitHub team. If the request does not reference a team, remove the
    `OWNERS_ALIASES` file.
    * Update the OWNERS file as per the request. If the repo is a
    [SIG Repository], add a labels entry for the SIG that the repo belongs to.
    * Update the `SECURITY_CONTACTS` file as per the request. Note that aliases
    cannot be used in this case so expand the team, if specified.
    * Create a new commit with the message *Update OWNERS, OWNERS_ALIASES and
    SECURITY_CONTACTS*.
  * Push the new commit directly to the master branch.
  * If the repo is a [SIG Repository], add a new topic of the form
  `k8s-sig-<sig-name-repo-belongs-to>` using the *Manage Topics* option.
  * Create a PR against [kubernetes/org] to add teams as per the [team guidance](#team-guidance)
  for alloting repo admin and write access.
  * Once the above PR is merged and the postsubmit has run, the new GitHub teams
  will be created. In the *Collaborators and Teams* section in Settings,
  assign the new teams appropriate access to the repo.
  * Ask the author of the repo creation request to add the repo
  as a part of a subproject in [`sigs.yaml`](/sigs.yaml).

#### Staging Repositories

If the repository is a staging repository, there are some deviations
from the above procedure:

  * The repository **must** have an initial empty commit. The contents of the
  repo will be populated from staging by the [publishing-bot].
  * Grant the [@kubernetes/stage-bots] team admin access to the repo.
  * Setup branch protection and enable access to the
  `stage-bots` team by adding the repo in
  [`prow/config.yaml`](https://git.k8s.io/test-infra/prow/config.yaml). See
  [kubernetes/test-infra#9292](https://github.com/kubernetes/test-infra/pull/9292)
  for an example.
  * Once the repo has been created, add the repo to
  [`hack/fetch-all-latest-and-push.sh`](https://git.k8s.io/publishing-bot/hack/fetch-all-latest-and-push.sh)
  in the [publishing-bot] repo.

<!-- TODO: Add suggestions for how to migrate existing repos -->

### Removing Repositories

When a repository has been deemed eligible for removal, we take the following
steps:

  * Ownership of the repo is transferred to the [kubernetes-retired] GitHub
    organization
  * The repo description is edited to start with the phrase "[EOL]"
  * All open issues and PRs are closed
  * All external collaborators are removed
  * All webhooks, apps, integrations or services are removed
  * GitHub Pages are disabled
  * Remove all teams associated with the repo
  * The repo is marked as archived using [GitHub's archive feature]
  * Remove the repo from [sigs.yaml]
  * The removal is announced on the kubernetes-dev mailing list and community
    meeting

This maintains the complete record of issues, PRs and other contributions,
leaves the repository read-only, and makes it clear that the repository should
be considered retired and unmaintained.

In case a repository has only the initial commits adding template files
and no additional activity, it can be completely deleted.

[GitHub Administration Team]:
/github-management/README.md#github-administration-team
[GitHub's archive feature]:
https://help.github.com/articles/archiving-a-github-repository/
[@kubernetes/owners]: https://github.com/orgs/kubernetes/teams/owners
[kubernetes/org]: https://github.com/kubernetes/org
[publishing-bot]: https://github.com/kubernetes/publishing-bot
[@kubernetes/stage-bots]: https://github.com/orgs/kubernetes/teams/stage-bots
[kubernetes-retired]: https://github.com/kubernetes-retired
[kubernetes-template-project]: https://github.com/kubernetes/kubernetes-template-project
[SIG Repository]: /github-management/kubernetes-repositories.md##sig-repositories
[template repo]: https://help.github.com/en/articles/creating-a-repository-from-a-template
