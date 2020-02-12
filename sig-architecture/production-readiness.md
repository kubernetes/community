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

#### Feature enablement and rollback

* **How can this feature be enabled / disabled in a live cluster?**
  - [ ] Feature gate
    - Feature gate name:
    - Components depending on the feature gate:
  - [ ] Other
    - Describe the mechanism:
    - Will enabling / disabling the feature require downtime of the control
      plane?
    - Will enabling / disabling the feature require downtime or reprovisioning
      of a node? (Do not assume `Dynamic Kubelet Config` feature is enabled).

* **Can the feature be disabled once it has been enabled (i.e. can we rollback
  the enablement)?**
  Describe the consequences on existing workloads (e.g. if this is runtime
  feature, can it break the existing applications?).

* **What happens if we reenable the feature if it was previously rolled back?**

* **Are there any tests for feature enablement/ disablement?**
  The e2e framework does not currently support enabling and disabling feature
  gates. However, unit tests in each component dealing with managing data created
  with and without the feature are necessary. At the very least, think about
  conversion tests if API types are being modified.

#### Scalability

* **Will enabling / using this feature result in any new API calls?**
  Describe them, providing:
  - API call type (e.g. PATCH pods)
  - estimated throughput
  - originating component(s) (e.g. Kubelet, Feature-X-controller)
  focusing mostly on:
  - components listing and/or watching resources they didn't before
  - API calls that may be triggered by changes of some Kubernetes resources
    (e.g. update of object X triggers new updates of object Y)
  - periodic API calls to reconcile state (e.g. periodic fetching state,
    heartbeats, leader election, etc.)

* **Will enabling / using this feature result in introducing new API types?**
  Describe them providing:
  - API type
  - Supported number of objects per cluster
  - Supported number of objects per namespace (for namespace-scoped objects)

* **Will enabling / using this feature result in any new calls to cloud
  provider?**

* **Will enabling / using this feature result in increasing size or count
  of the existing API objects?*
  Describe them providing:
  - API type(s):
  - Estimated increase in size: (e.g. new annotation of size 32B)
  - Estimated amount of new objects: (e.g. new Object X for every existing Pod)

* **Will enabling / using this feature result in increasing time taken by any
  operations covered by [existing SLIs/SLOs][]?**
  Think about adding additional work or introducing new steps in between
  (e.g. need to do X to start a container), etc. Please describe the details.

* **Will enabling / using this feature result in non-negligible increase of
  resource usage (CPU, RAM, disk, IO, ...) in any components?**
  Things to keep in mind include: additional in-memory state, additional
  non-trivial computations, excessive access to disks (including increased log
  volume), significant amount of data send and/or received over network, etc.
  This through this both in small and large cases, again with respect to the
  [supported limits][].

#### Rollout, Upgrade and Rollback Planning

#### Dependencies

* **Does this feature depend on any specific services running in the cluster?**
  Think about both cluster-level services (e.g. metrics-server) as well
  as node-level agents (e.g. specific version of CRI). Focus on external or
  optional services that are needed. For example, if this feature depends on
  a cloud provider API, or upon an external software-defined storage or network
  control plane.

* **How does this feature respond to complete failures of the services on which
  it depends?**
  Think about both running and newly created user workloads as well as
  cluster-level services (e.g. DNS).

* **How does this feature respond to degraded performance or high error rates
  from services on which it depends?**

#### Monitoring requirements

* **How can an operator determine if the feature is in use by workloads?**

* **How can an operator determine if the feature is functioning properly?**
  Focus on metrics that cluster operators may gather from different
  components and treat other signals as last resort.

* **What are the SLIs (Service Level Indicators) an operator can use to
  determine the health of the service?**
  - [ ] Metrics
    - Metric name:
    - [Optional] Aggregation method:
    - Components exposing the metric:
  - [ ] Other (treat as last resort)
    - Details:

* **What are the reasonable SLOs (Service Level Objectives) for the above SLIs?**

#### Troubleshooting
Troubleshooting section serves the `Playbook` role as of now. We may consider
splitting it into a dedicated `Playbook` document (potentially with some monitoring
details). For now we leave it here though, with some questions not required until
further stages (e.g. Beta/Ga) of feature lifecycle.

* **How does this feature react if the API server is unavailable?**

* **What are other known failure modes?**

* **How can those be detected via metrics or logs?**

* **What are the mitigations for each of those failure modes?**

* **What are the most useful log messages and what logging levels to they require?**
  Not required until feature graduates to Beta.

* **What steps should be taken if SLOs are not being met to determine the problem?**


[PRR KEP]: https://github.com/kubernetes/enhancements/blob/master/keps/sig-architecture/20190731-production-readiness-review-process.md
[supported limits]: https://github.com/kubernetes/community/blob/master/sig-scalability/configs-and-limits/thresholds.md
[existing SLIs/SLOs]: https://github.com/kubernetes/community/blob/master/sig-scalability/slos/slos.md#kubernetes-slisslos
