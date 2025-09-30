# GitHub Permissions

GitHub provides a limited permissions model for organizations and repositories.
It lacks granularity, and for the most part is "all or nothing". This doesn't
scale well with the size and velocity of the Kubernetes project.

We have created a number of automated systems/bots to allow us to work around
these limitations. Authorized users can issue [bot commands] to execute actions
against PRs and issues without having direct GitHub access to run these
actions. [OWNERS] files are used to gate approvals to merge, and most merges are
handled by automation. This allows us the flexibility to delegate access to
small or large groups of users, without providing direct write or admin access.

That said, there are some actions that are so infrequent, or so complex that
automation isn't a good fit. There is also a need for a small set of users that
can act as a backstop for setting up and maintaining this automation, and
manual intervention if needed.

## Organization Level Permissions

GitHub provides [two access levels][org permissions] to an organization: owner,
and member.

### Owner

Organization owners have full access to the organization, including the ability
to modify billing information and even delete the entire organization.
Therefore, we are very cautious about giving more people this access than
really need it.

There are certain actions that require org owner access:
- Invite or remove members from the organization (handled by [peribolos])
- Access the organization audit log
- Create new repositories
- Transfer repositories
- Approve GitHub application integrations

In the Kubernetes project, this role is held by the
[GitHub Administration Team].

### Member

Organization members are granted "read" access to all repositories in the org,
are able to be assigned issues and PRs, and are able to mention and join
teams. This is the base level of access to an organization.

Our automation tools look for organization membership as a permissions level
for running certain bot commands. The [bot commands] list details which
commands are restricted to org members.

Org membership is granted as a part of becoming a member of the Kubernetes
community as defined in the [community membership] document.

## Repository Level Permissions

GitHub provides [three access levels][repo permissions] to a repository: admin,
write, and read.

### Admin

A repository admin has full access to the repository, and is able to modify any
repository-scoped setting, including renaming or deleting a repository,
manually merge code, and override/change branch protection settings. This is a
trusted role, and should only be given to a limited number of senior
maintainers of a repository.

In most cases, this level of access should not be necessary as the majority of
actions will be able to be implemented by automation. Certain actions like
creating a release may still need this level of access.

<!--- TODO(cblecker): Define specific roles that need this. -->

### Write

Providing direct write access to a repository exposes a number of settings that
are not normally available, including:
- The ability to manually add and remove labels from issues/PRs
- The ability to push new branches
- Manually open and close issues/PRs

While users with write access cannot override branch protection settings and
manually merge PRs, they can manually apply labels like `lgtm`/`approve`,
bypassing normal processes. Write access is being phased out as the majority
of actions are implemented via automation.

### Read

This is the default level of access that is provided to org members on every
repo in the organization. Read access allows you to be assigned issues and PRs
in the repository, but that's about it. It is provided by default to every
member in the organization.


[bot commands]: https://go.k8s.io/bot-commands
[community membership]: /community-membership.md
[GitHub Administration Team]: /github-management/README.md#github-administration-team
[org permissions]:
https://help.github.com/articles/permission-levels-for-an-organization/
[OWNERS]: /contributors/guide/owners.md
[peribolos]: https://sigs.k8s.io/prow/cmd/peribolos
[repo permissions]:
https://help.github.com/articles/repository-permission-levels-for-an-organization/
