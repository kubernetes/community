This is a Work in Progress, documenting approximately how we have been operating up to this point.

# Principles

The Kubernetes community adheres to the following principles:
* Open: Kubernetes is open source. See repository guidelines and CLA, below.
* Welcoming and respectful: See Code of Conduct, below.
* Transparent and accessible: Work and collaboration should be done in public. See SIG governance, below.
* Merit: Ideas and contributions are accepted according to their technical merit and alignment with
  project objectives, [scope](http://kubernetes.io/docs/whatisk8s/), and [design
  principles](contributors/design-proposals/principles.md).

# Code of Conduct

The Kubernetes community abides by the CNCF [code of conduct](https://github.com/cncf/foundation/blob/master/code-of-conduct.md). Here is an excerpt:

_As contributors and maintainers of this project, and in the interest of fostering an open and welcoming community, we pledge to respect all people who contribute through reporting issues, posting feature requests, updating documentation, submitting pull requests or patches, and other activities._

As a member of the Kubernetes project, you represent the project and your fellow contributors. 
We value our community tremendously and we'd like to keep cultivating a friendly and collaborative
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

It is intended that contributors have the opportunity to grow in responsibilities,
privileges, and authority corresponding to the scope, quality, quantity, and duration
of their contributions. Definition of criteria and process is in progress, with preliminary
requirements below.
 
Roles that are currently assumed by project participants are described below,
Focusing on, but not limited to the `kubernetes/kubernetes` repo.

## Code and documentation contributors

New community members:

- [**NEWCOMER**](https://github.com/kubernetes/contrib/issues/1090):
  - Requirements
    - submitted first PR
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
    - we have [expectations](contributors/devel/community-expectations.md) that
      frequent contributors will assist in our code-review process and with project
      maintenance
  - Benefits
    - TBD
  
Established community members:

Established community members are expected to demonstrate their adherence to the principles in this
document, familiarity with project organization, conventions, policies & procedures, etc.,
and technical and/or writing abilities. Role-specific expectations, responsibilities, and requirements
are enumerated below.

- **MEMBER**:
  - Requirements
    - an active contributor for at least 3 months
    - at least 10 merged and/or assigned PRs
    - active enough to be assigned issues and/or PRs, and to be added to a github team
      (e.g., for a SIG) for notification purposes
  - Expectations
    - must enable [GitHub’s two-factor authentication](https://help.github.com/articles/about-two-factor-authentication/)
    - should subscribe to kubernetes-dev@googlegroups.com
    - should read the [developer guide](contributors/devel/README.md)
    - expected to be familiar with project organization, conventions, policies, etc.
  - Benefits
    - trusted enough to run tests on their PRs automatically
    - can issue `@k8s-bot ok to test` for other contributors
    - if they choose public membership, they get a badge on their github profile
- **REVIEWER**:
  - Requirements
    - org member for at least 3 months
    - at least 20 merged and/or assigned PRs, including at least 3 as the primary reviewer
    - familiar enough with some part of the codebase to be in an [OWNERS](contributors/devel/owners.md)
      file as a `reviewer` (in repos using the bot)
    - nominated by an APPROVER for that part of the codebase, with no objections from opant**: active in one or more areas of the project; wide variety of roles are represented
- **SIG Lead**: SIG organizer

## Management roles
- **Team Lead**: tech lead or manager of some team at some company working on K8s; can influence
  priorities of their team members; pragmatically, probably want label/assignment powers
- [**kubernetes-pm**](https://github.com/orgs/kubernetes/teams/kubernetes-pm): help to [manage and
  maintain the project](project-managers/README.md) in ways other than just writing code (e.g. managing
  issues); should subscribe to kubernetes-pm@googlegroups.com

## Rotations
- [**Build Cop**](contributors/devel/on-call-build-cop.md): ensure tests pass, submit queue is workinn
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
- New repos: New subprojects/repositories need to be able to add REVIEWERS, APPROVERS, and MAINTAINERS
  more rapidly than more mature subprojects. Subprojects less than 1 year old will have relaxed time and
  PR requirements (TBD).

## Removal from roles

Most of the above roles require continuous, significant involvement in the project. Kubernetes is a very
high-volume project -- hundreds of PRs and issues per week, with contributors in more than a dozen
timezones.

If someone becomes unable or unwilling to continue in their roles, they may retire. If someone doesn't
Fulfill their role for 90 days or violates the code of conduct, they may be removed from the role
(escalation/vote process TBD). If they wish to resume their role in the future, they may request to return
to it by asking the current members filling that role.

# SIG Governance

In order to standardize Special Interest Group efforts, create maximum transparency, and route contributors to the appropriate SIG, SIGs should follow the guidelines stated below:

- Meet regularly, at least for 30 minutes every 3 weeks, except November and December
- Keep up-to-date meeting notes, linked from the SIG's page in the community repo
- Announce meeting agenda and minutes after each meeting, on their SIG mailing list
- Record SIG meeting and make it publicly available
- Ensure the SIG's mailing list and slack channel are archived
- Report activity in the weekly community meeting at least once every 6 weeks
- Participate in release planning meetings and retrospectives, and burndown meetings, as needed
- Ensure related work happens in a project-owned github org and repository, with code and tests 
  explicitly owned and supported by the SIG, including issue triage, PR reviews, test-failure response,
  bug fixes, etc. 
- Use the above forums as the primary means of working, communicating, and collaborating, as opposed to
  private emails and meetings
- Represent the SIG for the PM group:
  - identify all features in the current release from the SIG
  - track all features (in the repo with all the fields complete)
  - attend your SIG meetings
  - attend the PM group meetings which occur 3-5 times per release
  - identify the annual roadmap
  - advise their SIG as needed

# CLA

All contributors must sign the CNCF CLA, as described [here](CLA.md).

# History

Discussion to finalize the initial content can be found in a [Google Doc](https://docs.google.com/document/d/1UKfV4Rdqi8JcrDYOYw9epRcXY17P2FDc2MENkJjMcas/edit).
Join kubernetes-dev or kubernetes-pm googlegroups to access it.

[![Analytics](https://kubernetes-site.appspot.com/UA-36037335-10/GitHub/governance.md?pixel)]()
