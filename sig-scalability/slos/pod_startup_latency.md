## Pod startup latency SLI/SLO details

### Definition

| Status | SLI | SLO |
| --- | --- | --- |
| __Official__ | Startup latency of schedulable<sup>[1](#footnote1)</sup> stateless<sup>[2](#footnote2)</sup> pods, excluding time to pull images and run init containers, measured from pod creation timestamp to when all its containers are reported as started and observed via watch, measured as 99th percentile over last 5 minutes | In default Kubernetes installation, 99th percentile per cluster-day <= 5s |
| __WIP__ | Startup latency of schedulable<sup>[1](#footnote1)</sup> stateful<sup>[3](#footnote3)</sup> pods, excluding time to pull images, run init containers, provision volumes (in delayed binding mode) and unmount/detach volumes (from previous pod if needed), measured from pod creation timestamp to when all its containers are reported as started and observed via watch, measured as 99th percentile over last 5 minutes | In default Kubernetes installation, 99th percentile per cluster-day <= X where X depends on storage provider |

<a name="footnote1">[1\]</a>By schedulable pod we mean a pod that can be
scheduled in the cluster without causing any preemption.

<a name="footnote2">[2\]</a>A `stateless pod` is defined as a pod that doesn't
mount volumes with sources other than secrets, config maps, downward API and
empty dir.

<a name="footnote3">[3\]</a>A `stateful pod` is defined as a pod that mounts
at least one volume with sources other than secrets, config maps, downward API
and empty dir.

### User stories
- As a user of vanilla Kubernetes, I want some guarantee how quickly my pods
will be started.

### Other notes
- Only schedulable pods contribute to the SLIs:
  - If there is no space in the cluster to place the pod, there is not much
    we can do about it (it is task for Cluster Autoscaler which should have
    separate SLIs/SLOs).
  - If placing a pod requires preempting other pods, that may heavily depend
    on the application (e.g. on their graceful termination period). We don't
    want that to contribute to this SLI.
- We are explicitly splitting stateless and stateful pods from each other:
  - Starting a stateful pod requires attaching and mounting volumes, that
    takes non-negligible amount of time and doesn't even fully depend on
    Kubernetes. However, even though it depends on chosen storage provider,
    it isn't application specific, thus we make that part of the SLI
    (though the exact SLO threshold may depend on chosen storage provider).
  - We also explicitly exclude time to provision a volume (in delayed volume
    binding mode), even though that also only depends on storage provider,
    not on the application itself. However, volume provisioning can be
    perceived as a bootstrapping operation in the lifetime of a stateful
    workload, being done only once at the beginning, not everytime a pod is
    created. As a result, we decided to exclude it.
  - We also explicitly exclude time to unmount and detach the volume (if it
    was previously mounted to a different pod). This situation is symetric to
    excluding pods that need to preempt others (it's kind of cleaning after
    predecessors).
- We are explicitly excluding image pulling from time the SLI. This is
because it highly depends on locality of the image, image registry performance
characteristic (e.g. throughput), image size itself, etc. Since we have
no control over any of those (and all of those would significantly affect SLI)
we decided to simply exclude it.
- We are also explicitly excluding time to run init containers, as, again, this
is heavily application-dependent (and does't depend on Kubernetes itself).
- The answer to question "when pod should be considered as started" is also
not obvious. We decided for the semantic of "when all its containers are
reported as started and observed via watch", because:
  - we require all containers to be started (not e.g. the first one) to ensure
    that the pod is started. We need to ensure that potential regressions like
    linearization of container startups within a pod will be catch by this SLI.
  - note that we don't require all container to be running - if some of them
    finished before the last one was started that is also fine. It is just
    required that all of them has been started (at least once).
  - we don't want to rely on "readiness checks", because they heavily
    depend on the application. If the application takes couple minutes to
    initialize before it starts responding to readiness checks, that shouldn't
		count towards Kubernetes performance.
  - even if your application started, many control loops in Kubernetes will
    not fire before they will observe that. If Kubelet is not able to report
    the status due to some reason, other parts of the system will not have
    a way to learn about it - this is why reporting part is so important
    here.
  - since watch is so centric to Kubernetes (and many control loops are
    triggered by specific watch events), observing the status of pod is
    also part of the SLI (as this is the moment when next control loops
    can potentially be fired).

### TODOs
- Revisit whether we want "watch pod status" part to be included in the SLI.
- Consider creating an SLI for pod deletion latency, given that for stateful
pods, detaching RWO volume is required before it can be attached to a
different node (i.e. slow pod deletion may block pod startup).
- While it's easy to exclude pods that require preempting other pods or
volume unmounting in tests (where we fully control the environment), we need
to figure out how to do that properly in production environments.

### Test scenario

__TODO: Descibe test scenario.__

Note: when running tests against clusters with nodes in multiple zones, the
preprovisioned volumes should be balanced across zones so that we don't make
pods unschedulable due to lack of resources in a single zone.
