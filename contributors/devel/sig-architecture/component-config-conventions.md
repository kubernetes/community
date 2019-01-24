# Component Configuration Conventions

# Objective

This document concerns the configuration of Kubernetes system components (as
opposed to the configuration of user workloads running on Kubernetes).
Component configuration is a major operational burden for operators of
Kubernetes clusters. To date, much literature has been written on and much
effort expended to improve component configuration. Despite this, the state of
component configuration remains dissonant. This document attempts to aggregate
that literature and propose a set of guidelines that component owners can
follow to improve consistency across the project.

# Background

Currently, component configuration is primarily driven through command line
flags. Command line driven configuration poses certain problems which are
discussed below. Attempts to improve component configuration as a whole have
been slow to make progress and have petered out (ref componentconfig api group,
configmap driven config issues). Some component owners have made use case
specific improvements on a per-need basis. Various comments in issues recommend
subsets of best design practice but no coherent, complete story exists.

## Pain Points of Current Configuration

Flag based configuration has poor qualities such as:

1.  Flags exist in a flat namespace, hampering the ability to organize them and expose them in helpful documentation. --help becomes useless as a reference as the number of knobs grows. It's impossible to distinguish useful knobs from cruft.
1.  Flags can't easily have different values for different instances of a class. To adjust the resync period in the informers of O(n) controllers requires O(n) different flags in a global namespace.
1.  Changing a process's command line necessitates a binary restart. This negatively impacts availability.
1.  Flags are unsuitable for passing confidential configuration. The command line of a process is available to unprivileged process running in the host pid namespace.
1.  Flags are a public API but are unversioned and unversionable.
1.  Many arguments against using global variables apply to flags.

Configuration in general has poor qualities such as:

1.  Configuration changes have the same forward/backward compatibility requirements as releases but rollout/rollback of configuration largely untested. Examples of configuration changes that might break a cluster: kubelet CNI plugin, etcd storage version.
1.  Configuration options often exist only to test a specific feature where the default is reasonable for all real use cases. Examples: many sync periods.
1.  Configuration options often exist to defer a "hard" design decision and to pay forward the "TODO(someone-else): think critically".
1.  Configuration options are often used to workaround deficiencies of the API. For example `--register-with-labels` and `--register-with-taints` could be solved with a node initializer, if initializers existed.
1.  Configuration options often exist to take testing shortcuts. There is a mentality that because a feature is opt-in, it can be released as a flag without robust testing.
1.  Configuration accumulates new knobs, knobs accumulate new behaviors, knobs are forgotten and bitrot reducing code quality over time.
1.  Number of configuration options is inversely proportional to test coverage. The size of the configuration state space grows >O(2^n) with the number of configuration bits. A handful of states in that space are ever tested.
1.  Configuration options hamper troubleshooting efforts. On github, users frequently file tickets from environments that are neither consistent nor reproducible.

## Types Of Configuration

Configuration can only come from three sources:

1.  Command line flags.
1.  API types serialized and stored on disk.
1.  API types serialized and stored in the kubernetes API.

Configuration options can be partitioned along certain lines. To name a few
important partitions:

1.  Bootstrap: This is configuration that is required before the component can contact the API. Examples include the kubeconfig and the filepath to the kubeconfig.
1.  Dynamic vs Static: Dynamic config is config that is expected to change as part of normal operations such as a scheduler configuration or a node entering maintenance mode. Static config is config that is unlikely to change over subsequent deployments and even releases of a component.
1.  Shared vs Per-Instance: Per-Instance configuration is configuration whose value is unique to the instance that the node runs on (e.g. Kubelet's `--hostname-override`).
1.  Feature Gates: Feature gates are configuration options that enable a feature that has been deemed unsafe to enable by default.
1.  Request context dependent: Request context dependent config is config that should probably be scoped to an attribute of the request (such as the user). We do a pretty good job of keeping these out of config and in policy objects (e.g. Quota, RBAC) but we could do more (e.g. rate limits).
1.  Environment information: This is configuration that is available through downwards and OS APIs, e.g. node name, pod name, number of cpus, IP address.

# Requirements

Desired qualities of a configuration solution:

1.  Secure: We need to control who can change configuration. We need to control who can read sensitive configuration.
1.  Manageable: We need to control which instances of a component uses which configuration, especially when those instances differ in version.
1.  Reliable: Configuration pushes should just work. If they fail, they should fail early in the rollout, rollback config if possible, and alert noisily.
1.  Recoverable: We need to be able to update (e.g. rollback) configuration when a component is down.
1.  Monitorable: Both humans and computers need to monitor configuration; humans through json interfaces like /configz, computers through interfaces like prometheus /streamz. Confidential configuration needs to be accounted for, but can also be useful to monitor in an unredacted or partially redacted (i.e. hashed) form.
1.  Verifiable: We need to be able to verify that a configuration is good. We need to verify the integrity of the received configuration and we need to validate that the encoded configuration state is sensible.
1.  Auditable: We need to be able to trace the origin of a configuration change.
1.  Accountable: We need to correlate a configuration push with its impact to the system. We need to be able to do this at the time of the push and later when analyzing logs.
1.  Available: We should avoid high frequency configuration updates that require service disruption. We need to take into account system component SLA.
1.  Scalable: We need to support distributing configuration to O(10,000) components at our current supported scalability limits.
1.  Consistent: There should exist conventions that hold across components.
1.  Composable: We should favor composition of configuration sources over layering/templating/inheritance.
1.  Normalized: Redundant specification of configuration data should be avoided.
1.  Testable: We need to be able to test the system under many different configurations. We also need to test configuration changes, both dynamic changes and those that require process restarts.
1.  Maintainable: We need to push back on ever increasing cyclomatic complexity in our codebase. Each if statement and function argument added to support a configuration option negatively impacts the maintainability of our code.
1.  Evolvable: We need to be able to extend our configuration API like we extend our other user facing APIs. We need to hold our configuration API to the same SLA and deprecation policy of public facing APIs. (e.g. [dynamic admission control](https://github.com/kubernetes/community/pull/611) and [hooks](https://github.com/kubernetes/kubernetes/issues/3585))

These don't need to be implemented immediately but are good to keep in mind. At
some point these should be ranked by priority and implemented.

# Two Part Solution:

## Part 1: Don't Make It Configuration

The most effective way to reduce the operational burden of configuration is to
minimize the amount of configuration. When adding a configuration option, ask
whether alternatives might be a better fit.

1.  Policy objects: Create first class Kubernetes objects to encompass how the system should behave. These are especially useful for request context dependent configuration. We do this already in places such as RBAC and ResourceQuota but we could do more such as rate limiting. We should never hardcode groups or usermaps in configuration.
1.  API features: Use (or implement) functionality of the API (e.g. think through and implement initializers instead of --register-with-label). Allowing for extension in the right places is a better way to give users control.
1.  Feature discovery: Write components that introspect the existing API to decide whether to enable a feature or not. E.g. controller-manager should start an app controller if the app API is available, kubelet should enable zram if zram is set in the node spec.
1.  Downwards API: Use the APIs that the OS and pod environment expose directly before opting to pass in new configuration options.
1.  const's: If you don't know whether tweaking a value will be useful, make the value const. Only give it a configuration option once there becomes a need to tweak the value at runtime.
1.  Autotuning: Build systems that incorporate feedback and do the best thing under the given circumstances. This makes the system more robust. (e.g. prefer congestion control, load shedding, backoff rather than explicit limiting).
1.  Avoid feature flags: Turn on features when they are tested and ready for production. Don't use feature flags as a fallback for poorly tested code.
1.  Configuration profiles: Instead of allowing individual configuration options to be modified, try to encompass a broader desire as a configuration profile. For example: instead of enabling individual alpha features, have an EnableAlpha option that enables all. Instead of allowing individual controller knobs to be modified, have a TestMode option that sets a broad number of parameters to be suitable for tests.

## Part 2: Component Configuration

### Versioning Configuration

We create configuration API groups per component that live in the source tree of
the component. Each component has its own API group for configuration.
Components will use the same API machinery that we use for other API groups.
Configuration API serialization doesn't have the same performance requirements
as other APIs so much of the codegen can be avoided (e.g. ugorji, generated
conversions) and we can instead fallback to the reflection based implementations
where they exist.

Configuration API groups for component config should be named according to the
scheme `<component>.config.k8s.io`. The `.config.k8s.io` suffix serves to
disambiguate types of config API groups from served APIs.

### Retrieving Configuration

The primary mechanism for retrieving static configuration should be
deserialization from files. For the majority of components (with the possible
exception of the kubelet, see
[here](https://github.com/kubernetes/kubernetes/pull/29459)), these files will
be source from the configmap API and managed by the kubelet. Reliability of
this mechanism is predicated on kubelet checkpointing of pod dependencies.


### Structuring Configuration

Group related options into distinct objects and subobjects. Instead of writing:


```yaml
kind: KubeProxyConfiguration
apiVersion: kubeproxy.config.k8s.io/v1beta3
ipTablesSyncPeriod: 2
ipTablesConntrackHashSize: 2
ipTablesConntrackTableSize: 2
```

Write:

```yaml
kind: KubeProxyConfiguration
apiVersion: kubeproxy.config.k8s.io/v1beta3
ipTables:
  syncPeriod: 2
  conntrack:
    hashSize: 2
    tableSize: 2
```

We should avoid passing around full configuration options to deeply constructed
modules. For example, instead of calling NewSomethingController in the
controller-manager with the full controller-manager config, group relevant
config into a subobject and only pass the subobject. We should expose the
smallest possible necessary configuration to the SomethingController.


### Handling Different Types Of Configuration

Above in "Type Of Configuration" we introduce a few ways to partition
configuration options. Environment information, request context depending
configuration, feature gates, and static configuration should be avoided if at
all possible using a configuration alternative. We should maintain separate
objects along these partitions and consider retrieving these configurations
from separate source (i.e. files). For example: kubeconfig (which falls into
the bootstrapping category) should not be part of the main config option (nor
should the filepath to the kubeconfig), per-instance config should be stored
separately from shared config. This allows for composition and obviates the
need for layering/templating solutions.


### In-Process Representation Of Configuration

We should separate structs for flags, serializable config, and runtime config.

1.  Structs for flags should have enough information for the process startup to retrieve its full configuration. Examples include: path the kubeconfig, path to configuration file, namespace and name of configmap to use for configuration.
1.  Structs for serializable configuration: This struct contains the full set of options in a serializable form (e.g. to represent an ip address instead of `net.IP`, use `string`). This is the struct that is versioned and serialized to disk using API machinery.
1.  Structs for runtime: This struct holds data in the most appropriate format for execution. This field can hold non-serializable types (e.g. have a `kubeClient` field instead of a `kubeConfig` field, store ip addresses as `net.IP`).

The flag struct is transformed into the configuration struct which is
transformed into the runtime struct.


### Migrating Away From Flags

Migrating to component configuration can happen incrementally (per component).
By versioning each component's API group separately, we can allow each API
group to advance to beta and GA independently. APIs should be approved by
component owners and reviewers familiar with the component configuration
conventions. We can incentivize operators to migrate away from flags by making
new configuration options only available through the component configuration
APIs.

# Caveats

Proposed are not laws but guidelines and as such we've favored completeness
over consistency. There will thus be need for exceptions.

1.  Components (especially those that are not self hosted such as the kubelet) will require custom rollout strategies of new config.
1.  Pod checkpointing by kubelet would allow this strategy to be simpler to make reliable.


# Miscellaneous Consideration

1.  **This document takes intentionally a very zealous stance against configuration.** Often configuration alternatives are not possible in Kubernetes as they are in proprietary software because Kubernetes has to run in diverse environments, with diverse users, managed by diverse operators.
1.  More frequent releases of kubernetes would make "skipping the config knob" more enticing because fixing a bad guess at a const wouldn't take O(4 months) best case to rollout. Factoring in our support for old versions, it takes closer to a year.
1.  Self-hosting resolves much of the distribution issue (except for maybe the Kubelet) but reliability is predicated on to-be-implemented features such as kubelet checkpointing of pod dependencies and sound operational practices such as incremental rollout of new configuration using Deployments/DaemonSets.
1.  Validating config is hard. Fatal logs lead to crash loops and error logs are ignored. Both options are suboptimal.
1.  Configuration needs to be updatable when components are down.
1.  Naming style guide:
    1.  No negatives, e.g. prefer --enable-foo over --disable-foo
    1.  Use the active voice
1.  We should actually enforce deprecation. Can we have a test that fails when a comment exists beyond its deadline to be removed? See [#44248](https://github.com/kubernetes/kubernetes/issues/44248)
1.  Use different implementations of the same interface rather than if statements to toggle features. This makes deprecation and deletion easy, improving maintainability.
1.  How does the proposed solution meet the requirements? Which desired qualities are missed?
1.  Configuration changes should trigger predictable and reproducible actions. From a given system state and a given component configuration, we should be able to simulate the actions that the system will take.
