# Management Cluster - SIG Multicluster Position Statement

Author: Corentin Debains (**[@corentone](https://github.com/corentone)**), Google  
Last Edit: 2024/12/09  
Status: DRAFT

## Goal
To establish a standard definition for a central cluster that is leveraged by multicluster 
controllers to manage multicluster applications or features across an inventory of clusters.

## Context
Multicluster controllers have always needed a place to run. This may happen in external
proprietary control-planes but for more generic platforms, it has been natural for the
Kubernetes community to leverage a Kubernetes Cluster and the existing api-machinery
available. There has been a variety of examples of which we can quote ArgoCD, MultiKueue
or any of the Federation effort (Karmada, KubeAdmiral), all of them not-naming the "location"
where they run or not aligning on the name (Admin cluster, Hub Cluster, Manager Cluster...). 
The [ClusterInventory](https://github.com/kubernetes/enhancements/blob/master/keps/sig-multicluster/4322-cluster-inventory/README.md)
(ClusterProfile CRDs) is also the starting point for a lot of multicluster controllers and, 
being a CRD, it requires an api-machinery to host it.

## Definition

A (multicluster) management cluster is a Kubernetes cluster that acts as a 
control-plane for other Kubernetes clusters (named Workload Clusters to differentiate
them). It MUST have visibility over the available clusters and MAY have administrative
privileges over them. It SHOULD not be part of workload clusters to provide a better
security isolation, especially when it has any administrative privileges over them.
There MAY be multiple management clusters overseeing the same set of Workload Clusters
and it is left to the administrator to guarantee that they don't compete in their
management tasks. There SHOULD be a single clusterset managed by a management cluster.
Management clusters can be used for both control-plane or data-plane features.


### Rationale on some specific points of the definition

* Multiple management clusters: While it often makes sense to have a single "Brain" overseeing
 a Fleet of Clusters, there is a need for flexibility over the number of management clusters. To
 allow redundancy to improve reliability, to allow sharding of responsibility (for regionalized
 controllers), to allow for separation of functionality (security-enforcer management cluster vs
 config-delivery management cluster), to allow for migrations (from old management cluster to new
 management cluster) and likely more.
* Management cluster also being part of the workload-running Fleet: We do recommend that the
 management cluster(s) be isolated from the running Workload Fleet for security and management
 concerns. But there may be specific cases or applications that require to mix the two. For example,
 controllers that take a "leader-election" approach and want a smaller footprint.
