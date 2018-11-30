# Kubernetes Scaling and Performance Goals

_by Quinton Hoole and Wojciech Tyczynski, Google Inc_

**April 2016**

This document is a markdown version converted from a working [Google Doc](https://docs.google.com/document/d/1pABjn69yeDJDl0Yrsneiwj5A_bUI7BZgvKTQtSwzjiE/edit).  Please refer to the original for extended commentary and discussion.

## Introduction

What size clusters do we think that we should support with Kubernetes in the short to medium term?  How performant do we think that the control system should be at scale?  What resource overhead should the Kubernetes control system reasonably consume?  This document attempts to answer these and related questions.

## Background
Based on our experience running very large production systems at Google, we have found that running excessively large compute and/or storage clusters has several downsides, most notably:

1. **Large-scale cluster failures can lead to large losses of capacity**: Despite our best efforts, correlated cluster-wide hardware or software failures do happen occasionally.  Even with built-in redundancy, load shedding, and good system design practises, the reality is that data center power supplies, and their backup systems do sometimes fail at the same time. Software rollouts, reconfigurations, system overloads and other cascading failures do occasionally go wrong at cluster-wide scale.  When they do (not if they do), it's better not to lose all (or even half of) your total capacity.
2. **Scale for the sake of scale can lead to unnecessarily complex system designs**:    Many relatively simple system designs have large, but not infinite, natural scaling limits, and in order to go beyond those limits, more complex designs become necessary.  Unless the need for the additional scale can be justified, the additional complexity often cannot be.

Where Google's internal requirements for compute capacity exceed the natural scaling limits of a cluster (or even multiple clusters, required for redundancy or geo-locality), we typically choose to build more clusters, rather than trying to build larger clusters.  While the resultant increase in the number of clusters can be seen as an unnecessary administrative burden, we have found that the benefits of this approach far outweigh this disadvantage. In addition we have built internal automated tooling which helps us to more easily manage larger numbers of clusters, and accommodate applications that span multiple clusters.

For Kubernetes and GKE, "Ubernetes" Cluster Federation will become our primary toolset in that regard.  Ubernetes is currently designed to federate up to 100 clusters together, to create an aggregate capacity of 100 times the maximum cluster capacity described in this document.

With the above considerations in mind, we define the following medium-term scaling and performance goals for individual Kubernetes clusters, with specific justifications where they may not be self-evident.

NOTES:
1. This document does not cover minimum scaling limits (e.g. minimum nodes per cluster, or RAM per node etc) - that topic is considered out of scope here.
2. These are not commitments that we plan to make for any specific release version (e.g. 1.3).  Rather, they represent what we believe to be reasonable short to medium-term scaling and performance objectives, absent any concrete and compelling use cases that suggest otherwise.

## Scaling and Performance Goals

### Primary Metrics

* Max cores per cluster
  * Goal: 200,000
  * See Background section above
* Max pods per core
  * Goal: 10
  * Assume RAM:core ratios of between 1:1 and 4:1 (GBytes/core).  Justification:
    1. Cores are the primary things that are shared. If too many pods (actually processes) try to share the same core, then the scheduling latency experienced by a given pod becomes too high. If, when a pod unblocks and becomes runnable it has to wait for too many other pods to block or be pre-empted, it becomes a lot less useful, particularly for pods in a serving path (or even a latency sensitive batch pipeline, like the ones that perform daily billing runs, hourly stock analyses, weather forecasts or other time-sensitive data processing).
    2. RAM: Practical experience has shown that on average many pods/processes require between 1 and 4 GB RAM per core. That fact is reflected in common commercial machine configurations and cloud provider offerings. The ratio does not change very fast over time (it has been roughly constant for 20 years). On average, useful processes require at least a few tens or hundreds of MB of RAM to do anything useful.
    3. The number of cores per machine changes dramatically over time. In 2005, mainstream servers had 2 cores, in 2010 they had 16 cores, in 2016 they have 48 cores or more. We don't want to have to adjust our pod density targets every year to keep up with that. Hence pods per core is a much more stable ratio to pursue than e.g. pods per machine.
* Management overhead per node
  * Goal: <5%, with a minimum of 0.5 core, 1GB RAM
  * includes Docker, KubeProxy, Kubelet, metrics-gathering, excludes kernel. (e.g. implies on 32 core, 64GB machine, ~1.6 cores, 3.2GB RAM for node management).
* Management overhead per cluster
  * Goal: <1%, with a minimum of 2 cores, 4GB RAM
  * Includes all non-node specific components including apiserver, scheduler, controllers, etcd, DNS, heapster etc, excludes HA; <3% for HA. This is in addition to the per node management overhead.

### Derived Metrics/Caveats

* Max cores per node
  * Goal: 64 now, hundreds in future
  * We plan to support the highest core count GCE instances, AWS instances and mainstream server platforms now and in the future.
* Max pods per machine
  * Goal: 500
  * This is marginally less than (max cores per node) x (max pods per core).  We might not, for example, in future support 256 cores per node times 10 pods per core.  But users can achieve max pods per core on today's  moderately large (e.g. 48 core, 96 GB RAM) machines.
* Max machines per cluster
  * Goal: 5,000
  * This is somewhat greater than (max cores per cluster) / (max cores per node), but not excessively so.  The primary metric is cores/cluster, but we do not believe that it is worth supporting the maximum core count per cluster on exclusively small (e.g. only single-core) machines. Conversely, we do not expect users to only be able to achieve the max cores per cluster limit by using the largest commercially available machines.  Using current numbers however, we plan to support clusters of e.g. 5,000 x 32 core machines = 160,000 cores.
* Max pods per cluster
  * Goal: 500,000
  * Note: this is somewhat less than (max cores per cluster) x (max pods per core).  I.e. The largest possible cluster, packed exclusively with the smallest possible pods, is not supported, as we have not (yet) found a good use case for this (see also Open Questions below).
* End-to-end pod startup time
  * Goal: <= 5 seconds (99th percentile)
  * With pre-pulled images, saturated cluster, no scheduling backlog.
* Scheduler throughput
  * Goal: 100 pods per second
  * Driven primarily by max pods per cluster, and max cluster saturation time (see below).
* Max cluster saturation time
  * Goal: 90 minutes
  * Time to restore a fully saturated large cluster is important for cluster-wide failure recovery, and/or related emergency capacity provisioning (e.g. building and populating a new cluster to replace capacity in a failed one). This number also needs to correlate with max pods per cluster, and max scheduler throughput (500,000 pods / 100 pods per second ~ 90 minutes).  We believe that this fulfills most real-world recovery requirements.  The required time to recovery is usually driven primarily by trying to reduce the probability of multiple uncorrelated cluster failures (e.g. "one of our 3 clusters has failed. We're just fine unless another one fails before we've repaired/replaced the first failed one").

## Control Plane Configurations for Testing

Configuration of the control plane for cluster testing varies by provider, and there are multiple reasonable configurations. Discussion and guideline of control plane configuration options and standards are documented [here](configs-and-limits/provider-configs.md).

## Open Questions

1. **What, if any, reasonable use cases exist for very large numbers of very small nodes (e.g. for isolation reasons - multitenant)?  Based on comments so far, it seems that the answer is yes, and needs to be addressed.**<br>
The above scaling goals explicitly accommodate a maximum of 5,000 nodes.  Do we need a special case for larger numbers of small nodes (e.g. 200,000  single-core nodes).  The latter already fits within our other scalability limits (cores per cluster, pods per cluster), so it might not be more difficult to achieve than those.  Two example use cases I've heard anecdotally are (a) running e.g. large numbers of customers' small, largely idle wordpress instances, one per node, and (b) giving away limited numbers of free containers to large numbers of customers for promotional reasons (similar to AWS free tier micro instances).
2. **What, if any, reasonable use cases exist for very large numbers of very small containers per core?**<br>
E.g. are perhaps hundreds of containers per core useful for specialized applications? We speculate that building large numbers of very small yet useful containers each within say 20MB of RAM, and <1% of a core is difficult (as opposed to very small supportive/sidecar pods alongside larger pods, which is totally legitimate, and supported).  
