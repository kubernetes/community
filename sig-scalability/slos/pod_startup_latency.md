## Pod startup latency SLI/SLO details

### User stories
- As a user of vanilla Kubernetes, I want some guarantee how quickly my pods
will be started.

### Other notes
- Only schedulable and stateless pods contribute to the SLI:
  - If there is no space in the cluster to place the pod, there is not much
    we can do about it (it is task for Cluster Autoscaler which should have
    separate SLIs/SLOs).
  - If placing a pod requires preempting other pods, that may heavily depend
    on the application (e.g. on their graceful termination period). We don't
    want that to contribute to this SLI.
  - Mounting disks required by non-stateless pods may potentially also require
    non-negligible time, not fully dependent on Kubernetes.
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
    that the pod is started. We need to ensure that pontential regressions like
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
- We should try to provide guarantees for non-stateless pods (the threshold
may be higher for them though).
- Revisit whether we want "watch pod status" part to be included in the SLI.

### Test scenario

__TODO: Descibe test scenario.__
