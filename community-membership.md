# Community membership

**Note:** This document is in progress

This doc outlines the various responsibilities of contributor roles in Kubernetes.  The Kubernetes
project is subdivided into subprojects under SIGs.  Responsibilities for most roles are scoped to these subprojects.

| Role | Responsibilities | Requirements | Defined by |
| -----| ---------------- | ------------ | -------|
| member | active contributor in the community | sponsored by 2 reviewers.  multiple contributions to the project. | Kubernetes GitHub org member. |
| reviewer | review contributions from other members | history of review and authorship in a subproject | [OWNERS] file reviewer entry. |
| approver | approve accepting contributions | highly experienced and active reviewer + contributor to a subproject | [OWNERS] file approver entry|
| subproject owner | set direction and priorities for a subproject | demonstrated responsibility and excellent technical judgement for the subproject | [sigs.yaml] subproject [OWNERS] file *owners* entry |

## New contributors

[New contributors] should be welcomed to the community by existing members, helped with PR workflow, and directed to
relevant documentation and communication channels.

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

- Enabled [two-factor authentication] on their GitHub account
- Have made multiple contributions to the project or community.  Contribution may include, but is not limited to:
    - Authoring or reviewing PRs on GitHub
    - Filing or commenting on issues on GitHub
    - Contributing to SIG, subproject, or community discussions (e.g. meetings, Slack, email discussion
      forums, Stack Overflow)
- Subscribed to [kubernetes-dev@googlegroups.com]
- Actively contributing to 1 or more subprojects.
- Sponsored by 2 reviewers. **Note the following requirements for sponsors**:
    - Sponsors must have close interactions with the prospective member - e.g. code/design/proposal review, coordinating
      on issues, etc.
    - Sponsors must be reviewers or approvers in at least 1 OWNERS file (in any repo in the Kubernetes GitHub
      organization)
    - Sponsors must be from multiple member companies to demonstrate integration across community.
- Send an email to *kubernetes-membership@googlegroups.com* with:
   - CC: your sponsors on the message
   - Subject: `REQUEST: New membership for <your-GH-handle>`
   - Body: Confirm that you have joined kubernetes-dev@googlegroups.com (e.g. `I have joined
     kubernetes-dev@googlegroups.com`)
   - Body: GitHub handles of sponsors
   - Body: List of contributions (PRs authored / reviewed, Issues responded to, etc)
- Have your sponsoring reviewers reply confirmation of sponsorship: `+1`
- Wait for response to the message
- Have read the [developer guide]

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
- Members can do `/lgtm` on open PRs.
- They can be assigned to issues and PRs, and people can ask members for reviews with a `/cc @username`.
- Tests can be run against their PRs automatically. No `/ok-to-test` needed.
- Members can do `/ok-to-test` for PRs that have a `needs-ok-to-test` label, and use commands like `/close` to close PRs as well.

**Note:** members who frequently contribute code are expected to proactively perform code reviews and work towards
becoming a primary *reviewer* for the subproject that they are active in.

## Reviewer

Reviewers are able to review code for quality and correctness on some part of a subproject.
They are knowledgeable about both the codebase and software engineering principles.

**Defined by:** *reviewers* entry in an OWNERS file in a repo owned by the Kubernetes project.

Reviewer status is scoped to a part of the codebase.

**Note:** Acceptance of code contributions requires at least one approver in addition to the assigned reviewers.

### Requirements

The following apply to the part of codebase for which one would be a reviewer in an [OWNERS] file
(for repos using the bot).

- member for at least 3 months
- Primary reviewer for at least 5 PRs to the codebase
- Reviewed or merged at least 20 substantial PRs to the codebase
- Knowledgeable about the codebase
- Sponsored by a subproject approver
  - With no objections from other approvers
  - Done through PR to update the OWNERS file
- May either self-nominate, be nominated by an approver in this subproject, or be nominated by a robot

### Responsibilities and privileges

The following apply to the part of codebase for which one would be a reviewer in an [OWNERS] file
(for repos using the bot).

- Tests are automatically run for PullRequests from members of the Kubernetes GitHub organization
- Code reviewer status may be a precondition to accepting large code contributions
- Responsible for project quality control via [code reviews]
  - Focus on code quality and correctness, including testing and factoring
  - May also review for more holistic issues, but not a requirement
- Expected to be responsive to review requests as per [community expectations]
- Assigned PRs to review related to subproject of expertise
- Assigned test bugs related to subproject of expertise
- Granted "read access" to kubernetes repo
- May get a badge on PR and issue comments

## Approver

Code approvers are able to both review and approve code contributions.  While code review is focused on
code quality and correctness, approval is focused on holistic acceptance of a contribution including:
backwards / forwards compatibility, adhering to API and flag conventions, subtle performance and correctness issues,
interactions with other parts of the system, etc.

**Defined by:** *approvers* entry in an OWNERS file in a repo owned by the Kubernetes project.

Approver status is scoped to a part of the codebase.

### Requirements

The following apply to the part of codebase for which one would be an approver in an [OWNERS] file
(for repos using the bot).

- Reviewer of the codebase for at least 3 months
- Primary reviewer for at least 10 substantial PRs to the codebase
- Reviewed or merged at least 30 PRs to the codebase
- Nominated by a subproject owner
  - With no objections from other subproject owners
  - Done through PR to update the top-level OWNERS file

### Responsibilities and privileges

The following apply to the part of codebase for which one would be an approver in an [OWNERS] file
(for repos using the bot).

- Approver status may be a precondition to accepting large code contributions
- Demonstrate sound technical judgement
- Responsible for project quality control via [code reviews]
  - Focus on holistic acceptance of contribution such as dependencies with other features, backwards / forwards
    compatibility, API and flag definitions, etc
- Expected to be responsive to review requests as per [community expectations]
- Mentor contributors and reviewers
- May approve code contributions for acceptance

## Subproject Owner

**Note:** This is a generalized high-level description of the role, and the specifics of the subproject owner role's
responsibilities and related processes *MUST* be defined for individual SIGs or subprojects.

Subproject Owners are the technical authority for a subproject in the Kubernetes project.  They *MUST* have demonstrated
both good judgement and responsibility towards the health of that subproject.  Subproject Owners *MUST* set technical
direction and make or approve design decisions for their subproject - either directly or through delegation
of these responsibilities.

**Defined by:** *owners* entry in subproject [OWNERS] files as defined by [sigs.yaml]  *subproject.owners*

### Requirements

The process for becoming an subproject Owner should be defined in the SIG charter of the SIG owning
the subproject.  Unlike the roles outlined above, the Owners of a subproject are typically limited
to a relatively small group of decision makers and updated as fits the needs of the subproject.

The following apply to the subproject for which one would be an owner.

- Deep understanding of the technical goals and direction of the subproject
- Deep understanding of the technical domain of the subproject
- Sustained contributions to design and direction by doing all of:
  - Authoring and reviewing proposals
  - Initiating, contributing and resolving discussions (emails, GitHub issues, meetings)
  - Identifying subtle or complex issues in designs and implementation PRs
- Directly contributed to the subproject through implementation and / or review

### Responsibilities and privileges

The following apply to the subproject for which one would be an owner.

- Make and approve technical design decisions for the subproject.
- Set technical direction and priorities for the subproject.
- Define milestones and releases.
- Mentor and guide approvers, reviewers, and contributors to the subproject.
- Ensure continued health of subproject
  - Adequate test coverage to confidently release
  - Tests are passing reliably (i.e. not flaky) and are fixed when they fail
- Ensure a healthy process for discussion and decision making is in place.
- Work with other subproject owners to maintain the project's overall health and success holistically

## ~~Maintainer~~

**Status:** Removed

The Maintainer role has been removed and replaced with a greater focus on [OWNERS].

[code reviews]: contributors/devel/collab.md
[community expectations]: contributors/guide/community-expectations.md
[developer guide]: contributors/devel/README.md
[two-factor authentication]: https://help.github.com/articles/about-two-factor-authentication
[kubernetes-dev@googlegroups.com]: https://groups.google.com/forum/#!forum/kubernetes-dev
[sigs.yaml]: sigs.yaml
[New contributors]: https://github.com/kubernetes/community/blob/master/CONTRIBUTING.md
[OWNERS]: contributors/guide/owners.md
