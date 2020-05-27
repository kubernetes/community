# Scheduler extender

There are three ways to add new scheduling rules (predicates and priority
functions) to Kubernetes: (1) by adding these rules to the scheduler and
recompiling, [described here](/contributors/devel/sig-scheduling/scheduler.md),
(2) implementing your own scheduler process that runs instead of, or alongside
of, the standard Kubernetes scheduler, (3) implementing a "scheduler extender"
process that the standard Kubernetes scheduler calls out to as a final pass when
making scheduling decisions.

This document describes the third approach. This approach is needed for use
cases where scheduling decisions need to be made on resources not directly
managed by the standard Kubernetes scheduler. The extender helps make scheduling
decisions based on such resources. (Note that the three approaches are not
mutually exclusive.)

When scheduling a pod, the extender allows an external process to filter and
prioritize nodes. Two separate http/https calls are issued to the extender, one
for "filter" and one for "prioritize" actions. If the pod cannot be scheduled,
the scheduler will try to preempt lower priority pods from nodes and send them
to extender "preempt" verb if configured. The extender can return a subset of
nodes and new victims back to the scheduler. Additionally, the extender can
choose to bind the pod to apiserver by implementing the "bind" action.

To use the extender, you must create a scheduler policy configuration file. The
configuration specifies how to reach the extender, whether to use http or https
and the timeout.

```go
// Holds the parameters used to communicate with the extender. If a verb is unspecified/empty,
// it is assumed that the extender chose not to provide that extension.
type ExtenderConfig struct {
	// URLPrefix at which the extender is available
	URLPrefix string `json:"urlPrefix"`
	// Verb for the filter call, empty if not supported. This verb is appended to the URLPrefix when issuing the filter call to extender.
	FilterVerb string `json:"filterVerb,omitempty"`
	// Verb for the preempt call, empty if not supported. This verb is appended to the URLPrefix when issuing the preempt call to extender.
	PreemptVerb string `json:"preemptVerb,omitempty"`
	// Verb for the prioritize call, empty if not supported. This verb is appended to the URLPrefix when issuing the prioritize call to extender.
	PrioritizeVerb string `json:"prioritizeVerb,omitempty"`
	// Verb for the bind call, empty if not supported. This verb is appended to the URLPrefix when issuing the bind call to extender.
	// If this method is implemented by the extender, it is the extender's responsibility to bind the pod to apiserver.
	BindVerb string `json:"bindVerb,omitempty"`
	// The numeric multiplier for the node scores that the prioritize call generates.
	// The weight should be a positive integer
	Weight int `json:"weight,omitempty"`
	// EnableHttps specifies whether https should be used to communicate with the extender
	EnableHttps bool `json:"enableHttps,omitempty"`
	// TLSConfig specifies the transport layer security config
	TLSConfig *ExtenderTLSConfig `json:"tlsConfig,omitempty"`
	// HTTPTimeout specifies the timeout duration for a call to the extender. Filter timeout fails the scheduling of the pod. Prioritize
	// timeout is ignored, k8s/other extenders priorities are used to select the node.
	HTTPTimeout time.Duration `json:"httpTimeout,omitempty"`
	// NodeCacheCapable specifies that the extender is capable of caching node information,
	// so the scheduler should only send minimal information about the eligible nodes
	// assuming that the extender already cached full details of all nodes in the cluster
	NodeCacheCapable bool `json:"nodeCacheCapable,omitempty"`
	// ManagedResources is a list of extended resources that are managed by
	// this extender.
	// - A pod will be sent to the extender on the Filter, Prioritize and Bind
	//   (if the extender is the binder) phases iff the pod requests at least
	//   one of the extended resources in this list. If empty or unspecified,
	//   all pods will be sent to this extender.
	// - If IgnoredByScheduler is set to true for a resource, kube-scheduler
	//   will skip checking the resource in predicates.
	// +optional
	ManagedResources []ExtenderManagedResource `json:"managedResources,omitempty"`
	// Ignorable specifies if the extender is ignorable, i.e. scheduling should not
	// fail when the extender returns an error or is not reachable.
	Ignorable bool `json:"ignorable,omitempty"
}
```

A sample scheduler policy file with extender configuration:

```json
{
  "predicates": [
    {
      "name": "HostName"
    },
    {
      "name": "MatchNodeSelector"
    },
    {
      "name": "PodFitsResources"
    }
  ],
  "priorities": [
    {
      "name": "LeastRequestedPriority",
      "weight": 1
    }
  ],
  "extenders": [
    {
      "urlPrefix": "http://127.0.0.1:12345/api/scheduler",
      "filterVerb": "filter",
      "enableHttps": false
    }
  ]
}
```

Arguments passed to the FilterVerb endpoint on the extender are the set of nodes
filtered through the k8s predicates and the pod. Arguments passed to the
PrioritizeVerb endpoint on the extender are the set of nodes filtered through
the k8s predicates and extender predicates and the pod.

```go
// ExtenderArgs represents the arguments needed by the extender to filter/prioritize
// nodes for a pod.
type ExtenderArgs struct {
	// Pod being scheduled
	Pod   api.Pod      `json:"pod"`
	// List of candidate nodes where the pod can be scheduled
	Nodes api.NodeList `json:"nodes"`
	// List of candidate node names where the pod can be scheduled; to be
	// populated only if Extender.NodeCacheCapable == true
	NodeNames *[]string
}
```

The "filter" call returns a list of nodes (schedulerapi.ExtenderFilterResult).

```go
// FailedNodesMap represents the filtered out nodes, with node names and failure messages
type FailedNodesMap map[string]string

// ExtenderFilterResult represents the results of a filter call to an extender
type ExtenderFilterResult struct {
	// Filtered set of nodes where the pod can be scheduled; to be populated
	// only if Extender.NodeCacheCapable == false
	Nodes *v1.NodeList
	// Filtered set of nodes where the pod can be scheduled; to be populated
	// only if Extender.NodeCacheCapable == true
	NodeNames *[]string
	// Filtered out nodes where the pod can't be scheduled and the failure messages
	FailedNodes FailedNodesMap
	// Error message indicating failure
	Error string
}
```

The "prioritize" call returns priorities for each node (schedulerapi.HostPriorityList).

```go
// HostPriority represents the priority of scheduling to a particular host, higher priority is better.
type HostPriority struct {
	// Name of the host
	Host string
	// Score associated with the host
	Score int64
}

// HostPriorityList declares a []HostPriority type.
type HostPriorityList []HostPriority
```

The "filter" call may prune the set of nodes based on its predicates. Scores
returned by the "prioritize" call are added to the k8s scores (computed through
its priority functions) and used for final host selection.

The "bind" call is used to delegate the bind of a pod to a node to the extender. It can
be optionally implemented by the extender. When it is implemented, it is the extender's
responsibility to issue the bind call to the apiserver. Pod name, namespace, UID and Node
name are passed to the extender.

```go
// ExtenderBindingArgs represents the arguments to an extender for binding a pod to a node.
type ExtenderBindingArgs struct {
	// PodName is the name of the pod being bound
	PodName string
	// PodNamespace is the namespace of the pod being bound
	PodNamespace string
	// PodUID is the UID of the pod being bound
	PodUID types.UID
	// Node selected by the scheduler
	Node string
}

// ExtenderBindingResult represents the result of binding of a pod to a node from an extender.
type ExtenderBindingResult struct {
	// Error message indicating failure
	Error string
}
```

The "preempt" call makes the extender to return a subset of given nodes and new victims on the nodes. 

```go
// ExtenderPreemptionArgs represents the arguments needed by the extender to preempt pods on nodes.
type ExtenderPreemptionArgs struct {
	// Pod being scheduled
	Pod *v1.Pod
	// Victims map generated by scheduler preemption phase
	// Only set NodeNameToMetaVictims if Extender.NodeCacheCapable == true. Otherwise, only set NodeNameToVictims.
	NodeNameToVictims     map[string]*Victims
	NodeNameToMetaVictims map[string]*MetaVictims
}

// ExtenderPreemptionResult represents the result returned by preemption phase of extender.
type ExtenderPreemptionResult struct {
	NodeNameToMetaVictims map[string]*MetaVictims
}
```
