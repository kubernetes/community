# Multicluster reserved namespaces

@perotinus

06/06/2018

## Background

sig-multicluster has identified the need for a canonical set of namespaces that
can be used for supporting multicluster applications and use cases. Initially,
an [issue](https://github.com/kubernetes/cluster-registry/issues/221) was filed
in the cluster-registry repository describing the need for a namespace that
would be used for public, global cluster records. This topic was further
discussed at the
[SIG meeting on June 5, 2018](https://www.youtube.com/watch?v=j6tHK8_mWz8&t=3012)
and in a
[thread](https://groups.google.com/forum/#!topic/kubernetes-sig-multicluster/8u-li_ZJpDI)
on the SIG mailing list.

## Reserved namespaces

We determined that there is currently a strong case for two reserved namespaces
for multicluster use:

-   `kube-multicluster-public`: a global, public namespace for storing cluster
    registry Cluster objects. If there are other custom resources that
    correspond with the global, public Cluster objects, they can also be stored
    here. For example, a custom resource that contains cloud-provider-specific
    metadata about a cluster. Tools built against the cluster registry can
    expect to find the canonical set of Cluster objects in this namespace[1].

-   `kube-multicluster-system`: an administrator-accessible namespace that
    contains components, such as multicluster controllers and their
    dependencies, that are not meant to be seen by most users directly.

The definition of these namespaces is not intended to be exhaustive: in the
future, there may be reason to define more multicluster namespaces, and
potentially conventions for namespaces that are replicated between clusters (for
example, to support a global cluster list that is replicated to all clusters
that are contained in the list).

## Conventions for reserved namespaces

By convention, resources in these namespaces are local to the clusters in which
they exist and will not be replicated to other clusters. In other words, these
namespaces are private to the clusters they are in, and multicluster operations
must not replicate them or their resources into other clusters.

[1] Tools are by no means compelled to look in this namespace for clusters, and
can choose to reference Cluster objects from other namespaces as is suitable to
their design and environment.
