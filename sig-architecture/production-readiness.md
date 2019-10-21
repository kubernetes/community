# Production Readiness Review Process

Production readiness reviews are intended to ensure that features merging into
Kubernetes are observable, scalable and supportable, can be safely operated in
production environments, and can be disabled or rolled back in the event they
cause increased failures in production.

## Status

The process and questoinnaire are currently under development as part of the
[PRR KEP][], with a target that reviews will be needed for features going into 1.18.

During the 1.17 cycle, the PRR team will be piloting the questionnaire and other
aspects of the process.

## Questionnaire

* Feature enablement and rollback
  - How can this feature be enabled / disabled in a live cluster?
  - Can the feature be disabled once it has been enabled (i.e., can we roll
    back the enablement)?
  - Will enabling / disabling the feature require downtime for the control
    plane?
  - Will enabling / disabling the feature require downtime or reprovisioning
    of a node?
  - What happens if a cluster with this feature enabled is rolled back? What
    happens if it is subsequently upgraded again?
  - Are there tests for this?
* Scalability
  - Will enabling / using the feature result in any new API calls?
    Describe them with their impact keeping in mind the [supported limits][]
    (e.g. 5000 nodes per cluster, 100 pods/s churn) focusing mostly on:
     - components listing and/or watching resources they didn't before
     - API calls that may be triggered by changes of some Kubernetes
       resources (e.g. update object X based on changes of object Y)
     - periodic API calls to reconcile state (e.g. periodic fetching state,
       heartbeats, leader election, etc.)
  - Will enabling / using the feature result in supporting new API types?
    How many objects of that type will be supported (and how that translates
    to limitations for users)?
  - Will enabling / using the feature result in increasing size or count
    of the existing API objects?
  - Will enabling / using the feature result in increasing time taken
    by any operations covered by [existing SLIs/SLOs][] (e.g. by adding
    additional work, introducing new steps in between, etc.)?
    Please describe the details if so.
  - Will enabling / using the feature result in non-negligible increase
    of resource usage (CPU, RAM, disk IO, ...) in any components?
    Things to keep in mind include: additional in-memory state, additional
    non-trivial computations, excessive access to disks (including increased
    log volume), significant amount of data sent and/or received over
    network, etc. Think through this in both small and large cases, again
    with respect to the [supported limits][].
* Rollout, Upgrade, and Rollback Planning
* Dependencies
  - Does this feature depend on any specific services running in the cluster
    (e.g., a metrics service)?
  - How does this feature respond to complete failures of the services on
    which it depends?
  - How does this feature respond to degraded performance or high error rates
    from services on which it depends?
* Monitoring requirements
  - How can an operator determine if the feature is in use by workloads?
  - How can an operator determine if the feature is functioning properly?
  - What are the service level indicators an operator can use to determine the
    health of the service?
  - What are reasonable service level objectives for the feature?
* Troubleshooting
  - What are the known failure modes?
  - How can those be detected via metrics or logs?
  - What are the mitigations for each of those failure modes?
  - What are the most useful log messages and what logging levels do they require?
  - What steps should be taken if SLOs are not being met to determine the
    problem?

[PRR KEP]: https://github.com/kubernetes/enhancements/blob/master/keps/sig-architecture/20190731-production-readiness-review-process.md
[supported limits]: https://github.com/kubernetes/community/blob/master/sig-scalability/configs-and-limits/thresholds.md
[existing SLIs/SLOs]: https://github.com/kubernetes/community/blob/master/sig-scalability/slos/slos.md#kubernetes-slisslos
