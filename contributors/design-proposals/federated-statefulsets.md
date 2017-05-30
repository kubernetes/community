# Federated StatefulSet design
irfan.rehman@huawei.com, dhilip.kumar.s@huawei.com

# Goal

To outline the use cases for the need of federated Statefulset and determine 
suitable design/implementation approach for the same.

# Background

The original design proposal of the statefulset in kubernetes, including the 
use cases and example applications, can be found [here](https://github.com/kubernetes/community/blob/master/contributors/design-proposals/stateful-apps.md).

# Use Cases

1 – A stateful app, wants replicas distributed in multiple clusters, such that it can 
form multiple smaller sized local app clusters using only local replicas, each serving 
users local to the cluster (using in cluster stateful app identities).

2 – A stateful app, for the reasons of high availability, wants the stateful pods 
distributed in different clusters, such that the set can withstand cluster failures. 
This represents an app with one single global quorum. This use case can be an extension
to the previous, where the same app instances get multiple identities (local and global).

# Concrete use case examples

1 - A federated etcd statefulset of 30 replicas across 10 clusters (using minreplicas=3, 
maxreplicas=3, PVC's) which creates a local quorum in each of the 10 clusters. The potential 
global quorum is ignored in this example. If a cluster goes down, the quorum in that cluster 
is dead, but the quora in all the other clusters remain up accessible locally to the in-cluster
apps. (Similar to use case 1 above).

2 - A federated etcd statefulset of 3 or more replicas across 2 or more then 2 clusters
which creates a global quorum able to withstand the temporary failure of any one cluster,
as well as the permanent failure of any node in any cluster. (Similar to use case 2 above).

3 - A federated cassandra db statefulset of 100 replicas across, for example, 5 clusters.
(The instances can have multiple identities each; apps and peer instances can choose to 
preferentially connect to local dns identities, but can also connect using global dns identity,
if need be, for example to connect to peers in other federated clusters).

4 - A federated statefulset of an app, which is designed to inherently support a geo-distributed 
app cluster, with replicas in mutiple federated k8s clusters. (for example galera db cluster).

# Design Requirements

If we consider the use cases listed above, the main design requirements can roughly be listed as:

1 – An unique, consistent and discoverable identity for each replica/instance across the 
federated clusters.

2 – It must be possible for the federated statefulset to first safely form an initial quorum 
(probably using cluster local replicas and identities).
It should further also be possible to add members (probably from other clusters and global 
identities) to this initial quorum.

3 – The ability to scale, across clusters in some deterministic fashion.

4 – An extended design requirement (which might not be a necessary requirement for this 
phase for work) can be an ability to migrate a pod from one node (in one cluster) to 
another node (possibly in another cluster), with stable network identity along with its attached volumes.

# Design Proposals

The Statefulset object and APIs are already defined in k8s.
The same needs to be exposed from federation to ensure that the interfaces remain consistent 
with a stand-alone k8s cluster.
A separate controller at federation would however be useful.

We considered two alternative designs for this proposal, both of which are further described 
in this document. We find the Design alternative 2, better in many aspects, and find it more 
suitable to be pursued for implementation further.

## Design Alternative 1

The federated replicset  controller should be able to spin off and monitor individual pods 
into federated clusters directly in a similar way to how an in-cluster controller would spin
off and monitor the pods in a stand-alone cluster.
The federation controller should be able to calculate and assign ordinal identities to these 
pods. The federation controller should also enforce the creation sequence, the same way 
an in-cluster controller would.

### Replica distribution (across federated clusters)

The federated controller will interact with each cluster and create and monitor individual 
pod objects in the clusters. The pod&#39;s spec and state will be stored in individual cluster&#39;s 
etcd, as it would happen on creating a pod in the given cluster.
The controller also assigns each ordinal replica to a specific cluster.
As the primary intention of this feature is to provide a solution for high availability, 
it makes great sense to ensure that the instances are distributed into maximum clusters possible.
The initial proposal is to simply hash the replica number into the number of healthy clusters.
A simple modulo calculation can be used to do the distribution.

### Instance identity and discovery

In the case of in-cluster statefulset, pods discover each other using the in-cluster dns names.
A headless service with selectors enables creation of dns records against the pod names.
This cannot work across clusters, as local pod IPs cannot be reached across the clusters.

The proposal to make it work across the clusters is to assign a service type &#39;LoadBalancer&#39; 
for each pod instance that is created in the clusters.
The service name will be based on the ordinal pod name assigned to each pod.
Additionally a set of label selectors will be added to each individual pod to ensure that the
cluster routes the specific service traffic to the specific pod.
The service type &#39;LoadBalancer&#39; will ensure that the data traffic can be routed across clusters.
The network dns names for each pod will remain very similar to the network names used by the in-cluster statefulsets.

As an example:

For a federation with 3 clusters all in the same DNS zone _federation.example_ and the statefulset 
name _web_ the dns names for instances might look like:
```
web-0.mynamespace.myfederation.svc.federation.example.com
web-1.mynamespace.myfederation.svc.federation.example.com
Web-2.mynamespace.myfederation.svc.federation.example.com
```

The difference compared to an in-cluster dns name for an instance of a statefulset can
be noted where the dns id starts with **$(podname).$(governing service domain)**, whereas when 
federated, in this approach, it will start only with **$(service name).**
This will be achieved by federated controller creating and monitoring individual services in 
specific clusters (not a federated service) against each pod, with the same name as that pod&#39;s 
ordinal identity. It will further need to be ensured that the pods be assigned the same hostname 
as the pod name (ordinal identity) and service name to keep it consistent with the in-cluster 
statefulsets topology.

There is one interesting catch about the service creation, here.
The service needs to exist locally in the cluster (1 service per replica pod across clusters), 
which is unlike a federated service. 
If we use the existing federated service mechanism, to create this service, and the service with 
a particular name (for example web-0 in example above) is created in federation, it will 
propagate the same to all federated clusters. 
This is unnecessary for the solution mentioned in this section (or for further reading in 
the following section), and would be waste of resources. 
Some simple proposals to tackle this problem are mentioned at the end of this document under 
section &quot; **Handling additional LB Services**&quot;.

## Design Alternative 2

In this approach, the federated statefulset controller will behave quite similar to federated 
replicaset or the federated deployment controller.
The federated controller would create and monitor individual statefulsets (rather then pods 
directly) partitioning and distributing the total stateful replicas across the federated clusters.

As a proposal in this design we suggest the possibility of the pods having multiple identities.
This facilitates application pods being able to participate in different quora, and leaves a 
choice with application replicas to join either the local or the global quora, as need be.
As of now we can envision the need of two possible identities for each pod (relevant to use case 2), 
enabled by this approach.
The ordinal number assigned to the pod will be calculated and assigned in exactly the same way, 
it is assigned in the k8s cluster as described 
[here](https://kubernetes.io/docs/concepts/abstractions/controllers/statefulsets/#pod-identity) and will be unique.
It however will be possible for each pod to take multiply dns identities.
This is detailed in the &#39;instance identity and discovery&#39; section below.

For this phase of implementation, it is proposed that the statefulset need not guarantee the 
order of creation of the pod instances across federation.

### Replica distribution (across federated clusters)

The proposed default behaviour of the federation controller, when a statefulset creation request is 
sent to federation API server, would be to partition the statefulset replicas and create a statefulset 
into each of the clusters with the reduced replica number (after partitioning), quite similar to the 
behaviour of the replicaset or daemonset controllers of the k8s federation.
If the replica count assigned to the federated statefulset is equal to or smaller than the total 
healthy clusters in the federation, then there will be at most one replica assigned to the statefulset 
created in each underlying cluster (and some clusters might not get a statefulset object at all 
in case of replica number is smaller than total healthy clusters).
For the case of replica numbers being larger than the total number of clusters, the proposed default 
behaviour will be, to partition the replica count equally into healthy clusters.
Annotations could be provided to override this behaviour in exactly the same fashion annotations 
are provided in the spec of federated replicasets ([_frs preferences_](https://github.com/kubernetes/community/blob/master/contributors/design-proposals/federated-replicasets.md#federatereplicaset-preferences)).
The applications which are being deployed as the statefulset can choose to connect to the other 
replica instances locally within the cluster or globally across the clusters to form the quorum 
using the respective dns names.
Applications (in case they support such a behaviour) can also communicate across quora (for example 
a local db cluster connects to another cluster/instance outside the cluster for periodic backups) 
using the global dns identities.

### Instance identity and discovery

Together with the statefulset object, the user will also need to submit the request for a headless 
service corresponding to the statefulset.
When such requests (one for statefulset and another for the headless service) are sent to the 
federation control plane, the federation service controller will create a matching headless service 
in each of the clusters.
It will further partition the total number of replicas and create statefulsets with partitioned 
replica numbers into at least 1 or more clusters.
The noteworthy point is the proposal that federated stateful controller would additionally modify 
the statefulset name by appending the cluster name to the statefulset name into whichever cluster 
the partitioned statefulset is getting created. 
This will ensure that each pod in the federated statefulset maintains an unique identity across 
all clusters, including a stable non-changing hostname for each pod even with the ordinal numbers 
being generated local to the clusters.
There is a possible issue, however unlikely, with this approach is that the name length might exceed 
the allowed k8s object name of 254 charaters. This can be left as an open issue (likely to be hit 
only with automated name generators, for both clusters joining federation and the statefulset name) 
for now, and can be subverted using an admission control plugin, to check name lengths in future.

In case of clusters getting statefulsets with more than 1 replica, pods will be able to discover 
each other using the in-cluster dns identity provided by the headless in-cluster service 
(as described [here](https://kubernetes.io/docs/tutorials/stateful-application/basic-stateful-set/#using-stable-network-identities)).

For an example, let&#39;s say we have 3 federated clusters and the following yamls are submitted 
to create a federated statefulset:
```
---
apiVersion: v1
kind: Service
metadata:
 name: nginx
 labels:
   app: nginx
spec:
 ports:
 - port: 80
   name: web
 clusterIP: None
 selector:
   app: nginx
---
apiVersion: apps/v1beta1
kind: StatefulSet
metadata:
 name: web
spec:
 serviceName: "nginx"
 replicas: 5
 template:
   metadata:
     labels:
       app: nginx
   spec:
     terminationGracePeriodSeconds: 10
     containers:
     - name: nginx
       image: gcr.io/google\_containers/nginx-slim:0.8
       ports:
       - containerPort: 80
         name: web
       volumeMounts:
       - name: www
         mountPath: /usr/share/nginx/html
 volumeClaimTemplates:
 - metadata:
     name: www
   spec:
     accessModes: ["ReadWriteOnce"]
     resources:
       requests:
         storage: 1Gi
```
The service will be created in each of the 3 clusters unmodified.

The default replica partitioning will mean a distribution of **2, 2 and 1** replicas to the 3 clusters.

The statefulset spec submitted to the clusters will have the statefulset metadata.name modified 
to _**web+&lt;cluster-name&gt;**_.

Using the example above, with statefulset name _web_, service name _nginx_ and domain name as
_federation.example_ the dns names created against each instance would be as below, discoverable 
both locally and globally (assuming the registered cluster names being c1, c2 and c3):
```
web-c1-0.nginx.mynamespace.svc.federation.example.com
web-c1-1.nginx.mynamespace.svc.federation.example.com
web-c2-0.nginx.mynamespace.svc.federation.example.com
web-c2-1.nginx.mynamespace.svc.federation.example.com
web-c3-0.nginx.mynamespace.svc.federation.example.com
```
This above is fine for the case of multiple local quora (one in each cluster), but these identities 
cannot work across the clusters, as the local pod IPs cannot be reached across the clusters.

Similar to _design alternative 1_, the proposal to make it work across the clusters is to create an 
additional service of type &#39;LoadBalancer&#39; for each pod instance that is created in the clusters. 
This will be automatically created by the federated statefulset controller, on demand, proposed to be 
controlled by an annotation on the statefulset.
If the annotation specifies so (users choice), the federation controller will create the additional 
service of type &#39;LoadBalancer&#39; against each pod replica.

The additional globally reachable dns identities (LB service) against each stateful pod once created will 
have dns names as below:
```
web-c1-0.mynamespace.svc.federation.example.com
web-c1-1.mynamespace.svc.federation.example.com
web-c2-0.mynamespace.svc.federation.example.com
web-c2-1.mynamespace.svc.federation.example.com
web-c3-0.mynamespace.svc.federation.example.com
```
The service name for the LB service, one each against the pod will be the same name as the pod name.
Also note that these particular LB services, although federated, will ideally resolve to 1 IP per service.
These services will be created with the label selector matching the statefulset pods.

The proposed solution is elaborated in section **Handling additional LB Services**.

## Handling additional LB services

As listed in the above approach, the current option of exposing the statefulset pods across clusters is to 
assign a LB service to each replica pod.
If the federated service controller is used (by creating the service in federation) to create these services, 
then as per the current behaviour the service will be created in each cluster. 
This leads to exactly one unused service per statefulset per cluster.
There are two alternatives to circumvent this problem:

1 - Cluster affinity is introduced for services such that the controller creates the service only 
in needed cluster. (The need of cluster affinity and anti-affinity in general is discussed 
[elsewhere](https://github.com/kubernetes/kubernetes/issues/41442) also)

2 - Federated statefulset controller handles the creation of cluster local services, rather than the 
federated service controller. It also will need to handle the dns records for each LB service into the 
federation (or cloud provider specific) public dns server.

We propose using alternative 1 listed here, as it fits broader scheme of things and is more consistent with 
the user expectation of being able to query all needed resources from the federation control plane and is 
less confusing to use at the same time.

The trick for individual pods be able to discover this additional dns identity through the LB service is for the 
federated controller to be able to update an A record of the IP resolved for the specific LB service 
(for example ```web-c1-0.mynamespace.svc.federation.example.com```)
against the statefulset governing service domain (```nginx.mynamespace.svc.federation.example.com```).
A peer pod in a stateful set then can discover its federated peers by doing a dns query against the governing
service domain ```nginx.mynamespace.svc.federation.example.com```.
It should be noted that the the additional dns record will be updated by the federated service controller, however 
the information linking the LB service and the target statefulset pod and the governing service domain will need to 
be passed from the federated statesfulset controller to federated service controller.
Doing it this was has one drawback that even if an additional LB service is available per stateful pod per cluster,
the dns A record information linking that particular LB service to the stateful pod endpoint will be available only 
when the dns query hits the federated dns server. This can be ignored, because the same statefulset pod can be 
discoverd locally also, using the giverning service domain.

## Storage volumes

There are two ways a federated statefulset can be assigned a persistent storage.

1 - A node local volume attached to the container

2 - A persistent volume claim from the cloud provider.

Either of the 2 options above can be used/specified by the user to attach a persistent store to the pods.

It is known, that as of now, the same persistent volume, even if can be used in a different cluster, 
k8s directly does not provide an API, yet, which can aid the same. In the absence of no direct way of 
quick migration of the storage data from one zone to another, or from one cloud provider to another if the 
case be, the proposal is to disallow migration of pods across clusters.

## Scale up/Scale down

When a federated statefulset is scaled up or down, the distribution of replicas will be recalculated and 
the replica numbers in each cluster will be updated independently.

The behaviour of the scaling on being targeted using a **federated hpa** will be discussed in a separate design.

# Conclusion

Option 1 above can solve only 1 of the listed use cases (scenario 1), and does not remain consistent with 
the existing federated controllers.

Option 2 above can solve both the use cases, is consistent with the existing federated controllers 
and makes more sense in the scheme of things.

# Federated statefulset updates

The in cluster statefulsets update proposal is under implementation as per [this proposal](https://github.com/kubernetes/community/pull/503)
The suggestion is to handle it as a separate design proposal, once the same is implemented and stabilised.


# Limitations (to be discussed)

_1 – In in case of a particular replica or the cluster containing particular replicas die, what should be the behavior._

It can be resurrected with the same identity in the original cluster only

_2 – Is there a migration scenario? If there is, in case a replica needs migration, how do we handle this?_

elaborated by @quinton-hoole

Strictly speaking, migration between clusters in the same zone is quite feasible.

I think we&#39;re left with a few options:

(1) only allow the use of node-local persistent (bad),

(2) disallow migration of replicas between clusters if they use non-node-local volumes (simple, but perhaps overly restrictive)

(3) allow migration of replicas between clusters provided that either the replica uses node-local-volumes, or the 
clusters are in the same zone (less restrictive, but more complex to implement, and for the user to understand). 
I suspect that the right answer here might be to start with (2) and add 3 in a later phase.

_3 – What happens if a cluster dies_

Nothing, statefulset would need to run with lesser replicas
