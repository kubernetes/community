Status: Draft

Created: 2018-04-09 / Last updated: 2019-01-09

Authors: bsalamat, misterikkit

--------------------------------------------------------------------------------

<!-- toc -->

- [SUMMARY](#summary)
- [OBJECTIVE](#objective)
- [BACKGROUND](#background)
- [OVERVIEW](#overview)
- [DETAILED DESIGN](#detailed-design)
  * [Scheduling Cycle](#scheduling-cycle)
  * [Extension points](#extension-points)
    + [Queue sort](#queue-sort)
    + [Pre-filter](#pre-filter)
    + [Filter](#filter)
    + [Post-filter](#post-filter)
    + [Scoring](#scoring)
    + [Normalize scoring](#normalize-scoring)
    + [Reserve](#reserve)
    + [Permit](#permit)
    + [Pre-bind](#pre-bind)
    + [Bind](#bind)
    + [Post-bind](#post-bind)
    + [Un-reserve](#un-reserve)
  * [Plugin API](#plugin-api)
    + [PluginContext](#plugincontext)
    + [PluginHandle](#pluginhandle)
    + [Plugin Registration](#plugin-registration)
  * [Plugin Lifecycle](#plugin-lifecycle)
    + [Initialization](#initialization)
    + [Concurrency](#concurrency)
  * [Configuring Plugins](#configuring-plugins)
    + [Enable/Disable](#enabledisable)
    + [Change Evaluation Order](#change-evaluation-order)
    + [Optional Args](#optional-args)
    + [Backward compatibility](#backward-compatibility)
- [USE CASES](#use-cases)
  * [Coscheduling](#coscheduling)
  * [Dynamic Resource Binding](#dynamic-resource-binding)
  * [Custom Scheduler Plugins (out of tree)](#custom-scheduler-plugins-out-of-tree)
- [DEVELOPMENT PLAN](#development-plan)

<!-- tocstop -->

# SUMMARY

This document describes the Kubernetes Scheduling Framework. The scheduling
framework is a new set of "plugin" APIs being added to the existing Kubernetes
Scheduler. Plugins are compiled into the scheduler, and these APIs allow many
scheduling features to be implemented as plugins, while keeping the scheduling
"core" simple and maintainable.

*Note: Previous versions of this document proposed replacing the existing
scheduler with a new implementation.*

# OBJECTIVE

-   Make scheduler more extendable.
-   Make scheduler core simpler by moving some of its features to plugins.
-   Propose extension points in the framework.
-   Propose a mechanism to receive plugin results and continue or abort based on
    the received results.
-   Propose a mechanism to handle errors and communicate them with plugins.

# BACKGROUND

Many features are being added to the Kubernetes Scheduler. They keep making the
code larger and the logic more complex. A more complex scheduler is harder to
maintain, its bugs are harder to find and fix, and those users running a custom
scheduler have a hard time catching up and integrating new changes. The current
Kubernetes scheduler provides [webhooks to extend](./scheduler_extender.md) its
functionality. However, these are limited in a few ways:

1.  The number of extension points are limited: "Filter" extenders are called
    after default predicate functions. "Prioritize" extenders are called after
    default priority functions. "Preempt" extenders are called after running
    default preemption mechanism. "Bind" verb of the extenders are used to bind
    a Pod. Only one of the extenders can be a binding extender, and that
    extender performs binding instead of the scheduler. Extenders cannot be
    invoked at other points, for example, they cannot be called before running
    predicate functions.
1.  Every call to the extenders involves marshaling and unmarshalling JSON.
    Calling a webhook (HTTP request) is also slower than calling native
    functions.
1.  It is hard to inform an extender that scheduler has aborted scheduling of a
    Pod. For example, if an extender provisions a cluster resource and scheduler
    contacts the extender and asks it to provision an instance of the resource
    for the pod being scheduled and then scheduler faces errors scheduling the
    pod and decides to abort the scheduling, it will be hard to communicate the
    error with the extender and ask it to undo the provisioning of the resource.
1.  Since current extenders run as a separate process, they cannot use
    scheduler's cache. They must either build their own cache from the API
    server or process only the information they receive from the default
    scheduler.

The above limitations hinder building high performance and versatile scheduler
features. We would ideally like to have an extension mechanism that is fast
enough to allow existing features to be converted into plugins, such as
predicate and priority functions. Such plugins will be compiled into the
scheduler binary. Additionally, authors of custom schedulers can compile a
custom scheduler using (unmodified) scheduler code and their own plugins.

# OVERVIEW

The Scheduling Framework defines new extension points and Go APIs in the
Kubernetes Scheduler for use by "plugins". Plugins add scheduling behaviors to
the scheduler, and are included at compile time. The scheduler's ComponentConfig
will allow plugins to be enabled, disabled, and reordered. Custom schedulers can
write their plugins "[out-of-tree](#custom-scheduler-plugins-out-of-tree)"and
compile a scheduler binary with their own plugins included.

***Non-goals***

-   Solve all scheduler limitations, although we would like to ensure that the
    new framework allows us to address known limitations in the future.
-   Provide implementation details of plugins and call-back functions, such as
    all of their arguments and return values.

# DETAILED DESIGN

## Scheduling Cycle

The main loop of the scheduler is referred to as a "scheduling cycle". Each
cycle covers the complete process of assigning one pod to a node (or determining
that the pod cannot be scheduled). Multiple scheduling cycles are started
serially, but some parts may run concurrently. (See [Concurrency](#concurrency))

## Extension points

The following picture shows the scheduling cycle of a pod and the extension
points that the scheduling framework exposes. In this picture "Filter" is
equivalent to "Predicate" and "Scoring" is equivalent to "Priority function".
Plugins are registered to be called at one or more of these extension points. In
the following sections we describe each extension point in the same order they
are called in a scheduling cycle.

One plugin may register at multiple extension points to perform more complex or
stateful tasks.

![image](images/scheduling-framework-extensions.png)

### Queue sort

These plugins are used to sort pods in the scheduling queue. A queue sort plugin
essentially will provide a "less(pod1, pod2)" function. Only one queue sort
plugin may be enabled at a time.

### Pre-filter

These plugins are used to pre-process info about the pod, or to check certain
conditions that the cluster or the pod must meet. If a pre-filter plugin returns
an error, the scheduling cycle is aborted. Pre-filter plugins are called
serially within a scheduling cycle.

### Filter

These plugins are used to filter out nodes that cannot run the Pod. For each
node, the scheduler will call filter plugins in their configured order. If any
filter plugin marks the node as infeasible, the remaining plugins will not be
called for that node. Nodes may be evaluated concurrently.

### Post-filter

This is an informational extension point. Plugins can will be called with a list
of nodes that passed the filtering phase. A plugin may use this data to update
internal state or to generate logs/metrics.

**Note:** Plugins wishing to perform "pre-scoring" work should use the
post-filter extension point.

### Scoring

These plugins are used to rank nodes that have passed the filtering phase. The
scheduler will call each scoring plugin for each node. There will be a well
defined range of integers representing the minimum and maximum scores. After the
[normalize scoring](#normalize-scoring) phase, the scheduler will combine node
scores from all plugins according to the configured plugin weights.

If a scoring plugin returns an error, the scheduler will treat it as a zero
score.

### Normalize scoring

These plugins are used to modify scores before the scheduler computes a final
ranking of Nodes. A plugin that registers for this extension point will be
called with the [scoring](#scoring) results from the same plugin. This is called
once per plugin per scheduling cycle.

For example, suppose a plugin `BlinkingLightScorer` ranks Nodes based on how
many blinking lights they have.

```go
func ScoreNode(_ *v1.Pod, n *v1.Node) (int, error) {
   return getBlinkingLightCount(n)
}
```

However, the maximum count of blinking lights may be small compared to
`NodeScoreMax`. To fix this, `BlinkingLightScorer` should also register for this
extension point.

```go
func NormalizeScores(scores map[string]int) {
   highest := 0
   for _, score := range scores {
      highest = max(highest, score)
   }
   for node, score := range scores {
      scores[node] = scores*NodeScoreMax/highest
   }
}
```

If any normalize-scoring plugin returns an error, the scheduling cycle is
aborted.

**Note:** Plugins wishing to perform "pre-reserve" work should use the
normalize-scoring extension point.

### Reserve

This is an informational extension point. Plugins which maintain runtime state
(aka "stateful plugins") should use this extension point to be notified by the
scheduler when resources on a node are being reserved for a given Pod. This
happens before the scheduler actually binds the pod to the Node, and it exists
to prevent race conditions while the scheduler waits for the bind to succeed.

Once a pod is in the reserved state, it will either trigger
[Un-reserve](#un-reserve) plugins (on failure) or [Post-bind](#post-bind)
plugins (on success).

*Note: This concept used to be referred to as "assume".*

### Permit

These plugins are used to prevent or delay the binding of a Pod. A permit plugin
can do one of three things.

1.  **approve** \
    Once all permit plugins approve a pod, it is sent for binding.

1.  **deny** \
    If any permit plugin denies a pod, it is returned to the scheduling queue.
    This will trigger [Un-reserve](#un-reserve) plugins.

1.  **wait** (with a timeout) \
    If a permit plugin returns "wait", then the pod is kept in the permit phase
    until a [plugin approves it](#pluginhandle). If a timeout occurs, **wait**
    becomes **deny** and the pod is returned to the scheduling queue, triggering
    [un-reserve](#un-reserve) plugins.

**Approving a pod binding**

While any plugin can receive the list of reserved pod from the cache and approve
them (see [`PluginHandle`](#pluginhandle)) we expect only the permit plugins to
approve binding of reserved Pods that are in "waiting" state. Once a pod is
approved, it is sent to the pre-bind phase.

### Pre-bind

These plugins are used to perform any work required before a pod is bound. For
example, a pre-bind plugin may provision a network volume and mount it on the
target node before allowing the pod to run there.

If any pre-bind plugin returns an error, the pod is [rejected](#un-reserve) and
returned to the scheduling queue.

### Bind

These plugins are used to bind a pod to a Node. Bind plugins will not be called
until all pre-bind plugins. Each bind plugin is called in the configured order.
A bind plugin may choose whether or not to handle the given Pod. If a bind
plugin chooses to handle a Pod, **the remaining bind plugins are skipped**.

### Post-bind

This is an informational extension point. Post-bind plugins are called after a
pod is successfully bound. This is the end of a scheduling cycle, and can be
used to clean up associated resources.

### Un-reserve

This is an informational extension point. If a pod was reserved and then
rejected in a later phase, then un-reserve plugins will be notified. Un-reserve
plugins should clean up state associated with the reserved Pod.

Plugins that use this extension point usually should also use
[Reserve](#reserve).

## Plugin API

There are two steps to the plugin API. First, plugins must register and get
configured, then they use the extension point interfaces. Extension point
interfaces have the following form.

```go
type Plugin interface {
   Name() string
}

type QueueSortPlugin interface {
   Plugin
   Less(*v1.Pod, *v1.Pod) bool
}

type PreFilterPlugin interface {
   Plugin
   PreFilter(PluginContext, *v1.Pod) error
}

// ...
```

### PluginContext

Most* plugin functions will be called with a `PluginContext` argument. A
`PluginContext` represents the current scheduling cycle.

A `PluginContext` provides read-only APIs for accessing the scheduler's cache of
cluster state. This is the preferred way for plugins to iterate over nodes,
iterate over pods on one node, check available resources, and other tasks. The
scheduler will provide a consistent view of the cluster through these APIs, even
if the data is a little stale. Since two scheduling cycles can overlap in time,
plugins should not assume that they will see the same data from two different
`PluginContext`s.

The `PluginContext` also provides an API similar to
[`context.WithValue`](https://godoc.org/context#WithValue) that can be used to
pass data between plugins at different extension points. Multiple plugins can
share the state or communicate via this mechanism. The state is preserved only
during a single scheduling cycle. It is worth noting that plugins are assumed to
be **trusted**. The scheduler does not prevent one plugin from accessing or
modifying another plugin's state.

\* *The only exception is for [queue sort](#queue-sort) plugins.*

**WARNING**: The data available through a `PluginContext` is not valid after a
scheduling cycle ends, and plugins should not hold references to that data
longer than necessary.

### PluginHandle

While the `PluginContext` provides APIs relevant to a single scheduling cycle,
the `PluginHandle` provides APIs relevant to the lifetime of a plugin.
Specifically, `PluginHandle` provides a client (`kubernetes.Interface`) and
`SharedInformerFactory`. The handle will also provide APIs to list and approve
or reject [waiting pods](#permit).

**WARNING**: `PluginHandle` provides access to both the kubernetes API server
and the scheduler's internal cache. The two are **not guaranteed to be in sync**
and extreme care should be taken when writing a plugin that uses data from both
of them.

Providing plugins access to the API server is necessary to implement useful
features, especially when those features consume object types that the scheduler
does not normally consider. Providing a `SharedInformerFactory` allows plugins
to share caches safely.

### Plugin Registration

Each plugin must define a constructor and add it to the hard-coded registry. For
more information about constructor args, see [Optional Args](#optional-args).

Example:

```go
type PluginFactory = func(json.RawMessage, PluginHandle) (Plugin, error)

type Registry map[string]PluginFactory

func NewRegistry() Registry {
   return Registry{
      fooplugin.Name: fooplugin.New,
      barplugin.Name: barplugin.New,
      // New plugins are registered here.
   }
}
```

It is also possible to add plugins to a `Registry` object and inject that into a
scheduler. See [Custom Scheduler Plugins](#custom-scheduler-plugins-out-of-tree)

## Plugin Lifecycle

### Initialization

There are two steps to plugin initialization. First,
[plugins are registered](#plugin-registration). Second, the scheduler uses its
configuration to decide which plugins to instantiate. If a plugin registers for
multiple extension points, *it is instantiated only once*.

When a plugin is instantiated, it is passed [config args](#optional-args) and a
[`PluginHandle`](#pluginhandle).

### Concurrency

There are two types of concurrency that plugin writers should consider. A plugin
might be invoked several times concurrently when evaluating multiple nodes, and
a plugin may be called concurrently from *different
[scheduling cycles](#scheduling-cycle)*.

In the main thread of the scheduler, only one scheduling cycle is processed at a
time. Any extension point up to and including [reserve](#reserve) will be
finished before the next scheduling cycle begins*. After the reserve phase, the
[permit](#permit) and [bind](#bind) phases are executed asynchronously. This
means that a plugin could be called concurrently from two different scheduling
cycles, provided that at least one of the calls is to an extension point after
reserve. Stateful plugins should take care to handle these situations.

Finally, [un-reserve](#un-reserve) plugins may be called from either the Permit
thread or the Bind thread, depending on how the pod was rejected.

\* *The queue sort extension point is a special case. It is not part of a
scheduling cycle and may be called concurrently for many pod pairs.*

![image](images/scheduling-framework-threads.png)

## Configuring Plugins

The scheduler's component configuration will allow for plugins to be enabled,
disabled, or otherwise configured. Plugin configuration is separated into two
parts.

1.  A list of enabled plugins for each extension point (and the order they
    should run in). If one of these lists is omitted, the default list will be
    used.
1.  An optional set of custom plugin arguments for each plugin. Omitting config
    args for a plugin is equivalent to using the default config for that plugin.

The plugin configuration is organized by extension points. A plugin that
registers with multiple points must be included in each list.

```go
type KubeSchedulerConfiguration struct {
    // ... other fields
    Plugins      Plugins
    PluginConfig []PluginConfig
}

type Plugins struct {
    QueueSort      []Plugin
    PreFilter      []Plugin
    Filter         []Plugin
    PostFilter     []Plugin
    Score          []Plugin
    NormalizeScore []Plugin
    Reserve        []Plugin
    Permit         []Plugin
    PreBind        []Plugin
    Bind           []Plugin
    PostBind       []Plugin
    UnReserve      []Plugin
}

type Plugin struct {
    Name   string
    Weight int // Only valid for Score plugins
}

type PluginConfig struct {
    Name string
    Args json.RawMessage
}
```

Example:

```json
{
  "plugins": {
    "preFilter": [
      {
        "name": "PluginA"
      },
      {
        "name": "PluginB"
      },
      {
        "name": "PluginC"
      }
    ],
    "score": [
      {
        "name": "PluginA",
        "weight": 30
      },
      {
        "name": "PluginX"
      },
      {
        "name": "PluginY"
      }
    ]
  },
  "pluginConfig": [
    {
      "name": "PluginX",
      "args": {
        "favorite_color": "#326CE5",
        "favorite_number": 7,
        "thanks_to": "thockin"
      }
    }
  ]
}
```

### Enable/Disable

When specified, the list of plugins for a particular extension point are the
only ones enabled. If an extension point is omitted from the config, then the
default set of plugins is used for that extension point.

### Change Evaluation Order

When relevant, plugin evaluation order is specified by the order the plugins
appear in the configuration. A plugin that registers for multiple extension
points can have different ordering at each extension point.

### Optional Args

Plugins may receive arguments from their config with arbitrary structure.
Because one plugin may appear in multiple extension points, the config is in a
separate list of `PluginConfig`.

For example,

```json
{
   "name": "ServiceAffinity",
   "args": {
      "LabelName": "app",
      "LabelValue": "mysql"
   }
}
```

```go
func NewServiceAffinity(args json.RawMessage, h PluginHandle) (Plugin, error) {
    var config struct {
        LabelName, LabelValue string
    }
    if err := json.Unmarshal(args, &config); err != nil {
        return nil, errors.Wrap(err, "could not parse args")
    }
    //...
}
```

### Backward compatibility

The current `KubeSchedulerConfiguration` kind has `apiVersion:
kubescheduler.config.k8s.io/v1alpha1`. This new config format will be either
`v1alpha2` or `v1beta1`. When a newer version of the scheduler parses a
`v1alpha1`, the "policy" section will be used to construct an equivalent plugin
configuration.

*Note: Moving `KubeSchedulerConfiguration` to `v1` is outside the scope of this
design, but see also
https://github.com/kubernetes/enhancements/blob/master/keps/sig-cluster-lifecycle/0032-create-a-k8s-io-component-repo.md
and https://github.com/kubernetes/community/pull/3008*

# USE CASES

These are just a few examples of how the scheduling framework can be used.

## Coscheduling

Functionality similar to
[kube-batch](https://github.com/kubernetes-sigs/kube-batch) (sometimes called
"gang scheduling") could be implemented as a plugin. For pods in a batch, the
plugin would "accumulate" pods in the [permit](#permit) phase by using the
"wait" option. Because the permit stage happens after [reserve](#reserve),
subsequent pods will be scheduled as if the waiting pod is using those
resources. Once enough pods from the batch are waiting, they can all be
approved.

## Dynamic Resource Binding

[Topology-Aware Volume Provisioning](https://kubernetes.io/blog/2018/10/11/topology-aware-volume-provisioning-in-kubernetes/)
can be (re)implemented as a plugin that registers for [filter](#filter) and
[pre-bind](#pre-bind) extension points. At the filtering phase, the plugin can
ensure that the pod will be scheduled in a zone which is capable of provisioning
the desired volume. Then at the pre-bind phase, the plugin can provision the
volume before letting scheduler bind the pod.

## Custom Scheduler Plugins (out of tree)

The scheduling framework allows people to write custom, performant scheduler
features without forking the scheduler's code. To accomplish this, developers
just need to write their own `main()` wrapper around the scheduler. Because
plugins must be compiled with the scheduler, writing a wrapper around `main()`
is necessary in order to avoid modifying code in `vendor/k8s.io/kubernetes`.

```go
import (
   "k8s.io/kubernetes/pkg/scheduler/plugins"
   scheduler "k8s.io/kubernetes/cmd/kube-scheduler/app"
)

func main() {
   registry := plugins.NewRegistry()
   registry.Add("MyPlugin", NewMyPlugin)
   scheduler.Main(registry)
}
```

*Note: The above code is an example, and might not match the implemented API.*

The custom plugin would be enabled in the scheduler config.

```json
{
   "name": "MyPlugin"
}
```

# DEVELOPMENT PLAN

TODO: write down milestones and target releases, and a plan for how we will
gracefully move to the new system
