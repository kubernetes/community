# WG Node Lifecycle Charter

This charter adheres to the conventions described in the [Kubernetes Charter README] and uses
the Roles and Organization Management outlined in [wg-governance].

[Kubernetes Charter README]: /committee-steering/governance/README.md

## Scope

The Kubernetes ecosystem currently faces challenges in node maintenance scenarios, with multiple
projects independently addressing similar issues. The goal of this working group is to develop
unified APIs that the entire ecosystem can depend on, reducing the maintenance burden across
projects and addressing scenarios that impede node drain or cause improper pod termination. Our
objective is to create easily configurable, out-of-the-box solutions that seamlessly integrate with
existing APIs and behaviors. We will strive to make these solutions minimalistic and extensible to
support advanced use cases across the ecosystem.

To properly solve the node drain, we must first understand the node lifecycle. This includes
provisioning/sunsetting of the nodes, PodDisruptionBudgets, API-initiated eviction and node
shutdown. This then impacts both the node and pod autoscaling, de/scheduling, load balancing, and
the applications running in the cluster. All of these areas have issues and would benefit from a
unified approach.

### In scope

- Explore a unified way of draining the nodes and managing node maintenance by introducing new APIs
  and extending the current ones. This includes exploring extension to or interactions with the Node
  object.
- Analyze the node lifecycle, the Node API, and possible interactions. We want to explore augmenting
  the Node API to expose additional state or status in order to coalesce other core Kubernetes and
  community APIs around node lifecycle management.
- Improve the disruption model that is currently implemented by API-initiated Eviction API and PDBs.
  Improve the descheduling, availability and migration capabilities of today's application
  workloads. Also explore the interactions with other eviction mechanisms.
- Coordinate pod termination and issues around de/scheduling, preemption and eviction.
- Improve the Graceful/Non-Graceful Node Shutdown and consider how this affects the node lifecycle.
  To graduate the [Graceful Node Shutdown](https://github.com/kubernetes/enhancements/issues/2000)
  feature to GA and resolve the associated node shutdown issues.
- Improve the scheduling and pod/node autoscaling to take into account ongoing node maintenance and
  the new disruption model/evictions. This includes balancing of the pods according to scheduling
  constraints. 
- Consider improving the pod lifecycle of DaemonSets and static pods during a node maintenance.
- Explore the cloud provider use cases and how they can hook into the node lifecycle. So that the
  users can use the same APIs or configurations across the board.
- Migrate users of the eviction based kubectl-like drain (kubectl, cluster autoscaler, karpenter,
  ...) and other scenarios to use the new unified node draining approach.
- Explore possible scenarios behind the reason why the node was terminated/drained/killed and how to
  track and react to each of them. Consider past discussions/historical perspective
  (e.g. "tombstones").

### Out of scope

- Implementing cloud provider specific logic, the goal is to have high-level API that the providers
  can use, hook into, or extend.
- Infrastructure provisioning, deprovisioning solution or physical infrastructure lifecycle
  management solution.

## Stakeholders

- SIG Apps
- SIG Autoscaling
- SIG CLI
- SIG Cloud Provider
- SIG Cluster Lifecycle
- SIG Network
- SIG Node
- SIG Scheduling
- SIG Storage

Stakeholders span from multiple SIGs to a broad set of end users,
public and private cloud providers, Kubernetes distribution providers,
and cloud provider end-users. Here are some user stories:

- As a cluster admin I want to have a simple interface to initiate a node drain/maintenance without
  any required manual interventions. I also want to be able to observe the node drain via the API
  and check on its progress. I also want to be able to discover workloads that are blocking the node
  drain.
- To support the new features, node maintenance, scheduler, descheduler, pod autoscaling, kubelet,
  and other actors should use a new eviction API to gracefully remove pods. This would enable new
  migration strategies that prefer to surge (upscale) pods first rather than downscale them. It
  would also allow other users/components to monitor pods that are gracefully removed/terminated
  and provide better behaviour in terms of de/scheduling, scaling and availability.
- As a cluster admin, I want to be able to perform arbitrary actions after the node drain is
  complete, such as resetting GPU drivers, resetting NICs, performing software updates or shutting
  down the machine.
- As an end user, I would like more alternatives to blue-green upgrades, especially with special
  hardware accelerators; it's far too expensive. I would like to choose a strategy on how to
  coordinate the node drain and the upgrade to achieve better cost-effectiveness.
- As a cloud provider, I need to perform regular maintenance on the hardware in my fleet. Enhancing
  Kubernetes to help CSPs safely remove hardware will reduce operational costs.
- Modelling the cost of doing accelerator maintenance in today's world can be massive. And since
  hardware accelerators tend to need more love and care, having software support to coordinate
  maintenance will reduce operational costs.
- As a cluster admin, I would like to use a mixture of on-demand and temporary spot instances in my
  clusters to reduce cloud expenditure. Having more reliable lifecycle and drain mechanisms for
  nodes will improve cluster stability in scenarios where instances may be terminated by the cloud
  provider due to cost-related thresholds.
- As a user, I want to prevent any disruption to my pet or expensive workloads (VMs, ML with
  accelerators) and either prevent termination altogether or have a reliable migration path. 
  Features like `terminationGracePeriodSeconds` are not sufficient as the termination/migration can
  take hours if not days.
- As a user, I want my application to finish all network and storage operations before terminating a
  pod. This includes closing pod connections, removing pods from endpoints, writing cached writes
  to the underlying storage and completing storage cleanup routines.

## Deliverables

The WG will coordinate requirement gathering and design, eventually leading to implementation.

Area we expect to explore:

- An API to express node drain/maintenance.
- An API to solve the problems wrt the API-initiated Eviction API and PDBs.
- An API/mechanism to gracefully terminate pods during a node shutdown.
- An API to deschedule pods that use DRA devices.
- An API to remove pods from endpoints before they terminate.
- Introduce enhancements across multiple Kubernetes SIGs to add support and integration for the new
  APIs to solve wide range of issues.

We expect to provide reference implementations of the new APIs including but not limited to
controllers, API validation, integration with existing core components and extension points for the
ecosystem. This should be accompanied by E2E / Conformance tests.

## Relevant KEPs and Documents

- https://github.com/kubernetes/enhancements/issues/4212
- https://github.com/kubernetes/enhancements/issues/4563
- https://github.com/kubernetes/enhancements/issues/2000
- https://github.com/kubernetes/enhancements/issues/5055
- https://docs.google.com/document/d/1t25jgO_-LRHhjRXf4KJ5xY_t8BZYdapv7MDAxVGY6R8/edit?tab=t.0#heading=h.i4lwa7rdng7y

## Relevant Projects

This is a list of known projects that solve similar problems in the ecosystem or would benefit from
the efforts of this WG:

- https://github.com/aws/aws-node-termination-handler
- https://github.com/foriequal0/pod-graceful-drain
- https://github.com/jukie/karpenter-deprovision-controller
- https://github.com/kubereboot/kured
- https://github.com/kubernetes/autoscaler/tree/master/cluster-autoscaler
- https://github.com/kubernetes-sigs/cluster-api/
- https://github.com/kubernetes-sigs/karpenter
- https://github.com/kubernetes-sigs/kubespray
- https://github.com/kubevirt/kubevirt
- https://github.com/medik8s/node-maintenance-operator
- https://github.com/Mellanox/maintenance-operator
- https://github.com/openshift/machine-config-operator
- https://github.com/planetlabs/draino
- https://github.com/strimzi/drain-cleaner

There are also internal custom solutions that companies use.

## Roles and Organization Management

This WG adheres to the Roles and Organization Management outlined in [wg-governance]
and opts-in to updates and modifications to [wg-governance].

[wg-governance]: /committee-steering/governance/wg-governance.md

## Timelines and Disbanding

The working group will disband once the features and core APIs defined in the KEPs have reached a
stable state (GA) and ongoing maintenance ownership is established within the relevant SIGs. We will
review whether the working group should disband if appropriate SIG ownership
can't be reached.
