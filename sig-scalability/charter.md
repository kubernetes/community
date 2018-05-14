# SIG Scalability Charter

## Mission
The SIG Scalability helps to define scalability goals for Kubernetes, ensures
that they all play well together and ensures that every Kubernetes release meets
them by measuring performance and scalability indicators and publishing the
results.

We also coordinate and contribute to general system-wide scalability and
performance improvements (that don’t fall into the charter of another individual
SIG) as well as provide consultations about any scalability and performance
related aspects of Kubernetes.

## What can we do/require from other SIGs
Scalability and performance are horizontal aspects of the system - changes in a
single place of Kubernetes may affect the whole system. As a result, to
effectively ensure Kubernetes scales, we need a special cross-SIG privileges.

- We can rollback any merged PR if it has been identified as a cause of any
  [performance/scalability SLOs] regression. The offending PR should only be
  merged again after proving to pass  tests at scale.
- We can pause the merge queue in case of a regression observed until a particular
  PR has been identified as cause of the regression and regression has been
  mitigated. The “Rules of engagement” of pausing merge-queue and rationale for
  necessity of its introduce are explained in a separate doc. <br/>
  TODO(wojtek-t, shyamjvs): Write it down and link here.
- We require significant changes (in terms of impact, such as: update of etcd,
  update of Go version, major architectural changes, etc.) may only be merged:
  - with an explicit approval from a [SIG-scalability approver](#sig-scalability-approvers)
    and
  - after having passed performance testing on biggest supported clusters (unless
    found unnecessary by scalability approver)
- We can block a feature from transitioning to Beta status if (when turned on) it
  causes a significant degradation of overall Kubernetes scalability/performance.
  (Ideally it would be “SLI degradation of more than X%” or “breaking SLO”, but
  initially it may also be SIG-scalability decision based on public test results).
- We can block a feature from transitioning to GA status if it cannot be used at
  scale.
- We can require a SIG to introduce a regression-catching benchmark test for a
  scalability-critical functionality.

For the record, by regression above we mean a regression identified by the set 
of release-blocking scalability/performance tests (as defined by
sig-release-master-blocking group of test suites).

[performance/scalability SLOs]: https://github.com/kubernetes/community/blob/master/sig-scalability/slos/slos.md

## SIG Values

- We are NOT firefighters, we are fire-prevention specialists.
- We promote deep technical understanding of the Kubernetes system and our tools.
- We strive to eliminate toil.
- We work towards building a scalable Kubernetes even in face of superlinear growth
  of number of contributors.

## Scope and subprojects
The scope of SIG Scalability covers all aspects of Kubernetes scalability and
performance. However, all issues that fully fall under a single SIG are implicitly
delegated to that SIG.

SIG scalability subprojects are as follows.

| Subproject | Description | Example Artifacts | OWNERS |
| --- | --- | --- | --- |
| Kubernetes scalability | Defining what does it mean that “Kubernetes scales”. This includes defining (or approving) individual performance SLIs/SLOs, ensuring they are all oriented on user experience and consistent with each other. | [SLIs/SLOs] | [OWNERS](https://github.com/kubernetes/community/blob/master/sig-scalability/slos/OWNERS) |
| Kubernetes performance validation | Ensuring that each official Kubernetes release satisfies all scalability and performance related requirements, as state in “Kubernetes scalability” definition | [1.9 validation report] | TODO |
| Scalability testing frameworks | Designing and creating frameworks to make scalability and performance testing of Kubernetes easy and available for all contributors. Different frameworks may help in different aspect of scalability testing enabling making conscious tradeoffs, e.g. cost of accuracy or real life vs more generalized benchmarking scenarios. | [Cluster loader] | [OWNERS](https://github.com/kubernetes/perf-tests/blob/master/OWNERS) [OWNERS](https://github.com/kubernetes/kubernetes/blob/master/test/kubemark/OWNERS) |
| Scalability and performance tests | Ensuring that all tests necessary to validate Kubernetes scalability and performance exist (ideally by providing easy-to-use framework and working with SIGs to provide them) have the environment and resources to run on and are being executed according to calendar enabling release validation. | [Scalability e2e tests] | [OWNERS](https://github.com/kubernetes/kubernetes/blob/master/test/e2e/scalability/OWNERS) |
| Scalability governance | Establishing and documenting best practises on how to design and/or implement Kubernetes features in scalable and performant way. Educating contributors and ensuring those are widely used. | [Regressions case study] | [OWNERS](https://github.com/kubernetes/community/blob/master/sig-scalability/governance/OWNERS) |

TODO: Figure out if we need subproject for finding bottlenecks, coordinating
improvements and architectural changes, etc.

[SLIs/SLOs]: https://github.com/kubernetes/community/blob/master/sig-scalability/slos/slos.md
[1.9 validation report]: https://github.com/kubernetes/sig-release/blob/master/releases/release-1.9/scalability_validation_report.md
[Cluster loader]: https://github.com/kubernetes/perf-tests/tree/master/clusterloader
[Scalability e2e tests]: https://github.com/kubernetes/kubernetes/tree/master/test/e2e/scalability
[Regressions case study]: https://github.com/kubernetes/community/blob/master/sig-scalability/blogs/scalability-regressions-case-studies.md

## Roles
The following roles are required for the SIG to properly function.
In the event that any role is unfilled, the SIG will make a best effort
to fill it and any decisions reliant on a missing role will be postponed
until the role is filled.

### Chair
- Number: 2-3
- Run operations and processes governing the SIG
- A majority of chairs cannot be from a single company.
- An initial set of chairs was established at the time the SIG was founded as:
  Wojciech Tyczynski and Bob Wise.
- Chairs may decide to step down and propose a replacement, who must be approved
  by all other chairs.
- Chairs may select additional chairs by consensus.
- Chairs may be removed by consensus of other Chairs and Technical Leads if not
  proactively working with other Chairs to fulfill responsibilities.

### Technical Lead
- Number: 2-3
- Establish new subprojects and retire existing ones
- Resolve cross-subprojects technical issues and decisions and escalations from
  subprojects.
- Decision making must be by consensus.
- An initial set of technical leads was set to long-standing group of SIG leads:
  Wojciech Tyczynski and Bob Wise.
- Technical leads must have demonstrated deep understanding of the whole system
  that is sufficient to assess impact of different changes on Kubernetes scalability.
- Technical leads must remain active in the role and are automatically removed
  from the position if they are unresponsive for >3 months.
- Technical leads may decide to step down at anytime and propose a replacement,
  who must be approved by all of the other technical leads.
- TODO: Diversity across companies?

### Subproject owners
- Number: at least 2
- The initial owners should be established at subproject founding from relevant
  OWNERS file wherever possible.
- Owners must be an escalation point for technical discussions and decisions within
  the subproject.
- Owners must set milestone priorities for their subprojects.
- Owners must remain active in the role and are automatically removed from the
  position if they are unresponsive for >3 months and may be removed by consensus
	of the other subproject owners and all of the Technical loeads if not proactively
	working to fulfill responsibilities.
- Owners may decide to step down at any time and propose replacement. Accepting
  replacement will be done by lazy-consensus from other subproject owners.
- Owners may select additional subproject owners through a super-majority vote
  amongst subproject owners.

### SIG Scalability approvers
- Number: at least 3
- Approve significant changes (in terms of potential impact, e.g. major architectural
  changes, upgrades of etcd or Go version) from scalability perspective.
- An initial set of approvers was set to:
  - Bob Wise
  - Clayton Coleman
  - Jordan Liggitt
  - Shyam Jeedigunta
  - Wojciech Tyczynski

## Organizational management
- Six months after this charter is first ratified, it must be reviewed and
  re-approved by the SIG in order to evaluate the assumptions made in its initial
  drafting.
- SIG meets bi-weekly on zoom with agenda in meeting nodes and should be
  facilitated by chair unless delegated.

## Project management

### Subproject creation
The initial set of subprojects owned by the SIG is defined above.
- New subprojects must be approved by consensus of SIG Technical Leads.

### Subproject retirement
Subprojects may be retired, when they are no longer supported based on the
following criteria:
- A subproject is no longer supported when there are no active owners with
  activity on the project:
  - for >3 months for subprojects with no known users
  - for >6 months for subprojects with known users after providing at least
    6 months notification
- Consensus amongst Technical Leads should be done to decide about retirement.

### Technical processes
- Decisions within the scope of individual subprojects should be made by lazy
  consensus by subproject owners;  if a decision can’t be made, it should be
  escalated to the SIG Technical leads.
- Issues impacting multiple subprojects in the SIG should be resolved by
  consensus of the owners of the involved subprojects; if a decision can’t be
  made, it should be escalated to the SIG Technical leads.
