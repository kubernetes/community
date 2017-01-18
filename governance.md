# Principles

The Kubernetes community adheres to the following principles:
* Open: Kubernetes is open source. See repository guidelines and CLA, below.
* Welcoming and respectful: See Code of Conduct, below.
* Transparent and accessible: Work and collaboration should be done in public. See SIG governance, below.
* Merit: Ideas and contributions are accepted according to their technical merit and alignment with project objectives, [scope](http://kubernetes.io/docs/whatisk8s/), and [design principles](contributors/design-proposals/principles.md).

# Code of Conduct

The Kubernetes community abides by the CNCF [code of conduct](https://github.com/cncf/foundation/blob/master/code-of-conduct.md). Here is an excerpt:

_As contributors and maintainers of this project, and in the interest of fostering an open and welcoming community, we pledge to respect all people who contribute through reporting issues, posting feature requests, updating documentation, submitting pull requests or patches, and other activities._

As a member of the Kubernetes project, you represent the project and your fellow contributors. 
We value our community tremendously and we'd like to keep cultivating a friendly and collaborative
environment for our contributors and users. We want everyone in the community to have 
[positive experiences](https://www.cncf.io/blog/2016/12/14/diversity-scholarship-series-one-software-engineers-unexpected-cloudnativecon-kubecon-experience).

# Repository guidelines

All repositories under Kubernetes github orgs, such as kubernetes and kubernetes-incubator,
should follow the procedures outlined in the [incubator document](incubator.md). All code projects
use the [Apache Licence version 2.0](LICENSE). Documentation repositories should use the
[Creative Commons License version 4.0](https://github.com/kubernetes/kubernetes.github.io/blob/master/LICENSE).

# Project Roles

Kubernetes is a large project. It is necessarily a group effort.

There are many ways to participate and contribute.
We value all forms of constructive contribution, no matter how small, even if not
explicitly described below.

It is intended that contributors have the opportunity to grow in responsibilities,
privileges, and authority corresponding to the scope, quality, quantity, and duration
of their contributions. Definition of criteria and process is in progress.

Why would someone want to perform and be accepted into a particular role?
 - To work more efficiently; more permissions reduce development friction
 - Status (have the Kubernetes badge on his/her profile and/or contributions)
 - Ownership
 - etc...
 
Roles that are currently assumed by project participants are described below,
with a focus on the `kubernetes/kubernetes` repo.

## Code and documentation contributors

New community members:

- [**New Contributor**](https://github.com/kubernetes/contrib/issues/1090): a
  couple of PRs; should be welcomed to the community, helped with PR workflow, and
  directed to relevant documentation
- **Active Contributor**: at least 3 merged and/or assigned PRs (which could include documentation
  contributions as well as code), including one in the past month; we have
  [expectations](contributors/devel/community-expectations.md)
  that frequent contributors will assist in our code-review process and with project
  maintenance
  
Established community members:

Established community members are expected to demonstrate their adherence to the principles in this
document, familiarity with project organization, roles, policies, procedures, conventions, etc.,
and technical and/or writing ability. Role-specific expectations, responsibilities, and requirements
are enumerated below.

- **Org Member**: an active contributor for at least 3 months; at least 10 merged and/or assigned PRs; active enough to be useful
  to assign issues to them and add them to a github team (e.g., for a SIG) for notification
  purposes; trusted enough to run tests on their PRs automatically; can issue "@k8s-bot ok to test"
  for other contributors; if they choose public membership, they get a badge on their github profile;
  should subscribe to kubernetes-dev@googlegroups.com; expected to be familiar with
  project organization, roles, policies, procedures, etc.; should read the [developer
  guide](contributors/devel/README.md); must enable
  [two-factor authentication](https://help.github.com/articles/about-two-factor-authentication/)
- **Reviewer**: org member for at least 3 months; at least 20 merged and/or assigned PRs, including 
  at least 3 as the primary reviewer; familiar enough with some part of the codebase to be in some
  [OWNERS](contributors/devel/owners.md) file as a reviewer (in repos using the bot),
  assigned related PRs, assigned relevant test bugs; responsible for project quality control via
  [code reviews](contributors/devel/collab.md); expected to be responsive to
  review requests as per [community expectations](contributors/devel/community-expectations.md);
  can champion incubator repos; must be nominated by an approver for that part of the codebase,
  with no objections from other approvers; should be added to
  [`kubernetes-reviewers`](https://github.com/orgs/kubernetes/teams/kubernetes-reviewers);
  "read access" to kubernetes repo; get a badge on PR and issue comments; may be asked to
  become a reviewer as a precondition for accepting a large code contribution
- **Approver**: in some [OWNERS](contributors/devel/owners.md) file as an approver, which
  will be needed to get code merged; previously a reviewer for that part of the
  codebase for at least 3 months; at least 30 merged and/or assigned PRs, including at least 10 as
  the primary reviewer; expected to be responsive to review requests as per
  [community expectations](contributors/devel/community-expectations.md); expected to 
  mentor contributors and reviewers; demonstrated sound technical judgement; nominated
  by an area/component owner, with no objections from other owners;  may be asked to
  become an approver as a precondition for accepting a large code contribution
- **Area/Component Owner**: in top-level [OWNERS](contributors/devel/owners.md) file for
  some area/component as an approver; design/proposal approval authority for some area 
  of the project, though escalation is still possible; expected to mentor and guide approvers,
  reviewers, and other contributors; may be asked to become an area/component owner as a precondition
  for accepting the contribution of a new component or other major function
- [**kubernetes-maintainers**](https://github.com/orgs/kubernetes/teams/kubernetes-maintainers):
  approver for some part of the codebase for at least 3 months; on project for at least 1 year;
  at least 50 merged and/or assigned PRs, including at least 20 as the primary reviewer;
  write access to repo (assign issues/PRs, add/remove labels and milestones, edit issues and PRs, edit wiki,
  create/delete labels and milestones); technically can lgtm any PR and cause it
  to be merged by the submit queue, but expected to respect OWNERS files; expected to review PRs, fix bugs, maintain and
  improve health and quality of the project, provide user support, mentor and guide approvers,
  reviewers, and other contributors; must apply to `kubernetes-maintainers@googlegroups.com`, with a
  [Champion](https://github.com/kubernetes/community/blob/master/incubator.md) from the existing
  kubernetes-maintainers members and a Sponsor from Project Approvers, with a summary
  of contributions to the project, current project responsibilities, and links to merged and assigned PRs;
  at least 3 of the maintainers must approve the application, with no objections; the application
  expires after 2 weeks if not enough approvals are granted
- **Project Approvers**: approver in [top-level OWNERS file in kubernetes repo](https://github.com/kubernetes/kubernetes/blob/master/OWNERS);
  de-facto project decision makers; technically can 
  approve virtually any PRs; can sponsor incubator repos
- [**API Approver**](https://github.com/orgs/kubernetes/teams/api-approvers):
  lead designers of the project, who are familiar with the 
  design, requirements, mechanics, conventions, style, scope, gotchas, etc. 
  of the API; most beta/GA API changes are vetted by the API approvers
- [**API Reviewer**](https://github.com/orgs/kubernetes/teams/api-reviewers):
  contributors familiar with design, requirements, mechanics, conventions, style,
  scope, gotchas, etc. of the API; have written and/or reviewed Kubernetes APIs

## SIG roles
- **SIG Participant**: active in one or more areas of the project; wide 
  variety of roles are represented
- **SIG Lead**: SIG organizer

## Management roles
- **Team Lead**: tech lead or manager of some team at some company working on 
  K8s; can influence priorities of their team members; pragmatically, 
  probably want label/assignment powers
- [**kubernetes-pm**](https://github.com/orgs/kubernetes/teams/kubernetes-pm): help to [manage and maintain the project](project-managers/README.md) in
  ways other than just writing code (e.g. managing issues); should subscribe to kubernetes-pm@googlegroups.com

## Rotations
- [**Build Cop**](contributors/devel/on-call-build-cop.md): ensure tests pass, submit queue is working, rollback PRs, 
  manually merge as necessary to fix build; should be members of appropriate repo's build-cops github team
  (e.g., [kubernetes-build-cops](https://github.com/orgs/kubernetes/teams/kubernetes-build-cops))
- [**User-Support Rotation**](contributors/devel/on-call-user-support.md): answer questions on stackoverflow, googlegroups, 
  slack, twitter, etc. full time while on duty

## Release roles
- The roles of the individuals/team responsible for major, minor, and patch releases is documented
  [here](https://github.com/kubernetes/community/tree/master/contributors/devel/release). Should be
  members of the appropriate release-managers github team (e.g., 
  [kubernetes-release-managers](https://github.com/orgs/kubernetes/teams/kubernetes-release-managers)).

## Other duty-specific github roles:
- **K8s Org Owner**: can create repos, do ~any github action; the number of
  owners shouldn't scale with the organization's growth, O(1), and optimally it
  should be less than 10 people who are very familiar with project workings and
  distributed across a few time zones and organizations The other repos will
  have distinct sets of people filling some of the above roles, also.

## Other repositories

Guidelines for roles in other repositories are TBD. New subprojects/repositories need to be
able to add reviewers, approvers, and maintainers more rapidly than more mature subprojects.
Subprojects less than 1 year old will have relaxed time and PR requirements.

## Stepping down from roles

Most of the above roles require continuous, significant involvement in the project. If someone
becomes unable or unwilling to continue in their roles, they may retire. If someone doesn't fulfill
their role for 90 days, they will be removed from the role. If they wish to resume their role in the
future, they may request to return to it by asking the current members filling that role.

# SIG Governance

In order to standardize Special Interest Group efforts, create maximum transparency, and route contributors to the appropriate SIG, SIGs should follow the guidelines stated below:

* Meet regularly, at least for 30 minutes every 3 weeks, except November and December
* Keep up-to-date meeting notes, linked from the SIG's page in the community repo
* Announce meeting agenda and minutes after each meeting, on their SIG mailing list
* Record SIG meeting and make it publicly available
* Ensure the SIG's mailing list and slack channel are archived
* Report activity in the weekly community meeting at least once every 6 weeks
* Participate in release planning meetings and retrospectives, and burndown meetings, as needed
* Ensure related work happens in a project-owned github org and repository, with code and tests explicitly owned and supported by the SIG, including issue triage, PR reviews, test-failure response, bug fixes, etc. 
* Use the above forums as the primary means of working, communicating, and collaborating, as opposed to private emails and meetings
* Represent the SIG for the PM group:
 * identify all features in the current release from the SIG
 * track all features (in the repo with all the fields complete)
 * attend your SIG meetings
 * attend the PM group meetings which occur 3-5 times per release
 * identify the annual roadmap
 * advise their SIG as needed

# CLA

All contributors must sign the CNCF CLA, as described [here](CLA.md).

[![Analytics](https://kubernetes-site.appspot.com/UA-36037335-10/GitHub/governance.md?pixel)]()
