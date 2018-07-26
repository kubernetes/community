# SIG Autoscaling Charter

This charter adheres to the conventions described in the [Kubernetes Charter README] and uses
the Roles and Organization Management outlined in [sig-governance].

## Scope

Covers development and maintenance of components for automated scaling in
Kubernetes.  This includes automated vertical and horizontal pod
autoscaling, initial resource estimation, cluster-proportional system
component autoscaling, and autoscaling of Kubernetes clusters themselves.

### In scope

- Autoscaling-related API objects, such as the HorizontalPodAutoscaler and
  VerticalPodAutoscaler

- Autoscaling-related tools, such as the cluster autoscaler,
  single-component scaling tools (e.g. pod-nanny), and
  cluster-proportional scaling tools

- Ensuring API interfaces (the scale subresource) are availble and usable
  to enable other SIG to write autoscalable objects, and enable people to
  interact with those interfaces.

[Link to SIG section in sigs.yaml][sigs.yaml]

#### Code, Binaries and Services

- Components and utilities that take automated action to scale a component
  on the cluster

- Components and utilities that take automated action to scale the cluster
  itself

- Special parts of client-go for interacting with with the scaling
  interfaces used by the HPA (e.g. the polymorphic scale client).

#### Cross-cutting and Externally Facing Processes

- Reviewing implementations of the scale subresource to ensure that
  autoscaling behaves properly

- Coordinating with SIG Instrumentation to ensure that metrics APIs are
  suitable for autoscaling on.

- Coordinating with SIG Scheduling to make sure scheduling descisions can
  interact well with the cluster autoscaler

### Out of scope

- Testing general cluster performance at scale (this falls under the
  purview of [SIG Scalability]).

- Owning metrics APIs (this falls under the purview of [SIG
  Instrumentation].  SIG Autoscaling should collaborate with [SIG
  Instrumentation] to ensure that metrics APIs are suitable for using in
  autoscaling.

## Roles and Organization Management

This sig follows adheres to the Roles and Organization Management outlined in [sig-governance]
and opts-in to updates and modifications to [sig-governance].

### Deviations from [sig-governance]

- SIG Autoscaling does not have chairs as a separate entity from tech
  leads.  The tech leads have the responsibility of chairs.

### Subproject Creation

SIG Technical Leads

[sig-governance]: https://github.com/kubernetes/community/blob/master/committee-steering/governance/sig-governance.md
[sigs.yaml]: https://github.com/kubernetes/community/blob/master/sigs.yaml#L305
[Kubernetes Charter README]: https://github.com/kubernetes/community/blob/master/committee-steering/governance/README.md
[SIG Scalability]: https://github.com/kubernetes/community/blob/master/sig-scalability
[SIG Instrumentation]: https://github.com/kubernetes/community/blob/master/sig-instrumentation
