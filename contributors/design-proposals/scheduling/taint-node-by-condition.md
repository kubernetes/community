# Taints Node according to NodeConditions

@k82cn, @gmarek, @jamiehannaford,  Jul 15, 2017

## Relevant issues:

* https://github.com/kubernetes/kubernetes/issues/42001
* https://github.com/kubernetes/kubernetes/issues/45717

## Motivation
In kubernetes 1.8 and before, there are six Node Conditions, each with three possible values: True, False or Unknown. Kubernetes components modify and check those node conditions without any consideration to pods and their specs. For example, the scheduler will filter out all nodes whose NetworkUnavailable condition is True, meaning that pods on the host network can not be scheduled to those nodes, even though a user might want that.  The motivation of this proposal is to taint Nodes based on certain conditions, so that other components can leverage Tolerations for more advanced scheduling.

## Functional Detail
Currently (1.8 and before), the conditions of nodes are updated by the kubelet and Node Controller. The kubelet updates the value to either True or False according to the node’s status. If the kubelet did not update the value, the Node Controller will set the value to Unknown after a specific grace period.

In addition to this, with taint-based-eviction, the Node Controller already taints nodes with either NotReady and Unreachable if certain conditions are met. In this proposal, the Node Controller will use additional taints on Nodes. The new taints are described below:

| ConditionType      | Condition Status   |Effect        | Key      |
| ------------------ | ------------------ | ------------ | -------- |
|Ready               |True                | -            | |
|                    |False               | NoExecute    | node.kubernetes.io/not-ready           |
|                    |Unknown             | NoExecute    | node.kubernetes.io/unreachable         |
|OutOfDisk           |True                | NoSchedule   | node.kubernetes.io/out-of-disk         |
|                    |False               | -            | |
|                    |Unknown             | -            | |
|MemoryPressure      |True                | NoSchedule   | node.kubernetes.io/memory-pressure     |
|                    |False               | -            | |
|                    |Unknown             | -            | |
|DiskPressure        |True                | NoSchedule   | node.kubernetes.io/disk-pressure       |
|                    |False               | -            | |
|                    |Unknown             | -            | |
|NetworkUnavailable  |True                | NoSchedule   | node.kubernetes.io/network-unavailable |
|                    |False               | -            | |
|                    |Unknown             | -            | |
|PIDPressure         |True                | NoSchedule   | node.kubernetes.io/pid-pressure        |
|                    |False               | -            | |
|                    |Unknown             | -            | |

For example, if a CNI network is not detected on the node (e.g. a network is unavailable), the Node Controller will taint the node with `node.kubernetes.io/network-unavailable=:NoSchedule`. This will then allow users to add a toleration to their `PodSpec`, ensuring that the pod can be scheduled to this node if necessary. If the kubelet did not update the node’s status after a grace period, the Node Controller will only taint the node with `node.kubernetes.io/unreachable`; it will not taint the node with any unknown condition.

