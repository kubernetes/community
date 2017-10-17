# Requirements

This document considers "customer facing" entities. For our purposes these are
the Kubernetes API, and running workloads.

1. Kubernetes needs to be safe in a multi-tenant environment
2. Securing Kubernetes should be easy

This document uses the "[CIA triad](https://en.wikipedia.org/wiki/Information_security#Key_concepts)" for analyzing these goals.

# Kubernetes API access controls

Kubernetes needs access controls capable of protecting the CIA of cluster
resources in a multi-tenant environment.

## Concerns with method based access controls

Many Kubernetes deployments provide method based access control (e.g.,
list pods). This is akin to securing \*nix with permissions on the "open"
syscall. To allow more than privileged user, ACLs should be placed on nouns, not
verbs.

## Concerns with implicit relationships

Kubernetes' implicit object relationships create subtle dangers.

1. The "create pod" privilege allows a user to mount a secret, even if the user
   cannot "get secrets". This is a failure of confidentiality.
2. The "create pod" privilege can be used to place a pod behind a service, even
   if the user cannot "update service". This is a failure of integrity.
3. The "create pod" privilege can be used to place a pod behind a deployment,
   which can be used to starve legitimate replicas. This is a failure of
   integrity, and availability.

This isn't a complete list of such dangers, but it's sufficient to show that
access controls on an object are insufficient to guarantee CIA. Access controls
are also needed on the relationships between objects.

## Label ACLs (implicit ownership)

Label ACLs are sometimes suggested as a method to provide access control on
object-object relationships. Special permissions would be required to create or
query an ACL'ed label. This approach is both complex and insufficient.

1. The accidental omission of a label means that the system fails open.
2. It becomes increasingly difficult to reason about permissions, especially if
   set-based label selectors are used.
3. Can ACLs be created on existing labels?
   1. Yes: clients can be locked out of their workloads.
   2. No: workloads must be restarted to change permissions.

This approach also perturbs the existing behavior of label selectors. Does the
"list pod" permission allow me to see pods with contentious ACLs? What about
"list secrets"?

Finally, this approach is insufficent, because not all relationships are defined
by label selectors. Notably, secrets cannot be protected by label ACLs.

## Explicit "ownership" (recommended)

Explicit owners can be used to limit access to both objects and relationships. A
service would reject both unauthorized changes to its body, and unauthorized
pods as backends.

The simplicity and correctness of this method recommend it.

### Namespace as "owner"

By scoping method permissions to namespaces, a crude ownership model can be
constructed (all objects are owned by a namespace, and object-object
relationships must have the same owner). The advantage of this approach is that
it's possible today.

There are also disadvantages to this method.

1. Not all resources are namespaced (e.g., nodes and volumes)
2. The myriad escalations and CIA failures make it harder to [separate
   privileges](https://en.wikipedia.org/wiki/Privilege_separation)
   within a namespace.
3. The mapping between namespace and true ownership must be enforced by the
   cluster admin.

### Intentional ownership

Rather than overloading namespace, Kubernetes can introduce an explicit concept
of ownership. For example, a service owned by group\* "foo" could only be updated by members
of "foo". The service's backends would also be limited to pods owned by a group.

1. All resources would be owned.
2. Scoping down permissions is simple with control over both object and relationship access\*\*.
3. These owners could derive from an identity management system, so the onus of
enforcement isn't on the cluster admin.

A true ownership model solves the disadvantages of namespace-based ownership.
But it's not clear that the improvements are worth the effort of implementation.

\* Group is only used as an example. The system could be built to support
users, roles, service accounts, or any other identifiable entity.

\*\* It's also reasonable to control relationships with a reflective or membership
constraint. This simplifies permissions management at the expense of
flexibility, and gives no real improvement over problem (2) with namespaces.

# Workloads

The other face of Kubernetes exposed to customers is the workload environment.
This section focuses on protecting workloads from the environment, and the
environment from workloads.

## Scheduled isolation

Many users have use cases for scheduled isolation. This is where Kubernetes
separates sensitive workloads from each other.

1. A sensitive application might not tolerate other workloads on the same node.
2. A user may bring a dedicated node.
3. An application may need to run privileged containers.

Both affinity and taints can achieve this today, but neither is protected. Both
affinities and tolerations are specific on the "PodSpec", so anyone with "create
pod" permissions can circumvent the mechanism.

### Dedicated nodes

Dedicated nodes require scheduling protection, because A's safety shouldn't be
dependent on B's goodwill. Taint ACLs would suffer from the same problem as
label ACLs. However, the "ownership" model can be applied successfully.

For instance, a special "allowed-namespaces" taint would forbid pods outside of
authorized namespaces. The pods would still be responsible for selecting the
nodes.

Alternatively, membership in a group, or a role, could be used to permit or
forbid execution on a node. This might make more sense than forcing namespaces
onto nodes, but probably requires the intentional ownership model.

### Privileged pods

A PodSecurityPolicy can prevent users from creating privileged containers.
However, some use cases may be that privileged pods can only be created on
dedicated nodes.

For instance, a game server may want to run with host networking. These will
need to be created with a privileged PSP. But many workloads may not tolerate
sharing a host with privileged user pods.

1. Privileged containers could taint a node.
2. A cluster-level policy could restrict privileged containers to dedicated
   nodes.
3. A "NodeSecurityPolicy" could be introduced, which could forbid privileged
   user containers.

## Identity and Secrets

Kubernetes is not intended to be a key management system. Customers should run a
key management system (KMS) that emphasizes security (e.g., pinned memory, encrypted
at rest, etc.). The chosen KMS can deliver appropriate secrets to containers.

Bootstrapping identity may still fall to Kubernetes. Either containers are
created with identifying resources (e.g., a cert), or the bootstrap protocol can
physically verify the origin of a request (e.g., by inspecting OS resources).

### Separation of privileges

Imagine a pod with the following members and permissions:

Permissions: logging (W), metrics (W), database (R/W)
Containers: logger, metric agent, web service

The pod's privileges can be further separated, giving each container only the
permissions needed to function. This suggests that a KMS needs to be capable of
distinguishing individual containers.

### Bootstrapping

An IP-based metadata server (e.g., kube2iam) is insufficient to separate
privileges, because all containers within a pod share a network namespace.

## Host isolation

### Security Modules

Many users are interested in using LSM plugins to further increase security in a
deployment, or even LSM replacements (e.g., RSBAC).

## Isolated networks

If namespace is repurposed as a fine-grained security boundary, then a
namespaced NetworkPolicy seems woefully inadequate. A large organization would
need at least one per project, and they would need to be kept in sync.
