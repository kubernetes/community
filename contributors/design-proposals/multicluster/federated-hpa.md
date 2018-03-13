# Federated Pod Autoscaler

# Requirements &amp; Design Document

irfan.rehman@huawei.com, quinton.hoole@huawei.com

# Use cases

1 – Users can schedule replicas of same application, across the 
federated clusters, using replicaset (or deployment). 
Users however further might need to let the replicas be scaled 
independently in each cluster, depending on the current usage metrics 
of the replicas; including the CPU, memory and application defined 
custom metrics.

2 - As stated in the previous use case, a federation user schedules 
replicas of same application, into federated clusters and subsequently
creates a horizontal pod autoscaler targeting the object responsible for 
the replicas. User would want the auto-scaling to continue based on 
the in-cluster metrics, even if for some reason, there is an outage at 
federation level. User (or other users) should still be able to access 
the deployed application into all federated clusters. Further, if the 
load on the deployed app varies, the autoscaler should continue taking 
care of scaling the replicas for a smooth user experience.

3 - A federation that consists of an on-premise cluster and  a cluster 
running in a public cloud has a user workload (eg. deployment or rs) 
preferentially running in the on-premise cluster. However if there are 
spikes in the app usage, such that the capacity in the on-premise cluster 
is not sufficient, the workload should be able to get scaled beyond the 
on-premise cluster boundary and into the other clusters which are part 
of this federation.

Please refer to some additional use cases, which partly led to the derivation 
of the above use case, and are listed in the **glossary** section of this document.

# User workflow

User wants to schedule a set of common workload across federated clusters. 
He creates a replicaset or a deployment to schedule the workload (with or 
without preferences). The federation then distributes the replicas of the 
given workload into the federated clusters. As the user at this point is 
unaware of the exact usage metrics of the individual pods created in the 
federated clusters, he creates an HPA into the federation, providing metric 
parameters to be used in the scale request for a resource. It is now the 
responsibility of this HPA to monitor the relevant resource metrics and the 
scaling of the pods per cluster then is controlled by the associated HPA.

# Alternative approaches

## Design Alternative 1

Make the autoscaling resource available and implement support for 
horizontalpodautoscalers objects at federation. The HPA API resource 
will need to be exposed at the federation level, which can follow the 
version similar to one implemented in the latest k8s cluster release.

Once the HPA object is created at federation, the federation controller
creates and monitors a similar HPA object (partitioning the min and max values)
in each of the federated clusters. Based on the metadata in spec of the HPA
describing the scaleTargetRef, the HPA will be applied on the already existing
target objects. If the target object is not present in the cluster (either 
because, its not created until now, or deleted for some reason), the HPA will 
still exist but no action will be taken. The HPA&#39;s action will become 
applicable when the target object is created in the given cluster anytime in 
future. Also as stated already the  federation controller will need to partition 
the min and max values appropriately into the federated clusters among the HPA 
objects such that the total of min and that of max replicas satisfies the 
constraints specified by the user at federation. The point of control over the 
scaling of replicas will lie locally with the federated hpa controller. The 
federated controller will however watch the cluster local HPAs wrt current
replicas of the target objects and will do intelligent dynamic adjustments of 
min and max values of the HPA replicas across the clusters based on the run time 
conditions.

The federation controller by default will distribute the min and max replicas of the 
HPA equally among all clusters. The min values will first be distributed such that 
any cluster into which the replicas are distributed does not get a min replicas 
lesser than 1. This means that HPA can actually be created in lesser number of 
ready clusters then available in federation. Once this distribution happens, the 
max replicas of the hpa will be distributed across all those clusters into which 
the HPA needs to be created. The default distribution can be overridden using the 
annotations on the HPA object, very similar to the annotations on federated 
replicaset object as described 
[here](https://github.com/kubernetes/community/blob/master/contributors/design-proposals/federated-replicasets.md#federatereplicaset-preferences).

One of the points to note here is that, doing this brings a two point control on 
number of replicas of the target object, one by the federated target object (rs or 
deployment) and other by the hpa local to the federated cluster. Solution to which 
is discussed in the following section. Another additional note here is that, the
preferences would consider use of only minreplicas and maxreplicas in this phase
of implementation and weights will be discarded for this alternative design.

### Rebalancing of workload replicas and control over the same.

The current implementation of federated replicasets (and deployments) first 
distributes the replicas into underlying clusters and then monitors the status 
of the pods in each cluster. In case there are clusters which have active pods 
lesser than what federation reconciler desires, federation control plane will 
trigger creation of the missing pods (which federation considers missing), or 
in other case would trigger removal of pods, if the control plane considers that 
the given cluster has more pods than needed. This is something which counters 
the role of HPA in individual cluster. To handle this, the knowledge that HPA 
is active separately targeting this object has to be percolated to the federation 
control plane monitoring the individual replicas such that, the federation control
plane stops reconciling the replicas in the individual clusters. In other words 
the link between the HPA wrt to the corresponding objects will need to be 
maintained and if an HPA is active, other federation controllers (aka replicaset 
and deployment controllers) reconcile process, would stop updating and/or 
rebalancing the replicas in and across the underlying clusters. The reconcile 
of the objects (rs or deployment) would still continue, to handle the scenario 
of the object missing from any given federated cluster.
The mechanism to achieve this behaviour shall be as below:
 - User creates a workload object (for example rs) in federation.
 - User then creates an HPA object in federation (this step and the previous
 step can follow either order of creation).
 - The rs as an object will exist in federation control plane with or without
 the user preferences and/or cluster selection annotations.
 - The HPA controller will first evaluate which cluster(s) get the replicas
 and which don't (if any). This list of clusters will be a subset of the
 cluster selector already applied on the hpa object.
 - The HPA controller will apply this list on the federated rs object as the
 cluster selection annotation overriding the user provided preferences (if any).
 The control over the placement of workload replicas and the add on preferences
 will thus lie completely with the HPA objects. This is an important assumption
 that the user of these federated objects interacting with each other should be
 aware of; and if the user needs to place replicas in specific clusters, together
 with workload autoscaling he/she should apply these preferences on the HPA
 object. Any preferences applied on the workload object (rs or deployment) will
 be overridden.
 - The target workload object (for example rs) replicas will be kept unchanged
 in the cluster which already has the replicas, will be created with one replica
 if the particular cluster does not have the same and HPA calculation resulted
 in some replicas for that cluster and deleted from the clusters which has the
 replicas and the federated HPA calculations result in no replicas for that
 particular cluster.
 - The desired replicas per cluster as per the federated HPA dynamic rebalance
 mechanism, elaborated in the next section, will be set on individual clusters
 local HPA, which in turn will set the same on the target local object.

### Dynamic HPA min/max rebalance

The proposal in this section can be used to improve the distribution of replicas 
across the clusters such that there are more replicas in those clusters, where 
they are needed more. The federation hpa controller will monitor the status of 
the local HPAs in the federated clusters and update the min and/or max values 
set on the local HPAs as below (assuming that all previous steps are done and 
local HPAs in federated clusters are active):

1. At some point, one or more of the cluster HPA&#39;s hit the upper limit of their 
allowed scaling such that _DesiredReplicas == MaxReplicas_; Or more appropriately 
_CurrentReplicas == DesiredReplicas == MaxReplicas_.

2. If the above is observed the Federation HPA tries to transfer allocation 
of _MaxReplicas_ from clusters where it is not needed (_DesiredReplicas < MaxReplicas_) 
or where it cannot be used, e.g. due to capacity constraints 
(_CurrentReplicas < DesiredReplicas <= MaxReplicas_) to the clusters which have 
reached their upper limit (1 above).

3. It will be taken care that the _MaxReplica_ does not become lesser than _MinReplica_ 
in any of the clusters in this redistribution. Additionally if the usage of the same 
could be established, _MinReplicas_ can also be distributed as in 4 below.

4. An exactly similar approach can also be applied to _MinReplicas_ of the local HPAs, 
so as to reduce the min from those clusters, where  
_CurrentReplicas == DesiredReplicas == MinReplicas_ and the observed average resource
metric usage (on the HPA) is lesser then a given threshold, to those clusters, 
where the _DesiredReplicas > MinReplicas_.

However, as stated in 3 above, the approach of distribution will first be implemented
only for _MaxReplicas_ to establish it utility, before implementing the same for _MinReplicas_.

## Design Alternative 2

Same as the previous one, the API will need to be exposed at federation.

However, when the request to create HPA is sent to federation, federation controller 
will not create the HPA into the federated clusters. The HPA object will reside in the 
federation API server only. The federation controller will need to get a metrics 
client to each of the federated clusters and collect all the relevant metrics 
periodically from all those clusters. The federation controller will further calculate 
the current average metrics utilisation across all clusters (using the collected metrics) 
of the given target object and calculate the replicas globally to attain the target 
utilisation as specified in the federation HPA. After arriving at the target replicas, 
the target replica number is set directly on the target object (replicaset, deployment, ..) 
using its scales sub-resource at federation. It will be left to the actual target object 
controller (for example RS controller) to distribute the replicas accordingly into the 
federated clusters. The point of control over the scaling of replicas will lie completely 
with the federation controllers.

### Algorithm (for alternative 2)

Federated HPA (FHPA), from every cluster gets:

- ```avg_i``` average metric value (like CPU utilization) for all pods matching the 
deployment/rs selector.
- ```count_i``` number of replicas that were used to calculate the average.

To calculate the target number of replicas HPA calculates the sum of all metrics from
all clusters:

```sum(avg_i * count_i)``` and divides it by target metric value. The target replica 
count (validated against HPA min/max and thresholds) is set on Federated 
Deployment/replica set. So the deployment has the correct number of replicas 
(that should match the desired metric value) and provides all of the rebalancing/failover 
mechanisms.

Further, this can be expanded such that FHPA places replicas where they are needed the
most (in cluster that have the most traffic). For that FHPA would play with weights in
Federated Deployment. Each cluster will get the weight of ```100 * avg_i/sum(avg_i)```.
Weights hint Federated Deployment where to put replicas. But they are only hints so 
if placing a replica in the desired cluster is not possible then it will be placed elsewhere, 
what is probably better than not having the replica at all.

# Other Scenario

Other scenario, for example rolling updates (when user updates the deployment or RS),
recreation of the object (when user specifies the strategy as recreate while updating
the object), will continue to be handled the way they are handled in an individual k8s
cluster. Additionally there is a shortcoming in the current implementation of the
federated deployments rolling update. There is an existing proposal as part of the
[federated deployment design doc](https://github.com/kubernetes/community/pull/325).
Given it is implemented, the rolling updates for a federated deployment while a
federated HPA is active on the same object will also work fine.

# Conclusion

The design alternative 2 has the following major drawbacks, which are sufficient to
discard it as a probable implementation option:
- This option needs the federation control plane controller to collect metrics
data from each cluster, which is an overhead with increasing gravity of the problem
with increasing number of federated clusters, in a given federation.
- The monitoring and update of objects which are targeted by the federated HPA object
(when needed) for a particular federated cluster would stop if for whatever reasons
the network link between the federated cluster and federation control plane is severed.
A bigger problem can happen in case of an outage of the federation control plane
altogether.

In Design Alternative 1 the autoscaling of replicas will continue, even if a given
cluster gets disconnected from federation or in case of the federation control plane
outage. This would happen because the local HPAs with the last know maxreplica and
minreplicas would exist in the local clusters. Additionally in this alternative there
is no need of collection and processing of the pod metrics for the target object from
each individual cluster.
This document proposes to use ***design alternative 1*** as the preferred implementation.

# Glossary

These use cases are specified using the terminology partly specific to telecom products/platforms:

1 - A telecom service provider has a large number of base stations, for a particular region, 
each with some set of virtualized resources each running some specific network functions. 
In a specific scenario the resources need to be treated logically separate (thus making large 
number of smaller clusters), but still a very similar workload needs to be deployed on each 
cluster (network function stacks, for example).

2 - In one of the architectures, the IOT matrix has IOT gateways, which aggregate a large 
number of IOT sensors in a small area (for example a shopping mall). The IOT gateway is 
envisioned as a virtualized resource, and in some cases multiple such resources need 
aggregation, each forming a small cluster. Each of these clusters might run very similar 
functions, but will independently scale based on the demand of that area.

3 - A telecom service provider has a large number of base stations, each with some set of 
virtualized resources, and each running specific network functions and each specifically 
catering to different network abilities (2g, 3g, 4g, etc). Each of these virtualized base 
stations, make small clusters and can cater to specific network abilities, such that one 
can cater to one or more network abilities. At a given point of time there would be some 
number of end user agents (cell phones) associated with each, and these UEs can come and 
go within the range of each. While the UEs move, a more centralized entity (read federation) 
needs to make a decision as to which exact base station cluster is suitable and with needed 
resources to handle the incoming UEs.
