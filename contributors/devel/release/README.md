# Kubernetes Release Roles
**Table of Contents**
* [Patch Release Manager](#patch-release-manager)
* [Kubernetes Release Management Team for Major/Minor Releases](#kubernetes-release-management-team-for-majorminor-releases)
* [Individual Contributors](#individual-contributors)

This document captures the requirements and duties of the individuals responsible for Kubernetes releases.

As documented in the [Kubernetes Versioning doc](https://github.com/kubernetes/community/blob/master/contributors/design-proposals/release/versioning.md), there are 3 types of Kubernetes releases:
* Major (x.0.0)
* Minor (x.x.0)
* Patch (x.x.x)

Major and minor releases are managed by a **Kubernetes Release Management Team**, and patch releases are managed by the **Patch Release Manager**. Exact roles and duties are defined below.

## Patch Release Manager

Patch releases are managed by the **Patch Release Manager**. Duties of the patch release manager include:
* Ensuring the release branch (e.g. `release-1.5`) remains in a healthy state.
  * If the build breaks or any CI for the release branch becomes unhealthy due to a bad merge or infrastructure issue, ensure that actions are taken ASAP to bring it back to a healthy state.
* Reviewing and approving [cherry picks](https://github.com/kubernetes/community/blob/master/contributors/devel/cherry-picks.md) to the release branch.
  * Patch releases should not contain new features, so ensure that cherry-picks are for bug/security fixes only.
  * Cherry picks should not destabilize the branch, so ensure that either the PR has had time to stabilize in master or will have time to stabilize in the release branch before the next patch release is cut.
* Setting the exact schedule (and cadence) for patch releases and actually cutting the [releases](https://github.com/kubernetes/kubernetes/releases).

See the [Patch Release Manager Playbook](patch-release-manager.md) for more details.

Current and past patch release managers are listed [here](https://github.com/kubernetes/community/wiki).

## Kubernetes Release Management Team for Major/Minor Releases

Major and Minor releases are managed by the **Kubernetes Release Management Team** which is responsible for ensuring Kubernetes releases go out on time (as scheduled) and with high quality (stable, with no major bugs).

Roles and responsibilities within the Kubernetes Release Management Team are as follows.

#### Release Management Team Lead
The Release Management Team Lead is the person ultimately responsible for ensuring the release goes out on-time with high-quality.  All the roles defined below report to the Release Management Team Lead.
* Establishes and communicates responsibilities and deadlines to release management team members, developers/feature owners, SIG leads, etc.
* Escalates and unblocks any issue that may jeopardise the release schedule or quality as quickly as possible.
* Finds people to take ownership of any release blocking issues that are not getting adequate attention.
* Keeps track of, and widely communicates, the status of the release (including status of all sub-leads, all release blockers, etc) and all deadlines leading up to release.
* Manages [exception](https://github.com/kubernetes/features/blob/master/EXCEPTIONS.md) process for features that want to merge after code freeze.

#### Release Branch Manager
* Manages (initiates and enforces) code freeze on main branch as scheduled for the release.
  * Ensures no new features are merged after code complete, unless they've been approved by the [exception process](https://github.com/kubernetes/features/blob/master/EXCEPTIONS.md).
* Cuts the `release-x.x` branch at the appropriate time during the milestone.
* Ensures release branch (e.g. `release-1.5`) remains in a healthy state for the duration of the major or minor release.
  * If the build breaks, or any CI for the release branch becomes unhealthy due to a bad merge or infrastructure issue, ensures that actions are taken ASAP to bring it back to a healthy state.
* Initiates automatic fast-forwards of the release branch to pick up all changes from master branch, when appropriate.
* Reviews and approves [cherry picks](https://github.com/kubernetes/community/blob/master/contributors/devel/cherry-picks.md) to the release branch.
  * Ensures onlyl bug/security fixes (but no new features) are cherry-picked after code complete unless approved by the [exception process](https://github.com/kubernetes/features/blob/master/EXCEPTIONS.md).
  * Ensures that cherry-picks do not destabilize the branch by either giving the PR enough time to stabilize in master or giving it enough time to stabilize in the release branch before cutting the release.
* Cuts the actual [release](https://github.com/kubernetes/kubernetes/releases).

#### Docs Lead
* Sets docs related deadlines for developers and works with Release Management Team Lead to ensure they are widely communicated.
* Sets up release branch for docs.
* Pings feature owners to ensure that release docs are created on time.
* Reviews/merges release doc PRs.
* Merges the docs release branch to master to make release docs live as soon as the release is official.

#### Features Lead
* Compiles the major themes, new features, known issues, actions required, notable changes to existing behavior, deprecations, etc. and edits them into a release doc checked in to the feature repository (ready to go out with the release).
* Collects and prepares the release notes

#### Bug Triage Lead
* Figures out which bugs (whether manually created or automatically generated) should be tracked for the release.
* Ensures all bugs being tracked for the release have owners that are responsive.
* Ensures all bugs are triaged as blocking or non-blocking.
* Ensures all bugs that are blocking are being actively worked on, esp after code complete.

#### Test Infra Lead
* Sets up and maintains all CI for the release branch.

#### Automated Upgrade Testing Lead
* Ensures that automated upgrade tests provide a clear go/no-go signal for the release.
* Tracks and finds owners for all issues with automated upgrade tests.

#### Manual Upgrade Testing Lead
* Ensures that any gaps in automated upgrade testing are covered by manual upgrade testing.
* Organizes the manual upgrade testing efforts, including setting up instructions for manual testing, finding manual testing volunteers, and ensuring any issues discovered are communicated widely and fixed quickly.

#### Testing Lead
* Ensures that all non-upgrade test CI provides a clear go/no-go signal for the release.
* Tracks and finds owners to fix any issues with any (non-upgrade) tests.

## Individual Contributors

Release responsiblites of indvidual contributors to the Kubernetes project are captured below.

### Patch Release

#### Cherry Picks
If you have a patch that needs to be ported back to a previous release (meaning it is a critical bug/security fix), once it is merged to the Kubernetes `master` branch:
* Mark your PR with the milestone corresponding to the release you want to port back to (e.g. `v1.5`), and add the `cherrypick-candidate` label to it.
* The Patch Release Manager will then review the PR and if it is ok for cherry-picking, will apply a `cherrypick-approved` label to it.
* Once your PR has been marked with the `cherrypick-approved` label by the Patch Release Manager, you should prepare a cherry-pick to the requested branch following the instructions [here](https://github.com/kubernetes/community/blob/master/contributors/devel/cherry-picks.md#how-do-cherrypick-candidates-make-it-to-the-release-branch).

### Major/Minor Release

#### Propose and Track Feature
If you are developing a feature for Kubernetes, make sure that an issue is open in the [features repository](https://github.com/kubernetes/features/issues). If you are targeting a particular release, make sure the issue is marked with the corresponding release milestone.

Ensure that all code for your feature is written, tested, reviewed, and merged before code freeze date for the target release.

During the code freeze period, fix any bugs discovered with you feature, and write feature documentation.

##### Writing Feature Documentation

1. Make sure your feature for the upcoming release is on the release tracking board (e.g. [link](https://docs.google.com/spreadsheets/d/1AFksRDgAt6BGA3OjRNIiO3IyKmA-GU7CXaxbihy48ns/edit?usp=sharing) for 1.8).
2. Create a PR with documentation for your feature in the [documents repo](https://github.com/kubernetes/kubernetes.github.io).
    * **Your PR should target the release branch (e.g. [`release-1.8`](https://github.com/kubernetes/kubernetes.github.io/tree/release-1.8)), not the [`master`](https://github.com/kubernetes/kubernetes.github.io/tree/master) branch.**
      * Any changes to the master branch become live on https://kubernetes.io/docs/ as soon as they are merged, and for releases we do not want docuemntation to go live until the release is cut.
3. Add link to your docs PR in the release tracking board, and notify the docs lead for the release (e.g. [Steve Perry](https://github.com/steveperry-53) for 1.8).
4. The docs lead will review your PR and give you feedback.
5. Once approved, the docs lead will merge your PR into the release branch.
6. When the release is cut, the docs lead will push the docs release branch to master, making your docs live on https://kubernetes.io/docs/.
