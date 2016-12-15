# Principles

The Kubernetes community adheres to the following principles:
* Open: Kubernetes is open source. See repository guidelines and CLA, below.
* Welcoming and respectful: See Code of Conduct, below.
* Transparent and accessible: Work and collaboration should be done in public. See SIG governance, below.

# Code of Conduct

The Kubernetes community abides by the CNCF [code of conduct](https://github.com/cncf/foundation/blob/master/code-of-conduct.md). Here is an excerpt:

_As contributors and maintainers of this project, and in the interest of fostering an open and welcoming community, we pledge to respect all people who contribute through reporting issues, posting feature requests, updating documentation, submitting pull requests or patches, and other activities._

We want everyone in the community to have [positive experiences](https://www.cncf.io/blog/2016/12/14/diversity-scholarship-series-one-software-engineers-unexpected-cloudnativecon-kubecon-experience).

# Repository guidelines

All repositories under Kubernetes github orgs, such as kubernetes and kubernetes-incubator,
should follow the procedures outlined in the [incubator document](incubator.md). All code projects
use the [Apache Licence version 2.0](LICENSE). Documentation repositories should use the
[Creative Commons License version 4.0](https://github.com/kubernetes/kubernetes.github.io/blob/master/LICENSE).

# Project Roles

Kubernetes is a large project. There are many ways to participate and contribute.

Roles that are currently assumed by different participants are described below.

Code and documentation contributors:
- **[New Contributor](https://github.com/kubernetes/contrib/issues/1090)**: a
  couple of PRs
- **Contributor**: more than a couple of PRs (which could include documentation
  contributions as well as code); we have [expectations](contributors/devel/community-expectations.md)
  that frequent contributors will assist in our code-review process and with project
  maintenance
- **Org Member**: active enough to be useful to assign issues to them and add
  them to a github team (e.g., for a SIG) for notification purposes; if they
  choose public membership, they get a badge on their github profile; should
  subscribe to kubernetes-dev@googlegroups.com; expected to be familiar with
  project organization, roles, policies, procedures, etc.; should read the [developer
  guide](contributors/devel/README.md)
- [**kubernetes-collaborators**](https://github.com/orgs/kubernetes/teams/kubernetes-collaborators): "read access" to kubernetes repo; get a badge 
  on PR and issue comments; trusted enough to run tests on their PRs 
  automatically; can issue "@k8s-bot ok to test" for other contributors
- **Reviewer**: In some [OWNERS](contributors/devel/owners.md) file as a reviewer (in repos using the bot),
  assigned related PRs, assigned relevant test bugs; can champion incubator
  repos
- **Approver**: some [OWNERS](contributors/devel/owners.md) file as an approver; will be needed to get code
  merged
  - **Area/Component Owner**: design/proposal approval authority for some area 
    of the project, though escalation is still possible, and most beta/GA API 
    changes are vetted by the API owners
- [**API Approver**](https://github.com/orgs/kubernetes/teams/api-approvers):
  lead designers of the project, who are familiar with the 
  design, requirements, mechanics, conventions, style, scope, gotchas, etc. 
  of the API
- [**kubernetes-maintainers**](https://github.com/orgs/kubernetes/teams/kubernetes-maintainers): write access to repo (assign issues/PRs,
  add/remove labels and milestones, edit issues and PRs, edit wiki,
  create/delete labels and milestones); technically can lgtm any PR and cause it
  to be merged by the submit queue; expected to review PRs and fix bugs related
  to their domains
- **Project Approvers**: approver in [top-level OWNERS file in kubernetes repo](https://github.com/kubernetes/kubernetes/blob/master/OWNERS);
  de-facto project decision makers; technically can 
  approve virtually any PRs; can sponsor incubator repos
  
SIG roles:
- **SIG Participant**: active in one or more areas of the project; wide 
  variety of roles are represented
- **SIG Lead**: SIG organizer

Management roles:
- **Team Lead**: tech lead or manager of some team at some company working on 
  K8s; can influence priorities of their team members; pragmatically, 
  probably want label/assignment powers
- [**kubernetes-pm**](https://github.com/orgs/kubernetes/teams/kubernetes-pm): help to [manage and maintain the project](project-managers/README.md) in
  ways other than just writing code (e.g. managing issues); should subscribe to kubernetes-pm@googlegroups.com

Rotations:
- [**Build Cop**](contributors/devel/on-call-build-cop.md): ensure tests pass, submit queue is working, rollback PRs, 
  manually merge as necessary to fix build
- [**User-Support Rotation**](contributors/devel/on-call-user-support.md): answer questions on stackoverflow, googlegroups, 
  slack, twitter, etc. full time while on duty

Release roles:
- **Minor Release Team**: help with minor release
- **Minor Release Manager**: drive minor release
- **Patch Release Team**: help with patch release
- **Patch Release Manager**: drive patch release

Duty-specific github roles:
- [**kubernetes-admin**](https://github.com/orgs/kubernetes/teams/kubernetes-admin): direct code write/merge access; for build cops and
  release czars only.
- **K8s Org Owner**: can create repos, do ~any github action; the number of
  owners shouldn't scale with the organization's growth, O(1), and optimally it
  should be less than 10 people who are very familiar with project workings and
  distributed across a few time zones and organizations The other repos will
  have distinct sets of people filling some of the above roles, also.
   
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

# CLA

All contributors must sign the CNCF CLA, as described [here](CLA.md).

[![Analytics](https://kubernetes-site.appspot.com/UA-36037335-10/GitHub/governance.md?pixel)]()
