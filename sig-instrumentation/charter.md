# SIG Instrumentation Charter

This charter adheres to the conventions described in the [Kubernetes Charter README] and uses
the Roles and Organization Management outlined in [sig-governance].

## Scope

Owns best practices for cluster observability through metrics and logging
across all Kubernetes components and development of components required for all
Kubernetes clusters (eg. heapster, metrics-server).

### In scope

SIG-Instrumentation revolves around the process of instrumenting and exposing
observability signals.

The act of instrumenting components not owned by SIG-Instrumentation is out of
scope, however SIG-Instrumentation is there to advise any contributors with
instrumentation decisions.

As well as giving advice in regards to instrumentation, SIG-Instrumentation
coordinates metric requirements of different SIGs for other components through
finding common APIs (such as the resource/core, custom and external metrics
APIs).


#### Code, Binaries and Services

[List of subprojects](https://github.com/kubernetes/community/tree/master/sig-instrumentation#subprojects)

- Components required for any Kubernetes cluster in regards to observability. Also referred to as the [core metrics pipeline][core-metrics-pipeline], meaning metrics that are to be consumed by the scheduler, kubectl and autoscaling. ([kubernetes-incubator/metrics-server](https://github.com/kubernetes-incubator/metrics-server), [kubernetes/heapster](https://github.com/kubernetes/heapster))
- Interfaces/API definitions required for any Kubernetes cluster in regards to observability. These the APIs defined in order to interface external system (such as Prometheus, Stackdriver, etc.) to be exposed to Kubernetes as a common interface, in order for Kubernetes to be able to treat metric sources as a generic metrics API. ([kubernetes/metrics](https://github.com/kubernetes/metrics), [kubernetes-incubator/custom-metrics-apiserver](https://github.com/kubernetes-incubator/custom-metrics-apiserver))
- Well established but optional components or adapters for Kubernetes clusters, if endorsed by members. Each component must have two or more members as maintainers. ([kubernetes/kube-state-metrics](https://github.com/kubernetes/kube-state-metrics), not yet officially owned by SIG-Instrumentation, but an example prospect for this category: [DirectXMan12/k8s-prometheus-adapter](https://github.com/DirectXMan12/k8s-prometheus-adapter))

#### Cross-cutting and Externally Facing Processes

- Guidance for instrumentation in order to ensure consistent and high quality instrumentation of core Kubernetes components. This includes:
  - Reviewing any instrumentation related changes and additions.
  - Guidance on what should be instrumented as well as dimensions of the same. (see [Instrumenting Kubernetes Guide][instrumenting-kubernetes])
  - Creating, adding and maintaining the Kubernetes instrumentation guidelines.
  - Coordinate cross SIG-Instrumentation efforts.
  - The interface of log files and their directory structure written out by container runtimes to be processed by other systems further, is shared responsibility between [SIG Node](sig-node) and SIG Instrumentation.

### Out of scope

- Processing of signals. For example ingesting metrics, logs, events into external systems.
- Dictating what states must result in an alert. Suggestions or opt-in alerts may be in scope.
- Cloud provider specific addons are out of scope and should be taken care of by the respective SIG.
- The act of writing log files, their format, or how they are to be processed afterwards.

## Roles and Organization Management

This SIG follows adheres to the Roles and Organization Management outlined in [sig-governance]
and opts-in to updates and modifications to [sig-governance].

### Additional responsibilities of Tech Leads

- No additional responsibilities of Tech Leads

### Deviations from [sig-governance]

Tech Leads must also fulfill all of the responsibilities of the Chair role as outlined in [sig-governance].

### Subproject Creation

By SIG Technical Leads

[sig-governance]: https://github.com/kubernetes/community/blob/master/committee-steering/governance/sig-governance.md
[sig-node]: https://github.com/kubernetes/community/tree/master/sig-node
[sigs.yaml]: https://github.com/kubernetes/community/blob/master/sigs.yaml#L964-L1018
[Kubernetes Charter README]: https://github.com/kubernetes/community/blob/master/committee-steering/governance/README.md
[instrumenting-kubernetes]: /contributors/devel/sig-instrumentation/instrumentation.md
[core-metrics-pipeline]: https://kubernetes.io/docs/tasks/debug-application-cluster/resource-metrics-pipeline/
