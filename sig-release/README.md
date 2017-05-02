# SIG Release

### Table of Contents
- [Mission](#mission)
- [Join us!](#join-us)
- [Broad responsibilities](#broad-responsibilities)
- [Specific responsibilities](#specific-responsibilities)
- [Relation to the Release Manager and Team](#relation-to-the-release-manager-and-team)
- [Kubernetes Release Team roles](#kubernetes-release-team-roles)
   - [Patch Release Manager](#patch-release-manager)
   - [Kubernetes Release Management Team for Major/Minor Releases](#kubernetes-release-management-team-for-majorminor-releases)
   - [Release Team Lead](#release-team-lead)
   - [Release Team Secondary](#release-team-secondary)
   - [Release Branch Manager](#release-branch-manager)
   - [Docs Lead](#docs-lead)
   - [Features Lead](#features-lead)
   - [Bug Triage Lead](#bug-triage-lead)
   - [Test Infra Lead](#test-infra-lead)
   - [Upgrade Testing Lead](#upgrade-testing-lead)
   - [CI Signal Lead](#ci-signal-lead)
   - [Release Team Shadow](#release-team-shadow)
   - [Individual Contributors](#individual-contributors)
      - [During a patch pelease](#during-a-patch-pelease)
      - [During a major/minor release](#during-a-majorminor-release)
         - [Propose and track feature](#propose-and-track-feature)
         - [Write feature documentation](#write-feature-documentation)
- [Other activities of the Release Team](#other-activities-of-the-release-team)
   - [During "Major" releases](#during-major-releases)
   - [During "Security" releases](#during-security-releases)
- [References](#references)

## Mission
- Production of high quality Kubernetes releases on a reliable schedule
- Ensure there is a [consistent group of community members](https://docs.google.com/document/d/1JAUqKl-lYdYLQ7GUT_9LzqxwQv-PcOdyAxNRZKItajo/edit) [1] in place to support the release process
- Provide guidance and tooling to facilitate the production of automated releases
- Serve as a tightly integrated partner with other SIGs to empower SIGs to integrate their repositories into the release process

## Join us!
- [Google group](https://groups.google.com/forum/#!forum/kubernetes-sig-release)
- [Slack channel](https://kubernetes.slack.com/messages/C56HWQ0TH/)
- [Events and meetings calendar](https://calendar.google.com/calendar/embed?src=coreos.com_regcvcrgvq98lua2ikijg1g1uk%40group.calendar.google.com&ctz=America/Los_Angeles)
- [Meeting agenda and notes](https://docs.google.com/document/d/1vhsixdT58iJFfoGZbpmvI_xnK59XyAjtadu3h6hHPpY/edit?usp=sharing)

## Broad responsibilities
- Ensuring high quality Kubernetes releases
  - Define and staff release roles to manage the resolution of release blocking criteria
  - Define and drive development processes (e.g. merge queues, cherrypicks) and release processes 
    (e.g. burndown meetings, cutting beta releases) with the intent of meeting the release schedule
  - Manage the creation of release specific artifacts including
    - Code branches
    - Binary distributions
    - Release notes
- Continually improving release and development processes
  - Work closely with SIG Contributor Experience to define and build tools to facilitate release process (e.g. dashboards)
  - Work with downstream communities responsible for packaging Kubernetes releases
  - Work with other SIGs to agree upon the responsibilities of thier SIG with respect to the release
  - Define and collect metrics related to the release in order to measure progress over each release
- Collaboration with downstream communities which build artifacts from Kubernetes releases
  - Ensure the appropriate level of integration with publishing the OSS build artifacts

## Specific responsibilities
- Support SIG PM by providing tooling and processes for the generation of release notes
- Coordinate with SIG PM to communicate feature burndown progress during a release cycle
- Manage repositories and tooling dedicated to releasing Kubernetes which at time of chartering these include     
  - kubernetes/release
  - deb/rpm packaging and hosting
  - Container image hosting
- Set and enforce criteria for repository inclusion in a Kubernetes release
  - Governance
  - Stabilization
  - Test coverage
- Set and enforce criteria for code inclusion in a Kubernetes release
  - Feature flags
  - Documentation
  - Test coverage
  - Dashboards
  - Status reports
- Define template and format for communication of release status
  - Ongoing status of the release process
  - Announcement of alpha, beta, release candidate availability
  - Announcement of release availability
- Deriving signal from the following sources
  - [test grid](https://k8s-testgrid.appspot.com/)
  - GitHub [flake issues](https://github.com/kubernetes/kubernetes/issues?q=is%3Aopen+is%3Aissue+label%3Akind%2Fflake)
  - GitHub [bug issues](https://github.com/kubernetes/kubernetes/issues?utf8=%E2%9C%93&q=is%3Aopen%20is%3Aissue%20label%3Akind%2Fbug)
- Identifying owning individuals and SIGs for blocking issues
- Working with SIGs and individuals to drive resolution of open issues
- Building summary of release criteria and status and publish to the community on a regular basis throughout the release cycle
- Manage the contents of `kubernetes/features` along with SIG PM
- Define burndown process
  - use of GitHub labels to signal release blocking status
  - use of GitHub milestones to communicate release blocking issues
  - use of flake GitHub issue count or CI signal as a release blocking status
- Coordinate the resolution of release blocking issues

## Relation to the Release Manager and Team

The Kubernetes Release Team is embedded within SIG Release and is responsible for the day to day work required to
successfully release while the SIG at large is focused on the continued improvement of the release process. Historically
the Release Manager (previously Release Czar)
and later Release Team has assumed the following responsibilities
- Authority to build and publish releases at the published release date under the auspices of the CNCF
- Authority to accept or reject cherrypick merge requests to the release branch
- Authority to accept or reject PRs to the kubernetes/kubernetes master branch during code slush period
- Changing the release timeline as needed if keeping it would materially impact the stability or quality of the release
these responsibilities will continue to be discharged by SIG release through the Release Team. This charter grants SIG 
Release the following additional responsibilities
- Authority to revert code changes which imperil the ability to produce a release by the communicated date or otherwise 
  negatively impact the quality of the release (e.g. insufficient testing, lack of documentation)
- Authority to guard code changes behind a feature flag which do not meet criteria for release
- Authority to close the submit queue to any changes which do not target the release as part of enforcing the code slush
  period
which shall be the discharged by the Release Team.

## Kubernetes Release Team roles

As documented in the [Kubernetes Versioning doc](https://github.com/kubernetes/kubernetes/blob/master/docs/design/versioning.md),
there are 3 types of Kubernetes releases:
- Major (x.0.0)
- Minor (x.x.0)
- Patch (x.x.x)

Major and minor releases are managed by a **Kubernetes Release Management Team**, and patch releases are managed by the **Patch Release Manager**.
Exact roles and duties are defined below.

### Patch Release Manager

Patch releases are managed by the **Patch Release Manager**. Duties of the patch release manager include:
- Ensuring the release branch (e.g. `release-1.5`) remains in a healthy state
  - If the build breaks or any CI for the release branch becomes unhealthy due to a bad merge or infrastructure issue,
    ensure that actions are taken ASAP to bring it back to a healthy state
- Reviewing and approving [cherry picks](https://github.com/kubernetes/community/blob/master/contributors/devel/cherry-picks.md) to the release branch
  - Patch releases should not contain new features, so ensure that cherry-picks are for bug/security fixes only
  - Cherry picks should not destabilize the branch, so ensure that either the PR has had time to stabilize in master
    or will have time to stabilize in the release branch before the next patch release is cut
- Setting the exact schedule (and cadence) for patch releases and actually cutting the [releases](https://github.com/kubernetes/kubernetes/releases)

Current and past patch release managers are listed [here](https://github.com/kubernetes/community/wiki).

### Kubernetes Release Management Team for Major/Minor Releases

Major and Minor releases are managed by the **Kubernetes Release Team** which is responsible for ensuring Kubernetes releases go out on time
(as scheduled) and with high quality (stable, with no major bugs).

Roles and responsibilities within the Kubernetes Release Management Team are as follows.

### Release Team Lead
The Release Team Lead is the person ultimately responsible for ensuring the release goes out on-time with high-quality.
All the roles defined below report to the Release Management Team Lead.
- Establishes and communicates responsibilities and deadlines to release management team members, developers/feature owners, SIG leads, etc
- Escalates and unblocks any issue that may jeopardise the release schedule or quality as quickly as possible
- Finds people to take ownership of any release blocking issues that are not getting adequate attention
- Keeps track of, and widely communicates, the status of the release (including status of all sub-leads, all release blockers, etc)
  and all deadlines leading up to release
- Manages [exception](https://github.com/kubernetes/features/blob/master/EXCEPTIONS.md) process for features that want to merge after code freeze

### Release Team Secondary
The Release Team Secondary is an assistant to the Release Team Lead and works to fill gaps in Release Team staffing or effort. The Secondary should
be familiar with the release process and remain ready to discharge the responsibilities of the Lead in the event the Lead is unavailable.

### Release Branch Manager
- Manages (initiates and enforces) code freeze on main branch as scheduled for the release
  - Ensures no new features are merged after code complete, unless they've been approved by the [exception process](https://github.com/kubernetes/features/blob/master/EXCEPTIONS.md)
- Cuts the `release-x.x` branch at the appropriate time during the milestone
- Ensures release branch (e.g. `release-1.5`) remains in a healthy state for the duration of the major or minor release
  - If the build breaks, or any CI for the release branch becomes unhealthy due to a bad merge or infrastructure issue,
    ensures that actions are taken ASAP to bring it back to a healthy state
- Initiates automatic fast-forwards of the release branch to pick up all changes from master branch, when appropriate
- Reviews and approves [cherry picks](https://github.com/kubernetes/community/blob/master/contributors/devel/cherry-picks.md) to the release branch
  - Ensures onlyl bug/security fixes (but no new features) are cherry-picked after code complete unless approved by the [exception process](https://github.com/kubernetes/features/blob/master/EXCEPTIONS.md)
  - Ensures that cherry-picks do not destabilize the branch by either giving the PR enough time to stabilize in master or giving it enough time to stabilize in the release branch before cutting the release
- Cuts the actual [release](https://github.com/kubernetes/kubernetes/releases)

### Docs Lead
- Sets docs related deadlines for developers and works with Release Management Team Lead to ensure they are widely communicated
- Sets up release branch for docs
- Pings feature owners to ensure that release docs are created on time
- Reviews/merges release doc PRs
- Merges the docs release branch to master to make release docs live as soon as the release is official

### Features Lead
- Compiles the major themes, new features, known issues, actions required, notable changes to existing behavior, deprecations, etc
  and edits them into a release doc checked in to the feature repository (ready to go out with the release)
- Collects and prepares the release notes

### Bug Triage Lead
- Figures out which bugs (whether manually created or automatically generated) should be tracked for the release
- Ensures all bugs being tracked for the release have owners that are responsive
- Ensures all bugs are triaged as blocking or non-blocking
- Ensures all bugs that are blocking are being actively worked on, esp after code complete

### Test Infra Lead
- Sets up and maintains all CI for the release branch

### Upgrade Testing Lead
- Ensures that automated upgrade tests provide a clear go/no-go signal for the release
- Tracks and finds owners for all issues with automated upgrade tests
- Ensures that any gaps in automated upgrade testing are covered by manual upgrade testing
- Organizes the manual upgrade testing efforts, including setting up instructions for manual testing,
  finding manual testing volunteers, and ensuring any issues discovered are communicated widely and fixed quickly

### CI Signal Lead
- Ensures that all non-upgrade test CI provides a clear go/no-go signal for the release
- Tracks and finds owners to fix any issues with any (non-upgrade) tests

### Release Team Shadow
Any Release Team member may select one or more mentees to shadow the release process in order to help fulfill future
Release Team staffing requirements and continue to grow the Kubernetes community in general. Potential mentees should
be prepared to devote a significant amount of time to shadowing their mentor during the release cycle. Successful Release Team Shadows
should be prepared to assume a lead role in a subsequent release.

### Individual Contributors
Release responsiblites of indvidual contributors to the Kubernetes project are captured below

#### During a patch pelease
If you have a patch that needs to be ported back to a previous release (meaning it is a critical bug/security fix), once it is merged to the Kubernetes `master` branch:
- Mark your PR with the milestone corresponding to the release you want to port back to (e.g. `v1.5`), and add the `cherrypick-candidate` label to it
- The Patch Release Manager will then review the PR and if it is ok for cherry-picking, will apply a `cherrypick-approved` label to it
- Once your PR has been marked with the `cherrypick-approved` label by the Patch Release Manager, you should prepare a cherry-pick to the requested branch following the instructions [here](https://github.com/kubernetes/community/blob/master/contributors/devel/cherry-picks.md#how-do-cherrypick-candidates-make-it-to-the-release-branch)

#### During a major/minor release

##### Propose and track feature
If you are developing a feature for Kubernetes, make sure that an issue is open in the [features repository](https://github.com/kubernetes/features/issues). If you are targeting a particular release, make sure the issue is marked with the corresponding release milestone.

Ensure that all code for your feature is written, tested, reviewed, and merged before code freeze date for the target release.

During the code freeze period, fix any bugs discovered with you feature, and write feature documentation.

##### Write feature documentation

1. Make sure your feature for the upcoming release is on the release tracking board (e.g. [link](https://docs.google.com/spreadsheets/d/1nspIeRVNjAQHRslHQD1-6gPv99OcYZLMezrBe3Pfhhg/edit#gid=0) for 1.6).
2. Create a PR with documentation for your feature in the [documents repo](https://github.com/kubernetes/kubernetes.github.io).
    - **Your PR should target the release branch (e.g. [release-1.6](https://github.com/kubernetes/kubernetes.github.io/tree/release-1.6)), not the [`master`](https://github.com/kubernetes/kubernetes.github.io/tree/master) branch.**
      - Any changes to the master branch become live on https://kubernetes.io/docs/ as soon as they are merged, and for releases we do not want docuemntation to go live until the release is cut.
3. Add link to your docs PR in the release tracking board, and notify the docs lead for the release (e.g. [Devin Donnelly](https://www.github.com/devin-donnelly) for 1.6).
4. The docs lead will review your PR and give you feedback.
5. Once approved, the docs lead will merge your PR into the release branch.
6. When the release is cut, the docs lead will push the docs release branch to master, making your docs live on https://kubernetes.io/docs/.

## Other activities of the Release Team

### During "Major" releases
To date no major release has been scheduled, however, SIG Release would be responsible for working closely with SIG PM
and SIG Testing to coordinate this effort across SIGs. The precise work required to produce a major relaese (e.g. 2.0, 3.0)
is undefined.

### During "Security" releases
For all unplanned or embargoed releases
- Facilitate security releases following the under the [Security Release Process](https://github.com/kubernetes/community/blob/master/contributors/devel/security-release-process.md)

## References
1. https://docs.google.com/a/google.com/document/d/1JAUqKl-lYdYLQ7GUT_9LzqxwQv-PcOdyAxNRZKItajo/edit?usp=sharing
