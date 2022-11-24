# Feature gate progression in Kubernetes

Feature flags are how Kubernetes tracks progression of functionality that is available when deployed.

Feature gates are intended to cover the development life cycle of a feature - they are not intended to be long-term APIs. As such, they are expected to be [deprecated](https://kubernetes.io/docs/reference/using-api/deprecation-policy) and removed after a feature becomes GA or is dropped. Optional features intended to be enabled or disabled by users even once the feature is GA should include a mechanism for enabling or disabling the feature (like a command-line flag or config file option) in addition to the associated feature gate.

## Lifecycle

Feature progress through `Alpha` -> `Beta` -> `GA`. Sometimes we end up deciding that a feature is not going to be supported and we end up marking them as `Deprecated`.

The majority of features will go through all three stages, but occasionally there are features which may skip stages.

When we add a feature flag, we basically add if/else conditions to ensure that a feature is ONLY activated when either the default is on or if the deployer has switched it on explicitly. When a feature gate is disabled, the system should behave as if the feature doesn't exist. The only exception to this is [API input validation](https://kubernetes.io/docs/reference/using-api/deprecation-policy/#deprecating-parts-of-the-api) on updates, which should preserve and validate data if and only if it was present before the update (which could occur in case of a version rollback).

There is no supported way to trigger a feature gate at runtime for production Kubernetes use. A feature gate is typically toggled by a component restart.

Unless an exception is granted for a particular feature, as documented and approved as part of [Production Readiness Review], it is expected that:

- Toggling a feature gate will not affect other components, i.e. disabling a feature gate in the API server will work independently of disabling that same feature gate in the kubelet or scheduler, and no coordination between components is needed.
- The effects of toggling a feature gate should be limited to the scope of the feature. Enabling or disabling a feature gate should not affect workloads that do not use the feature gate.
- Toggling feature gates should not result in fanout effects or cascading interactions in a cluster.

[Production Readiness Review]: /sig-architecture/production-readiness.md

## Specification Fields
	// Default is the default enablement state for the feature. Possible values are true/false.
	Default bool
	
	// LockToDefault indicates that the feature is locked to its default and cannot be changed. Possible values are true/false.
	LockToDefault bool
	
	// PreRelease indicates the maturity level of the feature. possible values are "featuregate.Alpha", "featuregate.Beta", "featuregate.GA", "featuregate.Deprecated"
	PreRelease prerelease


## Alpha Features

* `PreRelease` is set to `featuregate.Alpha`
* `Default` is always set to `false`
* `LockToDefault` is not set. Defaults to `false`

By default it is not switched on. This enables folks to switch on the feature using the command line. All API changes must start with an Alpha gate, which makes it possible to rollback from future versions.

## Beta Features

* `PreRelease` is set to `featuregate.Beta`
* `Default` is usually set to `true` (see below)
* `LockToDefault` is not set. Defaults to `false`

This enables the feature to be on and available in the default installation of Kubernetes. 

Sometimes (rarely) the `Default` is set to `false`. This tells folks that while this feature is in Beta, they will still need to do some work to switch it on and use it and potentially take some other explicit action outside of Kubernetes. For example, see the [CSIMigration feature gates](https://github.com/kubernetes/kubernetes/blob/5b0a2c3a29f6b5392e0f8f94ba5669bdc9eb73f6/pkg/features/kube_features.go#L792).

## GA Features

* `PreRelease` is set to `featuregate.GA`
* `Default` is always set to `true`
* `LockToDefault` is set to `true`

GA features are always on. 

Sometimes (rarely) we do not set `LockToDefault` and let it default to `false`. This enables folks to switch off the GA feature. We do this to tell folks that while the feature (for example `coredns`) is GA, they need to move off say `kubedns` to `coredns` in their infrastructures pretty soon and if they want to continue `kubedns` for a short time, they will have to switch off the GA flag. When we do remove the support for `kubedns` entirely we would set `LockToDefault` to `true` with some grace period for the transition.

[After at least two releases post-GA and deprecation](https://kubernetes.io/docs/reference/using-api/deprecation-policy/#deprecation), the feature gate is removed. Typically, we add a comment in [kubefeatures.go](https://github.com/kubernetes/kubernetes/blob/master/pkg/features/kube_features.go) such as: `// remove in 1.23` to signal when we plan to remove the feature gate. Remember when the feature gate is removed and the deployer has forgotten to drop the reference to the feature in the CLI flags (say the `kube-apiserver`), then they will see a hard failure. 

Also note that when we set `LockToDefault` to `true`, we remove all references (if/then conditions) to the feature gate from the codebase. 

## Deprecation

* `PreRelease` is set to `featuregate.Deprecated`
* See [Kubernetes Deprecation Policy](https://kubernetes.io/docs/reference/using-api/deprecation-policy/#deprecation) for more details

## Other scenarios

Sometimes we use`{Default: true, PreRelease: featuregate.Beta}` for keeping legacy behavior on while a new alternative is being implemented. When the new default behavior is transitioning to GA and it is time to drop the old legacy behavior, we will end up with `{Default: false, PreRelease: featuregate.GA, LockToDefault: true}`. For an example, see: [LegacyNodeRoleBehavior & ServiceNodeExclusion](https://github.com/kubernetes/kubernetes/pull/97543/files).
