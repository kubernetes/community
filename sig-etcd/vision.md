# SIG etcd Vision

The long-term success of the etcd project depends on the following:
- Etcd is a reliable key-value storage
- Etcd is simple to operate
- Etcd is a standalone solution for managing infrastructure
- Etcd scales beyond Kubernetes dimensions

The goals and milestones listed here are for future releases. 
The scope of release v3.6 has already been defined and is unlikely to change.

## Etcd is a reliable key-value storage service

Reliability remains the most important property of etcd.
The project cannot allow for another [data inconsistency incident].
If we could only pick one thing from the list of goals above, this would be it.
No matter what features we add in the future, 
they must not diminish etcd's reliability. 
We must establish processes and safeguards to prevent future incidents.

How?
- Etcd API guarantees are well understood, documented and tested.
- Etcd adopts a production readiness review process for new features, similar to Kubernetes one.
- Robustness tests should cover most of the API and most common failures.
- New features must have accompanying e2e tests and be covered by robustness tests.
- Etcd must be able to immediately detect corruption.
- Etcd must be able to automatically recover from data corruption.
 
[data inconsistency incident]: https://github.com/etcd-io/etcd/blob/main/Documentation/postmortems/v3.5-data-inconsistency.md

## Etcd is simple to operate

Etcd should be easy to operate.
Currently, there are many steps involved in operating etcd,
and some of these steps require external tools. 
For example, Kubernetes provides tools to [downgrade/upgrade etcd].
These tools are not part of the etcd,
but they are available as part of the Kubernetes distribution of etcd.

How?
- Etcd should not require users to run periodic defrag
- Etcd officially supports live upgrades and downgrades
- Disaster recovery for Etcd & Kubernetes
- Reliable cluster membership changes via learners with automated promotion
- Two node etcd clusters

## Etcd is a standalone solution for managing infrastructure configuration

Kubernetes is not the only way to manage infrastructure.
It was the first to introduce many concepts that have now become the standard,
but they are not unique to Kubernetes.
The most important design principle of Kubernetes,
the reconciliation protocol, is not something unique to it.

Reconciliation can be implemented solely on etcd,
as has been shown by projects like Cillium,
Calico Typha that support etcd-based control planes.
The reason why this idea has not propagated further is
the amount of work that was put into making 
the reconciliation protocol scale in Kubernetes.
The watch cache is a key part of this scaling,
and it is not part of the etcd project.

If etcd provided a Kubernetes-like storage interface
and primitives for the reconciliation protocol,
it would be a more viable solution for managing infrastructure.
This would allow users to build etcd-based control planes that
could scale to meet the needs of large and complex deployments.

How?
- Introduce Kubernetes like storage interface into etcd-client
- Provide etcd primitives for reconciliation protocol
- Strip out the Kubernetes watch cache and make it part of the etcd client.
- Use the watch cache in the client to build an eventually consistent etcd proxy.
 
[downgrade/upgrade etcd]: https://github.com/kubernetes/kubernetes/tree/master/cluster/images/etcd

## Etcd scales beyond Kubernetes dimensions

Etcd has proven its scalability by enabling Kubernetes clusters of up to 5,000 nodes.
However, as the cloud native ecosystem has evolved, new projects have been built on top of Kubernetes.
These projects, such as [KCP] (a multi-cluster control plane) and [Kueue] (a batch job queuing system),
have different scalability requirements than pure Kubernetes.
For example, they need support for larger storage sizes and higher throughput.

Etcd's strong points are its reliable raft and efficient watch implementation.
However, its storage capabilities are not as strong.
To address this, we should look into growing out storage capabilities and making them more flexible depending on the use case.

How?
- Well-defined and tested scalability dimensions
- Increase raft throughput (async and batch proposal handling)
- Increasing bbolt supported storage size
- Pluggable storage layer
- Hybrid clusters with write and read optimized members


[KCP]: https://cloud.redhat.com/blog/an-introduction-to-kcp
[Kueue]: https://github.com/kubernetes-sigs/kueue

