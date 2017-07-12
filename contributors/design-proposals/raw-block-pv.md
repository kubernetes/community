# [WIP][Proposal] Local Raw Block Consumption via Persistent Volume Source

Authors: erinboyd@, screeley44@, mtanino@

This document presents a proposal for managing raw block storage in Kubernetes using the persistent volume soruce API as a consistent model
of consumption.

# Goals# [Proposal] Local Raw Block Consumption via Persistent Volume Source

Authors: erinboyd@, screeley44@, mtanino@

This document presents a proposal for managing raw block storage in Kubernetes using the persistent volume soruce API as a consistent model
of consumption.

# Goals
* Enable durable access to block storage
* Support storage requirements for all workloads supported by Kubernetes
* Provide flexibility for users/vendors to utilize various types of storage devices
* Agree on API changes for block
* Provide a consistent security model for block devices 
* Provide block storage usage isolation
* Provide a means for running 
* Enable durable access to block storage
* Support storage requirements for all workloads supported by Kubernetes
* Provide flexibility for users/vendors to utilize various types of storage devices
* Agree on API changes for block
* Provide a consistent security model for block devices 
* Provide block storage usage isolation
* Provide a means for running containerized block storage offerings as non-privileged container

# Non Goals
* Support all storage devices natively in upstream Kubernetes. Non standard storage devices are expected to be managed using extension mechanisms.
   
# Value add to Kubernetes

  Before the advent of storage plugins, emptyDir and hostPath were widely used to quickly prototype stateless applications in Kube. 
  Both have limitations for their use in application that need to store persistent data or state. 
  EmptyDir, though quick and easy to use, provides no real guarantee of persistence for suitable amount of time. 
  Appropriately used as scratch space, it does not have the HostPath, became an initial offering for local storage, but had many drawbacks. Without having the ability to guarantee space & ensure ownership, one would lose data once a node was rescheduled. Therefore, the risk outweighed the reward when trying to  leverage the power of local storage needed for stateful applications like databases.
  Adding API changes to volume API (limits the use of raw block devices outside of storage)
  Current API proposal allow for overloading API with non-storage use cases (GPU) with PVC withrequested filesystem on top.
  Distributed filesystems and databases are the primary use cases for persistent local storage to fully utilize the performace of SSDs 
  and can in non-cloud environments provide a more cost effective solution.
    
# Design Overview

[TBD]
