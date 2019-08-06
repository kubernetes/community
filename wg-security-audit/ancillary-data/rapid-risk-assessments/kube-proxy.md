# Overview

- Component: kube-proxy
- Owner(s): [sig-network](https://github.com/kubernetes/community/tree/master/sig-network)
- SIG/WG(s) at meeting:
- Service Data Classification: Medium
- Highest Risk Impact:

# Service Notes

The portion should walk through the component and discuss connections, their relevant controls, and generally lay out how the component serves its relevant function. For example
a component that accepts an HTTP connection may have relevant questions about channel security (TLS and Cryptography), authentication, authorization, non-repudiation/auditing,
and logging. The questions aren't the *only* drivers as to what may be spoken about, the questions are meant to drive what we discuss and keep things on task for the duration
of a meeting/call.

## How does the service work?

- kubeproxy has several main modes of operation:
  - as a literal network proxy, handling networking between nodes
  - as a bridge between Container Network Interface (CNI) which handles the actual networking and the host operating system
    - `iptables` mode
    - `ipvs` mode
    - two Microsoft Windows-specific modes (not covered by the RRA)
- in any of these modes, kubeproxy interfaces with the host's routing table so as to achieve a seamless, flat network across the kubernetes cluster

## Are there any subcomponents or shared boundaries?

Yes.

- Similar to kubelet, kube-proxy run's on the node, with an implicit trust boundary between Worker components and Container components (i.e. pods)

## What communications protocols does it use?

- Direct IPC to `iptables` or `ipvs`
- HTTPS to the kube-apiserver
- HTTP Healthz port (which is a literal counter plus a `200 Ok` response)

## Where does it store data?

Minimal data should be stored by kube-proxy itself, this should mainly be handled by kubelet and some file system configuration

## What is the most sensitive data it stores?

N/A

## How is that data stored?

N/A

# Data Dictionary

| Name | Classification/Sensitivity | Comments |
| :--: | :--: | :--: |
| Data | Goes | Here |

# Control Families 

These are the areas of controls that we're interested in based on what the audit working group selected. 

When we say "controls," we mean a logical section of an application or system that handles a security requirement. Per CNSSI:

> The management, operational, and technical controls (i.e., safeguards or countermeasures) prescribed for an information system to protect the confidentiality, integrity, and availability of the system and its information.

For example, an system may have authorization requirements that say:

- users must be registered with a central authority
- all requests must be verified to be owned by the requesting user
- each account must have attributes associated with it to uniquely identify the user

and so on. 

For this assessment, we're looking at six basic control families:

- Networking
- Cryptography
- Secrets Management
- Authentication
- Authorization (Access Control)
- Multi-tenancy Isolation

Obviously we can skip control families as "not applicable" in the event that the component does not require it. For example,
something with the sole purpose of interacting with the local file system may have no meaningful Networking component; this
isn't a weakness, it's simply "not applicable."

For each control family we want to ask:

- What does the component do for this control?
- What sorts of data passes through that control? 
  - for example, a component may have sensitive data (Secrets Management), but that data never leaves the component's storage via Networking
- What can attacker do with access to this component?
- What's the simplest attack against it?
- Are there mitigations that we recommend (i.e. "Always use an interstitial firewall")?
- What happens if the component stops working (via DoS or other means)?
- Have there been similar vulnerabilities in the past? What were the mitigations?

# Threat Scenarios

- An External Attacker without access to the client application
- An External Attacker with valid access to the client application
- An Internal Attacker with access to cluster
- A Malicious Internal User

## Networking

- kube-proxy is actually five programs
- proxy: mostly deprecated, but a literal proxy, in that it intercepts requests and proxies them to backend services
- IPVS/iptables: very similar modes, handle connecting virtual IPs (VIPs) and the like via low-level routing (the preferred mode)
- two Windows-specific modes (out of scope for this discussion, but if there are details we can certainly add them)

Node ports:

- captures traffic from Host IP
- shuffles to backend (used for building load balancers)

- kube-proxy shells out to `iptables` or `ipvs`
- Also uses a netlink socket for IPVS (netlink are similar to Unix Domain Sockets)
- *Also* shells out to `ipset` under certain circumstances for IPVS (building sets of IPs and such) 


### User space proxy

Setup:

1. Connect to the kube-apiserver
1. Watch the API server for services/endpoints/&c
1. Build in-memory caching map: for services, for every port a service maps, open a port, write iptables rule for VIP & Virt Port
1. Watch for updates of services/endpoints/&c

when a consumer connects to the port:

1. Service is running VIP:VPort
1. Root NS -> iptable -> kube-proxy port
1. look at the src/dst port, check the map, pick a service on that port at random (if that fails, try another until either success or a retry count has exceeded)
1. Shuffle bytes back and forth between backend service and client until termination or failure

### iptables

1. Same initial setup (sans opening a port directly)
1. iptables restore command set
1. giant string of services
1. User VIP -> Random Backend -> Rewrite packets (at the kernel level, so kube-proxy never sees the data)
1. At the end of the sync loop, write (write in batches to avoid iptables contentions)
1. no more routing table touches until service updates (from watching kube-apiserver or a time out, expanded below)

**NOTE**: rate limited (bounded frequency) updates:
- no later than 10 minutes by default
- no sooner than 15s by default (if there are no service map updates)

this point came out of the following question: is having access to kube-proxy *worse* than having root access to the host machine?

### ipvs

1. Same setup as iptables & proxy mode
1. `ipvsadm` and `ipset` commands instead of `iptables`
1. This does have some strange changes:
  - ip address needs a dummy adapter
  - !NOTE Any service bound to 0.0.0.0 are also bound to _all_ adapters
    - somewhat expected because 0.0.0.0, but can still lead to interesting behavior

### concern points within networking

- !NOTE: ARP table attacks (such as if someone has `CAP_NET_RAW` in a container or host access) can impact kube-proxy
- Endpoint selection is namespace & pod-based, so injection could overwrite (I don't think this is worth a finding/note because kube-apiserver is the arbiter of truth)
- !FINDING (but low...): POD IP Reuse: (factor of 2 x max) cause a machine to churn thru IPS, you could cause a kube-proxy to forward ports to your pod if you win the race condition.
  - this would be limited to the window of routing updates
  - however, established connections would remain
  - kube-apiserver could be the arbiter of routing, but that may require more watch and connection to the central component
  - [editor] I think just noting this potential issue and maybe warning on it in kube-proxy logs would be enough 

### with root access?

Access to kube-proxy is mostly the same as root access

- set syscalls, route local, &c could gobble memory
- Node/VIP level
- Recommend `CAP_NET_BIND` (bind to low ports, don't need root for certain users) for containers/pods, alleviate concerns there
- Can map low ports to high ports in kube-proxy as well, but mucks with anything that pretends to be a VIP
  - LB forwards packets to service without new connection (based on srcport)
  - 2-hop LB, can't do direct LB 

## Cryptography

- kube-proxy itself does not handle cryptography other than the TLS connection to kube-apiserver

## Secrets Management

- kube-proxy itself does not handle secrets, but rather only consumes credentials from the command line (like all other k8s components)

## Authentication

- kube-proxy does not handle any authentication other than credentials to the kube-apiserver

## Authorization

- kube-proxy does not handle any authorization; the arbiters of authorization are kubelet and kube-proxy

## Multi-tenancy Isolation

- kube-proxy does not currently segment clients from one another, as clients on the same pod/host must use the same iptables/ipvs configuration
- kube-proxy does have conception of namespaces, but currently avoids enforcing much at that level
  - routes still must be added to iptables or the like
  - iptables contention could be problematic
  - much better to handle at higher-level components, namely kube-apiserver and kube-proxy

## Logging

- stderr directed to a file
- same as with kubelet
- !FINDING (but same as all other components) logs namespaces, service names (same as every other service)

# Additional Notes

## kubelet to iptables

- per pod network management
- pods can request a host port, docker style 
- kubenet and CNI plugins
- kubenet uses CNI
- setup kubenet iptable to map ports to a single pod
- overly broad, should be appended to iptables list
- all local IPs to the host 

!FINDING: don't use host ports, they can cause problems with services and such; we may recommend deprecating them

## Summary

# Recommendations
