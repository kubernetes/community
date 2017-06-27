# Federated StatefulSet design
irfan.rehman@huawei.com

# Goal

To outline the use cases for the need of federated Statefulset and determine
suitable design/implementation approach for the same.

# Background

The original design proposal of the statefulset in kubernetes, including the
use cases and example applications, can be found [here](https://github.com/kubernetes/community/blob/master/contributors/design-proposals/stateful-apps.md).

# Use Cases

1 – A stateful app, wants replicas distributed in multiple clusters, such that it can
form multiple smaller sized local app clusters using only local replicas, each serving
users local to the cluster (using in-cluster stateful app identities).

2 – A stateful app, with the goal of high availability, wants the stateful pods
distributed in different clusters, such that the set can withstand the failure of one
or more clusters. This represents an app with global quorum. This use case can be an
extension to the previous, where the same app instances get multiple identities (local
and global).

# Concrete use case examples

## Local identities
1 - A federated etcd statefulset of 30 replicas across 10 clusters which creates a local quorum
in each of the 10 clusters. The potential global quorum is ignored in this example. If a cluster
 goes down, the quorum in that cluster is dead, but the quora in all the other clusters remain
 up accessible locally to the in-cluster apps. (Similar to use case 1 above).

## Global identities
1 - A federated etcd statefulset of 3 or more replicas across 3 or more clusters which
creates a global quorum able to withstand the temporary failure of any one cluster,
as well as the permanent failure of any node in any cluster.

## Multiple identities
1 - A federated statefulset of an app, which is designed to inherently support a geo-distributed
app cluster, with replicas in multiple federated k8s clusters and inherent ability to
communicate across data centres (for example
[consul multi-dc cluster](https://www.consul.io/docs/internals/architecture.html#10-000-foot-view)).

2 - A federated cassandra db statefulset of 100 replicas across, for example, 5 clusters.
(The instances can have multiple identities each; apps and peer instances can choose to
preferentially connect to local dns identities, but can also connect using global dns identity,
if need be, for example to connect to peers in other federated clusters).

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
in this document. The detailed design for the primary alternative is elaborated in the next section.
A design alternative is also evaluated and listed later in the doc.

## Detailed Design (Primary Alternative)

In this approach, the federated statefulset controller will partly behave quite similar to federated
replicaset or the federated deployment controller.
The federated controller would create and monitor individual statefulsets partitioning and distributing
the total stateful replicas across the federated clusters.

As an additional proposal in this design we suggest the possibility of the pods having multiple identities.
This facilitates application pods being able to participate in different quora, and leaves a
choice with application replicas to join either the local or the global quora, as need be.
As of now we can envision the need of two possible identities for each stateful pod (relevant to use case 2),
enabled by this approach, but the same design can be extended for a stateful replica to get more then 2
identities also, in future, if the need be.
The ordinal number assigned to the pod will be calculated and assigned in exactly the same way,
it is assigned in the k8s cluster as described
[here](https://kubernetes.io/docs/concepts/abstractions/controllers/statefulsets/#pod-identity) and will
be unique.
The details of how the replicas will be distributed across federated clusters is elaborated in
&#39;Replica distribution&#39; section below.
The details of the multiple dns identities for each pod is separately elaborated in
&#39;instance identity and discovery&#39; sections below.

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
behaviour will be, to partition the replica count equally into healthy clusters. This default behaviour
in future can be overridden by applying user preferences very similar to replicaset preferences used for
federated replicasets today. It is proposed that these preferences are kept outside the scope of the
first phase (alpha) of the federated statefulsets implementation.

### Instance identity and discovery (local identities)

In the case of in-cluster statefulset, pods discover each other using the in-cluster dns names.
In k8s, a headless service with selectors enables creation of dns records against the pod names
for an in-cluster statefulset.
As an inherent design requirement in this proposal, the applications which are being deployed as
the statefulset can choose to connect to the other replica instances locally (just like the local
statefulset) within the cluster or globally across the clusters (additional federation identity) to
form the quorum using the respective dns names.
Applications (in case they support such a behaviour) can also communicate across quora (for example
a local db cluster connects to local replicas forming a quorum and then subsequently connects to other
cluster/instance outside the cluster for periodic backups) using the global dns identities.
This section deals with the cluster local stateful identities.

When a request for the statefulset object creation is submitted to the federation control plane,
the federated statefulset controller will partition the total number of stateful replicas
and create statefulsets with partitioned replica numbers into at least 1 or more clusters.
The noteworthy point is the proposal that federated stateful controller would additionally modify
the statefulset name by appending the cluster name to the statefulset name into whichever cluster
the partitioned statefulset is getting created.
This will ensure that each pod in the federated statefulset maintains an unique identity across
all clusters, including a stable non-changing hostname for each pod even with the ordinal numbers
being generated local to the clusters. This also seems reasonable due to the lack of portability
of storage and thus the dns identities across the clusters.
This is an important modification and will need to be documented well for the users of the
federated stateful replicas.

There is a possible issue, however quite unlikely, with this approach, that the name length might exceed
the allowed k8s object name of 254 characters. This can be left as an open issue (likely to be hit
only with automated name generators, used for both clusters joining federation and the statefulset name)
for now, and can be subverted using an admission control plugin, to check name lengths in future,
only if needed.

Together with the statefulset object, the user will also need to submit the request for a headless
service corresponding to the statefulset. When this request is sent to the federation control plane,
the federation service controller will create a matching headless service in each of the clusters.
As the partitioned statefulsets get created locally into federated clusters, for the clusters getting
statefulsets (those which get more than 1 replica in replica distribution), pods will be able to
discover each other using the in-cluster dns identity provided by the headless in-cluster service exactly
as it happens when a stateful set is created locally in a cluster
(as described [here](https://kubernetes.io/docs/tutorials/stateful-application/basic-stateful-set/#using-stable-network-identities)).

For an example, let&#39;s say we have 3 federated clusters and the following yamls are submitted
to create a federated statefulset:
```
---
apiVersion: v1
kind: Service
metadata:
 name: etcd
 labels:
   app: etcd
spec:
 ports:
 - port: 80
   name: store
 clusterIP: None
 selector:
   app: etcd
---
apiVersion: apps/v1beta1
kind: StatefulSet
metadata:
 name: store
spec:
 serviceName: "etcd"
 replicas: 11
 template:
   metadata:
     labels:
       app: etcd
   spec:
     terminationGracePeriodSeconds: 10
     containers:
     - name: etcd
       image: gcr.io/google_containers/etcd:3.0.17
       ports:
       - containerPort: 80
         name: store
       volumeMounts:
       - name: etcd_store
         mountPath: /usr/share/etcd
 volumeClaimTemplates:
 - metadata:
     name: etcd_store
   spec:
     accessModes: ["ReadWriteOnce"]
     resources:
       requests:
         storage: 1Gi
```
The service will be created in each of the 3 clusters unmodified.

The default replica partitioning will mean a distribution of **4, 4 and 3** replicas into the 3
clusters.

The statefulset spec submitted to the clusters will have the statefulset metadata.name modified
to _**store+&lt;cluster-name&gt;**_.

The details following extend the example yaml above, with statefulset name _store_, service name
_etcd_, federation name _myfederation_, domain name as _example.com_, and assume the registered
cluster names being _c1, c2 and c3_.

Peer pods can discover the stateful pods locally by querying the governing service domain
```etcd.mynamespace```, which will by default be expanded to a local CNAME of
```etcd.mynamespace.svc.example.com``` (assuming that the default domain name of the individual
clusters is overridden with example.com). The CNAME could as we be queried directly. This will
resolve to the SRV records local to the clusters stateful replicas for each clusters, something
like listed further. The name resolution is exactly as it happens for the local statefulsets in
k8s today. Additionally note that the federation name _myfederation_ has explicitly been dropped
in this example to demonstrate a local query. This is also how the in-cluster kube-dns works. It
parses and considers the service name query to be local or federated global based on the presence
of valid federation name segment in the queried string. [Refer here](https://github.com/kubernetes/dns/blob/master/pkg/dns/dns.go#L588).

The CNAME query will return SRV records as below.

within cluster c1
```
store-c1-0.etcd.mynamespace.svc.example.com
store-c1-1.etcd.mynamespace.svc.example.com
store-c1-2.etcd.mynamespace.svc.example.com
store-c1-3.etcd.mynamespace.svc.example.com
```
within cluster c2
```
store-c2-0.etcd.mynamespace.svc.example.com
store-c2-1.etcd.mynamespace.svc.example.com
store-c2-2.etcd.mynamespace.svc.example.com
store-c2-3.etcd.mynamespace.svc.example.com
```
within cluster c3
```
store-c3-0.etcd.mynamespace.svc.example.com
store-c3-1.etcd.mynamespace.svc.example.com
store-c3-2.etcd.mynamespace.svc.example.com
```
The SRV record will eventually map to an A record with a valid replica pod in-cluster IP if
the replica is up and running.

Additional pointer is that, if the given cluster is properly added to federation, this same
headless service and related resolved names can be queried including the federation name in
query too. For example the service CNAME record ```etcd.mynamespace.myfederation.svc.example.com```
will also result in similar local names as elaborated above, especially because at this point
the headless service, even though federated, will not get any dns names updated into the
global (federated) dns server. The results will however be different, if the global identities
are also requested by the user as elaborated in the next section.

### Instance identity and discovery (global identities)

The above is fine for the case of multiple local quora (one in each cluster), but these
identities cannot work across the clusters, as the local pod IPs cannot be reached across
the clusters.

The proposal to make it work across the clusters is to create an additional service of type
&#39;LoadBalancer&#39; for each stateful pod instance that is created in the clusters. The
name of each service instance will be same as the stateful instance name assigned to the
stateful replicas. This will be automatically created by the federated statefulset controller,
on demand, proposed to be controlled by an annotation on the statefulset. If the annotation
specifies so (users choice), the federation controller will create the additional service of
type &#39;LoadBalancer&#39; against each pod replica. The key name for the annotation will be:
```federation.kubernetes.io/statefulset-lb-needed```
The values could be ```true``` or ```false```. The default value will be ```false```, considering
that the external load balancer is an expensive resource.

The service name for the LB service, one each against the stateful pod will be the same name as
the stateful pod unique name. Also note that these particular LB services, although federated,
will ideally resolve to 1 IP per service. The LB services will be created with proper label
selectors to ensure the network traffic can reach the designated pods. This would also mean
that each stateful pod will need to be updated with an additional unique label selector which
can be matched with by the individual LB services. The suggestion is to use
```federated-statefulset: <stateful replica unique name>``` as the label selector.

The details of creation and handling the LB services are elaborated in section **Handling
additional LB Services** later in this doc.

The trick for individual pods being able to discover this additional dns identity through the LB
service is the federated controller first creating SRV records for each stateful pod in the
federation dns server against the governing service domain
```store.mynamespace.myfederation.svc.example.com``` and then later updating an A record of the IP
resolved for the specific LB service (for example ```store-c1-0.mynamespace.myfederation.svc.example.com```)
against each of these SRV records.
A peer pod in a stateful set then can discover its federated peers by doing a dns query against
the governing service domain ```store.mynamespace.myfederation.svc.example.com```.
It should be noted that as of now the federation service controller ignores service records which
are not of type Loadbalancer, and does not update them in the federation dns server. This behaviour
will need to be modified to ensure that the dns records against the headless services are also updated
in the federated dns server required. The federation service controller will determine this based
on the user annotation (```federation.kubernetes.io/statefulset-lb-needed```) with value ```true``` present
on the federated headless service. Please note that this is the same annotation which is used on the
federated statefulset to determine if this statefulset needs LB services or not. The federation
statefulset controller will synchronize the headless service with the annotation value.
Additionally the information about the stateful replicas (and the identically named LB services) for
which the service controller should update the dns records against the headless service will also
need to be passed on to the headless service. The proposal is to pass this information through
annotations on the headless service as below:
```
federation.kubernetes.io/stateful-replica-list: <replica-name-1>, <replica-name-2>, ..., <replica-name-n>
```
The federation service controller, if it finds the names in this list against a federated headless service,
will look for matching federated LB services and update the CNAME record for the headless service with the
relevant SRV records for all the LB services in the federation dns server.

Once this happens, the stateful pods can discover peers by querying the governing service domain (federated
headless service). Taking the same example listed earlier if the peer (or a user) queries the service
```store.mynamespace.myfederation.svc.example.com```, the below SRV records will be returned.
```
store-c1-0.etcd.mynamespace.myfederation.svc.example.com
store-c1-1.etcd.mynamespace.myfederation.svc.example.com
store-c1-2.etcd.mynamespace.myfederation.svc.example.com
store-c1-3.etcd.mynamespace.myfederation.svc.example.com
store-c2-0.etcd.mynamespace.myfederation.svc.example.com
store-c2-1.etcd.mynamespace.myfederation.svc.example.com
store-c2-2.etcd.mynamespace.myfederation.svc.example.com
store-c2-3.etcd.mynamespace.myfederation.svc.example.com
store-c3-0.etcd.mynamespace.myfederation.svc.example.com
store-c3-1.etcd.mynamespace.myfederation.svc.example.com
store-c3-2.etcd.mynamespace.myfederation.svc.example.com
```
The global federation dns service will have SRV records and relevant A records listing IPs of the LBs.
The local in-cluster kube-dns will have the SRV records and the relevant A records listing the local
stateful set pod IPs. The local kube-dns will be able to respond with both the local and global records,
when a query hits the local kube-dns. The peers can choose to connect to the local identities or the
global ones as the need be.

#### Handling additional LB services

As listed in the above approach, the current option of exposing the statefulset pods across clusters is to
assign a LB service to each replica pod.
If the federated service controller is used (by creating the service in federation) to create these services,
then as per the current behaviour the service will be created in each cluster.
This leads to exactly one unused service per statefulset per cluster.
There are two alternatives to circumvent this problem:

1 - Cluster affinity is introduced for services such that the controller creates the service only
in needed cluster. The need of cluster affinity and anti-affinity in general is discussed
[elsewhere](https://github.com/kubernetes/kubernetes/issues/41442) also, and as of today is already
available in federation.

2 - Federated statefulset controller handles the creation of cluster local services, rather than the
federated service controller. It also will need to handle the dns records for each LB service into the
federation (or cloud provider specific) public dns server.

We propose using alternative 1 listed here, as it fits broader scheme of things and is more consistent with
the user expectation of being able to query all needed resources from the federation control plane and is
less confusing to use at the same time.

Additionally, if a service with the identical name (with same or other properties exists already), it
will be overwritten, similar to the cascading objects behaviour (for example, a deployment creates pods).

## Storage volumes

There are two ways a federated statefulset can be assigned a persistent storage.

1 - A node local volume attached to the container

2 - A persistent volume claim from the cloud provider.

Either of the 2 options above can be used/specified by the user to attach a persistent store to the pods.

It is known, that there is no automated way or  k8s feature/API, yet, which can enable use of the same
persistent volume across different clusters, without users/admins intervention. In the absence of a
direct way of quick migration of storage data from one zone to another, or from one cloud provider
to another if the case be, the proposal is to disallow migration/movement of stateful replicas across
clusters.
This proposal will thus be implemented with a limitation that a particular stateful pod once created into a
particular cluster, will not be moved to another.

## Scale up/Scale down

When a federated statefulset is scaled up or down, the distribution of replicas will be recalculated and
the replica numbers in each cluster will be updated independently.

The behaviour of the scaling on being targeted using a **federated hpa** will be discussed in a separate design.

## Design Alternative

In this alternative the control over the stateful pods will lie only with the federation
statefulset controller. The federated replicset  controller should be able to spin off and
monitor individual pods into federated clusters directly in a similar way to how an in-cluster
controller would spin off and monitor the pods in a stand-alone cluster.
The federation controller should be able to calculate and assign ordinal identities to these
pods. The federation controller should also enforce the creation sequence, the same way
an in-cluster controller would.

### Replica distribution (across federated clusters)

The federated controller will interact with each cluster and create and monitor individual
pod objects in the clusters directly. The pod&#39;s spec and state will be stored in individual
cluster&#39;s etcd, as it would happen on creating a pod in the given cluster.
The controller also assigns each ordinal replica to a specific cluster.
As the primary intention of this feature is to provide a solution for high availability,
it makes great sense to ensure that the instances are distributed into maximum clusters possible.
The initial proposal is to simply hash the replica number into the number of healthy clusters.
A simple modulo calculation can be used to do the distribution.

### Instance identity and discovery

In the case of in-cluster statefulset, pods discover each other using the in-cluster dns names.
A headless service with selectors enables creation of dns records against the pod names.
This cannot work across clusters, as local pod IPs cannot be reached across the clusters.

Similar to the primary alternative, the proposal to make it work across the clusters is to
assign a service type &#39;LoadBalancer&#39; for each pod instance that is created in the clusters.
The service name will be based on the ordinal pod name assigned to each pod.
Additionally a set of label selectors will be added to each individual pod to ensure that the
cluster routes the specific service traffic to the specific pod.
The service type &#39;LoadBalancer&#39; will ensure that the data traffic can be routed across clusters.
The network dns names for each pod will remain very similar to the network names used by the in-cluster statefulsets.

As an example (using the same yamls and names as in the earlier par of this doc):

For a federation with 3 clusters all in the same DNS zone _example.com_ and the statefulset
name _store_ the dns names for instances might look like:
```
store-0.etcd.mynamespace.myfederation.svc.federation.example.com
store-1.etcd.mynamespace.myfederation.svc.federation.example.com
store-2.etcd.mynamespace.myfederation.svc.federation.example.com
store-3.etcd.mynamespace.myfederation.svc.federation.example.com
store-4.etcd.mynamespace.myfederation.svc.federation.example.com
store-5.etcd.mynamespace.myfederation.svc.federation.example.com
store-6.etcd.mynamespace.myfederation.svc.federation.example.com
store-7.etcd.mynamespace.myfederation.svc.federation.example.com
store-8.etcd.mynamespace.myfederation.svc.federation.example.com
store-9.etcd.mynamespace.myfederation.svc.federation.example.com
store-10.etcd.mynamespace.myfederation.svc.federation.example.com
store-11.etcd.mynamespace.myfederation.svc.federation.example.com
```
This will be achieved by federated controller creating and monitoring individual services in
specific clusters against each pod, with the same name as that pod&#39;s
ordinal identity. It will further need to be ensured that the pods be assigned the same hostname
as the pod name (ordinal identity) and service name to keep it consistent with the in-cluster
statefulsets topology.

The details about creation and handling of the LB services and updating the dns records will
be similar to that elaborated in the previous design.

# Conclusion
### Primary design
*** Pros ***
 - The federation controller logic is simpler, as the in-cluster statefulset controller
 functionality is reused (similar to other federated object controllers).
 - Consistent with other federated controllers (including pods).

*** Cons ***
 - The additional LB services are costly resources.

### Alternative design
*** Pros ***

*** Cons ***
 - The controller will be complex to implement and a good amount of statefulset functionality
 will either need to be reimplemented in federation controller or modifications/reorganisation
 to reuse the same portions of code will be needed.
 - As the control over the stateful pods is with the federation statefulset controller, the
 monitoring of the stateful pods will not happen, if for some reason a cluster (or all clusters)
 are disconnected from the federation control plane.
 - Even if this proposal is extended to have a separate local controllers to monitor the
 in-cluster objects (pods or statefulsets) locally, the below drawbacks would follow:
 -- One more controller over the same objects (k8s statefulset objects which can be queried
 from individual cluster), seems two point of control on the same object.
 -- If the new local controller monitors only stateful pods assigned to its cluster, then the
 statefulset object is not locally query-able.
 - Even if the app only needs local quorum, the communication across stateful pods is forced
 through public services, rather then the in-cluster traffic routing.
 - The additional LB services are costly resources.

Its thus proposed to use the primary design alternative listed earlier in this doc.

# Federated statefulset updates

The in-cluster statefulsets update proposal is under implementation as per [this proposal](https://github.com/kubernetes/community/pull/503)
The suggestion is to handle it as a separate design proposal, once the same is implemented and stabilised.


# Limitations (to be discussed)

_1 – In case of a particular replica or the cluster containing particular replicas die, what should be the behavior._

It can be resurrected with the same identity in the original cluster only

_2 – Is there a migration scenario? If there is, in case a replica needs migration, how do we handle this?_

elaborated by @quinton-hoole

Strictly speaking, migration between clusters in the same zone is quite feasible.

I think we&#39;re left with a few options:

(1) only allow the use of node-local persistent (bad),

(2) disallow migration of replicas between clusters (simple, but perhaps overly restrictive)

(3) allow migration of replicas between clusters provided that either the replica uses node-local-volumes, or the
clusters are in the same zone (less restrictive, but more complex to implement, and for the user to understand).
I suspect that the right answer here might be to start with (2) and add 3 in a later phase.

As a proposal in this design, and this phase of implementation replicas once created in a cluster cannot be moved
to another cluster. If a particular statefulset replica dies and the cluster and other stateful replicas in that particular cluster
are still alive, the behavior will be within the purview of local statefulset controller to handle.
If a cluster dies, which had stateful replicas, then the statefulset will need to run with remaining replicas.

_3 – What happens if a cluster dies_

Nothing, statefulset would need to run with lesser replicas# Federated StatefulSet design
