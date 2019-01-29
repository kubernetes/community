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
for "filter" and one for "prioritize" actions. Additionally, the extender can
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
	TLSConfig *client.TLSClientConfig `json:"tlsConfig,omitempty"`
	// HTTPTimeout specifies the timeout duration for a call to the extender. Filter timeout fails the scheduling of the pod. Prioritize
	// timeout is ignored, k8s/other extenders priorities are used to select the node.
	HTTPTimeout time.Duration `json:"httpTimeout,omitempty"`
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
}
```

The "filter" call returns a list of nodes (schedulerapi.ExtenderFilterResult). The "prioritize" call
returns priorities for each node (schedulerapi.HostPriorityList).

The "filter" call may prune the set of nodes based on its predicates. Scores
returned by the "prioritize" call are added to the k8s scores (computed through
its priority functions) and used for final host selection.

"bind" call is used to delegate the bind of a pod to a node to the extender. It can
be optionally implemented by the extender. When it is implemented, it is the extender's
responbility to issue the bind call to the apiserver. Pod name, namespace, UID and Node
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
```
