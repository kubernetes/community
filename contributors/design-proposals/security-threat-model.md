# Security Threat Model

As a cluster manager Kubernetes is expected to defend against a wide variety of threats. This document
attempts to define the high level actors, potential threats, and our primary layers of defense.


## Brief architectural description.

Kubernetes is a cluster manager that runs application workloads inside of containers.  Workloads run as
**pods** on **nodes** which communicate with the **api-server** to retrieve the description and list of
pods. Pods may communicate with each other over the network, or contact virtual proxies known as
**services** from each node at a unique IP or DNS address. Traffic to pods from outside the cluster
may flow via an **ingress** router/proxy. Pods may have additional resources attached at runtime,
including persistent volumes, credentials that let them act against the API, secrets used for contacting
other systems, or configuration files for their processes.

Users and administrators access the apiserver to deploy applications, control access, or make
changes to existing applications. **Controllers** are automated agents of the system that use the
definitions in the api-server to ensure the desired state matches reality. Controllers act on-behalf-of
users to make these changes and have access privileges regular users will not.


## Actors

The primary actors and their roles on the system are:

1. Users deploy applications onto the platform via the APIs related to application definition.
   1. Some users are highly limited in the actions they may take, such as only creating a subset of available resources and being unable to edit.
2. Administrators manage security, allocate resources, and conduct maintenance on cluster components.
   1. Organization admins control a subset of resources on the cluster, but are not privileged to act on the whole cluster.
   2. Infrastucture admins manage both the cluster and the underlying systems that run the cluster.
3. Each node offers resources to the cluster and accepts pods that have been scheduled on it.
4. Scheduler controllers assign pods to nodes based on workload definitions and must know details of both to effectively schedule.
5. Workload controllers create pods that have certain lifecycle rules.
6. Network providers use information about the topology of the cluster (physical and logical) to give pods IP address and control their access.
7. Secret controllers create or update secrets for use by users and applications.
8. Federations need to run workloads on behalf of users on individual clusters.
9. Other integrations may need access to read or write definitions on some or all of the namespaces in the cluster.


## Security boundaries

Several important security boundaries exist in Kubernetes that reduce the possibility of interference or exploit by consumers.

1. Between the container processes within a pod and the node, and between different pods on a node
2. At the network interfaces between different pods, from the outside of the cluster to inside (ingress), and from the inside of the cluster to outside (egress)
3. Between infrastructure components and the API
4. Between end users and the API
5. Between the API servers and etcd
6. Between infrastructure monitoring infrastructure and the infrastructure components
7. Between different tenants, primarily modelled as different namespaces or nodes
8. Between the federation control plane and the cluster control plane


## Multi-tenancy

Kubernetes should be capable of subdividing the cluster for use by distinct participants with their own security boundaries.

1. Single-tenant: A small number of administrators who are also users deploying applications, and are often infrastructure admins
2. Collaborative Multi-tenant: Two or more teams acting together on the same cluster without strong boundaries between their roles
3. Multi-tenant: Many users with limited permissions whose applications may need to interact or be isolated, and administrators to configure and enforce that tenancy
4. Organizational multi-tenant: Several large groups of users and applications (organizations) where the organization admin is granted a subset of cluster administrative permissions in order to subdivide responsibility.

TODO: link to multi-tenancy doc


## Threats

The primary threats to Kubernetes are described here:

1. Escaping container and pod isolation to the node or to other containers on the node
2. Accessing APIs as a user without proper authorization or authentication
3. Consuming disproportionate resources on the system to deny other users access
4. Crashing or wedging the system so that no workloads can be processed
5. Encouraging controllers to act as confused deputies and perform actions at a higher level of trust
6. Accessing secret data entrusted to the system without appropriate permission for escalating access to the cluster or other systems
7. Using access to the cluster API to gain elevated permission on the nodes
8. Using access to the federation API to gain elevated permission on the nodes
9. Disguising or concealing malicious actions to prevent forensic assessment after an incident
10. Reusing cluster resources (like services, namespaces, or secret names) after they have been deleted to masquerade as the previous user

TODO: list threats that are not considered


### Defending against threats

TODO: categorize how we approach this

TODO: list important upcoming work