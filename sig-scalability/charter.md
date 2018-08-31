# SIG Scalability Charter

This charter adheres to the conventions described in the [Kubernetes Charter README]
and uses the Roles and Organization Management outlined in [sig-governance].

[sig-governance]: https://github.com/kubernetes/community/blob/master/committee-steering/governance/sig-governance.md
[Kubernetes Charter README]: https://github.com/kubernetes/community/blob/master/committee-steering/governance/README.md

## Scope

SIG Scalability's primary responsibilities are to define and drive scalability
goals for Kubernetes. This involves defining, testing and measuring performance and
scalability related Service Level Indicators (SLIs) and ensuring that every
Kubernetes release meets Service Level Objectives (SLOs) built on top of those
SLIs.

We also coordinate and contribute to general system-wide scalability and
performance improvements (that don't fall into the charter of another individual
SIG) by driving large architectural changes and finding bottlenecks, as well as
provide consultations about any scalability and performance related aspects of
Kubernetes.

### In Scope

#### Code, Binaries and Services:

- Scalability and performance testing frameworks. Examples include:
  - [Cluster loader](https://github.com/kubernetes/perf-tests/tree/master/clusterloader2)
  - [Kubemark](https://github.com/kubernetes/kubernetes/tree/master/cmd/kubemark)
- Scalability and performance tests:
  - [Tests](https://github.com/kubernetes/kubernetes/blob/master/test/e2e/scalability/)
  - [Jobs running those](https://github.com/kubernetes/test-infra/tree/master/config/jobs/kubernetes/sig-scalability)

#### Cross-cutting and Externally Facing Processes

- Defining what does “Kubernetes scales” mean by defining (or approving)
individual performance SLIs/SLOs, ensuring they are all oriented on user
experience and consistent with each other:
  - [SLIs/SLOs](https://github.com/kubernetes/community/blob/master/sig-scalability/slos/slos.md)
- Ensuring that each official Kubernetes release satisfies all scalability and
performance related requirements, as stated in "Kubernetes scalability" definition.
- Establishing and documenting best practises on how to design and/or implement
Kubernetes features in scalable and performant way. Educating contributors and
consulting individual designs/implementations to ensure that those are widely used.
Example artifacts:
  - [Scalability governance](https://github.com/kubernetes/community/blob/master/sig-scalability/governance)
- Finding system bottlenecks and coordinating improvement on cross-cutting
architectural changes.

### Out of scope

- Improving performance/scalability of features falling into charters of
individual SIGs.

## What can we do/require from other SIGs

Scalability and performance are horizontal aspects of the system - changes in a
single place of Kubernetes may affect the whole system. As a result, to
effectively ensure Kubernetes scales, we need a special cross-SIG privileges.

- We can rollback any merged PR if it has been identified as a cause of any
  [performance/scalability SLOs] regression (identified by the set of release
  blocking scalability/performance tests). The offending PR should only be
  merged again after proving to pass  tests at scale.
- In the event of a performance regression, we can block all PRs from being
  merged into the relevant repos until the cause of the regression is
  identified and mitigated.
  The “Rules of engagement” of pausing merge-queue and rationale for
  necessity of its introduce are explained in [a separate doc](./block_merges.md).
- We require significant changes (in terms of impact, such as: update of etcd,
  update of Go version, major architectural changes, etc.) may only be merged:
  - with an explicit approval from a SIG-scalability tech lead and
  - after having passed performance testing on biggest supported clusters (unless
    found unnecessary by approver)
- We can block a feature from transitioning:
  - to Beta status, if (when turned on) it causes violation of already existing
    performance/scalability SLOs;
  - to GA status, when it can be used scale. That means:
    - in rare cases, introducing a new SLI and SLO and ensuring it is met at scale
    - in most of cases, extending scalability tests to use it and ensuring that
      existing SLOs are still met
- We can require a SIG to introduce a regression-catching benchmark test for a
  scalability-critical functionality.

[performance/scalability SLOs]: https://github.com/kubernetes/community/blob/master/sig-scalability/slos/slos.md

## Roles and Organization Management

This sig follows adheres to the Roles and Organization Management outlined in
[sig-governance] and opts-in to updates and modifications to [sig-governance].

[sig-governance]: https://github.com/kubernetes/community/blob/master/committee-steering/governance/sig-governance.md

### Subproject Creation

SIG Scalability delegates subproject approval to Technical Leads. See [Subproject creation - Option 1].

[Subproject creation - Option 1]: https://github.com/kubernetes/community/blob/master/committee-steering/governance/sig-governance.md#subproject-creation
