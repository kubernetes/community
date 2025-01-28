# Hub Cluster - SIG Multicluster Position Statement

Author: Corentin Debains (**[@corentone](https://github.com/corentone)**), Google  
Last Edit: 2025/01/25  
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
where they run or not aligning on the name (Admin cluster, Management Cluster, Command Cluster, Manager Cluster...). 
The [ClusterInventory](https://github.com/kubernetes/enhancements/blob/master/keps/sig-multicluster/4322-cluster-inventory/README.md)
(ClusterProfile CRDs) is also the starting point for a lot of multicluster controllers and, 
being a CRD, it requires an api-machinery to host it. The functionality of this cluster is also
defined in separation to what a "workload" cluster does, which is to run the business applications,
when hub runs infrastructure components.

## Definition

A (multicluster) hub cluster is a Kubernetes cluster that acts as a 
control-plane for other Kubernetes clusters (named Workload Clusters to differentiate
them). It MUST have the ClusterProfiles written on it MAY have access to the api, metrics or
workloads of the workload clusters and MAY have administrative privileges over them. It 
SHOULD not be part of workload clusters or running as mixed mode (workload and hub) to provide a better
security isolation, especially when it has any administrative privileges over them.
There MAY be multiple hub clusters overseeing the same set of Workload Clusters
and it is left to the administrator to guarantee that they don't compete in their
management tasks. There SHOULD be a single [clusterset](https://multicluster.sigs.k8s.io/api-types/cluster-set/)
managed by a hub cluster. Hub clusters can be used for multicluster controllers relative to platform-running features,
for example: managing the clusters, or application-running features, for example: scheduling business
applications dynamically.

### Rationale on some specific points of the definition

* Multiple hub clusters: While it often makes sense to have a single "Brain" overseeing
 a Fleet of Clusters, there is a need for flexibility over the number of hub clusters. To
 allow redundancy to improve reliability, to allow sharding of responsibility (for regionalized
 controllers), to allow for separation of functionality (security-enforcer hub cluster vs
 config-delivery hub cluster), to allow for migrations (from old hub cluster to new
 hub cluster) and likely more.
* Hub cluster also being part of the workload-running Fleet: We do recommend that the
 hub cluster(s) be isolated from the running Workload Fleet for security and hub
 concerns. But there may be specific cases or applications that require to mix the two. For example,
 controllers that take a "leader-election" approach and want a smaller footprint.
* Application-running features vs platform-running features: Hub clusters can runcontrollers
  that are catering to a "Platform" type of user, effectively using a central cluster to manage other clusters and
  other infrastructure. For example, centrally monitoring health of clusters of a clusterset. It can also run
  controllers that are helping run business applications globally. For example, having a definition of a multicluster
  application and scheduling replicas of the application to the different clusters of the clusterset.
  This means that access control to the hub cluster and permissions given to controllers on the hub
  clusters must be carefully designed.
