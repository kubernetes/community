# Community membership

**Note:** This document is in progress

This doc outlines the various responsibilities of contributor roles in Kubernetes.  The Kubernetes
project is subdivided into various sub-areas.  Responsibilities for many contributor roles are
scoped by these sub-areas.

| Role | Responsibilities | Requirements | Defined by |
| -----| ---------------- | ------------ | -------|
| member | active contributor in the community | sponsored by 2 reviewers.  multiple contributions to the project. | Kubernetes GitHub org member. |
| reviewer | review contribution from other members | history of review and authorship in an area | OWNERS file reviewer entry. |
| approver | approve accepting contributions | highly experienced and active reviewer + contributor to an area | OWNERS file approver entry|
| owner | set priorities and approve proposals | demonstrated responsibility and good judgement for entire area | OWNERS file approver entries for entire area.  "# owner" comment next to entry. |
| maintainer | cross area ownership of project health | highly experienced contributor active in multiple areas and roles. | GitHub repo write access |

## New contributors

[**New contributors**](https://github.com/kubernetes/contrib/issues/1090) should be welcomed to the community
by existing members, helped with PR workflow, and directed to relevant documentation and communication channels.

**Note:** Individuals may be added as an outside collaborator (with READ access) to a repo in the Kubernetes GitHub
organization without becoming a member.  This will allow them to be assigned issues and PRs until they become a member,
but will not allow tests to be run against their PRs automatically nor allow them to interact with the PR bot.

### Requirements for outside collaborators

- Working on some contribution to the project that would benefit from
  the abillity to have PRs or Issues to be assigned to the contributor
- Have the support of 1 member
  - Find a member who will sponsor you
  - Send an email to kubernetes-membership@googlegroups.com
    - CC: your sponsor
    - Subject: `REQUEST: New outside collaborator for <your-GH-handle>`
    - Body: GitHub handle of sponsor
    - Body: Justification - any contributions or what you will be working on
  - Have your sponsoring member reply confirmation of sponsorship: `+1`
  - Wait for response to the message

## Established community members

Established community members are expected to demonstrate their adherence to the principles in this
document, familiarity with project organization, roles, policies, procedures, conventions, etc.,
and technical and/or writing ability. Role-specific expectations, responsibilities, and requirements
are enumerated below.

## Member

Members are continuously active contributors in the community.  They can have issues and PRs assigned to them,
participate in SIGs through GitHub teams, and pre-submit tests are automatically run for their PRs.
Members are expected to remain active contributors to the community.

**Defined by:** Member of the Kubernetes GitHub organization

### Requirements

- Enabled [two-factor authentication](https://help.github.com/articles/about-two-factor-authentication/) on their GitHub account
- Have made multiple contributions to the project or community.  Contribution may include, but is not limited to:
    - Authoring or reviewing PRs on GitHub
    - Filing or commenting on issues on GitHub
    - Contributing to SIG or community discussions (e.g. meetings, Slack, email discussion forums, Stack Overflow)
- Subscribed to [`kubernetes-dev@googlegroups.com`](https://groups.google.com/forum/#!forum/kubernetes-dev)
- Are actively contributing to 1 or more areas.
- Sponsored by 2 reviewers. **Note the following requirements for sponsors**:
    - Sponsors must have close interactions with the prospective member - e.g. code/design/proposal review, coordinating on issues, etc.
    - Sponsors must be reviewers or approvers in at least 1 OWNERS file (in any repo in the Kubernetes GitHub organization)
    - Not a requirement, but having sponsorship from a reviewer from another company is encouraged (you get a gold star).
- Send an email to *kubernetes-membership@googlegroups.com* with:
   - CC: your sponsors on the message
   - Subject: `REQUEST: New membership for <your-GH-handle>`
   - Body: Confirm that you have joined kubernetes-dev@googlegroups.com (e.g. `I have joined kubernetes-dev@googlegroups.com`)
   - Body: GitHub handles of sponsors
   - Body: List of contributions (PRs authored / reviewed, Issues responded to, etc)
- Have your sponsoring reviewers reply confirmation of sponsorship: `+1`
- Wait for response to the message
- Have read the [developer guide](contributors/devel/README.md)

Example message:

```
To: kubernetes-membership@googlegroups.com
CC: <sponsor1>, <sponsor2>
Subject: REQUEST: New membership for <your-GH-handle>
Body:

I have joined kubernetes-dev@googlegroups.com.

Sponsors:
- <GH handle> / <email>
- <GH handle> / <email>

List of contributions:
- <PR reviewed / authored>
- <PR reviewed / authored>
- <PR reviewed / authored>
- <Issue responded to>
- <Issue responded to>

```

### Responsibilities and privileges

- Responsive to issues and PRs assigned to them
- Responsive to mentions of SIG teams they are members of
- Active owner of code they have contributed (unless ownership is explicitly transferred)
  - Code is well tested
  - Tests consistently pass
  - Addresses bugs or issues discovered after code is accepted

**Note:** members who frequently contribute code are expected to proactively perform code reviews and work towards
becoming a primary *reviewer* for the area that they are active in.

## Reviewer

Reviewers are able to review code for quality and correctness on some part of the project.
They are knowledgeable about both the codebase and software engineering principles.

**Defined by:** *reviewer* entry in an OWNERS file in the Kubernetes codebase.

Reviewer status is scoped to a part of the codebase.

**Note:** Acceptance of code contributions requires at least one approver in addition to the assigned reviewers.

### Requirements

The following apply to the part of codebase for which one would be a reviewer in an
[OWNERS](contributors/devel/owners.md) file (for repos using the bot).

- member for at least 3 months
- Primary reviewer for at least 5 PRs to the codebase
- Reviewed or merged at least 20 substantial PRs to the codebase
- Knowledgeable about the codebase
- Sponsored by an area approver
  - With no objections from other approvers
  - Done through PR to update the OWNERS file
- May either self-nominate, be nominated by an approver in this area, or be nominated by a robot

### Responsibilities and privileges

The following apply to the part of codebase for which one would be a reviewer in an
[OWNERS](contributors/devel/owners.md) file (for repos using the bot).

- Tests are automatically run for PullRequests from members of the Kubernetes GitHub organization
- Code reviewer status may be a precondition to accepting large code contributions
- Responsible for project quality control via [code reviews](contributors/devel/collab.md)
  - Focus on code quality and correctness, including testing and factoring
  - May also review for more holistic issues, but not a requirement
- Expected to be responsive to review requests as per [community expectations](contributors/devel/community-expectations.md)
- Assigned PRs to review related to area of expertise
- Assigned test bugs related to area of expertise
- Added to [`kubernetes-reviewers`](https://github.com/orgs/kubernetes/teams/kubernetes-reviewers)
- Granted "read access" to kubernetes repo
- Can champion incubator repos
- May get a badge on PR and issue comments

## Approver

Code approvers are able to both review and approve code contributions.  While code review is focused on
code quality and correctness, approval is focused on holistic acceptance of a contribution including:
backwards / forwards compatibility, adhering to API and flag conventions, subtle performance and correctness issues,
interactions with other parts of the system, etc.

**Defined by:** *approver* entry in an OWNERS file in the kubernetes codebase

Approver status is scoped to a part of the codebase.

### Requirements

The following apply to the part of codebase for which one would be an approver in an
[OWNERS](contributors/devel/owners.md) file (for repos using the bot).

- Reviewer of the codebase for at least 3 months
- Primary reviewer for at least 10 substantial PRs to the codebase
- Reviewed or merged at least 30 PRs to the codebase
- Nominated by an area/component owner
  - With no objections from other owners
  - Done through PR to update the top-level OWNERS file

### Responsibilities and privileges

The following apply to the part of codebase for which one would be an approver in an
[OWNERS](contributors/devel/owners.md) file (for repos using the bot).

- Approver status may be a precondition to accepting large code contributions
- Demonstrate sound technical judgement
- Responsible for project quality control via [code reviews](contributors/devel/collab.md)
  - Focus on holistic acceptance of contribution such as dependencies with other features, backwards / forwards compatibility, API and flag definitions, etc
- Expected to be responsive to review requests as per [community expectations](contributors/devel/community-expectations.md);
- Mentor contributors and reviewers
- May approve code contributions for acceptance

## Owner

Owners of an area / component are approvers of an entire area that have demonstrated good judgement and
responsibility.  Owners accept design proposals and approve design decisions for their area of ownership.

**Defined by:** *approver* entry in the top-level OWNERS file for the area in the kubernetes codebase.
May have a comment (e.g. `# owner`) next to the approver entry indicating that the individual is an owner.

### Requirements

The following apply to the area / component for which one would be an owner.

- Originally authored or contributed major functionality to an area
- An approver in the **top-level** [OWNERS](contributors/devel/owners.md) files

### Responsibilities and privileges

The following apply to the area / component for which one would be an owner.

- Owner status may be a precondition to accepting a new component or piece of major functionality
- Design/proposal approval authority over the area / component, though escalation to [`kubernetes-maintainers`](https://groups.google.com/forum/#!forum/kubernetes-maintainers) is possible.
- Mentor and guide approvers, reviewers, and members.

## Maintainer

[**Kubernetes project maintainers**](https://github.com/orgs/kubernetes/teams/kubernetes-maintainers) work holistically
across the project to maintain its health and success.  They are typically involved in multiple different areas, and
have made substantial contributions both through code and broader organizational efforts.

**Defined by:** *write* access to the kubernetes GitHub repo

TODO: Determine if this role is outdated and needs to be redefined or merged into owner role.

### Requirements

- Approver for some part of the codebase for at least 3 months
- Member for at least 1 year
- Primary reviewer for 20 substantial PRs
- Reviewed or merged at least 50 PRs
- Apply to [`kubernetes-maintainers`](https://github.com/orgs/kubernetes/teams/kubernetes-maintainers), with:
  - A [Champion](https://github.com/kubernetes/community/blob/master/incubator.md#faq) from the existing
    kubernetes-maintainers members
  - A Sponsor from Project Approvers
  - Summary of contributions to the project
  - Current project responsibilities
  - Links to merged and assigned PRs
- At least 3 of the maintainers must approve the application, with no objections
- Application expires after 2 weeks if not enough approvals are granted

### Responsibilities and privileges

- Write access to repo (assign issues/PRs, add/remove labels and milestones, edit issues and PRs, edit wiki, create/delete labels and milestones)
- Capable of directly applying lgtm + approve labels for any PR, causing it to be merged by the submit queue - but permitted only in rare instances
   - Expected to respect OWNERS files approvals and use standard procedure for merging code
- Expected to work to holistically maintain the health of the project through:
  - Reviewing PRs
  - Fixing bugs
  - Providing user support
  - Mentoring and guiding approvers, reviewers, and other contributors
