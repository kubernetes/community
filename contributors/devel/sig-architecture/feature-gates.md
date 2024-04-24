# Feature gate progression in Kubernetes

Feature flags are how Kubernetes tracks progression of functionality that is available when deployed.

Feature gates are intended to cover the development life cycle of a feature - they are not intended to be long-term APIs. As such, they are expected to be [deprecated](https://kubernetes.io/docs/reference/using-api/deprecation-policy) and removed after a feature becomes GA or is dropped. Optional features intended to be enabled or disabled by users even once the feature is GA should include a mechanism for enabling or disabling the feature (like a command-line flag or config file option) in addition to the associated feature gate.

## Lifecycle

Features generally progress through `Alpha` -> `Beta` -> `GA`. Sometimes we end up deciding that a feature is not going to be supported and we end up marking them as `Deprecated`.

The majority of features will go through all three stages, but occasionally there are features which may skip stages.
While some exceptions may happen, approvers should use the following guidance:
- features that involve [API changes] must progress through all `Alpha`, `Beta`, `GA` stages
- features that are unproven at achieving their goals, have significant complexity,
  risk of defects/problematic edge cases, or performance/scalability implications
  should progress through all `Alpha`, `Beta`, `GA` stages
- features which achieve their goals with minimal complexity and performance/scalability
  implications, but still carry non-trivial risk (e.g. due to changing user-facing behavior
  or problematic edge cases) might skip `Alpha` and start directly in `Beta`
  (provided the appropriate `Beta` quality is achieved) but should be off by-default until
  proven in representative production environment that utilizes the feature with the scale
  or variety of use to prove it's working
- more generally, changes that carry a risk of making previously working functionality
  no longer work in certain edge cases should always start in off-by-default state
- smaller changes with low enough risk that still may need to be disabled using the
  feature gate without introducing a new long term configuration option, might skip
  `Alpha` and start directly in `Beta` (provided the appropriate `Beta` quality is achieved)
  and can be enabled by-default from the very beginning
- bug fixes that have a sufficient level of risk that being able to turn off the fix via a
  feature gate is justified are recommended to go directly to `Beta` and should be enabled
  by-default from the very beginning; an alternative for bug fixes that could be perceived
  as "removal" is to use Deprecated state, however still ensuring that the fix can be
  disabled

[API changes]: https://github.com/kubernetes/community/blob/master/sig-architecture/api-review-process.md#what-parts-of-a-pr-are-api-changes

When we add a feature flag, we basically add if/else conditions to ensure that a feature is ONLY activated when either the default is on or if the deployer has switched it on explicitly. When a feature gate is disabled, the system should behave as if the feature doesn't exist. The only exception to this is [API input validation](https://kubernetes.io/docs/reference/using-api/deprecation-policy/#deprecating-parts-of-the-api) on updates, which should preserve and validate data if and only if it was present before the update (which could occur in case of a version rollback).

There is no supported way to change a feature gate at runtime for production Kubernetes use. A feature gate is typically toggled by a component restart.

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

By default Alpha features are not switched on. This enables folks to switch on the feature using the command line. All API changes must start with an Alpha gate, which makes it possible to rollback from future versions.

## Beta Features

* `PreRelease` is set to `featuregate.Beta`
* `Default` is usually set to `true` (see below)
* `LockToDefault` is not set. Defaults to `false`

Beta features are usually on by default. This enables the feature to be available in the default installation of Kubernetes.

Sometimes (rarely) the `Default` is set to `false` in Beta. This tells folks that while this feature is Beta, they will still need to do some work to switch it on and use it, and potentially take some other explicit action outside of Kubernetes. For example, the CSIMigration feature gates looked like this:

```
	CSIMigration:                                   {Default: true, PreRelease: featuregate.Beta},
	CSIMigrationGCE:                                {Default: false, PreRelease: featuregate.Beta}, // Off by default (requires GCE PD CSI Driver)
	CSIMigrationAWS:                                {Default: false, PreRelease: featuregate.Beta}, // Off by default (requires AWS EBS CSI driver)
	CSIMigrationAzureDisk:                          {Default: false, PreRelease: featuregate.Beta}, // Off by default (requires Azure Disk CSI driver)
	CSIMigrationAzureFile:                          {Default: false, PreRelease: featuregate.Beta}, // Off by default (requires Azure File CSI driver)
	CSIMigrationvSphere:                            {Default: false, PreRelease: featuregate.Beta}, // Off by default (requires vSphere CSI driver)
```

## GA Features

* `PreRelease` is set to `featuregate.GA`
* `Default` is always set to `true`
* `LockToDefault` is set to `true`

GA features are always on by default, and usually cannot be disabled.

Sometimes (rarely) we do not set `LockToDefault` (thus defaulting to `false`, meaning "not locked"). This enables folks to switch off the GA feature. We do this to indicate that while the feature (for example `coredns`) is GA, they need to move off `kubedns` to `coredns` in their infrastructure. If they want to continue using `kubedns` for a short time, they can choose to switch off the GA flag. When we eventually remove the support for `kubedns` entirely we would set `LockToDefault` to `true` with some grace period for the transition.

[After at least two releases post-GA and deprecation](https://kubernetes.io/docs/reference/using-api/deprecation-policy/#deprecation), the feature gate is removed. Typically, we add a comment in [kubefeatures.go](https://github.com/kubernetes/kubernetes/blob/master/pkg/features/kube_features.go) such as: `// remove in 1.23` to signal when we plan to remove the feature gate. Remember when the feature gate is removed and the deployer has forgotten to drop the reference to the feature in the CLI flags (say the `kube-apiserver`), then they will see a hard failure.

Also note that when we set `LockToDefault` to `true`, we remove all references (if/then conditions) to the feature gate from the codebase.

## Deprecation

* `PreRelease` is set to `featuregate.Deprecated`
* `Default` is set to `false`
* See [Kubernetes Deprecation Policy](https://kubernetes.io/docs/reference/using-api/deprecation-policy/#deprecation) for more details

Very rarely we will deprecate some aspect of Kubernetes (almost always something that has no actual impact).  When we do that, the pattern is to name the gate so it describes the functionality being deprecated and to default the value to `false`.  If some user is impacted by the deprecation, they can set that gate to `true` to unbreak themselves (and then file a bug).  If this happens, we must reconsider the deprecation and may choose to abandon it entirely by changing the gate back to `true` for a release or two and eventually removing it.

## Other scenarios

Sometimes we use`{Default: true, PreRelease: featuregate.Beta}` for keeping legacy behavior on while a new alternative is being implemented. When the new default behavior is transitioning to GA and it is time to drop the old legacy behavior, we will end up with `{Default: false, PreRelease: featuregate.GA, LockToDefault: true}`. For an example, see: [LegacyNodeRoleBehavior & ServiceNodeExclusion](https://github.com/kubernetes/kubernetes/pull/97543/files).
