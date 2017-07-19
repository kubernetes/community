# Kubernetes Scalability thresholds

## Background

Since 1.6 release Kubernetes officially supports 5000-node clusters. However,
the question is what that actually means. As of early Q3 2017 we are in the
process of defining set of performance-related SLIs ([Service Level Indicators])
and SLOs ([Service Level Objectives]).

However, no matter what SLIs and SLOs we have, there will always be some users
coming and saying that their cluster is not meeting the SLOs. And in most cases
it appears that the reason behind is that we (as developers) have silently
assumed something (e.g. there will be no more than 10000 services in the
cluster) and users were not aware of that.

This document is trying to explicitly summarize limits for the number of objects
in the system that we are aware of and state if we will try to relax them in the
future or not.

## Kubernetes thresholds

We start with explicit definition of quantities and thresholds we assume are
satisfied in the cluster. This is followed by an explanations for some of those.
Important notes about the numbers:
1. In most cases, exceeding these thresholds doesn’t mean that the cluster
   fails over - it just means that its overall performance degrades.
1. **Some thresholds below (e.g. total number of all objects, or total number of
   pods or namespaces) are given for the largest possible cluster. For smaller
   clusters, the limits are proportionally lower.**
1. The thresholds obviously differ between different Kubernetes releases
   (hopefully each of them is non-decreasing). The numbers we present are for
   the current release (Kubernetes 1.7 release).
1. There are a lot of factors that influence the thresholds, e.g. etcd version
   or storage data format. For each of those we assume the default from the
   release to avoid providing numbers for huge number of combinations of those.
1. The “Head threshold” is representing the status of Kubernetes head. This
   column should be snapshotted at every release to produce per-release
   thresholds (and dedicated column for each release should then be added).

| Quantity                            | Head threshold | 1.8 release | Long term goal |
|-------------------------------------|----------------|-------------|----------------|
| Total number of all objects         | 250000         |             | 1000000        |
| Number of nodes                     | 5000           |             | 5000           |
| Number of pods                      | 150000         |             | 500000         |
| Number of pods per node<sup>1</sup> | 100            |             | 100            |
| Number of pods per core<sup>1</sup> | 10             |             | 10             |
| Number of namespaces (ns)           | 10000          |             | 100000         |
| Number of pods per ns               | 15000          |             | 50000          |
| Number of services                  | 10000          |             | 100000         |
| Number of all services backends     | TBD            |             | 500000         |
| Number of backends per service      | 5000           |             | 5000           |
| Number of deployments per ns        | 20000          |             | 10000          |
| Number of pods per deployment       | TBD            |             | 10000          |
| Number of jobs per ns               | TBD            |             | 1000           |
| Number of daemon sets per ns        | TBD            |             | 100            |
| Number of stateful sets per ns      | TBD            |             | 100            |
| Number of secrets per ns            | TBD            |             | TBD            |
| Number of secrets per pod           | TBD            |             | TBD            |
| Number of config maps per ns        | TBD            |             | TBD            |
| Number of config maps per pod       | TBD            |             | TBD            |
| Number of storageclasses            | TBD            |             | TBD            |
| Number of roles and rolebindings    | TBD            |             | TBD            |

There are also thresholds for other types, but for those the numbers depend
also on the environment (bare metal or which cloud provider) the cluster is
running in. These include:

| Quantity                                  | Head threshold | 1.8 release | Long term goal |
|-------------------------------------------|----------------|-------------|----------------|
| Number of ingresses                       | TBD            |             | TBD            |
| Number of PersistentVolumes               | TBD            |             | TBD            |
| Number of PersistentVolumeClaims per ns   | TBD            |             | TBD            |
| Number of PersistentVolumeClaims per node | TBD            |             | TBD            |


The rationale for some of those numbers:
1. Total number of objects <br/>
There is a limitation on the total number of objects on the system, as this
affects among others etcd and its resource consumption.
1. Number of nodes <br/>
We believe that having clusters with more than 5000 nodes is not the best
option and users should consider splitting into multiple clusters. However,
we may consider bumping the long term goal at some time in the future.
1. Number of services and endpoints <br/>
Each service port and each service backend has a corresponding entry in
iptables. Number of backends of a given service impact the size of the
`Endpoints` objects, which impacts size of data that is being sent all over
the system.
1. Number of objects of a given type per namespace <br/>
This holds for different objects (pods, secrets, deployments, ...). There are
a number of control loops in the system that need to iterate over all objects
in a given namespace as a reaction to some changes in state. Having large
number of objects of a given type in a single namespace can make those loops
expensive and slow down processing given state changes.

---
<sup>1</sup> The limit for number of pods on a given node is in fact minimum from the “pod per node” and “pods per core times number of cores of a node”.

[Service Level Indicators]: https://en.wikipedia.org/wiki/Service_level_indicator
[Service Level Objectives]: https://en.wikipedia.org/wiki/Service_level_objective
