# StatefulSets Session

Topics to talk about:
* local volumes
* requests for the storage sig
* reclaim policies
* Filtering APIs for scheduler
* Data locality
* State of the StateFulSet
* Portable IPs
* Sticky Regions
* Renaming Pods

## State of the StatefulSet

1.5 will come out soon, we'll go beta for StatefulSets in that one.  One of the questions is what are the next steps for Statefulsets?  One thing is a long beta, so that we know we can trust statefulsets and they're safe.

Missed some discussion here about force deletion.

The pod isn't done until the kubelet says it's done.  The issue is what happens when we have a netsplit, because the master doesn't know what's happening with the pods.  In the future we'll maybe add some kind of fencer to make sure that they can't rejoin.  Fencing is probably a topic for the Bare-Metal Sig.

Are we going to sacrifice availability for consistency?  We won't explicitly take actions which aren't safe automatically.  Question: should the kubelet delete automatically if it can't contact the master?  No, because it can't contact the master to say it did it.

When are we going to finish the rename from PetSet to StatefulSet?  The PR is merged for renaming, but the documentation changes aren't.  

Storage provisioning?  The assumption is that you will be able to preallocate a lot of storage for dynamic storage so that you can stamp out PVCs.  If dynamic volumes aren't simple to use this is a lot more annoying.  

Building initial quorums issue?

It would be great to have a developer storage class which ties back to a fake NFS.  For testing and dev.  The idea behind local volumes is that it should be easy to create throwaway storage on local disk.  So that you can write things which run on every kube cluster.

Will there be a API for the application?  To communicate members joining and leaving.  Answer today is that's what the KubeAPI is for.

The hard problem is configchange.  You can't do config change unless you bootstrap it correctly.  If kube is changing things under me I can't maintant quorum (as an app).  This happens when expanding the set of nodes.  You need to figure out who's in and who's out.  

Where does the glue software which relates the statefulset to the application?  But different applications handle things like consensus and quorum very differently.  What about notifying the service that you're available for traffic.  Example for this with etcd with readiness vs. membership service.  You can have two states, one where the node is ready, and one where the application is ready.  Readiness vs. liveness check could differentiate?

Is rapid spin-up a real issue?  Nobody thinks so,
