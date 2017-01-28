Last update: 1/27/2017

This is a Work in Progress, documenting approximately how we have been operating up to this point,
and formalizing some previously informal conventions.

# Principles

The Kubernetes community adheres to the following principles:
* Open: Kubernetes is open source. See repository guidelines and CLA, below.
* Welcoming and respectful: See Code of Conduct, below.
* Transparent and accessible: Work and collaboration should be done in public. See SIG governance, below.
* Merit: Ideas and contributions are accepted according to their technical merit and alignment with
  project objectives, [scope](http://kubernetes.io/docs/whatisk8s/), and [design
  principles](contributors/design-proposals/principles.md). People are promoted in responsibility
  based on the scope, quality, quantity, and duration of past contributions. See Project Roles, below.

# Code of Conduct

The Kubernetes community abides by the CNCF [code of conduct](https://github.com/cncf/foundation/blob/master/code-of-conduct.md). Here is an excerpt:

_As contributors and maintainers of this project, and in the interest of fostering an open and welcoming community, we pledge to respect all people who contribute through reporting issues, posting feature requests, updating documentation, submitting pull requests or patches, and other activities._

As a member of the Kubernetes project, you represent the project and your fellow contributors. 
We value our community tremendously and it is critical that we cultivate a friendly and collaborative
environment for our contributors and users. We want everyone in the community to have 
[positive experiences](https://www.cncf.io/blog/2016/12/14/diversity-scholarship-series-one-software-engineers-unexpected-cloudnativecon-kubecon-experience).

# Repository guidelines

All repositories under Kubernetes github orgs, such as kubernetes and kubernetes-incubator,
should follow the procedures outlined in the [incubator document](incubator.md). All code projects must
use the [Apache Licence version 2.0](LICENSE). Documentation repositories must use the
[Creative Commons License version 4.0](https://github.com/kubernetes/kubernetes.github.io/blob/master/LICENSE).

TODO: Need policies/procedures for donated code (e.g., helm, kubernetes-anywhere, kompose, kargo).

# Project Roles

Kubernetes is a large project. It is necessarily a group effort.

There are many ways to participate and contribute. We value all forms of constructive contribution
no matter how small, even if not explicitly described below.

Contributors have the opportunity to grow in responsibilities, privileges, and authority corresponding
to the scope, quality, quantity, and duration of their contributions. Definition of criteria and process
is in progress, with preliminary requirements below.
 
Roles that are currently assumed by project participants are described below, focusing on, but not
limited to the `kubernetes/kubernetes` repo.

## Code and documentation contributors


The “contributor ladder”.

New community members:

- [**NEWCOMER**](https://github.com/kubernetes/contrib/issues/1090):
  - Requirements
    - submitted first PR
      - TODO: investigate feasibility of extending this to first issue filed
  - Expectations
    - hopefully will contribute again in the future
  - Benefits
    - welcomed to the community
    - helped with PR workflow
    - directed to relevant documentation
- **CONTRIBUTOR**:
  - Requirements
    - at least 3 merged and/or assigned PRs (which could include documentation
      contributions as well as code), including one in the past month
  - Expectations
    - we [expect](contributors/devel/community-expectations.md) that
      frequent contributors will assist in our code-review process and with project
      maintenance
  - Benefits
    - TBD: swag
  
Established community members:

Established community members are expected to demonstrate their adherence to the principles in this
document, familiarity with project organization, conventions, and policies & procedures, as well as technical and/or writing abilities. Role-specific expectations, responsibilities, and requirements
are enumerated below. Because github teams are not visible by nonmembers of the org and because changes
to them are not transparent, the primary documentation of role membership beyond MEMBER should be 
maintained in OWNERS files in the repository.

- **MEMBER**:
  - Requirements
    - an active contributor for at least 3 months
    - authored and/or reviewed at least 10 merged PRs
      - TODO: figure out how to count reviews
    - active enough to be assigned issues and/or PRs, and to be added to a github team
      (a SIG, for example) for notification purposes
    - nomination process TBD
  - Expectations
    - must enable [GitHub’s two-factor authentication](https://help.github.com/articles/about-two-factor-authentication/)
    - should subscribe to kubernetes-dev@googlegroups.com and kubernetes-dev-announce@googlegroups.com
    - should read the [developer guide](contributors/devel/README.md)
    - expected to be familiar with project organization, convchnical vision and judgement
    - demonstrated empathy for the user and open-source developer perspective
    - committed to project's mission and culture
    - In `leads` list in [top-level OWNERS file in kubernetes
      repo](https://github.com/kubernetes/kubernetes/blob/master/OWNERS)
      - Proposed initial list comprised of long-time senior project leads (formerly top-level approvers):
        bgrant0607, brendandburns, dchen1107, jbeda, lavalamp, smarterclayton, thockin
    - nomination/application process TBD
    - cap on number of members TBD
  - Expectations
    - TBD
    - provide overall technical guidance and vision for the project
    - maintain the [definition of the project](https://kubernetes.io/docs/whatisk8s/)
    - decide project structure, such as system layers (e.g., core) and repository breakdown
    - assist SIGs in identifying areas of overlapping technical and ownership responsibility
    - resolve technical escalations in the cases of OWNER and SIG lead disagreements 
    - meet monthly by videoconference or in person
    - TODO: office hours?
  - Benefits
    - project decision makers
    - technically can approve virtually any PRs
    - can [Sponsor incubator repos](incubator.md)
    - can Sponsor MAINTAINERS


## Orthogonal technical roles

API REVIEWER and APPROVER are called out specifically because the API is critical to the
identity of Kubernetes and is a horizontal area that crosses directories and SIGs.

- [**API REVIEWER**]:
  - Requirements
    - tenure requirement TBD
    - nomination process TBD
    - initial members TBD
    - have written and/or reviewed Kubernetes APIs
    - familiar enough with design, requirements, mechanics, conventions, style,
      scope, gotchas, etc. of the API to be in `kubernetes/pkg/api` and `kubernetes/pkg/apis`
      OWNERS files `reviewers` lists
  - Expectations
    - review API changes and proposals in their functional area
    - ensure Kubernetes has a consistent [API style](contributors/devel/api-conventions.md),
      patterns, and philosophies
    - guide development of new APIs
  - Benefits
    - TBD
    - added to [`api-reviewers`](https://github.com/orgs/kubernetes/teams/api-reviewers) github team
- **API APPROVER**:
  - Requirements
    - tenure requirement TBD
    - nomination process TBD
    - initial members TBD (historically bgrant0607, thockin, smarterclayton, erictune)
    - designed and reviewed several APIs in the system
    - familiar with the design, requirements, mechanics, conventions, style,
      scope, gotchas, etc. of the API to be in `kubernetes/pkg/api` and `kubernetes/pkg/apis`
      OWNERS files `approvers` lists
    - ensure Kubernetes has a consistent [API style](contributors/devel/api-conventions.md),
      patterns, and philosophies
    - guide development of new APIs
    - codify new patterns for novel problems
  - Expectations
    - review and approve API changes and proposals in their functional area
  - Benefits
    - TBD
    - added to [`api-approvers`](https://github.com/orgs/kubernetes/teams/api-approvers) github team

## SIG roles
- **SIG PARTICIPANT**: active in one or more areas of the project; wide variety of roles are represented
- **SIG LEAD**: SIG organizer

## Management roles
- [**PM**](https://github.com/orgs/kubernetes/teams/kubernetes-pm): help to [manage and
  maintain the project](project-managers/README.md) in ways other than just writing code (e.g. managing
  issues); should subscribe to kubernetes-pm@googlegroups.com
  - Requirements, expectations, benefits TBD
  - TODO: owns feature tracking and roadmap planning
- **TEAM LEAD**: tech lead or manager of some team at some company working on K8s; can influence
  priorities of their team members; pragmatically, probably want label/assignment powers
  - Requirements, expectations, benefits TBD
    - e.g., identify and resolve staffing gaps (engineering, docs, test, release, ...), effort gaps
     (tragedy of the commons), expertise mismatches, priority conflicts, personnel conflicts
  - meeting requirements TBD
- TODO: ownership of process and organization improvement
  - The project needs managers -- people to think about and improve the organization of the project’s
    contributors, processes, etc. There is more description in the 
    [three-branches proposal](https://github.com/kubernetes/community/issues/295) and in the
    [elders proposal discussion](https://github.com/kubernetes/community/pull/267#issuecomment-273715158).
    - TODO: what processes should be covered by PM, SIG Contributor Experience, and release managers
- **PROCESS REVIEWER**:
  - Requirements
    - Have designed, driven, implemented, and rolled out new processes for the project
  - Expectations
    - Review proposals for new processes
    - Guide someone who is designing, driving, implementing, and/or rolling out a new process
  - Benefits
    - TBD
- **PROCESS APPROVER**:
  - Requirements
    - TBD
    - Have designed, driven, implemented, and rolled out new processes for the project
  - Expectations
    - Approve proposals for new processes
    - Ensure processes are consistent and effective, and that we have adequate means of measuring
      their effectiveness and efficiency
    - Ensure that decisions are made in a rational and transparent way
    - Ensure that the policies of the project are documented, communicated, and followed
    - Mentor new PROCESS REVIEWERS
  - Benefits
    - TBD
    - Gratitude of contributors for a smoothly run project

## Rotations
- [**Build Cop**](contributors/devel/on-call-build-cop.md): ensure tests pass, submit queue is working,
  rollback PRs, manually merge as necessary to fix build; should be members of appropriate repo's
  build-cops github team (e.g.,
  [kubernetes-build-cops](https://github.com/orgs/kubernetes/teams/kubernetes-build-cops))
- [**User-Support Rotation**](contributors/devel/on-call-user-support.md): answer questions on
  stackoverflow, googlegroups, slack, twitter, etc. full time while on duty

## Release roles
- The roles of the individuals/team responsible for major, minor, and patch releases is documented
  [here](https://github.com/kubernetes/community/tree/master/contributors/devel/release). Should be
  members of the appropriate release-managers github team (e.g., 
  [kubernetes-release-managers](https://github.com/orgs/kubernetes/teams/kubernetes-release-managers)).

## Other duty-specific github roles:
- [**Github Org Owner**](https://github.com/orgs/kubernetes/people?utf8=%E2%9C%93&query=%20role%3Aowner):
  can create repos, and do any github action; the number of
  owners shouldn't scale with the organization's growth, O(1), and optimally it
  should be less than 20 people who are very familiar with project workings and
  distributed across a few time zones and organizations The other repos will
  have distinct sets of people filling some of the above roles, also.

## Procedural roles

- **Champion**
  - A Champion is the primary point of contact for guiding someone through a process, such
    as [creating a new incubator repo](incubator.md) or becoming a maintainer
  - The majority of the mentorship, review, and advice regarding Kubernetes community norms and
    processes will come from the Champion
  - Potential Champions come from a group of existing Kubernetes contributors, such as REVIEWERS
    or MAINTAINERS -- which group depends on the particular process 
- **Sponsor**
  - A Sponsor is an approver for initiating a process, such as creating a new incubator repo or
    adding a new maintainer
  - Potential Sponsors come from a very small set of senior Kubernetes contributors (typically LEADS)
  - The idea is that by relying on this small set of Kubernetes Community members to approve
    will ensure consistency and preserve the culture and integrity of the project.
  - Being a Sponsor is a minor advisory role

## Other repositories

Guidelines for roles in other repositories are TBD, but some known special cases are called out below. 

Notable repositories/categories:

- `kubernetes.github.io`: In addition to the typical repo-oriented roles, `kubernetes-pm` has write
  permission to this repo in order to manage issues (labels, milestones, etc.), and
  `kubernetes-maintainers` and `kubernetes-reviewers` have write and read permissions, as well
- `features`: In addition to the typical repo-oriented roles, `kubernetes-pm` has write permission
  to this repo in order to manage issues (labels, milestones, etc.) and the issue template, and
  `kubernetes-maintainers` and `kubernetes-reviewers` have write and read permissions, as well
- `community`: This repository is intended to be fairly open to the community. `kubernetes-maintainers`,
  `kubernetes-pm`, SIG-related github teams, and other trusted groups may be given write access in
  order to help maintain SIG and contributor documentation
- `test-infra`: In addition to the typical repo-oriented roles, `kubernetes-build-cops` have admin
  access to this repository in order to fix CI problems.
- `release`: In addition to the typical repo-oriented roles, `kubernetes-build-cops` have admin access
  to this repository, `kubernetes-release-managers` have admin access, `kubernetes-maintainers` has
  write permissions, and `kubernetes-reviewers` has read permissions
- Donated repos (e.g., `heapster`, `helm`, `kompose`): TBD
- [Incubator repos](incubator.md): New subprojects/repositories need to be able to add REVIEWERS,
  APPROVERS, and MAINTAINERS more rapidly than more mature subprojects. Subprojects less than 1 year old
  will have relaxed time and PR requirements (TBD).

## Removal from roles

Most of the above roles require continuous, significant involvement in the project. Kubernetes is a very
high-volume project -- hundreds of PRs and issues per week, with contributors in more than a dozen
timezones.

If someone becomes unable or unwilling to continue in their roles, they may retire. If someone doesn't
fulfill their role for 90 days or violates the code of conduct, they may be removed from the role
(escalation/vote process TBD). If they wish to resume their role in the future, they may request to return
to it by asking the current members filling that role.

# Special Interest Group (SIG) Governance


[SIGs](README.md#special-interest-groups-sig-and-working-groups) are the subteams of the project.
SIGs own code and/or documentation of the project, share knowledge, bring new members up to speed, and
more.

We have a couple dozen SIGs, and contributors in more than a dozen timezones. People want to be able
to follow what's going on without attending every meeting, which is impossible, especially for
non-American timezones.

In order to standardize Special Interest Group efforts, maximize transparency, and route contributors
to the appropriate SIG, SIGs should follow the guidelines stated below:

- Meet regularly, at least for 30 minutes every month except November and December
- Keep up-to-date meeting notes, linked from the SIG's page in the community repo
- Announce meeting agenda before each meeting and post minutes after, on their SIG mailing list
- Record SIG meeting and make it publicly available
- Ensure the SIG's mailing list and slack channel are archived
- Report activity in the weekly community meeting at least once every 6 weeks
- Participate in release planning meetings and retrospectives, and burndown meetings, as needed.
  When the right people aren't present in such meetings, it can put the project at risk, such as
  by slipping the release.
- Development of new code  happens in a project-owned github org and repository, with code and tests 
  explicitly owned and supported by the SIG, including issue triage, PR reviews, test-failure response,
  bug fixes, etc.  An exception is when the SIG contributes to existing OSS projects not on github (e.g.
  Apache Spark) or which have their own github org (e.g. Tensorflow) to improve integration with
  Kubernetes 
- Use the above forums as the primary means of working, communicating, and collaborating, as opposed to
  private emails and meetings
- Google Docs should be shared with at least the SIG discussion group and kubernetes-dev
- Represent the SIG for the PM group (either a SIG liaison to the PM group or a PM liaison to the SIG):
  - identify all features in the current release from the SIG
  - track all features (in the repo with all the fields complete)
  - attend your SIG meetings
  - attend the PM group meetings which occur 3-5 times per release
  - identify the annual roadmap
  - advise their SIG as needed


Not all of the responsibilities need to be shouldered by the SIG lead.

# CLA

All contributors must sign the CNCF CLA, as described [here](CLA.md).

# Process for changing this document

TBD

# History

Discussion to finalize the initial content can be found in a [Google Doc](https://docs.google.com/document/d/1UKfV4Rdqi8JcrDYOYw9epRcXY17P2FDc2MENkJjMcas/edit).
Join kubernetes-dev or kubernetes-pm googlegroups to access it.

[![Analytics](https://kubernetes-site.appspot.com/UA-36037335-10/GitHub/governance.md?pixel)]()
