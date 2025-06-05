# Feature Gates in Kubernetes

Feature "gates" are how Kubernetes manages behavioral changes as they progress
through the development lifecycle. This gives cluster operators a way to try
new features or to disable them when something goes wrong.

This last point bears repeating. all software has bugs, and new software tends
to have more (and worse) bugs than mature software. A feature gate is intended
to be a fast way to neutralize a new feature and mitigate damage caused by
bugs. Feature authors and reviewers should invest time into thinking about
whether their gates meet this goal.

Feature gates are _not_ intended to be long-term APIs. Individual gates are
expected to be
[deprecated](https://kubernetes.io/docs/reference/using-api/deprecation-policy)
and removed after a feature becomes GA (or is dropped). Truly optional
capabilities which are permanently intended to be enabled or disabled by users
(even once the feature is GA) should include a mechanism for enabling or
disabling the feature (like a command-line flag or config file option) in
addition to the associated feature gate.

NOTE: This document is fairly high-level. For more precise details, including
code snippets, see [api_changes.md](api_changes.md).

## Defining a feature gate

A feature gate definition is comprised of the following struct fields:

```
	// Default is the default enablement state for the feature. If the value of
	// the gate is not explicitly set, this value is used.
	Default bool
	
	// LockToDefault indicates that the feature is locked to its default value
	// and cannot be changed.
	LockToDefault bool
	
	// PreRelease indicates the maturity level of the feature. Possible values
	// are "featuregate.Alpha", "featuregate.Beta", "featuregate.GA", or
	// "featuregate.Deprecated"
	PreRelease prerelease
```

Unless an exception is granted for a particular feature, as documented and
approved as part of [Production Readiness Review], it is expected that:

- Toggling a feature gate will not affect other components, e.g. disabling a
  feature gate in the API server will work independently of disabling that same
  feature gate in the kubelet or scheduler, and no coordination between
  components is needed.
- The effects of toggling a feature gate should be limited to the scope of the
  feature. Enabling or disabling a feature gate should not affect workloads
  that do not use the feature gate.
- Toggling feature gates should not result in fanout effects or cascading
  interactions in a cluster.
- Disabling a feature gate should prevent any further damage caused by bugs in
  the feature.

[Production Readiness Review]: /sig-architecture/production-readiness.md

## Enabling and disabling feature gates

Gates can be enabled or disabled either via flags (see `--feature-gates`) or
component config files (see `featureGates`), and some can be enabled via
environment variables.

There is no supported way to change a feature gate at runtime for production
Kubernetes use. A feature gate is typically toggled by a component restart.

## Lifecycle

Features generally progress through `Alpha` -> `Beta` -> `GA`. Sometimes we end
up deciding that a feature is not going to be supported and we end up marking
them as `Deprecated`.

The majority of features will go through all three stages, but occasionally
there are features which may skip a stage.

While some exceptions may happen, approvers should use the following guidance:
- Features that involve [API changes] must progress through all `Alpha`,
  `Beta`, `GA` stages.
- Features that are unproven at achieving their goals, have significant
  complexity, risk of defects/problematic edge cases, or
  performance/scalability implications should progress through all `Alpha`,
  `Beta`, `GA` stages.
- Features which achieve their goals with minimal complexity and
  performance/scalability implications, but still carry non-trivial risk (e.g.
  due to changing user-facing behavior or problematic edge cases) might skip
  `Alpha` and start directly in `Beta` (provided the appropriate `Beta` quality
  is achieved) but should be off by default until proven in representative
  production environment that utilizes the feature with the scale or variety of
  use to prove it's working.
- Changes that carry a risk of making previously working functionality no
  longer work in certain edge cases should always start as off by default.
- Smaller changes with low enough risk that still may need to be disabled using
  the feature gate without introducing a new long term configuration option,
  might skip `Alpha` and start directly in `Beta` (provided the appropriate
  `Beta` quality is achieved) and can be enabled by default from the very
  beginning.
- Bug fixes that have a sufficient level of risk that being able to turn off
  the fix via a feature gate is justified are recommended to go directly to
  `Beta` and should be enabled by default from the very beginning; an
  alternative for bug fixes that could be perceived as "removal" is to use
  Deprecated state, however still ensuring that the fix can be disabled.

[API changes]: https://github.com/kubernetes/community/blob/master/sig-architecture/api-review-process.md#what-parts-of-a-pr-are-api-changes

### Compatibility versions

The Kubernetes "compatibility version" feature promises that control-plane
components can be configured to behave as if they were any of the three previous
releases, which includes (non-alpha) feature gates and gate-controlled APIs and logic.

As a feature progresses through the lifecycle, we must preserve enough
information to allow such compatible configuration, including both the old and
new states, along with the version that the transition occurred.

Example:

```
RetryGenerateName: {  
  {Version: version.MustParse("1.30"), Default: false, PreRelease: featuregate.Alpha},  
  {Version: version.MustParse("1.31"), Default: true, PreRelease: featuregate.Beta},  
  {Version: version.MustParse("1.32"), Default: true, LockToDefault: true, PreRelease: featuregate.GA},  
},  
```

In this example, a component at v1.33 can be configured to behave as v1.32+ or
to be compatible with v1.31 (where the gate was disableable) or v1.30 (where the
gate was off by default). That component at v1.34 can be configured to behave as
v1.32+ or to be compatible with v1.31 (where the gate was disableable), but
v1.30 falls outside the "three releases" window.

At v1.35, the component can only be configured to be compatible as far back as
1.32 (where the gate was locked on), and so the lifecycle is complete.

### Alpha features

* `PreRelease` is set to `featuregate.Alpha`
* `Default` is always set to `false`
* `LockToDefault` is set to `false` (or not specified)

Alpha features are never enabled by default, but users may switch them on
explicitly.

### Beta features

* `PreRelease` is set to `featuregate.Beta`
* `Default` is usually set to `true` (see below)
* `LockToDefault` is set to `false` (or not specified)

Beta features are usually enabled by default (note that beta features are not
the same thing as beta APIs). This allows them to be available in the default
installation of Kubernetes. Users may switch them off explicitly.

Sometimes (rarely) a beta feature will be disabled by default. This tells users
that while this feature is Beta, they will still need to do some work to switch
it on and use it, which may include taking some other explicit action outside of
Kubernetes. For example, the CSIMigration feature gates looked like this:

```
	CSIMigration:          {Default: true,  PreRelease: featuregate.Beta},
	CSIMigrationGCE:       {Default: false, PreRelease: featuregate.Beta}, // Off by default (requires GCE PD CSI Driver)
	CSIMigrationAWS:       {Default: false, PreRelease: featuregate.Beta}, // Off by default (requires AWS EBS CSI driver)
	CSIMigrationAzureDisk: {Default: false, PreRelease: featuregate.Beta}, // Off by default (requires Azure Disk CSI driver)
	CSIMigrationAzureFile: {Default: false, PreRelease: featuregate.Beta}, // Off by default (requires Azure File CSI driver)
	CSIMigrationvSphere:   {Default: false, PreRelease: featuregate.Beta}, // Off by default (requires vSphere CSI driver)
```

### GA features

* `PreRelease` is set to `featuregate.GA`
* `Default` is always set to `true`
* `LockToDefault` is usually set to `true` (see below)

GA features are always on by default, and usually cannot be disabled.

Sometimes (rarely) a GA feature is allowed to be disabled (`LockToDefault: false`).
This is used for features that are being enabled by default for the first time.
Such features must be disable-able in the first release they are on by default.
The most common case is when an entirely new API is created and is not enabled by
default until the serialization is stable (GA).
This can also indicate that while this feature is GA, they need to take some
other action outside of Kubernetes to use it. This gives some grace period for
users to take action, but such feature gates will eventually set
`LockToDefault` to `true` and then be retired, like normal.

To achieve our [compatibility version](#compatibility-versions) promise, after
three releases where a feature has been locked to the default value (whether
that feature is GA or deprecated), feature gates and references should be
removed. We use three releases because it corresponds to roughly one year in the
Kubernetes development cycle which is our [support
period](https://kubernetes.io/releases/patch-releases/#support-period). For
example, if a gate was `LockToDefault: true` in kubernetes version `X`, it may
be removed in version `X+3` (which must be compatible with `X+2`, `X+1`, and
`X`, all of which were also locked).

Typically, we add a comment in
[the code](https://github.com/kubernetes/kubernetes/blob/master/pkg/features/kube_features.go)
such as: `// remove in 1.23` to signal when we plan to remove the feature gate.
We provide this grace period to give users time to stop referencing "finished"
gates. If a feature gate is removed and a user has forgotten to drop the
reference to it (e.g. in the CLI flags of `kube-apiserver`), then they will see
a hard failure.

#### Disablement Tests

Typically for full coverage, unit and integration tests exist for both when a
feature is enabled and disabled. When a feature is promoted to GA, it is usually
locked to true by default and cannot be unset in testing. For integration with
compatibility version, feature disablement tests should be maintained until the
GA feature is fully removed three releases after promotion. To test scenarios
where the feature gate may still be disabled, emulation version should be set in
disablement tests. For these disablement tests, simply set emulation version to
the version before the GA promotion to support disablement. The emulation
version would be set the line before the feature gate is set.

For example, if the "CustomResourceFieldSelectors" becomes GA in version 1.32,
the emulation version is set to v1.31.

```go
featuregatetesting.SetFeatureGateEmulationVersionDuringTest(t, utilfeature.DefaultFeatureGate, version.MustParse("1.31"))
featuregatetesting.SetFeatureGateDuringTest(t, utilfeature.DefaultFeatureGate, apiextensionsfeatures.CustomResourceFieldSelectors, false)
```

When the feature gate is removed three releases later in v1.35, the disablement
feature gate test is then removed. For tests using a matrix, emulation version
should only be set on tests that disable the feature.


```go
testcases := []struct{
  ...
  featureEnabled bool
}

for _, tc := range testcases {
  if !tc.featureEnabled {
    featuregatetesting.SetFeatureGateEmulationVersionDuringTest(t, utilfeature.DefaultFeatureGate, version.MustParse("1.31"))
  }
  featuregatetesting.SetFeatureGateDuringTest(t, utilfeature.DefaultFeatureGate, features.MyFeature, tc.featureEnabled)
}
```

Disablement tests are only required to be preserved for components and libraries
that support compatibility version. Tests for node and kubelet are unaffected by
compatibility version.

### Deprecation

* `PreRelease` is set to `featuregate.Deprecated`
* `Default` is set to `false`
* See [Kubernetes Deprecation Policy](https://kubernetes.io/docs/reference/using-api/deprecation-policy/#deprecation) for more details

Very rarely we will deprecate some user-visible aspect of Kubernetes (almost
always something that has no actual impact or which is truly a bug). When we
do that, we also use a feature gate. The pattern is to name the gate so it
describes the functionality being deprecated and to default the value to
`false`.

If some user is impacted by the deprecation, they can set that gate to `true`
to unbreak themselves (and then file a bug). If this happens, we must
reconsider the deprecation and may choose to abandon it entirely by changing
the gate back to `true` for a release or two and eventually removing it.

Once the [deprecation period](https://kubernetes.io/docs/reference/using-api/deprecation-policy/#deprecation)
has passed, the gate should be locked to the default value (`LockToDefault:
true`). As with GA features, all references to the feature gate must be kept for
a minimum of three releases after gate has been locked to the default value. The
gate, all references to it, and all gated logic may be removed after those three
releases. See [compatibility version](#compatibility-versions) for more details.

For example:

```
DeprecatedFeature: {  
  {Version: version.MustParse("1.29"), Default: false, PreRelease: featuregate.Alpha}, // feature graduated to alpha.
  {Version: version.MustParse("1.30"), Default: true, PreRelease: featuregate.Beta},  // feature graduated to beta.
  {Version: version.MustParse("1.32"), Default: false, PreRelease: featuregate.Deprecated}, // feature is deprecated and turned off. LockToDefault is unset to give users time to transition.
  {Version: version.MustParse("1.34"), Default: false, LockToDefault: true, PreRelease: featuregate.Deprecated}, // feature is deprecated, off, and locked. remove in v1.37 once versions prior to v1.34 cannot be emulated anymore.
},  
```

NOTE: We do not remove GA fields from the API.

### Other scenarios

Sometimes we use `{Default: true, PreRelease: featuregate.Beta}` for keeping
legacy behavior on while a new alternative is being implemented. When the new
default behavior is transitioning to GA and it is time to drop the old legacy
behavior, we will end up with `{Default: false, PreRelease: featuregate.GA,
LockToDefault: true}`. For an example, see: [LegacyNodeRoleBehavior &
ServiceNodeExclusion](https://github.com/kubernetes/kubernetes/pull/97543/files).

## Using a feature gate in code

Feature gates are basically fancy `bool` variables. They are either enabled or
disabled. When implementing a feature gate, there are a few patterns you might
follow.

### Features which add a new API field

All features which add API fields *must* start in alpha. This ensures that once
the feature becomes beta (and the fields get used) it is possible to roll-back
one Kubernetes version and not have the data be discarded. Users who enable and
use alpha features (which are disabled by default) do not have a safe rollback
target.

When a feature gate is disabled, the system should behave as if the API fields
do not exist. API operations which try to use the field are expected to proceed
as if the field was unknown and the "extra" data was discarded.

The API's registry code (`pkg/registry/...`) must check the gate before
validation. If the gate is disabled and the operation is a CREATE, the new
field must be removed (set to `nil`). If the gate is disabled and the
operation is an UPDATE, the previous form of the object must be checked. Only
if this object was not already using this field should it be removed. This
usually manifests as something like:

```
if disabled(gate) && !newFieldInUse(oldObj) {
    obj.NewField = nil
}
```

#### Validation: fields with no default value

For optional fields without a default value, API validation should *not* check
the gate. Instead, the usual pattern for an optional field applies: if the
field has a value, the value must be validated. This ensures that once an
object is validated and accepted with a feature gate enabled, subsequent
changes to the gate will not cause the saved object to fail validation.

#### Validation: fields with a default value

Some fields are optional in the API, but have a default value, which means they
are effectively required fields as far as the implementation is concerned. Once
the feature reaches GA and the feature gate is removed, there should never be a
case where such a field does not have a value.

Validation for these fields usually looks something like:

```
if obj.NewField == nil {
    allErrs = append(allErrs, field.Required(...))
} else {
    if !newFieldValid(obj.NewField, fldPath.Child("newField")) {
        allErrs = append(allErrs, field.Invalid(...))
    }
}
```

While a feature gate is active, the validation must also consider the gate's
state, the API operation in question, and the previous state of the object (in
case of UPDATE). It would be wrong to assert that the field is required when
the feature gate is disabled, but also wrong not to require it when the feature
gate is enabled.

The API's registry code (`pkg/registry/...`) must check the gate before
validation and pass a flag into the validation logic (usually as a field in an
"options" struct) to indicate whether the validation code should allow the new
field or not.

```
if enabled(gate) || newFieldInUse(oldObj) {
    options.EnableNewField = true
}
ValidateThisObject(obj, oldObj, options)
```

The validation code then looks something like:

```
if opts.EnableNewField {
    if obj.NewField == nil {
        allErrs = append(allErrs, field.Required(...))
    } else {
        if !newFieldValid(obj.NewField, fldPath.Child("newField")) {
            allErrs = append(allErrs, field.Invalid(...))
        }
    }
}
```

#### Implementation logic

The implementation of such a feature will usually want to check the gate, and
neutralize itself when disabled. In general, when a feature gate is disabled,
the system should behave as if the feature doesn't exist. This may result in
the API field carrying a value but not actually triggering any functionality.
This is preferable to a bug in the implementation causing harm that cannot be
disabled by the gate.

### Features which change an existing API field

For features which don't add a new field, but do expand what values are
allowed in the API (e.g. adding a value to an enum or loosening validation),
feature gates *must* start in alpha, for the same reasons as those with new
fields. Features which do not change the API schema, but do change API
operations (e.g. allowing updates to previously immutable fields) may be able
to skip alpha.

Unlike new fields, which can be removed entirely, we almost never want to
change the value of user input. For these features, the API validation logic
is the first time we have an opportunity to act.

#### Validation

The API's registry code (`pkg/registry/...`) must check the gate before
validation. As with the case of a new field, this logic must consider the
value of the field, as well as the current operation and (in case of UPDATE)
the previous state of the object. Similar to new fields with default values, it
must pass a flag into the validation logic (usually as a field in an "options"
struct) to indicate whether the validation code should allow the new value or
not.

```
if enabled(gate) || newFieldValueInUse(oldObj) {
    options.AllowNewFieldValue = true
}
ValidateThisObject(obj, oldObj, options)
```

#### Implementation logic

The implementation of such a feature may or may not need to check the gate,
depending on the feature. Unlike a new field, it is impossible to act like the
feature doesn't exist when the feature gate is disabled. The feature
implementation must decide what to do when the gate is disabled but the feature
gated value has already been used and stored.

Some features can fall back on a substantially similar value, and some must use
the new value. Emphasis should be placed on risk mitigation - if the feature
has a bug, disabling the gate *should* stop or at least bound the damage.

### Features which do not change the API

Features which have no API surface but still change behavior enough to warrant
a feature gate may start in alpha or (rarely) in beta. Unlike features with API
surface, the implementation logic is the only place where a feature gate can be
applied. This usually manifests as a simple `if`/`else` block:

```
if enabled(gate) {
  doNewThing()
} else {
  doOldThing()
}
```

As with API based features, the system should behave as if the feature doesn't
exist when the feature gate is disabled. Given the wide variety of features,
the exact meaning of "behave as if the feature doesn't exist" must be
determined by the implementation of each feature. Emphasis should be placed on
risk mitigation - if the feature has a bug, disabling the gate *should* stop or
at least bound the damage.
