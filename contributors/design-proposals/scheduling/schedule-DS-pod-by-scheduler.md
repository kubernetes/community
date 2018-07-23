# Schedule DaemonSet Pods by default scheduler, not DaemonSet controller

[@k82cn](http://github.com/k82cn), Feb 2018, [#42002](https://github.com/kubernetes/kubernetes/issues/42002).

## Motivation

A DaemonSet ensures that all (or some) nodes run a copy of a pod. As nodes are added to the cluster, pods are added to them. As nodes are removed from the cluster, those pods are garbage collected. Normally, the machine that a pod runs on is selected by the Kubernetes scheduler; however, pods of DaemonSet are created and scheduled by DaemonSet controller who leveraged kube-scheduler’s predicates policy. That introduces the following issues:

* DaemonSet can not respect Node’s resource changes, e.g. more resources after other Pods exit ([#46935](https://github.com/kubernetes/kubernetes/issues/46935), [#58868](https://github.com/kubernetes/kubernetes/issues/58868))
* DaemonSet can not respect Pod Affinity and Pod AntiAffinity ([#29276](https://github.com/kubernetes/kubernetes/issues/29276))
* Duplicated logic to respect scheduler features, e.g. critical pods ([#42028](https://github.com/kubernetes/kubernetes/issues/42028)),  tolerant/taint
* Hard to debug why DaemonSet’s Pod is not created, e.g. not enough resources; it’s better to have a pending Pods with predicates’ event
* Hard to support preemption in different components, e.g. DS and default scheduler

After [discussions](https://docs.google.com/document/d/1v7hsusMaeImQrOagktQb40ePbK6Jxp1hzgFB9OZa_ew/edit#), SIG scheduling approved changing DaemonSet controller to create DaemonSet Pods and set their node-affinity and let them be scheduled by default scheduler. After this change, DaemonSet controller will no longer schedule DaemonSet Pods directly.

## Solutions

Before the discussion of solutions/options, there’s some requirements/questions on DaemonSet:

* **Q**: DaemonSet controller can make pods even if the network of node is unavailable, e.g. CNI network providers (Calico, Flannel),
Will this impact bootstrapping, such as in the case that a DaemonSet is being used to provide the pod network?

  **A**: This will be handled by supporting scheduling tolerating workloads on NotReady Nodes ([#45717](https://github.com/kubernetes/kubernetes/issues/45717)); after moving to check node’s taint, the DaemonSet pods will tolerate `NetworkUnavailable` taint. 

* **Q**: DaemonSet controller can make pods even if when the scheduler has not been started, which can help cluster bootstrap.

  **A**: As the scheduling logic is moved to default scheduler, the kube-scheduler must be started during cluster start-up.

* **Q**: Will this change/constrain update strategies, such as scheduling an updated pod to a node before the previous pod is gone?

  **A**: no, this will NOT change update strategies.

* **Q**: How would Daemons be integrated into Node lifecycle, such as being scheduled before any other nodes and/or remaining after all others are evicted? This isn't currently implemented, but was planned.

  **A**:  Similar to the other Pods; DaemonSet Pods only has attributes to make sure one Pod per Node, DaemonSet controller will create Pods based on node number (by considering ‘nodeSelector’).


Currently, pods of DaemonSet are created and scheduled by DaemonSet controller:

1. DS controller filter nodes by nodeSelector and scheduler’s predicates
2. For each node, create a Pod for it by setting spec.hostName directly; it’ll skip default scheduler

This option is to leverage NodeAffinity feature to avoid introducing scheduler’s predicates in DS controller:

1. DS controller filter nodes by nodeSelector, but does NOT check against scheduler’s predicates (e.g. PodFitHostResources)
2. For each node, DS controller creates a Pod for it with the following NodeAffinity
    ```yaml
    nodeAffinity:
      requiredDuringSchedulingIgnoredDuringExecution:
      - nodeSelectorTerms:
          matchExpressions:
          - key: kubernetes.io/hostname
            operator: in
            values:
            - dest_hostname
    ```
3. When sync Pods, DS controller will map nodes and pods by this NodeAffinity to check whether Pods are started for nodes
4. In scheduler, DaemonSet Pods will stay pending if scheduling predicates fail. To avoid this, an appropriate priority must
 be set to all critical DaemonSet Pods. Scheduler will preempt other pods to ensure critical pods were scheduled even when
 the cluster is under resource pressure.

## Reference

* [DaemonsetController can't feel it when node has more resources, e.g. other Pod exits](https://github.com/kubernetes/kubernetes/issues/46935)
* [DaemonsetController can't feel it when node recovered from outofdisk state](https://github.com/kubernetes/kubernetes/issues/45628)
* [DaemonSet pods should be scheduled by default scheduler, not DaemonSet controller](https://github.com/kubernetes/kubernetes/issues/42002)
* [NodeController should add NoSchedule taints and we should get rid of getNodeConditionPredicate()](https://github.com/kubernetes/kubernetes/issues/42001)
* [DaemonSet should respect Pod Affinity and Pod AntiAffinity](https://github.com/kubernetes/kubernetes/issues/29276)
* [Make DaemonSet respect critical pods annotation when scheduling](https://github.com/kubernetes/kubernetes/pull/42028)
