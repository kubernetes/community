# Node Declared Features

This guide is for Kubernetes contributors adding a node-level feature that
the control plane must know about. It explains when a feature must be
registered as a _node declared feature_ and how to do it.

Node Declared Features
([KEP-5328](https://git.k8s.io/enhancements/keps/sig-node/5328-node-declared-features))
lets each node publish the node-level
features it supports in `node.status.declaredFeatures`, so the control plane
can keep Pods off nodes that cannot run them: the kube-scheduler, via its
`NodeDeclaredFeatures` plugin, filters out nodes missing a feature the Pod
requires, the `NodeDeclaredFeatureValidator` admission plugin rejects updates
to bound Pods that would require a feature their node does not declare, and
the kubelet re-checks requirements at admission as a final safeguard.

## When to add a declared feature

Add one when all of the following hold:

1. **Version-skew sensitivity.** Pods opt into behavior (usually via a
  `PodSpec` field) that only works when the node's kubelet implements it,
  so in a skewed cluster a Pod could land on a node that silently ignores
  or fails it.
2. **Feature gate association.** The feature is guarded by at least one
  feature gate (enforced by a registry consistency test).
3. **Control plane actionability.** A control plane component — the
  kube-scheduler, an admission plugin, or the API server — can use the
  declaration to make a decision.
4. **Static determinability.** Node support is decidable at kubelet
  startup from the kubelet's feature gates and static configuration — not
  from dynamic runtime state, hardware, or topology.

Do **not** add one for:

1. **Permanent node attributes** — hardware, architecture, or other
  capabilities that are not feature-gate dependent.
2. **Features that have already graduated to Beta or GA.** The declaration
  must be introduced with the feature at Alpha, before any release ships
  the behavior without declaring it; otherwise released nodes that support
  the feature but do not declare it are indistinguishable from nodes that
  lack it. This constrains when the declaration is _introduced_ — a feature
  that starts declaring at Alpha keeps its declared feature through Beta
  and GA.
3. **Conditionally gate-dependent behavior.** If the feature depends on its
  gate only under certain configurations (for example, in-place CPU resize
  for Guaranteed QoS Pods depends on
  `InPlacePodVerticalScalingExclusiveCPUs` only when the static CPU manager
  policy is enabled), it cannot be declared.

## How it works

- The kubelet discovers supported features once at startup and publishes the
  sorted list in `node.status.declaredFeatures`. The list does not change
  without a kubelet restart.
- The scheduler's `NodeDeclaredFeatures` plugin infers a Pod's required
  features from its spec and filters out nodes whose declared set does not
  satisfy them. The kubelet re-checks the same requirements at admission,
  protecting against scheduler bypass (for example, Pods created with
  `nodeName` set).
- The `NodeDeclaredFeatureValidator` admission plugin infers requirements
  from updates to bound Pods (the main resource and the `resize`
  subresource) and rejects the update if the Pod's node lacks them. It also
  rejects ResourceSlices that set `NodeAllocatableResources` on devices
  targeting nodes that do not declare `DRANodeAllocatableResources`.
- Some declarations are consumed by components reading
  `node.status.declaredFeatures` directly, outside the generic scheduler
  and admission paths. For example, the API server checks a node's declared
  features to decide whether it can use WebSockets when connecting to the
  kubelet for exec, attach, and port-forward
  (`ExtendWebSocketsToKubelet`). Such features register
  with both inference methods returning false (the declaration-only
  archetype).
- Every declared feature is temporary: during post-GA cleanup, the feature
  sets `MaxVersion`. On versions newer than `MaxVersion`, the kubelet stops
  declaring the feature and the scheduler and admission plugins treat it as
  universally available. The feature is deleted from the registry when its
  feature gate is removed (see [Lifecycle](#lifecycle)).

## Adding a declared feature

The change is three steps plus tests:

1.  A feature package implementing the `Feature` interface, in the
    framework's directory
    [`k8s.io/component-helpers/nodedeclaredfeatures`](https://git.k8s.io/kubernetes/staging/src/k8s.io/component-helpers/nodedeclaredfeatures).
2.  Append your feature to the registry
    (`k8s.io/component-helpers/nodedeclaredfeatures/features/registry.go`).
3.  Implement the `Feature` interface, defined in
    [`k8s.io/component-helpers/nodedeclaredfeatures/types/types.go`](https://git.k8s.io/kubernetes/staging/src/k8s.io/component-helpers/nodedeclaredfeatures/types/types.go).
    This is the expected interface for each new declared feature. See the inline comments for more details:

```go
type Feature interface {
	// Name returns the feature's well-known name, published in
	// node.status.declaredFeatures. By convention this is the feature
	// gate name. The name is part of the Node API surface and must be
	// unique across all the registered features. Adding a name is a
	// formal API change and requires API review.
	Name() string

	// Discover reports whether this node provides the feature, based on
	// the node's feature gates and static configuration. The feature
	// gates listed in Requirements() are determining: if any of them is
	// disabled, Discover must return false. Called by kubelet during startup.
	Discover(cfg *NodeConfiguration) bool

	// Requirements declares the feature gates the feature depends on
	// (at least one; mandatory). It must stay consistent with what
	// Discover actually checks; both are enforced by framework tests.
	Requirements() *FeatureRequirements

	// InferForScheduling reports whether scheduling this Pod requires
	// the feature. Called at the PreFilter extension point of the
	// NodeDeclaredFeatures scheduler plugin and by the kubelet's
	// pod admission handler, for every Pod.
	InferForScheduling(podInfo *PodInfo) bool

	// InferForUpdate reports whether this update to a bound Pod
	// introduces a requirement on the feature. Called by the
	// NodeDeclaredFeatureValidator admission plugin.
	InferForUpdate(oldPodInfo, newPodInfo *PodInfo) bool

	// MaxVersion is the last Kubernetes version (inclusive) at which
	// the feature is still a scheduling factor. Set based on the feature's
	// GA version plus the supported version skew, after which components
	// stop declaring and stop requiring the feature.
	MaxVersion() *version.Version
}
```

## Lifecycle

- **Alpha / Beta:** the kubelet declares the feature wherever its gate is
  enabled; nothing to do beyond the steps above.
- **GA:** keep declaring — this is what lets the control plane recognize
  older nodes that predate the GA.
- **Post-GA cleanup:** once every kubelet version within the
  [supported skew](https://kubernetes.io/releases/version-skew-policy/)
  has the feature locked on, set `MaxVersion` to the last
  version where the feature is still a scheduling factor (gate locked on in
  1.N with a three-version kubelet skew → `MaxVersion` is 1.N+2).
  Components newer than `MaxVersion` stop declaring and stop requiring the
  feature automatically. When the gate is removed, delete the feature
  implementation and its registry entry. Declared features must not
  accumulate.

## Reference

- [KEP-5328: Node Declared Features](https://github.com/kubernetes/enhancements/tree/master/keps/sig-node/5328-node-declared-features)
