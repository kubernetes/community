# Provide Metrics to Volume Controllers

## Background

Both Persistent Volume Controller and Attach-Detach Volume Controller call out Cloud and storage providers APIs to provision, delete, attach, detach volumes. As observed, latency from such API calls have usability issues. 

Metrics used by [Kubernetes resource monitoring tool](https://kubernetes.io/docs/user-guide/monitoring/) is the mechanism to instrument Kubernetes components. Providing metrics to volume controllers helps to identify and reduce latencies.

## Volume Actions to track

###  Persistent Volume Controller

* Provision

Contrller interacts with Cloud or storage providers to create a volume that matches the claim specification. In many cases, PVs must be created dynamically in order for Pod to start. Provision latencies have direct impact on Pod startup.


###  Attach and Detach Controller

* Attach

Similar to PV Provision, Volume Attach contributes to Pod startup latency. Certain platforms observed attach latency in minutes. 


## Metrics type

Histogram metrics, starting from 100 milliseconds, are used to measure these latencies.
