---
kep-number: 1a
title: Meta KEP Implementation
authors:
  - "@justaugustus"
  - "@calebamiles"
  - "@jdumars"
owning-sig: sig-pm
participating-sigs:
  - sig-architecture
reviewers:
  - "@erictune"
  - "@idvoretskyi"
  - "@spiffxp"
  - "@timothysc"
approvers:
  - "@bgrant0607"
  - "@jbeda"
editor:
creation-date: 2018-09-08
last-updated: 2018-09-29
status: implementable
see-also:
  - KEP-1
replaces:
superseded-by:
---

# Meta KEP Implementation

## Table of Contents

- [Table of Contents](#table-of-contents)
- [Summary](#summary)
- [Motivation](#motivation)
  - [Why another KEP?](#why-another-kep)
  - [Non-Goals](#non-goals)
- [Proposal](#proposal)
  - [Implementation Details / Notes / Constraints](#implementation-details--notes--constraints)
    - [Define](#define)
    - [Organize](#organize)
    - [Visibility and Automation](#visibility-and-automation)
    - [Constraints](#constraints)
- [Graduation Criteria](#graduation-criteria)
- [Implementation History](#implementation-history)

## Summary

Drive KEP adoption through improved process, documentation, visibility, and automation.

## Motivation

The KEP process is the standardized structure for proposing changes to the Kubernetes project.

In order to graduate KEPs to GA, we must iterate over the implementation details.

This KEP seeks to define actionable / delegable items to move the process forward.

Finally, by submitting a KEP, we gain an opportunity to dogfood the process and further identify areas for improvement.

### Why another KEP?

As [KEP-1] is currently de facto for the project, we must be careful to make changes to it in an iterative and atomic fashion.

When proposing a KEP, we action over some unit of work, usually some area of code.

In this instance, we treat [KEP-1] as the unit of work. That said, this would be considered a meta-KEP of the meta-KEP.

### Non-Goals

- API Review process
- Feature request triage
- Developer guide

## Proposal

### Implementation Details / Notes / Constraints

#### Define

- Refine existing KEP documentation
- Define KEP [DACI]
- Glossary of terms (enhancement, KEP, feature, etc.)
- KEP Workflow
  - KEP states
  - Entry / exit criteria

#### Organize

- Move KEPs from flat-files to a directory structure:
```
├── keps # top-level directory
│   ├── sig-beard # SIG directory
|   |   ├── 9000-beard-implementation-api # KEP directory
|   |   |   ├── kep.md (required) # KEP (multi-release work)
|   |   |   ├── experience_reports (required) # user feedback
|   |   |   │   ├── alpha-feedback.md
|   |   |   │   └── beta-feedback.md
|   |   |   ├── features # units of work that span approximately one release cycle
|   |   |   │   ├── feature-01.md
|   |   |   │   ├── ...
|   |   |   │   └── feature-n.md
|   |   |   ├── guides
|   |   |   |   ├── guide-for-developers.md (required)
|   |   |   |   ├── guide-for-teachers.md (required)
|   |   |   |   ├── guide-for-operators.md (required)
|   |   |   |   └── guide-for-project-maintainers.md
|   |   |   ├── index.json (required) # used for site generation e.g., Hugo
|   |   |   ├── metadata.yaml (required) # used for automation / project tracking
|   |   └── └── OWNERS
│   ├── sig-foo
|   ├── ...
|   └── sig-n
```

metadata.yaml would contain information that was previously in a KEP's YAML front-matter:

```
---
authors: # required
  - "calebamiles" # just a GitHub handle for now
  - "jbeda"
title: "Kubernetes Enhancement Proposal process"
number: 42 # required
owning-sig: "sig-pm" # required
participating-sigs:
  - "sig-architecture"
  - "sig-contributor-experience"
approvers: # required
  - "bgrant0607" # just a GitHub handle for now
reviewers:
  - "justaugustus"  # just a GitHub handle for now
  - "jdumars"
editors:
  - null # generally omit empty/null fields
status: "active" # required
github:
  issues:
    - null # GitHub url
  pull_requests:
    - null # GitHub url
  projects:
    - project_id: null
      card_id: null
releases: # required
  - k8s_version: v1.9
    kep_status: "active"
    k8s_status: "alpha" # one of alpha|beta|GA
  - k8s_version: v1.10
    kep_status: "active"
    k8s_status: "alpha"
replaces:
  - kep_location: null
superseded-by:
  - kep_location: null
created: 2018-01-22 # in YYYY-MM-DD
updated: 2018-09-04
```

- Move existing KEPs into [k/features]
- Create a `kind/kep` label for [k/community] and [k/features]
  - For `k/community`:
    - Label incoming KEPs as `kind/kep`
    - Enable searches of `org:kubernetes label:kind/kep`, so we can identify active PRs to `k/community` and reroute the PR authors to `k/enhancements` (depending on the state)
  - For `k/enhancements` (fka `k/features`):
    - Label incoming KEPs as `kind/kep`
    - Classify KEP submissions / tracking issues as `kind/kep`, differentiating them from `kind/feature`
- Move existing design proposals into [k/features]
- Move existing architectural documents into [k/features] (process TBD)
- Deprecate design proposals
- Rename [k/features] to [k/enhancements]
- Create tombstones / redirects to [k/enhancements]
- Prevent new KEPs and design proposals from landing in [k/community]
- Remove `kind/kep` from [k/community] once KEP migration is complete
- Correlate existing Feature tracking issues with links to KEPs
- Fix [KEP numbering races] by using the GitHub issue number of the KEP tracking issue
- Coordination of existing KEPs to use new directory structure (with SIG PM guidance per SIG)

#### Visibility and Automation

- Create tooling to:
  - Generate KEP directories and associated metadata
  - Present KEPs, through some easy to use mechanism e.g., https://enhancements.k8s.io. This would be a redesigned version of https://contributor.kubernetes.io/keps/. We envision this site / repo having at least three directories:
    - `keps/` (KEPs)
    - `design-proposals/` (historical design proposals from https://git.k8s.io/community/contributors/design-proposals)
    - `arch[itecture]|design/` (design principles of Kubernetes, derived from reorganizing https://git.k8s.io/community/contributors/devel, mentioned [here](https://github.com/kubernetes/community/issues/2565#issuecomment-419185591))
  - Enable project tracking across SIGs

#### Constraints

- Preserve git history
- Preserve issues
- Preserve PRs

## Graduation Criteria

Throughout implementation, we will be reaching out across the project to SIG leadership, approvers, and reviewers to capture feedback.

While graduation criteria has not strictly been defined at this stage, we will define it in future updates to this KEP.

## Implementation History

- 2018-08-20: (@timothysc) Issue filed about repo separation: https://github.com/kubernetes/community/issues/2565
- 2018-08-30: SIG Architecture meeting mentioning the need for a clearer KEP process - https://youtu.be/MMJ-zAR_GbI
- 2018-09-06: SIG Architecture meeting agreeing to move forward with a KEP process improvement effort to be co-led with SIG PM (@justaugustus / @jdumars) - https://youtu.be/fmlXkN4DJy0
- 2018-09-10: KEP-1a submitted for review
- 2018-09-25: Rationale discussion in SIG PM meeting
- 2018-09-28: Merged as `provisional`
- 2018-09-29: KEP implementation started
- 2018-09-29: [KEP Implementation Tracking issue] created
- 2018-09-29: [KEP Implementation Tracking board] created
- 2018-09-29: Submitted as `implementable`

[DACI]: https://www.atlassian.com/team-playbook/plays/daci
[KEP-1]: 0001-kubernetes-enhancement-proposal-process.md
[KEP Implementation Tracking board]: https://github.com/orgs/kubernetes/projects/5
[KEP Implementation Tracking issue]: https://github.com/kubernetes/features/issues/617
[KEP numbering races]: https://github.com/kubernetes/community/issues/2245
[k/community]: http://git.k8s.io/community
[k/enhancements]: http://git.k8s.io/enhancements
[k/features]: http://git.k8s.io/features
