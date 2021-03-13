# Feature gate progression in Kubernetes

Feature flags are how Kubernetes tracks progression of functionality that is available when deployed.

## Lifecycle

Feature progress through `Alpha` -> `Beta` -> `GA`. Sometimes we end up deciding that a feature is not going to be supported and we end up marking them as `Deprecated`.

When we add a feature flag, we basically add if/else conditions to ensure that a feature is ONLY activated when either the default is on or if the deployer has switched it on explicitly.

## Fields in the specification are
	// Default is the default enablement state for the feature. possible values are true/false.
	Default bool
	
    // LockToDefault indicates that the feature is locked to its default and cannot be changed. possible values are true/false.
	LockToDefault bool
	
    // PreRelease indicates the maturity level of the feature. possible values are "featuregate.Alpha", "featuregate.Beta", "featuregate.GA", "featuregate.Deprecated"
	PreRelease prerelease


## Alpha Features

* `PreRelease` is set to `featuregate.Alpha`
* `Default` is always set to `false`
* `LockToDefault` is not set. Defaults to `false`

This enables folks to switch on the feature using the command line. By default it is not switched on.

## Beta Features

* `PreRelease` is set to `featuregate.Beta`
* `Default` is always set to `true`
* `LockToDefault` is not set. Defaults to `false`

This enables the feature to be on and available in the default installation of kubernetes. 

Sometimes (rarely) the `Default` is set to `false`. For example when this flag needs some explicit action or support outside of Kubernetes. This tells folks that while this feature is Beta, you still need to do some work to switch it on and use it.

## GA Features

* `PreRelease` is set to `featuregate.GA`
* `Default` is always set to `true`
* `LockToDefault` is set to `true`

GA features are always on. 

Sometimes (rarely) we do not set `LockToDefault` and let it default to false. This enables folks to switch off the GA feature. We do this to tell folks that while the feature (for example `coredns`) is GA, they need to move off say `kubedns` to `coredns` in their infrastructures pretty soon and if they want to continue `kubedns` for a short time, they will have to switch off the GA flag. When we do remove the support for `kubedns` entirely we would set `LockToDefault` to `true` with some grace period for the transition


Typically we add a comment `// remove in 1.23` to signal when we would entirely remove the feature gate. Remember when the feature gate is removed and the deployer has forgotten to drop the reference to the feature in the CLI flags (say the `kube-apiserver`), then they will see a hard failure. 

Also note that when we set `LockToDefault` to `true`, we remove all references (if/then conditions) to the feature gate from the codebase. 

## Deprecation

* `PreRelease` is set to `featuregate.Deprecated`
* `Default` is always set to `true`

## Other scenarios

Sometimes we use`{Default: true, PreRelease: featuregate.Beta}` for keeping legacy behavior (see `LegacyNodeRoleBehavior`) on and when we basically want to drop the legacy behavior with new default behavior. So at GA we will end up with `{Default: false, PreRelease: featuregate.GA, LockToDefault: true}`

