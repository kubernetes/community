Kubernetes should have a mechanism to restrict pod exection on nodes. The
difficulty is two-fold. First, nodes aren't namespaced, so it's difficult to
give selective access to only a subset of nodes. Second, placement of a pod on a
node is equivalent to write access on the pod/node resource.

#Use cases for sole-tenancy

There are many use cases for dedicated nodes, so a non-comprehensive list is
provided.

1. Many applications require elevated privileges, which could compromise other
applications on the same node.
1. Many applications cannot trust a shared platform.
1. Many applications are designed to be 1:1 with a node (e.g., storage servers).
1. Many applications may want access to special hardware.
1. Many applications are critical, and cannot tolerate interference from
   neighbors.
1. Many applications may need isolation for compliance.

#Requirements for restricted nodes

1. Nodes must be capable of rejecting non-conforming workloads.
1. Nodes must allow conforming workloads.

#Requirements for dedicated workloads

1. Placement must be guaranteeable.
1. Placement must be restrictable.

If placement cannot be guaranteed, then the application may execute in an
unexpected environment, with varying results. For instance, if must run with
Skylake nodes, but is placed on Haswell nodes, then performance may suffer.

If placement cannot be restricted to dedicated nodes, then privileged workloads
may compromise the rest of the cluster. For instance, if the container must load
a kernel driver (e.g., NFS), then administrators may want to quarantine it on
dedicated nodes.

#Proposal: Taints

Node access restriction can be *cooperatively* implemented today with taints and
tolerations. Administrators can taint a node, and restrict tolerations to the
pods intended to run on those nodes.

Unfortunately, tolerations are settable on the pod, so anyone could circumvent
the dedicated nodes' protections either maliciously, or accidentally.

Importantly, this does not keep allowable pods off of untainted nodes. This is
important, because isolation is a prominent use case for dedicated nodes.

## PodSecurityPolicy

Kubernetes already contains a mechanism to restrict admission of privileged
pods: the PodSecurityPolciy admission controller. By considering tolerations a
new type of privilege, a PSP can be used to enforce the integrity of a dedicated
node.

This becomes awkward when the application is intended to run as privileged,
because taints don't force the pod onto a dedicated node. A naive application of
PSP to dedicated nodes could compromise all shared nodes in a cluster.

To protect nodes from errant pods, a "pod taint" is required. These "pod
taints" would correspond to "node tolerations". Simple examination reveals that
a single, symmetric taint/toleration pairing is sufficient to enforce dedication
and placement.

Nodes would reject pods with a "required" toleration, unless that node holds a
corresponding taint. The PodSecurityPolicy then has a list of both allowed
tolerations, and another of required tolerations. This way an administrator can
grant the ability to execute privileged pods, but only on dedicated nodes.

## Namespaced/ServiceAccount toleration admission controller

Given that namespaces are the current security boundary, it might be convenient
to apply a default, namespaced toleration. This could be performed by an
admission controller, so that all pods within a namespace are granted the
correct tolerations.

Because creating a namespace creates a new taint/toleration pairing, the naming
scheme must be protected to prevent collisions. Otherwise creating a namespace
with a well-crafted name could be an escalation path to dedicated nodes.

This feature is entirely a convenience, and possibly not worth the risks
introduced by the complexity. If added, the same design should be considered for
ServiceAccounts.
