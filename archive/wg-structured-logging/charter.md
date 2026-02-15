# WG Structured Logging Charter

This charter adheres to the conventions described in the [Kubernetes Charter README]
and uses the Roles and Organization Management outlined in [wg-governance].

[Kubernetes Charter README]: https://github.com/kubernetes/community/blob/master/committee-steering/governance/README.md

## Scope

Modernize logging in Kubernetes core components, allowing users to efficiently consume, process, store and analyse 
information stored in logs.

### In Scope

- Define the standard for logging - propose libraries, interfaces, metadata schema
- Reduce friction for using logging - reduce dependencies and performance overhead
- Give more choice over logging - allow pluggable logging implementation
- Ensure quality consistent logging - overview migration, create documentation, tooling and educate reviewers
- Prevent regressions caused by logging - measure performance overhead and log volume changes

For all of the above, we will focus on core Kubernetes components and addons.
Other SIG subprojects/components (e.g. SIG Scheduling descheduler) are out of
scope.

### Out of scope

- Logging outside the kubernetes/kubernetes repository
- Non core Kubernetes component binaries like kubectl and kubeadm
- Application logs read by kubectl

## Stakeholders

Stakeholders in this working group span multiple SIGs that own parts of 
the code in core Kubernetes components and addons.

  - API Machinery
  - Architecture
  - Cloud Provider
  - Instrumentation
  - Network
  - Node
  - Scheduling
  - Storage

## Deliverables

The artifacts the group is supposed to deliver include:
- Completion of [Structured Logging migration]
- Graduation of JSON logging format to GA
- Documented guidelines on using Structured Logging
- Replacement for non-structured logging library (klog)
- Automated tooling to prevent regressions caused by logging

[Structured Logging migration]: https://github.com/kubernetes/enhancements/tree/master/keps/sig-instrumentation/1602-structured-logging

## Roles and Organization Management

This wg follows adheres to the Roles and Organization Management outlined in
[wg-governance] and opts-in to updates and modifications to [wg-governance].

[wg-governance]: https://github.com/kubernetes/community/blob/master/committee-steering/governance/wg-governance.md

## Timelines and Disbanding

The exact timeline for existing of this working group is hard to predict at
this time.

The group will start working on the deliverables mentioned above. Once the
group we will be satisfied with the current shape of them and no additional
coordination on their execution will be needed, we will retire Working Group
and pass oversight of logging to SIG Instrumentation.
