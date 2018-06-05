# The Kubernetes Scheduler

The Kubernetes scheduler runs as a process alongside the other master components such as the API server.
Its interface to the API server is to watch for Pods with an empty PodSpec.NodeName,
and for each Pod, it posts a binding indicating where the Pod should be scheduled.

## Exploring the code

We are dividing scheduler into three layers from high level:
- [cmd/kube-scheduler/scheduler.go](http://releases.k8s.io/HEAD/cmd/kube-scheduler/scheduler.go):
  This is the main() entry that does initialization before calling the scheduler framework.
- [pkg/scheduler/scheduler.go](http://releases.k8s.io/HEAD/pkg/scheduler/scheduler.go):
  This is the scheduler framework that handles stuff (e.g. binding) beyond the scheduling algorithm.
- [pkg/scheduler/core/generic_scheduler.go](http://releases.k8s.io/HEAD/pkg/scheduler/core/generic_scheduler.go):
  The scheduling algorithm that assigns nodes for pods.

## The scheduling algorithm

```
For given pod:

    +---------------------------------------------+
    |               Schedulable nodes:            |
    |                                             |
    | +--------+    +--------+      +--------+    |
    | | node 1 |    | node 2 |      | node 3 |    |
    | +--------+    +--------+      +--------+    |
    |                                             |
    +-------------------+-------------------------+
                        |
                        |
                        v
    +-------------------+-------------------------+

    Pred. filters: node 3 doesn't have enough resource

    +-------------------+-------------------------+
                        |
                        |
                        v
    +-------------------+-------------------------+
    |             remaining nodes:                |
    |   +--------+                 +--------+     |
    |   | node 1 |                 | node 2 |     |
    |   +--------+                 +--------+     |
    |                                             |
    +-------------------+-------------------------+
                        |
                        |
                        v
    +-------------------+-------------------------+

    Priority function:    node 1: p=2
                          node 2: p=5

    +-------------------+-------------------------+
                        |
                        |
                        v
            select max{node priority} = node 2
```

The Scheduler tries to find a node for each Pod, one at a time.
- First it applies a set of "predicates" to filter out inappropriate nodes. For example, if the PodSpec specifies resource requests, then the scheduler will filter out nodes that don't have at least that much resources available (computed as the capacity of the node minus the sum of the resource requests of the containers that are already running on the node).
- Second, it applies a set of "priority functions"
that rank the nodes that weren't filtered out by the predicate check. For example, it tries to spread Pods across nodes and zones while at the same time favoring the least (theoretically) loaded nodes (where "load" - in theory - is measured as the sum of the resource requests of the containers running on the node, divided by the node's capacity).
- Finally, the node with the highest priority is chosen (or, if there are multiple such nodes, then one of them is chosen at random). The code for this main scheduling loop is in the function `Schedule()` in [pkg/scheduler/core/generic_scheduler.go](http://releases.k8s.io/HEAD/pkg/scheduler/core/generic_scheduler.go)

### Predicates and priorities policies

Predicates are a set of policies applied one by one to filter out inappropriate nodes.
Priorities are a set of policies applied one by one to rank nodes (that made it through the filter of the predicates).
By default, Kubernetes provides built-in predicates and priorities policies documented in [scheduler_algorithm.md](scheduler_algorithm.md).
The predicates and priorities code are defined in [pkg/scheduler/algorithm/predicates/predicates.go](http://releases.k8s.io/HEAD/pkg/scheduler/algorithm/predicates/predicates.go) and [pkg/scheduler/algorithm/priorities](http://releases.k8s.io/HEAD/pkg/scheduler/algorithm/priorities/) , respectively.


## Scheduler extensibility

The scheduler is extensible: the cluster administrator can choose which of the pre-defined
scheduling policies to apply, and can add new ones.

### Modifying policies

The policies that are applied when scheduling can be chosen in one of two ways.
The default policies used are selected by the functions `defaultPredicates()` and `defaultPriorities()` in
[pkg/scheduler/algorithmprovider/defaults/defaults.go](http://releases.k8s.io/HEAD/pkg/scheduler/algorithmprovider/defaults/defaults.go).
However, the choice of policies can be overridden by passing the command-line flag `--policy-config-file` to the scheduler, pointing to a JSON file specifying which scheduling policies to use. See [examples/scheduler-policy-config.json](https://git.k8s.io/examples/staging/scheduler-policy-config.json) for an example
config file. (Note that the config file format is versioned; the API is defined in [pkg/scheduler/api](http://releases.k8s.io/HEAD/pkg/scheduler/api/)).
Thus to add a new scheduling policy, you should modify  [pkg/scheduler/algorithm/predicates/predicates.go](http://releases.k8s.io/HEAD/pkg/scheduler/algorithm/predicates/predicates.go) or add to the directory  [pkg/scheduler/algorithm/priorities](http://releases.k8s.io/HEAD/pkg/scheduler/algorithm/priorities/), and either register the policy in `defaultPredicates()` or `defaultPriorities()`, or use a policy config file.

