# Kubernetes Repository Guidelines

This document attempts to outline a structure for creating and associating
GitHub repositories with the Kubernetes project. It also describes how and when
repositories are removed.

The document presents a tiered system of repositories with increasingly strict
requirements in an attempt to provide the right level of oversight and
flexibility for a variety of different projects.

Requests for creating, transferring, modifying, or archiving repositories can be
made by [opening a request](https://github.com/kubernetes/org/issues/new/choose)
against the kubernetes/org repo.

- [Associated Repositories](#associated-repositories)
  * [Goals](#goals)
  * [Rules](#rules)
- [SIG repositories](#sig-repositories)
  * [Goals](#goals-1)
  * [Rules for new repositories](#rules-for-new-repositories)
  * [Rules for donated repositories](#rules-for-donated-repositories)
- [Core Repositories](#core-repositories)
  * [Goals](#goals-2)
  * [Rules](#rules-1)
- [Maintenance Mode](#maintenance-mode)
  * [Roadmap](#roadmap)
  * [How-tos](#how-tos)
- [Removing Repositories](#removing-repositories)
  * [Grounds for removal](#grounds-for-removal)
- [FAQ](#faq)

## Associated Repositories

Associated repositories conform to the Kubernetes community standards for a
repository, but otherwise have no restrictions. Associated repositories exist
solely for the purpose of making it easier for the Kubernetes community to work
together. There is no implication of support or endorsement of any kind by the
Kubernetes project, the goals are purely logistical.

### Goals

To facilitate contributions and collaboration from the broader Kubernetes
community.  Contributions to random projects with random CLAs (or DCOs) can be
logistically difficult, so associated repositories should be easier.


### Rules

   * Must adopt the Kubernetes Code of Conduct statement in their repo.
   * All code projects use the Apache License version 2.0. Documentation
     repositories must use the Creative Commons License version 4.0.
   * Must adopt the CNCF CLA bot automation for pull requests.


## SIG repositories

SIG repositories serve as temporary homes for SIG-sponsored experimental
projects or prototypes of new core functionality, or as permanent homes for
SIG-specific projects and tools.

### Goals

To provide a place for SIGs to collaborate on projects endorsed by and actively
worked on by members of the SIG. SIGs should be able to approve and create new
repositories for SIG-sponsored projects without requiring higher level approval
from a central body (e.g. steering committee or sig-architecture)

### Rules for new repositories

   * For now all repos will live in `github.com/kubernetes-sigs/\<project-name\>`.
      * For project names: see [project-naming](./../committee-steering/governance/project-naming.md)
   * Must contain the topic for the sponsoring SIG - e.g.
     `k8s-sig-api-machinery`.  (Added through the *Manage topics* link on the
repo page.)
   * Must adopt the Kubernetes Code of Conduct
   * All code projects use the Apache License version 2.0. Documentation
     repositories must use the Creative Commons License version 4.0.
   * Must adopt the CNCF CLA bot, merge bot and Kubernetes PR commands/bots.
   * All OWNERS of the project must also be active SIG members.
   * Must be approved by the process spelled out in the SIG's charter and a
   publicly linkable written decision should be available for the same.
   * SIG must already have identified all of their existing subprojects and
     code, with valid OWNERS files, in
[`sigs.yaml`](https://github.com/kubernetes/community/blob/master/sigs.yaml)

### Rules for donated repositories

The `kubernetes-sigs` organization is primarily intended to house net-new
projects originally created in that organization. However, projects that a SIG
adopts may also be donated.

In addition to the requirements for new repositories, donated repositories must
demonstrate that:

   * All contributors must have signed the [CNCF Individual
     CLA](https://github.com/cncf/cla/blob/master/individual-cla.pdf) or [CNCF
Corporate CLA](https://github.com/cncf/cla/blob/master/corporate-cla.pdf)
   * If (a) contributor(s) have not signed the CLA and could not be reached, a
     NOTICE file should be added referencing section 7 of the CLA with a list of
the developers who could not be reached
   * Licenses of dependencies are acceptable; please review the [allowed-third-party-license-policy.md](https://github.com/cncf/foundation/blob/main/policies-guidance/allowed-third-party-license-policy.md)
     and [exceptions](https://github.com/cncf/foundation/tree/main/license-exceptions).
     If your dependencies are not covered, then please open a `License Exception Request` issue in
     [cncf/foundation](https://github.com/cncf/foundation/issues) repository.
   * Boilerplate text across all files should attribute copyright as follows:
     `"Copyright <Project Authors>"` if no CLA was in place prior to donation
   * Additions of [the standard Kubernetes header](https://git.k8s.io/kubernetes/hack/boilerplate/boilerplate.go.txt)
     to code created by the contributors can occur post-transfer, but should
     ideally occur shortly thereafter.
   * Should contain template files as per the
     [kubernetes-template-project](https://github.com/kubernetes/kubernetes-template-project).
   * For project names: see [project-naming](./../committee-steering/governance/project-naming.md)

Note that copyright notices should only be modified or removed by the people or
organizations named in the notice. See [the FAQ below](#faq) for more information
regarding copyrights and copyright notices.

#### Additional requirements for forking external repositories

In addition to all standard requirements for donated repositories:

* When forking an external repository into the Kubernetes GitHub organization (particularly if the repository owners cannot be contacted), a NOTICE file must be included listing all existing contributors to the external repository. An example of the NOTICE file is [here](https://github.com/kubernetes-sigs/randfill/blob/ba7cb247249023527ba6c4d05aeeb8f8faf22c74/NOTICE).

This is to ensures proper attribution and compliance with the CLA.

## Core Repositories

Core repositories are considered core components of Kubernetes. They are
utilities, tools, applications, or libraries that are expected to be present in
every or nearly every Kubernetes cluster, such as components and tools included
in official Kubernetes releases. Additionally, the kubernetes.io website, k8s.io
machinery, and other project-wide infrastructure will remain in the kubernetes
github organization.

### Goals

Create a broader base of repositories than the existing
gh/kubernetes/kubernetes so that the project can scale. Present expectations
about the centrality and importance of the repository in the Kubernetes
ecosystem. Carries the endorsement of the Kubernetes community.

### Rules

   * Must live under `github.com/kubernetes/<project-name>`
   * Must adopt the Kubernetes Code of Conduct
   * All code projects use the Apache Licence version 2.0. Documentation
     repositories must use the Creative Commons License version 4.0.
   * Must adopt the CNCF CLA bot
   * Must adopt all Kubernetes automation (e.g. /lgtm, etc)
   * All OWNERS must be members of standing as defined by ability to vote in
     Kubernetes steering committee elections. in the Kubernetes community
   * Repository must be approved by SIG-Architecture

## Maintenance Mode

Projects that are considered "done" or not pursuing the development of new features but are relied on as a dependency can be considered in Maintenance mode.

When in maintenance mode, the community can expect the following from the owners
of the project/repository:

  * No concrete plan on introduction of new features
  * Incoming Issues and PR(s) will not be looked at in any regular cadence
  * Minimal upkeep for project language and  dependency updates
  * Security-related features/updates to be taken care of
  * Explicit guidance around removal/transition to active state

### Process for transitioning to maintenance mode

  * SIG Chairs or TLs can open PRs to add the label and a preamble in the README.md
 for the project.
  * In addition, a notice will be sent out by the SIG Chairs or TLs to the [#k-dev]() 
  mailing list with a fortnight for lazy consensus once the PR is opened.

### Rules

 * Steering committee liaisons can also recommend a project/repository to transition to
 maintenance mode during the annual reporting process.
 * When in doubt, SIG-Architecture will be the decision-maker.
 * Specific members who can help with the process will be identified & added to the
 OWNERS file for performing the required activities. 

## Removing Repositories

As important as it is to add new repositories, it is equally important to prune
old repositories that are no longer relevant or useful.

It is in the best interests of everyone involved in the Kubernetes community
that our various projects and repositories are active and healthy. This ensures
that repositories are kept up to date with the latest Kubernetes wide processes,
it ensures a rapid response to potential required fixes (e.g. critical security
problems) and (most importantly) it ensures that contributors and users receive
quick feedback on their issues and contributions.

### Grounds for removal

SIG repositories and core repositories may be removed from the project if they
are deemed _inactive_. Inactive repositories are those that meet any of the
following criteria:

   * There are no longer any active maintainers for the project and no
     replacements can be found.
   * All PRs or Issues have gone un-addressed for longer than six months.
   * There have been no new commits or other changes in more than a year.
   * The contents have been folded into another actively maintained project.

Associated repositories are much more loosely associated with the Kubernetes
project and are generally not subject to removal, except under exceptional
circumstances (e.g. a code of conduct violation).

## FAQ

**My project is currently in core, but doesn’t seem to fit these guidelines,
what’s going to happen?**

For now, nothing. Eventually, we may redistribute projects, but for now the goal
is to adapt the process going forward, not re-legislate past decisions.

**I’m starting a new project, what should I do?**

Is this a SIG-sponsored project? If so, convince some SIG to host it, take it to
the SIG mailing list, meeting and get consensus, then the SIG can create a repo
for you in the SIG organization.

Is this a small-group or personal project? If so, create a repository wherever
you’d like, and make it an associated project.

We suggest starting with the kubernetes-template-project to ensure you have the
correct code of conduct, license, etc.

**Much of the things needed (e.g. CLA Bot integration) is missing to support
associated projects. Many things seem vague. Help!**

True, we need to improve these things. For now, do the best you can to conform
to the spirit of the proposal (e.g. post the code of conduct, etc)

**When I donate my project, am I transferring my copyrights?**

No. All contributors retain ownership of their copyrights in the code they donate.
Instead, they are granting a license to the project (that's the 'L' in 'CLA').

For consistency and efficiency in complying with notice requirements, code that is
donated to a Kubernetes repo should use [the standard header](https://git.k8s.io/kubernetes/hack/boilerplate/boilerplate.go.txt)
referencing "The Kubernetes Authors". That doesn't mean you are transferring your
copyright. Instead, it's a general reference to the fact that the copyrights remain
owned by the authors of Kubernetes.

Note that you should _never_ modify or remove a third party's copyright notice if
you are not authorized by them to do so.
