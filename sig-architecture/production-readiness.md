# Production Readiness Review Process

Production readiness reviews are intended to ensure that features merging into
Kubernetes are observable, scalable and supportable, can be safely operated in
production environments, and can be disabled or rolled back in the event they
cause increased failures in production.

## Status

The process and questoinnaire are currently under development as part of the
[PRR KEP](https://github.com/kubernetes/enhancements/blob/master/keps/sig-architecture/20190731-production-readiness-review-process.md), with a target that reviews will be needed for features
going into 1.18.

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
