# Networking
**Lead:** thockin  
**Slides:** [here](https://docs.google.com/presentation/d/1Qb2fbyTClpl-_DYJtNSReIllhetlOSxFWYei4Zt0qFU/edit#slide=id.g2264d16f0b_0_14)  
**Thanks to our notetakers:** onyiny-ang, mrbobbytales, tpepper


This session is not declaring what's being implemented next, but rather laying out the problems that loom.

## Coming soon
 - kube-proxy with IPVS
     - currently beta
 - core DNS replacing kube DNS
     - currently beta
 - pod "ready++"
     - allow external systems to participate in rolling updates. Say your load-balancer takes 5-10 seconds to program, when you bring up new pod and take down old pod the load balancer has lost old backends but hasn't yet added new backends. The external dependency like this becomes a gating pod decorator.
     - adds configuration to pod to easily verify readiness
     - design agreed upon, alpha (maybe) in 1.11

## Ingress
* The lowest common-denominator API. This is really limiting for users, especially compared to modern software L7 proxies.
* annotation model of markup limits portability
* ingress survey reports:
  * people want portability
  * everyone uses non-portable features…
  * 2018 L7 requirements are dramatically higher than what they were and many vendors don’t support that level of functionality.
* Possible Solution? Routes
  * openshift uses routes
  * heptio prototyping routes currently
* All things considered, requirements are driving it closer and closer to istio
Possibility, poach some of the ideas and add them to kubernetes native.

## Istio
(as a potential solution)
- maturing rapidly with good APIs and support
- Given that plus istio is not part of kubernetes, it's unlikely near term to become a default or required part of a k8s deployment. The general ideas around istio style service mesh could be more native in k8s.

## Topology and node-local Services
- demand for node-local network and service discovery but how to go about it?
    - e.g. “I want to talk to the logging daemon on my current host”
    - special-case topology?
    - client-side choice
- These types of services should not be a service proper.

## Multi-network

- certain scenarios demand multi-network
- A pod can be in multiple networks at once. You might have different quality of service on different networks (eg: fast/expensive, slower/cheaper), or different connectivity (eg: the rack-internal network).
- Tackling scenarios like NFV
- need deeper changes like multiple pod IPs but also need to avoid repeating old mistakes
- SIG-Network WG designing a PoC -- If interested jump on SIG-network WG weekly call
- Q: Would this PoC help if virtual-kubelets were used to span cloud providers? Spanning latency domains in networks is also complicated. Many parts of k8s are chatty, assuming a cluster internal low-latency connectivity.

## Net Plugins vs Device Plugins
- These plugins do not coordinate today and are difficult to work around
- gpu that is also an infiniband device
- causes problems because network and device are very different with verbs etc
- problems encountered with having to schedule devices and network together at the same time.
“I want a gpu on this host that has a gpu attached and I want it to be the same deviec”
PoC available to make this work, but its rough and a problem right now.
- Resources WG and networking SIG are discussing this challenging problem
- SIGs/WGs. Conversation may feel like a cycle, but @thockin feels it is a spiral that is slowly converging and he has a doc he can share covering the evolving thinking.

## Net Plugins, gRPC, Services
- tighter coupling between netplugins and kube-proxy could be useful
- grpc is awesome for plugins, why not use a grpc network plugin
- pass services to network plugin to bypass kube-proxy, give more awareness to the network plugin and enable more functionality.

## IPv6
- beta but **no** support for dual-stack (v4 & v6 at the same time)
- Need deeper changes like multiple pod IPs (need to change the pod API--see Multi-network)
- https://github.com/kubernetes/features/issues/563

## Services v3

- Services + Endpoints have a grab-bag of features which is not ideal; "grew organically"
- Need to start segmenting the "core" API group
    - write API in a way that is more obvious
    - split things out and reflect it in API
- Opportunity to rethink and refactor:
    - Endpoints -> Endpoint?
    - split the grouping construct from the “gazintas”
      - virtualIP, network, dns name moves into the service
    - EOL troublesome features
      - port remapping

## DNS Reboot
- We abuse DNS and mess up our DNS schema
    - it's possible to write queries in DNS that take over names
    - @thockin has a doc with more information about the details of this
    - Why can't I use more than 6 web domains? bugzilla circa 1996
- problem: its possible to write queries in dns that write over names
  - create a namespace called “com” and an app named “google” and it’ll cause a problem
- “svc” is an artifact and should not be a part of dns
- issues with certain underlying libraries
- Changing it is hard (if we care about compatibility)
- Can we fix DNS spec or use "enlightened" DNS servers
  - Smart proxies on behalf of pods that do the searching and become a “better” dns
- External DNS
- Creates DNS entries in external system (route53)
- Currently in incubator, not sure on status, possibly might move out of incubator, but unsure on path forward

Perf and Scalability  
- iptables is krufty. nftables implementation should be better.
- ebpf implementation (eg; Cilium) has potential

## Questions:

- Consistent mechanism to continue progress but maintain backwards compatibility
- External DNS was not mentioned -- blue/green traffic switching
    - synchronizes kubernetes resources into various Kubernetes services
    - it's in incubator right now (deprecated)
    - unsure of the future trajectory
    - widely used in production
    - relies sometimes on annotations and ingress
- Q: Device plugins. . .spiraling around and hoping for eventual convergence/simplification
  - A: Resource management on device/net plugin, feels like things are going in a spiral, but progress is being made, it is a very difficult problem and hard to keep all design points tracked. Trying to come to consensus on it all.
- Q: Would CoreDNS be the best place for the plugins and other modes for DNS proxy etc.
    - loss of packets are a problem -- long tail of latency
    - encourage cloud providers to support gRPC
- Q: With the issues talked about earlier, why can’t istio be integrated natively?
  - A: Istio can't be required/default: still green
    - today we can't proclaim that Kubernetes must support Istio
    - probably not enough community support this year (not everyone is using it at this point)
- Q: Thoughts on k8s v2?
  - A: Things will not just be turned off, things must be phased out and over the course of years, especially for services which have been core for some time.

## Take Aways:
- This is not a comprehensive list of everything that is up and coming
- A lot of work went into all of these projects
