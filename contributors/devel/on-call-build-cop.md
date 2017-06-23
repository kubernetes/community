# Kubernetes BuildCop Workflow

June 2017

## Objective

This document describes the responsibilities and the workflow of a person assuming the buildcop role. 
The current buildcop can be found [here](https://storage.googleapis.com/kubernetes-jenkins/oncall.html).

## Prerequisites for build-copping

- Ensure you have admin access to [http://github.com/kubernetes/kubernetes](http://github.com/kubernetes/kubernetes)
  - Check your membership in the GitHub team: [kubernetes-build-cops](https://github.com/orgs/kubernetes/teams/kubernetes-build-cops/members). 
  If you are not a member contact one of the team maintainers to get yourself added to it.
  - Test your admin access by e.g. adding a label to an issue.
- You must communicate any concerns/actions via the **#sig-release** slack channel to ensure that 
the release team has context on the current state of the submit queue.
- You must attend the release burndown meeting to provide an update on the current state of the submit-queue

## Responsibilities

The build-cop's primary responsibility is to ensure that automatic merges are happening at a 
**reasonable** rate. This may include performing merging of test flake PRs when the pre-submits 
are failing repeatedly. The buildcop must be familiar with the 
[queue labels](https://submit-queue.k8s.io/#/info) and apply them as necessary to critical fixes. 
The priority labels are defunct and no longer respected by the submit-queue. As of June 2017, 
the merge rate is ~30 PRs per day if there are that many PRs in the queue. The previous 
responsibilities of this role included classification of incoming issues, but that is no 
longer a part of the mandate.

## Workflow

1. Check the Prow batch dashboard: [https://prow.k8s.io/?type=batch](https://prow.k8s.io/?type=batch) 
to ensure that batch jobs are running regularly. It's okay to see occasional flakes. Do not worry
about manually re-running individual tests, since Prow will rerun them.
2. If there are post-submit blocking jobs (see [link](https://submit-queue.k8s.io/#/e2e)), ensure 
that those builds are green and allowing merges to occur.
3. If several batch merges are failing, file an issue for that job and describe the possible 
causes for the failure. Debug if possible, else triage and assign to a particular SIG, and 
@-mention the maintainers. For example, see: 
[#47135](https://github.com/kubernetes/kubernetes/issues/47135)
4. Communicate the actions to **#sig-release** via slack and ensure that the issue is being worked on.
5. If the issue is not worked on for several hours, please escalate to the release team. 
  The release team members can be found via the [features](https://github.com/kubernetes/features) repo.
  For example, the Kubernetes 1.7 release team members are listed [here](https://github.com/kubernetes/features/blob/master/release-1.7/release_team.md).
  Notify the release manager/release team members via GitHub mentions and slack. 
6. When the SIG member sends a fix, manually merge if necessary, after verifying that pre-submits pass, 
or use the 'retest-not-required' label with the appropriate 'queue/*' label to ensure merge of the 
flake fix.
7. Issue an update to the **#sig-release** channel on the merge rate and the PR that was used to fix the queue.
